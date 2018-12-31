// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mobile_service.proto

package mobile_service_grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Gateway struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Hostname             string   `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Gateway) Reset()         { *m = Gateway{} }
func (m *Gateway) String() string { return proto.CompactTextString(m) }
func (*Gateway) ProtoMessage()    {}
func (*Gateway) Descriptor() ([]byte, []int) {
	return fileDescriptor_mobile_service_ed1046ca35c1b1b3, []int{0}
}
func (m *Gateway) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Gateway.Unmarshal(m, b)
}
func (m *Gateway) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Gateway.Marshal(b, m, deterministic)
}
func (dst *Gateway) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gateway.Merge(dst, src)
}
func (m *Gateway) XXX_Size() int {
	return xxx_messageInfo_Gateway.Size(m)
}
func (m *Gateway) XXX_DiscardUnknown() {
	xxx_messageInfo_Gateway.DiscardUnknown(m)
}

var xxx_messageInfo_Gateway proto.InternalMessageInfo

func (m *Gateway) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Gateway) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

type Creds struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Creds) Reset()         { *m = Creds{} }
func (m *Creds) String() string { return proto.CompactTextString(m) }
func (*Creds) ProtoMessage()    {}
func (*Creds) Descriptor() ([]byte, []int) {
	return fileDescriptor_mobile_service_ed1046ca35c1b1b3, []int{1}
}
func (m *Creds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Creds.Unmarshal(m, b)
}
func (m *Creds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Creds.Marshal(b, m, deterministic)
}
func (dst *Creds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Creds.Merge(dst, src)
}
func (m *Creds) XXX_Size() int {
	return xxx_messageInfo_Creds.Size(m)
}
func (m *Creds) XXX_DiscardUnknown() {
	xxx_messageInfo_Creds.DiscardUnknown(m)
}

var xxx_messageInfo_Creds proto.InternalMessageInfo

func (m *Creds) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Creds) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Stats struct {
	NumChangesFeeds      string   `protobuf:"bytes,1,opt,name=num_changes_feeds,json=numChangesFeeds,proto3" json:"num_changes_feeds,omitempty"`
	Gateway              *Gateway `protobuf:"bytes,2,opt,name=gateway,proto3" json:"gateway,omitempty"`
	Creds                *Creds   `protobuf:"bytes,3,opt,name=creds,proto3" json:"creds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Stats) Reset()         { *m = Stats{} }
func (m *Stats) String() string { return proto.CompactTextString(m) }
func (*Stats) ProtoMessage()    {}
func (*Stats) Descriptor() ([]byte, []int) {
	return fileDescriptor_mobile_service_ed1046ca35c1b1b3, []int{2}
}
func (m *Stats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stats.Unmarshal(m, b)
}
func (m *Stats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stats.Marshal(b, m, deterministic)
}
func (dst *Stats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stats.Merge(dst, src)
}
func (m *Stats) XXX_Size() int {
	return xxx_messageInfo_Stats.Size(m)
}
func (m *Stats) XXX_DiscardUnknown() {
	xxx_messageInfo_Stats.DiscardUnknown(m)
}

var xxx_messageInfo_Stats proto.InternalMessageInfo

func (m *Stats) GetNumChangesFeeds() string {
	if m != nil {
		return m.NumChangesFeeds
	}
	return ""
}

func (m *Stats) GetGateway() *Gateway {
	if m != nil {
		return m.Gateway
	}
	return nil
}

func (m *Stats) GetCreds() *Creds {
	if m != nil {
		return m.Creds
	}
	return nil
}

type StatsReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatsReply) Reset()         { *m = StatsReply{} }
func (m *StatsReply) String() string { return proto.CompactTextString(m) }
func (*StatsReply) ProtoMessage()    {}
func (*StatsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_mobile_service_ed1046ca35c1b1b3, []int{3}
}
func (m *StatsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsReply.Unmarshal(m, b)
}
func (m *StatsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsReply.Marshal(b, m, deterministic)
}
func (dst *StatsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsReply.Merge(dst, src)
}
func (m *StatsReply) XXX_Size() int {
	return xxx_messageInfo_StatsReply.Size(m)
}
func (m *StatsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsReply.DiscardUnknown(m)
}

