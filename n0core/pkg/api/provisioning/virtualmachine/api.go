package virtualmachine

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/koding/websocketproxy"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/n0stack/n0stack/n0core/pkg/api/pool/node"
	"github.com/n0stack/n0stack/n0core/pkg/api/provisioning"
	"github.com/n0stack/n0stack/n0core/pkg/api/provisioning/block_storage"
	"github.com/n0stack/n0stack/n0core/pkg/datastore"
	"github.com/n0stack/n0stack/n0core/pkg/util/grpc"
	"github.com/n0stack/n0stack/n0core/pkg/util/net"
	"github.com/n0stack/n0stack/n0proto.go/pkg/transaction"
	"github.com/n0stack/n0stack/n0proto.go/pool/v0"
	"github.com/n0stack/n0stack/n0proto.go/provisioning/v0"
)

var N0coreVirtualMachineNamespace uuid.UUID

func init() {
	N0coreVirtualMachineNamespace, _ = uuid.FromString("a015d18d-b2c3-4181-8028-6f707ef31c95")
}

type VirtualMachineAPI struct {
	dataStore datastore.Datastore

	// dependency APIs
	nodeAPI         ppool.NodeServiceClient
	networkAPI      ppool.NetworkServiceClient
	blockstorageAPI pprovisioning.BlockStorageServiceClient
}

func CreateVirtualMachineAPI(ds datastore.Datastore, noa ppool.NodeServiceClient, nea ppool.NetworkServiceClient, bsa pprovisioning.BlockStorageServiceClient) *VirtualMachineAPI {
	a := &VirtualMachineAPI{
		dataStore:       ds.AddPrefix("virtual_machine"),
		nodeAPI:         noa,
		networkAPI:      nea,
		blockstorageAPI: bsa,
	}

	return a
}

func (a *VirtualMachineAPI) addDefaultGateway(ctx context.Context, network *ppool.Network) (string, error) {
	ipn := netutil.ParseCIDR(network.Ipv4Cidr)
	ip := netutil.GetEndIP(ipn.Network())

	a.networkAPI.ReserveNetworkInterface(ctx, &ppool.ReserveNetworkInterfaceRequest{
		NetworkName:          network.Name,
		NetworkInterfaceName: "default-gateway",
		Ipv4Address:          ip.String(),
		Annotations: map[string]string{
			provisioning.AnnotationVirtualMachineVncWebSocketPort: "true",
		},
	})

	return ip.String(), nil
}

// PENDINGステートにすることで楽観的なロックを行う
func (a *VirtualMachineAPI) lockOptimistically(vm *pprovisioning.VirtualMachine) (func() error, error) {
	// PENDINGステートにすることで楽観的なロックを行う
	if vm.State == pprovisioning.VirtualMachine_PENDING {
		return nil, grpcutil.WrapGrpcErrorf(codes.ResourceExhausted, "State is PENDING, so cannnot do any actions") // これで State がいいのか自信ない
	}

	vm.State = pprovisioning.VirtualMachine_PENDING
	if err := a.dataStore.Apply(vm.Name, vm); err != nil {
		return nil, grpcutil.WrapGrpcErrorf(codes.Internal, "Failed to apply data for db: err='%s'", err.Error())
	}

	f := func() error {
		if vm.State == pprovisioning.VirtualMachine_PENDING {
			vm.State = pprovisioning.VirtualMachine_UNKNOWN
		}

		return a.dataStore.Apply(vm.Name, vm)
	}

	return f, nil
}

