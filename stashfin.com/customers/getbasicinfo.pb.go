// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.19.6
// source: customers/getbasicinfo.proto

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

type GetBasicInfoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBasicInfoRequest) Reset() {
	*x = GetBasicInfoRequest{}
	mi := &file_customers_getbasicinfo_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBasicInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBasicInfoRequest) ProtoMessage() {}

func (x *GetBasicInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customers_getbasicinfo_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBasicInfoRequest.ProtoReflect.Descriptor instead.
func (*GetBasicInfoRequest) Descriptor() ([]byte, []int) {
	return file_customers_getbasicinfo_proto_rawDescGZIP(), []int{0}
}

type GetBasicInfoResponse struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Mobile           string                 `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Pan              string                 `protobuf:"bytes,2,opt,name=pan,proto3" json:"pan,omitempty"`
	Aadhaar          string                 `protobuf:"bytes,3,opt,name=aadhaar,proto3" json:"aadhaar,omitempty"`
	PersonalEmail    string                 `protobuf:"bytes,4,opt,name=personal_email,json=personalEmail,proto3" json:"personal_email,omitempty"`
	CurrAddress      *Address               `protobuf:"bytes,5,opt,name=curr_address,json=currAddress,proto3" json:"curr_address,omitempty"`
	PermanentAddress *Address               `protobuf:"bytes,6,opt,name=permanent_address,json=permanentAddress,proto3" json:"permanent_address,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *GetBasicInfoResponse) Reset() {
	*x = GetBasicInfoResponse{}
	mi := &file_customers_getbasicinfo_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBasicInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBasicInfoResponse) ProtoMessage() {}

func (x *GetBasicInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customers_getbasicinfo_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBasicInfoResponse.ProtoReflect.Descriptor instead.
func (*GetBasicInfoResponse) Descriptor() ([]byte, []int) {
	return file_customers_getbasicinfo_proto_rawDescGZIP(), []int{1}
}

func (x *GetBasicInfoResponse) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *GetBasicInfoResponse) GetPan() string {
	if x != nil {
		return x.Pan
	}
	return ""
}

func (x *GetBasicInfoResponse) GetAadhaar() string {
	if x != nil {
		return x.Aadhaar
	}
	return ""
}

func (x *GetBasicInfoResponse) GetPersonalEmail() string {
	if x != nil {
		return x.PersonalEmail
	}
	return ""
}

func (x *GetBasicInfoResponse) GetCurrAddress() *Address {
	if x != nil {
		return x.CurrAddress
	}
	return nil
}

func (x *GetBasicInfoResponse) GetPermanentAddress() *Address {
	if x != nil {
		return x.PermanentAddress
	}
	return nil
}

