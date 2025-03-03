// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: customers/cb/insurancedetails.proto

package cb

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

type Offer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	IconUrl     string `protobuf:"bytes,3,opt,name=icon_url,json=iconUrl,proto3" json:"icon_url,omitempty"`
}

func (x *Offer) Reset() {
	*x = Offer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_cb_insurancedetails_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Offer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Offer) ProtoMessage() {}

func (x *Offer) ProtoReflect() protoreflect.Message {
	mi := &file_customers_cb_insurancedetails_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Offer.ProtoReflect.Descriptor instead.
func (*Offer) Descriptor() ([]byte, []int) {
	return file_customers_cb_insurancedetails_proto_rawDescGZIP(), []int{0}
}

func (x *Offer) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Offer) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Offer) GetIconUrl() string {
	if x != nil {
		return x.IconUrl
	}
	return ""
}

type InsuranceDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InsuranceDetailsRequest) Reset() {
	*x = InsuranceDetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_cb_insurancedetails_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsuranceDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsuranceDetailsRequest) ProtoMessage() {}

func (x *InsuranceDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customers_cb_insurancedetails_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsuranceDetailsRequest.ProtoReflect.Descriptor instead.
func (*InsuranceDetailsRequest) Descriptor() ([]byte, []int) {
	return file_customers_cb_insurancedetails_proto_rawDescGZIP(), []int{1}
}

type InsuranceDetailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offers []*Offer `protobuf:"bytes,1,rep,name=offers,proto3" json:"offers,omitempty"`
}

func (x *InsuranceDetailsResponse) Reset() {
	*x = InsuranceDetailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_cb_insurancedetails_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsuranceDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsuranceDetailsResponse) ProtoMessage() {}

func (x *InsuranceDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customers_cb_insurancedetails_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsuranceDetailsResponse.ProtoReflect.Descriptor instead.
func (*InsuranceDetailsResponse) Descriptor() ([]byte, []int) {
	return file_customers_cb_insurancedetails_proto_rawDescGZIP(), []int{2}
}

func (x *InsuranceDetailsResponse) GetOffers() []*Offer {
	if x != nil {
		return x.Offers
	}
	return nil
}

var File_customers_cb_insurancedetails_proto protoreflect.FileDescriptor

var file_customers_cb_insurancedetails_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x63, 0x62, 0x2f, 0x69,
	0x6e, 0x73, 0x75, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73,
	0x2e, 0x63, 0x62, 0x2e, 0x69, 0x6e, 0x73, 0x75, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x22, 0x5a, 0x0a, 0x05, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x63, 0x6f, 0x6e, 0x55, 0x72, 0x6c,
	0x22, 0x19, 0x0a, 0x17, 0x69, 0x6e, 0x73, 0x75, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x58, 0x0a, 0x18, 0x69,
	0x6e, 0x73, 0x75, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x73, 0x2e, 0x63, 0x62, 0x2e, 0x69, 0x6e, 0x73, 0x75, 0x72, 0x61, 0x6e, 0x63, 0x65,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_customers_cb_insurancedetails_proto_rawDescOnce sync.Once
	file_customers_cb_insurancedetails_proto_rawDescData = file_customers_cb_insurancedetails_proto_rawDesc
)

func file_customers_cb_insurancedetails_proto_rawDescGZIP() []byte {
	file_customers_cb_insurancedetails_proto_rawDescOnce.Do(func() {
		file_customers_cb_insurancedetails_proto_rawDescData = protoimpl.X.CompressGZIP(file_customers_cb_insurancedetails_proto_rawDescData)
	})
	return file_customers_cb_insurancedetails_proto_rawDescData
}

var file_customers_cb_insurancedetails_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_customers_cb_insurancedetails_proto_goTypes = []interface{}{
	(*Offer)(nil),                    // 0: customers.cb.insurancedetails.Offer
	(*InsuranceDetailsRequest)(nil),  // 1: customers.cb.insurancedetails.insuranceDetailsRequest
	(*InsuranceDetailsResponse)(nil), // 2: customers.cb.insurancedetails.insuranceDetailsResponse
}
var file_customers_cb_insurancedetails_proto_depIdxs = []int32{
	0, // 0: customers.cb.insurancedetails.insuranceDetailsResponse.offers:type_name -> customers.cb.insurancedetails.Offer
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_customers_cb_insurancedetails_proto_init() }
func file_customers_cb_insurancedetails_proto_init() {
	if File_customers_cb_insurancedetails_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_customers_cb_insurancedetails_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Offer); i {
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
		file_customers_cb_insurancedetails_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsuranceDetailsRequest); i {
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
		file_customers_cb_insurancedetails_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsuranceDetailsResponse); i {
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
			RawDescriptor: file_customers_cb_insurancedetails_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_customers_cb_insurancedetails_proto_goTypes,
		DependencyIndexes: file_customers_cb_insurancedetails_proto_depIdxs,
		MessageInfos:      file_customers_cb_insurancedetails_proto_msgTypes,
	}.Build()
	File_customers_cb_insurancedetails_proto = out.File
	file_customers_cb_insurancedetails_proto_rawDesc = nil
	file_customers_cb_insurancedetails_proto_goTypes = nil
	file_customers_cb_insurancedetails_proto_depIdxs = nil
}