func (a *VirtualMachineAPI) CreateVirtualMachine(ctx context.Context, req *pprovisioning.CreateVirtualMachineRequest) (*pprovisioning.VirtualMachine, error) {
	// validation
	var id uuid.UUID
	{
		switch {
		case req.Name == "":
			return nil, grpcutil.WrapGrpcErrorf(codes.InvalidArgument, "Set name")

		case req.LimitCpuMilliCore%1000 != 0:
			return nil, grpcutil.WrapGrpcErrorf(codes.InvalidArgument, "Make limit_cpu_milli_core '%d' a multiple of 1000", req.LimitCpuMilliCore)

		case req.RequestCpuMilliCore == 0 || req.RequestMemoryBytes == 0:
			return nil, grpcutil.WrapGrpcErrorf(codes.InvalidArgument, "Set request_*")
		}

		var err error
		if req.Uuid == "" {
			id = uuid.NewV5(N0coreVirtualMachineNamespace, req.Name)
		} else if id, err = uuid.FromString(req.Uuid); err != nil {
			return nil, grpcutil.WrapGrpcErrorf(codes.InvalidArgument, "Set valid uuid")
		}
	}

	tx := transaction.Begin()
	defer transaction.WrapRollbackError(tx.Rollback())

	vm := &pprovisioning.VirtualMachine{
		Name:        req.Name,
		Annotations: req.Annotations,
		Uuid:        id.String(),

		RequestCpuMilliCore: req.RequestCpuMilliCore,
		LimitCpuMilliCore:   req.LimitCpuMilliCore,
		RequestMemoryBytes:  req.RequestMemoryBytes,
		LimitMemoryBytes:    req.LimitMemoryBytes,

		BlockStorageNames: req.BlockStorageNames,
		Nics:              req.Nics,

		LoginUsername:     req.LoginUsername,
		SshAuthorizedKeys: req.SshAuthorizedKeys,
	}
	if vm.Annotations == nil {
		vm.Annotations = make(map[string]string)
	}

	{
		prev := &pprovisioning.VirtualMachine{}
		if err := a.dataStore.Get(vm.Name, prev); err != nil {
			return nil, grpcutil.WrapGrpcErrorf(codes.Internal, "Failed to get data for db: err='%s'", err.Error())
		} else if prev.Name != "" {
			return nil, grpcutil.WrapGrpcErrorf(codes.AlreadyExists, "BlockStorage '%s' is already exists", vm.Name)
		}

		unlock, err := a.lockOptimistically(vm)
		if err != nil {
			return nil, err
		}
		tx.PushRollback("free optimistic lock", unlock)
	}

	{
		vm.ComputeName = vm.Name

		var err error
		var n *ppool.Node
		if node, ok := vm.Annotations[AnnotationVirtualMachineRequestNodeName]; !ok {
			n, err = a.nodeAPI.ScheduleCompute(ctx, &ppool.ScheduleComputeRequest{
				ComputeName: vm.ComputeName,
				Annotations: map[string]string{
					AnnotationComputeReservedBy: vm.Name,
				},
				RequestCpuMilliCore: vm.RequestCpuMilliCore,
				LimitCpuMilliCore:   vm.LimitCpuMilliCore,
				RequestMemoryBytes:  vm.RequestMemoryBytes,
				LimitMemoryBytes:    vm.LimitMemoryBytes,
			})
			if err != nil {
				return nil, grpcutil.WrapGrpcErrorf(grpc.Code(err), "Failed to ScheduleCompute: desc=%s", grpc.ErrorDesc(err))
			}
		} else {
			n, err = a.nodeAPI.ReserveCompute(ctx, &ppool.ReserveComputeRequest{
				NodeName:    node,
				ComputeName: vm.ComputeName,
				Annotations: map[string]string{
					AnnotationComputeReservedBy: vm.Name,
				},
				RequestCpuMilliCore: vm.RequestCpuMilliCore,
				LimitCpuMilliCore:   vm.LimitCpuMilliCore,
				RequestMemoryBytes:  vm.RequestMemoryBytes,
				LimitMemoryBytes:    vm.LimitMemoryBytes,
			})
			if err != nil {
				return nil, grpcutil.WrapGrpcErrorf(grpc.Code(err), "Failed to ReserveCompute: desc=%s", grpc.ErrorDesc(err))
			}
		}

		vm.ComputeNodeName = n.Name

		tx.PushRollback(fmt.Sprintf("ReleaseCompute '%s'", vm.Name), func() error {
			_, err := a.nodeAPI.ReleaseCompute(ctx, &ppool.ReleaseComputeRequest{
				NodeName:    vm.ComputeNodeName,
				ComputeName: vm.ComputeName,
			})
			return err
		})
	}

	blockdevs := make([]*BlockDev, len(vm.BlockStorageNames))
	{
		for i, n := range vm.BlockStorageNames {
			v, err := a.blockstorageAPI.SetInuseBlockStorage(ctx, &pprovisioning.SetInuseBlockStorageRequest{Name: n})
			if err != nil {
				return nil, grpcutil.WrapGrpcErrorf(grpc.Code(err), "Failed to SetInuseBlockStorage: desc=%s", grpc.ErrorDesc(err))
			}
			tx.PushRollback(fmt.Sprintf("SetAvailableBlockStorage '%s'", n), func() error {
				_, err := a.blockstorageAPI.SetAvailableBlockStorage(ctx, &pprovisioning.SetAvailableBlockStorageRequest{Name: n})
				return err
			})

			blockdevs[i] = &BlockDev{
				Name:      n,
				Url:       v.Annotations[blockstorage.AnnotationBlockStorageURL],
				BootIndex: uint32(i),
			}
		}
	}

	netdevs := make([]*NetDev, len(vm.Nics))
	{
		vm.NetworkInterfaceNames = make([]string, len(vm.Nics))

		for i, nic := range vm.Nics {
			vm.NetworkInterfaceNames[i] = vm.Name + strconv.Itoa(i)

			network, err := a.networkAPI.ReserveNetworkInterface(ctx, &ppool.ReserveNetworkInterfaceRequest{
				NetworkName:          nic.NetworkName,
				NetworkInterfaceName: vm.NetworkInterfaceNames[i],
				Annotations: map[string]string{
					AnnotationComputeReservedBy: vm.Name,
				},
				HardwareAddress: nic.HardwareAddress,
				Ipv4Address:     nic.Ipv4Address,
				Ipv6Address:     nic.Ipv6Address,
			})
			if err != nil {
				return nil, grpcutil.WrapGrpcErrorf(grpc.Code(err), "Failed to ReserveNetworkInterface: desc=%s", grpc.ErrorDesc(err))
			}
			tx.PushRollback("ReleaseNetworkInterface", func() error {
				_, err := a.networkAPI.ReleaseNetworkInterface(ctx, &ppool.ReleaseNetworkInterfaceRequest{
					NetworkName:          nic.NetworkName,
					NetworkInterfaceName: vm.NetworkInterfaceNames[i],
				})
				return err
			})

			vm.Nics[i].HardwareAddress = network.ReservedNetworkInterfaces[vm.NetworkInterfaceNames[i]].HardwareAddress
			vm.Nics[i].Ipv4Address = network.ReservedNetworkInterfaces[vm.NetworkInterfaceNames[i]].Ipv4Address
			vm.Nics[i].Ipv6Address = network.ReservedNetworkInterfaces[vm.NetworkInterfaceNames[i]].Ipv6Address

			ip := netutil.ParseCIDR(network.Ipv4Cidr)
			gateway := ""
			for _, ni := range network.ReservedNetworkInterfaces {
				if v, ok := ni.Annotations[AnnotationVirtualMachineVncWebSocketPort]; ok {
					gateway = v
				}
			}
			if gateway == "" {
				if gateway, err = a.addDefaultGateway(ctx, network); err != nil {
					return nil, grpcutil.WrapGrpcErrorf(codes.Internal, errors.Wrapf(err, "Failed to add default gateway").Error())
				}
			}

			netdevs[i] = &NetDev{
				Name:            vm.NetworkInterfaceNames[i],
				NetworkName:     vm.Nics[i].NetworkName,
				HardwareAddress: vm.Nics[i].HardwareAddress,
				Ipv4AddressCidr: fmt.Sprintf("%s/%d", vm.Nics[i].Ipv4Address, ip.SubnetMaskBits()),
				Ipv4Gateway:     gateway,
				Nameservers:     []string{"8.8.8.8"}, // TODO: 取るようにする
				// TODO: domain searchはnetworkのdomainから取る
			}
		}
	}

	{
		conn, err := node.GetConnection(ctx, a.nodeAPI, vm.ComputeNodeName)
		cli := NewVirtualMachineAgentServiceClient(conn)
		if err != nil {
			return nil, grpcutil.WrapGrpcErrorf(codes.Internal, "Failed to dial to node: err=%s", err.Error())
		}
		if conn == nil {
			return nil, grpcutil.WrapGrpcErrorf(codes.FailedPrecondition, "Node '%s' is not ready, so cannot delete: please wait a moment", vm.ComputeNodeName)
		}
		defer conn.Close()

		res, err := cli.BootVirtualMachine(ctx, &BootVirtualMachineRequest{
			Name:              vm.Name,
			Uuid:              vm.Uuid,
			Vcpus:             vm.LimitCpuMilliCore / 1000,
			MemoryBytes:       vm.LimitMemoryBytes,
			Netdevs:           netdevs,
			Blockdevs:         blockdevs,
			LoginUsername:     vm.LoginUsername,
			SshAuthorizedKeys: vm.SshAuthorizedKeys,
		})
		if err != nil {
			return nil, grpcutil.WrapGrpcErrorf(grpc.Code(err), "Failed to CreateVirtualMachineAgent: desc=%s", grpc.ErrorDesc(err))
		}
		tx.PushRollback("", func() error {
			_, err := cli.DeleteVirtualMachine(ctx, &DeleteVirtualMachineRequest{Name: vm.Name})
			return err
		})

		vm.Annotations[AnnotationNetworkInterfaceIsGateway] = strconv.Itoa(int(res.WebsocketPort))
		vm.State = GetAPIStateFromAgentState(res.State)
	}

	if err := a.dataStore.Apply(vm.Name, vm); err != nil {
		return nil, grpcutil.WrapGrpcErrorf(codes.Internal, "Failed to apply data for db: err='%s'", err.Error())
	}

	tx.Commit()
	return vm, nil
}

