// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: playground.proto

package boundary

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
	Playground_CreateUser_FullMethodName = "/proto.Playground/CreateUser"
	Playground_LoginUser_FullMethodName  = "/proto.Playground/LoginUser"
)

// PlaygroundClient is the client API for Playground service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlaygroundClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
}

type playgroundClient struct {
	cc grpc.ClientConnInterface
}

func NewPlaygroundClient(cc grpc.ClientConnInterface) PlaygroundClient {
	return &playgroundClient{cc}
}

func (c *playgroundClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, Playground_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playgroundClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, Playground_LoginUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlaygroundServer is the server API for Playground service.
// All implementations must embed UnimplementedPlaygroundServer
// for forward compatibility
type PlaygroundServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	mustEmbedUnimplementedPlaygroundServer()
}

// UnimplementedPlaygroundServer must be embedded to have forward compatible implementations.
type UnimplementedPlaygroundServer struct {
}

func (UnimplementedPlaygroundServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedPlaygroundServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedPlaygroundServer) mustEmbedUnimplementedPlaygroundServer() {}

// UnsafePlaygroundServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlaygroundServer will
// result in compilation errors.
type UnsafePlaygroundServer interface {
	mustEmbedUnimplementedPlaygroundServer()
}

func RegisterPlaygroundServer(s grpc.ServiceRegistrar, srv PlaygroundServer) {
	s.RegisterService(&Playground_ServiceDesc, srv)
}

func _Playground_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaygroundServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Playground_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaygroundServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Playground_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaygroundServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Playground_LoginUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaygroundServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Playground_ServiceDesc is the grpc.ServiceDesc for Playground service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Playground_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Playground",
	HandlerType: (*PlaygroundServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _Playground_CreateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _Playground_LoginUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "playground.proto",
}
