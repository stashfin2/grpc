// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.19.6
// source: customers/updatesanctionamount.proto

package customers

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

type UpdateSanctionAmountRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	SanctionUpgrade bool                   `protobuf:"varint,1,opt,name=sanction_upgrade,json=sanctionUpgrade,proto3" json:"sanction_upgrade,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *UpdateSanctionAmountRequest) Reset() {
	*x = UpdateSanctionAmountRequest{}
	mi := &file_customers_updatesanctionamount_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateSanctionAmountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSanctionAmountRequest) ProtoMessage() {}

func (x *UpdateSanctionAmountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customers_updatesanctionamount_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSanctionAmountRequest.ProtoReflect.Descriptor instead.
func (*UpdateSanctionAmountRequest) Descriptor() ([]byte, []int) {
	return file_customers_updatesanctionamount_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateSanctionAmountRequest) GetSanctionUpgrade() bool {
	if x != nil {
		return x.SanctionUpgrade
	}
	return false
}

type UpdateSanctionAmountResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        bool                   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateSanctionAmountResponse) Reset() {
	*x = UpdateSanctionAmountResponse{}
	mi := &file_customers_updatesanctionamount_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateSanctionAmountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSanctionAmountResponse) ProtoMessage() {}

func (x *UpdateSanctionAmountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customers_updatesanctionamount_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSanctionAmountResponse.ProtoReflect.Descriptor instead.
func (*UpdateSanctionAmountResponse) Descriptor() ([]byte, []int) {
	return file_customers_updatesanctionamount_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateSanctionAmountResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

var File_customers_updatesanctionamount_proto protoreflect.FileDescriptor

var file_customers_updatesanctionamount_proto_rawDesc = string([]byte{
	0x0a, 0x24, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x73, 0x61, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x73, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x61, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x48, 0x0a, 0x1b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x61, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x61, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0f, 0x73, 0x61, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x22, 0x36, 0x0a, 0x1c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_customers_updatesanctionamount_proto_rawDescOnce sync.Once
	file_customers_updatesanctionamount_proto_rawDescData []byte
)

func file_customers_updatesanctionamount_proto_rawDescGZIP() []byte {
	file_customers_updatesanctionamount_proto_rawDescOnce.Do(func() {
		file_customers_updatesanctionamount_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_customers_updatesanctionamount_proto_rawDesc), len(file_customers_updatesanctionamount_proto_rawDesc)))
	})
	return file_customers_updatesanctionamount_proto_rawDescData
}

var file_customers_updatesanctionamount_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_customers_updatesanctionamount_proto_goTypes = []any{
	(*UpdateSanctionAmountRequest)(nil),  // 0: customers.updatesanctionamount.updateSanctionAmountRequest
	(*UpdateSanctionAmountResponse)(nil), // 1: customers.updatesanctionamount.updateSanctionAmountResponse
}
var file_customers_updatesanctionamount_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_customers_updatesanctionamount_proto_init() }
func file_customers_updatesanctionamount_proto_init() {
	if File_customers_updatesanctionamount_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_customers_updatesanctionamount_proto_rawDesc), len(file_customers_updatesanctionamount_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_customers_updatesanctionamount_proto_goTypes,
		DependencyIndexes: file_customers_updatesanctionamount_proto_depIdxs,
		MessageInfos:      file_customers_updatesanctionamount_proto_msgTypes,
	}.Build()
	File_customers_updatesanctionamount_proto = out.File
	file_customers_updatesanctionamount_proto_goTypes = nil
	file_customers_updatesanctionamount_proto_depIdxs = nil
}
