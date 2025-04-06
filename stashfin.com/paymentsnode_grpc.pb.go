// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.6
// source: paymentsnode.proto

package stashfin_com

import (
	context "context"
	payments "github.com/stashfin2/grpc/stashfin.com/payments"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Paymentsnode_Getpaymentoptions_FullMethodName  = "/service.paymentsnode/getpaymentoptions"
	Paymentsnode_Initiatepayment_FullMethodName    = "/service.paymentsnode/initiatepayment"
	Paymentsnode_Checkpaymentstatus_FullMethodName = "/service.paymentsnode/checkpaymentstatus"
	Paymentsnode_Nachregistration_FullMethodName   = "/service.paymentsnode/nachregistration"
	Paymentsnode_Nachpresentation_FullMethodName   = "/service.paymentsnode/nachpresentation"
	Paymentsnode_Fundtransfer_FullMethodName       = "/service.paymentsnode/fundtransfer"
)

// PaymentsnodeClient is the client API for Paymentsnode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentsnodeClient interface {
	Getpaymentoptions(ctx context.Context, in *payments.Request, opts ...grpc.CallOption) (*payments.Response, error)
	Initiatepayment(ctx context.Context, in *payments.Request, opts ...grpc.CallOption) (*payments.Response, error)
	Checkpaymentstatus(ctx context.Context, in *payments.Request, opts ...grpc.CallOption) (*payments.Response, error)
	Nachregistration(ctx context.Context, in *payments.Request, opts ...grpc.CallOption) (*payments.Response, error)
	Nachpresentation(ctx context.Context, in *payments.Request, opts ...grpc.CallOption) (*payments.Response, error)
	Fundtransfer(ctx context.Context, in *payments.Request, opts ...grpc.CallOption) (*payments.Response, error)
}

type paymentsnodeClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentsnodeClient(cc grpc.ClientConnInterface) PaymentsnodeClient {
	return &paymentsnodeClient{cc}
}

