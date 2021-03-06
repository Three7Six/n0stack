// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pool/v0/network.proto

package ppool // import "github.com/n0stack/n0stack/n0proto.go/pool/v0"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import v0 "github.com/n0stack/n0stack/n0proto.go/budget/v0"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type Network_NetworkState int32

const (
	// falied state because failed some process on API.
	Network_FAILED Network_NetworkState = 0
	// unknown state because failed to connect for scheduled node after RUNNING.
	Network_UNKNOWN   Network_NetworkState = 1
	Network_AVAILABLE Network_NetworkState = 2
)

var Network_NetworkState_name = map[int32]string{
	0: "FAILED",
	1: "UNKNOWN",
	2: "AVAILABLE",
}
var Network_NetworkState_value = map[string]int32{
	"FAILED":    0,
	"UNKNOWN":   1,
	"AVAILABLE": 2,
}

func (x Network_NetworkState) String() string {
	return proto.EnumName(Network_NetworkState_name, int32(x))
}
func (Network_NetworkState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_network_faeccbe0a81f7f5a, []int{0, 0}
}

type Network struct {
	Name                      string                          `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Annotations               map[string]string               `protobuf:"bytes,3,rep,name=annotations" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Ipv4Cidr                  string                          `protobuf:"bytes,10,opt,name=ipv4_cidr,json=ipv4Cidr" json:"ipv4_cidr,omitempty"`
	Ipv6Cidr                  string                          `protobuf:"bytes,11,opt,name=ipv6_cidr,json=ipv6Cidr" json:"ipv6_cidr,omitempty"`
	Domain                    string                          `protobuf:"bytes,12,opt,name=domain" json:"domain,omitempty"`
	State                     Network_NetworkState            `protobuf:"varint,50,opt,name=state,enum=n0stack.pool.Network_NetworkState" json:"state,omitempty"`
	ReservedNetworkInterfaces map[string]*v0.NetworkInterface `protobuf:"bytes,51,rep,name=reserved_network_interfaces,json=reservedNetworkInterfaces" json:"reserved_network_interfaces,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral      struct{}                        `json:"-"`
	XXX_unrecognized          []byte                          `json:"-"`
	XXX_sizecache             int32                           `json:"-"`
}

func (m *Network) Reset()         { *m = Network{} }
func (m *Network) String() string { return proto.CompactTextString(m) }
func (*Network) ProtoMessage()    {}
func (*Network) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_faeccbe0a81f7f5a, []int{0}
}
func (m *Network) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Network.Unmarshal(m, b)
}
func (m *Network) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Network.Marshal(b, m, deterministic)
}
func (dst *Network) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Network.Merge(dst, src)
}
func (m *Network) XXX_Size() int {
	return xxx_messageInfo_Network.Size(m)
}
func (m *Network) XXX_DiscardUnknown() {
	xxx_messageInfo_Network.DiscardUnknown(m)
}

var xxx_messageInfo_Network proto.InternalMessageInfo

func (m *Network) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Network) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *Network) GetIpv4Cidr() string {
	if m != nil {
		return m.Ipv4Cidr
	}
	return ""
}

func (m *Network) GetIpv6Cidr() string {
	if m != nil {
		return m.Ipv6Cidr
	}
	return ""
}

func (m *Network) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Network) GetState() Network_NetworkState {
	if m != nil {
		return m.State
	}
	return Network_FAILED
}

func (m *Network) GetReservedNetworkInterfaces() map[string]*v0.NetworkInterface {
	if m != nil {
		return m.ReservedNetworkInterfaces
	}
	return nil
}

type ListNetworksRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListNetworksRequest) Reset()         { *m = ListNetworksRequest{} }
func (m *ListNetworksRequest) String() string { return proto.CompactTextString(m) }
func (*ListNetworksRequest) ProtoMessage()    {}
func (*ListNetworksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_faeccbe0a81f7f5a, []int{1}
}
func (m *ListNetworksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNetworksRequest.Unmarshal(m, b)
}
func (m *ListNetworksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNetworksRequest.Marshal(b, m, deterministic)
}
func (dst *ListNetworksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNetworksRequest.Merge(dst, src)
}
func (m *ListNetworksRequest) XXX_Size() int {
	return xxx_messageInfo_ListNetworksRequest.Size(m)
}
func (m *ListNetworksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNetworksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListNetworksRequest proto.InternalMessageInfo

type ListNetworksResponse struct {
	Networks             []*Network `protobuf:"bytes,1,rep,name=networks" json:"networks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListNetworksResponse) Reset()         { *m = ListNetworksResponse{} }
func (m *ListNetworksResponse) String() string { return proto.CompactTextString(m) }
func (*ListNetworksResponse) ProtoMessage()    {}
func (*ListNetworksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_faeccbe0a81f7f5a, []int{2}
}
func (m *ListNetworksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNetworksResponse.Unmarshal(m, b)
}
func (m *ListNetworksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNetworksResponse.Marshal(b, m, deterministic)
}
func (dst *ListNetworksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNetworksResponse.Merge(dst, src)
}
func (m *ListNetworksResponse) XXX_Size() int {
	return xxx_messageInfo_ListNetworksResponse.Size(m)
}
func (m *ListNetworksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNetworksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListNetworksResponse proto.InternalMessageInfo

func (m *ListNetworksResponse) GetNetworks() []*Network {
	if m != nil {
		return m.Networks
	}
	return nil
}

type GetNetworkRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNetworkRequest) Reset()         { *m = GetNetworkRequest{} }
func (m *GetNetworkRequest) String() string { return proto.CompactTextString(m) }
func (*GetNetworkRequest) ProtoMessage()    {}
func (*GetNetworkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_faeccbe0a81f7f5a, []int{3}
}
func (m *GetNetworkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNetworkRequest.Unmarshal(m, b)
}
func (m *GetNetworkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNetworkRequest.Marshal(b, m, deterministic)
}
func (dst *GetNetworkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNetworkRequest.Merge(dst, src)
}
func (m *GetNetworkRequest) XXX_Size() int {
	return xxx_messageInfo_GetNetworkRequest.Size(m)
}
func (m *GetNetworkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNetworkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetNetworkRequest proto.InternalMessageInfo

func (m *GetNetworkRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ApplyNetworkRequest struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Annotations          map[string]string `protobuf:"bytes,3,rep,name=annotations" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Ipv4Cidr             string            `protobuf:"bytes,10,opt,name=ipv4_cidr,json=ipv4Cidr" json:"ipv4_cidr,omitempty"`
	Ipv6Cidr             string            `protobuf:"bytes,11,opt,name=ipv6_cidr,json=ipv6Cidr" json:"ipv6_cidr,omitempty"`
	Domain               string            `protobuf:"bytes,12,opt,name=domain" json:"domain,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ApplyNetworkRequest) Reset()         { *m = ApplyNetworkRequest{} }
func (m *ApplyNetworkRequest) String() string { return proto.CompactTextString(m) }
func (*ApplyNetworkRequest) ProtoMessage()    {}
func (*ApplyNetworkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_faeccbe0a81f7f5a, []int{4}
}
func (m *ApplyNetworkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyNetworkRequest.Unmarshal(m, b)
}
func (m *ApplyNetworkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyNetworkRequest.Marshal(b, m, deterministic)
}
func (dst *ApplyNetworkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyNetworkRequest.Merge(dst, src)
}
func (m *ApplyNetworkRequest) XXX_Size() int {
	return xxx_messageInfo_ApplyNetworkRequest.Size(m)
}
func (m *ApplyNetworkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyNetworkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyNetworkRequest proto.InternalMessageInfo

func (m *ApplyNetworkRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ApplyNetworkRequest) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *ApplyNetworkRequest) GetIpv4Cidr() string {
	if m != nil {
		return m.Ipv4Cidr
	}
	return ""
}

func (m *ApplyNetworkRequest) GetIpv6Cidr() string {
	if m != nil {
		return m.Ipv6Cidr
	}
	return ""
}

func (m *ApplyNetworkRequest) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

type DeleteNetworkRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteNetworkRequest) Reset()         { *m = DeleteNetworkRequest{} }
func (m *DeleteNetworkRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteNetworkRequest) ProtoMessage()    {}
func (*DeleteNetworkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_faeccbe0a81f7f5a, []int{5}
}
func (m *DeleteNetworkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteNetworkRequest.Unmarshal(m, b)
}
func (m *DeleteNetworkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteNetworkRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteNetworkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteNetworkRequest.Merge(dst, src)
}
func (m *DeleteNetworkRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteNetworkRequest.Size(m)
}
func (m *DeleteNetworkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteNetworkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteNetworkRequest proto.InternalMessageInfo

func (m *DeleteNetworkRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ReserveNetworkInterfaceRequest struct {
	NetworkName          string            `protobuf:"bytes,1,opt,name=network_name,json=networkName" json:"network_name,omitempty"`
	NetworkInterfaceName string            `protobuf:"bytes,2,opt,name=network_interface_name,json=networkInterfaceName" json:"network_interface_name,omitempty"`
	Annotations          map[string]string `protobuf:"bytes,3,rep,name=annotations" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	HardwareAddress      string            `protobuf:"bytes,4,opt,name=hardware_address,json=hardwareAddress" json:"hardware_address,omitempty"`
	Ipv4Address          string            `protobuf:"bytes,5,opt,name=ipv4_address,json=ipv4Address" json:"ipv4_address,omitempty"`
	Ipv6Address          string            `protobuf:"bytes,6,opt,name=ipv6_address,json=ipv6Address" json:"ipv6_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ReserveNetworkInterfaceRequest) Reset()         { *m = ReserveNetworkInterfaceRequest{} }
func (m *ReserveNetworkInterfaceRequest) String() string { return proto.CompactTextString(m) }
func (*ReserveNetworkInterfaceRequest) ProtoMessage()    {}
func (*ReserveNetworkInterfaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_faeccbe0a81f7f5a, []int{6}
}
func (m *ReserveNetworkInterfaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReserveNetworkInterfaceRequest.Unmarshal(m, b)
}
func (m *ReserveNetworkInterfaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReserveNetworkInterfaceRequest.Marshal(b, m, deterministic)
}
func (dst *ReserveNetworkInterfaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReserveNetworkInterfaceRequest.Merge(dst, src)
}
func (m *ReserveNetworkInterfaceRequest) XXX_Size() int {
	return xxx_messageInfo_ReserveNetworkInterfaceRequest.Size(m)
}
func (m *ReserveNetworkInterfaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReserveNetworkInterfaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReserveNetworkInterfaceRequest proto.InternalMessageInfo

func (m *ReserveNetworkInterfaceRequest) GetNetworkName() string {
	if m != nil {
		return m.NetworkName
	}
	return ""
}

func (m *ReserveNetworkInterfaceRequest) GetNetworkInterfaceName() string {
	if m != nil {
		return m.NetworkInterfaceName
	}
	return ""
}

func (m *ReserveNetworkInterfaceRequest) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *ReserveNetworkInterfaceRequest) GetHardwareAddress() string {
	if m != nil {
		return m.HardwareAddress
	}
	return ""
}

func (m *ReserveNetworkInterfaceRequest) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *ReserveNetworkInterfaceRequest) GetIpv6Address() string {
	if m != nil {
		return m.Ipv6Address
	}
	return ""
}

type ReleaseNetworkInterfaceRequest struct {
	NetworkName          string   `protobuf:"bytes,1,opt,name=network_name,json=networkName" json:"network_name,omitempty"`
	NetworkInterfaceName string   `protobuf:"bytes,2,opt,name=network_interface_name,json=networkInterfaceName" json:"network_interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReleaseNetworkInterfaceRequest) Reset()         { *m = ReleaseNetworkInterfaceRequest{} }
func (m *ReleaseNetworkInterfaceRequest) String() string { return proto.CompactTextString(m) }
func (*ReleaseNetworkInterfaceRequest) ProtoMessage()    {}
func (*ReleaseNetworkInterfaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_faeccbe0a81f7f5a, []int{7}
}
func (m *ReleaseNetworkInterfaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleaseNetworkInterfaceRequest.Unmarshal(m, b)
}
func (m *ReleaseNetworkInterfaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleaseNetworkInterfaceRequest.Marshal(b, m, deterministic)
}
func (dst *ReleaseNetworkInterfaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleaseNetworkInterfaceRequest.Merge(dst, src)
}
func (m *ReleaseNetworkInterfaceRequest) XXX_Size() int {
	return xxx_messageInfo_ReleaseNetworkInterfaceRequest.Size(m)
}
func (m *ReleaseNetworkInterfaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleaseNetworkInterfaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReleaseNetworkInterfaceRequest proto.InternalMessageInfo

func (m *ReleaseNetworkInterfaceRequest) GetNetworkName() string {
	if m != nil {
		return m.NetworkName
	}
	return ""
}

func (m *ReleaseNetworkInterfaceRequest) GetNetworkInterfaceName() string {
	if m != nil {
		return m.NetworkInterfaceName
	}
	return ""
}

func init() {
	proto.RegisterType((*Network)(nil), "n0stack.pool.Network")
	proto.RegisterMapType((map[string]string)(nil), "n0stack.pool.Network.AnnotationsEntry")
	proto.RegisterMapType((map[string]*v0.NetworkInterface)(nil), "n0stack.pool.Network.ReservedNetworkInterfacesEntry")
	proto.RegisterType((*ListNetworksRequest)(nil), "n0stack.pool.ListNetworksRequest")
	proto.RegisterType((*ListNetworksResponse)(nil), "n0stack.pool.ListNetworksResponse")
	proto.RegisterType((*GetNetworkRequest)(nil), "n0stack.pool.GetNetworkRequest")
	proto.RegisterType((*ApplyNetworkRequest)(nil), "n0stack.pool.ApplyNetworkRequest")
	proto.RegisterMapType((map[string]string)(nil), "n0stack.pool.ApplyNetworkRequest.AnnotationsEntry")
	proto.RegisterType((*DeleteNetworkRequest)(nil), "n0stack.pool.DeleteNetworkRequest")
	proto.RegisterType((*ReserveNetworkInterfaceRequest)(nil), "n0stack.pool.ReserveNetworkInterfaceRequest")
	proto.RegisterMapType((map[string]string)(nil), "n0stack.pool.ReserveNetworkInterfaceRequest.AnnotationsEntry")
	proto.RegisterType((*ReleaseNetworkInterfaceRequest)(nil), "n0stack.pool.ReleaseNetworkInterfaceRequest")
	proto.RegisterEnum("n0stack.pool.Network_NetworkState", Network_NetworkState_name, Network_NetworkState_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NetworkService service

type NetworkServiceClient interface {
	ListNetworks(ctx context.Context, in *ListNetworksRequest, opts ...grpc.CallOption) (*ListNetworksResponse, error)
	GetNetwork(ctx context.Context, in *GetNetworkRequest, opts ...grpc.CallOption) (*Network, error)
	ApplyNetwork(ctx context.Context, in *ApplyNetworkRequest, opts ...grpc.CallOption) (*Network, error)
	DeleteNetwork(ctx context.Context, in *DeleteNetworkRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ReserveNetworkInterface(ctx context.Context, in *ReserveNetworkInterfaceRequest, opts ...grpc.CallOption) (*Network, error)
	ReleaseNetworkInterface(ctx context.Context, in *ReleaseNetworkInterfaceRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type networkServiceClient struct {
	cc *grpc.ClientConn
}

func NewNetworkServiceClient(cc *grpc.ClientConn) NetworkServiceClient {
	return &networkServiceClient{cc}
}

func (c *networkServiceClient) ListNetworks(ctx context.Context, in *ListNetworksRequest, opts ...grpc.CallOption) (*ListNetworksResponse, error) {
	out := new(ListNetworksResponse)
	err := grpc.Invoke(ctx, "/n0stack.pool.NetworkService/ListNetworks", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) GetNetwork(ctx context.Context, in *GetNetworkRequest, opts ...grpc.CallOption) (*Network, error) {
	out := new(Network)
	err := grpc.Invoke(ctx, "/n0stack.pool.NetworkService/GetNetwork", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) ApplyNetwork(ctx context.Context, in *ApplyNetworkRequest, opts ...grpc.CallOption) (*Network, error) {
	out := new(Network)
	err := grpc.Invoke(ctx, "/n0stack.pool.NetworkService/ApplyNetwork", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) DeleteNetwork(ctx context.Context, in *DeleteNetworkRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/n0stack.pool.NetworkService/DeleteNetwork", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) ReserveNetworkInterface(ctx context.Context, in *ReserveNetworkInterfaceRequest, opts ...grpc.CallOption) (*Network, error) {
	out := new(Network)
	err := grpc.Invoke(ctx, "/n0stack.pool.NetworkService/ReserveNetworkInterface", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) ReleaseNetworkInterface(ctx context.Context, in *ReleaseNetworkInterfaceRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/n0stack.pool.NetworkService/ReleaseNetworkInterface", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NetworkService service

type NetworkServiceServer interface {
	ListNetworks(context.Context, *ListNetworksRequest) (*ListNetworksResponse, error)
	GetNetwork(context.Context, *GetNetworkRequest) (*Network, error)
	ApplyNetwork(context.Context, *ApplyNetworkRequest) (*Network, error)
	DeleteNetwork(context.Context, *DeleteNetworkRequest) (*empty.Empty, error)
	ReserveNetworkInterface(context.Context, *ReserveNetworkInterfaceRequest) (*Network, error)
	ReleaseNetworkInterface(context.Context, *ReleaseNetworkInterfaceRequest) (*empty.Empty, error)
}

func RegisterNetworkServiceServer(s *grpc.Server, srv NetworkServiceServer) {
	s.RegisterService(&_NetworkService_serviceDesc, srv)
}

func _NetworkService_ListNetworks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNetworksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).ListNetworks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.pool.NetworkService/ListNetworks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).ListNetworks(ctx, req.(*ListNetworksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_GetNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).GetNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.pool.NetworkService/GetNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).GetNetwork(ctx, req.(*GetNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_ApplyNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).ApplyNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.pool.NetworkService/ApplyNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).ApplyNetwork(ctx, req.(*ApplyNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_DeleteNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).DeleteNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.pool.NetworkService/DeleteNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).DeleteNetwork(ctx, req.(*DeleteNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_ReserveNetworkInterface_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReserveNetworkInterfaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).ReserveNetworkInterface(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.pool.NetworkService/ReserveNetworkInterface",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).ReserveNetworkInterface(ctx, req.(*ReserveNetworkInterfaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_ReleaseNetworkInterface_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseNetworkInterfaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).ReleaseNetworkInterface(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.pool.NetworkService/ReleaseNetworkInterface",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).ReleaseNetworkInterface(ctx, req.(*ReleaseNetworkInterfaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "n0stack.pool.NetworkService",
	HandlerType: (*NetworkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListNetworks",
			Handler:    _NetworkService_ListNetworks_Handler,
		},
		{
			MethodName: "GetNetwork",
			Handler:    _NetworkService_GetNetwork_Handler,
		},
		{
			MethodName: "ApplyNetwork",
			Handler:    _NetworkService_ApplyNetwork_Handler,
		},
		{
			MethodName: "DeleteNetwork",
			Handler:    _NetworkService_DeleteNetwork_Handler,
		},
		{
			MethodName: "ReserveNetworkInterface",
			Handler:    _NetworkService_ReserveNetworkInterface_Handler,
		},
		{
			MethodName: "ReleaseNetworkInterface",
			Handler:    _NetworkService_ReleaseNetworkInterface_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pool/v0/network.proto",
}

func init() { proto.RegisterFile("pool/v0/network.proto", fileDescriptor_network_faeccbe0a81f7f5a) }

var fileDescriptor_network_faeccbe0a81f7f5a = []byte{
	// 780 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x5d, 0x6f, 0x12, 0x4d,
	0x14, 0x2e, 0x50, 0x68, 0x7b, 0xa0, 0x2d, 0x9d, 0x52, 0xba, 0xef, 0xf2, 0xa6, 0xc2, 0x5e, 0x68,
	0x6d, 0xcc, 0x6e, 0xa5, 0x95, 0x34, 0x35, 0x9a, 0x50, 0x8b, 0x4a, 0x24, 0x98, 0xe0, 0x57, 0xe2,
	0x0d, 0x2e, 0xec, 0x94, 0x6e, 0x80, 0xdd, 0x75, 0x67, 0xa0, 0x21, 0xc6, 0x1b, 0x13, 0xaf, 0xbc,
	0x32, 0xfe, 0x32, 0xe3, 0x5f, 0xf0, 0x87, 0x98, 0x9d, 0x19, 0xe8, 0x02, 0x0b, 0xd5, 0x34, 0xe9,
	0x15, 0x3b, 0xe7, 0x3c, 0xe7, 0x39, 0x73, 0x3e, 0x07, 0xd8, 0x72, 0x6c, 0xbb, 0xa3, 0xf5, 0xf7,
	0x35, 0x0b, 0xd3, 0x0b, 0xdb, 0x6d, 0xab, 0x8e, 0x6b, 0x53, 0x1b, 0x25, 0xac, 0x7d, 0x42, 0xf5,
	0x66, 0x5b, 0xf5, 0xd4, 0xf2, 0xff, 0x2d, 0xdb, 0x6e, 0x75, 0xb0, 0xa6, 0x3b, 0xa6, 0xa6, 0x5b,
	0x96, 0x4d, 0x75, 0x6a, 0xda, 0x16, 0xe1, 0x58, 0x39, 0x23, 0xb4, 0xec, 0xd4, 0xe8, 0x9d, 0x69,
	0xb8, 0xeb, 0xd0, 0x81, 0x50, 0xe6, 0x1a, 0x3d, 0xa3, 0x85, 0xa9, 0xcf, 0x43, 0xdd, 0xb4, 0x28,
	0x76, 0xcf, 0xf4, 0x26, 0xe6, 0x10, 0xe5, 0xe7, 0x22, 0x2c, 0x55, 0xb9, 0x0e, 0x21, 0x58, 0xb4,
	0xf4, 0x2e, 0x96, 0x42, 0xd9, 0xd0, 0xee, 0x4a, 0x8d, 0x7d, 0xa3, 0xe7, 0x10, 0xf7, 0x39, 0x95,
	0x22, 0xd9, 0xc8, 0x6e, 0x3c, 0x7f, 0x5b, 0xf5, 0xdf, 0x50, 0x15, 0xf6, 0x6a, 0xf1, 0x12, 0x58,
	0xb2, 0xa8, 0x3b, 0xa8, 0xf9, 0x4d, 0x51, 0x06, 0x56, 0x4c, 0xa7, 0x7f, 0x58, 0x6f, 0x9a, 0x86,
	0x2b, 0x01, 0x73, 0xb1, 0xec, 0x09, 0x9e, 0x98, 0x86, 0x2b, 0x94, 0x05, 0xae, 0x8c, 0x8f, 0x94,
	0x05, 0xa6, 0x4c, 0x43, 0xcc, 0xb0, 0xbb, 0xba, 0x69, 0x49, 0x09, 0xa6, 0x11, 0x27, 0x74, 0x04,
	0x51, 0x42, 0x75, 0x8a, 0xa5, 0x7c, 0x36, 0xb4, 0xbb, 0x96, 0x57, 0x82, 0x6f, 0x25, 0x7e, 0x5f,
	0x79, 0xc8, 0x1a, 0x37, 0x40, 0x14, 0x32, 0x2e, 0x26, 0xd8, 0xed, 0x63, 0xa3, 0x3e, 0x95, 0x19,
	0x22, 0x1d, 0xb0, 0x28, 0x0f, 0x83, 0xf9, 0x6a, 0xc2, 0x50, 0x9c, 0xcb, 0x23, 0x33, 0x1e, 0xf3,
	0x7f, 0xee, 0x2c, 0xbd, 0xfc, 0x18, 0x92, 0x93, 0x29, 0x42, 0x49, 0x88, 0xb4, 0xf1, 0x40, 0xa4,
	0xdc, 0xfb, 0x44, 0x29, 0x88, 0xf6, 0xf5, 0x4e, 0x0f, 0x4b, 0x61, 0x26, 0xe3, 0x87, 0xe3, 0xf0,
	0x51, 0x48, 0xb6, 0x60, 0x67, 0xbe, 0xf3, 0x00, 0xb6, 0x82, 0x9f, 0x2d, 0x9e, 0xcf, 0x8e, 0x62,
	0xe2, 0xad, 0xa1, 0x4e, 0x12, 0xf9, 0xfc, 0x29, 0x05, 0x48, 0xf8, 0x93, 0x87, 0x00, 0x62, 0x4f,
	0x8b, 0xe5, 0x4a, 0xe9, 0x34, 0xb9, 0x80, 0xe2, 0xb0, 0xf4, 0xa6, 0xfa, 0xa2, 0xfa, 0xf2, 0x5d,
	0x35, 0x19, 0x42, 0xab, 0xb0, 0x52, 0x7c, 0x5b, 0x2c, 0x57, 0x8a, 0x27, 0x95, 0x52, 0x32, 0xac,
	0x6c, 0xc1, 0x66, 0xc5, 0x24, 0x54, 0xd8, 0x92, 0x1a, 0xfe, 0xd8, 0xc3, 0x84, 0x2a, 0x65, 0x48,
	0x8d, 0x8b, 0x89, 0x63, 0x5b, 0x04, 0xa3, 0xfb, 0xb0, 0x2c, 0x6a, 0x40, 0xa4, 0x10, 0xcb, 0xfc,
	0x56, 0x60, 0xe6, 0x6b, 0x23, 0x98, 0x72, 0x07, 0x36, 0x9e, 0xe1, 0x21, 0x93, 0xe0, 0x0f, 0x6a,
	0x5f, 0xe5, 0x7b, 0x18, 0x36, 0x8b, 0x8e, 0xd3, 0x19, 0x5c, 0x8d, 0x45, 0xaf, 0x83, 0x5a, 0x3d,
	0x3f, 0x7e, 0x95, 0x00, 0xae, 0x9b, 0x6e, 0xfb, 0xeb, 0xb6, 0x91, 0xb2, 0x07, 0xa9, 0x53, 0xdc,
	0xc1, 0x14, 0xff, 0x45, 0xfe, 0xbe, 0x45, 0x46, 0x3d, 0x37, 0xd5, 0x29, 0xc2, 0x2c, 0x07, 0x89,
	0xe1, 0x08, 0xf9, 0xcc, 0xe3, 0x42, 0x56, 0xf5, 0x32, 0x7b, 0x08, 0xe9, 0xa9, 0x29, 0xe3, 0x60,
	0x7e, 0xb9, 0x94, 0x35, 0xc1, 0xcd, 0xac, 0xea, 0x41, 0xf5, 0x78, 0x34, 0x5e, 0x8f, 0xf9, 0x77,
	0xbb, 0xa2, 0x34, 0x77, 0x21, 0x79, 0xae, 0xbb, 0xc6, 0x85, 0xee, 0xe2, 0xba, 0x6e, 0x18, 0x2e,
	0x26, 0x44, 0x5a, 0x64, 0x17, 0x5a, 0x1f, 0xca, 0x8b, 0x5c, 0xec, 0x05, 0xc9, 0xaa, 0x38, 0x84,
	0x45, 0x79, 0x90, 0x9e, 0x6c, 0x1c, 0x52, 0x18, 0x41, 0x62, 0x23, 0x48, 0x41, 0x40, 0xae, 0x5d,
	0xb9, 0x81, 0x57, 0x8c, 0x0e, 0xd6, 0xc9, 0x8d, 0x17, 0x23, 0xff, 0x35, 0x0a, 0x6b, 0xc3, 0x65,
	0x80, 0xdd, 0xbe, 0xd9, 0xc4, 0xc8, 0x82, 0x84, 0x7f, 0x9e, 0x51, 0x6e, 0xbc, 0x34, 0x01, 0x2b,
	0x40, 0x56, 0xe6, 0x41, 0xf8, 0x3a, 0x50, 0xb6, 0xbf, 0xfc, 0xfa, 0xfd, 0x23, 0xbc, 0x81, 0xd6,
	0xd9, 0x8b, 0x77, 0xf9, 0x74, 0x21, 0x03, 0xe0, 0x72, 0xe8, 0xd1, 0xad, 0x71, 0xaa, 0xa9, 0x75,
	0x20, 0x07, 0x2f, 0x11, 0x65, 0x87, 0xd1, 0x4b, 0x28, 0x3d, 0x41, 0xaf, 0x7d, 0xf2, 0x72, 0xf0,
	0x19, 0xb5, 0x21, 0xe1, 0x1f, 0xf2, 0xc9, 0xa8, 0x02, 0x16, 0xc0, 0x2c, 0x4f, 0x39, 0xe6, 0x29,
	0x23, 0xcf, 0xf0, 0x74, 0x1c, 0xda, 0x43, 0x6d, 0x58, 0x1d, 0x1b, 0x45, 0x34, 0x91, 0xa0, 0xa0,
	0x39, 0x95, 0xd3, 0x2a, 0x7f, 0xf3, 0xd5, 0xe1, 0x9b, 0xaf, 0x96, 0xbc, 0x37, 0x7f, 0x18, 0xd9,
	0xde, 0xac, 0xc8, 0x3e, 0xc0, 0xf6, 0x8c, 0x71, 0x41, 0xf7, 0xfe, 0x65, 0xaa, 0x66, 0xc5, 0xbb,
	0x80, 0x74, 0xcf, 0x43, 0x60, 0x7f, 0x4e, 0x7b, 0x98, 0xd7, 0xc6, 0x33, 0x43, 0x5c, 0x38, 0x79,
	0xf0, 0xfe, 0xa0, 0x65, 0xd2, 0xf3, 0x5e, 0x43, 0x6d, 0xda, 0x5d, 0x4d, 0x70, 0xfa, 0x7e, 0x99,
	0x81, 0xda, 0xb2, 0x35, 0xf1, 0xd7, 0xea, 0xa1, 0xe3, 0x7d, 0x34, 0x62, 0x4c, 0x7e, 0xf0, 0x27,
	0x00, 0x00, 0xff, 0xff, 0x2e, 0xe2, 0xe2, 0x6a, 0x72, 0x09, 0x00, 0x00,
}
