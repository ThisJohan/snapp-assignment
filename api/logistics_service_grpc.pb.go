// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.6.1
// source: logistics_service.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	LogisticsService_Todo_FullMethodName = "/api.LogisticsService/Todo"
)

// LogisticsServiceClient is the client API for LogisticsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogisticsServiceClient interface {
	Todo(ctx context.Context, in *TodoRequest, opts ...grpc.CallOption) (*TodoResponse, error)
}

type logisticsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogisticsServiceClient(cc grpc.ClientConnInterface) LogisticsServiceClient {
	return &logisticsServiceClient{cc}
}

func (c *logisticsServiceClient) Todo(ctx context.Context, in *TodoRequest, opts ...grpc.CallOption) (*TodoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TodoResponse)
	err := c.cc.Invoke(ctx, LogisticsService_Todo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogisticsServiceServer is the server API for LogisticsService service.
// All implementations must embed UnimplementedLogisticsServiceServer
// for forward compatibility
type LogisticsServiceServer interface {
	Todo(context.Context, *TodoRequest) (*TodoResponse, error)
	mustEmbedUnimplementedLogisticsServiceServer()
}

// UnimplementedLogisticsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLogisticsServiceServer struct {
}

func (UnimplementedLogisticsServiceServer) Todo(context.Context, *TodoRequest) (*TodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Todo not implemented")
}
func (UnimplementedLogisticsServiceServer) mustEmbedUnimplementedLogisticsServiceServer() {}

// UnsafeLogisticsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogisticsServiceServer will
// result in compilation errors.
type UnsafeLogisticsServiceServer interface {
	mustEmbedUnimplementedLogisticsServiceServer()
}

func RegisterLogisticsServiceServer(s grpc.ServiceRegistrar, srv LogisticsServiceServer) {
	s.RegisterService(&LogisticsService_ServiceDesc, srv)
}

func _LogisticsService_Todo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogisticsServiceServer).Todo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogisticsService_Todo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogisticsServiceServer).Todo(ctx, req.(*TodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LogisticsService_ServiceDesc is the grpc.ServiceDesc for LogisticsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogisticsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.LogisticsService",
	HandlerType: (*LogisticsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Todo",
			Handler:    _LogisticsService_Todo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "logistics_service.proto",
}
