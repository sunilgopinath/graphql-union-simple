// Code generated by protoc-gen-go. DO NOT EDIT.
// source: spark.proto

/*
Package sparkAnalytics is a generated protocol buffer package.

It is generated from these files:
	spark.proto

It has these top-level messages:
	AssetRequest
	AssetReply
*/
package spark

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

// The request message containing the user's name.
type AssetRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *AssetRequest) Reset()                    { *m = AssetRequest{} }
func (m *AssetRequest) String() string            { return proto.CompactTextString(m) }
func (*AssetRequest) ProtoMessage()               {}
func (*AssetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AssetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type AssetReply struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Current  int32  `protobuf:"varint,2,opt,name=current" json:"current,omitempty"`
	Trending bool   `protobuf:"varint,3,opt,name=trending" json:"trending,omitempty"`
}

func (m *AssetReply) Reset()                    { *m = AssetReply{} }
func (m *AssetReply) String() string            { return proto.CompactTextString(m) }
func (*AssetReply) ProtoMessage()               {}
func (*AssetReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AssetReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AssetReply) GetCurrent() int32 {
	if m != nil {
		return m.Current
	}
	return 0
}

func (m *AssetReply) GetTrending() bool {
	if m != nil {
		return m.Trending
	}
	return false
}

func init() {
	proto.RegisterType((*AssetRequest)(nil), "sparkAnalytics.AssetRequest")
	proto.RegisterType((*AssetReply)(nil), "sparkAnalytics.AssetReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SparkAnalytics service

type SparkAnalyticsClient interface {
	// Sends a greeting
	AssetStatistics(ctx context.Context, in *AssetRequest, opts ...grpc.CallOption) (*AssetReply, error)
}

type sparkAnalyticsClient struct {
	cc *grpc.ClientConn
}

func NewSparkAnalyticsClient(cc *grpc.ClientConn) SparkAnalyticsClient {
	return &sparkAnalyticsClient{cc}
}

func (c *sparkAnalyticsClient) AssetStatistics(ctx context.Context, in *AssetRequest, opts ...grpc.CallOption) (*AssetReply, error) {
	out := new(AssetReply)
	err := grpc.Invoke(ctx, "/sparkAnalytics.SparkAnalytics/AssetStatistics", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SparkAnalytics service

type SparkAnalyticsServer interface {
	// Sends a greeting
	AssetStatistics(context.Context, *AssetRequest) (*AssetReply, error)
}

func RegisterSparkAnalyticsServer(s *grpc.Server, srv SparkAnalyticsServer) {
	s.RegisterService(&_SparkAnalytics_serviceDesc, srv)
}

func _SparkAnalytics_AssetStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SparkAnalyticsServer).AssetStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sparkAnalytics.SparkAnalytics/AssetStatistics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SparkAnalyticsServer).AssetStatistics(ctx, req.(*AssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SparkAnalytics_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sparkAnalytics.SparkAnalytics",
	HandlerType: (*SparkAnalyticsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AssetStatistics",
			Handler:    _SparkAnalytics_AssetStatistics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spark.proto",
}

func init() { proto.RegisterFile("spark.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0xcd, 0xf8, 0x57, 0xaf, 0x32, 0x85, 0xac, 0x86, 0xd1, 0xc5, 0x90, 0xd5, 0xac, 0x82,
	0xe8, 0x13, 0xb4, 0x2b, 0x37, 0x42, 0x49, 0x17, 0x2e, 0x25, 0x4e, 0x2e, 0x35, 0x98, 0x26, 0x31,
	0xb9, 0x45, 0xe7, 0xed, 0x65, 0x42, 0x15, 0xbb, 0xe8, 0xee, 0x7c, 0x39, 0x81, 0xfb, 0x71, 0xe0,
	0x3a, 0x47, 0x9d, 0x3e, 0x64, 0x4c, 0x81, 0x02, 0xaf, 0x0b, 0x2c, 0xbc, 0x76, 0x23, 0xd9, 0x21,
	0x0b, 0x01, 0x37, 0x8b, 0x9c, 0x91, 0x14, 0x7e, 0xee, 0x30, 0x13, 0xe7, 0x70, 0xe6, 0xf5, 0x16,
	0x1b, 0xd6, 0xb1, 0xfe, 0x4a, 0x95, 0x2c, 0x14, 0xc0, 0xfe, 0x4f, 0x74, 0x23, 0xaf, 0xa1, 0xb2,
	0x66, 0xdf, 0x57, 0xd6, 0xf0, 0x06, 0x2e, 0x87, 0x5d, 0x4a, 0xe8, 0xa9, 0xa9, 0x3a, 0xd6, 0x9f,
	0xab, 0x5f, 0xe4, 0x2d, 0xcc, 0x28, 0xa1, 0x37, 0xd6, 0x6f, 0x9a, 0xd3, 0x8e, 0xf5, 0x33, 0xf5,
	0xc7, 0x0f, 0xaf, 0x50, 0xaf, 0x0f, 0x4c, 0xf8, 0x33, 0xcc, 0xcb, 0x95, 0x35, 0x69, 0xb2, 0xb9,
	0x3c, 0xdd, 0xc9, 0x43, 0x5b, 0xf9, 0x5f, 0xb5, 0x6d, 0x8f, 0xb4, 0xd1, 0x8d, 0xe2, 0x64, 0x79,
	0x0f, 0xb7, 0x36, 0xc8, 0x4d, 0x8a, 0x83, 0xc4, 0x6f, 0xbd, 0x8d, 0x0e, 0xb3, 0x7c, 0x47, 0xe7,
	0xc2, 0x57, 0x48, 0xce, 0x2c, 0xe7, 0x4f, 0x53, 0x7e, 0x99, 0xf2, 0x6a, 0x1a, 0x66, 0xc5, 0xde,
	0x2e, 0xca, 0x42, 0x8f, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x7c, 0x8e, 0xa4, 0x30, 0x01,
	0x00, 0x00,
}
