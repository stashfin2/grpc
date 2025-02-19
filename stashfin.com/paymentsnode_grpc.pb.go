// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
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

// PaymentsnodeClient is the client API for Paymentsnode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentsnodeClient interface {
	Getpaymentoptions(ctx context.Context, in *payments.GetpaymentoptionsRequest, opts ...grpc.CallOption) (*payments.GetpaymentoptionsResponse, error)
	Initiatepayment(ctx context.Context, in *payments.InitiatepaymentRequest, opts ...grpc.CallOption) (*payments.InitiatepaymentResponse, error)
	Checkpaymentstatus(ctx context.Context, in *payments.CheckpaymentstatusRequest, opts ...grpc.CallOption) (*payments.CheckpaymentstatusResponse, error)
	Nachregistration(ctx context.Context, in *payments.NachregistrationRequest, opts ...grpc.CallOption) (*payments.NachregistrationResponse, error)
	Nachpresentation(ctx context.Context, in *payments.NachpresentationRequest, opts ...grpc.CallOption) (*payments.NachpresentationResponse, error)
	Fundtransfer(ctx context.Context, in *payments.FundtransferRequest, opts ...grpc.CallOption) (*payments.FundtransferResponse, error)
}

type paymentsnodeClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentsnodeClient(cc grpc.ClientConnInterface) PaymentsnodeClient {
	return &paymentsnodeClient{cc}
}

func (c *paymentsnodeClient) Getpaymentoptions(ctx context.Context, in *payments.GetpaymentoptionsRequest, opts ...grpc.CallOption) (*payments.GetpaymentoptionsResponse, error) {
	out := new(payments.GetpaymentoptionsResponse)
	err := c.cc.Invoke(ctx, "/service.paymentsnode/getpaymentoptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentsnodeClient) Initiatepayment(ctx context.Context, in *payments.InitiatepaymentRequest, opts ...grpc.CallOption) (*payments.InitiatepaymentResponse, error) {
	out := new(payments.InitiatepaymentResponse)
	err := c.cc.Invoke(ctx, "/service.paymentsnode/initiatepayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentsnodeClient) Checkpaymentstatus(ctx context.Context, in *payments.CheckpaymentstatusRequest, opts ...grpc.CallOption) (*payments.CheckpaymentstatusResponse, error) {
	out := new(payments.CheckpaymentstatusResponse)
	err := c.cc.Invoke(ctx, "/service.paymentsnode/checkpaymentstatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentsnodeClient) Nachregistration(ctx context.Context, in *payments.NachregistrationRequest, opts ...grpc.CallOption) (*payments.NachregistrationResponse, error) {
	out := new(payments.NachregistrationResponse)
	err := c.cc.Invoke(ctx, "/service.paymentsnode/nachregistration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentsnodeClient) Nachpresentation(ctx context.Context, in *payments.NachpresentationRequest, opts ...grpc.CallOption) (*payments.NachpresentationResponse, error) {
	out := new(payments.NachpresentationResponse)
	err := c.cc.Invoke(ctx, "/service.paymentsnode/nachpresentation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentsnodeClient) Fundtransfer(ctx context.Context, in *payments.FundtransferRequest, opts ...grpc.CallOption) (*payments.FundtransferResponse, error) {
	out := new(payments.FundtransferResponse)
	err := c.cc.Invoke(ctx, "/service.paymentsnode/fundtransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentsnodeServer is the server API for Paymentsnode service.
// All implementations must embed UnimplementedPaymentsnodeServer
// for forward compatibility
type PaymentsnodeServer interface {
	Getpaymentoptions(context.Context, *payments.GetpaymentoptionsRequest) (*payments.GetpaymentoptionsResponse, error)
	Initiatepayment(context.Context, *payments.InitiatepaymentRequest) (*payments.InitiatepaymentResponse, error)
	Checkpaymentstatus(context.Context, *payments.CheckpaymentstatusRequest) (*payments.CheckpaymentstatusResponse, error)
	Nachregistration(context.Context, *payments.NachregistrationRequest) (*payments.NachregistrationResponse, error)
	Nachpresentation(context.Context, *payments.NachpresentationRequest) (*payments.NachpresentationResponse, error)
	Fundtransfer(context.Context, *payments.FundtransferRequest) (*payments.FundtransferResponse, error)
	mustEmbedUnimplementedPaymentsnodeServer()
}

// UnimplementedPaymentsnodeServer must be embedded to have forward compatible implementations.
type UnimplementedPaymentsnodeServer struct {
}

func (UnimplementedPaymentsnodeServer) Getpaymentoptions(context.Context, *payments.GetpaymentoptionsRequest) (*payments.GetpaymentoptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Getpaymentoptions not implemented")
}
func (UnimplementedPaymentsnodeServer) Initiatepayment(context.Context, *payments.InitiatepaymentRequest) (*payments.InitiatepaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Initiatepayment not implemented")
}
func (UnimplementedPaymentsnodeServer) Checkpaymentstatus(context.Context, *payments.CheckpaymentstatusRequest) (*payments.CheckpaymentstatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Checkpaymentstatus not implemented")
}
func (UnimplementedPaymentsnodeServer) Nachregistration(context.Context, *payments.NachregistrationRequest) (*payments.NachregistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Nachregistration not implemented")
}
func (UnimplementedPaymentsnodeServer) Nachpresentation(context.Context, *payments.NachpresentationRequest) (*payments.NachpresentationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Nachpresentation not implemented")
}
func (UnimplementedPaymentsnodeServer) Fundtransfer(context.Context, *payments.FundtransferRequest) (*payments.FundtransferResponse, error) {
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
	in := new(payments.GetpaymentoptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentsnodeServer).Getpaymentoptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.paymentsnode/getpaymentoptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentsnodeServer).Getpaymentoptions(ctx, req.(*payments.GetpaymentoptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Paymentsnode_Initiatepayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(payments.InitiatepaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentsnodeServer).Initiatepayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.paymentsnode/initiatepayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentsnodeServer).Initiatepayment(ctx, req.(*payments.InitiatepaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Paymentsnode_Checkpaymentstatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(payments.CheckpaymentstatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentsnodeServer).Checkpaymentstatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.paymentsnode/checkpaymentstatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentsnodeServer).Checkpaymentstatus(ctx, req.(*payments.CheckpaymentstatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Paymentsnode_Nachregistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(payments.NachregistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentsnodeServer).Nachregistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.paymentsnode/nachregistration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentsnodeServer).Nachregistration(ctx, req.(*payments.NachregistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Paymentsnode_Nachpresentation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(payments.NachpresentationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentsnodeServer).Nachpresentation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.paymentsnode/nachpresentation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentsnodeServer).Nachpresentation(ctx, req.(*payments.NachpresentationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Paymentsnode_Fundtransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(payments.FundtransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentsnodeServer).Fundtransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.paymentsnode/fundtransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentsnodeServer).Fundtransfer(ctx, req.(*payments.FundtransferRequest))
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