func (a *VirtualMachineAPI) ListVirtualMachines(ctx context.Context, req *pprovisioning.ListVirtualMachinesRequest) (*pprovisioning.ListVirtualMachinesResponse, error) {
	res := &pprovisioning.ListVirtualMachinesResponse{}
	f := func(s int) []proto.Message {
		res.VirtualMachines = make([]*pprovisioning.VirtualMachine, s)
		for i := range res.VirtualMachines {
			res.VirtualMachines[i] = &pprovisioning.VirtualMachine{}
		}

		m := make([]proto.Message, s)
		for i, v := range res.VirtualMachines {
			m[i] = v
		}

		return m
	}

	if err := a.dataStore.List(f); err != nil {
		log.Printf("[WARNING] Failed to list data from db: err='%s'", err.Error())
		return nil, grpcutil.WrapGrpcErrorf(codes.Internal, "Failed to list from db, please retry or contact for the administrator of this cluster")
	}
	if len(res.VirtualMachines) == 0 {
		return nil, grpcutil.WrapGrpcErrorf(codes.NotFound, "")
	}

	return res, nil
}

func (a *VirtualMachineAPI) GetVirtualMachine(ctx context.Context, req *pprovisioning.GetVirtualMachineRequest) (*pprovisioning.VirtualMachine, error) {
	vm := &pprovisioning.VirtualMachine{}
	if err := a.dataStore.Get(req.Name, vm); err != nil {
		log.Printf("[WARNING] Failed to get data from db: err='%s'", err.Error())
		return nil, grpcutil.WrapGrpcErrorf(codes.Internal, "Failed to get '%s' from db, please retry or contact for the administrator of this cluster", req.Name)
	}
	if vm.Name == "" {
		return nil, grpcutil.WrapGrpcErrorf(codes.NotFound, "")
	}

	return vm, nil
}

