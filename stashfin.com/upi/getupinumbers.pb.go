// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.19.6
// source: upi/getupinumbers.proto

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

type GetUpiNumbersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUpiNumbersRequest) Reset() {
	*x = GetUpiNumbersRequest{}
	mi := &file_upi_getupinumbers_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUpiNumbersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUpiNumbersRequest) ProtoMessage() {}

func (x *GetUpiNumbersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_upi_getupinumbers_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUpiNumbersRequest.ProtoReflect.Descriptor instead.
func (*GetUpiNumbersRequest) Descriptor() ([]byte, []int) {
	return file_upi_getupinumbers_proto_rawDescGZIP(), []int{0}
}

type UpiNumbersList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UpiNumber     *string                `protobuf:"bytes,1,opt,name=upi_number,json=upiNumber,proto3,oneof" json:"upi_number,omitempty"`
	IsPrimary     *string                `protobuf:"bytes,2,opt,name=is_primary,json=isPrimary,proto3,oneof" json:"is_primary,omitempty"`
	IsDeleted     *string                `protobuf:"bytes,3,opt,name=is_deleted,json=isDeleted,proto3,oneof" json:"is_deleted,omitempty"`
	Status        *string                `protobuf:"bytes,4,opt,name=status,proto3,oneof" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpiNumbersList) Reset() {
	*x = UpiNumbersList{}
	mi := &file_upi_getupinumbers_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpiNumbersList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpiNumbersList) ProtoMessage() {}

func (x *UpiNumbersList) ProtoReflect() protoreflect.Message {
	mi := &file_upi_getupinumbers_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpiNumbersList.ProtoReflect.Descriptor instead.
func (*UpiNumbersList) Descriptor() ([]byte, []int) {
	return file_upi_getupinumbers_proto_rawDescGZIP(), []int{1}
}

func (x *UpiNumbersList) GetUpiNumber() string {
	if x != nil && x.UpiNumber != nil {
		return *x.UpiNumber
	}
	return ""
}

func (x *UpiNumbersList) GetIsPrimary() string {
	if x != nil && x.IsPrimary != nil {
		return *x.IsPrimary
	}
	return ""
}

func (x *UpiNumbersList) GetIsDeleted() string {
	if x != nil && x.IsDeleted != nil {
		return *x.IsDeleted
	}
	return ""
}

func (x *UpiNumbersList) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

type UpiNumbersListData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Upinumbers    []*UpiNumbersList      `protobuf:"bytes,1,rep,name=upinumbers,proto3" json:"upinumbers,omitempty"`
	Vpa           *string                `protobuf:"bytes,2,opt,name=vpa,proto3,oneof" json:"vpa,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpiNumbersListData) Reset() {
	*x = UpiNumbersListData{}
	mi := &file_upi_getupinumbers_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpiNumbersListData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpiNumbersListData) ProtoMessage() {}

