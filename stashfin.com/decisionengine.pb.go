// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.19.6
// source: decisionengine.proto

package stashfin_com

import (
	decisionengine "github.com/stashfin2/grpc/stashfin.com/decisionengine"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_decisionengine_proto protoreflect.FileDescriptor

var file_decisionengine_proto_rawDesc = string([]byte{
	0x0a, 0x14, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a,
	0x29, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f,
	0x6c, 0x6f, 0x63, 0x64, 0x69, 0x73, 0x62, 0x75, 0x72, 0x73, 0x61, 0x6c, 0x61, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x64, 0x65, 0x63, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x64, 0x65, 0x63, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x73, 0x61, 0x76, 0x65, 0x63, 0x6f, 0x6e, 0x73, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc2, 0x03, 0x0a, 0x0f, 0x64, 0x65, 0x63,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x97, 0x01, 0x0a,
	0x14, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x44, 0x69, 0x73, 0x62,
	0x75, 0x72, 0x73, 0x61, 0x6c, 0x12, 0x3d, 0x2e, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4c,
	0x6f, 0x63, 0x44, 0x69, 0x73, 0x62, 0x75, 0x72, 0x73, 0x61, 0x6c, 0x2e, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x62, 0x75, 0x72, 0x73, 0x61, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x3e, 0x2e, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f,
	0x63, 0x44, 0x69, 0x73, 0x62, 0x75, 0x72, 0x73, 0x61, 0x6c, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x62, 0x75, 0x72, 0x73, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xa2, 0x01, 0x0a, 0x15, 0x64, 0x65, 0x63, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x12, 0x42, 0x2e, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x43, 0x2e, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x64, 0x65, 0x63, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x70, 0x0a, 0x0b, 0x73,
	0x61, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x2e, 0x64, 0x65, 0x63,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x73, 0x61, 0x76, 0x65,
	0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x73,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x64, 0x65, 0x63,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x73, 0x61, 0x76, 0x65,
	0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x73,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var file_decisionengine_proto_goTypes = []any{
	(*decisionengine.ValidateDisbursalRequest)(nil),      // 0: decisionengine.validateLocDisbursal.validateDisbursalRequest
	(*decisionengine.DecisionEngineTriggerRequest)(nil),  // 1: decisionengine.decisionEngineTrigger.decisionEngineTriggerRequest
	(*decisionengine.SaveConsentRequest)(nil),            // 2: decisionengine.saveConsent.saveConsentRequest
	(*decisionengine.ValidateDisbursalResponse)(nil),     // 3: decisionengine.validateLocDisbursal.validateDisbursalResponse
	(*decisionengine.DecisionEngineTriggerResponse)(nil), // 4: decisionengine.decisionEngineTrigger.decisionEngineTriggerResponse
	(*decisionengine.SaveConsentResponse)(nil),           // 5: decisionengine.saveConsent.saveConsentResponse
}
var file_decisionengine_proto_depIdxs = []int32{
	0, // 0: service.decision_engine.validateLocDisbursal:input_type -> decisionengine.validateLocDisbursal.validateDisbursalRequest
	1, // 1: service.decision_engine.decisionEngineTrigger:input_type -> decisionengine.decisionEngineTrigger.decisionEngineTriggerRequest
	2, // 2: service.decision_engine.saveConsent:input_type -> decisionengine.saveConsent.saveConsentRequest
	3, // 3: service.decision_engine.validateLocDisbursal:output_type -> decisionengine.validateLocDisbursal.validateDisbursalResponse
	4, // 4: service.decision_engine.decisionEngineTrigger:output_type -> decisionengine.decisionEngineTrigger.decisionEngineTriggerResponse
	5, // 5: service.decision_engine.saveConsent:output_type -> decisionengine.saveConsent.saveConsentResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_decisionengine_proto_init() }
func file_decisionengine_proto_init() {
	if File_decisionengine_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_decisionengine_proto_rawDesc), len(file_decisionengine_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_decisionengine_proto_goTypes,
		DependencyIndexes: file_decisionengine_proto_depIdxs,
	}.Build()
	File_decisionengine_proto = out.File
	file_decisionengine_proto_goTypes = nil
	file_decisionengine_proto_depIdxs = nil
}
