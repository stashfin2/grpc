// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: eqxcustomers/getdashboard.proto

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

type GetDashboardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductPage string `protobuf:"bytes,1,opt,name=product_page,json=productPage,proto3" json:"product_page,omitempty"`
}

func (x *GetDashboardRequest) Reset() {
	*x = GetDashboardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eqxcustomers_getdashboard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDashboardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDashboardRequest) ProtoMessage() {}

func (x *GetDashboardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eqxcustomers_getdashboard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDashboardRequest.ProtoReflect.Descriptor instead.
func (*GetDashboardRequest) Descriptor() ([]byte, []int) {
	return file_eqxcustomers_getdashboard_proto_rawDescGZIP(), []int{0}
}

func (x *GetDashboardRequest) GetProductPage() string {
	if x != nil {
		return x.ProductPage
	}
	return ""
}

type GetDashboardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Groups []*BlockGroups `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *GetDashboardResponse) Reset() {
	*x = GetDashboardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eqxcustomers_getdashboard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDashboardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDashboardResponse) ProtoMessage() {}

func (x *GetDashboardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eqxcustomers_getdashboard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDashboardResponse.ProtoReflect.Descriptor instead.
func (*GetDashboardResponse) Descriptor() ([]byte, []int) {
	return file_eqxcustomers_getdashboard_proto_rawDescGZIP(), []int{1}
}

func (x *GetDashboardResponse) GetGroups() []*BlockGroups {
	if x != nil {
		return x.Groups
	}
	return nil
}

type BlockGroups struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupName string   `protobuf:"bytes,1,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	SeqNo     *int32   `protobuf:"varint,2,opt,name=seq_no,json=seqNo,proto3,oneof" json:"seq_no,omitempty"`
	Blocks    []*Block `protobuf:"bytes,3,rep,name=blocks,proto3" json:"blocks,omitempty"`
}

func (x *BlockGroups) Reset() {
	*x = BlockGroups{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eqxcustomers_getdashboard_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockGroups) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockGroups) ProtoMessage() {}

func (x *BlockGroups) ProtoReflect() protoreflect.Message {
	mi := &file_eqxcustomers_getdashboard_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockGroups.ProtoReflect.Descriptor instead.
func (*BlockGroups) Descriptor() ([]byte, []int) {
	return file_eqxcustomers_getdashboard_proto_rawDescGZIP(), []int{2}
}

func (x *BlockGroups) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *BlockGroups) GetSeqNo() int32 {
	if x != nil && x.SeqNo != nil {
		return *x.SeqNo
	}
	return 0
}

func (x *BlockGroups) GetBlocks() []*Block {
	if x != nil {
		return x.Blocks
	}
	return nil
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockName     string   `protobuf:"bytes,1,opt,name=block_name,json=blockName,proto3" json:"block_name,omitempty"`
	BlockSequence *int32   `protobuf:"varint,2,opt,name=block_sequence,json=blockSequence,proto3,oneof" json:"block_sequence,omitempty"`
	Items         []*Items `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	BlockType     *string  `protobuf:"bytes,4,opt,name=block_type,json=blockType,proto3,oneof" json:"block_type,omitempty"`
	BlockTitle    *string  `protobuf:"bytes,5,opt,name=block_title,json=blockTitle,proto3,oneof" json:"block_title,omitempty"`
	BlockCardSize *int32   `protobuf:"varint,6,opt,name=block_card_size,json=blockCardSize,proto3,oneof" json:"block_card_size,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eqxcustomers_getdashboard_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_eqxcustomers_getdashboard_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_eqxcustomers_getdashboard_proto_rawDescGZIP(), []int{3}
}

func (x *Block) GetBlockName() string {
	if x != nil {
		return x.BlockName
	}
	return ""
}

func (x *Block) GetBlockSequence() int32 {
	if x != nil && x.BlockSequence != nil {
		return *x.BlockSequence
	}
	return 0
}

func (x *Block) GetItems() []*Items {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Block) GetBlockType() string {
	if x != nil && x.BlockType != nil {
		return *x.BlockType
	}
	return ""
}

func (x *Block) GetBlockTitle() string {
	if x != nil && x.BlockTitle != nil {
		return *x.BlockTitle
	}
	return ""
}

func (x *Block) GetBlockCardSize() int32 {
	if x != nil && x.BlockCardSize != nil {
		return *x.BlockCardSize
	}
	return 0
}

