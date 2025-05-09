// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: payments/nachpresentation.proto

package payments

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NachPresentationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId      int64   `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Amount          float32 `protobuf:"fixed32,2,opt,name=amount,proto3" json:"amount,omitempty"`
	ClientId        string  `protobuf:"bytes,3,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ExternalOrderId string  `protobuf:"bytes,4,opt,name=external_order_id,json=externalOrderId,proto3" json:"external_order_id,omitempty"`
}

func (x *NachPresentationRequest) Reset() {
	*x = NachPresentationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payments_nachpresentation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NachPresentationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NachPresentationRequest) ProtoMessage() {}

func (x *NachPresentationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payments_nachpresentation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NachPresentationRequest.ProtoReflect.Descriptor instead.
func (*NachPresentationRequest) Descriptor() ([]byte, []int) {
	return file_payments_nachpresentation_proto_rawDescGZIP(), []int{0}
}

func (x *NachPresentationRequest) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *NachPresentationRequest) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *NachPresentationRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *NachPresentationRequest) GetExternalOrderId() string {
	if x != nil {
		return x.ExternalOrderId
	}
	return ""
}

type NachPresentationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	StatusCode int32  `protobuf:"varint,2,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
}

func (x *NachPresentationResponse) Reset() {
	*x = NachPresentationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payments_nachpresentation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NachPresentationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NachPresentationResponse) ProtoMessage() {}

func (x *NachPresentationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payments_nachpresentation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NachPresentationResponse.ProtoReflect.Descriptor instead.
func (*NachPresentationResponse) Descriptor() ([]byte, []int) {
	return file_payments_nachpresentation_proto_rawDescGZIP(), []int{1}
}

func (x *NachPresentationResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *NachPresentationResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

var File_payments_nachpresentation_proto protoreflect.FileDescriptor

var file_payments_nachpresentation_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x6e, 0x61, 0x63, 0x68, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x19, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x6e, 0x61, 0x63, 0x68,
	0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x9b, 0x01, 0x0a,
	0x17, 0x6e, 0x61, 0x63, 0x68, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2a,
	0x0a, 0x11, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x53, 0x0a, 0x18, 0x6e, 0x61,
	0x63, 0x68, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payments_nachpresentation_proto_rawDescOnce sync.Once
	file_payments_nachpresentation_proto_rawDescData = file_payments_nachpresentation_proto_rawDesc
)

func file_payments_nachpresentation_proto_rawDescGZIP() []byte {
	file_payments_nachpresentation_proto_rawDescOnce.Do(func() {
		file_payments_nachpresentation_proto_rawDescData = protoimpl.X.CompressGZIP(file_payments_nachpresentation_proto_rawDescData)
	})
	return file_payments_nachpresentation_proto_rawDescData
}

var file_payments_nachpresentation_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_payments_nachpresentation_proto_goTypes = []interface{}{
	(*NachPresentationRequest)(nil),  // 0: payments.nachpresentation.nachPresentationRequest
	(*NachPresentationResponse)(nil), // 1: payments.nachpresentation.nachPresentationResponse
}
var file_payments_nachpresentation_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_payments_nachpresentation_proto_init() }
func file_payments_nachpresentation_proto_init() {
	if File_payments_nachpresentation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payments_nachpresentation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NachPresentationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_payments_nachpresentation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NachPresentationResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_payments_nachpresentation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_payments_nachpresentation_proto_goTypes,
		DependencyIndexes: file_payments_nachpresentation_proto_depIdxs,
		MessageInfos:      file_payments_nachpresentation_proto_msgTypes,
	}.Build()
	File_payments_nachpresentation_proto = out.File
	file_payments_nachpresentation_proto_rawDesc = nil
	file_payments_nachpresentation_proto_goTypes = nil
	file_payments_nachpresentation_proto_depIdxs = nil
}
