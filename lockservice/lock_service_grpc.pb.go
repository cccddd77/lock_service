// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: lockservice/lock_service.proto

package lockservice

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

// LockServiceClient is the client API for LockService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LockServiceClient interface {
	DoLock(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Rsp, error)
	UnLock(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Rsp, error)
}

type lockServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLockServiceClient(cc grpc.ClientConnInterface) LockServiceClient {
	return &lockServiceClient{cc}
}

func (c *lockServiceClient) DoLock(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Rsp, error) {
	out := new(Rsp)
	err := c.cc.Invoke(ctx, "/lockservice.LockService/DoLock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lockServiceClient) UnLock(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Rsp, error) {
	out := new(Rsp)
	err := c.cc.Invoke(ctx, "/lockservice.LockService/UnLock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LockServiceServer is the server API for LockService service.
// All implementations must embed UnimplementedLockServiceServer
// for forward compatibility
type LockServiceServer interface {
	DoLock(context.Context, *Req) (*Rsp, error)
	UnLock(context.Context, *Req) (*Rsp, error)
	mustEmbedUnimplementedLockServiceServer()
}

// UnimplementedLockServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLockServiceServer struct {
}

func (UnimplementedLockServiceServer) DoLock(context.Context, *Req) (*Rsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoLock not implemented")
}
func (UnimplementedLockServiceServer) UnLock(context.Context, *Req) (*Rsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnLock not implemented")
}
func (UnimplementedLockServiceServer) mustEmbedUnimplementedLockServiceServer() {}

// UnsafeLockServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LockServiceServer will
// result in compilation errors.
type UnsafeLockServiceServer interface {
	mustEmbedUnimplementedLockServiceServer()
}

func RegisterLockServiceServer(s grpc.ServiceRegistrar, srv LockServiceServer) {
	s.RegisterService(&LockService_ServiceDesc, srv)
}

func _LockService_DoLock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LockServiceServer).DoLock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lockservice.LockService/DoLock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LockServiceServer).DoLock(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _LockService_UnLock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LockServiceServer).UnLock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lockservice.LockService/UnLock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LockServiceServer).UnLock(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

// LockService_ServiceDesc is the grpc.ServiceDesc for LockService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LockService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lockservice.LockService",
	HandlerType: (*LockServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoLock",
			Handler:    _LockService_DoLock_Handler,
		},
		{
			MethodName: "UnLock",
			Handler:    _LockService_UnLock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lockservice/lock_service.proto",
}