var xxx_messageInfo_StatsReply proto.InternalMessageInfo

type MetaKVPath struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetaKVPath) Reset()         { *m = MetaKVPath{} }
func (m *MetaKVPath) String() string { return proto.CompactTextString(m) }
func (*MetaKVPath) ProtoMessage()    {}
func (*MetaKVPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_mobile_service_ed1046ca35c1b1b3, []int{4}
}
func (m *MetaKVPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetaKVPath.Unmarshal(m, b)
}
func (m *MetaKVPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetaKVPath.Marshal(b, m, deterministic)
}
func (dst *MetaKVPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaKVPath.Merge(dst, src)
}
func (m *MetaKVPath) XXX_Size() int {
	return xxx_messageInfo_MetaKVPath.Size(m)
}
func (m *MetaKVPath) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaKVPath.DiscardUnknown(m)
}

var xxx_messageInfo_MetaKVPath proto.InternalMessageInfo

func (m *MetaKVPath) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type MetaKVPair struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Rev                  string   `protobuf:"bytes,3,opt,name=rev,proto3" json:"rev,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetaKVPair) Reset()         { *m = MetaKVPair{} }
func (m *MetaKVPair) String() string { return proto.CompactTextString(m) }
func (*MetaKVPair) ProtoMessage()    {}
func (*MetaKVPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_mobile_service_ed1046ca35c1b1b3, []int{5}
}
func (m *MetaKVPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetaKVPair.Unmarshal(m, b)
}
func (m *MetaKVPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetaKVPair.Marshal(b, m, deterministic)
}
func (dst *MetaKVPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaKVPair.Merge(dst, src)
}
func (m *MetaKVPair) XXX_Size() int {
	return xxx_messageInfo_MetaKVPair.Size(m)
}
func (m *MetaKVPair) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaKVPair.DiscardUnknown(m)
}

var xxx_messageInfo_MetaKVPair proto.InternalMessageInfo

func (m *MetaKVPair) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *MetaKVPair) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *MetaKVPair) GetRev() string {
	if m != nil {
		return m.Rev
	}
	return ""
}

type MetaKVPairs struct {
	Items                []*MetaKVPair `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *MetaKVPairs) Reset()         { *m = MetaKVPairs{} }
func (m *MetaKVPairs) String() string { return proto.CompactTextString(m) }
func (*MetaKVPairs) ProtoMessage()    {}
func (*MetaKVPairs) Descriptor() ([]byte, []int) {
	return fileDescriptor_mobile_service_ed1046ca35c1b1b3, []int{6}
}
func (m *MetaKVPairs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetaKVPairs.Unmarshal(m, b)
}
func (m *MetaKVPairs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetaKVPairs.Marshal(b, m, deterministic)
}
func (dst *MetaKVPairs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaKVPairs.Merge(dst, src)
}
func (m *MetaKVPairs) XXX_Size() int {
	return xxx_messageInfo_MetaKVPairs.Size(m)
}
func (m *MetaKVPairs) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaKVPairs.DiscardUnknown(m)
}

var xxx_messageInfo_MetaKVPairs proto.InternalMessageInfo

