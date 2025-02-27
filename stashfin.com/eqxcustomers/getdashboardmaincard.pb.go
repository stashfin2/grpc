// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: eqxcustomers/getdashboardmaincard.proto

package eqxcustomers

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

type GetDashboardMainCardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDashboardMainCardRequest) Reset() {
	*x = GetDashboardMainCardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eqxcustomers_getdashboardmaincard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDashboardMainCardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDashboardMainCardRequest) ProtoMessage() {}

func (x *GetDashboardMainCardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eqxcustomers_getdashboardmaincard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDashboardMainCardRequest.ProtoReflect.Descriptor instead.
func (*GetDashboardMainCardRequest) Descriptor() ([]byte, []int) {
	return file_eqxcustomers_getdashboardmaincard_proto_rawDescGZIP(), []int{0}
}

type GetDashboardMainCardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Code      string        `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	BlockData *MainCardData `protobuf:"bytes,3,opt,name=block_data,json=blockData,proto3" json:"block_data,omitempty"`
	UserState string        `protobuf:"bytes,4,opt,name=user_state,json=userState,proto3" json:"user_state,omitempty"`
}

func (x *GetDashboardMainCardResponse) Reset() {
	*x = GetDashboardMainCardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eqxcustomers_getdashboardmaincard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDashboardMainCardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDashboardMainCardResponse) ProtoMessage() {}

func (x *GetDashboardMainCardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eqxcustomers_getdashboardmaincard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDashboardMainCardResponse.ProtoReflect.Descriptor instead.
func (*GetDashboardMainCardResponse) Descriptor() ([]byte, []int) {
	return file_eqxcustomers_getdashboardmaincard_proto_rawDescGZIP(), []int{1}
}

func (x *GetDashboardMainCardResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetDashboardMainCardResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GetDashboardMainCardResponse) GetBlockData() *MainCardData {
	if x != nil {
		return x.BlockData
	}
	return nil
}

func (x *GetDashboardMainCardResponse) GetUserState() string {
	if x != nil {
		return x.UserState
	}
	return ""
}

type MainCardData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text             string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	SubText          string `protobuf:"bytes,2,opt,name=sub_text,json=subText,proto3" json:"sub_text,omitempty"`
	ActionButtonText string `protobuf:"bytes,3,opt,name=action_button_text,json=actionButtonText,proto3" json:"action_button_text,omitempty"`
	Action           string `protobuf:"bytes,4,opt,name=action,proto3" json:"action,omitempty"`
	ActionType       string `protobuf:"bytes,5,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty"`
	Comment          string `protobuf:"bytes,6,opt,name=comment,proto3" json:"comment,omitempty"`
	LandingPage      string `protobuf:"bytes,7,opt,name=landing_page,json=landingPage,proto3" json:"landing_page,omitempty"`
	Timer            string `protobuf:"bytes,8,opt,name=timer,proto3" json:"timer,omitempty"`
	Image            string `protobuf:"bytes,9,opt,name=image,proto3" json:"image,omitempty"`
	ShowButton       bool   `protobuf:"varint,10,opt,name=show_button,json=showButton,proto3" json:"show_button,omitempty"`
	StaticText       string `protobuf:"bytes,11,opt,name=static_text,json=staticText,proto3" json:"static_text,omitempty"`
}

func (x *MainCardData) Reset() {
	*x = MainCardData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eqxcustomers_getdashboardmaincard_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MainCardData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MainCardData) ProtoMessage() {}

func (x *MainCardData) ProtoReflect() protoreflect.Message {
	mi := &file_eqxcustomers_getdashboardmaincard_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MainCardData.ProtoReflect.Descriptor instead.
func (*MainCardData) Descriptor() ([]byte, []int) {
	return file_eqxcustomers_getdashboardmaincard_proto_rawDescGZIP(), []int{2}
}

func (x *MainCardData) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *MainCardData) GetSubText() string {
	if x != nil {
		return x.SubText
	}
	return ""
}

func (x *MainCardData) GetActionButtonText() string {
	if x != nil {
		return x.ActionButtonText
	}
	return ""
}

func (x *MainCardData) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *MainCardData) GetActionType() string {
	if x != nil {
		return x.ActionType
	}
	return ""
}

func (x *MainCardData) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *MainCardData) GetLandingPage() string {
	if x != nil {
		return x.LandingPage
	}
	return ""
}

func (x *MainCardData) GetTimer() string {
	if x != nil {
		return x.Timer
	}
	return ""
}

func (x *MainCardData) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *MainCardData) GetShowButton() bool {
	if x != nil {
		return x.ShowButton
	}
	return false
}

func (x *MainCardData) GetStaticText() string {
	if x != nil {
		return x.StaticText
	}
	return ""
}

var File_eqxcustomers_getdashboardmaincard_proto protoreflect.FileDescriptor

var file_eqxcustomers_getdashboardmaincard_proto_rawDesc = []byte{
	0x0a, 0x27, 0x65, 0x71, 0x78, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x67,
	0x65, 0x74, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x6d, 0x61, 0x69, 0x6e, 0x63,
	0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x65, 0x71, 0x78, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x67, 0x65, 0x74, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x6d, 0x61, 0x69, 0x6e, 0x63, 0x61, 0x72, 0x64, 0x22, 0x1d, 0x0a, 0x1b,
	0x67, 0x65, 0x74, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x4d, 0x61, 0x69, 0x6e,
	0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xb5, 0x01, 0x0a, 0x1c,
	0x67, 0x65, 0x74, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x4d, 0x61, 0x69, 0x6e,
	0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x4e, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x65, 0x71, 0x78, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x67, 0x65, 0x74, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x6d, 0x61, 0x69, 0x6e, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x4d, 0x61, 0x69,
	0x6e, 0x43, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x22, 0xcf, 0x02, 0x0a, 0x0c, 0x4d, 0x61, 0x69, 0x6e, 0x43, 0x61, 0x72, 0x64,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x5f,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x54,
	0x65, 0x78, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x75,
	0x74, 0x74, 0x6f, 0x6e, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x54, 0x65, 0x78,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x61, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x50, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x72,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x62, 0x75, 0x74, 0x74,
	0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x68, 0x6f, 0x77, 0x42, 0x75,
	0x74, 0x74, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x63, 0x54, 0x65, 0x78, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eqxcustomers_getdashboardmaincard_proto_rawDescOnce sync.Once
	file_eqxcustomers_getdashboardmaincard_proto_rawDescData = file_eqxcustomers_getdashboardmaincard_proto_rawDesc
)

func file_eqxcustomers_getdashboardmaincard_proto_rawDescGZIP() []byte {
	file_eqxcustomers_getdashboardmaincard_proto_rawDescOnce.Do(func() {
		file_eqxcustomers_getdashboardmaincard_proto_rawDescData = protoimpl.X.CompressGZIP(file_eqxcustomers_getdashboardmaincard_proto_rawDescData)
	})
	return file_eqxcustomers_getdashboardmaincard_proto_rawDescData
}

var file_eqxcustomers_getdashboardmaincard_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_eqxcustomers_getdashboardmaincard_proto_goTypes = []interface{}{
	(*GetDashboardMainCardRequest)(nil),  // 0: eqxcustomers.getdashboardmaincard.getDashboardMainCardRequest
	(*GetDashboardMainCardResponse)(nil), // 1: eqxcustomers.getdashboardmaincard.getDashboardMainCardResponse
	(*MainCardData)(nil),                 // 2: eqxcustomers.getdashboardmaincard.MainCardData
}
var file_eqxcustomers_getdashboardmaincard_proto_depIdxs = []int32{
	2, // 0: eqxcustomers.getdashboardmaincard.getDashboardMainCardResponse.block_data:type_name -> eqxcustomers.getdashboardmaincard.MainCardData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eqxcustomers_getdashboardmaincard_proto_init() }
func file_eqxcustomers_getdashboardmaincard_proto_init() {
	if File_eqxcustomers_getdashboardmaincard_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eqxcustomers_getdashboardmaincard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDashboardMainCardRequest); i {
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
		file_eqxcustomers_getdashboardmaincard_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDashboardMainCardResponse); i {
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
		file_eqxcustomers_getdashboardmaincard_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MainCardData); i {
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
			RawDescriptor: file_eqxcustomers_getdashboardmaincard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eqxcustomers_getdashboardmaincard_proto_goTypes,
		DependencyIndexes: file_eqxcustomers_getdashboardmaincard_proto_depIdxs,
		MessageInfos:      file_eqxcustomers_getdashboardmaincard_proto_msgTypes,
	}.Build()
	File_eqxcustomers_getdashboardmaincard_proto = out.File
	file_eqxcustomers_getdashboardmaincard_proto_rawDesc = nil
	file_eqxcustomers_getdashboardmaincard_proto_goTypes = nil
	file_eqxcustomers_getdashboardmaincard_proto_depIdxs = nil
}
