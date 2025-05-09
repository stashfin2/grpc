// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: customers/addmoredetails.proto

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

type AddMoreDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FatherName       string  `protobuf:"bytes,1,opt,name=father_name,json=fatherName,proto3" json:"father_name,omitempty"`
	Occupation       string  `protobuf:"bytes,2,opt,name=occupation,proto3" json:"occupation,omitempty"`
	CurHouse         *string `protobuf:"bytes,3,opt,name=cur_house,json=curHouse,proto3,oneof" json:"cur_house,omitempty"`
	CurAddressLine1  string  `protobuf:"bytes,4,opt,name=cur_address_line1,json=curAddressLine1,proto3" json:"cur_address_line1,omitempty"`
	CurAddressLine2  string  `protobuf:"bytes,5,opt,name=cur_address_line2,json=curAddressLine2,proto3" json:"cur_address_line2,omitempty"`
	CurPincode       string  `protobuf:"bytes,6,opt,name=cur_pincode,json=curPincode,proto3" json:"cur_pincode,omitempty"`
	IsSamePerAddress bool    `protobuf:"varint,7,opt,name=is_same_per_address,json=isSamePerAddress,proto3" json:"is_same_per_address,omitempty"`
	PerHouse         *string `protobuf:"bytes,8,opt,name=per_house,json=perHouse,proto3,oneof" json:"per_house,omitempty"`
	PerAddressLine1  *string `protobuf:"bytes,9,opt,name=per_address_line1,json=perAddressLine1,proto3,oneof" json:"per_address_line1,omitempty"`
	PerAddressLine2  *string `protobuf:"bytes,10,opt,name=per_address_line2,json=perAddressLine2,proto3,oneof" json:"per_address_line2,omitempty"`
	PerPincode       *string `protobuf:"bytes,11,opt,name=per_pincode,json=perPincode,proto3,oneof" json:"per_pincode,omitempty"`
}

func (x *AddMoreDetailsRequest) Reset() {
	*x = AddMoreDetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_addmoredetails_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMoreDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMoreDetailsRequest) ProtoMessage() {}

func (x *AddMoreDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customers_addmoredetails_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMoreDetailsRequest.ProtoReflect.Descriptor instead.
func (*AddMoreDetailsRequest) Descriptor() ([]byte, []int) {
	return file_customers_addmoredetails_proto_rawDescGZIP(), []int{0}
}

func (x *AddMoreDetailsRequest) GetFatherName() string {
	if x != nil {
		return x.FatherName
	}
	return ""
}

func (x *AddMoreDetailsRequest) GetOccupation() string {
	if x != nil {
		return x.Occupation
	}
	return ""
}

func (x *AddMoreDetailsRequest) GetCurHouse() string {
	if x != nil && x.CurHouse != nil {
		return *x.CurHouse
	}
	return ""
}

func (x *AddMoreDetailsRequest) GetCurAddressLine1() string {
	if x != nil {
		return x.CurAddressLine1
	}
	return ""
}

func (x *AddMoreDetailsRequest) GetCurAddressLine2() string {
	if x != nil {
		return x.CurAddressLine2
	}
	return ""
}

func (x *AddMoreDetailsRequest) GetCurPincode() string {
	if x != nil {
		return x.CurPincode
	}
	return ""
}

func (x *AddMoreDetailsRequest) GetIsSamePerAddress() bool {
	if x != nil {
		return x.IsSamePerAddress
	}
	return false
}

func (x *AddMoreDetailsRequest) GetPerHouse() string {
	if x != nil && x.PerHouse != nil {
		return *x.PerHouse
	}
	return ""
}

func (x *AddMoreDetailsRequest) GetPerAddressLine1() string {
	if x != nil && x.PerAddressLine1 != nil {
		return *x.PerAddressLine1
	}
	return ""
}

func (x *AddMoreDetailsRequest) GetPerAddressLine2() string {
	if x != nil && x.PerAddressLine2 != nil {
		return *x.PerAddressLine2
	}
	return ""
}

func (x *AddMoreDetailsRequest) GetPerPincode() string {
	if x != nil && x.PerPincode != nil {
		return *x.PerPincode
	}
	return ""
}

type AddMoreDetailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AddMoreDetailsResponse) Reset() {
	*x = AddMoreDetailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_addmoredetails_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMoreDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMoreDetailsResponse) ProtoMessage() {}