type Address struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	HouseFlatNo   string                 `protobuf:"bytes,1,opt,name=house_flat_no,json=houseFlatNo,proto3" json:"house_flat_no,omitempty"`
	AddressLine_1 string                 `protobuf:"bytes,2,opt,name=address_line_1,json=addressLine1,proto3" json:"address_line_1,omitempty"`
	AddressLine_2 string                 `protobuf:"bytes,3,opt,name=address_line_2,json=addressLine2,proto3" json:"address_line_2,omitempty"`
	Landmark      string                 `protobuf:"bytes,4,opt,name=landmark,proto3" json:"landmark,omitempty"`
	City          string                 `protobuf:"bytes,5,opt,name=city,proto3" json:"city,omitempty"`
	State         string                 `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty"`
	PinCode       string                 `protobuf:"bytes,7,opt,name=pin_code,json=pinCode,proto3" json:"pin_code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Address) Reset() {
	*x = Address{}
	mi := &file_customers_getbasicinfo_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_customers_getbasicinfo_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_customers_getbasicinfo_proto_rawDescGZIP(), []int{2}
}

func (x *Address) GetHouseFlatNo() string {
	if x != nil {
		return x.HouseFlatNo
	}
	return ""
}

func (x *Address) GetAddressLine_1() string {
	if x != nil {
		return x.AddressLine_1
	}
	return ""
}

func (x *Address) GetAddressLine_2() string {
	if x != nil {
		return x.AddressLine_2
	}
	return ""
}

func (x *Address) GetLandmark() string {
	if x != nil {
		return x.Landmark
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Address) GetPinCode() string {
	if x != nil {
		return x.PinCode
	}
	return ""
}

var File_customers_getbasicinfo_proto protoreflect.FileDescriptor

var file_customers_getbasicinfo_proto_rawDesc = string([]byte{
	0x0a, 0x1c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x62,
	0x61, 0x73, 0x69, 0x63, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x67, 0x65, 0x74, 0x62, 0x61, 0x73,
	0x69, 0x63, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x15, 0x0a, 0x13, 0x67, 0x65, 0x74, 0x42, 0x61, 0x73,
	0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x93, 0x02,
	0x0a, 0x14, 0x67, 0x65, 0x74, 0x42, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x70, 0x61, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x61, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x61, 0x64, 0x68, 0x61, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x61, 0x64, 0x68, 0x61, 0x61, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x42, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x73, 0x2e, 0x67, 0x65, 0x74, 0x62, 0x61, 0x73, 0x69, 0x63, 0x69, 0x6e, 0x66, 0x6f,
	0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x4c, 0x0a, 0x11, 0x70, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x65,
	0x6e, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x67, 0x65, 0x74,
	0x62, 0x61, 0x73, 0x69, 0x63, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x10, 0x70, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0xda, 0x01, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x22, 0x0a, 0x0d, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x66, 0x6c, 0x61, 0x74, 0x5f, 0x6e, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x46, 0x6c, 0x61,
	0x74, 0x4e, 0x6f, 0x12, 0x24, 0x0a, 0x0e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x6c,
	0x69, 0x6e, 0x65, 0x5f, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x6e, 0x65, 0x31, 0x12, 0x24, 0x0a, 0x0e, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x6e, 0x65, 0x32, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x64, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x64, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x69, 0x6e, 0x43, 0x6f, 0x64, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_customers_getbasicinfo_proto_rawDescOnce sync.Once
	file_customers_getbasicinfo_proto_rawDescData []byte
)

func file_customers_getbasicinfo_proto_rawDescGZIP() []byte {
	file_customers_getbasicinfo_proto_rawDescOnce.Do(func() {
		file_customers_getbasicinfo_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_customers_getbasicinfo_proto_rawDesc), len(file_customers_getbasicinfo_proto_rawDesc)))
	})
	return file_customers_getbasicinfo_proto_rawDescData
}

var file_customers_getbasicinfo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_customers_getbasicinfo_proto_goTypes = []any{
	(*GetBasicInfoRequest)(nil),  // 0: customers.getbasicinfo.getBasicInfoRequest
	(*GetBasicInfoResponse)(nil), // 1: customers.getbasicinfo.getBasicInfoResponse
	(*Address)(nil),              // 2: customers.getbasicinfo.Address
}
var file_customers_getbasicinfo_proto_depIdxs = []int32{
	2, // 0: customers.getbasicinfo.getBasicInfoResponse.curr_address:type_name -> customers.getbasicinfo.Address
	2, // 1: customers.getbasicinfo.getBasicInfoResponse.permanent_address:type_name -> customers.getbasicinfo.Address
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_customers_getbasicinfo_proto_init() }
func file_customers_getbasicinfo_proto_init() {
	if File_customers_getbasicinfo_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_customers_getbasicinfo_proto_rawDesc), len(file_customers_getbasicinfo_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_customers_getbasicinfo_proto_goTypes,
		DependencyIndexes: file_customers_getbasicinfo_proto_depIdxs,
		MessageInfos:      file_customers_getbasicinfo_proto_msgTypes,
	}.Build()
	File_customers_getbasicinfo_proto = out.File
	file_customers_getbasicinfo_proto_goTypes = nil
	file_customers_getbasicinfo_proto_depIdxs = nil
}
