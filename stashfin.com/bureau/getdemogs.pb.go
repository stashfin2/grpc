// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: bureau/getdemogs.proto

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

type DemogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportId int64 `protobuf:"varint,1,opt,name=report_id,json=reportId,proto3" json:"report_id,omitempty"`
}

func (x *DemogsRequest) Reset() {
	*x = DemogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bureau_getdemogs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemogsRequest) ProtoMessage() {}

func (x *DemogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bureau_getdemogs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemogsRequest.ProtoReflect.Descriptor instead.
func (*DemogsRequest) Descriptor() ([]byte, []int) {
	return file_bureau_getdemogs_proto_rawDescGZIP(), []int{0}
}

func (x *DemogsRequest) GetReportId() int64 {
	if x != nil {
		return x.ReportId
	}
	return 0
}

type DemogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Data   *Demogs `protobuf:"bytes,2,opt,name=data,proto3,oneof" json:"data,omitempty"`
}

func (x *DemogsResponse) Reset() {
	*x = DemogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bureau_getdemogs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemogsResponse) ProtoMessage() {}

func (x *DemogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bureau_getdemogs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemogsResponse.ProtoReflect.Descriptor instead.
func (*DemogsResponse) Descriptor() ([]byte, []int) {
	return file_bureau_getdemogs_proto_rawDescGZIP(), []int{1}
}

func (x *DemogsResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *DemogsResponse) GetData() *Demogs {
	if x != nil {
		return x.Data
	}
	return nil
}

type Demogs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Personal *Personal  `protobuf:"bytes,1,opt,name=personal,proto3" json:"personal,omitempty"`
	Ids      []*Ids     `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids,omitempty"`
	Address  []*Address `protobuf:"bytes,3,rep,name=address,proto3" json:"address,omitempty"`
	Phone    []*Phone   `protobuf:"bytes,4,rep,name=phone,proto3" json:"phone,omitempty"`
	Emails   []*Emails  `protobuf:"bytes,5,rep,name=emails,proto3" json:"emails,omitempty"`
}

func (x *Demogs) Reset() {
	*x = Demogs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bureau_getdemogs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Demogs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Demogs) ProtoMessage() {}

func (x *Demogs) ProtoReflect() protoreflect.Message {
	mi := &file_bureau_getdemogs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Demogs.ProtoReflect.Descriptor instead.
func (*Demogs) Descriptor() ([]byte, []int) {
	return file_bureau_getdemogs_proto_rawDescGZIP(), []int{2}
}

func (x *Demogs) GetPersonal() *Personal {
	if x != nil {
		return x.Personal
	}
	return nil
}

func (x *Demogs) GetIds() []*Ids {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *Demogs) GetAddress() []*Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Demogs) GetPhone() []*Phone {
	if x != nil {
		return x.Phone
	}
	return nil
}

func (x *Demogs) GetEmails() []*Emails {
	if x != nil {
		return x.Emails
	}
	return nil
}

type Personal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Dob        string `protobuf:"bytes,2,opt,name=dob,proto3" json:"dob,omitempty"`
	Gender     string `protobuf:"bytes,3,opt,name=gender,proto3" json:"gender,omitempty"`
	Occupation string `protobuf:"bytes,4,opt,name=occupation,proto3" json:"occupation,omitempty"`
	Income     int64  `protobuf:"varint,5,opt,name=income,proto3" json:"income,omitempty"`
}

func (x *Personal) Reset() {
	*x = Personal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bureau_getdemogs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Personal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Personal) ProtoMessage() {}

func (x *Personal) ProtoReflect() protoreflect.Message {
	mi := &file_bureau_getdemogs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Personal.ProtoReflect.Descriptor instead.
func (*Personal) Descriptor() ([]byte, []int) {
	return file_bureau_getdemogs_proto_rawDescGZIP(), []int{3}
}

func (x *Personal) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Personal) GetDob() string {
	if x != nil {
		return x.Dob
	}
	return ""
}

