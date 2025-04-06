// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.19.6
// source: upi/listblockedvpa.proto

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

type ListBlockedVPARequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BlockType     *string                `protobuf:"bytes,1,opt,name=block_type,json=blockType,proto3,oneof" json:"block_type,omitempty"`
	BlockValue    *string                `protobuf:"bytes,2,opt,name=block_value,json=blockValue,proto3,oneof" json:"block_value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListBlockedVPARequest) Reset() {
	*x = ListBlockedVPARequest{}
	mi := &file_upi_listblockedvpa_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListBlockedVPARequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBlockedVPARequest) ProtoMessage() {}

func (x *ListBlockedVPARequest) ProtoReflect() protoreflect.Message {
	mi := &file_upi_listblockedvpa_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBlockedVPARequest.ProtoReflect.Descriptor instead.
func (*ListBlockedVPARequest) Descriptor() ([]byte, []int) {
	return file_upi_listblockedvpa_proto_rawDescGZIP(), []int{0}
}

func (x *ListBlockedVPARequest) GetBlockType() string {
	if x != nil && x.BlockType != nil {
		return *x.BlockType
	}
	return ""
}

func (x *ListBlockedVPARequest) GetBlockValue() string {
	if x != nil && x.BlockValue != nil {
		return *x.BlockValue
	}
	return ""
}

type BlockedList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BlockType     string                 `protobuf:"bytes,1,opt,name=block_type,json=blockType,proto3" json:"block_type,omitempty"`
	BlockValue    string                 `protobuf:"bytes,2,opt,name=block_value,json=blockValue,proto3" json:"block_value,omitempty"`
	Reason        string                 `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty"`
	BlockedSince  string                 `protobuf:"bytes,4,opt,name=blocked_since,json=blockedSince,proto3" json:"blocked_since,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BlockedList) Reset() {
	*x = BlockedList{}
	mi := &file_upi_listblockedvpa_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BlockedList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockedList) ProtoMessage() {}

func (x *BlockedList) ProtoReflect() protoreflect.Message {
	mi := &file_upi_listblockedvpa_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockedList.ProtoReflect.Descriptor instead.
func (*BlockedList) Descriptor() ([]byte, []int) {
	return file_upi_listblockedvpa_proto_rawDescGZIP(), []int{1}
}

func (x *BlockedList) GetBlockType() string {
	if x != nil {
		return x.BlockType
	}
	return ""
}

func (x *BlockedList) GetBlockValue() string {
	if x != nil {
		return x.BlockValue
	}
	return ""
}

func (x *BlockedList) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *BlockedList) GetBlockedSince() string {
	if x != nil {
		return x.BlockedSince
	}
	return ""
}

type Data struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BlockedList   []*BlockedList         `protobuf:"bytes,1,rep,name=blocked_list,json=blockedList,proto3" json:"blocked_list,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Data) Reset() {
	*x = Data{}
	mi := &file_upi_listblockedvpa_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_upi_listblockedvpa_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_upi_listblockedvpa_proto_rawDescGZIP(), []int{2}
}

func (x *Data) GetBlockedList() []*BlockedList {
	if x != nil {
		return x.BlockedList
	}
	return nil
}

type ListBlockedVPAResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data          *Data                  `protobuf:"bytes,3,opt,name=data,proto3,oneof" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListBlockedVPAResponse) Reset() {
	*x = ListBlockedVPAResponse{}
	mi := &file_upi_listblockedvpa_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListBlockedVPAResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBlockedVPAResponse) ProtoMessage() {}

func (x *ListBlockedVPAResponse) ProtoReflect() protoreflect.Message {
	mi := &file_upi_listblockedvpa_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBlockedVPAResponse.ProtoReflect.Descriptor instead.
func (*ListBlockedVPAResponse) Descriptor() ([]byte, []int) {
	return file_upi_listblockedvpa_proto_rawDescGZIP(), []int{3}
}

func (x *ListBlockedVPAResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ListBlockedVPAResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ListBlockedVPAResponse) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_upi_listblockedvpa_proto protoreflect.FileDescriptor

var file_upi_listblockedvpa_proto_rawDesc = string([]byte{
	0x0a, 0x18, 0x75, 0x70, 0x69, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65,
	0x64, 0x76, 0x70, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x75, 0x70, 0x69, 0x2e,
	0x6c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x56, 0x50, 0x41, 0x22, 0x80,
	0x01, 0x0a, 0x15, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x56, 0x50,
	0x41, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x01, 0x52, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x88,
	0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x8a, 0x01, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x53, 0x69, 0x6e, 0x63, 0x65, 0x22, 0x4a,
	0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x42, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65,
	0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x75,
	0x70, 0x69, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x56, 0x50,
	0x41, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0b, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x86, 0x01, 0x0a, 0x16, 0x6c,
	0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x56, 0x50, 0x41, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x75, 0x70, 0x69, 0x2e, 0x6c, 0x69, 0x73, 0x74,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x56, 0x50, 0x41, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x48,
	0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_upi_listblockedvpa_proto_rawDescOnce sync.Once
	file_upi_listblockedvpa_proto_rawDescData []byte
)

func file_upi_listblockedvpa_proto_rawDescGZIP() []byte {
	file_upi_listblockedvpa_proto_rawDescOnce.Do(func() {
		file_upi_listblockedvpa_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_upi_listblockedvpa_proto_rawDesc), len(file_upi_listblockedvpa_proto_rawDesc)))
	})
	return file_upi_listblockedvpa_proto_rawDescData
}

var file_upi_listblockedvpa_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_upi_listblockedvpa_proto_goTypes = []any{
	(*ListBlockedVPARequest)(nil),  // 0: upi.listBlockedVPA.listBlockedVPARequest
	(*BlockedList)(nil),            // 1: upi.listBlockedVPA.BlockedList
	(*Data)(nil),                   // 2: upi.listBlockedVPA.Data
	(*ListBlockedVPAResponse)(nil), // 3: upi.listBlockedVPA.listBlockedVPAResponse
}
var file_upi_listblockedvpa_proto_depIdxs = []int32{
	1, // 0: upi.listBlockedVPA.Data.blocked_list:type_name -> upi.listBlockedVPA.BlockedList
	2, // 1: upi.listBlockedVPA.listBlockedVPAResponse.data:type_name -> upi.listBlockedVPA.Data
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_upi_listblockedvpa_proto_init() }
func file_upi_listblockedvpa_proto_init() {
	if File_upi_listblockedvpa_proto != nil {
		return
	}
	file_upi_listblockedvpa_proto_msgTypes[0].OneofWrappers = []any{}
	file_upi_listblockedvpa_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_upi_listblockedvpa_proto_rawDesc), len(file_upi_listblockedvpa_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_upi_listblockedvpa_proto_goTypes,
		DependencyIndexes: file_upi_listblockedvpa_proto_depIdxs,
		MessageInfos:      file_upi_listblockedvpa_proto_msgTypes,
	}.Build()
	File_upi_listblockedvpa_proto = out.File
	file_upi_listblockedvpa_proto_goTypes = nil
	file_upi_listblockedvpa_proto_depIdxs = nil
}