func (c *paymentsnodeClient) Getpaymentoptions(ctx context.Context, in *payments.Request, opts ...grpc.CallOption) (*payments.Response, error) {
	out := new(payments.Response)
	err := c.cc.Invoke(ctx, Paymentsnode_Getpaymentoptions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentsnodeClient) Initiatepayment(ctx context.Context, in *payments.Request, opts ...grpc.CallOption) (*payments.Response, error) {
	out := new(payments.Response)
	err := c.cc.Invoke(ctx, Paymentsnode_Initiatepayment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentsnodeClient) Checkpaymentstatus(ctx context.Context, in *payments.Request, opts ...grpc.CallOption) (*payments.Response, error) {
	out := new(payments.Response)
	err := c.cc.Invoke(ctx, Paymentsnode_Checkpaymentstatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentsnodeClient) Nachregistration(ctx context.Context, in *payments.Request, opts ...grpc.CallOption) (*payments.Response, error) {
	out := new(payments.Response)
	err := c.cc.Invoke(ctx, Paymentsnode_Nachregistration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentsnodeClient) Nachpresentation(ctx context.Context, in *payments.Request, opts ...grpc.CallOption) (*payments.Response, error) {
	out := new(payments.Response)
	err := c.cc.Invoke(ctx, Paymentsnode_Nachpresentation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentsnodeClient) Fundtransfer(ctx context.Context, in *payments.Request, opts ...grpc.CallOption) (*payments.Response, error) {
	out := new(payments.Response)
	err := c.cc.Invoke(ctx, Paymentsnode_Fundtransfer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentsnodeServer is the server API for Paymentsnode service.
// All implementations must embed UnimplementedPaymentsnodeServer
// for forward compatibility
type PaymentsnodeServer interface {
	Getpaymentoptions(context.Context, *payments.Request) (*payments.Response, error)
	Initiatepayment(context.Context, *payments.Request) (*payments.Response, error)
	Checkpaymentstatus(context.Context, *payments.Request) (*payments.Response, error)
	Nachregistration(context.Context, *payments.Request) (*payments.Response, error)
	Nachpresentation(context.Context, *payments.Request) (*payments.Response, error)
	Fundtransfer(context.Context, *payments.Request) (*payments.Response, error)
	mustEmbedUnimplementedPaymentsnodeServer()
}

// UnimplementedPaymentsnodeServer must be embedded to have forward compatible implementations.
type UnimplementedPaymentsnodeServer struct {
}

func (UnimplementedPaymentsnodeServer) Getpaymentoptions(context.Context, *payments.Request) (*payments.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Getpaymentoptions not implemented")
}
func (UnimplementedPaymentsnodeServer) Initiatepayment(context.Context, *payments.Request) (*payments.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Initiatepayment not implemented")
}
func (UnimplementedPaymentsnodeServer) Checkpaymentstatus(context.Context, *payments.Request) (*payments.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Checkpaymentstatus not implemented")
}
func (UnimplementedPaymentsnodeServer) Nachregistration(context.Context, *payments.Request) (*payments.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Nachregistration not implemented")
}
func (UnimplementedPaymentsnodeServer) Nachpresentation(context.Context, *payments.Request) (*payments.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Nachpresentation not implemented")
}
func (UnimplementedPaymentsnodeServer) Fundtransfer(context.Context, *payments.Request) (*payments.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fundtransfer not implemented")
}
func (UnimplementedPaymentsnodeServer) mustEmbedUnimplementedPaymentsnodeServer() {}

// UnsafePaymentsnodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentsnodeServer will
// result in compilation errors.
type UnsafePaymentsnodeServer interface {
	mustEmbedUnimplementedPaymentsnodeServer()
}

func RegisterPaymentsnodeServer(s grpc.ServiceRegistrar, srv PaymentsnodeServer) {
	s.RegisterService(&Paymentsnode_ServiceDesc, srv)
}

func _Paymentsnode_Getpaymentoptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(payments.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentsnodeServer).Getpaymentoptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Paymentsnode_Getpaymentoptions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentsnodeServer).Getpaymentoptions(ctx, req.(*payments.Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Paymentsnode_Initiatepayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(payments.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentsnodeServer).Initiatepayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Paymentsnode_Initiatepayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentsnodeServer).Initiatepayment(ctx, req.(*payments.Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Paymentsnode_Checkpaymentstatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(payments.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentsnodeServer).Checkpaymentstatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Paymentsnode_Checkpaymentstatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentsnodeServer).Checkpaymentstatus(ctx, req.(*payments.Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Paymentsnode_Nachregistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(payments.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentsnodeServer).Nachregistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Paymentsnode_Nachregistration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentsnodeServer).Nachregistration(ctx, req.(*payments.Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Paymentsnode_Nachpresentation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(payments.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentsnodeServer).Nachpresentation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Paymentsnode_Nachpresentation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentsnodeServer).Nachpresentation(ctx, req.(*payments.Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Paymentsnode_Fundtransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(payments.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentsnodeServer).Fundtransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Paymentsnode_Fundtransfer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentsnodeServer).Fundtransfer(ctx, req.(*payments.Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Paymentsnode_ServiceDesc is the grpc.ServiceDesc for Paymentsnode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Paymentsnode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.paymentsnode",
	HandlerType: (*PaymentsnodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getpaymentoptions",
			Handler:    _Paymentsnode_Getpaymentoptions_Handler,
		},
		{
			MethodName: "initiatepayment",
			Handler:    _Paymentsnode_Initiatepayment_Handler,
		},
		{
			MethodName: "checkpaymentstatus",
			Handler:    _Paymentsnode_Checkpaymentstatus_Handler,
		},
		{
			MethodName: "nachregistration",
			Handler:    _Paymentsnode_Nachregistration_Handler,
		},
		{
			MethodName: "nachpresentation",
			Handler:    _Paymentsnode_Nachpresentation_Handler,
		},
		{
			MethodName: "fundtransfer",
			Handler:    _Paymentsnode_Fundtransfer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "paymentsnode.proto",
}
