// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cloudrungrpc/greeter.proto

package greeter

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetGreetingRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGreetingRequest) Reset()         { *m = GetGreetingRequest{} }
func (m *GetGreetingRequest) String() string { return proto.CompactTextString(m) }
func (*GetGreetingRequest) ProtoMessage()    {}
func (*GetGreetingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd1af738d44a1a9c, []int{0}
}

func (m *GetGreetingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGreetingRequest.Unmarshal(m, b)
}
func (m *GetGreetingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGreetingRequest.Marshal(b, m, deterministic)
}
func (m *GetGreetingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGreetingRequest.Merge(m, src)
}
func (m *GetGreetingRequest) XXX_Size() int {
	return xxx_messageInfo_GetGreetingRequest.Size(m)
}
func (m *GetGreetingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGreetingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGreetingRequest proto.InternalMessageInfo

func (m *GetGreetingRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Greeting struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Greeting) Reset()         { *m = Greeting{} }
func (m *Greeting) String() string { return proto.CompactTextString(m) }
func (*Greeting) ProtoMessage()    {}
func (*Greeting) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd1af738d44a1a9c, []int{1}
}

func (m *Greeting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Greeting.Unmarshal(m, b)
}
func (m *Greeting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Greeting.Marshal(b, m, deterministic)
}
func (m *Greeting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Greeting.Merge(m, src)
}
func (m *Greeting) XXX_Size() int {
	return xxx_messageInfo_Greeting.Size(m)
}
func (m *Greeting) XXX_DiscardUnknown() {
	xxx_messageInfo_Greeting.DiscardUnknown(m)
}

var xxx_messageInfo_Greeting proto.InternalMessageInfo

func (m *Greeting) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*GetGreetingRequest)(nil), "cloudrungrpc.greeter.GetGreetingRequest")
	proto.RegisterType((*Greeting)(nil), "cloudrungrpc.greeter.Greeting")
}

func init() { proto.RegisterFile("cloudrungrpc/greeter.proto", fileDescriptor_fd1af738d44a1a9c) }

var fileDescriptor_fd1af738d44a1a9c = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x3d, 0x4b, 0xc4, 0x40,
	0x10, 0x86, 0x0d, 0x88, 0x1f, 0x63, 0xb7, 0x58, 0x48, 0x90, 0x43, 0x52, 0xa5, 0xda, 0x05, 0x2d,
	0xb5, 0xb2, 0x89, 0xc5, 0x15, 0x72, 0x85, 0xa0, 0xdd, 0x5e, 0x6e, 0x18, 0x17, 0x72, 0x33, 0x71,
	0x33, 0x0b, 0xfa, 0xef, 0x25, 0x5f, 0x10, 0x30, 0x56, 0xfb, 0xc2, 0xf3, 0xec, 0xbe, 0x3b, 0x03,
	0x79, 0xdd, 0x48, 0x3a, 0xc4, 0xc4, 0x14, 0xdb, 0xda, 0x51, 0x44, 0x54, 0x8c, 0xb6, 0x8d, 0xa2,
	0x62, 0xae, 0x97, 0xcc, 0x4e, 0x2c, 0xbf, 0x25, 0x11, 0x6a, 0xd0, 0xf9, 0x36, 0x38, 0xcf, 0x2c,
	0xea, 0x35, 0x08, 0x77, 0xe3, 0x9d, 0xa2, 0x04, 0x53, 0xa1, 0x56, 0xbd, 0x1b, 0x98, 0x76, 0xf8,
	0x95, 0xb0, 0x53, 0x63, 0xe0, 0x94, 0xfd, 0x11, 0x6f, 0xb2, 0xbb, 0xac, 0xbc, 0xdc, 0x0d, 0xb9,
	0xd8, 0xc0, 0xc5, 0xac, 0xf5, 0x5c, 0xf1, 0x5b, 0x67, 0xde, 0xe7, 0xfb, 0x03, 0x9c, 0x57, 0x63,
	0xa5, 0x79, 0x87, 0xab, 0xc5, 0xa3, 0xa6, 0xb4, 0x6b, 0x1f, 0xb3, 0x7f, 0x7b, 0xf3, 0xcd, 0x3f,
	0xe6, 0xa4, 0x15, 0x27, 0xcf, 0x5b, 0x00, 0x6a, 0x66, 0xf0, 0x9a, 0x7d, 0x3c, 0x51, 0xd0, 0xcf,
	0xb4, 0xb7, 0xb5, 0x1c, 0xdd, 0x8b, 0xff, 0x91, 0x37, 0xcf, 0x5b, 0x11, 0x76, 0x84, 0x3c, 0x0c,
	0xe8, 0xd6, 0xf6, 0xf5, 0x38, 0x9d, 0xfb, 0xb3, 0xc1, 0x79, 0xf8, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0x9f, 0x68, 0xdb, 0x2c, 0x56, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	GetGreeting(ctx context.Context, in *GetGreetingRequest, opts ...grpc.CallOption) (*Greeting, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) GetGreeting(ctx context.Context, in *GetGreetingRequest, opts ...grpc.CallOption) (*Greeting, error) {
	out := new(Greeting)
	err := c.cc.Invoke(ctx, "/cloudrungrpc.greeter.Greeter/GetGreeting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	GetGreeting(context.Context, *GetGreetingRequest) (*Greeting, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_GetGreeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGreetingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).GetGreeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloudrungrpc.greeter.Greeter/GetGreeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).GetGreeting(ctx, req.(*GetGreetingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cloudrungrpc.greeter.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGreeting",
			Handler:    _Greeter_GetGreeting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloudrungrpc/greeter.proto",
}
