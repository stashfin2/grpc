// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.6
// source: kyc/getselfieurl.proto

package kyc

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

type GetSelfieRedirectionUrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RedirectUrl string `protobuf:"bytes,1,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
	CustomerId  int64  `protobuf:"varint,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *GetSelfieRedirectionUrlRequest) Reset() {
	*x = GetSelfieRedirectionUrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kyc_getselfieurl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSelfieRedirectionUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSelfieRedirectionUrlRequest) ProtoMessage() {}

func (x *GetSelfieRedirectionUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kyc_getselfieurl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSelfieRedirectionUrlRequest.ProtoReflect.Descriptor instead.
func (*GetSelfieRedirectionUrlRequest) Descriptor() ([]byte, []int) {
	return file_kyc_getselfieurl_proto_rawDescGZIP(), []int{0}
}

func (x *GetSelfieRedirectionUrlRequest) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

func (x *GetSelfieRedirectionUrlRequest) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

type GetSelfieRedirectionUrlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status                 string        `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	ProfilePicUploadUrl    string        `protobuf:"bytes,2,opt,name=profile_pic_upload_url,json=profilePicUploadUrl,proto3" json:"profile_pic_upload_url,omitempty"`
	LandingPage            string        `protobuf:"bytes,3,opt,name=landing_page,json=landingPage,proto3" json:"landing_page,omitempty"`
	RelationalStashfinSite string        `protobuf:"bytes,4,opt,name=relational_stashfin_site,json=relationalStashfinSite,proto3" json:"relational_stashfin_site,omitempty"`
	R                      *ShortUrlData `protobuf:"bytes,5,opt,name=r,proto3,oneof" json:"r,omitempty"`
}

func (x *GetSelfieRedirectionUrlResponse) Reset() {
	*x = GetSelfieRedirectionUrlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kyc_getselfieurl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSelfieRedirectionUrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSelfieRedirectionUrlResponse) ProtoMessage() {}

func (x *GetSelfieRedirectionUrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kyc_getselfieurl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSelfieRedirectionUrlResponse.ProtoReflect.Descriptor instead.
func (*GetSelfieRedirectionUrlResponse) Descriptor() ([]byte, []int) {
	return file_kyc_getselfieurl_proto_rawDescGZIP(), []int{1}
}

func (x *GetSelfieRedirectionUrlResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetSelfieRedirectionUrlResponse) GetProfilePicUploadUrl() string {
	if x != nil {
		return x.ProfilePicUploadUrl
	}
	return ""
}

func (x *GetSelfieRedirectionUrlResponse) GetLandingPage() string {
	if x != nil {
		return x.LandingPage
	}
	return ""
}

func (x *GetSelfieRedirectionUrlResponse) GetRelationalStashfinSite() string {
	if x != nil {
		return x.RelationalStashfinSite
	}
	return ""
}

func (x *GetSelfieRedirectionUrlResponse) GetR() *ShortUrlData {
	if x != nil {
		return x.R
	}
	return nil
}

type ShortUrlData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UrlCode     string  `protobuf:"bytes,1,opt,name=url_code,json=urlCode,proto3" json:"url_code,omitempty"`
	RedirectUrl string  `protobuf:"bytes,2,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
	IsActive    int32   `protobuf:"varint,3,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	UserId      int32   `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OfferType   *string `protobuf:"bytes,5,opt,name=offer_type,json=offerType,proto3,oneof" json:"offer_type,omitempty"`
	ExpireDate  string  `protobuf:"bytes,6,opt,name=expire_date,json=expireDate,proto3" json:"expire_date,omitempty"`
	CreateDate  string  `protobuf:"bytes,7,opt,name=create_date,json=createDate,proto3" json:"create_date,omitempty"`
	LandingPage *string `protobuf:"bytes,8,opt,name=landing_page,json=landingPage,proto3,oneof" json:"landing_page,omitempty"`
}

func (x *ShortUrlData) Reset() {
	*x = ShortUrlData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kyc_getselfieurl_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortUrlData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortUrlData) ProtoMessage() {}

func (x *ShortUrlData) ProtoReflect() protoreflect.Message {
	mi := &file_kyc_getselfieurl_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortUrlData.ProtoReflect.Descriptor instead.
func (*ShortUrlData) Descriptor() ([]byte, []int) {
	return file_kyc_getselfieurl_proto_rawDescGZIP(), []int{2}
}

func (x *ShortUrlData) GetUrlCode() string {
	if x != nil {
		return x.UrlCode
	}
	return ""
}

func (x *ShortUrlData) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

func (x *ShortUrlData) GetIsActive() int32 {
	if x != nil {
		return x.IsActive
	}
	return 0
}

func (x *ShortUrlData) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ShortUrlData) GetOfferType() string {
	if x != nil && x.OfferType != nil {
		return *x.OfferType
	}
	return ""
}

