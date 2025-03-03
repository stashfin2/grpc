// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: loans/getpendingemiamount.proto

package loans

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

type GetPendingEmiAmountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customerid int32 `protobuf:"varint,1,opt,name=customerid,proto3" json:"customerid,omitempty"`
}

func (x *GetPendingEmiAmountRequest) Reset() {
	*x = GetPendingEmiAmountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loans_getpendingemiamount_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPendingEmiAmountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPendingEmiAmountRequest) ProtoMessage() {}

func (x *GetPendingEmiAmountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loans_getpendingemiamount_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPendingEmiAmountRequest.ProtoReflect.Descriptor instead.
func (*GetPendingEmiAmountRequest) Descriptor() ([]byte, []int) {
	return file_loans_getpendingemiamount_proto_rawDescGZIP(), []int{0}
}

func (x *GetPendingEmiAmountRequest) GetCustomerid() int32 {
	if x != nil {
		return x.Customerid
	}
	return 0
}

type GetPendingEmiAmountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PendingAmount int32 `protobuf:"varint,1,opt,name=pending_amount,json=pendingAmount,proto3" json:"pending_amount,omitempty"`
}

func (x *GetPendingEmiAmountResponse) Reset() {
	*x = GetPendingEmiAmountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loans_getpendingemiamount_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPendingEmiAmountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPendingEmiAmountResponse) ProtoMessage() {}

func (x *GetPendingEmiAmountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loans_getpendingemiamount_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPendingEmiAmountResponse.ProtoReflect.Descriptor instead.
func (*GetPendingEmiAmountResponse) Descriptor() ([]byte, []int) {
	return file_loans_getpendingemiamount_proto_rawDescGZIP(), []int{1}
}

func (x *GetPendingEmiAmountResponse) GetPendingAmount() int32 {
	if x != nil {
		return x.PendingAmount
	}
	return 0
}

var File_loans_getpendingemiamount_proto protoreflect.FileDescriptor

var file_loans_getpendingemiamount_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6c, 0x6f, 0x61, 0x6e, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x70, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x65, 0x6d, 0x69, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x19, 0x6c, 0x6f, 0x61, 0x6e, 0x73, 0x2e, 0x67, 0x65, 0x74, 0x70, 0x65, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x65, 0x6d, 0x69, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3c, 0x0a, 0x1a,
	0x67, 0x65, 0x74, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x45, 0x6d, 0x69, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x69, 0x64, 0x22, 0x44, 0x0a, 0x1b, 0x67, 0x65,
	0x74, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x45, 0x6d, 0x69, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x65, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0d, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_loans_getpendingemiamount_proto_rawDescOnce sync.Once
	file_loans_getpendingemiamount_proto_rawDescData = file_loans_getpendingemiamount_proto_rawDesc
)

func file_loans_getpendingemiamount_proto_rawDescGZIP() []byte {
	file_loans_getpendingemiamount_proto_rawDescOnce.Do(func() {
		file_loans_getpendingemiamount_proto_rawDescData = protoimpl.X.CompressGZIP(file_loans_getpendingemiamount_proto_rawDescData)
	})
	return file_loans_getpendingemiamount_proto_rawDescData
}

var file_loans_getpendingemiamount_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_loans_getpendingemiamount_proto_goTypes = []interface{}{
	(*GetPendingEmiAmountRequest)(nil),  // 0: loans.getpendingemiamount.getPendingEmiAmountRequest
	(*GetPendingEmiAmountResponse)(nil), // 1: loans.getpendingemiamount.getPendingEmiAmountResponse
}
var file_loans_getpendingemiamount_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_loans_getpendingemiamount_proto_init() }
func file_loans_getpendingemiamount_proto_init() {
	if File_loans_getpendingemiamount_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_loans_getpendingemiamount_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPendingEmiAmountRequest); i {
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
		file_loans_getpendingemiamount_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPendingEmiAmountResponse); i {
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
			RawDescriptor: file_loans_getpendingemiamount_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_loans_getpendingemiamount_proto_goTypes,
		DependencyIndexes: file_loans_getpendingemiamount_proto_depIdxs,
		MessageInfos:      file_loans_getpendingemiamount_proto_msgTypes,
	}.Build()
	File_loans_getpendingemiamount_proto = out.File
	file_loans_getpendingemiamount_proto_rawDesc = nil
	file_loans_getpendingemiamount_proto_goTypes = nil
	file_loans_getpendingemiamount_proto_depIdxs = nil
}