func (a *VirtualMachineAPI) UpdateVirtualMachine(ctx context.Context, req *pprovisioning.UpdateVirtualMachineRequest) (*pprovisioning.VirtualMachine, error) {
	return nil, grpcutil.WrapGrpcErrorf(codes.Unimplemented, "")
}

func (a *VirtualMachineAPI) DeleteVirtualMachine(ctx context.Context, req *pprovisioning.DeleteVirtualMachineRequest) (*empty.Empty, error) {
	vm := &pprovisioning.VirtualMachine{}
	if err := a.dataStore.Get(req.Name, vm); err != nil {
		return nil, grpcutil.WrapGrpcErrorf(codes.Internal, "Failed to get '%s' from db: err='%s'", req.Name, err.Error())
	} else if vm.Name == "" {
		return nil, grpcutil.WrapGrpcErrorf(codes.NotFound, "")
	}

	{
		conn, err := node.GetConnection(ctx, a.nodeAPI, vm.ComputeNodeName)
		if err != nil {
			return nil, grpcutil.WrapGrpcErrorf(codes.Internal, "") // TODO: #89
		}
		if conn == nil {
			return nil, grpcutil.WrapGrpcErrorf(codes.FailedPrecondition, "Node '%s' is not ready, so cannot delete: please wait a moment", vm.ComputeNodeName)
		}
		defer conn.Close()
		cli := NewVirtualMachineAgentServiceClient(conn)

		netdevs := make([]*NetDev, len(vm.Nics))
		for i, n := range vm.Nics {
			netdevs[i] = &NetDev{
				Name:            vm.NetworkInterfaceNames[i],
				NetworkName:     n.NetworkName,
				HardwareAddress: n.HardwareAddress,
			}
		}

		_, err = cli.DeleteVirtualMachine(context.Background(), &DeleteVirtualMachineRequest{
			Name:    vm.Name,
			Netdevs: netdevs,
		})
		if err != nil {
			return nil, grpcutil.WrapGrpcErrorf(grpc.Code(err), "Failed to DeleteVirtualMachineAgent: desc=%s", grpc.ErrorDesc(err))
		}
	}

	_, err := a.nodeAPI.ReleaseCompute(context.Background(), &ppool.ReleaseComputeRequest{
		NodeName:    vm.ComputeNodeName,
		ComputeName: vm.ComputeName,
	})
	if err != nil {
		return nil, grpcutil.WrapGrpcErrorf(grpc.Code(err), "Failed to ReleaseCompute: desc=%s", grpc.ErrorDesc(err))
	}

	for i, nic := range vm.Nics {
		_, err := a.networkAPI.ReleaseNetworkInterface(context.Background(), &ppool.ReleaseNetworkInterfaceRequest{
			NetworkName:          nic.NetworkName,
			NetworkInterfaceName: vm.NetworkInterfaceNames[i],
		})
		if err != nil {
			return nil, grpcutil.WrapGrpcErrorf(grpc.Code(err), "Failed to ReleaseNetworkInterface: desc=%s", grpc.ErrorDesc(err))
		}
	}

	for _, n := range vm.BlockStorageNames {
		_, err := a.blockstorageAPI.SetAvailableBlockStorage(context.Background(), &pprovisioning.SetAvailableBlockStorageRequest{Name: n})
		if err != nil {
			return nil, grpcutil.WrapGrpcErrorf(grpc.Code(err), "Failed to SetAvailableBlockStorage: desc=%s", grpc.ErrorDesc(err))
		}
	}

	if err := a.dataStore.Delete(req.Name); err != nil {
		return nil, grpcutil.WrapGrpcErrorf(codes.Internal, "message:Failed to delete from db.\tgot:%v", err.Error())
	}

	return &empty.Empty{}, nil
}

