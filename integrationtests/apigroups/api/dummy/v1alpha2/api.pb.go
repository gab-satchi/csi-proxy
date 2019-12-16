// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package v1alpha2

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

type ComputeDoubleRequest struct {
	// we changed in favor of an int64 field here
	Input64              int64    `protobuf:"varint,1,opt,name=input64,proto3" json:"input64,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeDoubleRequest) Reset()         { *m = ComputeDoubleRequest{} }
func (m *ComputeDoubleRequest) String() string { return proto.CompactTextString(m) }
func (*ComputeDoubleRequest) ProtoMessage()    {}
func (*ComputeDoubleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *ComputeDoubleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeDoubleRequest.Unmarshal(m, b)
}
func (m *ComputeDoubleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeDoubleRequest.Marshal(b, m, deterministic)
}
func (m *ComputeDoubleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeDoubleRequest.Merge(m, src)
}
func (m *ComputeDoubleRequest) XXX_Size() int {
	return xxx_messageInfo_ComputeDoubleRequest.Size(m)
}
func (m *ComputeDoubleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeDoubleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeDoubleRequest proto.InternalMessageInfo

func (m *ComputeDoubleRequest) GetInput64() int64 {
	if m != nil {
		return m.Input64
	}
	return 0
}

type ComputeDoubleResponse struct {
	Response             int64    `protobuf:"varint,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeDoubleResponse) Reset()         { *m = ComputeDoubleResponse{} }
func (m *ComputeDoubleResponse) String() string { return proto.CompactTextString(m) }
func (*ComputeDoubleResponse) ProtoMessage()    {}
func (*ComputeDoubleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *ComputeDoubleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeDoubleResponse.Unmarshal(m, b)
}
func (m *ComputeDoubleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeDoubleResponse.Marshal(b, m, deterministic)
}
func (m *ComputeDoubleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeDoubleResponse.Merge(m, src)
}
func (m *ComputeDoubleResponse) XXX_Size() int {
	return xxx_messageInfo_ComputeDoubleResponse.Size(m)
}
func (m *ComputeDoubleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeDoubleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeDoubleResponse proto.InternalMessageInfo

func (m *ComputeDoubleResponse) GetResponse() int64 {
	if m != nil {
		return m.Response
	}
	return 0
}

func init() {
	proto.RegisterType((*ComputeDoubleRequest)(nil), "v1alpha2.ComputeDoubleRequest")
	proto.RegisterType((*ComputeDoubleResponse)(nil), "v1alpha2.ComputeDoubleResponse")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x28, 0x33, 0x4c, 0xcc, 0x29, 0xc8, 0x48, 0x34, 0x52,
	0x32, 0xe0, 0x12, 0x71, 0xce, 0xcf, 0x2d, 0x28, 0x2d, 0x49, 0x75, 0xc9, 0x2f, 0x4d, 0xca, 0x49,
	0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe0, 0x62, 0xcf, 0xcc, 0x2b, 0x28, 0x2d,
	0x31, 0x33, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x82, 0x71, 0x95, 0x8c, 0xb9, 0x44, 0xd1,
	0x74, 0x14, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x49, 0x71, 0x71, 0x14, 0x41, 0xd9, 0x50, 0x3d,
	0x70, 0xbe, 0x51, 0x34, 0x17, 0xab, 0x4b, 0x69, 0x6e, 0x6e, 0xa5, 0x50, 0x10, 0x17, 0x2f, 0x8a,
	0x6e, 0x21, 0x39, 0x3d, 0x98, 0x5b, 0xf4, 0xb0, 0x39, 0x44, 0x4a, 0x1e, 0xa7, 0x3c, 0xc4, 0x68,
	0x25, 0x86, 0x24, 0x36, 0xb0, 0xa7, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2b, 0xdc, 0xda,
	0x04, 0xe1, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DummyClient is the client API for Dummy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DummyClient interface {
	// ComputeDouble computes the double of the input. Real smart stuff!
	ComputeDouble(ctx context.Context, in *ComputeDoubleRequest, opts ...grpc.CallOption) (*ComputeDoubleResponse, error)
}

type dummyClient struct {
	cc *grpc.ClientConn
}

func NewDummyClient(cc *grpc.ClientConn) DummyClient {
	return &dummyClient{cc}
}

func (c *dummyClient) ComputeDouble(ctx context.Context, in *ComputeDoubleRequest, opts ...grpc.CallOption) (*ComputeDoubleResponse, error) {
	out := new(ComputeDoubleResponse)
	err := c.cc.Invoke(ctx, "/v1alpha2.Dummy/ComputeDouble", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DummyServer is the server API for Dummy service.
type DummyServer interface {
	// ComputeDouble computes the double of the input. Real smart stuff!
	ComputeDouble(context.Context, *ComputeDoubleRequest) (*ComputeDoubleResponse, error)
}

// UnimplementedDummyServer can be embedded to have forward compatible implementations.
type UnimplementedDummyServer struct {
}

func (*UnimplementedDummyServer) ComputeDouble(ctx context.Context, req *ComputeDoubleRequest) (*ComputeDoubleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComputeDouble not implemented")
}

func RegisterDummyServer(s *grpc.Server, srv DummyServer) {
	s.RegisterService(&_Dummy_serviceDesc, srv)
}

func _Dummy_ComputeDouble_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComputeDoubleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DummyServer).ComputeDouble(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha2.Dummy/ComputeDouble",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DummyServer).ComputeDouble(ctx, req.(*ComputeDoubleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dummy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1alpha2.Dummy",
	HandlerType: (*DummyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ComputeDouble",
			Handler:    _Dummy_ComputeDouble_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}