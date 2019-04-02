// Code generated by protoc-gen-go. DO NOT EDIT.
// source: voltha_protos/afrouter.proto

package afrouter

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Result struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_be7e2f565b9e1c96, []int{0}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Result) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_be7e2f565b9e1c96, []int{1}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Count struct {
	Count                uint32   `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Count) Reset()         { *m = Count{} }
func (m *Count) String() string { return proto.CompactTextString(m) }
func (*Count) ProtoMessage()    {}
func (*Count) Descriptor() ([]byte, []int) {
	return fileDescriptor_be7e2f565b9e1c96, []int{2}
}

func (m *Count) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Count.Unmarshal(m, b)
}
func (m *Count) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Count.Marshal(b, m, deterministic)
}
func (m *Count) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Count.Merge(m, src)
}
func (m *Count) XXX_Size() int {
	return xxx_messageInfo_Count.Size(m)
}
func (m *Count) XXX_DiscardUnknown() {
	xxx_messageInfo_Count.DiscardUnknown(m)
}

var xxx_messageInfo_Count proto.InternalMessageInfo

func (m *Count) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Conn struct {
	Server               string   `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Pkg                  string   `protobuf:"bytes,2,opt,name=pkg,proto3" json:"pkg,omitempty"`
	Svc                  string   `protobuf:"bytes,3,opt,name=svc,proto3" json:"svc,omitempty"`
	Cluster              string   `protobuf:"bytes,4,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Backend              string   `protobuf:"bytes,5,opt,name=backend,proto3" json:"backend,omitempty"`
	Connection           string   `protobuf:"bytes,6,opt,name=connection,proto3" json:"connection,omitempty"`
	Addr                 string   `protobuf:"bytes,7,opt,name=addr,proto3" json:"addr,omitempty"`
	Port                 uint64   `protobuf:"varint,8,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Conn) Reset()         { *m = Conn{} }
func (m *Conn) String() string { return proto.CompactTextString(m) }
func (*Conn) ProtoMessage()    {}
func (*Conn) Descriptor() ([]byte, []int) {
	return fileDescriptor_be7e2f565b9e1c96, []int{3}
}

func (m *Conn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Conn.Unmarshal(m, b)
}
func (m *Conn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Conn.Marshal(b, m, deterministic)
}
func (m *Conn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Conn.Merge(m, src)
}
func (m *Conn) XXX_Size() int {
	return xxx_messageInfo_Conn.Size(m)
}
func (m *Conn) XXX_DiscardUnknown() {
	xxx_messageInfo_Conn.DiscardUnknown(m)
}

var xxx_messageInfo_Conn proto.InternalMessageInfo

func (m *Conn) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func (m *Conn) GetPkg() string {
	if m != nil {
		return m.Pkg
	}
	return ""
}

func (m *Conn) GetSvc() string {
	if m != nil {
		return m.Svc
	}
	return ""
}

func (m *Conn) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *Conn) GetBackend() string {
	if m != nil {
		return m.Backend
	}
	return ""
}

func (m *Conn) GetConnection() string {
	if m != nil {
		return m.Connection
	}
	return ""
}

func (m *Conn) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *Conn) GetPort() uint64 {
	if m != nil {
		return m.Port
	}
	return 0
}

