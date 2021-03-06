// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/api/dummy/v1alpha1/api.proto

// Deprecated: Do not use.
// v1alpha1 is no longer maintained, and will be removed soon.

package v1alpha1

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
	Input32              int32    `protobuf:"varint,1,opt,name=input32,proto3" json:"input32,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeDoubleRequest) Reset()         { *m = ComputeDoubleRequest{} }
func (m *ComputeDoubleRequest) String() string { return proto.CompactTextString(m) }
func (*ComputeDoubleRequest) ProtoMessage()    {}
func (*ComputeDoubleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_07d739019a1ce9cc, []int{0}
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

// Deprecated: Do not use.
func (m *ComputeDoubleRequest) GetInput32() int32 {
	if m != nil {
		return m.Input32
	}
	return 0
}

type ComputeDoubleResponse struct {
	Response32           int32    `protobuf:"varint,1,opt,name=response32,proto3" json:"response32,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeDoubleResponse) Reset()         { *m = ComputeDoubleResponse{} }
func (m *ComputeDoubleResponse) String() string { return proto.CompactTextString(m) }
func (*ComputeDoubleResponse) ProtoMessage()    {}
func (*ComputeDoubleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_07d739019a1ce9cc, []int{1}
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

func (m *ComputeDoubleResponse) GetResponse32() int32 {
	if m != nil {
		return m.Response32
	}
	return 0
}

func init() {
	proto.RegisterType((*ComputeDoubleRequest)(nil), "v1alpha1.ComputeDoubleRequest")
	proto.RegisterType((*ComputeDoubleResponse)(nil), "v1alpha1.ComputeDoubleResponse")
}

func init() {
	proto.RegisterFile("github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/api/dummy/v1alpha1/api.proto", fileDescriptor_07d739019a1ce9cc)
}

var fileDescriptor_07d739019a1ce9cc = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0xa5, 0x2a, 0x03, 0x5e, 0x16, 0x85, 0x22, 0x52, 0x25, 0x27, 0x2f, 0xcd, 0xd2,
	0x56, 0xf0, 0xae, 0x7d, 0x82, 0x80, 0x27, 0x41, 0x48, 0xe2, 0x90, 0x2e, 0x36, 0x3b, 0xe3, 0xce,
	0x8c, 0xd8, 0x37, 0xf0, 0xb1, 0xa5, 0xd1, 0x80, 0x8a, 0xbd, 0xfd, 0x7c, 0xc3, 0xff, 0xf3, 0x0d,
	0x3c, 0xb6, 0x41, 0xd7, 0x56, 0x17, 0x0d, 0x75, 0xfe, 0xc5, 0x6a, 0x4c, 0x11, 0x15, 0x65, 0xd6,
	0x48, 0xf0, 0x8d, 0x84, 0x19, 0x27, 0x7a, 0xdf, 0xfa, 0x10, 0x15, 0xdb, 0x54, 0x69, 0xa0, 0xa8,
	0x28, 0x2a, 0xbe, 0xe2, 0xd0, 0x26, 0x32, 0xee, 0x93, 0x7f, 0xb6, 0xae, 0xdb, 0xfa, 0xb7, 0x79,
	0xb5, 0xe1, 0x75, 0x35, 0xdf, 0xa1, 0x82, 0x13, 0x29, 0xb9, 0xe3, 0x81, 0xe5, 0x37, 0x70, 0x7a,
	0x4f, 0x1d, 0x9b, 0xe2, 0x8a, 0xac, 0xde, 0x60, 0x89, 0xaf, 0x86, 0xa2, 0xee, 0x02, 0x8e, 0x42,
	0x64, 0xd3, 0xe5, 0x62, 0x92, 0x5d, 0x65, 0xd7, 0xe3, 0xbb, 0xd1, 0x24, 0x2b, 0x07, 0x94, 0xdf,
	0xc2, 0xd9, 0x9f, 0x96, 0x30, 0x45, 0x41, 0x37, 0x05, 0x48, 0xdf, 0x79, 0x68, 0x96, 0x3f, 0xc8,
	0xe2, 0x09, 0xc6, 0xab, 0x9d, 0x94, 0x7b, 0x80, 0x93, 0x5f, 0x0b, 0x6e, 0x5a, 0x0c, 0x4e, 0xc5,
	0x7f, 0x42, 0xe7, 0x97, 0x7b, 0xef, 0x5f, 0xe3, 0xf9, 0xc1, 0xc7, 0x28, 0xab, 0x0f, 0xfb, 0xff,
	0x96, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x93, 0xe6, 0xf6, 0x57, 0x3e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DummyClient is the client API for Dummy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DummyClient interface {
	// ComputeDouble computes the double of the input. Real smart stuff!
	//
	// Deprecated: Do not use.
	ComputeDouble(ctx context.Context, in *ComputeDoubleRequest, opts ...grpc.CallOption) (*ComputeDoubleResponse, error)
}

type dummyClient struct {
	cc grpc.ClientConnInterface
}

func NewDummyClient(cc grpc.ClientConnInterface) DummyClient {
	return &dummyClient{cc}
}

// Deprecated: Do not use.
func (c *dummyClient) ComputeDouble(ctx context.Context, in *ComputeDoubleRequest, opts ...grpc.CallOption) (*ComputeDoubleResponse, error) {
	out := new(ComputeDoubleResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.Dummy/ComputeDouble", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DummyServer is the server API for Dummy service.
type DummyServer interface {
	// ComputeDouble computes the double of the input. Real smart stuff!
	//
	// Deprecated: Do not use.
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
		FullMethod: "/v1alpha1.Dummy/ComputeDouble",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DummyServer).ComputeDouble(ctx, req.(*ComputeDoubleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dummy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1alpha1.Dummy",
	HandlerType: (*DummyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ComputeDouble",
			Handler:    _Dummy_ComputeDouble_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/api/dummy/v1alpha1/api.proto",
}
