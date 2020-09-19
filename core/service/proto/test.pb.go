// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package proto // import "."

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

type EState int32

const (
	EState_Stopped EState = 0
	EState_Normal  EState = 1
)

var EState_name = map[int32]string{
	0: "Stopped",
	1: "Normal",
}
var EState_value = map[string]int32{
	"Stopped": 0,
	"Normal":  1,
}

func (x EState) String() string {
	return proto.EnumName(EState_name, int32(x))
}
func (EState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_test_0660f64e8b550f38, []int{0}
}

type User1 struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	GroupId              int64             `protobuf:"zigzag64,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Extra                map[string]string `protobuf:"bytes,3,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Roles                []string          `protobuf:"bytes,4,rep,name=roles,proto3" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *User1) Reset()         { *m = User1{} }
func (m *User1) String() string { return proto.CompactTextString(m) }
func (*User1) ProtoMessage()    {}
func (*User1) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_0660f64e8b550f38, []int{0}
}
func (m *User1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User1.Unmarshal(m, b)
}
func (m *User1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User1.Marshal(b, m, deterministic)
}
func (dst *User1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User1.Merge(dst, src)
}
func (m *User1) XXX_Size() int {
	return xxx_messageInfo_User1.Size(m)
}
func (m *User1) XXX_DiscardUnknown() {
	xxx_messageInfo_User1.DiscardUnknown(m)
}

var xxx_messageInfo_User1 proto.InternalMessageInfo

func (m *User1) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User1) GetGroupId() int64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *User1) GetExtra() map[string]string {
	if m != nil {
		return m.Extra
	}
	return nil
}

func (m *User1) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

type UserResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	State                EState   `protobuf:"varint,2,opt,name=state,proto3,enum=EState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_0660f64e8b550f38, []int{1}
}
func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (dst *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(dst, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserResponse) GetState() EState {
	if m != nil {
		return m.State
	}
	return EState_Stopped
}

func init() {
	proto.RegisterType((*User1)(nil), "User1")
	proto.RegisterMapType((map[string]string)(nil), "User1.ExtraEntry")
	proto.RegisterType((*UserResponse)(nil), "UserResponse")
	proto.RegisterEnum("EState", EState_name, EState_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterServiceClient is the client API for GreeterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterServiceClient interface {
	Hello(ctx context.Context, in *User1, opts ...grpc.CallOption) (*UserResponse, error)
}

type greeterServiceClient struct {
	cc *grpc.ClientConn
}

func NewGreeterServiceClient(cc *grpc.ClientConn) GreeterServiceClient {
	return &greeterServiceClient{cc}
}

func (c *greeterServiceClient) Hello(ctx context.Context, in *User1, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/GreeterService/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServiceServer is the server API for GreeterService service.
type GreeterServiceServer interface {
	Hello(context.Context, *User1) (*UserResponse, error)
}

func RegisterGreeterServiceServer(s *grpc.Server, srv GreeterServiceServer) {
	s.RegisterService(&_GreeterService_serviceDesc, srv)
}

func _GreeterService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GreeterService/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServiceServer).Hello(ctx, req.(*User1))
	}
	return interceptor(ctx, in, info, handler)
}

var _GreeterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "GreeterService",
	HandlerType: (*GreeterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _GreeterService_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_test_0660f64e8b550f38) }

var fileDescriptor_test_0660f64e8b550f38 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x97, 0x75, 0x6d, 0xed, 0x9b, 0x8e, 0xf9, 0xf0, 0x50, 0x07, 0x42, 0xed, 0xc5, 0xe2,
	0x21, 0x60, 0xbd, 0x0c, 0x3d, 0x29, 0x14, 0xf5, 0xe2, 0x21, 0xc5, 0x8b, 0x17, 0xa9, 0xf6, 0x21,
	0xc3, 0xae, 0x29, 0x49, 0x36, 0xdc, 0x5f, 0xe5, 0xbf, 0x28, 0x49, 0x04, 0x3d, 0x78, 0xca, 0xfb,
	0x7d, 0xc9, 0x97, 0xef, 0x4b, 0x00, 0x0c, 0x69, 0xc3, 0x07, 0x25, 0x8d, 0xcc, 0xbf, 0x18, 0x84,
	0x4f, 0x9a, 0xd4, 0x05, 0x22, 0x4c, 0xfa, 0x66, 0x4d, 0x29, 0xcb, 0x58, 0x91, 0x08, 0x37, 0xe3,
	0x31, 0xec, 0xbd, 0x2b, 0xb9, 0x19, 0x5e, 0x56, 0x6d, 0x3a, 0xce, 0x58, 0x81, 0x22, 0x76, 0xfc,
	0xd0, 0xe2, 0x19, 0x84, 0xf4, 0x69, 0x54, 0x93, 0x06, 0x59, 0x50, 0x4c, 0xcb, 0x43, 0xee, 0x6e,
	0xe1, 0x95, 0xd5, 0xaa, 0xde, 0xa8, 0x9d, 0xf0, 0xfb, 0x78, 0x04, 0xa1, 0x92, 0x1d, 0xe9, 0x74,
	0x92, 0x05, 0x45, 0x22, 0x3c, 0x2c, 0x96, 0x00, 0xbf, 0x47, 0x71, 0x0e, 0xc1, 0x07, 0xed, 0x7e,
	0xa2, 0xed, 0x68, 0x5d, 0xdb, 0xa6, 0xdb, 0x90, 0x8b, 0x4d, 0x84, 0x87, 0xab, 0xf1, 0x92, 0xe5,
	0x37, 0xb0, 0x6f, 0xa3, 0x04, 0xe9, 0x41, 0xf6, 0x9a, 0xfe, 0xed, 0x7d, 0x02, 0xa1, 0x36, 0x8d,
	0xf1, 0xee, 0x59, 0x19, 0xf3, 0xaa, 0xb6, 0x28, 0xbc, 0x7a, 0x7e, 0x0a, 0x91, 0x17, 0x70, 0x0a,
	0x71, 0x6d, 0xe4, 0x30, 0x50, 0x3b, 0x1f, 0x21, 0x40, 0xf4, 0x28, 0xd5, 0xba, 0xe9, 0xe6, 0xac,
	0x2c, 0x61, 0x76, 0xa7, 0x88, 0x0c, 0xa9, 0x9a, 0xd4, 0x76, 0xf5, 0x46, 0x98, 0x41, 0x78, 0x4f,
	0x5d, 0x27, 0x31, 0xf2, 0x4f, 0x5d, 0x1c, 0xf0, 0xbf, 0x3d, 0xf2, 0xd1, 0x6d, 0xf2, 0x1c, 0xf3,
	0x6b, 0xf7, 0xad, 0xaf, 0x91, 0x5b, 0x2e, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x91, 0x9e, 0x97,
	0xc7, 0x6b, 0x01, 0x00, 0x00,
}
