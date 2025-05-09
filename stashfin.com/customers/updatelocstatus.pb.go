// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: customers/updatelocstatus.proto

package customers

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

type LocStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocDisabled bool `protobuf:"varint,1,opt,name=loc_disabled,json=locDisabled,proto3" json:"loc_disabled,omitempty"`
}

func (x *LocStatusRequest) Reset() {
	*x = LocStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_updatelocstatus_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocStatusRequest) ProtoMessage() {}

func (x *LocStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customers_updatelocstatus_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocStatusRequest.ProtoReflect.Descriptor instead.
func (*LocStatusRequest) Descriptor() ([]byte, []int) {
	return file_customers_updatelocstatus_proto_rawDescGZIP(), []int{0}
}

func (x *LocStatusRequest) GetLocDisabled() bool {
	if x != nil {
		return x.LocDisabled
	}
	return false
}

type LocStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status            string  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	PendingLoanAmount float32 `protobuf:"fixed32,2,opt,name=pending_loan_amount,json=pendingLoanAmount,proto3" json:"pending_loan_amount,omitempty"`
}

func (x *LocStatusResponse) Reset() {
	*x = LocStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_updatelocstatus_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocStatusResponse) ProtoMessage() {}

func (x *LocStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customers_updatelocstatus_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocStatusResponse.ProtoReflect.Descriptor instead.
func (*LocStatusResponse) Descriptor() ([]byte, []int) {
	return file_customers_updatelocstatus_proto_rawDescGZIP(), []int{1}
}

func (x *LocStatusResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *LocStatusResponse) GetPendingLoanAmount() float32 {
	if x != nil {
		return x.PendingLoanAmount
	}
	return 0
}

var File_customers_updatelocstatus_proto protoreflect.FileDescriptor

var file_customers_updatelocstatus_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x6c, 0x6f, 0x63, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x19, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x6c, 0x6f, 0x63, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x35, 0x0a, 0x10,
	0x6c, 0x6f, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x6f, 0x63, 0x5f, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x6c, 0x6f, 0x63, 0x44, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x22, 0x5b, 0x0a, 0x11, 0x6c, 0x6f, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x2e, 0x0a, 0x13, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x6f, 0x61, 0x6e,
	0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x11, 0x70,
	0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x61, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_customers_updatelocstatus_proto_rawDescOnce sync.Once
	file_customers_updatelocstatus_proto_rawDescData = file_customers_updatelocstatus_proto_rawDesc
)

func file_customers_updatelocstatus_proto_rawDescGZIP() []byte {
	file_customers_updatelocstatus_proto_rawDescOnce.Do(func() {
		file_customers_updatelocstatus_proto_rawDescData = protoimpl.X.CompressGZIP(file_customers_updatelocstatus_proto_rawDescData)
	})
	return file_customers_updatelocstatus_proto_rawDescData
}

var file_customers_updatelocstatus_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_customers_updatelocstatus_proto_goTypes = []interface{}{
	(*LocStatusRequest)(nil),  // 0: customers.updatelocstatus.locStatusRequest
	(*LocStatusResponse)(nil), // 1: customers.updatelocstatus.locStatusResponse
}
var file_customers_updatelocstatus_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_customers_updatelocstatus_proto_init() }
func file_customers_updatelocstatus_proto_init() {
	if File_customers_updatelocstatus_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_customers_updatelocstatus_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocStatusRequest); i {
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
		file_customers_updatelocstatus_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocStatusResponse); i {
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
			RawDescriptor: file_customers_updatelocstatus_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_customers_updatelocstatus_proto_goTypes,
		DependencyIndexes: file_customers_updatelocstatus_proto_depIdxs,
		MessageInfos:      file_customers_updatelocstatus_proto_msgTypes,
	}.Build()
	File_customers_updatelocstatus_proto = out.File
	file_customers_updatelocstatus_proto_rawDesc = nil
	file_customers_updatelocstatus_proto_goTypes = nil
	file_customers_updatelocstatus_proto_depIdxs = nil
}
