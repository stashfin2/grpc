// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.6
// source: deveqxcustomers.proto

package stashfin_com

import (
	context "context"
	eqxcustomers "github.com/stashfin2/grpc/stashfin.com/eqxcustomers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Deveqxcustomers_SendOtp_FullMethodName              = "/service.deveqxcustomers/sendOtp"
	Deveqxcustomers_VerifyOtp_FullMethodName            = "/service.deveqxcustomers/verifyOtp"
	Deveqxcustomers_GetUserById_FullMethodName          = "/service.deveqxcustomers/getUserById"
	Deveqxcustomers_GetDashboard_FullMethodName         = "/service.deveqxcustomers/getDashboard"
	Deveqxcustomers_GetDashboardMainCard_FullMethodName = "/service.deveqxcustomers/getDashboardMainCard"
	Deveqxcustomers_UpdatePushId_FullMethodName         = "/service.deveqxcustomers/updatePushId"
	Deveqxcustomers_Getnotifications_FullMethodName     = "/service.deveqxcustomers/getnotifications"
	Deveqxcustomers_UpdateNotifications_FullMethodName  = "/service.deveqxcustomers/updateNotifications"
	Deveqxcustomers_VerifyToken_FullMethodName          = "/service.deveqxcustomers/VerifyToken"
	Deveqxcustomers_GetTokens_FullMethodName            = "/service.deveqxcustomers/getTokens"
	Deveqxcustomers_GetCustomerByMobile_FullMethodName  = "/service.deveqxcustomers/getCustomerByMobile"
)

// DeveqxcustomersClient is the client API for Deveqxcustomers service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeveqxcustomersClient interface {
	SendOtp(ctx context.Context, in *eqxcustomers.SendOtpRequest, opts ...grpc.CallOption) (*eqxcustomers.SendOtpResponse, error)
	VerifyOtp(ctx context.Context, in *eqxcustomers.VerifyOtpReqeust, opts ...grpc.CallOption) (*eqxcustomers.VerifyOtpResponse, error)
	GetUserById(ctx context.Context, in *eqxcustomers.GetCustomerByIdRequest, opts ...grpc.CallOption) (*eqxcustomers.GetCustomerByIdResponse, error)
	GetDashboard(ctx context.Context, in *eqxcustomers.GetDashboardRequest, opts ...grpc.CallOption) (*eqxcustomers.GetDashboardResponse, error)
	GetDashboardMainCard(ctx context.Context, in *eqxcustomers.GetDashboardMainCardRequest, opts ...grpc.CallOption) (*eqxcustomers.GetDashboardMainCardResponse, error)
	UpdatePushId(ctx context.Context, in *eqxcustomers.UpdatePushIdRequest, opts ...grpc.CallOption) (*eqxcustomers.UpdatePushIdResponse, error)
	Getnotifications(ctx context.Context, in *eqxcustomers.GetNotificationsRequest, opts ...grpc.CallOption) (*eqxcustomers.GetNotificationsResponse, error)
	UpdateNotifications(ctx context.Context, in *eqxcustomers.UpdateNotificationRequest, opts ...grpc.CallOption) (*eqxcustomers.UpdateNotificationResponse, error)
	VerifyToken(ctx context.Context, in *eqxcustomers.VerifyTokenRequest, opts ...grpc.CallOption) (*eqxcustomers.VerifyTokenResponse, error)
	GetTokens(ctx context.Context, in *eqxcustomers.GetTokensRequest, opts ...grpc.CallOption) (*eqxcustomers.GetTokensResponse, error)
	GetCustomerByMobile(ctx context.Context, in *eqxcustomers.GetCustomerByMobileRequest, opts ...grpc.CallOption) (*eqxcustomers.GetCustomerByMobileResponse, error)
}

type deveqxcustomersClient struct {
	cc grpc.ClientConnInterface
}

func NewDeveqxcustomersClient(cc grpc.ClientConnInterface) DeveqxcustomersClient {
	return &deveqxcustomersClient{cc}
}

