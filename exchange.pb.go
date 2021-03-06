// Code generated by protoc-gen-go. DO NOT EDIT.
// source: exchange.proto

/*
Package exchange is a generated protocol buffer package.

It is generated from these files:
	exchange.proto

It has these top-level messages:
	RateRequest
	RateReply
	RateList
*/
package exchange

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

type RateRequest struct {
	Base   string `protobuf:"bytes,1,opt,name=base" json:"base,omitempty"`
	Target string `protobuf:"bytes,2,opt,name=target" json:"target,omitempty"`
}

func (m *RateRequest) Reset()                    { *m = RateRequest{} }
func (m *RateRequest) String() string            { return proto.CompactTextString(m) }
func (*RateRequest) ProtoMessage()               {}
func (*RateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RateRequest) GetBase() string {
	if m != nil {
		return m.Base
	}
	return ""
}

func (m *RateRequest) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

type RateReply struct {
	Base   string  `protobuf:"bytes,1,opt,name=base" json:"base,omitempty"`
	Target string  `protobuf:"bytes,2,opt,name=target" json:"target,omitempty"`
	Rate   float64 `protobuf:"fixed64,3,opt,name=rate" json:"rate,omitempty"`
}

func (m *RateReply) Reset()                    { *m = RateReply{} }
func (m *RateReply) String() string            { return proto.CompactTextString(m) }
func (*RateReply) ProtoMessage()               {}
func (*RateReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RateReply) GetBase() string {
	if m != nil {
		return m.Base
	}
	return ""
}

func (m *RateReply) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *RateReply) GetRate() float64 {
	if m != nil {
		return m.Rate
	}
	return 0
}

