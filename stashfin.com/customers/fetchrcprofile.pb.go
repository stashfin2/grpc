// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: customers/fetchrcprofile.proto

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

type FetchRCProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RcNumber             string `protobuf:"bytes,1,opt,name=rcNumber,proto3" json:"rcNumber,omitempty"`
	DrivingLicenseNumber string `protobuf:"bytes,2,opt,name=drivingLicenseNumber,proto3" json:"drivingLicenseNumber,omitempty"`
	DateOfBirth          string `protobuf:"bytes,3,opt,name=dateOfBirth,proto3" json:"dateOfBirth,omitempty"`
}

func (x *FetchRCProfileRequest) Reset() {
	*x = FetchRCProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_fetchrcprofile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchRCProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchRCProfileRequest) ProtoMessage() {}

func (x *FetchRCProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customers_fetchrcprofile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchRCProfileRequest.ProtoReflect.Descriptor instead.
func (*FetchRCProfileRequest) Descriptor() ([]byte, []int) {
	return file_customers_fetchrcprofile_proto_rawDescGZIP(), []int{0}
}

func (x *FetchRCProfileRequest) GetRcNumber() string {
	if x != nil {
		return x.RcNumber
	}
	return ""
}

func (x *FetchRCProfileRequest) GetDrivingLicenseNumber() string {
	if x != nil {
		return x.DrivingLicenseNumber
	}
	return ""
}

func (x *FetchRCProfileRequest) GetDateOfBirth() string {
	if x != nil {
		return x.DateOfBirth
	}
	return ""
}

type FetchRCProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId       int32  `protobuf:"varint,1,opt,name=customerId,proto3" json:"customerId,omitempty"`
	RequestId        string `protobuf:"bytes,2,opt,name=requestId,proto3" json:"requestId,omitempty"`
	RcNumber         string `protobuf:"bytes,3,opt,name=rcNumber,proto3" json:"rcNumber,omitempty"`
	Name             string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	PresentAddress   string `protobuf:"bytes,5,opt,name=presentAddress,proto3" json:"presentAddress,omitempty"`
	PermanentAddress string `protobuf:"bytes,6,opt,name=permanentAddress,proto3" json:"permanentAddress,omitempty"`
	FatherName       string `protobuf:"bytes,7,opt,name=fatherName,proto3" json:"fatherName,omitempty"`
	Financed         bool   `protobuf:"varint,8,opt,name=financed,proto3" json:"financed,omitempty"`
	Financier        string `protobuf:"bytes,9,opt,name=financier,proto3" json:"financier,omitempty"`
}

func (x *FetchRCProfileResponse) Reset() {
	*x = FetchRCProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_fetchrcprofile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchRCProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchRCProfileResponse) ProtoMessage() {}

func (x *FetchRCProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customers_fetchrcprofile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchRCProfileResponse.ProtoReflect.Descriptor instead.
func (*FetchRCProfileResponse) Descriptor() ([]byte, []int) {
	return file_customers_fetchrcprofile_proto_rawDescGZIP(), []int{1}
}

func (x *FetchRCProfileResponse) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *FetchRCProfileResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *FetchRCProfileResponse) GetRcNumber() string {
	if x != nil {
		return x.RcNumber
	}
	return ""
}

func (x *FetchRCProfileResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FetchRCProfileResponse) GetPresentAddress() string {
	if x != nil {
		return x.PresentAddress
	}
	return ""
}

func (x *FetchRCProfileResponse) GetPermanentAddress() string {
	if x != nil {
		return x.PermanentAddress
	}
	return ""
}

func (x *FetchRCProfileResponse) GetFatherName() string {
	if x != nil {
		return x.FatherName
	}
	return ""
}

func (x *FetchRCProfileResponse) GetFinanced() bool {
	if x != nil {
		return x.Financed
	}
	return false
}

func (x *FetchRCProfileResponse) GetFinancier() string {
	if x != nil {
		return x.Financier
	}
	return ""
}

var File_customers_fetchrcprofile_proto protoreflect.FileDescriptor

var file_customers_fetchrcprofile_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x66, 0x65, 0x74, 0x63,
	0x68, 0x72, 0x63, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x18, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x66, 0x65, 0x74, 0x63,
	0x68, 0x72, 0x63, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x89, 0x01, 0x0a, 0x15, 0x66,
	0x65, 0x74, 0x63, 0x68, 0x52, 0x43, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x63, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x63, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x32, 0x0a, 0x14, 0x64, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x63, 0x65, 0x6e,
	0x73, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14,
	0x64, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69,
	0x72, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x4f,
	0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x22, 0xb4, 0x02, 0x0a, 0x16, 0x66, 0x65, 0x74, 0x63, 0x68,
	0x52, 0x43, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x63, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x63, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x26, 0x0a, 0x0e, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x70, 0x65, 0x72, 0x6d, 0x61,
	0x6e, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x70, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x61, 0x74, 0x68, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x61, 0x74, 0x68, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x69, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x69, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_customers_fetchrcprofile_proto_rawDescOnce sync.Once
	file_customers_fetchrcprofile_proto_rawDescData = file_customers_fetchrcprofile_proto_rawDesc
)

func file_customers_fetchrcprofile_proto_rawDescGZIP() []byte {
	file_customers_fetchrcprofile_proto_rawDescOnce.Do(func() {
		file_customers_fetchrcprofile_proto_rawDescData = protoimpl.X.CompressGZIP(file_customers_fetchrcprofile_proto_rawDescData)
	})
	return file_customers_fetchrcprofile_proto_rawDescData
}

var file_customers_fetchrcprofile_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_customers_fetchrcprofile_proto_goTypes = []interface{}{
	(*FetchRCProfileRequest)(nil),  // 0: customers.fetchrcprofile.fetchRCProfileRequest
	(*FetchRCProfileResponse)(nil), // 1: customers.fetchrcprofile.fetchRCProfileResponse
}
var file_customers_fetchrcprofile_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_customers_fetchrcprofile_proto_init() }
func file_customers_fetchrcprofile_proto_init() {
	if File_customers_fetchrcprofile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_customers_fetchrcprofile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchRCProfileRequest); i {
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
		file_customers_fetchrcprofile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchRCProfileResponse); i {
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
			RawDescriptor: file_customers_fetchrcprofile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_customers_fetchrcprofile_proto_goTypes,
		DependencyIndexes: file_customers_fetchrcprofile_proto_depIdxs,
		MessageInfos:      file_customers_fetchrcprofile_proto_msgTypes,
	}.Build()
	File_customers_fetchrcprofile_proto = out.File
	file_customers_fetchrcprofile_proto_rawDesc = nil
	file_customers_fetchrcprofile_proto_goTypes = nil
	file_customers_fetchrcprofile_proto_depIdxs = nil
}
