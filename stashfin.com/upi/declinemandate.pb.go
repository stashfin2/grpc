// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.19.6
// source: upi/declinemandate.proto

package upi

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DeclineMandateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TxnId         string                 `protobuf:"bytes,1,opt,name=txn_id,json=txnId,proto3" json:"txn_id,omitempty"`
	Umn           string                 `protobuf:"bytes,2,opt,name=umn,proto3" json:"umn,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeclineMandateRequest) Reset() {
	*x = DeclineMandateRequest{}
	mi := &file_upi_declinemandate_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeclineMandateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeclineMandateRequest) ProtoMessage() {}

func (x *DeclineMandateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_upi_declinemandate_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeclineMandateRequest.ProtoReflect.Descriptor instead.
func (*DeclineMandateRequest) Descriptor() ([]byte, []int) {
	return file_upi_declinemandate_proto_rawDescGZIP(), []int{0}
}

func (x *DeclineMandateRequest) GetTxnId() string {
	if x != nil {
		return x.TxnId
	}
	return ""
}

func (x *DeclineMandateRequest) GetUmn() string {
	if x != nil {
		return x.Umn
	}
	return ""
}

type DeclineMandateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeclineMandateResponse) Reset() {
	*x = DeclineMandateResponse{}
	mi := &file_upi_declinemandate_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeclineMandateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeclineMandateResponse) ProtoMessage() {}

func (x *DeclineMandateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_upi_declinemandate_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeclineMandateResponse.ProtoReflect.Descriptor instead.
func (*DeclineMandateResponse) Descriptor() ([]byte, []int) {
	return file_upi_declinemandate_proto_rawDescGZIP(), []int{1}
}

func (x *DeclineMandateResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *DeclineMandateResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_upi_declinemandate_proto protoreflect.FileDescriptor

var file_upi_declinemandate_proto_rawDesc = string([]byte{
	0x0a, 0x18, 0x75, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x6e,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x75, 0x70, 0x69, 0x2e,
	0x64, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65, 0x4d, 0x61, 0x6e, 0x64, 0x61, 0x74, 0x65, 0x22, 0x40,
	0x0a, 0x15, 0x64, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65, 0x4d, 0x61, 0x6e, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x78, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x78, 0x6e, 0x49, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x6d, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x6d, 0x6e,
	0x22, 0x4a, 0x0a, 0x16, 0x64, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65, 0x4d, 0x61, 0x6e, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_upi_declinemandate_proto_rawDescOnce sync.Once
	file_upi_declinemandate_proto_rawDescData []byte
)

func file_upi_declinemandate_proto_rawDescGZIP() []byte {
	file_upi_declinemandate_proto_rawDescOnce.Do(func() {
		file_upi_declinemandate_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_upi_declinemandate_proto_rawDesc), len(file_upi_declinemandate_proto_rawDesc)))
	})
	return file_upi_declinemandate_proto_rawDescData
}

var file_upi_declinemandate_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_upi_declinemandate_proto_goTypes = []any{
	(*DeclineMandateRequest)(nil),  // 0: upi.declineMandate.declineMandateRequest
	(*DeclineMandateResponse)(nil), // 1: upi.declineMandate.declineMandateResponse
}
var file_upi_declinemandate_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_upi_declinemandate_proto_init() }
func file_upi_declinemandate_proto_init() {
	if File_upi_declinemandate_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_upi_declinemandate_proto_rawDesc), len(file_upi_declinemandate_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_upi_declinemandate_proto_goTypes,
		DependencyIndexes: file_upi_declinemandate_proto_depIdxs,
		MessageInfos:      file_upi_declinemandate_proto_msgTypes,
	}.Build()
	File_upi_declinemandate_proto = out.File
	file_upi_declinemandate_proto_goTypes = nil
	file_upi_declinemandate_proto_depIdxs = nil
}
