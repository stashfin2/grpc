// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: banking.proto

package stashfin_com

import (
	aa "github.com/stashfin2/grpc/stashfin.com/banking/aa"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_banking_proto protoreflect.FileDescriptor

var file_banking_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1c, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e,
	0x67, 0x2f, 0x61, 0x61, 0x2f, 0x67, 0x65, 0x74, 0x62, 0x61, 0x6e, 0x6b, 0x6c, 0x69, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2f,
	0x61, 0x61, 0x2f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x61, 0x2f, 0x67, 0x65,
	0x74, 0x72, 0x70, 0x64, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x61, 0x2f, 0x67, 0x65, 0x74, 0x61, 0x61,
	0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x62,
	0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x61, 0x2f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x74, 0x65, 0x70, 0x65, 0x6e, 0x6e, 0x79, 0x64, 0x72, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x61, 0x2f, 0x67, 0x65,
	0x74, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x61, 0x2f,
	0x6e, 0x65, 0x74, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e,
	0x67, 0x2f, 0x61, 0x61, 0x2f, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x72, 0x70, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa1, 0x07, 0x0a, 0x07, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e,
	0x67, 0x12, 0x68, 0x0a, 0x0b, 0x67, 0x65, 0x74, 0x42, 0x61, 0x6e, 0x6b, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x2a, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x61, 0x2e, 0x67, 0x65,
	0x74, 0x62, 0x61, 0x6e, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x67, 0x65, 0x74, 0x42, 0x61, 0x6e,
	0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x62,
	0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x61, 0x2e, 0x67, 0x65, 0x74, 0x62, 0x61, 0x6e,
	0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x67, 0x65, 0x74, 0x42, 0x61, 0x6e, 0x6b, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x60, 0x0a, 0x0f, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24,
	0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x61, 0x2e, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x74, 0x65, 0x2e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x61,
	0x61, 0x2e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x2e, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x69, 0x0a,
	0x0a, 0x67, 0x65, 0x74, 0x52, 0x70, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x2b, 0x2e, 0x62, 0x61,
	0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x61, 0x2e, 0x67, 0x65, 0x74, 0x72, 0x70, 0x64, 0x6c,
	0x69, 0x6e, 0x6b, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x70, 0x64, 0x4c, 0x69, 0x6e,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x61, 0x61, 0x2e, 0x67, 0x65, 0x74, 0x72, 0x70, 0x64, 0x6c, 0x69, 0x6e, 0x6b,
	0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x70, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x0c, 0x67, 0x65, 0x74, 0x41,
	0x41, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x2c, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x61, 0x61, 0x2e, 0x67, 0x65, 0x74, 0x61, 0x61, 0x73, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x67, 0x65, 0x74, 0x41, 0x41, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x61, 0x61, 0x2e, 0x67, 0x65, 0x74, 0x61, 0x61, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x67, 0x65, 0x74, 0x41, 0x41, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x76, 0x0a, 0x11, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x74, 0x65, 0x50, 0x65, 0x6e, 0x6e, 0x79, 0x44, 0x72, 0x6f, 0x70, 0x12, 0x2e, 0x2e, 0x62,
	0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x61, 0x2e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x74, 0x65, 0x70, 0x65, 0x6e, 0x6e, 0x79, 0x64, 0x72, 0x6f, 0x70, 0x2e, 0x70, 0x65, 0x6e, 0x6e,
	0x79, 0x44, 0x72, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x62,
	0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x61, 0x2e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x74, 0x65, 0x70, 0x65, 0x6e, 0x6e, 0x79, 0x64, 0x72, 0x6f, 0x70, 0x2e, 0x70, 0x65, 0x6e, 0x6e,
	0x79, 0x44, 0x72, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x7c, 0x0a, 0x0f, 0x67, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x61,
	0x6e, 0x6b, 0x12, 0x32, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x61, 0x2e,
	0x67, 0x65, 0x74, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x62, 0x61, 0x6e, 0x6b, 0x2e,
	0x67, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x61, 0x6e, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x61, 0x61, 0x2e, 0x67, 0x65, 0x74, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x62,
	0x61, 0x6e, 0x6b, 0x2e, 0x67, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42,
	0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x80, 0x01,
	0x0a, 0x12, 0x6e, 0x65, 0x74, 0x42, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x74, 0x65, 0x12, 0x32, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x61,
	0x61, 0x2e, 0x6e, 0x65, 0x74, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x74, 0x65, 0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x50, 0x65, 0x72, 0x66, 0x69, 0x6f,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x61, 0x61, 0x2e, 0x6e, 0x65, 0x74, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x50,
	0x65, 0x72, 0x66, 0x69, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x77, 0x0a, 0x0c, 0x67, 0x65, 0x74, 0x52, 0x50, 0x44, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x31, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x61, 0x2e, 0x70, 0x6f,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x72, 0x70, 0x64, 0x2e, 0x67, 0x65, 0x74, 0x52, 0x50, 0x44, 0x50,
	0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x61,
	0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x72, 0x70, 0x64, 0x2e, 0x67, 0x65, 0x74, 0x52,
	0x50, 0x44, 0x50, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_banking_proto_goTypes = []interface{}{
	(*aa.GetBankListRequest)(nil),          // 0: banking.aa.getbanklist.getBankListRequest
	(*aa.InitiateRequest)(nil),             // 1: banking.aa.initiate.initiateRequest
	(*aa.CreateRpdLinkRequest)(nil),        // 2: banking.aa.getrpdlink.createRpdLinkRequest
	(*aa.GetAASupportRequest)(nil),         // 3: banking.aa.getaasupport.getAASupportRequest
	(*aa.PennyDropRequest)(nil),            // 4: banking.aa.initiatepennydrop.pennyDropRequest
	(*aa.GetCustomerBankRequest)(nil),      // 5: banking.aa.getcustomerbank.getCustomerBankRequest
	(*aa.InputPerfiosRequest)(nil),         // 6: banking.aa.netbankinginitiate.inputPerfiosRequest
	(*aa.GetRPDPollingStatusRequest)(nil),  // 7: banking.aa.pollingrpd.getRPDPollingStatusRequest
	(*aa.GetBankListResponse)(nil),         // 8: banking.aa.getbanklist.getBankListResponse
	(*aa.InitiateResponse)(nil),            // 9: banking.aa.initiate.initiateResponse
	(*aa.CreateRpdLinkResponse)(nil),       // 10: banking.aa.getrpdlink.createRpdLinkResponse
	(*aa.GetAASupportResponse)(nil),        // 11: banking.aa.getaasupport.getAASupportResponse
	(*aa.PennyDropResponse)(nil),           // 12: banking.aa.initiatepennydrop.pennyDropResponse
	(*aa.GetCustomerBankResponse)(nil),     // 13: banking.aa.getcustomerbank.getCustomerBankResponse
	(*aa.OutputPerfiosResponse)(nil),       // 14: banking.aa.netbankinginitiate.OutputPerfiosResponse
	(*aa.GetRPDPollingStatusResponse)(nil), // 15: banking.aa.pollingrpd.getRPDPollingStatusResponse
}
var file_banking_proto_depIdxs = []int32{
	0,  // 0: service.banking.getBankList:input_type -> banking.aa.getbanklist.getBankListRequest
	1,  // 1: service.banking.initiateRequest:input_type -> banking.aa.initiate.initiateRequest
	2,  // 2: service.banking.getRpdLink:input_type -> banking.aa.getrpdlink.createRpdLinkRequest
	3,  // 3: service.banking.getAASupport:input_type -> banking.aa.getaasupport.getAASupportRequest
	4,  // 4: service.banking.initiatePennyDrop:input_type -> banking.aa.initiatepennydrop.pennyDropRequest
	5,  // 5: service.banking.getCustomerBank:input_type -> banking.aa.getcustomerbank.getCustomerBankRequest
	6,  // 6: service.banking.netBankingInitiate:input_type -> banking.aa.netbankinginitiate.inputPerfiosRequest
	7,  // 7: service.banking.getRPDStatus:input_type -> banking.aa.pollingrpd.getRPDPollingStatusRequest
	8,  // 8: service.banking.getBankList:output_type -> banking.aa.getbanklist.getBankListResponse
	9,  // 9: service.banking.initiateRequest:output_type -> banking.aa.initiate.initiateResponse
	10, // 10: service.banking.getRpdLink:output_type -> banking.aa.getrpdlink.createRpdLinkResponse
	11, // 11: service.banking.getAASupport:output_type -> banking.aa.getaasupport.getAASupportResponse
	12, // 12: service.banking.initiatePennyDrop:output_type -> banking.aa.initiatepennydrop.pennyDropResponse
	13, // 13: service.banking.getCustomerBank:output_type -> banking.aa.getcustomerbank.getCustomerBankResponse
	14, // 14: service.banking.netBankingInitiate:output_type -> banking.aa.netbankinginitiate.OutputPerfiosResponse
	15, // 15: service.banking.getRPDStatus:output_type -> banking.aa.pollingrpd.getRPDPollingStatusResponse
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_banking_proto_init() }
func file_banking_proto_init() {
	if File_banking_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_banking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_banking_proto_goTypes,
		DependencyIndexes: file_banking_proto_depIdxs,
	}.Build()
	File_banking_proto = out.File
	file_banking_proto_rawDesc = nil
	file_banking_proto_goTypes = nil
	file_banking_proto_depIdxs = nil
}
