// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.19.6
// source: bureau/getdbrcolending.proto

package bureau

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

type DbrRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CustomerId    int64                  `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DbrRequest) Reset() {
	*x = DbrRequest{}
	mi := &file_bureau_getdbrcolending_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DbrRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DbrRequest) ProtoMessage() {}

func (x *DbrRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bureau_getdbrcolending_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DbrRequest.ProtoReflect.Descriptor instead.
func (*DbrRequest) Descriptor() ([]byte, []int) {
	return file_bureau_getdbrcolending_proto_rawDescGZIP(), []int{0}
}

func (x *DbrRequest) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

type DbrResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Dbr           int64                  `protobuf:"varint,2,opt,name=dbr,proto3" json:"dbr,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DbrResponse) Reset() {
	*x = DbrResponse{}
	mi := &file_bureau_getdbrcolending_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DbrResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DbrResponse) ProtoMessage() {}

func (x *DbrResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bureau_getdbrcolending_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DbrResponse.ProtoReflect.Descriptor instead.
func (*DbrResponse) Descriptor() ([]byte, []int) {
	return file_bureau_getdbrcolending_proto_rawDescGZIP(), []int{1}
}

func (x *DbrResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *DbrResponse) GetDbr() int64 {
	if x != nil {
		return x.Dbr
	}
	return 0
}

var File_bureau_getdbrcolending_proto protoreflect.FileDescriptor

var file_bureau_getdbrcolending_proto_rawDesc = string([]byte{
	0x0a, 0x1c, 0x62, 0x75, 0x72, 0x65, 0x61, 0x75, 0x2f, 0x67, 0x65, 0x74, 0x64, 0x62, 0x72, 0x63,
	0x6f, 0x6c, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16,
	0x62, 0x75, 0x72, 0x65, 0x61, 0x75, 0x2e, 0x67, 0x65, 0x74, 0x64, 0x62, 0x72, 0x63, 0x6f, 0x6c,
	0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x2d, 0x0a, 0x0a, 0x64, 0x62, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x37, 0x0a, 0x0b, 0x64, 0x62, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x64, 0x62, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x64, 0x62, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_bureau_getdbrcolending_proto_rawDescOnce sync.Once
	file_bureau_getdbrcolending_proto_rawDescData []byte
)

func file_bureau_getdbrcolending_proto_rawDescGZIP() []byte {
	file_bureau_getdbrcolending_proto_rawDescOnce.Do(func() {
		file_bureau_getdbrcolending_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_bureau_getdbrcolending_proto_rawDesc), len(file_bureau_getdbrcolending_proto_rawDesc)))
	})
	return file_bureau_getdbrcolending_proto_rawDescData
}

var file_bureau_getdbrcolending_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bureau_getdbrcolending_proto_goTypes = []any{
	(*DbrRequest)(nil),  // 0: bureau.getdbrcolending.dbrRequest
	(*DbrResponse)(nil), // 1: bureau.getdbrcolending.dbrResponse
}
var file_bureau_getdbrcolending_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bureau_getdbrcolending_proto_init() }
func file_bureau_getdbrcolending_proto_init() {
	if File_bureau_getdbrcolending_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_bureau_getdbrcolending_proto_rawDesc), len(file_bureau_getdbrcolending_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bureau_getdbrcolending_proto_goTypes,
		DependencyIndexes: file_bureau_getdbrcolending_proto_depIdxs,
		MessageInfos:      file_bureau_getdbrcolending_proto_msgTypes,
	}.Build()
	File_bureau_getdbrcolending_proto = out.File
	file_bureau_getdbrcolending_proto_goTypes = nil
	file_bureau_getdbrcolending_proto_depIdxs = nil
}
