// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.19.6
// source: customers/freedom/addressconfirmation.proto

package freedom

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

type ConfirmAddressRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Address       string                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	AddressType   string                 `protobuf:"bytes,2,opt,name=address_type,json=addressType,proto3" json:"address_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConfirmAddressRequest) Reset() {
	*x = ConfirmAddressRequest{}
	mi := &file_customers_freedom_addressconfirmation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfirmAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmAddressRequest) ProtoMessage() {}

func (x *ConfirmAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customers_freedom_addressconfirmation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmAddressRequest.ProtoReflect.Descriptor instead.
func (*ConfirmAddressRequest) Descriptor() ([]byte, []int) {
	return file_customers_freedom_addressconfirmation_proto_rawDescGZIP(), []int{0}
}

func (x *ConfirmAddressRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ConfirmAddressRequest) GetAddressType() string {
	if x != nil {
		return x.AddressType
	}
	return ""
}

type ConfirmAddressResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Status        string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConfirmAddressResponse) Reset() {
	*x = ConfirmAddressResponse{}
	mi := &file_customers_freedom_addressconfirmation_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfirmAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmAddressResponse) ProtoMessage() {}

func (x *ConfirmAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customers_freedom_addressconfirmation_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmAddressResponse.ProtoReflect.Descriptor instead.
func (*ConfirmAddressResponse) Descriptor() ([]byte, []int) {
	return file_customers_freedom_addressconfirmation_proto_rawDescGZIP(), []int{1}
}

func (x *ConfirmAddressResponse) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *ConfirmAddressResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_customers_freedom_addressconfirmation_proto protoreflect.FileDescriptor

var file_customers_freedom_addressconfirmation_proto_rawDesc = string([]byte{
	0x0a, 0x2b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x66, 0x72, 0x65, 0x65,
	0x64, 0x6f, 0x6d, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x66, 0x72, 0x65, 0x65, 0x64, 0x6f, 0x6d,
	0x2e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x54, 0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4b, 0x0a, 0x16, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_customers_freedom_addressconfirmation_proto_rawDescOnce sync.Once
	file_customers_freedom_addressconfirmation_proto_rawDescData []byte
)

func file_customers_freedom_addressconfirmation_proto_rawDescGZIP() []byte {
	file_customers_freedom_addressconfirmation_proto_rawDescOnce.Do(func() {
		file_customers_freedom_addressconfirmation_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_customers_freedom_addressconfirmation_proto_rawDesc), len(file_customers_freedom_addressconfirmation_proto_rawDesc)))
	})
	return file_customers_freedom_addressconfirmation_proto_rawDescData
}

var file_customers_freedom_addressconfirmation_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_customers_freedom_addressconfirmation_proto_goTypes = []any{
	(*ConfirmAddressRequest)(nil),  // 0: customers.freedom.addressconfirmation.confirmAddressRequest
	(*ConfirmAddressResponse)(nil), // 1: customers.freedom.addressconfirmation.confirmAddressResponse
}
var file_customers_freedom_addressconfirmation_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_customers_freedom_addressconfirmation_proto_init() }
func file_customers_freedom_addressconfirmation_proto_init() {
	if File_customers_freedom_addressconfirmation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_customers_freedom_addressconfirmation_proto_rawDesc), len(file_customers_freedom_addressconfirmation_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_customers_freedom_addressconfirmation_proto_goTypes,
		DependencyIndexes: file_customers_freedom_addressconfirmation_proto_depIdxs,
		MessageInfos:      file_customers_freedom_addressconfirmation_proto_msgTypes,
	}.Build()
	File_customers_freedom_addressconfirmation_proto = out.File
	file_customers_freedom_addressconfirmation_proto_goTypes = nil
	file_customers_freedom_addressconfirmation_proto_depIdxs = nil
}