func (m *MetaKVPairs) GetItems() []*MetaKVPair {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*Gateway)(nil), "mobile_service_grpc.Gateway")
	proto.RegisterType((*Creds)(nil), "mobile_service_grpc.Creds")
	proto.RegisterType((*Stats)(nil), "mobile_service_grpc.Stats")
	proto.RegisterType((*StatsReply)(nil), "mobile_service_grpc.StatsReply")
	proto.RegisterType((*MetaKVPath)(nil), "mobile_service_grpc.MetaKVPath")
	proto.RegisterType((*MetaKVPair)(nil), "mobile_service_grpc.MetaKVPair")
	proto.RegisterType((*MetaKVPairs)(nil), "mobile_service_grpc.MetaKVPairs")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MobileServiceClient is the client API for MobileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MobileServiceClient interface {
	MetaKVGet(ctx context.Context, in *MetaKVPath, opts ...grpc.CallOption) (*MetaKVPair, error)
	MetaKVListAllChildren(ctx context.Context, in *MetaKVPath, opts ...grpc.CallOption) (*MetaKVPairs, error)
	MetaKVSet(ctx context.Context, in *MetaKVPair, opts ...grpc.CallOption) (*Empty, error)
	MetaKVAdd(ctx context.Context, in *MetaKVPair, opts ...grpc.CallOption) (*Empty, error)
	MetaKVDelete(ctx context.Context, in *MetaKVPair, opts ...grpc.CallOption) (*Empty, error)
	MetaKVRecursiveDelete(ctx context.Context, in *MetaKVPath, opts ...grpc.CallOption) (*Empty, error)
	// Server-side streaming for all changes to MetaKV, stays open indefinitely
	MetaKVObserveChildren(ctx context.Context, in *MetaKVPath, opts ...grpc.CallOption) (MobileService_MetaKVObserveChildrenClient, error)
	// The client uses this to stream stats
	SendStats(ctx context.Context, opts ...grpc.CallOption) (MobileService_SendStatsClient, error)
}

type mobileServiceClient struct {
	cc *grpc.ClientConn
}

func NewMobileServiceClient(cc *grpc.ClientConn) MobileServiceClient {
	return &mobileServiceClient{cc}
}

func (c *mobileServiceClient) MetaKVGet(ctx context.Context, in *MetaKVPath, opts ...grpc.CallOption) (*MetaKVPair, error) {
	out := new(MetaKVPair)
	err := c.cc.Invoke(ctx, "/mobile_service_grpc.MobileService/MetaKVGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileServiceClient) MetaKVListAllChildren(ctx context.Context, in *MetaKVPath, opts ...grpc.CallOption) (*MetaKVPairs, error) {
	out := new(MetaKVPairs)
	err := c.cc.Invoke(ctx, "/mobile_service_grpc.MobileService/MetaKVListAllChildren", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileServiceClient) MetaKVSet(ctx context.Context, in *MetaKVPair, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/mobile_service_grpc.MobileService/MetaKVSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileServiceClient) MetaKVAdd(ctx context.Context, in *MetaKVPair, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/mobile_service_grpc.MobileService/MetaKVAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileServiceClient) MetaKVDelete(ctx context.Context, in *MetaKVPair, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/mobile_service_grpc.MobileService/MetaKVDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileServiceClient) MetaKVRecursiveDelete(ctx context.Context, in *MetaKVPath, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/mobile_service_grpc.MobileService/MetaKVRecursiveDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileServiceClient) MetaKVObserveChildren(ctx context.Context, in *MetaKVPath, opts ...grpc.CallOption) (MobileService_MetaKVObserveChildrenClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MobileService_serviceDesc.Streams[0], "/mobile_service_grpc.MobileService/MetaKVObserveChildren", opts...)
	if err != nil {
		return nil, err
	}
	x := &mobileServiceMetaKVObserveChildrenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MobileService_MetaKVObserveChildrenClient interface {
	Recv() (*MetaKVPair, error)
	grpc.ClientStream
}

type mobileServiceMetaKVObserveChildrenClient struct {
	grpc.ClientStream
}

func (x *mobileServiceMetaKVObserveChildrenClient) Recv() (*MetaKVPair, error) {
	m := new(MetaKVPair)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mobileServiceClient) SendStats(ctx context.Context, opts ...grpc.CallOption) (MobileService_SendStatsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MobileService_serviceDesc.Streams[1], "/mobile_service_grpc.MobileService/SendStats", opts...)
	if err != nil {
		return nil, err
	}
	x := &mobileServiceSendStatsClient{stream}
	return x, nil
}

type MobileService_SendStatsClient interface {
	Send(*Stats) error
	CloseAndRecv() (*StatsReply, error)
	grpc.ClientStream
}

type mobileServiceSendStatsClient struct {
	grpc.ClientStream
}

func (x *mobileServiceSendStatsClient) Send(m *Stats) error {
	return x.ClientStream.SendMsg(m)
}

func (x *mobileServiceSendStatsClient) CloseAndRecv() (*StatsReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StatsReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MobileServiceServer is the server API for MobileService service.
type MobileServiceServer interface {
	MetaKVGet(context.Context, *MetaKVPath) (*MetaKVPair, error)
	MetaKVListAllChildren(context.Context, *MetaKVPath) (*MetaKVPairs, error)
	MetaKVSet(context.Context, *MetaKVPair) (*Empty, error)
	MetaKVAdd(context.Context, *MetaKVPair) (*Empty, error)
	MetaKVDelete(context.Context, *MetaKVPair) (*Empty, error)
	MetaKVRecursiveDelete(context.Context, *MetaKVPath) (*Empty, error)
	// Server-side streaming for all changes to MetaKV, stays open indefinitely
	MetaKVObserveChildren(*MetaKVPath, MobileService_MetaKVObserveChildrenServer) error
	// The client uses this to stream stats
	SendStats(MobileService_SendStatsServer) error
}

func RegisterMobileServiceServer(s *grpc.Server, srv MobileServiceServer) {
	s.RegisterService(&_MobileService_serviceDesc, srv)
}

func _MobileService_MetaKVGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetaKVPath)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileServiceServer).MetaKVGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mobile_service_grpc.MobileService/MetaKVGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileServiceServer).MetaKVGet(ctx, req.(*MetaKVPath))
	}
	return interceptor(ctx, in, info, handler)
}