func (x *UpiNumbersListData) ProtoReflect() protoreflect.Message {
	mi := &file_upi_getupinumbers_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpiNumbersListData.ProtoReflect.Descriptor instead.
func (*UpiNumbersListData) Descriptor() ([]byte, []int) {
	return file_upi_getupinumbers_proto_rawDescGZIP(), []int{2}
}

func (x *UpiNumbersListData) GetUpinumbers() []*UpiNumbersList {
	if x != nil {
		return x.Upinumbers
	}
	return nil
}

func (x *UpiNumbersListData) GetVpa() string {
	if x != nil && x.Vpa != nil {
		return *x.Vpa
	}
	return ""
}

type GetUpiNumbersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data          *UpiNumbersListData    `protobuf:"bytes,3,opt,name=data,proto3,oneof" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUpiNumbersResponse) Reset() {
	*x = GetUpiNumbersResponse{}
	mi := &file_upi_getupinumbers_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUpiNumbersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUpiNumbersResponse) ProtoMessage() {}

func (x *GetUpiNumbersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_upi_getupinumbers_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUpiNumbersResponse.ProtoReflect.Descriptor instead.
func (*GetUpiNumbersResponse) Descriptor() ([]byte, []int) {
	return file_upi_getupinumbers_proto_rawDescGZIP(), []int{3}
}

func (x *GetUpiNumbersResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetUpiNumbersResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetUpiNumbersResponse) GetData() *UpiNumbersListData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_upi_getupinumbers_proto protoreflect.FileDescriptor

var file_upi_getupinumbers_proto_rawDesc = string([]byte{
	0x0a, 0x17, 0x75, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x74, 0x75, 0x70, 0x69, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x75, 0x70, 0x69, 0x2e, 0x67,
	0x65, 0x74, 0x55, 0x70, 0x69, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x22, 0x16, 0x0a, 0x14,
	0x67, 0x65, 0x74, 0x55, 0x70, 0x69, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0xd1, 0x01, 0x0a, 0x0e, 0x75, 0x70, 0x69, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0a, 0x75, 0x70, 0x69, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x75,
	0x70, 0x69, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x69,
	0x73, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x09, 0x69, 0x73, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x88, 0x01, 0x01, 0x12,
	0x22, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x75, 0x70, 0x69, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x76, 0x0a, 0x12, 0x75, 0x70, 0x69, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x41,
	0x0a, 0x0a, 0x75, 0x70, 0x69, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x75, 0x70, 0x69, 0x2e, 0x67, 0x65, 0x74, 0x55, 0x70, 0x69, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x75, 0x70, 0x69, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0a, 0x75, 0x70, 0x69, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x12, 0x15, 0x0a, 0x03, 0x76, 0x70, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x03, 0x76, 0x70, 0x61, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x76, 0x70, 0x61,
	0x22, 0x92, 0x01, 0x0a, 0x15, 0x67, 0x65, 0x74, 0x55, 0x70, 0x69, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3e, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x75, 0x70, 0x69,
	0x2e, 0x67, 0x65, 0x74, 0x55, 0x70, 0x69, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x75,
	0x70, 0x69, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_upi_getupinumbers_proto_rawDescOnce sync.Once
	file_upi_getupinumbers_proto_rawDescData []byte
)

func file_upi_getupinumbers_proto_rawDescGZIP() []byte {
	file_upi_getupinumbers_proto_rawDescOnce.Do(func() {
		file_upi_getupinumbers_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_upi_getupinumbers_proto_rawDesc), len(file_upi_getupinumbers_proto_rawDesc)))
	})
	return file_upi_getupinumbers_proto_rawDescData
}

var file_upi_getupinumbers_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_upi_getupinumbers_proto_goTypes = []any{
	(*GetUpiNumbersRequest)(nil),  // 0: upi.getUpiNumbers.getUpiNumbersRequest
	(*UpiNumbersList)(nil),        // 1: upi.getUpiNumbers.upiNumbersList
	(*UpiNumbersListData)(nil),    // 2: upi.getUpiNumbers.upiNumbersListData
	(*GetUpiNumbersResponse)(nil), // 3: upi.getUpiNumbers.getUpiNumbersResponse
}
var file_upi_getupinumbers_proto_depIdxs = []int32{
	1, // 0: upi.getUpiNumbers.upiNumbersListData.upinumbers:type_name -> upi.getUpiNumbers.upiNumbersList
	2, // 1: upi.getUpiNumbers.getUpiNumbersResponse.data:type_name -> upi.getUpiNumbers.upiNumbersListData
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_upi_getupinumbers_proto_init() }
func file_upi_getupinumbers_proto_init() {
	if File_upi_getupinumbers_proto != nil {
		return
	}
	file_upi_getupinumbers_proto_msgTypes[1].OneofWrappers = []any{}
	file_upi_getupinumbers_proto_msgTypes[2].OneofWrappers = []any{}
	file_upi_getupinumbers_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_upi_getupinumbers_proto_rawDesc), len(file_upi_getupinumbers_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_upi_getupinumbers_proto_goTypes,
		DependencyIndexes: file_upi_getupinumbers_proto_depIdxs,
		MessageInfos:      file_upi_getupinumbers_proto_msgTypes,
	}.Build()
	File_upi_getupinumbers_proto = out.File
	file_upi_getupinumbers_proto_goTypes = nil
	file_upi_getupinumbers_proto_depIdxs = nil
}
