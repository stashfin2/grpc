// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.6
// source: documents.proto

package magneto

import (
	context "context"
	documents "github.com/magneto/documents"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DocumentsClient is the client API for Documents service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DocumentsClient interface {
	Getagreement(ctx context.Context, in *documents.Documentrequest, opts ...grpc.CallOption) (*documents.Documentresponse, error)
	Getfinancialdoc(ctx context.Context, in *documents.Financialdocrequest, opts ...grpc.CallOption) (*documents.Financialdocresponse, error)
}

type documentsClient struct {
	cc grpc.ClientConnInterface
}

func NewDocumentsClient(cc grpc.ClientConnInterface) DocumentsClient {
	return &documentsClient{cc}
}

func (c *documentsClient) Getagreement(ctx context.Context, in *documents.Documentrequest, opts ...grpc.CallOption) (*documents.Documentresponse, error) {
	out := new(documents.Documentresponse)
	err := c.cc.Invoke(ctx, "/service.documents/getagreement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentsClient) Getfinancialdoc(ctx context.Context, in *documents.Financialdocrequest, opts ...grpc.CallOption) (*documents.Financialdocresponse, error) {
	out := new(documents.Financialdocresponse)
	err := c.cc.Invoke(ctx, "/service.documents/getfinancialdoc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DocumentsServer is the server API for Documents service.
// All implementations must embed UnimplementedDocumentsServer
// for forward compatibility
type DocumentsServer interface {
	Getagreement(context.Context, *documents.Documentrequest) (*documents.Documentresponse, error)
	Getfinancialdoc(context.Context, *documents.Financialdocrequest) (*documents.Financialdocresponse, error)
	mustEmbedUnimplementedDocumentsServer()
}

// UnimplementedDocumentsServer must be embedded to have forward compatible implementations.
type UnimplementedDocumentsServer struct {
}

func (UnimplementedDocumentsServer) Getagreement(context.Context, *documents.Documentrequest) (*documents.Documentresponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Getagreement not implemented")
}
func (UnimplementedDocumentsServer) Getfinancialdoc(context.Context, *documents.Financialdocrequest) (*documents.Financialdocresponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Getfinancialdoc not implemented")
}
func (UnimplementedDocumentsServer) mustEmbedUnimplementedDocumentsServer() {}

// UnsafeDocumentsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DocumentsServer will
// result in compilation errors.
type UnsafeDocumentsServer interface {
	mustEmbedUnimplementedDocumentsServer()
}

func RegisterDocumentsServer(s grpc.ServiceRegistrar, srv DocumentsServer) {
	s.RegisterService(&Documents_ServiceDesc, srv)
}

func _Documents_Getagreement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(documents.Documentrequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentsServer).Getagreement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.documents/getagreement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentsServer).Getagreement(ctx, req.(*documents.Documentrequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Documents_Getfinancialdoc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(documents.Financialdocrequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentsServer).Getfinancialdoc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.documents/getfinancialdoc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentsServer).Getfinancialdoc(ctx, req.(*documents.Financialdocrequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Documents_ServiceDesc is the grpc.ServiceDesc for Documents service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Documents_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.documents",
	HandlerType: (*DocumentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getagreement",
			Handler:    _Documents_Getagreement_Handler,
		},
		{
			MethodName: "getfinancialdoc",
			Handler:    _Documents_Getfinancialdoc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "documents.proto",
}
