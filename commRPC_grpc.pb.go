// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: commRPC.proto

package v1

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
	CommRPC_ExecNodeProxyRPC_FullMethodName = "/api.commRPC.v1.CommRPC/ExecNodeProxyRPC"
)

// CommRPCClient is the client API for CommRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommRPCClient interface {
	ExecNodeProxyRPC(ctx context.Context, in *ExecNodeProxyRPCRequest, opts ...grpc.CallOption) (*ExecNodeProxyRPCReply, error)
}

type commRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewCommRPCClient(cc grpc.ClientConnInterface) CommRPCClient {
	return &commRPCClient{cc}
}

func (c *commRPCClient) ExecNodeProxyRPC(ctx context.Context, in *ExecNodeProxyRPCRequest, opts ...grpc.CallOption) (*ExecNodeProxyRPCReply, error) {
	out := new(ExecNodeProxyRPCReply)
	err := c.cc.Invoke(ctx, CommRPC_ExecNodeProxyRPC_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommRPCServer is the server API for CommRPC service.
// All implementations must embed UnimplementedCommRPCServer
// for forward compatibility
type CommRPCServer interface {
	ExecNodeProxyRPC(context.Context, *ExecNodeProxyRPCRequest) (*ExecNodeProxyRPCReply, error)
	mustEmbedUnimplementedCommRPCServer()
}

// UnimplementedCommRPCServer must be embedded to have forward compatible implementations.
type UnimplementedCommRPCServer struct {
}

func (UnimplementedCommRPCServer) ExecNodeProxyRPC(context.Context, *ExecNodeProxyRPCRequest) (*ExecNodeProxyRPCReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecNodeProxyRPC not implemented")
}
func (UnimplementedCommRPCServer) mustEmbedUnimplementedCommRPCServer() {}

// UnsafeCommRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommRPCServer will
// result in compilation errors.
type UnsafeCommRPCServer interface {
	mustEmbedUnimplementedCommRPCServer()
}

func RegisterCommRPCServer(s grpc.ServiceRegistrar, srv CommRPCServer) {
	s.RegisterService(&CommRPC_ServiceDesc, srv)
}

func _CommRPC_ExecNodeProxyRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecNodeProxyRPCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommRPCServer).ExecNodeProxyRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommRPC_ExecNodeProxyRPC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommRPCServer).ExecNodeProxyRPC(ctx, req.(*ExecNodeProxyRPCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommRPC_ServiceDesc is the grpc.ServiceDesc for CommRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.commRPC.v1.CommRPC",
	HandlerType: (*CommRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecNodeProxyRPC",
			Handler:    _CommRPC_ExecNodeProxyRPC_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "commRPC.proto",
}