func (a *VirtualMachineAPI) BootVirtualMachine(ctx context.Context, req *pprovisioning.BootVirtualMachineRequest) (*pprovisioning.VirtualMachine, error) {
	return nil, grpcutil.WrapGrpcErrorf(codes.Unimplemented, "")
}

func (a *VirtualMachineAPI) RebootVirtualMachine(ctx context.Context, req *pprovisioning.RebootVirtualMachineRequest) (*pprovisioning.VirtualMachine, error) {
	return nil, grpcutil.WrapGrpcErrorf(codes.Unimplemented, "")
}

func (a *VirtualMachineAPI) ShutdownVirtualMachine(ctx context.Context, req *pprovisioning.ShutdownVirtualMachineRequest) (*pprovisioning.VirtualMachine, error) {
	return nil, grpcutil.WrapGrpcErrorf(codes.Unimplemented, "")
}

func (a *VirtualMachineAPI) SaveVirtualMachine(ctx context.Context, req *pprovisioning.SaveVirtualMachineRequest) (*pprovisioning.VirtualMachine, error) {
	return nil, grpcutil.WrapGrpcErrorf(codes.Unimplemented, "")
}

// TODO: めんどくさいので n0core コマンドで定義した URL に一時的に依存している、治す必要あり
func (a *VirtualMachineAPI) OpenConsole(ctx context.Context, req *pprovisioning.OpenConsoleRequest) (*pprovisioning.OpenConsoleResponse, error) {
	vm := &pprovisioning.VirtualMachine{}
	if err := a.dataStore.Get(req.Name, vm); err != nil {
		log.Printf("[WARNING] Failed to get data from db: err='%s'", err.Error())
		return nil, grpc.Errorf(codes.Internal, "Failed to get '%s' from db, please retry or contact for the administrator of this cluster", req.Name)
	} else if vm.Name == "" {
		return nil, grpc.Errorf(codes.NotFound, "")
	}

	u := &url.URL{
		Scheme:   "http",
		Path:     "/static/virtual_machines/vnc.html",
		RawQuery: fmt.Sprintf("path=api/v0/virtual_machines/%s/vncwebsocket", vm.Name),
	}

	return &pprovisioning.OpenConsoleResponse{
		ConsoleUrl: u.String(),
	}, nil
}

