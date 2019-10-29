// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bobsknobshop/truth/v1/truth.proto

package classy

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

type GetServiceKpiRequest struct {
	StartTimestamp       int64    `protobuf:"varint,1,opt,name=start_timestamp,json=startTimestamp,proto3" json:"start_timestamp,omitempty"`
	EndTimestamp         int64    `protobuf:"varint,2,opt,name=end_timestamp,json=endTimestamp,proto3" json:"end_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServiceKpiRequest) Reset()         { *m = GetServiceKpiRequest{} }
func (m *GetServiceKpiRequest) String() string { return proto.CompactTextString(m) }
func (*GetServiceKpiRequest) ProtoMessage()    {}
func (*GetServiceKpiRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ed36ddfdb54d891, []int{0}
}

func (m *GetServiceKpiRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServiceKpiRequest.Unmarshal(m, b)
}
func (m *GetServiceKpiRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServiceKpiRequest.Marshal(b, m, deterministic)
}
func (m *GetServiceKpiRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServiceKpiRequest.Merge(m, src)
}
func (m *GetServiceKpiRequest) XXX_Size() int {
	return xxx_messageInfo_GetServiceKpiRequest.Size(m)
}
func (m *GetServiceKpiRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServiceKpiRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetServiceKpiRequest proto.InternalMessageInfo

func (m *GetServiceKpiRequest) GetStartTimestamp() int64 {
	if m != nil {
		return m.StartTimestamp
	}
	return 0
}

func (m *GetServiceKpiRequest) GetEndTimestamp() int64 {
	if m != nil {
		return m.EndTimestamp
	}
	return 0
}

type GetServiceKpiResponse struct {
	Versions             []*GetServiceKpiResponse_Version `protobuf:"bytes,1,rep,name=versions,proto3" json:"versions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *GetServiceKpiResponse) Reset()         { *m = GetServiceKpiResponse{} }
func (m *GetServiceKpiResponse) String() string { return proto.CompactTextString(m) }
func (*GetServiceKpiResponse) ProtoMessage()    {}
func (*GetServiceKpiResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ed36ddfdb54d891, []int{1}
}

func (m *GetServiceKpiResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServiceKpiResponse.Unmarshal(m, b)
}
func (m *GetServiceKpiResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServiceKpiResponse.Marshal(b, m, deterministic)
}
func (m *GetServiceKpiResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServiceKpiResponse.Merge(m, src)
}
func (m *GetServiceKpiResponse) XXX_Size() int {
	return xxx_messageInfo_GetServiceKpiResponse.Size(m)
}
func (m *GetServiceKpiResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServiceKpiResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetServiceKpiResponse proto.InternalMessageInfo

func (m *GetServiceKpiResponse) GetVersions() []*GetServiceKpiResponse_Version {
	if m != nil {
		return m.Versions
	}
	return nil
}

type GetServiceKpiResponse_Version struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	StartTimestamp       int64    `protobuf:"varint,2,opt,name=start_timestamp,json=startTimestamp,proto3" json:"start_timestamp,omitempty"`
	EndTimestamp         int64    `protobuf:"varint,3,opt,name=end_timestamp,json=endTimestamp,proto3" json:"end_timestamp,omitempty"`
	Unit                 string   `protobuf:"bytes,4,opt,name=unit,proto3" json:"unit,omitempty"`
	Value                float32  `protobuf:"fixed32,5,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServiceKpiResponse_Version) Reset()         { *m = GetServiceKpiResponse_Version{} }
func (m *GetServiceKpiResponse_Version) String() string { return proto.CompactTextString(m) }
func (*GetServiceKpiResponse_Version) ProtoMessage()    {}
func (*GetServiceKpiResponse_Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ed36ddfdb54d891, []int{1, 0}
}

func (m *GetServiceKpiResponse_Version) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServiceKpiResponse_Version.Unmarshal(m, b)
}
func (m *GetServiceKpiResponse_Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServiceKpiResponse_Version.Marshal(b, m, deterministic)
}
func (m *GetServiceKpiResponse_Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServiceKpiResponse_Version.Merge(m, src)
}
func (m *GetServiceKpiResponse_Version) XXX_Size() int {
	return xxx_messageInfo_GetServiceKpiResponse_Version.Size(m)
}
func (m *GetServiceKpiResponse_Version) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServiceKpiResponse_Version.DiscardUnknown(m)
}

var xxx_messageInfo_GetServiceKpiResponse_Version proto.InternalMessageInfo

func (m *GetServiceKpiResponse_Version) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetServiceKpiResponse_Version) GetStartTimestamp() int64 {
	if m != nil {
		return m.StartTimestamp
	}
	return 0
}

func (m *GetServiceKpiResponse_Version) GetEndTimestamp() int64 {
	if m != nil {
		return m.EndTimestamp
	}
	return 0
}

func (m *GetServiceKpiResponse_Version) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *GetServiceKpiResponse_Version) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*GetServiceKpiRequest)(nil), "bobsknobshop.classy.v1.GetServiceKpiRequest")
	proto.RegisterType((*GetServiceKpiResponse)(nil), "bobsknobshop.classy.v1.GetServiceKpiResponse")
	proto.RegisterType((*GetServiceKpiResponse_Version)(nil), "bobsknobshop.classy.v1.GetServiceKpiResponse.Version")
}

func init() { proto.RegisterFile("bobsknobshop/truth/v1/truth.proto", fileDescriptor_1ed36ddfdb54d891) }

var fileDescriptor_1ed36ddfdb54d891 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x4d, 0xfa, 0x47, 0x5d, 0xad, 0xc2, 0x52, 0x25, 0x16, 0x0f, 0xb5, 0x1e, 0xec, 0x41,
	0xb3, 0xb4, 0xe2, 0x49, 0xbc, 0xe8, 0x41, 0x41, 0x0f, 0x1a, 0xa5, 0x07, 0x11, 0x64, 0xd3, 0x0e,
	0xe9, 0x62, 0xb2, 0x13, 0xb3, 0x9b, 0x60, 0xaf, 0x7e, 0x07, 0xbf, 0xaf, 0x74, 0x57, 0x4b, 0xab,
	0x15, 0xea, 0xed, 0x65, 0x78, 0x6f, 0x66, 0x7e, 0xd9, 0x21, 0x7b, 0x21, 0x86, 0xea, 0x45, 0x62,
	0xa8, 0x86, 0x98, 0x32, 0x9d, 0xe5, 0x7a, 0xc8, 0x8a, 0x8e, 0x15, 0x7e, 0x9a, 0xa1, 0x46, 0xba,
	0x3d, 0x6d, 0xf1, 0xfb, 0x31, 0x57, 0x6a, 0xe4, 0x17, 0x9d, 0xc6, 0x6e, 0x84, 0x18, 0xc5, 0xc0,
	0x78, 0x2a, 0x18, 0x97, 0x12, 0x35, 0xd7, 0x02, 0xa5, 0xb2, 0xa9, 0xd6, 0x80, 0xd4, 0x2f, 0x41,
	0xdf, 0x43, 0x56, 0x88, 0x3e, 0x5c, 0xa7, 0x22, 0x80, 0xd7, 0x1c, 0x94, 0xa6, 0x07, 0x64, 0x53,
	0x69, 0x9e, 0xe9, 0x67, 0x2d, 0x12, 0x50, 0x9a, 0x27, 0xa9, 0xe7, 0x34, 0x9d, 0x76, 0x29, 0xd8,
	0x30, 0xe5, 0x87, 0xef, 0x2a, 0xdd, 0x27, 0x35, 0x90, 0x83, 0x29, 0x9b, 0x6b, 0x6c, 0xeb, 0x20,
	0x07, 0x13, 0x53, 0xeb, 0xdd, 0x25, 0x5b, 0x3f, 0xc6, 0xa8, 0x14, 0xa5, 0x02, 0x7a, 0x47, 0x56,
	0x0a, 0xc8, 0xd4, 0x78, 0x23, 0xcf, 0x69, 0x96, 0xda, 0x6b, 0xdd, 0x13, 0x7f, 0x3e, 0x88, 0x3f,
	0xb7, 0x81, 0xdf, 0xb3, 0xe9, 0x60, 0xd2, 0xa6, 0xf1, 0xe1, 0x90, 0xe5, 0xaf, 0x2a, 0xa5, 0xa4,
	0x2c, 0x79, 0x02, 0x66, 0xf7, 0xd5, 0xc0, 0xe8, 0x79, 0x68, 0xee, 0x62, 0x68, 0xa5, 0xdf, 0x68,
	0xe3, 0x09, 0xb9, 0x14, 0xda, 0x2b, 0xdb, 0x09, 0x63, 0x4d, 0xeb, 0xa4, 0x52, 0xf0, 0x38, 0x07,
	0xaf, 0xd2, 0x74, 0xda, 0x6e, 0x60, 0x3f, 0xba, 0x6f, 0xa4, 0x7a, 0x61, 0x60, 0xa8, 0x24, 0xb5,
	0x19, 0x18, 0x7a, 0xb8, 0x20, 0xb3, 0x79, 0x9b, 0xc6, 0xd1, 0xbf, 0xfe, 0x50, 0x6b, 0xe9, 0xfc,
	0x89, 0xec, 0x44, 0xf1, 0x1f, 0xa1, 0x5b, 0xe7, 0xf1, 0x2c, 0x12, 0x7a, 0x98, 0x87, 0x7e, 0x1f,
	0x13, 0x76, 0xc5, 0x47, 0xd8, 0xe3, 0xf2, 0x06, 0x51, 0xb2, 0x08, 0xa4, 0x39, 0x12, 0x36, 0x73,
	0x7c, 0x36, 0xc9, 0x8a, 0xce, 0xa9, 0x55, 0x61, 0xd5, 0x98, 0x8e, 0x3f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x12, 0x8b, 0x5a, 0xf7, 0xa4, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClassyClient is the client API for Classy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClassyClient interface {
	// Returns a classification for a comment
	GetServiceKpi(ctx context.Context, in *GetServiceKpiRequest, opts ...grpc.CallOption) (*GetServiceKpiResponse, error)
}

type classyClient struct {
	cc *grpc.ClientConn
}

func NewClassyClient(cc *grpc.ClientConn) ClassyClient {
	return &classyClient{cc}
}

func (c *classyClient) GetServiceKpi(ctx context.Context, in *GetServiceKpiRequest, opts ...grpc.CallOption) (*GetServiceKpiResponse, error) {
	out := new(GetServiceKpiResponse)
	err := c.cc.Invoke(ctx, "/bobsknobshop.classy.v1.Classy/GetServiceKpi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClassyServer is the server API for Classy service.
type ClassyServer interface {
	// Returns a classification for a comment
	GetServiceKpi(context.Context, *GetServiceKpiRequest) (*GetServiceKpiResponse, error)
}

func RegisterClassyServer(s *grpc.Server, srv ClassyServer) {
	s.RegisterService(&_Classy_serviceDesc, srv)
}

func _Classy_GetServiceKpi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceKpiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassyServer).GetServiceKpi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bobsknobshop.classy.v1.Classy/GetServiceKpi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassyServer).GetServiceKpi(ctx, req.(*GetServiceKpiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Classy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bobsknobshop.classy.v1.Classy",
	HandlerType: (*ClassyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServiceKpi",
			Handler:    _Classy_GetServiceKpi_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bobsknobshop/truth/v1/truth.proto",
}