type Affinity struct {
	Router               string   `protobuf:"bytes,1,opt,name=router,proto3" json:"router,omitempty"`
	Route                string   `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
	Cluster              string   `protobuf:"bytes,3,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Backend              string   `protobuf:"bytes,4,opt,name=backend,proto3" json:"backend,omitempty"`
	Id                   string   `protobuf:"bytes,5,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Affinity) Reset()         { *m = Affinity{} }
func (m *Affinity) String() string { return proto.CompactTextString(m) }
func (*Affinity) ProtoMessage()    {}
func (*Affinity) Descriptor() ([]byte, []int) {
	return fileDescriptor_be7e2f565b9e1c96, []int{4}
}

func (m *Affinity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Affinity.Unmarshal(m, b)
}
func (m *Affinity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Affinity.Marshal(b, m, deterministic)
}
func (m *Affinity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Affinity.Merge(m, src)
}
func (m *Affinity) XXX_Size() int {
	return xxx_messageInfo_Affinity.Size(m)
}
func (m *Affinity) XXX_DiscardUnknown() {
	xxx_messageInfo_Affinity.DiscardUnknown(m)
}

var xxx_messageInfo_Affinity proto.InternalMessageInfo

func (m *Affinity) GetRouter() string {
	if m != nil {
		return m.Router
	}
	return ""
}

func (m *Affinity) GetRoute() string {
	if m != nil {
		return m.Route
	}
	return ""
}

func (m *Affinity) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *Affinity) GetBackend() string {
	if m != nil {
		return m.Backend
	}
	return ""
}

func (m *Affinity) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Result)(nil), "afrouter.Result")
	proto.RegisterType((*Empty)(nil), "afrouter.Empty")
	proto.RegisterType((*Count)(nil), "afrouter.Count")
	proto.RegisterType((*Conn)(nil), "afrouter.Conn")
	proto.RegisterType((*Affinity)(nil), "afrouter.Affinity")
}

func init() { proto.RegisterFile("voltha_protos/afrouter.proto", fileDescriptor_be7e2f565b9e1c96) }

var fileDescriptor_be7e2f565b9e1c96 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x41, 0x8b, 0xdb, 0x30,
	0x10, 0x85, 0xeb, 0xc4, 0x76, 0x9c, 0x29, 0x09, 0x41, 0x94, 0x22, 0x4a, 0x5b, 0x82, 0x4f, 0xb9,
	0x24, 0x86, 0x86, 0x42, 0xaf, 0xad, 0xe9, 0x1f, 0x70, 0x6e, 0x7b, 0x59, 0x1c, 0x59, 0x76, 0x4c,
	0x12, 0x8d, 0x91, 0x64, 0x43, 0xd8, 0xfd, 0x6d, 0xfb, 0xdb, 0x16, 0x49, 0x76, 0x92, 0x85, 0xdd,
	0xdb, 0x7c, 0xef, 0x59, 0xe6, 0x3d, 0x69, 0xe0, 0x7b, 0x87, 0x27, 0x7d, 0xc8, 0x1f, 0x1b, 0x89,
	0x1a, 0x55, 0x92, 0x97, 0x12, 0x5b, 0xcd, 0xe5, 0xc6, 0x32, 0x89, 0x06, 0x8e, 0xff, 0x40, 0x98,
	0x71, 0xd5, 0x9e, 0x34, 0xa1, 0x30, 0x51, 0x2d, 0x63, 0x5c, 0x29, 0xea, 0x2d, 0xbd, 0x55, 0x94,
	0x0d, 0x48, 0xbe, 0x40, 0xc0, 0xa5, 0x44, 0x49, 0x47, 0x4b, 0x6f, 0x35, 0xcd, 0x1c, 0xc4, 0x13,
	0x08, 0xfe, 0x9f, 0x1b, 0x7d, 0x89, 0x7f, 0x40, 0x90, 0x62, 0x2b, 0xb4, 0xf9, 0x8e, 0x99, 0xc1,
	0x9e, 0x9f, 0x65, 0x0e, 0xe2, 0x17, 0x0f, 0xfc, 0x14, 0x85, 0x20, 0x5f, 0x21, 0x54, 0x5c, 0x76,
	0x5c, 0x5a, 0x7f, 0x9a, 0xf5, 0x44, 0x16, 0x30, 0x6e, 0x8e, 0x55, 0xff, 0x73, 0x33, 0x1a, 0x45,
	0x75, 0x8c, 0x8e, 0x9d, 0xa2, 0x3a, 0x66, 0xc2, 0xb1, 0x53, 0xab, 0x34, 0x97, 0xd4, 0xb7, 0xea,
	0x80, 0xc6, 0xd9, 0xe7, 0xec, 0xc8, 0x45, 0x41, 0x03, 0xe7, 0xf4, 0x48, 0x7e, 0x02, 0x30, 0x14,
	0x82, 0x33, 0x5d, 0xa3, 0xa0, 0xa1, 0x35, 0xef, 0x14, 0x42, 0xc0, 0xcf, 0x8b, 0x42, 0xd2, 0x89,
	0x75, 0xec, 0x6c, 0xb4, 0x06, 0xa5, 0xa6, 0xd1, 0xd2, 0x5b, 0xf9, 0x99, 0x9d, 0xe3, 0x67, 0x88,
	0xfe, 0x96, 0x65, 0x2d, 0x6a, 0x7d, 0x31, 0x1d, 0xdc, 0xc5, 0x0d, 0x1d, 0x1c, 0x99, 0xea, 0x76,
	0x1a, 0xae, 0xc8, 0xc2, 0x7d, 0xea, 0xf1, 0x87, 0xa9, 0xfd, 0xb7, 0xa9, 0xe7, 0x30, 0xaa, 0x87,
	0x2a, 0xa3, 0xba, 0xf8, 0xf5, 0x04, 0xb3, 0x14, 0x45, 0x59, 0x57, 0xad, 0xcc, 0x6d, 0xec, 0x2d,
	0xcc, 0x76, 0x5c, 0xa7, 0xb7, 0x1e, 0xf3, 0xcd, 0xf5, 0x75, 0x8d, 0xfa, 0x6d, 0x71, 0x63, 0xf7,
	0xb4, 0xf1, 0x27, 0xf2, 0x1b, 0x3e, 0xef, 0xb8, 0xbe, 0xd6, 0x20, 0xb7, 0x4f, 0x06, 0xed, 0xbd,
	0x63, 0xff, 0x92, 0x87, 0x75, 0x55, 0xeb, 0x43, 0xbb, 0xdf, 0x30, 0x3c, 0x27, 0xd8, 0x70, 0xc1,
	0x50, 0x16, 0x89, 0xdb, 0xad, 0x75, 0xbf, 0x5b, 0x15, 0x5e, 0xd7, 0x6b, 0x1f, 0x5a, 0x6d, 0xfb,
	0x1a, 0x00, 0x00, 0xff, 0xff, 0x40, 0x18, 0x4d, 0xa7, 0x7f, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConfigurationClient is the client API for Configuration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConfigurationClient interface {
	SetConnection(ctx context.Context, in *Conn, opts ...grpc.CallOption) (*Result, error)
	SetAffinity(ctx context.Context, in *Affinity, opts ...grpc.CallOption) (*Result, error)
}

type configurationClient struct {
	cc *grpc.ClientConn
}

func NewConfigurationClient(cc *grpc.ClientConn) ConfigurationClient {
	return &configurationClient{cc}
}

func (c *configurationClient) SetConnection(ctx context.Context, in *Conn, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/afrouter.Configuration/SetConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configurationClient) SetAffinity(ctx context.Context, in *Affinity, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/afrouter.Configuration/SetAffinity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigurationServer is the server API for Configuration service.
type ConfigurationServer interface {
	SetConnection(context.Context, *Conn) (*Result, error)
	SetAffinity(context.Context, *Affinity) (*Result, error)
}

// UnimplementedConfigurationServer can be embedded to have forward compatible implementations.
type UnimplementedConfigurationServer struct {
}

func (*UnimplementedConfigurationServer) SetConnection(ctx context.Context, req *Conn) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetConnection not implemented")
}
func (*UnimplementedConfigurationServer) SetAffinity(ctx context.Context, req *Affinity) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAffinity not implemented")
}

func RegisterConfigurationServer(s *grpc.Server, srv ConfigurationServer) {
	s.RegisterService(&_Configuration_serviceDesc, srv)
}

func _Configuration_SetConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Conn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServer).SetConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/afrouter.Configuration/SetConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServer).SetConnection(ctx, req.(*Conn))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configuration_SetAffinity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Affinity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServer).SetAffinity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/afrouter.Configuration/SetAffinity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServer).SetAffinity(ctx, req.(*Affinity))
	}
	return interceptor(ctx, in, info, handler)
}

var _Configuration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "afrouter.Configuration",
	HandlerType: (*ConfigurationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetConnection",
			Handler:    _Configuration_SetConnection_Handler,
		},
		{
			MethodName: "SetAffinity",
			Handler:    _Configuration_SetAffinity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "voltha_protos/afrouter.proto",
}