type Items struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text          *string `protobuf:"bytes,1,opt,name=text,proto3,oneof" json:"text,omitempty"`
	SubText       *string `protobuf:"bytes,2,opt,name=sub_text,json=subText,proto3,oneof" json:"sub_text,omitempty"`
	ImageUrl      *string `protobuf:"bytes,3,opt,name=image_url,json=imageUrl,proto3,oneof" json:"image_url,omitempty"`
	VideoUrl      *string `protobuf:"bytes,4,opt,name=video_url,json=videoUrl,proto3,oneof" json:"video_url,omitempty"`
	ActionUrl     *string `protobuf:"bytes,5,opt,name=action_url,json=actionUrl,proto3,oneof" json:"action_url,omitempty"`
	AppVersion    *string `protobuf:"bytes,6,opt,name=app_version,json=appVersion,proto3,oneof" json:"app_version,omitempty"`
	Badge         *string `protobuf:"bytes,7,opt,name=badge,proto3,oneof" json:"badge,omitempty"`
	ActionType    *string `protobuf:"bytes,8,opt,name=action_type,json=actionType,proto3,oneof" json:"action_type,omitempty"`
	ItemsMetadata *string `protobuf:"bytes,9,opt,name=items_metadata,json=itemsMetadata,proto3,oneof" json:"items_metadata,omitempty"`
	ItemSequence  *int32  `protobuf:"varint,10,opt,name=item_sequence,json=itemSequence,proto3,oneof" json:"item_sequence,omitempty"`
	Rules         *string `protobuf:"bytes,11,opt,name=rules,proto3,oneof" json:"rules,omitempty"`
}

func (x *Items) Reset() {
	*x = Items{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eqxcustomers_getdashboard_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Items) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Items) ProtoMessage() {}

func (x *Items) ProtoReflect() protoreflect.Message {
	mi := &file_eqxcustomers_getdashboard_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Items.ProtoReflect.Descriptor instead.
func (*Items) Descriptor() ([]byte, []int) {
	return file_eqxcustomers_getdashboard_proto_rawDescGZIP(), []int{4}
}

func (x *Items) GetText() string {
	if x != nil && x.Text != nil {
		return *x.Text
	}
	return ""
}

func (x *Items) GetSubText() string {
	if x != nil && x.SubText != nil {
		return *x.SubText
	}
	return ""
}

func (x *Items) GetImageUrl() string {
	if x != nil && x.ImageUrl != nil {
		return *x.ImageUrl
	}
	return ""
}

func (x *Items) GetVideoUrl() string {
	if x != nil && x.VideoUrl != nil {
		return *x.VideoUrl
	}
	return ""
}

func (x *Items) GetActionUrl() string {
	if x != nil && x.ActionUrl != nil {
		return *x.ActionUrl
	}
	return ""
}

func (x *Items) GetAppVersion() string {
	if x != nil && x.AppVersion != nil {
		return *x.AppVersion
	}
	return ""
}

func (x *Items) GetBadge() string {
	if x != nil && x.Badge != nil {
		return *x.Badge
	}
	return ""
}

func (x *Items) GetActionType() string {
	if x != nil && x.ActionType != nil {
		return *x.ActionType
	}
	return ""
}

func (x *Items) GetItemsMetadata() string {
	if x != nil && x.ItemsMetadata != nil {
		return *x.ItemsMetadata
	}
	return ""
}

func (x *Items) GetItemSequence() int32 {
	if x != nil && x.ItemSequence != nil {
		return *x.ItemSequence
	}
	return 0
}

func (x *Items) GetRules() string {
	if x != nil && x.Rules != nil {
		return *x.Rules
	}
	return ""
}

var File_eqxcustomers_getdashboard_proto protoreflect.FileDescriptor

var file_eqxcustomers_getdashboard_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x65, 0x71, 0x78, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x67,
	0x65, 0x74, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x19, 0x65, 0x71, 0x78, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e,
	0x67, 0x65, 0x74, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x22, 0x38, 0x0a, 0x13,
	0x67, 0x65, 0x74, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x50, 0x61, 0x67, 0x65, 0x22, 0x57, 0x0a, 0x14, 0x67, 0x65, 0x74, 0x44, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f,
	0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x65, 0x71, 0x78, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x67, 0x65,
	0x74, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22,
	0x8e, 0x01, 0x0a, 0x0c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x06, 0x73, 0x65, 0x71, 0x5f, 0x6e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x00, 0x52, 0x05, 0x73, 0x65, 0x71, 0x4e, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x38, 0x0a, 0x06, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x71,
	0x78, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x67, 0x65, 0x74, 0x64, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x06, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x65, 0x71, 0x5f, 0x6e, 0x6f,
	0x22, 0xc7, 0x02, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x0e, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x00, 0x52, 0x0d, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x36, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x71, 0x78, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x73, 0x2e, 0x67, 0x65, 0x74, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x22, 0x0a,
	0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x01, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x24, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x03, 0x52, 0x0d, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x53, 0x69, 0x7a,
	0x65, 0x88, 0x01, 0x01, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x9a, 0x04, 0x0a, 0x05, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x12, 0x17, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a,
	0x08, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x07, 0x73, 0x75, 0x62, 0x54, 0x65, 0x78, 0x74, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a,
	0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x02, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x12,
	0x20, 0x0a, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x03, 0x52, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x55, 0x72, 0x6c, 0x88, 0x01,
	0x01, 0x12, 0x22, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x09, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55,
	0x72, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x0a, 0x61, 0x70,
	0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x62,
	0x61, 0x64, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x05, 0x62, 0x61,
	0x64, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52, 0x0a, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x0e,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x08, 0x52, 0x0d, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x69, 0x74, 0x65, 0x6d,
	0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x09, 0x52, 0x0c, 0x69, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x0a, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x73, 0x75, 0x62, 0x5f, 0x74,
	0x65, 0x78, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72,
	0x6c, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x42, 0x0e,
	0x0a, 0x0c, 0x5f, 0x61, 0x70, 0x70, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x62, 0x61, 0x64, 0x67, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x10, 0x0a, 0x0e, 0x5f,
	0x69, 0x74, 0x65, 0x6d, 0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eqxcustomers_getdashboard_proto_rawDescOnce sync.Once
	file_eqxcustomers_getdashboard_proto_rawDescData = file_eqxcustomers_getdashboard_proto_rawDesc
)

