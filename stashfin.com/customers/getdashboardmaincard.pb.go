// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.19.6
// source: customers/getdashboardmaincard.proto

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

type GetDashboardMainCardRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetDashboardMainCardRequest) Reset() {
	*x = GetDashboardMainCardRequest{}
	mi := &file_customers_getdashboardmaincard_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDashboardMainCardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDashboardMainCardRequest) ProtoMessage() {}

func (x *GetDashboardMainCardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customers_getdashboardmaincard_proto_msgTypes[0]
	if x != nil {
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
	return file_customers_getdashboardmaincard_proto_rawDescGZIP(), []int{0}
}

type GetDashboardMainCardResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Code          string                 `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	BlockData     *MainCardData          `protobuf:"bytes,3,opt,name=block_data,json=blockData,proto3" json:"block_data,omitempty"`
	UserState     string                 `protobuf:"bytes,4,opt,name=user_state,json=userState,proto3" json:"user_state,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetDashboardMainCardResponse) Reset() {
	*x = GetDashboardMainCardResponse{}
	mi := &file_customers_getdashboardmaincard_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDashboardMainCardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDashboardMainCardResponse) ProtoMessage() {}

func (x *GetDashboardMainCardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customers_getdashboardmaincard_proto_msgTypes[1]
	if x != nil {
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
	return file_customers_getdashboardmaincard_proto_rawDescGZIP(), []int{1}
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
	state            protoimpl.MessageState `protogen:"open.v1"`
	Text             string                 `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	AmountText       string                 `protobuf:"bytes,2,opt,name=amount_text,json=amountText,proto3" json:"amount_text,omitempty"`
	SubText          string                 `protobuf:"bytes,3,opt,name=sub_text,json=subText,proto3" json:"sub_text,omitempty"`
	ActionButtonText string                 `protobuf:"bytes,4,opt,name=action_button_text,json=actionButtonText,proto3" json:"action_button_text,omitempty"`
	Action           string                 `protobuf:"bytes,5,opt,name=action,proto3" json:"action,omitempty"`
	ActionType       string                 `protobuf:"bytes,6,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty"`
	TotalLoc         int32                  `protobuf:"varint,7,opt,name=total_loc,json=totalLoc,proto3" json:"total_loc,omitempty"`
	UsedLoc          int32                  `protobuf:"varint,8,opt,name=used_loc,json=usedLoc,proto3" json:"used_loc,omitempty"`
	AvailableSpLoc   int32                  `protobuf:"varint,9,opt,name=available_sp_loc,json=availableSpLoc,proto3" json:"available_sp_loc,omitempty"`
	CanReloadCard    bool                   `protobuf:"varint,10,opt,name=can_reload_card,json=canReloadCard,proto3" json:"can_reload_card,omitempty"`
	DisbursedAmount  string                 `protobuf:"bytes,11,opt,name=disbursed_amount,json=disbursedAmount,proto3" json:"disbursed_amount,omitempty"`
	RemainingEmi     string                 `protobuf:"bytes,12,opt,name=remaining_emi,json=remainingEmi,proto3" json:"remaining_emi,omitempty"`
	BalanceEmiAmount int32                  `protobuf:"varint,13,opt,name=balance_emi_amount,json=balanceEmiAmount,proto3" json:"balance_emi_amount,omitempty"`
	Comment          string                 `protobuf:"bytes,14,opt,name=comment,proto3" json:"comment,omitempty"`
	LandingPage      string                 `protobuf:"bytes,15,opt,name=landing_page,json=landingPage,proto3" json:"landing_page,omitempty"`
	Timer            string                 `protobuf:"bytes,16,opt,name=timer,proto3" json:"timer,omitempty"`
	Image            string                 `protobuf:"bytes,17,opt,name=image,proto3" json:"image,omitempty"`
	LocError         string                 `protobuf:"bytes,18,opt,name=loc_error,json=locError,proto3" json:"loc_error,omitempty"`
	ShowButton       bool                   `protobuf:"varint,19,opt,name=show_button,json=showButton,proto3" json:"show_button,omitempty"`
	TotalStashCash   string                 `protobuf:"bytes,20,opt,name=total_stash_cash,json=totalStashCash,proto3" json:"total_stash_cash,omitempty"`
	SuppTitle        string                 `protobuf:"bytes,21,opt,name=supp_title,json=suppTitle,proto3" json:"supp_title,omitempty"`
	SuppLocStatus    int32                  `protobuf:"varint,22,opt,name=supp_loc_status,json=suppLocStatus,proto3" json:"supp_loc_status,omitempty"`
	SuppSubTitle     string                 `protobuf:"bytes,23,opt,name=supp_sub_title,json=suppSubTitle,proto3" json:"supp_sub_title,omitempty"`
	CardType         string                 `protobuf:"bytes,24,opt,name=card_type,json=cardType,proto3" json:"card_type,omitempty"`
	SanctionAmount   float32                `protobuf:"fixed32,25,opt,name=sanction_amount,json=sanctionAmount,proto3" json:"sanction_amount,omitempty"`
	StaticText       string                 `protobuf:"bytes,26,opt,name=static_text,json=staticText,proto3" json:"static_text,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *MainCardData) Reset() {
	*x = MainCardData{}
	mi := &file_customers_getdashboardmaincard_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MainCardData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MainCardData) ProtoMessage() {}

func (x *MainCardData) ProtoReflect() protoreflect.Message {
	mi := &file_customers_getdashboardmaincard_proto_msgTypes[2]
	if x != nil {
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
	return file_customers_getdashboardmaincard_proto_rawDescGZIP(), []int{2}
}

func (x *MainCardData) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *MainCardData) GetAmountText() string {
	if x != nil {
		return x.AmountText
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

func (x *MainCardData) GetTotalLoc() int32 {
	if x != nil {
		return x.TotalLoc
	}
	return 0
}

func (x *MainCardData) GetUsedLoc() int32 {
	if x != nil {
		return x.UsedLoc
	}
	return 0
}

func (x *MainCardData) GetAvailableSpLoc() int32 {
	if x != nil {
		return x.AvailableSpLoc
	}
	return 0
}

func (x *MainCardData) GetCanReloadCard() bool {
	if x != nil {
		return x.CanReloadCard
	}
	return false
}

func (x *MainCardData) GetDisbursedAmount() string {
	if x != nil {
		return x.DisbursedAmount
	}
	return ""
}

func (x *MainCardData) GetRemainingEmi() string {
	if x != nil {
		return x.RemainingEmi
	}
	return ""
}

func (x *MainCardData) GetBalanceEmiAmount() int32 {
	if x != nil {
		return x.BalanceEmiAmount
	}
	return 0
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

func (x *MainCardData) GetLocError() string {
	if x != nil {
		return x.LocError
	}
	return ""
}

func (x *MainCardData) GetShowButton() bool {
	if x != nil {
		return x.ShowButton
	}
	return false
}

func (x *MainCardData) GetTotalStashCash() string {
	if x != nil {
		return x.TotalStashCash
	}
	return ""
}

func (x *MainCardData) GetSuppTitle() string {
	if x != nil {
		return x.SuppTitle
	}
	return ""
}

func (x *MainCardData) GetSuppLocStatus() int32 {
	if x != nil {
		return x.SuppLocStatus
	}
	return 0
}

func (x *MainCardData) GetSuppSubTitle() string {
	if x != nil {
		return x.SuppSubTitle
	}
	return ""
}

func (x *MainCardData) GetCardType() string {
	if x != nil {
		return x.CardType
	}
	return ""
}

func (x *MainCardData) GetSanctionAmount() float32 {
	if x != nil {
		return x.SanctionAmount
	}
	return 0
}

func (x *MainCardData) GetStaticText() string {
	if x != nil {
		return x.StaticText
	}
	return ""
}

var File_customers_getdashboardmaincard_proto protoreflect.FileDescriptor

var file_customers_getdashboardmaincard_proto_rawDesc = string([]byte{
	0x0a, 0x24, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x64,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x6d, 0x61, 0x69, 0x6e, 0x63, 0x61, 0x72, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x73, 0x2e, 0x67, 0x65, 0x74, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x6d, 0x61,
	0x69, 0x6e, 0x63, 0x61, 0x72, 0x64, 0x22, 0x1d, 0x0a, 0x1b, 0x67, 0x65, 0x74, 0x44, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x4d, 0x61, 0x69, 0x6e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xb2, 0x01, 0x0a, 0x1c, 0x67, 0x65, 0x74, 0x44, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x4d, 0x61, 0x69, 0x6e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x4b,
	0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x67,
	0x65, 0x74, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x6d, 0x61, 0x69, 0x6e, 0x63,
	0x61, 0x72, 0x64, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x43, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0xf2, 0x06, 0x0a, 0x0c, 0x4d,
	0x61, 0x69, 0x6e, 0x43, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x65, 0x78, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x54, 0x65, 0x78, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x5f, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x75, 0x74, 0x74, 0x6f, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6c, 0x6f, 0x63, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4c, 0x6f, 0x63, 0x12,
	0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x6c, 0x6f, 0x63, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x75, 0x73, 0x65, 0x64, 0x4c, 0x6f, 0x63, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x70, 0x5f, 0x6c, 0x6f, 0x63, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53,
	0x70, 0x4c, 0x6f, 0x63, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x61, 0x6e, 0x5f, 0x72, 0x65, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x63,
	0x61, 0x6e, 0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x72, 0x64, 0x12, 0x29, 0x0a, 0x10,
	0x64, 0x69, 0x73, 0x62, 0x75, 0x72, 0x73, 0x65, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x69, 0x73, 0x62, 0x75, 0x72, 0x73, 0x65,
	0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x6d, 0x61, 0x69,
	0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x6d, 0x69, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x45, 0x6d, 0x69, 0x12, 0x2c, 0x0a, 0x12,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x65, 0x6d, 0x69, 0x5f, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x45, 0x6d, 0x69, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x61, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x50, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x72,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x68, 0x6f, 0x77, 0x42, 0x75, 0x74, 0x74, 0x6f,
	0x6e, 0x12, 0x28, 0x0a, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x73, 0x68,
	0x5f, 0x63, 0x61, 0x73, 0x68, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x53, 0x74, 0x61, 0x73, 0x68, 0x43, 0x61, 0x73, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x75, 0x70, 0x70, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x75, 0x70, 0x70, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x75,
	0x70, 0x70, 0x5f, 0x6c, 0x6f, 0x63, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x16, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0d, 0x73, 0x75, 0x70, 0x70, 0x4c, 0x6f, 0x63, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x75, 0x70, 0x70, 0x5f, 0x73, 0x75, 0x62, 0x5f, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x70, 0x70,
	0x53, 0x75, 0x62, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x64,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x72,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x61, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x19, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e,
	0x73, 0x61, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x1a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x54, 0x65, 0x78, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_customers_getdashboardmaincard_proto_rawDescOnce sync.Once
	file_customers_getdashboardmaincard_proto_rawDescData []byte
)

func file_customers_getdashboardmaincard_proto_rawDescGZIP() []byte {
	file_customers_getdashboardmaincard_proto_rawDescOnce.Do(func() {
		file_customers_getdashboardmaincard_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_customers_getdashboardmaincard_proto_rawDesc), len(file_customers_getdashboardmaincard_proto_rawDesc)))
	})
	return file_customers_getdashboardmaincard_proto_rawDescData
}

var file_customers_getdashboardmaincard_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_customers_getdashboardmaincard_proto_goTypes = []any{
	(*GetDashboardMainCardRequest)(nil),  // 0: customers.getdashboardmaincard.getDashboardMainCardRequest
	(*GetDashboardMainCardResponse)(nil), // 1: customers.getdashboardmaincard.getDashboardMainCardResponse
	(*MainCardData)(nil),                 // 2: customers.getdashboardmaincard.MainCardData
}
var file_customers_getdashboardmaincard_proto_depIdxs = []int32{
	2, // 0: customers.getdashboardmaincard.getDashboardMainCardResponse.block_data:type_name -> customers.getdashboardmaincard.MainCardData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_customers_getdashboardmaincard_proto_init() }
func file_customers_getdashboardmaincard_proto_init() {
	if File_customers_getdashboardmaincard_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_customers_getdashboardmaincard_proto_rawDesc), len(file_customers_getdashboardmaincard_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_customers_getdashboardmaincard_proto_goTypes,
		DependencyIndexes: file_customers_getdashboardmaincard_proto_depIdxs,
		MessageInfos:      file_customers_getdashboardmaincard_proto_msgTypes,
	}.Build()
	File_customers_getdashboardmaincard_proto = out.File
	file_customers_getdashboardmaincard_proto_goTypes = nil
	file_customers_getdashboardmaincard_proto_depIdxs = nil
}