func _MobileService_MetaKVListAllChildren_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetaKVPath)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileServiceServer).MetaKVListAllChildren(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mobile_service_grpc.MobileService/MetaKVListAllChildren",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileServiceServer).MetaKVListAllChildren(ctx, req.(*MetaKVPath))
	}
	return interceptor(ctx, in, info, handler)
}

func _MobileService_MetaKVSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetaKVPair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileServiceServer).MetaKVSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mobile_service_grpc.MobileService/MetaKVSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileServiceServer).MetaKVSet(ctx, req.(*MetaKVPair))
	}
	return interceptor(ctx, in, info, handler)
}

func _MobileService_MetaKVAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetaKVPair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileServiceServer).MetaKVAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mobile_service_grpc.MobileService/MetaKVAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileServiceServer).MetaKVAdd(ctx, req.(*MetaKVPair))
	}
	return interceptor(ctx, in, info, handler)
}

func _MobileService_MetaKVDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetaKVPair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileServiceServer).MetaKVDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mobile_service_grpc.MobileService/MetaKVDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileServiceServer).MetaKVDelete(ctx, req.(*MetaKVPair))
	}
	return interceptor(ctx, in, info, handler)
}

func _MobileService_MetaKVRecursiveDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetaKVPath)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileServiceServer).MetaKVRecursiveDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mobile_service_grpc.MobileService/MetaKVRecursiveDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileServiceServer).MetaKVRecursiveDelete(ctx, req.(*MetaKVPath))
	}
	return interceptor(ctx, in, info, handler)
}

func _MobileService_MetaKVObserveChildren_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MetaKVPath)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MobileServiceServer).MetaKVObserveChildren(m, &mobileServiceMetaKVObserveChildrenServer{stream})
}

type MobileService_MetaKVObserveChildrenServer interface {
	Send(*MetaKVPair) error
	grpc.ServerStream
}

type mobileServiceMetaKVObserveChildrenServer struct {
	grpc.ServerStream
}

func (x *mobileServiceMetaKVObserveChildrenServer) Send(m *MetaKVPair) error {
	return x.ServerStream.SendMsg(m)
}

func _MobileService_SendStats_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MobileServiceServer).SendStats(&mobileServiceSendStatsServer{stream})
}