func (x *ShortUrlData) GetExpireDate() string {
	if x != nil {
		return x.ExpireDate
	}
	return ""
}

func (x *ShortUrlData) GetCreateDate() string {
	if x != nil {
		return x.CreateDate
	}
	return ""
}

func (x *ShortUrlData) GetLandingPage() string {
	if x != nil && x.LandingPage != nil {
		return *x.LandingPage
	}
	return ""
}

var File_kyc_getselfieurl_proto protoreflect.FileDescriptor

var file_kyc_getselfieurl_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6b, 0x79, 0x63, 0x2f, 0x67, 0x65, 0x74, 0x73, 0x65, 0x6c, 0x66, 0x69, 0x65, 0x75,
	0x72, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6b, 0x79, 0x63, 0x2e, 0x67, 0x65,
	0x74, 0x73, 0x65, 0x6c, 0x66, 0x69, 0x65, 0x75, 0x72, 0x6c, 0x22, 0x64, 0x0a, 0x1e, 0x47, 0x65,
	0x74, 0x53, 0x65, 0x6c, 0x66, 0x69, 0x65, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x84, 0x02, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6c, 0x66, 0x69, 0x65, 0x52, 0x65,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x33, 0x0a, 0x16,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x69, 0x63, 0x5f, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72,
	0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x50, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x18, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x73, 0x68, 0x66, 0x69, 0x6e, 0x5f, 0x73, 0x69, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x53, 0x74, 0x61, 0x73, 0x68, 0x66, 0x69, 0x6e, 0x53, 0x69, 0x74, 0x65, 0x12, 0x31,
	0x0a, 0x01, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6b, 0x79, 0x63, 0x2e,
	0x67, 0x65, 0x74, 0x73, 0x65, 0x6c, 0x66, 0x69, 0x65, 0x75, 0x72, 0x6c, 0x2e, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x55, 0x72, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x01, 0x72, 0x88, 0x01,
	0x01, 0x42, 0x04, 0x0a, 0x02, 0x5f, 0x72, 0x22, 0xb0, 0x02, 0x0a, 0x0c, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x55, 0x72, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x72, 0x6c, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x72, 0x6c, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0a,
	0x6f, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x09, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x6c, 0x61, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x50, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6f,
	0x66, 0x66, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6c, 0x61,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_kyc_getselfieurl_proto_rawDescOnce sync.Once
	file_kyc_getselfieurl_proto_rawDescData = file_kyc_getselfieurl_proto_rawDesc
)

func file_kyc_getselfieurl_proto_rawDescGZIP() []byte {
	file_kyc_getselfieurl_proto_rawDescOnce.Do(func() {
		file_kyc_getselfieurl_proto_rawDescData = protoimpl.X.CompressGZIP(file_kyc_getselfieurl_proto_rawDescData)
	})
	return file_kyc_getselfieurl_proto_rawDescData
}

var file_kyc_getselfieurl_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_kyc_getselfieurl_proto_goTypes = []interface{}{
	(*GetSelfieRedirectionUrlRequest)(nil),  // 0: kyc.getselfieurl.GetSelfieRedirectionUrlRequest
	(*GetSelfieRedirectionUrlResponse)(nil), // 1: kyc.getselfieurl.GetSelfieRedirectionUrlResponse
	(*ShortUrlData)(nil),                    // 2: kyc.getselfieurl.shortUrlData
}
var file_kyc_getselfieurl_proto_depIdxs = []int32{
	2, // 0: kyc.getselfieurl.GetSelfieRedirectionUrlResponse.r:type_name -> kyc.getselfieurl.shortUrlData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_kyc_getselfieurl_proto_init() }
func file_kyc_getselfieurl_proto_init() {
	if File_kyc_getselfieurl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kyc_getselfieurl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSelfieRedirectionUrlRequest); i {
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
		file_kyc_getselfieurl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSelfieRedirectionUrlResponse); i {
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
		file_kyc_getselfieurl_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortUrlData); i {
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
	file_kyc_getselfieurl_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_kyc_getselfieurl_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kyc_getselfieurl_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kyc_getselfieurl_proto_goTypes,
		DependencyIndexes: file_kyc_getselfieurl_proto_depIdxs,
		MessageInfos:      file_kyc_getselfieurl_proto_msgTypes,
	}.Build()
	File_kyc_getselfieurl_proto = out.File
	file_kyc_getselfieurl_proto_rawDesc = nil
	file_kyc_getselfieurl_proto_goTypes = nil
	file_kyc_getselfieurl_proto_depIdxs = nil
}
