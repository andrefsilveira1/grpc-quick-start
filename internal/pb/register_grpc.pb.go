// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: proto/register.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RegisterService_CreateRegister_FullMethodName              = "/pb.RegisterService/CreateRegister"
	RegisterService_CreateRegisterStream_FullMethodName        = "/pb.RegisterService/CreateRegisterStream"
	RegisterService_CreateRegisterBidirectional_FullMethodName = "/pb.RegisterService/CreateRegisterBidirectional"
)

// RegisterServiceClient is the client API for RegisterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegisterServiceClient interface {
	CreateRegister(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Register, error)
	CreateRegisterStream(ctx context.Context, opts ...grpc.CallOption) (RegisterService_CreateRegisterStreamClient, error)
	CreateRegisterBidirectional(ctx context.Context, opts ...grpc.CallOption) (RegisterService_CreateRegisterBidirectionalClient, error)
}

type registerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRegisterServiceClient(cc grpc.ClientConnInterface) RegisterServiceClient {
	return &registerServiceClient{cc}
}

func (c *registerServiceClient) CreateRegister(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Register, error) {
	out := new(Register)
	err := c.cc.Invoke(ctx, RegisterService_CreateRegister_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registerServiceClient) CreateRegisterStream(ctx context.Context, opts ...grpc.CallOption) (RegisterService_CreateRegisterStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &RegisterService_ServiceDesc.Streams[0], RegisterService_CreateRegisterStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &registerServiceCreateRegisterStreamClient{stream}
	return x, nil
}

type RegisterService_CreateRegisterStreamClient interface {
	Send(*CreateRequest) error
	CloseAndRecv() (*Registers, error)
	grpc.ClientStream
}

type registerServiceCreateRegisterStreamClient struct {
	grpc.ClientStream
}

func (x *registerServiceCreateRegisterStreamClient) Send(m *CreateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *registerServiceCreateRegisterStreamClient) CloseAndRecv() (*Registers, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Registers)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *registerServiceClient) CreateRegisterBidirectional(ctx context.Context, opts ...grpc.CallOption) (RegisterService_CreateRegisterBidirectionalClient, error) {
	stream, err := c.cc.NewStream(ctx, &RegisterService_ServiceDesc.Streams[1], RegisterService_CreateRegisterBidirectional_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &registerServiceCreateRegisterBidirectionalClient{stream}
	return x, nil
}

type RegisterService_CreateRegisterBidirectionalClient interface {
	Send(*CreateRequest) error
	Recv() (*Register, error)
	grpc.ClientStream
}

type registerServiceCreateRegisterBidirectionalClient struct {
	grpc.ClientStream
}

func (x *registerServiceCreateRegisterBidirectionalClient) Send(m *CreateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *registerServiceCreateRegisterBidirectionalClient) Recv() (*Register, error) {
	m := new(Register)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RegisterServiceServer is the server API for RegisterService service.
// All implementations must embed UnimplementedRegisterServiceServer
// for forward compatibility
type RegisterServiceServer interface {
	CreateRegister(context.Context, *CreateRequest) (*Register, error)
	CreateRegisterStream(RegisterService_CreateRegisterStreamServer) error
	CreateRegisterBidirectional(RegisterService_CreateRegisterBidirectionalServer) error
	mustEmbedUnimplementedRegisterServiceServer()
}

// UnimplementedRegisterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRegisterServiceServer struct {
}

func (UnimplementedRegisterServiceServer) CreateRegister(context.Context, *CreateRequest) (*Register, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRegister not implemented")
}
func (UnimplementedRegisterServiceServer) CreateRegisterStream(RegisterService_CreateRegisterStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateRegisterStream not implemented")
}
func (UnimplementedRegisterServiceServer) CreateRegisterBidirectional(RegisterService_CreateRegisterBidirectionalServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateRegisterBidirectional not implemented")
}
func (UnimplementedRegisterServiceServer) mustEmbedUnimplementedRegisterServiceServer() {}

// UnsafeRegisterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegisterServiceServer will
// result in compilation errors.
type UnsafeRegisterServiceServer interface {
	mustEmbedUnimplementedRegisterServiceServer()
}

func RegisterRegisterServiceServer(s grpc.ServiceRegistrar, srv RegisterServiceServer) {
	s.RegisterService(&RegisterService_ServiceDesc, srv)
}

func _RegisterService_CreateRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServiceServer).CreateRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RegisterService_CreateRegister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServiceServer).CreateRegister(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegisterService_CreateRegisterStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RegisterServiceServer).CreateRegisterStream(&registerServiceCreateRegisterStreamServer{stream})
}

type RegisterService_CreateRegisterStreamServer interface {
	SendAndClose(*Registers) error
	Recv() (*CreateRequest, error)
	grpc.ServerStream
}

type registerServiceCreateRegisterStreamServer struct {
	grpc.ServerStream
}

func (x *registerServiceCreateRegisterStreamServer) SendAndClose(m *Registers) error {
	return x.ServerStream.SendMsg(m)
}

func (x *registerServiceCreateRegisterStreamServer) Recv() (*CreateRequest, error) {
	m := new(CreateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RegisterService_CreateRegisterBidirectional_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RegisterServiceServer).CreateRegisterBidirectional(&registerServiceCreateRegisterBidirectionalServer{stream})
}

type RegisterService_CreateRegisterBidirectionalServer interface {
	Send(*Register) error
	Recv() (*CreateRequest, error)
	grpc.ServerStream
}

type registerServiceCreateRegisterBidirectionalServer struct {
	grpc.ServerStream
}

func (x *registerServiceCreateRegisterBidirectionalServer) Send(m *Register) error {
	return x.ServerStream.SendMsg(m)
}

func (x *registerServiceCreateRegisterBidirectionalServer) Recv() (*CreateRequest, error) {
	m := new(CreateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RegisterService_ServiceDesc is the grpc.ServiceDesc for RegisterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RegisterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.RegisterService",
	HandlerType: (*RegisterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRegister",
			Handler:    _RegisterService_CreateRegister_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateRegisterStream",
			Handler:       _RegisterService_CreateRegisterStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "CreateRegisterBidirectional",
			Handler:       _RegisterService_CreateRegisterBidirectional_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/register.proto",
}