type MobileService_SendStatsServer interface {
	SendAndClose(*StatsReply) error
	Recv() (*Stats, error)
	grpc.ServerStream
}

type mobileServiceSendStatsServer struct {
	grpc.ServerStream
}

func (x *mobileServiceSendStatsServer) SendAndClose(m *StatsReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *mobileServiceSendStatsServer) Recv() (*Stats, error) {
	m := new(Stats)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MobileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mobile_service_grpc.MobileService",
	HandlerType: (*MobileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MetaKVGet",
			Handler:    _MobileService_MetaKVGet_Handler,
		},
		{
			MethodName: "MetaKVListAllChildren",
			Handler:    _MobileService_MetaKVListAllChildren_Handler,
		},
		{
			MethodName: "MetaKVSet",
			Handler:    _MobileService_MetaKVSet_Handler,
		},
		{
			MethodName: "MetaKVAdd",
			Handler:    _MobileService_MetaKVAdd_Handler,
		},
		{
			MethodName: "MetaKVDelete",
			Handler:    _MobileService_MetaKVDelete_Handler,
		},
		{
			MethodName: "MetaKVRecursiveDelete",
			Handler:    _MobileService_MetaKVRecursiveDelete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MetaKVObserveChildren",
			Handler:       _MobileService_MetaKVObserveChildren_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendStats",
			Handler:       _MobileService_SendStats_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "mobile_service.proto",
}

func init() {
	proto.RegisterFile("mobile_service.proto", fileDescriptor_mobile_service_ed1046ca35c1b1b3)
}

var fileDescriptor_mobile_service_ed1046ca35c1b1b3 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xd1, 0x8a, 0x13, 0x31,
	0x14, 0x86, 0x3b, 0x5b, 0xc7, 0xb5, 0xa7, 0x15, 0x35, 0xae, 0x50, 0x06, 0xc1, 0x92, 0xab, 0xe2,
	0x45, 0x59, 0x2a, 0x0a, 0x5e, 0xc9, 0xd2, 0xd5, 0x15, 0xdd, 0xb2, 0x32, 0x85, 0xbd, 0x52, 0x4a,
	0x3a, 0x39, 0x76, 0x02, 0x93, 0x99, 0x21, 0xc9, 0x74, 0xe9, 0xc3, 0xf8, 0x02, 0x3e, 0xa5, 0x24,
	0xd9, 0x69, 0x57, 0x18, 0x9d, 0x85, 0xde, 0xe5, 0x9c, 0xfc, 0xff, 0x37, 0x7f, 0x4e, 0xc8, 0xc0,
	0x89, 0x2c, 0x56, 0x22, 0xc3, 0xa5, 0x46, 0xb5, 0x11, 0x09, 0x4e, 0x4a, 0x55, 0x98, 0x82, 0x3c,
	0xff, 0xbb, 0xbb, 0x5c, 0xab, 0x32, 0x89, 0x06, 0x49, 0x21, 0x65, 0x91, 0x7b, 0x09, 0x7d, 0x0f,
	0xc7, 0x17, 0xcc, 0xe0, 0x0d, 0xdb, 0x12, 0x02, 0x0f, 0xaa, 0x4a, 0xf0, 0x61, 0x30, 0x0a, 0xc6,
	0xbd, 0xd8, 0xad, 0x49, 0x04, 0x8f, 0xd2, 0x42, 0x9b, 0x9c, 0x49, 0x1c, 0x1e, 0xb9, 0xfe, 0xae,
	0xa6, 0x1f, 0x20, 0x9c, 0x29, 0xe4, 0xda, 0x8a, 0x2a, 0x8d, 0xca, 0x89, 0xbc, 0x79, 0x57, 0xdb,
	0xbd, 0x92, 0x69, 0x7d, 0x53, 0x28, 0x5e, 0x03, 0xea, 0x9a, 0xfe, 0x0a, 0x20, 0x5c, 0x18, 0x66,
	0x34, 0x79, 0x0d, 0xcf, 0xf2, 0x4a, 0x2e, 0x93, 0x94, 0xe5, 0x6b, 0xd4, 0xcb, 0x9f, 0x88, 0x5c,
	0xdf, 0xa2, 0x9e, 0xe4, 0x95, 0x9c, 0xf9, 0xfe, 0x27, 0xdb, 0x26, 0xef, 0xe0, 0x78, 0xed, 0x13,
	0x3b, 0x60, 0x7f, 0xfa, 0x72, 0xd2, 0x70, 0xcc, 0xc9, 0xed, 0xa9, 0xe2, 0x5a, 0x4c, 0x4e, 0x21,
	0x4c, 0x6c, 0xdc, 0x61, 0xd7, 0xb9, 0xa2, 0x46, 0x97, 0x3b, 0x50, 0xec, 0x85, 0x74, 0x00, 0xe0,
	0xe2, 0xc5, 0x58, 0x66, 0x5b, 0x3a, 0x02, 0x98, 0xa3, 0x61, 0x5f, 0xaf, 0xbf, 0x31, 0x93, 0xda,
	0x61, 0x95, 0xcc, 0xa4, 0xf5, 0xb0, 0xec, 0x9a, 0x7e, 0xde, 0x2b, 0x84, 0x22, 0x27, 0x10, 0x6e,
	0x58, 0x56, 0xf9, 0x91, 0x0c, 0x62, 0x5f, 0xec, 0x7c, 0x47, 0x7b, 0x1f, 0x79, 0x0a, 0x5d, 0x85,
	0x1b, 0x97, 0xab, 0x17, 0xdb, 0x25, 0x3d, 0x87, 0xfe, 0x9e, 0xa4, 0xc9, 0x5b, 0x08, 0x85, 0x41,
	0x69, 0x47, 0xd2, 0x1d, 0xf7, 0xa7, 0xaf, 0x1a, 0xa3, 0xef, 0x0d, 0xb1, 0x57, 0x4f, 0x7f, 0x87,
	0xf0, 0x78, 0xee, 0x94, 0x0b, 0x2f, 0x24, 0x57, 0xd0, 0xf3, 0xb2, 0x0b, 0x34, 0xe4, 0xff, 0x18,
	0x93, 0x46, 0x6d, 0xdf, 0xa1, 0x1d, 0xf2, 0x1d, 0x5e, 0xf8, 0xfa, 0x52, 0x68, 0x73, 0x96, 0x65,
	0xb3, 0x54, 0x64, 0x5c, 0x61, 0xde, 0x0e, 0x1f, 0xb5, 0xc0, 0x35, 0xed, 0x90, 0x2f, 0x75, 0xdc,
	0x45, 0x6b, 0x5c, 0xa1, 0xa2, 0xe6, 0x1b, 0xfd, 0x28, 0x4b, 0xb3, 0xbd, 0xcb, 0x3a, 0xe3, 0xfc,
	0x50, 0xd6, 0x1c, 0x06, 0x5e, 0x7b, 0x8e, 0x19, 0x1a, 0x3c, 0x14, 0x77, 0x5d, 0x0f, 0x31, 0xc6,
	0xa4, 0x52, 0x5a, 0x6c, 0xf0, 0x5e, 0x5c, 0x93, 0xb6, 0x70, 0x7f, 0xd4, 0xdc, 0xab, 0x95, 0xdd,
	0xc7, 0xfb, 0x5f, 0x4e, 0xfb, 0xcd, 0x9f, 0x06, 0xe4, 0x12, 0x7a, 0x0b, 0xcc, 0xb9, 0x7f, 0xc1,
	0xcd, 0x49, 0xdc, 0xde, 0x3f, 0x68, 0x77, 0x9e, 0x56, 0x67, 0x1c, 0xac, 0x1e, 0xba, 0xff, 0xd1,
	0x9b, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x65, 0xeb, 0xd0, 0x52, 0xca, 0x04, 0x00, 0x00,
}