func (x *AddMoreDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customers_addmoredetails_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMoreDetailsResponse.ProtoReflect.Descriptor instead.
func (*AddMoreDetailsResponse) Descriptor() ([]byte, []int) {
	return file_customers_addmoredetails_proto_rawDescGZIP(), []int{1}
}

func (x *AddMoreDetailsResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_customers_addmoredetails_proto protoreflect.FileDescriptor

var file_customers_addmoredetails_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x64, 0x64, 0x6d,
	0x6f, 0x72, 0x65, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x18, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x61, 0x64, 0x64, 0x6d,
	0x6f, 0x72, 0x65, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0xa4, 0x04, 0x0a, 0x15, 0x61,
	0x64, 0x64, 0x4d, 0x6f, 0x72, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x61, 0x74, 0x68, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x63, 0x63, 0x75, 0x70,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x09, 0x63, 0x75, 0x72, 0x5f, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x63, 0x75, 0x72, 0x48,
	0x6f, 0x75, 0x73, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x75, 0x72, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x31, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x63, 0x75, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69,
	0x6e, 0x65, 0x31, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x75, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x32, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x63, 0x75, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x6e, 0x65, 0x32, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x72, 0x5f, 0x70, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x72, 0x50, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x2d, 0x0a, 0x13, 0x69, 0x73, 0x5f, 0x73, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x69,
	0x73, 0x53, 0x61, 0x6d, 0x65, 0x50, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x20, 0x0a, 0x09, 0x70, 0x65, 0x72, 0x5f, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x70, 0x65, 0x72, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x2f, 0x0a, 0x11, 0x70, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x31, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0f,
	0x70, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x6e, 0x65, 0x31, 0x88,
	0x01, 0x01, 0x12, 0x2f, 0x0a, 0x11, 0x70, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x32, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52,
	0x0f, 0x70, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x6e, 0x65, 0x32,
	0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x69, 0x6e, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x50,
	0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x63, 0x75,
	0x72, 0x5f, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x65, 0x72, 0x5f,
	0x68, 0x6f, 0x75, 0x73, 0x65, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x31, 0x42, 0x14, 0x0a, 0x12, 0x5f,
	0x70, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x69, 0x6e, 0x65,
	0x32, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x69, 0x6e, 0x63, 0x6f, 0x64,
	0x65, 0x22, 0x32, 0x0a, 0x16, 0x61, 0x64, 0x64, 0x4d, 0x6f, 0x72, 0x65, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_customers_addmoredetails_proto_rawDescOnce sync.Once
	file_customers_addmoredetails_proto_rawDescData = file_customers_addmoredetails_proto_rawDesc
)

func file_customers_addmoredetails_proto_rawDescGZIP() []byte {
	file_customers_addmoredetails_proto_rawDescOnce.Do(func() {
		file_customers_addmoredetails_proto_rawDescData = protoimpl.X.CompressGZIP(file_customers_addmoredetails_proto_rawDescData)
	})
	return file_customers_addmoredetails_proto_rawDescData
}

var file_customers_addmoredetails_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_customers_addmoredetails_proto_goTypes = []interface{}{
	(*AddMoreDetailsRequest)(nil),  // 0: customers.addmoredetails.addMoreDetailsRequest
	(*AddMoreDetailsResponse)(nil), // 1: customers.addmoredetails.addMoreDetailsResponse
}
var file_customers_addmoredetails_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_customers_addmoredetails_proto_init() }
func file_customers_addmoredetails_proto_init() {
	if File_customers_addmoredetails_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_customers_addmoredetails_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddMoreDetailsRequest); i {
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
		file_customers_addmoredetails_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddMoreDetailsResponse); i {
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
	file_customers_addmoredetails_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_customers_addmoredetails_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_customers_addmoredetails_proto_goTypes,
		DependencyIndexes: file_customers_addmoredetails_proto_depIdxs,
		MessageInfos:      file_customers_addmoredetails_proto_msgTypes,
	}.Build()
	File_customers_addmoredetails_proto = out.File
	file_customers_addmoredetails_proto_rawDesc = nil
	file_customers_addmoredetails_proto_goTypes = nil
	file_customers_addmoredetails_proto_depIdxs = nil
}
