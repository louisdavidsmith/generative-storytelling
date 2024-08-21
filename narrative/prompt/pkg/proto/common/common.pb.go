// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: proto/common.proto

package common

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

type CheckOutcome int32

const (
	CheckOutcome_UNEXPECTED_POSITIVE_OUTCOME  CheckOutcome = 0
	CheckOutcome_SUCCESSFUL                   CheckOutcome = 1
	CheckOutcome_SUCCESSFUL_WITH_COMPLICATION CheckOutcome = 2
	CheckOutcome_FAIL_WITH_BOON               CheckOutcome = 3
	CheckOutcome_FAIL_SIMPLE                  CheckOutcome = 4
	CheckOutcome_FAIL_WITH_COMPLICATIONS      CheckOutcome = 5
)

// Enum value maps for CheckOutcome.
var (
	CheckOutcome_name = map[int32]string{
		0: "UNEXPECTED_POSITIVE_OUTCOME",
		1: "SUCCESSFUL",
		2: "SUCCESSFUL_WITH_COMPLICATION",
		3: "FAIL_WITH_BOON",
		4: "FAIL_SIMPLE",
		5: "FAIL_WITH_COMPLICATIONS",
	}
	CheckOutcome_value = map[string]int32{
		"UNEXPECTED_POSITIVE_OUTCOME":  0,
		"SUCCESSFUL":                   1,
		"SUCCESSFUL_WITH_COMPLICATION": 2,
		"FAIL_WITH_BOON":               3,
		"FAIL_SIMPLE":                  4,
		"FAIL_WITH_COMPLICATIONS":      5,
	}
)

func (x CheckOutcome) Enum() *CheckOutcome {
	p := new(CheckOutcome)
	*p = x
	return p
}

func (x CheckOutcome) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CheckOutcome) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_common_proto_enumTypes[0].Descriptor()
}

func (CheckOutcome) Type() protoreflect.EnumType {
	return &file_proto_common_proto_enumTypes[0]
}

func (x CheckOutcome) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CheckOutcome.Descriptor instead.
func (CheckOutcome) EnumDescriptor() ([]byte, []int) {
	return file_proto_common_proto_rawDescGZIP(), []int{0}
}

var File_proto_common_proto protoreflect.FileDescriptor

var file_proto_common_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2a, 0xa3, 0x01, 0x0a,
	0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x1f, 0x0a,
	0x1b, 0x55, 0x4e, 0x45, 0x58, 0x50, 0x45, 0x43, 0x54, 0x45, 0x44, 0x5f, 0x50, 0x4f, 0x53, 0x49,
	0x54, 0x49, 0x56, 0x45, 0x5f, 0x4f, 0x55, 0x54, 0x43, 0x4f, 0x4d, 0x45, 0x10, 0x00, 0x12, 0x0e,
	0x0a, 0x0a, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x46, 0x55, 0x4c, 0x10, 0x01, 0x12, 0x20,
	0x0a, 0x1c, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x46, 0x55, 0x4c, 0x5f, 0x57, 0x49, 0x54,
	0x48, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02,
	0x12, 0x12, 0x0a, 0x0e, 0x46, 0x41, 0x49, 0x4c, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x42, 0x4f,
	0x4f, 0x4e, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x41, 0x49, 0x4c, 0x5f, 0x53, 0x49, 0x4d,
	0x50, 0x4c, 0x45, 0x10, 0x04, 0x12, 0x1b, 0x0a, 0x17, 0x46, 0x41, 0x49, 0x4c, 0x5f, 0x57, 0x49,
	0x54, 0x48, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53,
	0x10, 0x05, 0x42, 0x0e, 0x5a, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_common_proto_rawDescOnce sync.Once
	file_proto_common_proto_rawDescData = file_proto_common_proto_rawDesc
)

func file_proto_common_proto_rawDescGZIP() []byte {
	file_proto_common_proto_rawDescOnce.Do(func() {
		file_proto_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_common_proto_rawDescData)
	})
	return file_proto_common_proto_rawDescData
}

var file_proto_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_common_proto_goTypes = []any{
	(CheckOutcome)(0), // 0: common.CheckOutcome
}
var file_proto_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_common_proto_init() }
func file_proto_common_proto_init() {
	if File_proto_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_common_proto_goTypes,
		DependencyIndexes: file_proto_common_proto_depIdxs,
		EnumInfos:         file_proto_common_proto_enumTypes,
	}.Build()
	File_proto_common_proto = out.File
	file_proto_common_proto_rawDesc = nil
	file_proto_common_proto_goTypes = nil
	file_proto_common_proto_depIdxs = nil
}
