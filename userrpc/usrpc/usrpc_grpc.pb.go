// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: usrpc.proto

package usrpc

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

// UsrpcClient is the client API for Usrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsrpcClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type usrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewUsrpcClient(cc grpc.ClientConnInterface) UsrpcClient {
	return &usrpcClient{cc}
}

func (c *usrpcClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/usrpc.Usrpc/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsrpcServer is the server API for Usrpc service.
// All implementations must embed UnimplementedUsrpcServer
// for forward compatibility
type UsrpcServer interface {
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedUsrpcServer()
}

// UnimplementedUsrpcServer must be embedded to have forward compatible implementations.
type UnimplementedUsrpcServer struct {
}

func (UnimplementedUsrpcServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedUsrpcServer) mustEmbedUnimplementedUsrpcServer() {}

// UnsafeUsrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsrpcServer will
// result in compilation errors.
type UnsafeUsrpcServer interface {
	mustEmbedUnimplementedUsrpcServer()
}

func RegisterUsrpcServer(s grpc.ServiceRegistrar, srv UsrpcServer) {
	s.RegisterService(&Usrpc_ServiceDesc, srv)
}

func _Usrpc_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsrpcServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usrpc.Usrpc/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsrpcServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Usrpc_ServiceDesc is the grpc.ServiceDesc for Usrpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Usrpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "usrpc.Usrpc",
	HandlerType: (*UsrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Usrpc_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usrpc.proto",
}
