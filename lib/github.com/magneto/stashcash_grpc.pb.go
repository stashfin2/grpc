// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.6
// source: stashcash.proto

package magneto

import (
	context "context"
	stashcash "github.com/magneto/stashcash"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StashcashClient is the client API for Stashcash service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StashcashClient interface {
	Getscbalance(ctx context.Context, in *stashcash.Request, opts ...grpc.CallOption) (*stashcash.Response, error)
	Creditsc(ctx context.Context, in *stashcash.Request, opts ...grpc.CallOption) (*stashcash.Response, error)
	Reversesc(ctx context.Context, in *stashcash.Request, opts ...grpc.CallOption) (*stashcash.Response, error)
	Getschistory(ctx context.Context, in *stashcash.Request, opts ...grpc.CallOption) (*stashcash.Response, error)
	Debitsc(ctx context.Context, in *stashcash.Request, opts ...grpc.CallOption) (*stashcash.Response, error)
}

type stashcashClient struct {
	cc grpc.ClientConnInterface
}

func NewStashcashClient(cc grpc.ClientConnInterface) StashcashClient {
	return &stashcashClient{cc}
}

func (c *stashcashClient) Getscbalance(ctx context.Context, in *stashcash.Request, opts ...grpc.CallOption) (*stashcash.Response, error) {
	out := new(stashcash.Response)
	err := c.cc.Invoke(ctx, "/service.stashcash/getscbalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stashcashClient) Creditsc(ctx context.Context, in *stashcash.Request, opts ...grpc.CallOption) (*stashcash.Response, error) {
	out := new(stashcash.Response)
	err := c.cc.Invoke(ctx, "/service.stashcash/creditsc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stashcashClient) Reversesc(ctx context.Context, in *stashcash.Request, opts ...grpc.CallOption) (*stashcash.Response, error) {
	out := new(stashcash.Response)
	err := c.cc.Invoke(ctx, "/service.stashcash/reversesc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stashcashClient) Getschistory(ctx context.Context, in *stashcash.Request, opts ...grpc.CallOption) (*stashcash.Response, error) {
	out := new(stashcash.Response)
	err := c.cc.Invoke(ctx, "/service.stashcash/getschistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stashcashClient) Debitsc(ctx context.Context, in *stashcash.Request, opts ...grpc.CallOption) (*stashcash.Response, error) {
	out := new(stashcash.Response)
	err := c.cc.Invoke(ctx, "/service.stashcash/debitsc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StashcashServer is the server API for Stashcash service.
// All implementations must embed UnimplementedStashcashServer
// for forward compatibility
type StashcashServer interface {
	Getscbalance(context.Context, *stashcash.Request) (*stashcash.Response, error)
	Creditsc(context.Context, *stashcash.Request) (*stashcash.Response, error)
	Reversesc(context.Context, *stashcash.Request) (*stashcash.Response, error)
	Getschistory(context.Context, *stashcash.Request) (*stashcash.Response, error)
	Debitsc(context.Context, *stashcash.Request) (*stashcash.Response, error)
	mustEmbedUnimplementedStashcashServer()
}

// UnimplementedStashcashServer must be embedded to have forward compatible implementations.
type UnimplementedStashcashServer struct {
}

func (UnimplementedStashcashServer) Getscbalance(context.Context, *stashcash.Request) (*stashcash.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Getscbalance not implemented")
}
func (UnimplementedStashcashServer) Creditsc(context.Context, *stashcash.Request) (*stashcash.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Creditsc not implemented")
}
func (UnimplementedStashcashServer) Reversesc(context.Context, *stashcash.Request) (*stashcash.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reversesc not implemented")
}
func (UnimplementedStashcashServer) Getschistory(context.Context, *stashcash.Request) (*stashcash.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Getschistory not implemented")
}
func (UnimplementedStashcashServer) Debitsc(context.Context, *stashcash.Request) (*stashcash.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Debitsc not implemented")
}
func (UnimplementedStashcashServer) mustEmbedUnimplementedStashcashServer() {}

// UnsafeStashcashServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StashcashServer will
// result in compilation errors.
type UnsafeStashcashServer interface {
	mustEmbedUnimplementedStashcashServer()
}

func RegisterStashcashServer(s grpc.ServiceRegistrar, srv StashcashServer) {
	s.RegisterService(&Stashcash_ServiceDesc, srv)
}

func _Stashcash_Getscbalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(stashcash.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StashcashServer).Getscbalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.stashcash/getscbalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StashcashServer).Getscbalance(ctx, req.(*stashcash.Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stashcash_Creditsc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(stashcash.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StashcashServer).Creditsc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.stashcash/creditsc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StashcashServer).Creditsc(ctx, req.(*stashcash.Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stashcash_Reversesc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(stashcash.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StashcashServer).Reversesc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.stashcash/reversesc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StashcashServer).Reversesc(ctx, req.(*stashcash.Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stashcash_Getschistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(stashcash.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StashcashServer).Getschistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.stashcash/getschistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StashcashServer).Getschistory(ctx, req.(*stashcash.Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stashcash_Debitsc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(stashcash.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StashcashServer).Debitsc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.stashcash/debitsc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StashcashServer).Debitsc(ctx, req.(*stashcash.Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Stashcash_ServiceDesc is the grpc.ServiceDesc for Stashcash service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Stashcash_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.stashcash",
	HandlerType: (*StashcashServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getscbalance",
			Handler:    _Stashcash_Getscbalance_Handler,
		},
		{
			MethodName: "creditsc",
			Handler:    _Stashcash_Creditsc_Handler,
		},
		{
			MethodName: "reversesc",
			Handler:    _Stashcash_Reversesc_Handler,
		},
		{
			MethodName: "getschistory",
			Handler:    _Stashcash_Getschistory_Handler,
		},
		{
			MethodName: "debitsc",
			Handler:    _Stashcash_Debitsc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stashcash.proto",
}
