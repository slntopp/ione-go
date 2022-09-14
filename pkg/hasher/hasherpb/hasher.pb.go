// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: pkg/hasher/hasherpb/hasher.proto

package hasherpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_pkg_hasher_hasherpb_hasher_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50000,
		Name:          "nocloud.hasher.hash",
		Tag:           "varint,50000,opt,name=hash",
		Filename:      "pkg/hasher/hasherpb/hasher.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50010,
		Name:          "nocloud.hasher.hashed",
		Tag:           "varint,50010,opt,name=hashed",
		Filename:      "pkg/hasher/hasherpb/hasher.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50020,
		Name:          "nocloud.hasher.skipped",
		Tag:           "varint,50020,opt,name=skipped",
		Filename:      "pkg/hasher/hasherpb/hasher.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional bool hash = 50000;
	E_Hash = &file_pkg_hasher_hasherpb_hasher_proto_extTypes[0]
	// optional bool hashed = 50010;
	E_Hashed = &file_pkg_hasher_hasherpb_hasher_proto_extTypes[1]
	// optional bool skipped = 50020;
	E_Skipped = &file_pkg_hasher_hasherpb_hasher_proto_extTypes[2]
)

var File_pkg_hasher_hasherpb_hasher_proto protoreflect.FileDescriptor

var file_pkg_hasher_hasherpb_hasher_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x6b, 0x67, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x65, 0x72, 0x2f, 0x68, 0x61, 0x73,
	0x68, 0x65, 0x72, 0x70, 0x62, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x68, 0x61, 0x73, 0x68,
	0x65, 0x72, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x33, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x1d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd0, 0x86, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x3a, 0x37, 0x0a, 0x06, 0x68, 0x61, 0x73,
	0x68, 0x65, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xda, 0x86, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x68, 0x61, 0x73, 0x68,
	0x65, 0x64, 0x3a, 0x39, 0x0a, 0x07, 0x73, 0x6b, 0x69, 0x70, 0x70, 0x65, 0x64, 0x12, 0x1d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe4, 0x86, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x6b, 0x69, 0x70, 0x70, 0x65, 0x64, 0x42, 0xaa, 0x01,
	0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x68, 0x61,
	0x73, 0x68, 0x65, 0x72, 0x42, 0x0b, 0x48, 0x61, 0x73, 0x68, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6c, 0x6e, 0x74, 0x6f, 0x70, 0x70, 0x2f, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x65, 0x72, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x65,
	0x72, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x4e, 0x48, 0x58, 0xaa, 0x02, 0x0e, 0x4e, 0x6f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x65, 0x72, 0xca, 0x02, 0x0e, 0x4e, 0x6f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x48, 0x61, 0x73, 0x68, 0x65, 0x72, 0xe2, 0x02, 0x1a, 0x4e, 0x6f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x48, 0x61, 0x73, 0x68, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x4e, 0x6f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x3a, 0x3a, 0x48, 0x61, 0x73, 0x68, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_pkg_hasher_hasherpb_hasher_proto_goTypes = []interface{}{
	(*descriptorpb.FieldOptions)(nil), // 0: google.protobuf.FieldOptions
}
var file_pkg_hasher_hasherpb_hasher_proto_depIdxs = []int32{
	0, // 0: nocloud.hasher.hash:extendee -> google.protobuf.FieldOptions
	0, // 1: nocloud.hasher.hashed:extendee -> google.protobuf.FieldOptions
	0, // 2: nocloud.hasher.skipped:extendee -> google.protobuf.FieldOptions
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	0, // [0:3] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_hasher_hasherpb_hasher_proto_init() }
func file_pkg_hasher_hasherpb_hasher_proto_init() {
	if File_pkg_hasher_hasherpb_hasher_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_hasher_hasherpb_hasher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 3,
			NumServices:   0,
		},
		GoTypes:           file_pkg_hasher_hasherpb_hasher_proto_goTypes,
		DependencyIndexes: file_pkg_hasher_hasherpb_hasher_proto_depIdxs,
		ExtensionInfos:    file_pkg_hasher_hasherpb_hasher_proto_extTypes,
	}.Build()
	File_pkg_hasher_hasherpb_hasher_proto = out.File
	file_pkg_hasher_hasherpb_hasher_proto_rawDesc = nil
	file_pkg_hasher_hasherpb_hasher_proto_goTypes = nil
	file_pkg_hasher_hasherpb_hasher_proto_depIdxs = nil
}
