// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: proto/generation.proto

package generation

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

type Role int32

const (
	Role_USER      Role = 0
	Role_ASSISTANT Role = 1
	Role_SYSTEM    Role = 2
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "USER",
		1: "ASSISTANT",
		2: "SYSTEM",
	}
	Role_value = map[string]int32{
		"USER":      0,
		"ASSISTANT": 1,
		"SYSTEM":    2,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_generation_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_proto_generation_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_proto_generation_proto_rawDescGZIP(), []int{0}
}

type Chat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Role    Role   `protobuf:"varint,2,opt,name=role,proto3,enum=generation.Role" json:"role,omitempty"`
}

func (x *Chat) Reset() {
	*x = Chat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_generation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chat) ProtoMessage() {}

func (x *Chat) ProtoReflect() protoreflect.Message {
	mi := &file_proto_generation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chat.ProtoReflect.Descriptor instead.
func (*Chat) Descriptor() ([]byte, []int) {
	return file_proto_generation_proto_rawDescGZIP(), []int{0}
}

func (x *Chat) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Chat) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_USER
}

type GenerateResponseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chat    []*Chat `protobuf:"bytes,1,rep,name=chat,proto3" json:"chat,omitempty"`
	ModelId string  `protobuf:"bytes,2,opt,name=model_id,json=modelId,proto3" json:"model_id,omitempty"`
}

func (x *GenerateResponseRequest) Reset() {
	*x = GenerateResponseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_generation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateResponseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateResponseRequest) ProtoMessage() {}

func (x *GenerateResponseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_generation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateResponseRequest.ProtoReflect.Descriptor instead.
func (*GenerateResponseRequest) Descriptor() ([]byte, []int) {
	return file_proto_generation_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateResponseRequest) GetChat() []*Chat {
	if x != nil {
		return x.Chat
	}
	return nil
}

func (x *GenerateResponseRequest) GetModelId() string {
	if x != nil {
		return x.ModelId
	}
	return ""
}

type StreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *StreamResponse) Reset() {
	*x = StreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_generation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamResponse) ProtoMessage() {}

func (x *StreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_generation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamResponse.ProtoReflect.Descriptor instead.
func (*StreamResponse) Descriptor() ([]byte, []int) {
	return file_proto_generation_proto_rawDescGZIP(), []int{2}
}

func (x *StreamResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type BatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *BatchResponse) Reset() {
	*x = BatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_generation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchResponse) ProtoMessage() {}

func (x *BatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_generation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchResponse.ProtoReflect.Descriptor instead.
func (*BatchResponse) Descriptor() ([]byte, []int) {
	return file_proto_generation_proto_rawDescGZIP(), []int{3}
}

func (x *BatchResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_proto_generation_proto protoreflect.FileDescriptor

var file_proto_generation_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x46, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x5a, 0x0a, 0x17,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x63, 0x68, 0x61, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x04, 0x63, 0x68, 0x61, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x64, 0x22, 0x26, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x2b, 0x0a, 0x0d, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x2b, 0x0a,
	0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x53, 0x45, 0x52, 0x10, 0x00, 0x12,
	0x0d, 0x0a, 0x09, 0x41, 0x53, 0x53, 0x49, 0x53, 0x54, 0x41, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0a,
	0x0a, 0x06, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x10, 0x02, 0x32, 0xc9, 0x01, 0x0a, 0x11, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x5b, 0x0a, 0x16, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x23, 0x2e, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x57, 0x0a,
	0x15, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x73, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x23, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x12, 0x5a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_generation_proto_rawDescOnce sync.Once
	file_proto_generation_proto_rawDescData = file_proto_generation_proto_rawDesc
)

func file_proto_generation_proto_rawDescGZIP() []byte {
	file_proto_generation_proto_rawDescOnce.Do(func() {
		file_proto_generation_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_generation_proto_rawDescData)
	})
	return file_proto_generation_proto_rawDescData
}

var file_proto_generation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_generation_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_generation_proto_goTypes = []any{
	(Role)(0),                       // 0: generation.Role
	(*Chat)(nil),                    // 1: generation.Chat
	(*GenerateResponseRequest)(nil), // 2: generation.GenerateResponseRequest
	(*StreamResponse)(nil),          // 3: generation.StreamResponse
	(*BatchResponse)(nil),           // 4: generation.BatchResponse
}
var file_proto_generation_proto_depIdxs = []int32{
	0, // 0: generation.Chat.role:type_name -> generation.Role
	1, // 1: generation.GenerateResponseRequest.chat:type_name -> generation.Chat
	2, // 2: generation.GenerationService.GenerateResponseStream:input_type -> generation.GenerateResponseRequest
	2, // 3: generation.GenerationService.GenerateRepsonseBatch:input_type -> generation.GenerateResponseRequest
	3, // 4: generation.GenerationService.GenerateResponseStream:output_type -> generation.StreamResponse
	4, // 5: generation.GenerationService.GenerateRepsonseBatch:output_type -> generation.BatchResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_generation_proto_init() }
func file_proto_generation_proto_init() {
	if File_proto_generation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_generation_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Chat); i {
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
		file_proto_generation_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GenerateResponseRequest); i {
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
		file_proto_generation_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*StreamResponse); i {
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
		file_proto_generation_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*BatchResponse); i {
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
			RawDescriptor: file_proto_generation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_generation_proto_goTypes,
		DependencyIndexes: file_proto_generation_proto_depIdxs,
		EnumInfos:         file_proto_generation_proto_enumTypes,
		MessageInfos:      file_proto_generation_proto_msgTypes,
	}.Build()
	File_proto_generation_proto = out.File
	file_proto_generation_proto_rawDesc = nil
	file_proto_generation_proto_goTypes = nil
	file_proto_generation_proto_depIdxs = nil
}
