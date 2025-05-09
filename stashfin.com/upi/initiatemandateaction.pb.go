// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: upi/initiatemandateaction.proto

package upi

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

type InitiateMandateActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxnId string `protobuf:"bytes,1,opt,name=txn_id,json=txnId,proto3" json:"txn_id,omitempty"`
	Umn   string `protobuf:"bytes,2,opt,name=umn,proto3" json:"umn,omitempty"`
}

func (x *InitiateMandateActionRequest) Reset() {
	*x = InitiateMandateActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upi_initiatemandateaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitiateMandateActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitiateMandateActionRequest) ProtoMessage() {}

func (x *InitiateMandateActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_upi_initiatemandateaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitiateMandateActionRequest.ProtoReflect.Descriptor instead.
func (*InitiateMandateActionRequest) Descriptor() ([]byte, []int) {
	return file_upi_initiatemandateaction_proto_rawDescGZIP(), []int{0}
}

func (x *InitiateMandateActionRequest) GetTxnId() string {
	if x != nil {
		return x.TxnId
	}
	return ""
}

func (x *InitiateMandateActionRequest) GetUmn() string {
	if x != nil {
		return x.Umn
	}
	return ""
}

type InitiateMandateActionData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeqNo   string `protobuf:"bytes,1,opt,name=seq_no,json=seqNo,proto3" json:"seq_no,omitempty"`
	TxnHash string `protobuf:"bytes,2,opt,name=txn_hash,json=txnHash,proto3" json:"txn_hash,omitempty"`
}

func (x *InitiateMandateActionData) Reset() {
	*x = InitiateMandateActionData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upi_initiatemandateaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitiateMandateActionData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitiateMandateActionData) ProtoMessage() {}

func (x *InitiateMandateActionData) ProtoReflect() protoreflect.Message {
	mi := &file_upi_initiatemandateaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitiateMandateActionData.ProtoReflect.Descriptor instead.
func (*InitiateMandateActionData) Descriptor() ([]byte, []int) {
	return file_upi_initiatemandateaction_proto_rawDescGZIP(), []int{1}
}

func (x *InitiateMandateActionData) GetSeqNo() string {
	if x != nil {
		return x.SeqNo
	}
	return ""
}

func (x *InitiateMandateActionData) GetTxnHash() string {
	if x != nil {
		return x.TxnHash
	}
	return ""
}

type InitiateMandateActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string                     `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string                     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *InitiateMandateActionData `protobuf:"bytes,3,opt,name=data,proto3,oneof" json:"data,omitempty"`
}

func (x *InitiateMandateActionResponse) Reset() {
	*x = InitiateMandateActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upi_initiatemandateaction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitiateMandateActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitiateMandateActionResponse) ProtoMessage() {}

func (x *InitiateMandateActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_upi_initiatemandateaction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitiateMandateActionResponse.ProtoReflect.Descriptor instead.
func (*InitiateMandateActionResponse) Descriptor() ([]byte, []int) {
	return file_upi_initiatemandateaction_proto_rawDescGZIP(), []int{2}
}

func (x *InitiateMandateActionResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *InitiateMandateActionResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *InitiateMandateActionResponse) GetData() *InitiateMandateActionData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_upi_initiatemandateaction_proto protoreflect.FileDescriptor

var file_upi_initiatemandateaction_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x75, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x6d, 0x61,
	0x6e, 0x64, 0x61, 0x74, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x19, 0x75, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x4d,
	0x61, 0x6e, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x47, 0x0a, 0x1c,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06,
	0x74, 0x78, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x78,
	0x6e, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x6d, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x6d, 0x6e, 0x22, 0x4d, 0x0a, 0x19, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74,
	0x65, 0x4d, 0x61, 0x6e, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x65, 0x71, 0x5f, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x65, 0x71, 0x4e, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x78, 0x6e,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x78, 0x6e,
	0x48, 0x61, 0x73, 0x68, 0x22, 0xa9, 0x01, 0x0a, 0x1d, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74,
	0x65, 0x4d, 0x61, 0x6e, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x4d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x75, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_upi_initiatemandateaction_proto_rawDescOnce sync.Once
	file_upi_initiatemandateaction_proto_rawDescData = file_upi_initiatemandateaction_proto_rawDesc
)

func file_upi_initiatemandateaction_proto_rawDescGZIP() []byte {
	file_upi_initiatemandateaction_proto_rawDescOnce.Do(func() {
		file_upi_initiatemandateaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_upi_initiatemandateaction_proto_rawDescData)
	})
	return file_upi_initiatemandateaction_proto_rawDescData
}

var file_upi_initiatemandateaction_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_upi_initiatemandateaction_proto_goTypes = []interface{}{
	(*InitiateMandateActionRequest)(nil),  // 0: upi.initiateMandateAction.initiateMandateActionRequest
	(*InitiateMandateActionData)(nil),     // 1: upi.initiateMandateAction.InitiateMandateActionData
	(*InitiateMandateActionResponse)(nil), // 2: upi.initiateMandateAction.initiateMandateActionResponse
}
var file_upi_initiatemandateaction_proto_depIdxs = []int32{
	1, // 0: upi.initiateMandateAction.initiateMandateActionResponse.data:type_name -> upi.initiateMandateAction.InitiateMandateActionData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_upi_initiatemandateaction_proto_init() }
func file_upi_initiatemandateaction_proto_init() {
	if File_upi_initiatemandateaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_upi_initiatemandateaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitiateMandateActionRequest); i {
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
		file_upi_initiatemandateaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitiateMandateActionData); i {
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
		file_upi_initiatemandateaction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitiateMandateActionResponse); i {
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
	file_upi_initiatemandateaction_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_upi_initiatemandateaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_upi_initiatemandateaction_proto_goTypes,
		DependencyIndexes: file_upi_initiatemandateaction_proto_depIdxs,
		MessageInfos:      file_upi_initiatemandateaction_proto_msgTypes,
	}.Build()
	File_upi_initiatemandateaction_proto = out.File
	file_upi_initiatemandateaction_proto_rawDesc = nil
	file_upi_initiatemandateaction_proto_goTypes = nil
	file_upi_initiatemandateaction_proto_depIdxs = nil
}
