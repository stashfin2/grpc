// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.6
// source: customers/createcustomerreference.proto

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

type CreateCustomerReferenceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReferenceName string `protobuf:"bytes,1,opt,name=reference_name,json=referenceName,proto3" json:"reference_name,omitempty"`
	Relationship  string `protobuf:"bytes,2,opt,name=relationship,proto3" json:"relationship,omitempty"`
	MobileNumber  string `protobuf:"bytes,3,opt,name=mobile_number,json=mobileNumber,proto3" json:"mobile_number,omitempty"`
	Email         string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"` //optional
	Address       string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Pincode       string `protobuf:"bytes,6,opt,name=pincode,proto3" json:"pincode,omitempty"`
	StateId       int32  `protobuf:"varint,7,opt,name=state_id,json=stateId,proto3" json:"state_id,omitempty"`
	CityId        int32  `protobuf:"varint,8,opt,name=city_id,json=cityId,proto3" json:"city_id,omitempty"`
}

func (x *CreateCustomerReferenceRequest) Reset() {
	*x = CreateCustomerReferenceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_createcustomerreference_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCustomerReferenceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCustomerReferenceRequest) ProtoMessage() {}

func (x *CreateCustomerReferenceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customers_createcustomerreference_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCustomerReferenceRequest.ProtoReflect.Descriptor instead.
func (*CreateCustomerReferenceRequest) Descriptor() ([]byte, []int) {
	return file_customers_createcustomerreference_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCustomerReferenceRequest) GetReferenceName() string {
	if x != nil {
		return x.ReferenceName
	}
	return ""
}

func (x *CreateCustomerReferenceRequest) GetRelationship() string {
	if x != nil {
		return x.Relationship
	}
	return ""
}

func (x *CreateCustomerReferenceRequest) GetMobileNumber() string {
	if x != nil {
		return x.MobileNumber
	}
	return ""
}

func (x *CreateCustomerReferenceRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateCustomerReferenceRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CreateCustomerReferenceRequest) GetPincode() string {
	if x != nil {
		return x.Pincode
	}
	return ""
}

func (x *CreateCustomerReferenceRequest) GetStateId() int32 {
	if x != nil {
		return x.StateId
	}
	return 0
}

func (x *CreateCustomerReferenceRequest) GetCityId() int32 {
	if x != nil {
		return x.CityId
	}
	return 0
}

type CreateCustomerReferenceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReferenceId int32 `protobuf:"varint,1,opt,name=reference_id,json=referenceId,proto3" json:"reference_id,omitempty"`
}

func (x *CreateCustomerReferenceResponse) Reset() {
	*x = CreateCustomerReferenceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_createcustomerreference_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCustomerReferenceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCustomerReferenceResponse) ProtoMessage() {}

func (x *CreateCustomerReferenceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customers_createcustomerreference_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCustomerReferenceResponse.ProtoReflect.Descriptor instead.
func (*CreateCustomerReferenceResponse) Descriptor() ([]byte, []int) {
	return file_customers_createcustomerreference_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCustomerReferenceResponse) GetReferenceId() int32 {
	if x != nil {
		return x.ReferenceId
	}
	return 0
}

var File_customers_createcustomerreference_proto protoreflect.FileDescriptor

var file_customers_createcustomerreference_proto_rawDesc = []byte{
	0x0a, 0x27, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x73, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x8e, 0x02, 0x0a,
	0x1e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x25, 0x0a, 0x0e, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x69, 0x74, 0x79, 0x49, 0x64, 0x22, 0x44, 0x0a,
	0x1f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x49, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_customers_createcustomerreference_proto_rawDescOnce sync.Once
	file_customers_createcustomerreference_proto_rawDescData = file_customers_createcustomerreference_proto_rawDesc
)

func file_customers_createcustomerreference_proto_rawDescGZIP() []byte {
	file_customers_createcustomerreference_proto_rawDescOnce.Do(func() {
		file_customers_createcustomerreference_proto_rawDescData = protoimpl.X.CompressGZIP(file_customers_createcustomerreference_proto_rawDescData)
	})
	return file_customers_createcustomerreference_proto_rawDescData
}

var file_customers_createcustomerreference_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_customers_createcustomerreference_proto_goTypes = []interface{}{
	(*CreateCustomerReferenceRequest)(nil),  // 0: customers.createcustomerreference.createCustomerReferenceRequest
	(*CreateCustomerReferenceResponse)(nil), // 1: customers.createcustomerreference.createCustomerReferenceResponse
}
var file_customers_createcustomerreference_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_customers_createcustomerreference_proto_init() }
func file_customers_createcustomerreference_proto_init() {
	if File_customers_createcustomerreference_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_customers_createcustomerreference_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCustomerReferenceRequest); i {
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
		file_customers_createcustomerreference_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCustomerReferenceResponse); i {
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
			RawDescriptor: file_customers_createcustomerreference_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_customers_createcustomerreference_proto_goTypes,
		DependencyIndexes: file_customers_createcustomerreference_proto_depIdxs,
		MessageInfos:      file_customers_createcustomerreference_proto_msgTypes,
	}.Build()
	File_customers_createcustomerreference_proto = out.File
	file_customers_createcustomerreference_proto_rawDesc = nil
	file_customers_createcustomerreference_proto_goTypes = nil
	file_customers_createcustomerreference_proto_depIdxs = nil
}