func file_eqxcustomers_getdashboard_proto_rawDescGZIP() []byte {
	file_eqxcustomers_getdashboard_proto_rawDescOnce.Do(func() {
		file_eqxcustomers_getdashboard_proto_rawDescData = protoimpl.X.CompressGZIP(file_eqxcustomers_getdashboard_proto_rawDescData)
	})
	return file_eqxcustomers_getdashboard_proto_rawDescData
}

var file_eqxcustomers_getdashboard_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_eqxcustomers_getdashboard_proto_goTypes = []interface{}{
	(*GetDashboardRequest)(nil),  // 0: eqxcustomers.getdashboard.getDashboardRequest
	(*GetDashboardResponse)(nil), // 1: eqxcustomers.getdashboard.getDashboardResponse
	(*BlockGroups)(nil),          // 2: eqxcustomers.getdashboard.Block_groups
	(*Block)(nil),                // 3: eqxcustomers.getdashboard.Block
	(*Items)(nil),                // 4: eqxcustomers.getdashboard.Items
}
var file_eqxcustomers_getdashboard_proto_depIdxs = []int32{
	2, // 0: eqxcustomers.getdashboard.getDashboardResponse.groups:type_name -> eqxcustomers.getdashboard.Block_groups
	3, // 1: eqxcustomers.getdashboard.Block_groups.blocks:type_name -> eqxcustomers.getdashboard.Block
	4, // 2: eqxcustomers.getdashboard.Block.items:type_name -> eqxcustomers.getdashboard.Items
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_eqxcustomers_getdashboard_proto_init() }
func file_eqxcustomers_getdashboard_proto_init() {
	if File_eqxcustomers_getdashboard_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eqxcustomers_getdashboard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDashboardRequest); i {
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
		file_eqxcustomers_getdashboard_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDashboardResponse); i {
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
		file_eqxcustomers_getdashboard_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockGroups); i {
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
		file_eqxcustomers_getdashboard_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_eqxcustomers_getdashboard_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Items); i {
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
	file_eqxcustomers_getdashboard_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_eqxcustomers_getdashboard_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_eqxcustomers_getdashboard_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eqxcustomers_getdashboard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eqxcustomers_getdashboard_proto_goTypes,
		DependencyIndexes: file_eqxcustomers_getdashboard_proto_depIdxs,
		MessageInfos:      file_eqxcustomers_getdashboard_proto_msgTypes,
	}.Build()
	File_eqxcustomers_getdashboard_proto = out.File
	file_eqxcustomers_getdashboard_proto_rawDesc = nil
	file_eqxcustomers_getdashboard_proto_goTypes = nil
	file_eqxcustomers_getdashboard_proto_depIdxs = nil
}