type RateList struct {
	Count    int32        `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
	Rates    []*RateReply `protobuf:"bytes,2,rep,name=rates" json:"rates,omitempty"`
	CostTime int32        `protobuf:"varint,3,opt,name=cost_time,json=costTime" json:"cost_time,omitempty"`
}

func (m *RateList) Reset()                    { *m = RateList{} }
func (m *RateList) String() string            { return proto.CompactTextString(m) }
func (*RateList) ProtoMessage()               {}
func (*RateList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RateList) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *RateList) GetRates() []*RateReply {
	if m != nil {
		return m.Rates
	}
	return nil
}

func (m *RateList) GetCostTime() int32 {
	if m != nil {
		return m.CostTime
	}
	return 0
}

func init() {
	proto.RegisterType((*RateRequest)(nil), "exchange.RateRequest")
	proto.RegisterType((*RateReply)(nil), "exchange.RateReply")
	proto.RegisterType((*RateList)(nil), "exchange.RateList")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ExchangeService service

type ExchangeServiceClient interface {
	GetRate(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateReply, error)
	ListRate(ctx context.Context, opts ...grpc.CallOption) (ExchangeService_ListRateClient, error)
}

type exchangeServiceClient struct {
	cc *grpc.ClientConn
}

func NewExchangeServiceClient(cc *grpc.ClientConn) ExchangeServiceClient {
	return &exchangeServiceClient{cc}
}

func (c *exchangeServiceClient) GetRate(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateReply, error) {
	out := new(RateReply)
	err := grpc.Invoke(ctx, "/exchange.ExchangeService/GetRate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeServiceClient) ListRate(ctx context.Context, opts ...grpc.CallOption) (ExchangeService_ListRateClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ExchangeService_serviceDesc.Streams[0], c.cc, "/exchange.ExchangeService/ListRate", opts...)
	if err != nil {
		return nil, err
	}
	x := &exchangeServiceListRateClient{stream}
	return x, nil
}

type ExchangeService_ListRateClient interface {
	Send(*RateRequest) error
	CloseAndRecv() (*RateList, error)
	grpc.ClientStream
}

type exchangeServiceListRateClient struct {
	grpc.ClientStream
}

func (x *exchangeServiceListRateClient) Send(m *RateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *exchangeServiceListRateClient) CloseAndRecv() (*RateList, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(RateList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ExchangeService service

type ExchangeServiceServer interface {
	GetRate(context.Context, *RateRequest) (*RateReply, error)
	ListRate(ExchangeService_ListRateServer) error
}

func RegisterExchangeServiceServer(s *grpc.Server, srv ExchangeServiceServer) {
	s.RegisterService(&_ExchangeService_serviceDesc, srv)
}

func _ExchangeService_GetRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).GetRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exchange.ExchangeService/GetRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).GetRate(ctx, req.(*RateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeService_ListRate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExchangeServiceServer).ListRate(&exchangeServiceListRateServer{stream})
}

type ExchangeService_ListRateServer interface {
	SendAndClose(*RateList) error
	Recv() (*RateRequest, error)
	grpc.ServerStream
}

type exchangeServiceListRateServer struct {
	grpc.ServerStream
}

func (x *exchangeServiceListRateServer) SendAndClose(m *RateList) error {
	return x.ServerStream.SendMsg(m)
}

func (x *exchangeServiceListRateServer) Recv() (*RateRequest, error) {
	m := new(RateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ExchangeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "exchange.ExchangeService",
	HandlerType: (*ExchangeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRate",
			Handler:    _ExchangeService_GetRate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListRate",
			Handler:       _ExchangeService_ListRate_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "exchange.proto",
}

func init() { proto.RegisterFile("exchange.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0x4d, 0x4b, 0x43, 0x31,
	0x10, 0x6c, 0x5a, 0x5f, 0x7d, 0x6f, 0x0b, 0x0a, 0xeb, 0x07, 0x0f, 0xbd, 0x3c, 0x72, 0x8a, 0x97,
	0x1e, 0xea, 0x41, 0xfa, 0x03, 0xc4, 0x83, 0x9e, 0xa2, 0x77, 0x49, 0xc3, 0xd2, 0x06, 0x6c, 0x53,
	0x93, 0xad, 0xe8, 0x0f, 0xf0, 0x7f, 0x4b, 0x92, 0x5a, 0x50, 0x44, 0xe8, 0x6d, 0x67, 0x98, 0x99,
	0x9d, 0x64, 0xe1, 0x88, 0xde, 0xed, 0xc2, 0xac, 0xe6, 0x34, 0x5e, 0x07, 0xcf, 0x1e, 0xeb, 0x6f,
	0x2c, 0xa7, 0x30, 0xd2, 0x86, 0x49, 0xd3, 0xeb, 0x86, 0x22, 0x23, 0xc2, 0xc1, 0xcc, 0x44, 0x6a,
	0x45, 0x27, 0x54, 0xa3, 0xf3, 0x8c, 0xe7, 0x30, 0x64, 0x13, 0xe6, 0xc4, 0x6d, 0x3f, 0xb3, 0x5b,
	0x24, 0xef, 0xa1, 0x29, 0xd6, 0xf5, 0xcb, 0xc7, 0x3e, 0xc6, 0xa4, 0x0d, 0x86, 0xa9, 0x1d, 0x74,
	0x42, 0x09, 0x9d, 0x67, 0xb9, 0x80, 0x3a, 0x85, 0x3d, 0xb8, 0xc8, 0x78, 0x0a, 0x95, 0xf5, 0x9b,
	0x15, 0xe7, 0xb0, 0x4a, 0x17, 0x80, 0x57, 0x50, 0x25, 0x65, 0x6c, 0xfb, 0xdd, 0x40, 0x8d, 0x26,
	0x27, 0xe3, 0xdd, 0x9b, 0x76, 0x2d, 0x74, 0x51, 0xe0, 0x25, 0x34, 0xd6, 0x47, 0x7e, 0x66, 0xb7,
	0x2c, 0x5b, 0x2a, 0x5d, 0x27, 0xe2, 0xc9, 0x2d, 0x69, 0xf2, 0x29, 0xe0, 0xf8, 0x76, 0x6b, 0x7d,
	0xa4, 0xf0, 0xe6, 0x2c, 0xe1, 0x0d, 0x1c, 0xde, 0x11, 0xa7, 0x1c, 0x3c, 0xfb, 0x9d, 0x9b, 0x3f,
	0xe6, 0xe2, 0xaf, 0x75, 0xb2, 0x87, 0x53, 0xa8, 0x53, 0xe5, 0xff, 0x9c, 0xf8, 0x93, 0x4e, 0x72,
	0xd9, 0x53, 0x62, 0x36, 0xcc, 0xa7, 0xb8, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x61, 0x37, 0x6a,
	0xd3, 0x9c, 0x01, 0x00, 0x00,
}
