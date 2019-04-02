// Code generated by protoc-gen-go. DO NOT EDIT.
// source: voltha_protos/health.proto

package voltha

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/opencord/voltha-protos/go/common"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Health states
type HealthStatus_HealthState int32

const (
	HealthStatus_HEALTHY    HealthStatus_HealthState = 0
	HealthStatus_OVERLOADED HealthStatus_HealthState = 1
	HealthStatus_DYING      HealthStatus_HealthState = 2
)

var HealthStatus_HealthState_name = map[int32]string{
	0: "HEALTHY",
	1: "OVERLOADED",
	2: "DYING",
}

var HealthStatus_HealthState_value = map[string]int32{
	"HEALTHY":    0,
	"OVERLOADED": 1,
	"DYING":      2,
}

func (x HealthStatus_HealthState) String() string {
	return proto.EnumName(HealthStatus_HealthState_name, int32(x))
}

func (HealthStatus_HealthState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dd1fc2b2d96d69b8, []int{0, 0}
}

// Encode health status of a Voltha instance
type HealthStatus struct {
	// Current state of health of this Voltha instance
	State                HealthStatus_HealthState `protobuf:"varint,1,opt,name=state,proto3,enum=voltha.HealthStatus_HealthState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *HealthStatus) Reset()         { *m = HealthStatus{} }
func (m *HealthStatus) String() string { return proto.CompactTextString(m) }
func (*HealthStatus) ProtoMessage()    {}
func (*HealthStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd1fc2b2d96d69b8, []int{0}
}

func (m *HealthStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthStatus.Unmarshal(m, b)
}
func (m *HealthStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthStatus.Marshal(b, m, deterministic)
}
func (m *HealthStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthStatus.Merge(m, src)
}
func (m *HealthStatus) XXX_Size() int {
	return xxx_messageInfo_HealthStatus.Size(m)
}
func (m *HealthStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthStatus.DiscardUnknown(m)
}

var xxx_messageInfo_HealthStatus proto.InternalMessageInfo

func (m *HealthStatus) GetState() HealthStatus_HealthState {
	if m != nil {
		return m.State
	}
	return HealthStatus_HEALTHY
}

func init() {
	proto.RegisterEnum("voltha.HealthStatus_HealthState", HealthStatus_HealthState_name, HealthStatus_HealthState_value)
	proto.RegisterType((*HealthStatus)(nil), "voltha.HealthStatus")
}

func init() { proto.RegisterFile("voltha_protos/health.proto", fileDescriptor_dd1fc2b2d96d69b8) }

var fileDescriptor_dd1fc2b2d96d69b8 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0xcb, 0xcf, 0x29,
	0xc9, 0x48, 0x8c, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x2f, 0xd6, 0xcf, 0x48, 0x4d, 0xcc, 0x29, 0xc9,
	0xd0, 0x03, 0xf3, 0x84, 0xd8, 0x20, 0x72, 0x52, 0x32, 0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0xfa,
	0x89, 0x05, 0x99, 0xfa, 0x89, 0x79, 0x79, 0xf9, 0x25, 0x89, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0x10,
	0x55, 0x52, 0xd2, 0x50, 0x59, 0x30, 0x2f, 0xa9, 0x34, 0x4d, 0x3f, 0x35, 0xb7, 0xa0, 0xa4, 0x12,
	0x2a, 0x29, 0x81, 0x6a, 0x7c, 0x6e, 0x6a, 0x49, 0x22, 0x44, 0x46, 0xa9, 0x85, 0x91, 0x8b, 0xc7,
	0x03, 0x6c, 0x5b, 0x70, 0x49, 0x62, 0x49, 0x69, 0xb1, 0x90, 0x2d, 0x17, 0x6b, 0x71, 0x49, 0x62,
	0x49, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x9f, 0x91, 0x82, 0x1e, 0x44, 0xab, 0x1e, 0xb2, 0x22,
	0x24, 0x4e, 0xaa, 0x13, 0xeb, 0x8b, 0x6f, 0x67, 0x65, 0x19, 0x83, 0x20, 0xba, 0x94, 0x4c, 0xb9,
	0xb8, 0x91, 0x24, 0x85, 0xb8, 0xb9, 0xd8, 0x3d, 0x5c, 0x1d, 0x7d, 0x42, 0x3c, 0x22, 0x05, 0x18,
	0x84, 0xf8, 0xb8, 0xb8, 0xfc, 0xc3, 0x5c, 0x83, 0x7c, 0xfc, 0x1d, 0x5d, 0x5c, 0x5d, 0x04, 0x18,
	0x85, 0x38, 0xb9, 0x58, 0x5d, 0x22, 0x3d, 0xfd, 0xdc, 0x05, 0x98, 0x8c, 0x12, 0xb9, 0x78, 0xa1,
	0xda, 0x52, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0x02, 0xb8, 0xf8, 0xdd, 0x53, 0x4b, 0x50, 0x5c,
	0x26, 0xa6, 0x07, 0xf1, 0xa2, 0x1e, 0xcc, 0x8b, 0x7a, 0xae, 0x20, 0x2f, 0x4a, 0x89, 0x60, 0x73,
	0xa2, 0x12, 0x7f, 0xd3, 0xe5, 0x27, 0x93, 0x99, 0x38, 0x85, 0xd8, 0xa1, 0x81, 0xe9, 0xa4, 0x1b,
	0xa5, 0x9d, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x5f, 0x90, 0x9a,
	0x97, 0x9c, 0x5f, 0x94, 0xa2, 0x0f, 0xd1, 0xab, 0x0b, 0x0d, 0x99, 0xf4, 0x7c, 0xa8, 0x40, 0x12,
	0x1b, 0x58, 0xc4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x06, 0x8f, 0xa7, 0x7d, 0x9a, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HealthServiceClient is the client API for HealthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HealthServiceClient interface {
	// Return current health status of a Voltha instance
	GetHealthStatus(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HealthStatus, error)
}

type healthServiceClient struct {
	cc *grpc.ClientConn
}

func NewHealthServiceClient(cc *grpc.ClientConn) HealthServiceClient {
	return &healthServiceClient{cc}
}

func (c *healthServiceClient) GetHealthStatus(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HealthStatus, error) {
	out := new(HealthStatus)
	err := c.cc.Invoke(ctx, "/voltha.HealthService/GetHealthStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthServiceServer is the server API for HealthService service.
type HealthServiceServer interface {
	// Return current health status of a Voltha instance
	GetHealthStatus(context.Context, *empty.Empty) (*HealthStatus, error)
}

// UnimplementedHealthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHealthServiceServer struct {
}

func (*UnimplementedHealthServiceServer) GetHealthStatus(ctx context.Context, req *empty.Empty) (*HealthStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHealthStatus not implemented")
}

func RegisterHealthServiceServer(s *grpc.Server, srv HealthServiceServer) {
	s.RegisterService(&_HealthService_serviceDesc, srv)
}

func _HealthService_GetHealthStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServiceServer).GetHealthStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voltha.HealthService/GetHealthStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServiceServer).GetHealthStatus(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _HealthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "voltha.HealthService",
	HandlerType: (*HealthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHealthStatus",
			Handler:    _HealthService_GetHealthStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "voltha_protos/health.proto",
}
