// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/kvm/kvm.proto

/*
Package kvm is a generated protocol buffer package.

It is generated from these files:
	proto/kvm/kvm.proto

It has these top-level messages:
	KVM
	ApplyKVMRequest
	DeleteKVMRequest
	ActionKVMRequest
*/
package kvm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type KVM struct {
	Uuid string `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// CPU
	CpuCores uint32 `protobuf:"varint,3,opt,name=cpu_cores,json=cpuCores" json:"cpu_cores,omitempty"`
	// Memory
	MemoryBytes uint64                  `protobuf:"varint,4,opt,name=memory_bytes,json=memoryBytes" json:"memory_bytes,omitempty"`
	Storages    map[string]*KVM_Storage `protobuf:"bytes,5,rep,name=storages" json:"storages,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Nics        map[string]*KVM_NIC     `protobuf:"bytes,6,rep,name=nics" json:"nics,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// TCP port of websocket vnc which is opened by qemu
	VncWebsocketPort uint32 `protobuf:"varint,7,opt,name=vnc_websocket_port,json=vncWebsocketPort" json:"vnc_websocket_port,omitempty"`
	QmpPath          string `protobuf:"bytes,8,opt,name=qmp_path,json=qmpPath" json:"qmp_path,omitempty"`
}

func (m *KVM) Reset()                    { *m = KVM{} }
func (m *KVM) String() string            { return proto.CompactTextString(m) }
func (*KVM) ProtoMessage()               {}
func (*KVM) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *KVM) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *KVM) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KVM) GetCpuCores() uint32 {
	if m != nil {
		return m.CpuCores
	}
	return 0
}

func (m *KVM) GetMemoryBytes() uint64 {
	if m != nil {
		return m.MemoryBytes
	}
	return 0
}

func (m *KVM) GetStorages() map[string]*KVM_Storage {
	if m != nil {
		return m.Storages
	}
	return nil
}

func (m *KVM) GetNics() map[string]*KVM_NIC {
	if m != nil {
		return m.Nics
	}
	return nil
}

func (m *KVM) GetVncWebsocketPort() uint32 {
	if m != nil {
		return m.VncWebsocketPort
	}
	return 0
}

func (m *KVM) GetQmpPath() string {
	if m != nil {
		return m.QmpPath
	}
	return ""
}

// Storage
// definition(label, url)
type KVM_Storage struct {
	Url       string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	BootIndex uint32 `protobuf:"varint,2,opt,name=boot_index,json=bootIndex" json:"boot_index,omitempty"`
}

func (m *KVM_Storage) Reset()                    { *m = KVM_Storage{} }
func (m *KVM_Storage) String() string            { return proto.CompactTextString(m) }
func (*KVM_Storage) ProtoMessage()               {}
func (*KVM_Storage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *KVM_Storage) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *KVM_Storage) GetBootIndex() uint32 {
	if m != nil {
		return m.BootIndex
	}
	return 0
}

// Network
// definition(tap_name, hwaddr)
type KVM_NIC struct {
	TapName string `protobuf:"bytes,1,opt,name=tap_name,json=tapName" json:"tap_name,omitempty"`
	HwAddr  string `protobuf:"bytes,2,opt,name=hw_addr,json=hwAddr" json:"hw_addr,omitempty"`
}

func (m *KVM_NIC) Reset()                    { *m = KVM_NIC{} }
func (m *KVM_NIC) String() string            { return proto.CompactTextString(m) }
func (*KVM_NIC) ProtoMessage()               {}
func (*KVM_NIC) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 2} }

func (m *KVM_NIC) GetTapName() string {
	if m != nil {
		return m.TapName
	}
	return ""
}

func (m *KVM_NIC) GetHwAddr() string {
	if m != nil {
		return m.HwAddr
	}
	return ""
}

type ApplyKVMRequest struct {
	Kvm *KVM `protobuf:"bytes,1,opt,name=kvm" json:"kvm,omitempty"`
}

