// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.5.1
// source: accounts/grpc.proto

package accounts

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ClientClient is the client API for Client service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientClient interface {
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*Account, error)
}

type clientClient struct {
	cc grpc.ClientConnInterface
}

func NewClientClient(cc grpc.ClientConnInterface) ClientClient {
	return &clientClient{cc}
}

func (c *clientClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/accounts.Client/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/accounts.Client/GetAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServer is the server API for Client service.
// All implementations should embed UnimplementedClientServer
// for forward compatibility
type ClientServer interface {
	CreateAccount(context.Context, *CreateAccountRequest) (*empty.Empty, error)
	GetAccount(context.Context, *GetAccountRequest) (*Account, error)
}

// UnimplementedClientServer should be embedded to have forward compatible implementations.
type UnimplementedClientServer struct {
}

func (UnimplementedClientServer) CreateAccount(context.Context, *CreateAccountRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedClientServer) GetAccount(context.Context, *GetAccountRequest) (*Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}

// UnsafeClientServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientServer will
// result in compilation errors.
type UnsafeClientServer interface {
	mustEmbedUnimplementedClientServer()
}

func RegisterClientServer(s grpc.ServiceRegistrar, srv ClientServer) {
	s.RegisterService(&Client_ServiceDesc, srv)
}

func _Client_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounts.Client/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounts.Client/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).GetAccount(ctx, req.(*GetAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Client_ServiceDesc is the grpc.ServiceDesc for Client service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Client_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accounts.Client",
	HandlerType: (*ClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _Client_CreateAccount_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _Client_GetAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accounts/grpc.proto",
}