func (a *VirtualMachineAPI) ProxyWebsocket() func(echo.Context) error {
	return func(c echo.Context) error {
		vmName := c.Param("name")

		vm := &pprovisioning.VirtualMachine{}
		if err := a.dataStore.Get(vmName, vm); err != nil {
			log.Printf("[WARNING] Failed to get data from db: err='%s'", err.Error())
			return fmt.Errorf("db error")
		} else if vm.Name == "" {
			return err
		}

		node, err := a.nodeAPI.GetNode(context.Background(), &ppool.GetNodeRequest{Name: vm.ComputeNodeName})
		if err != nil {
			return err
		}

		nodeIP := node.Address
		websocketPort, err := strconv.Atoi(vm.Annotations[AnnotationNetworkInterfaceIsGateway])
		if err != nil {
			return err
		}

		u := &url.URL{
			Scheme: "ws",
			Host:   fmt.Sprintf("%s:%d", nodeIP, websocketPort),
			Path:   "/",
		}

		ws := &websocketproxy.WebsocketProxy{
			Backend: func(r *http.Request) *url.URL {
				return u
			},
		}
		// TODO: セキュリティを無視して、とりあえず動かす https://github.com/koding/websocketproxy/issues/9
		delete(c.Request().Header, "Origin")
		log.Printf("[DEBUG] websocket proxy requesting to backend '%s'", ws.Backend(c.Request()))
		ws.ServeHTTP(c.Response(), c.Request())

		return nil
	}
}