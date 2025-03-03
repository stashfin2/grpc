// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: customers/pd/getbanklist.proto

package pd

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

type GetPdBankListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPdBankListRequest) Reset() {
	*x = GetPdBankListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_pd_getbanklist_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPdBankListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPdBankListRequest) ProtoMessage() {}

func (x *GetPdBankListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customers_pd_getbanklist_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPdBankListRequest.ProtoReflect.Descriptor instead.
func (*GetPdBankListRequest) Descriptor() ([]byte, []int) {
	return file_customers_pd_getbanklist_proto_rawDescGZIP(), []int{0}
}

type GetPdBankListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PopularBanks []*BankList `protobuf:"bytes,1,rep,name=popular_banks,json=popularBanks,proto3" json:"popular_banks,omitempty"`
	AllBanks     []*BankList `protobuf:"bytes,2,rep,name=all_banks,json=allBanks,proto3" json:"all_banks,omitempty"`
}

func (x *GetPdBankListResponse) Reset() {
	*x = GetPdBankListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_pd_getbanklist_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPdBankListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPdBankListResponse) ProtoMessage() {}

func (x *GetPdBankListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customers_pd_getbanklist_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPdBankListResponse.ProtoReflect.Descriptor instead.
func (*GetPdBankListResponse) Descriptor() ([]byte, []int) {
	return file_customers_pd_getbanklist_proto_rawDescGZIP(), []int{1}
}

func (x *GetPdBankListResponse) GetPopularBanks() []*BankList {
	if x != nil {
		return x.PopularBanks
	}
	return nil
}

func (x *GetPdBankListResponse) GetAllBanks() []*BankList {
	if x != nil {
		return x.AllBanks
	}
	return nil
}

type BankList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BankName *string `protobuf:"bytes,1,opt,name=bank_name,json=bankName,proto3,oneof" json:"bank_name,omitempty"`
	BankCode *string `protobuf:"bytes,2,opt,name=bank_code,json=bankCode,proto3,oneof" json:"bank_code,omitempty"`
	IconUrl  *string `protobuf:"bytes,3,opt,name=icon_url,json=iconUrl,proto3,oneof" json:"icon_url,omitempty"`
}

func (x *BankList) Reset() {
	*x = BankList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_pd_getbanklist_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankList) ProtoMessage() {}

func (x *BankList) ProtoReflect() protoreflect.Message {
	mi := &file_customers_pd_getbanklist_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankList.ProtoReflect.Descriptor instead.
func (*BankList) Descriptor() ([]byte, []int) {
	return file_customers_pd_getbanklist_proto_rawDescGZIP(), []int{2}
}

func (x *BankList) GetBankName() string {
	if x != nil && x.BankName != nil {
		return *x.BankName
	}
	return ""
}

func (x *BankList) GetBankCode() string {
	if x != nil && x.BankCode != nil {
		return *x.BankCode
	}
	return ""
}

func (x *BankList) GetIconUrl() string {
	if x != nil && x.IconUrl != nil {
		return *x.IconUrl
	}
	return ""
}

var File_customers_pd_getbanklist_proto protoreflect.FileDescriptor

var file_customers_pd_getbanklist_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x64, 0x2f, 0x67,
	0x65, 0x74, 0x62, 0x61, 0x6e, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x18, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x64, 0x2e, 0x67,
	0x65, 0x74, 0x62, 0x61, 0x6e, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x16, 0x0a, 0x14, 0x67, 0x65,
	0x74, 0x50, 0x64, 0x42, 0x61, 0x6e, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0xa1, 0x01, 0x0a, 0x15, 0x67, 0x65, 0x74, 0x50, 0x64, 0x42, 0x61, 0x6e, 0x6b,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0d,
	0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x62, 0x61, 0x6e, 0x6b, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x64, 0x2e, 0x67, 0x65, 0x74, 0x62, 0x61, 0x6e, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x42,
	0x61, 0x6e, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0c, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72,
	0x42, 0x61, 0x6e, 0x6b, 0x73, 0x12, 0x3f, 0x0a, 0x09, 0x61, 0x6c, 0x6c, 0x5f, 0x62, 0x61, 0x6e,
	0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x64, 0x2e, 0x67, 0x65, 0x74, 0x62, 0x61, 0x6e, 0x6b, 0x6c,
	0x69, 0x73, 0x74, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x08, 0x61, 0x6c,
	0x6c, 0x42, 0x61, 0x6e, 0x6b, 0x73, 0x22, 0x97, 0x01, 0x0a, 0x08, 0x42, 0x61, 0x6e, 0x6b, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x09, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x62, 0x61, 0x6e, 0x6b, 0x4e, 0x61,
	0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x62, 0x61, 0x6e, 0x6b,
	0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a, 0x08, 0x69, 0x63, 0x6f, 0x6e, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x07, 0x69, 0x63, 0x6f,
	0x6e, 0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x62, 0x61, 0x6e, 0x6b,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_customers_pd_getbanklist_proto_rawDescOnce sync.Once
	file_customers_pd_getbanklist_proto_rawDescData = file_customers_pd_getbanklist_proto_rawDesc
)

func file_customers_pd_getbanklist_proto_rawDescGZIP() []byte {
	file_customers_pd_getbanklist_proto_rawDescOnce.Do(func() {
		file_customers_pd_getbanklist_proto_rawDescData = protoimpl.X.CompressGZIP(file_customers_pd_getbanklist_proto_rawDescData)
	})
	return file_customers_pd_getbanklist_proto_rawDescData
}

var file_customers_pd_getbanklist_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_customers_pd_getbanklist_proto_goTypes = []interface{}{
	(*GetPdBankListRequest)(nil),  // 0: customers.pd.getbanklist.getPdBankListRequest
	(*GetPdBankListResponse)(nil), // 1: customers.pd.getbanklist.getPdBankListResponse
	(*BankList)(nil),              // 2: customers.pd.getbanklist.BankList
}
var file_customers_pd_getbanklist_proto_depIdxs = []int32{
	2, // 0: customers.pd.getbanklist.getPdBankListResponse.popular_banks:type_name -> customers.pd.getbanklist.BankList
	2, // 1: customers.pd.getbanklist.getPdBankListResponse.all_banks:type_name -> customers.pd.getbanklist.BankList
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_customers_pd_getbanklist_proto_init() }
func file_customers_pd_getbanklist_proto_init() {
	if File_customers_pd_getbanklist_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_customers_pd_getbanklist_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPdBankListRequest); i {
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
		file_customers_pd_getbanklist_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPdBankListResponse); i {
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
		file_customers_pd_getbanklist_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankList); i {
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
	file_customers_pd_getbanklist_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_customers_pd_getbanklist_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_customers_pd_getbanklist_proto_goTypes,
		DependencyIndexes: file_customers_pd_getbanklist_proto_depIdxs,
		MessageInfos:      file_customers_pd_getbanklist_proto_msgTypes,
	}.Build()
	File_customers_pd_getbanklist_proto = out.File
	file_customers_pd_getbanklist_proto_rawDesc = nil
	file_customers_pd_getbanklist_proto_goTypes = nil
	file_customers_pd_getbanklist_proto_depIdxs = nil
}