func (m *ApplyKVMRequest) Reset()                    { *m = ApplyKVMRequest{} }
func (m *ApplyKVMRequest) String() string            { return proto.CompactTextString(m) }
func (*ApplyKVMRequest) ProtoMessage()               {}
func (*ApplyKVMRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ApplyKVMRequest) GetKvm() *KVM {
	if m != nil {
		return m.Kvm
	}
	return nil
}

type DeleteKVMRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteKVMRequest) Reset()                    { *m = DeleteKVMRequest{} }
func (m *DeleteKVMRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteKVMRequest) ProtoMessage()               {}
func (*DeleteKVMRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DeleteKVMRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ActionKVMRequest struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	QmpPath string `protobuf:"bytes,8,opt,name=qmp_path,json=qmpPath" json:"qmp_path,omitempty"`
}

func (m *ActionKVMRequest) Reset()                    { *m = ActionKVMRequest{} }
func (m *ActionKVMRequest) String() string            { return proto.CompactTextString(m) }
func (*ActionKVMRequest) ProtoMessage()               {}
func (*ActionKVMRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ActionKVMRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ActionKVMRequest) GetQmpPath() string {
	if m != nil {
		return m.QmpPath
	}
	return ""
}

func init() {
	proto.RegisterType((*KVM)(nil), "n0stack.n0core.kvm.KVM")
	proto.RegisterType((*KVM_Storage)(nil), "n0stack.n0core.kvm.KVM.Storage")
	proto.RegisterType((*KVM_NIC)(nil), "n0stack.n0core.kvm.KVM.NIC")
	proto.RegisterType((*ApplyKVMRequest)(nil), "n0stack.n0core.kvm.ApplyKVMRequest")
	proto.RegisterType((*DeleteKVMRequest)(nil), "n0stack.n0core.kvm.DeleteKVMRequest")
	proto.RegisterType((*ActionKVMRequest)(nil), "n0stack.n0core.kvm.ActionKVMRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for KVMService service

type KVMServiceClient interface {
	ApplyKVM(ctx context.Context, in *ApplyKVMRequest, opts ...grpc.CallOption) (*KVM, error)
	DeleteKVM(ctx context.Context, in *DeleteKVMRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	// VM actions
	Boot(ctx context.Context, in *ActionKVMRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	Reboot(ctx context.Context, in *ActionKVMRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	HardReboot(ctx context.Context, in *ActionKVMRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	Shutdown(ctx context.Context, in *ActionKVMRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	HardShutdown(ctx context.Context, in *ActionKVMRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	Save(ctx context.Context, in *ActionKVMRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type kVMServiceClient struct {
	cc *grpc.ClientConn
}

func NewKVMServiceClient(cc *grpc.ClientConn) KVMServiceClient {
	return &kVMServiceClient{cc}
}

func (c *kVMServiceClient) ApplyKVM(ctx context.Context, in *ApplyKVMRequest, opts ...grpc.CallOption) (*KVM, error) {
	out := new(KVM)
	err := grpc.Invoke(ctx, "/n0stack.n0core.kvm.KVMService/ApplyKVM", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVMServiceClient) DeleteKVM(ctx context.Context, in *DeleteKVMRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/n0stack.n0core.kvm.KVMService/DeleteKVM", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVMServiceClient) Boot(ctx context.Context, in *ActionKVMRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/n0stack.n0core.kvm.KVMService/Boot", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVMServiceClient) Reboot(ctx context.Context, in *ActionKVMRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/n0stack.n0core.kvm.KVMService/Reboot", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVMServiceClient) HardReboot(ctx context.Context, in *ActionKVMRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/n0stack.n0core.kvm.KVMService/HardReboot", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVMServiceClient) Shutdown(ctx context.Context, in *ActionKVMRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/n0stack.n0core.kvm.KVMService/Shutdown", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVMServiceClient) HardShutdown(ctx context.Context, in *ActionKVMRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/n0stack.n0core.kvm.KVMService/HardShutdown", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVMServiceClient) Save(ctx context.Context, in *ActionKVMRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/n0stack.n0core.kvm.KVMService/Save", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for KVMService service

type KVMServiceServer interface {
	ApplyKVM(context.Context, *ApplyKVMRequest) (*KVM, error)
	DeleteKVM(context.Context, *DeleteKVMRequest) (*google_protobuf.Empty, error)
	// VM actions
	Boot(context.Context, *ActionKVMRequest) (*google_protobuf.Empty, error)
	Reboot(context.Context, *ActionKVMRequest) (*google_protobuf.Empty, error)
	HardReboot(context.Context, *ActionKVMRequest) (*google_protobuf.Empty, error)
	Shutdown(context.Context, *ActionKVMRequest) (*google_protobuf.Empty, error)
	HardShutdown(context.Context, *ActionKVMRequest) (*google_protobuf.Empty, error)
	Save(context.Context, *ActionKVMRequest) (*google_protobuf.Empty, error)
}

func RegisterKVMServiceServer(s *grpc.Server, srv KVMServiceServer) {
	s.RegisterService(&_KVMService_serviceDesc, srv)
}

func _KVMService_ApplyKVM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyKVMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVMServiceServer).ApplyKVM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.n0core.kvm.KVMService/ApplyKVM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVMServiceServer).ApplyKVM(ctx, req.(*ApplyKVMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVMService_DeleteKVM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteKVMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVMServiceServer).DeleteKVM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.n0core.kvm.KVMService/DeleteKVM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVMServiceServer).DeleteKVM(ctx, req.(*DeleteKVMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVMService_Boot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionKVMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVMServiceServer).Boot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.n0core.kvm.KVMService/Boot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVMServiceServer).Boot(ctx, req.(*ActionKVMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVMService_Reboot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionKVMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVMServiceServer).Reboot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.n0core.kvm.KVMService/Reboot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVMServiceServer).Reboot(ctx, req.(*ActionKVMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVMService_HardReboot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionKVMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVMServiceServer).HardReboot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.n0core.kvm.KVMService/HardReboot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVMServiceServer).HardReboot(ctx, req.(*ActionKVMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVMService_Shutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionKVMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVMServiceServer).Shutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.n0core.kvm.KVMService/Shutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVMServiceServer).Shutdown(ctx, req.(*ActionKVMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVMService_HardShutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionKVMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVMServiceServer).HardShutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.n0core.kvm.KVMService/HardShutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVMServiceServer).HardShutdown(ctx, req.(*ActionKVMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVMService_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionKVMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVMServiceServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.n0core.kvm.KVMService/Save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVMServiceServer).Save(ctx, req.(*ActionKVMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KVMService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "n0stack.n0core.kvm.KVMService",
	HandlerType: (*KVMServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ApplyKVM",
			Handler:    _KVMService_ApplyKVM_Handler,
		},
		{
			MethodName: "DeleteKVM",
			Handler:    _KVMService_DeleteKVM_Handler,
		},
		{
			MethodName: "Boot",
			Handler:    _KVMService_Boot_Handler,
		},
		{
			MethodName: "Reboot",
			Handler:    _KVMService_Reboot_Handler,
		},
		{
			MethodName: "HardReboot",
			Handler:    _KVMService_HardReboot_Handler,
		},
		{
			MethodName: "Shutdown",
			Handler:    _KVMService_Shutdown_Handler,
		},
		{
			MethodName: "HardShutdown",
			Handler:    _KVMService_HardShutdown_Handler,
		},
		{
			MethodName: "Save",
			Handler:    _KVMService_Save_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/kvm/kvm.proto",
}

func init() { proto.RegisterFile("proto/kvm/kvm.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 626 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x94, 0x6d, 0x6b, 0xdb, 0x30,
	0x10, 0xc7, 0x9b, 0xc6, 0xcd, 0xc3, 0xa5, 0x65, 0x41, 0x83, 0xd5, 0x73, 0x19, 0x4b, 0xb3, 0x07,
	0x32, 0x18, 0x76, 0x9b, 0x51, 0xd8, 0xca, 0xde, 0xa4, 0x5d, 0x4b, 0xdb, 0x2c, 0xa1, 0x38, 0x23,
	0x83, 0x31, 0x30, 0x7e, 0xd0, 0x12, 0xe3, 0x58, 0x52, 0x65, 0xd9, 0x99, 0x3f, 0xeb, 0x3e, 0xc1,
	0xbe, 0xc5, 0x90, 0x9d, 0x74, 0x6d, 0x97, 0x74, 0x2f, 0x9a, 0x17, 0x06, 0xe9, 0x7c, 0xf7, 0xe3,
	0x7f, 0xff, 0x93, 0x04, 0x8f, 0x19, 0xa7, 0x82, 0x1a, 0x41, 0x12, 0xca, 0x4f, 0xcf, 0x76, 0x08,
	0x91, 0xbd, 0x48, 0xd8, 0x6e, 0xa0, 0x93, 0x3d, 0x97, 0x72, 0xac, 0x07, 0x49, 0xa8, 0xed, 0x8c,
	0x28, 0x1d, 0x4d, 0xb0, 0x91, 0x65, 0x38, 0xf1, 0x0f, 0x03, 0x87, 0x4c, 0xa4, 0x79, 0x41, 0xf3,
	0x97, 0x02, 0xc5, 0xee, 0xb0, 0x87, 0x10, 0x28, 0x71, 0xec, 0x7b, 0x6a, 0xa1, 0x51, 0x68, 0x55,
	0xcd, 0x6c, 0x2d, 0x63, 0xc4, 0x0e, 0xb1, 0xba, 0x9e, 0xc7, 0xe4, 0x1a, 0xed, 0x40, 0xd5, 0x65,
	0xb1, 0x25, 0xe1, 0x91, 0x5a, 0x6c, 0x14, 0x5a, 0x5b, 0x66, 0xc5, 0x65, 0xf1, 0xb1, 0xdc, 0xa3,
	0x5d, 0xd8, 0x0c, 0x71, 0x48, 0x79, 0x6a, 0x39, 0xa9, 0xc0, 0x91, 0xaa, 0x34, 0x0a, 0x2d, 0xc5,
	0xac, 0xe5, 0xb1, 0x23, 0x19, 0x42, 0x1d, 0xa8, 0x44, 0x82, 0x72, 0x7b, 0x84, 0x23, 0x75, 0xa3,
	0x51, 0x6c, 0xd5, 0xda, 0xaf, 0xf4, 0x7f, 0x35, 0xeb, 0xdd, 0x61, 0x4f, 0x1f, 0xcc, 0xf2, 0x4e,
	0x88, 0xe0, 0xa9, 0x79, 0x5d, 0x86, 0x0e, 0x40, 0x21, 0xbe, 0x1b, 0xa9, 0xa5, 0xac, 0x7c, 0x77,
	0x59, 0x79, 0xdf, 0x77, 0x67, 0xa5, 0x59, 0x3a, 0x7a, 0x0b, 0x28, 0x21, 0xae, 0x35, 0xc5, 0x4e,
	0x44, 0xdd, 0x00, 0x0b, 0x8b, 0x51, 0x2e, 0xd4, 0x72, 0xd6, 0x42, 0x3d, 0x21, 0xee, 0xd7, 0xf9,
	0x8f, 0x4b, 0xca, 0x05, 0x7a, 0x0a, 0x95, 0xab, 0x90, 0x59, 0xcc, 0x16, 0x63, 0xb5, 0x92, 0xf5,
	0x5f, 0xbe, 0x0a, 0xd9, 0xa5, 0x2d, 0xc6, 0xda, 0x21, 0x94, 0x67, 0xd2, 0x50, 0x1d, 0x8a, 0x31,
	0x9f, 0xcc, 0x4c, 0x93, 0x4b, 0xf4, 0x0c, 0xc0, 0xa1, 0x54, 0x58, 0x3e, 0xf1, 0xf0, 0xcf, 0xcc,
	0xb9, 0x2d, 0xb3, 0x2a, 0x23, 0xe7, 0x32, 0xa0, 0x7d, 0x87, 0xad, 0x5b, 0x6d, 0x49, 0x42, 0x80,
	0xd3, 0x39, 0x21, 0xc0, 0x29, 0x3a, 0x80, 0x8d, 0xc4, 0x9e, 0xc4, 0xb9, 0xed, 0xb5, 0xf6, 0xf3,
	0xff, 0xd8, 0x63, 0xe6, 0xd9, 0x87, 0xeb, 0xef, 0x0b, 0xda, 0x07, 0x28, 0xf6, 0xcf, 0x8f, 0xa5,
	0x76, 0x61, 0x33, 0x2b, 0x9b, 0x5d, 0x0e, 0x2e, 0x0b, 0x9b, 0xf5, 0xe5, 0xf8, 0xb6, 0xa1, 0x3c,
	0x9e, 0x5a, 0xb6, 0xe7, 0xf1, 0xd9, 0x54, 0x4b, 0xe3, 0x69, 0xc7, 0xf3, 0xb8, 0xf6, 0x05, 0xaa,
	0xd7, 0x86, 0x2d, 0x10, 0xb5, 0x7f, 0x5b, 0xd4, 0xce, 0x52, 0xd3, 0xcf, 0x8f, 0x6f, 0x08, 0x6a,
	0x7e, 0x84, 0x47, 0x1d, 0xc6, 0x26, 0x69, 0x77, 0xd8, 0x33, 0xf1, 0x55, 0x8c, 0x23, 0x81, 0xde,
	0x40, 0x31, 0x48, 0xc2, 0x8c, 0x5d, 0x6b, 0x6f, 0x2f, 0xe1, 0x98, 0x32, 0xa7, 0xf9, 0x1a, 0xea,
	0x9f, 0xf0, 0x04, 0x0b, 0x7c, 0xa3, 0x7c, 0x7e, 0x26, 0x0b, 0x7f, 0xcf, 0x64, 0xb3, 0x03, 0xf5,
	0x8e, 0x2b, 0x7c, 0x4a, 0xee, 0xcf, 0xbb, 0x67, 0xa6, 0xed, 0xdf, 0x0a, 0x40, 0x77, 0xd8, 0x1b,
	0x60, 0x9e, 0xf8, 0x2e, 0x46, 0x17, 0x50, 0x99, 0xeb, 0x46, 0x2f, 0x16, 0x69, 0xbc, 0xd3, 0x95,
	0xb6, 0xac, 0x91, 0xe6, 0x1a, 0xea, 0x42, 0xf5, 0xba, 0x0b, 0xf4, 0x72, 0x51, 0xde, 0xdd, 0x26,
	0xb5, 0x27, 0x7a, 0x7e, 0x65, 0xf5, 0xf9, 0x95, 0xd5, 0x4f, 0xe4, 0x95, 0x6d, 0xae, 0xa1, 0x53,
	0x50, 0x8e, 0x28, 0x15, 0x8b, 0x39, 0x77, 0x4d, 0xb8, 0x87, 0x73, 0x06, 0x25, 0x13, 0x3b, 0xab,
	0x20, 0x7d, 0x06, 0x38, 0xb3, 0xb9, 0xb7, 0x22, 0xda, 0x05, 0x54, 0x06, 0xe3, 0x58, 0x78, 0x74,
	0x4a, 0x1e, 0xcc, 0xea, 0xc3, 0xa6, 0x54, 0xb6, 0x32, 0xde, 0x29, 0x28, 0x03, 0x3b, 0xc1, 0x0f,
	0xe5, 0x1c, 0xed, 0x7f, 0x33, 0x46, 0xbe, 0x18, 0xc7, 0x8e, 0xee, 0xd2, 0xd0, 0x98, 0xb1, 0x8c,
	0x9c, 0x25, 0x1f, 0xe9, 0xc4, 0x8f, 0x7c, 0x4a, 0x7c, 0x32, 0x32, 0x08, 0xf5, 0xb0, 0x7c, 0xdc,
	0x9d, 0x52, 0x06, 0x79, 0xf7, 0x27, 0x00, 0x00, 0xff, 0xff, 0x11, 0x52, 0xf8, 0x0a, 0xf4, 0x05,
	0x00, 0x00,
}