func (x *Personal) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *Personal) GetOccupation() string {
	if x != nil {
		return x.Occupation
	}
	return ""
}

func (x *Personal) GetIncome() int64 {
	if x != nil {
		return x.Income
	}
	return 0
}

type Ids struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Ids) Reset() {
	*x = Ids{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bureau_getdemogs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ids) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ids) ProtoMessage() {}

func (x *Ids) ProtoReflect() protoreflect.Message {
	mi := &file_bureau_getdemogs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ids.ProtoReflect.Descriptor instead.
func (*Ids) Descriptor() ([]byte, []int) {
	return file_bureau_getdemogs_proto_rawDescGZIP(), []int{4}
}

func (x *Ids) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Ids) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	State   string `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	Pincode int64  `protobuf:"varint,3,opt,name=pincode,proto3" json:"pincode,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bureau_getdemogs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_bureau_getdemogs_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_bureau_getdemogs_proto_rawDescGZIP(), []int{5}
}

func (x *Address) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Address) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Address) GetPincode() int64 {
	if x != nil {
		return x.Pincode
	}
	return 0
}

type Phone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *Phone) Reset() {
	*x = Phone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bureau_getdemogs_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Phone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Phone) ProtoMessage() {}

func (x *Phone) ProtoReflect() protoreflect.Message {
	mi := &file_bureau_getdemogs_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Phone.ProtoReflect.Descriptor instead.
func (*Phone) Descriptor() ([]byte, []int) {
	return file_bureau_getdemogs_proto_rawDescGZIP(), []int{6}
}

func (x *Phone) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type Emails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *Emails) Reset() {
	*x = Emails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bureau_getdemogs_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Emails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Emails) ProtoMessage() {}

