// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.6
// source: bureau/getbasicdetails.proto

package bureau

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

type DetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId int64  `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Range      int32  `protobuf:"varint,2,opt,name=range,proto3" json:"range,omitempty"`
	Source     string `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	BureauType int32  `protobuf:"varint,4,opt,name=bureau_type,json=bureauType,proto3" json:"bureau_type,omitempty"`
	PullType   int32  `protobuf:"varint,5,opt,name=pull_type,json=pullType,proto3" json:"pull_type,omitempty"`
}

func (x *DetailsRequest) Reset() {
	*x = DetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bureau_getbasicdetails_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailsRequest) ProtoMessage() {}

func (x *DetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bureau_getbasicdetails_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailsRequest.ProtoReflect.Descriptor instead.
func (*DetailsRequest) Descriptor() ([]byte, []int) {
	return file_bureau_getbasicdetails_proto_rawDescGZIP(), []int{0}
}

func (x *DetailsRequest) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *DetailsRequest) GetRange() int32 {
	if x != nil {
		return x.Range
	}
	return 0
}

func (x *DetailsRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *DetailsRequest) GetBureauType() int32 {
	if x != nil {
		return x.BureauType
	}
	return 0
}

func (x *DetailsRequest) GetPullType() int32 {
	if x != nil {
		return x.PullType
	}
	return 0
}

type DetailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Data   *Report `protobuf:"bytes,2,opt,name=data,proto3,oneof" json:"data,omitempty"`
}

func (x *DetailsResponse) Reset() {
	*x = DetailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bureau_getbasicdetails_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailsResponse) ProtoMessage() {}

func (x *DetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bureau_getbasicdetails_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailsResponse.ProtoReflect.Descriptor instead.
func (*DetailsResponse) Descriptor() ([]byte, []int) {
	return file_bureau_getbasicdetails_proto_rawDescGZIP(), []int{1}
}

func (x *DetailsResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *DetailsResponse) GetData() *Report {
	if x != nil {
		return x.Data
	}
	return nil
}

type Report struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Score      int32  `protobuf:"varint,1,opt,name=score,proto3" json:"score,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Occupation string `protobuf:"bytes,3,opt,name=occupation,proto3" json:"occupation,omitempty"`
	Dob        string `protobuf:"bytes,4,opt,name=dob,proto3" json:"dob,omitempty"`
	Income     int32  `protobuf:"varint,5,opt,name=income,proto3" json:"income,omitempty"`
	PullDate   string `protobuf:"bytes,6,opt,name=pull_date,json=pullDate,proto3" json:"pull_date,omitempty"`
	BureauType string `protobuf:"bytes,7,opt,name=bureau_type,json=bureauType,proto3" json:"bureau_type,omitempty"`
	Gender     string `protobuf:"bytes,8,opt,name=gender,proto3" json:"gender,omitempty"`
}

func (x *Report) Reset() {
	*x = Report{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bureau_getbasicdetails_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Report) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Report) ProtoMessage() {}

func (x *Report) ProtoReflect() protoreflect.Message {
	mi := &file_bureau_getbasicdetails_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Report.ProtoReflect.Descriptor instead.
func (*Report) Descriptor() ([]byte, []int) {
	return file_bureau_getbasicdetails_proto_rawDescGZIP(), []int{2}
}

func (x *Report) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Report) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Report) GetOccupation() string {
	if x != nil {
		return x.Occupation
	}
	return ""
}

func (x *Report) GetDob() string {
	if x != nil {
		return x.Dob
	}
	return ""
}

func (x *Report) GetIncome() int32 {
	if x != nil {
		return x.Income
	}
	return 0
}

func (x *Report) GetPullDate() string {
	if x != nil {
		return x.PullDate
	}
	return ""
}

func (x *Report) GetBureauType() string {
	if x != nil {
		return x.BureauType
	}
	return ""
}

func (x *Report) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

var File_bureau_getbasicdetails_proto protoreflect.FileDescriptor

var file_bureau_getbasicdetails_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x62, 0x75, 0x72, 0x65, 0x61, 0x75, 0x2f, 0x67, 0x65, 0x74, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16,
	0x62, 0x75, 0x72, 0x65, 0x61, 0x75, 0x2e, 0x67, 0x65, 0x74, 0x62, 0x61, 0x73, 0x69, 0x63, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x9d, 0x01, 0x0a, 0x0e, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x61,
	0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x72, 0x65,
	0x61, 0x75, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x62,
	0x75, 0x72, 0x65, 0x61, 0x75, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x75, 0x6c,
	0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x75,
	0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x22, 0x6b, 0x0a, 0x0f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x37, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x62, 0x75, 0x72, 0x65, 0x61, 0x75, 0x2e, 0x67, 0x65, 0x74, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x48,
	0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x22, 0xd2, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x63, 0x63, 0x75,
	0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x63,
	0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x6f, 0x62, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x6f, 0x62, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6e,
	0x63, 0x6f, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x69, 0x6e, 0x63, 0x6f,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x75, 0x6c, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x72, 0x65, 0x61, 0x75, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x72, 0x65, 0x61, 0x75, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bureau_getbasicdetails_proto_rawDescOnce sync.Once
	file_bureau_getbasicdetails_proto_rawDescData = file_bureau_getbasicdetails_proto_rawDesc
)

func file_bureau_getbasicdetails_proto_rawDescGZIP() []byte {
	file_bureau_getbasicdetails_proto_rawDescOnce.Do(func() {
		file_bureau_getbasicdetails_proto_rawDescData = protoimpl.X.CompressGZIP(file_bureau_getbasicdetails_proto_rawDescData)
	})
	return file_bureau_getbasicdetails_proto_rawDescData
}

var file_bureau_getbasicdetails_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_bureau_getbasicdetails_proto_goTypes = []interface{}{
	(*DetailsRequest)(nil),  // 0: bureau.getbasicdetails.detailsRequest
	(*DetailsResponse)(nil), // 1: bureau.getbasicdetails.detailsResponse
	(*Report)(nil),          // 2: bureau.getbasicdetails.Report
}
var file_bureau_getbasicdetails_proto_depIdxs = []int32{
	2, // 0: bureau.getbasicdetails.detailsResponse.data:type_name -> bureau.getbasicdetails.Report
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_bureau_getbasicdetails_proto_init() }
func file_bureau_getbasicdetails_proto_init() {
	if File_bureau_getbasicdetails_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bureau_getbasicdetails_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailsRequest); i {
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
		file_bureau_getbasicdetails_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailsResponse); i {
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
		file_bureau_getbasicdetails_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Report); i {
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
	file_bureau_getbasicdetails_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bureau_getbasicdetails_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bureau_getbasicdetails_proto_goTypes,
		DependencyIndexes: file_bureau_getbasicdetails_proto_depIdxs,
		MessageInfos:      file_bureau_getbasicdetails_proto_msgTypes,
	}.Build()
	File_bureau_getbasicdetails_proto = out.File
	file_bureau_getbasicdetails_proto_rawDesc = nil
	file_bureau_getbasicdetails_proto_goTypes = nil
	file_bureau_getbasicdetails_proto_depIdxs = nil
}