func (c *deveqxcustomersClient) SendOtp(ctx context.Context, in *eqxcustomers.SendOtpRequest, opts ...grpc.CallOption) (*eqxcustomers.SendOtpResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(eqxcustomers.SendOtpResponse)
	err := c.cc.Invoke(ctx, Deveqxcustomers_SendOtp_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deveqxcustomersClient) VerifyOtp(ctx context.Context, in *eqxcustomers.VerifyOtpReqeust, opts ...grpc.CallOption) (*eqxcustomers.VerifyOtpResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(eqxcustomers.VerifyOtpResponse)
	err := c.cc.Invoke(ctx, Deveqxcustomers_VerifyOtp_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deveqxcustomersClient) GetUserById(ctx context.Context, in *eqxcustomers.GetCustomerByIdRequest, opts ...grpc.CallOption) (*eqxcustomers.GetCustomerByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(eqxcustomers.GetCustomerByIdResponse)
	err := c.cc.Invoke(ctx, Deveqxcustomers_GetUserById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deveqxcustomersClient) GetDashboard(ctx context.Context, in *eqxcustomers.GetDashboardRequest, opts ...grpc.CallOption) (*eqxcustomers.GetDashboardResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(eqxcustomers.GetDashboardResponse)
	err := c.cc.Invoke(ctx, Deveqxcustomers_GetDashboard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deveqxcustomersClient) GetDashboardMainCard(ctx context.Context, in *eqxcustomers.GetDashboardMainCardRequest, opts ...grpc.CallOption) (*eqxcustomers.GetDashboardMainCardResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(eqxcustomers.GetDashboardMainCardResponse)
	err := c.cc.Invoke(ctx, Deveqxcustomers_GetDashboardMainCard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deveqxcustomersClient) UpdatePushId(ctx context.Context, in *eqxcustomers.UpdatePushIdRequest, opts ...grpc.CallOption) (*eqxcustomers.UpdatePushIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(eqxcustomers.UpdatePushIdResponse)
	err := c.cc.Invoke(ctx, Deveqxcustomers_UpdatePushId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deveqxcustomersClient) Getnotifications(ctx context.Context, in *eqxcustomers.GetNotificationsRequest, opts ...grpc.CallOption) (*eqxcustomers.GetNotificationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(eqxcustomers.GetNotificationsResponse)
	err := c.cc.Invoke(ctx, Deveqxcustomers_Getnotifications_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deveqxcustomersClient) UpdateNotifications(ctx context.Context, in *eqxcustomers.UpdateNotificationRequest, opts ...grpc.CallOption) (*eqxcustomers.UpdateNotificationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(eqxcustomers.UpdateNotificationResponse)
	err := c.cc.Invoke(ctx, Deveqxcustomers_UpdateNotifications_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deveqxcustomersClient) VerifyToken(ctx context.Context, in *eqxcustomers.VerifyTokenRequest, opts ...grpc.CallOption) (*eqxcustomers.VerifyTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(eqxcustomers.VerifyTokenResponse)
	err := c.cc.Invoke(ctx, Deveqxcustomers_VerifyToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deveqxcustomersClient) GetTokens(ctx context.Context, in *eqxcustomers.GetTokensRequest, opts ...grpc.CallOption) (*eqxcustomers.GetTokensResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(eqxcustomers.GetTokensResponse)
	err := c.cc.Invoke(ctx, Deveqxcustomers_GetTokens_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deveqxcustomersClient) GetCustomerByMobile(ctx context.Context, in *eqxcustomers.GetCustomerByMobileRequest, opts ...grpc.CallOption) (*eqxcustomers.GetCustomerByMobileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(eqxcustomers.GetCustomerByMobileResponse)
	err := c.cc.Invoke(ctx, Deveqxcustomers_GetCustomerByMobile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeveqxcustomersServer is the server API for Deveqxcustomers service.
// All implementations must embed UnimplementedDeveqxcustomersServer
// for forward compatibility.
type DeveqxcustomersServer interface {
	SendOtp(context.Context, *eqxcustomers.SendOtpRequest) (*eqxcustomers.SendOtpResponse, error)
	VerifyOtp(context.Context, *eqxcustomers.VerifyOtpReqeust) (*eqxcustomers.VerifyOtpResponse, error)
	GetUserById(context.Context, *eqxcustomers.GetCustomerByIdRequest) (*eqxcustomers.GetCustomerByIdResponse, error)
	GetDashboard(context.Context, *eqxcustomers.GetDashboardRequest) (*eqxcustomers.GetDashboardResponse, error)
	GetDashboardMainCard(context.Context, *eqxcustomers.GetDashboardMainCardRequest) (*eqxcustomers.GetDashboardMainCardResponse, error)
	UpdatePushId(context.Context, *eqxcustomers.UpdatePushIdRequest) (*eqxcustomers.UpdatePushIdResponse, error)
	Getnotifications(context.Context, *eqxcustomers.GetNotificationsRequest) (*eqxcustomers.GetNotificationsResponse, error)
	UpdateNotifications(context.Context, *eqxcustomers.UpdateNotificationRequest) (*eqxcustomers.UpdateNotificationResponse, error)
	VerifyToken(context.Context, *eqxcustomers.VerifyTokenRequest) (*eqxcustomers.VerifyTokenResponse, error)
	GetTokens(context.Context, *eqxcustomers.GetTokensRequest) (*eqxcustomers.GetTokensResponse, error)
	GetCustomerByMobile(context.Context, *eqxcustomers.GetCustomerByMobileRequest) (*eqxcustomers.GetCustomerByMobileResponse, error)
	mustEmbedUnimplementedDeveqxcustomersServer()
}

// UnimplementedDeveqxcustomersServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDeveqxcustomersServer struct{}

func (UnimplementedDeveqxcustomersServer) SendOtp(context.Context, *eqxcustomers.SendOtpRequest) (*eqxcustomers.SendOtpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendOtp not implemented")
}
func (UnimplementedDeveqxcustomersServer) VerifyOtp(context.Context, *eqxcustomers.VerifyOtpReqeust) (*eqxcustomers.VerifyOtpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyOtp not implemented")
}
func (UnimplementedDeveqxcustomersServer) GetUserById(context.Context, *eqxcustomers.GetCustomerByIdRequest) (*eqxcustomers.GetCustomerByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedDeveqxcustomersServer) GetDashboard(context.Context, *eqxcustomers.GetDashboardRequest) (*eqxcustomers.GetDashboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDashboard not implemented")
}
func (UnimplementedDeveqxcustomersServer) GetDashboardMainCard(context.Context, *eqxcustomers.GetDashboardMainCardRequest) (*eqxcustomers.GetDashboardMainCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDashboardMainCard not implemented")
}
func (UnimplementedDeveqxcustomersServer) UpdatePushId(context.Context, *eqxcustomers.UpdatePushIdRequest) (*eqxcustomers.UpdatePushIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePushId not implemented")
}
func (UnimplementedDeveqxcustomersServer) Getnotifications(context.Context, *eqxcustomers.GetNotificationsRequest) (*eqxcustomers.GetNotificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Getnotifications not implemented")
}
func (UnimplementedDeveqxcustomersServer) UpdateNotifications(context.Context, *eqxcustomers.UpdateNotificationRequest) (*eqxcustomers.UpdateNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNotifications not implemented")
}
func (UnimplementedDeveqxcustomersServer) VerifyToken(context.Context, *eqxcustomers.VerifyTokenRequest) (*eqxcustomers.VerifyTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyToken not implemented")
}
func (UnimplementedDeveqxcustomersServer) GetTokens(context.Context, *eqxcustomers.GetTokensRequest) (*eqxcustomers.GetTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokens not implemented")
}
func (UnimplementedDeveqxcustomersServer) GetCustomerByMobile(context.Context, *eqxcustomers.GetCustomerByMobileRequest) (*eqxcustomers.GetCustomerByMobileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerByMobile not implemented")
}
func (UnimplementedDeveqxcustomersServer) mustEmbedUnimplementedDeveqxcustomersServer() {}
func (UnimplementedDeveqxcustomersServer) testEmbeddedByValue()                         {}

// UnsafeDeveqxcustomersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeveqxcustomersServer will
// result in compilation errors.
type UnsafeDeveqxcustomersServer interface {
	mustEmbedUnimplementedDeveqxcustomersServer()
}

func RegisterDeveqxcustomersServer(s grpc.ServiceRegistrar, srv DeveqxcustomersServer) {
	// If the following call pancis, it indicates UnimplementedDeveqxcustomersServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Deveqxcustomers_ServiceDesc, srv)
}

func _Deveqxcustomers_SendOtp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(eqxcustomers.SendOtpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeveqxcustomersServer).SendOtp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Deveqxcustomers_SendOtp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeveqxcustomersServer).SendOtp(ctx, req.(*eqxcustomers.SendOtpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deveqxcustomers_VerifyOtp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(eqxcustomers.VerifyOtpReqeust)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeveqxcustomersServer).VerifyOtp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Deveqxcustomers_VerifyOtp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeveqxcustomersServer).VerifyOtp(ctx, req.(*eqxcustomers.VerifyOtpReqeust))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deveqxcustomers_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(eqxcustomers.GetCustomerByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeveqxcustomersServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Deveqxcustomers_GetUserById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeveqxcustomersServer).GetUserById(ctx, req.(*eqxcustomers.GetCustomerByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deveqxcustomers_GetDashboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(eqxcustomers.GetDashboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeveqxcustomersServer).GetDashboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Deveqxcustomers_GetDashboard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeveqxcustomersServer).GetDashboard(ctx, req.(*eqxcustomers.GetDashboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deveqxcustomers_GetDashboardMainCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(eqxcustomers.GetDashboardMainCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeveqxcustomersServer).GetDashboardMainCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Deveqxcustomers_GetDashboardMainCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeveqxcustomersServer).GetDashboardMainCard(ctx, req.(*eqxcustomers.GetDashboardMainCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deveqxcustomers_UpdatePushId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(eqxcustomers.UpdatePushIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeveqxcustomersServer).UpdatePushId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Deveqxcustomers_UpdatePushId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeveqxcustomersServer).UpdatePushId(ctx, req.(*eqxcustomers.UpdatePushIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deveqxcustomers_Getnotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(eqxcustomers.GetNotificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeveqxcustomersServer).Getnotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Deveqxcustomers_Getnotifications_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeveqxcustomersServer).Getnotifications(ctx, req.(*eqxcustomers.GetNotificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deveqxcustomers_UpdateNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(eqxcustomers.UpdateNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeveqxcustomersServer).UpdateNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Deveqxcustomers_UpdateNotifications_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeveqxcustomersServer).UpdateNotifications(ctx, req.(*eqxcustomers.UpdateNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deveqxcustomers_VerifyToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(eqxcustomers.VerifyTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeveqxcustomersServer).VerifyToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Deveqxcustomers_VerifyToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeveqxcustomersServer).VerifyToken(ctx, req.(*eqxcustomers.VerifyTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deveqxcustomers_GetTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(eqxcustomers.GetTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeveqxcustomersServer).GetTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Deveqxcustomers_GetTokens_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeveqxcustomersServer).GetTokens(ctx, req.(*eqxcustomers.GetTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deveqxcustomers_GetCustomerByMobile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(eqxcustomers.GetCustomerByMobileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeveqxcustomersServer).GetCustomerByMobile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Deveqxcustomers_GetCustomerByMobile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeveqxcustomersServer).GetCustomerByMobile(ctx, req.(*eqxcustomers.GetCustomerByMobileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Deveqxcustomers_ServiceDesc is the grpc.ServiceDesc for Deveqxcustomers service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Deveqxcustomers_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.deveqxcustomers",
	HandlerType: (*DeveqxcustomersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "sendOtp",
			Handler:    _Deveqxcustomers_SendOtp_Handler,
		},
		{
			MethodName: "verifyOtp",
			Handler:    _Deveqxcustomers_VerifyOtp_Handler,
		},
		{
			MethodName: "getUserById",
			Handler:    _Deveqxcustomers_GetUserById_Handler,
		},
		{
			MethodName: "getDashboard",
			Handler:    _Deveqxcustomers_GetDashboard_Handler,
		},
		{
			MethodName: "getDashboardMainCard",
			Handler:    _Deveqxcustomers_GetDashboardMainCard_Handler,
		},
		{
			MethodName: "updatePushId",
			Handler:    _Deveqxcustomers_UpdatePushId_Handler,
		},
		{
			MethodName: "getnotifications",
			Handler:    _Deveqxcustomers_Getnotifications_Handler,
		},
		{
			MethodName: "updateNotifications",
			Handler:    _Deveqxcustomers_UpdateNotifications_Handler,
		},
		{
			MethodName: "VerifyToken",
			Handler:    _Deveqxcustomers_VerifyToken_Handler,
		},
		{
			MethodName: "getTokens",
			Handler:    _Deveqxcustomers_GetTokens_Handler,
		},
		{
			MethodName: "getCustomerByMobile",
			Handler:    _Deveqxcustomers_GetCustomerByMobile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "deveqxcustomers.proto",
}