func (x *Emails) ProtoReflect() protoreflect.Message {
	mi := &file_bureau_getdemogs_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Emails.ProtoReflect.Descriptor instead.
func (*Emails) Descriptor() ([]byte, []int) {
	return file_bureau_getdemogs_proto_rawDescGZIP(), []int{7}
}

func (x *Emails) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_bureau_getdemogs_proto protoreflect.FileDescriptor

var file_bureau_getdemogs_proto_rawDesc = []byte{
	0x0a, 0x16, 0x62, 0x75, 0x72, 0x65, 0x61, 0x75, 0x2f, 0x67, 0x65, 0x74, 0x64, 0x65, 0x6d, 0x6f,
	0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x62, 0x75, 0x72, 0x65, 0x61, 0x75,
	0x2e, 0x67, 0x65, 0x74, 0x64, 0x65, 0x6d, 0x6f, 0x67, 0x73, 0x22, 0x2c, 0x0a, 0x0d, 0x64, 0x65,
	0x6d, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x22, 0x64, 0x0a, 0x0e, 0x64, 0x65, 0x6d, 0x6f,
	0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x31, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x62, 0x75, 0x72, 0x65, 0x61, 0x75, 0x2e, 0x67, 0x65, 0x74, 0x64, 0x65, 0x6d,
	0x6f, 0x67, 0x73, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x67, 0x73, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x22, 0xff,
	0x01, 0x0a, 0x06, 0x44, 0x65, 0x6d, 0x6f, 0x67, 0x73, 0x12, 0x36, 0x0a, 0x08, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x75,
	0x72, 0x65, 0x61, 0x75, 0x2e, 0x67, 0x65, 0x74, 0x64, 0x65, 0x6d, 0x6f, 0x67, 0x73, 0x2e, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61,
	0x6c, 0x12, 0x27, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x62, 0x75, 0x72, 0x65, 0x61, 0x75, 0x2e, 0x67, 0x65, 0x74, 0x64, 0x65, 0x6d, 0x6f, 0x67,
	0x73, 0x2e, 0x49, 0x64, 0x73, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x75,
	0x72, 0x65, 0x61, 0x75, 0x2e, 0x67, 0x65, 0x74, 0x64, 0x65, 0x6d, 0x6f, 0x67, 0x73, 0x2e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x2d, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x62, 0x75, 0x72, 0x65, 0x61, 0x75, 0x2e, 0x67, 0x65, 0x74, 0x64, 0x65, 0x6d, 0x6f, 0x67,
	0x73, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x30,
	0x0a, 0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x62, 0x75, 0x72, 0x65, 0x61, 0x75, 0x2e, 0x67, 0x65, 0x74, 0x64, 0x65, 0x6d, 0x6f, 0x67,
	0x73, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73,
	0x22, 0x80, 0x01, 0x0a, 0x08, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x6f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x64, 0x6f, 0x62, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x6f,
	0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x69,
	0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x6e, 0x63,
	0x6f, 0x6d, 0x65, 0x22, 0x2f, 0x0a, 0x03, 0x49, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x53, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x70, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x1d, 0x0a, 0x05, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x1e, 0x0a, 0x06, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bureau_getdemogs_proto_rawDescOnce sync.Once
	file_bureau_getdemogs_proto_rawDescData = file_bureau_getdemogs_proto_rawDesc
)

func file_bureau_getdemogs_proto_rawDescGZIP() []byte {
	file_bureau_getdemogs_proto_rawDescOnce.Do(func() {
		file_bureau_getdemogs_proto_rawDescData = protoimpl.X.CompressGZIP(file_bureau_getdemogs_proto_rawDescData)
	})
	return file_bureau_getdemogs_proto_rawDescData
}

var file_bureau_getdemogs_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_bureau_getdemogs_proto_goTypes = []interface{}{
	(*DemogsRequest)(nil),  // 0: bureau.getdemogs.demogsRequest
	(*DemogsResponse)(nil), // 1: bureau.getdemogs.demogsResponse
	(*Demogs)(nil),         // 2: bureau.getdemogs.Demogs
	(*Personal)(nil),       // 3: bureau.getdemogs.Personal
	(*Ids)(nil),            // 4: bureau.getdemogs.Ids
	(*Address)(nil),        // 5: bureau.getdemogs.Address
	(*Phone)(nil),          // 6: bureau.getdemogs.Phone
	(*Emails)(nil),         // 7: bureau.getdemogs.Emails
}
var file_bureau_getdemogs_proto_depIdxs = []int32{
	2, // 0: bureau.getdemogs.demogsResponse.data:type_name -> bureau.getdemogs.Demogs
	3, // 1: bureau.getdemogs.Demogs.personal:type_name -> bureau.getdemogs.Personal
	4, // 2: bureau.getdemogs.Demogs.ids:type_name -> bureau.getdemogs.Ids
	5, // 3: bureau.getdemogs.Demogs.address:type_name -> bureau.getdemogs.Address
	6, // 4: bureau.getdemogs.Demogs.phone:type_name -> bureau.getdemogs.Phone
	7, // 5: bureau.getdemogs.Demogs.emails:type_name -> bureau.getdemogs.Emails
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_bureau_getdemogs_proto_init() }
func file_bureau_getdemogs_proto_init() {
	if File_bureau_getdemogs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bureau_getdemogs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemogsRequest); i {
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
		file_bureau_getdemogs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemogsResponse); i {
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
		file_bureau_getdemogs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Demogs); i {
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
		file_bureau_getdemogs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Personal); i {
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
		file_bureau_getdemogs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ids); i {
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
		file_bureau_getdemogs_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_bureau_getdemogs_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Phone); i {
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
		file_bureau_getdemogs_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Emails); i {
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
	file_bureau_getdemogs_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bureau_getdemogs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bureau_getdemogs_proto_goTypes,
		DependencyIndexes: file_bureau_getdemogs_proto_depIdxs,
		MessageInfos:      file_bureau_getdemogs_proto_msgTypes,
	}.Build()
	File_bureau_getdemogs_proto = out.File
	file_bureau_getdemogs_proto_rawDesc = nil
	file_bureau_getdemogs_proto_goTypes = nil
	file_bureau_getdemogs_proto_depIdxs = nil
}
