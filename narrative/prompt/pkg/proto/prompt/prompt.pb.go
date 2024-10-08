// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: proto/prompt.proto

package prompt

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	common "prompt/pkg/proto/common"
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
	return file_proto_prompt_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_proto_prompt_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_proto_prompt_proto_rawDescGZIP(), []int{0}
}

type CurrentSanity int32

const (
	CurrentSanity_LOW    CurrentSanity = 0
	CurrentSanity_MEDIUM CurrentSanity = 1
	CurrentSanity_HIGH   CurrentSanity = 2
)

// Enum value maps for CurrentSanity.
var (
	CurrentSanity_name = map[int32]string{
		0: "LOW",
		1: "MEDIUM",
		2: "HIGH",
	}
	CurrentSanity_value = map[string]int32{
		"LOW":    0,
		"MEDIUM": 1,
		"HIGH":   2,
	}
)

func (x CurrentSanity) Enum() *CurrentSanity {
	p := new(CurrentSanity)
	*p = x
	return p
}

func (x CurrentSanity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CurrentSanity) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_prompt_proto_enumTypes[1].Descriptor()
}

func (CurrentSanity) Type() protoreflect.EnumType {
	return &file_proto_prompt_proto_enumTypes[1]
}

func (x CurrentSanity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CurrentSanity.Descriptor instead.
func (CurrentSanity) EnumDescriptor() ([]byte, []int) {
	return file_proto_prompt_proto_rawDescGZIP(), []int{1}
}

type Chat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Role    Role   `protobuf:"varint,2,opt,name=role,proto3,enum=prompt.Role" json:"role,omitempty"`
}

func (x *Chat) Reset() {
	*x = Chat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_prompt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chat) ProtoMessage() {}

func (x *Chat) ProtoReflect() protoreflect.Message {
	mi := &file_proto_prompt_proto_msgTypes[0]
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
	return file_proto_prompt_proto_rawDescGZIP(), []int{0}
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

type GetNarrationPromptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoreContext string  `protobuf:"bytes,1,opt,name=lore_context,json=loreContext,proto3" json:"lore_context,omitempty"`
	GameState   string  `protobuf:"bytes,2,opt,name=game_state,json=gameState,proto3" json:"game_state,omitempty"`
	UserInput   string  `protobuf:"bytes,3,opt,name=user_input,json=userInput,proto3" json:"user_input,omitempty"`
	ChatHistory []*Chat `protobuf:"bytes,4,rep,name=chat_history,json=chatHistory,proto3" json:"chat_history,omitempty"`
}

func (x *GetNarrationPromptRequest) Reset() {
	*x = GetNarrationPromptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_prompt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNarrationPromptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNarrationPromptRequest) ProtoMessage() {}

func (x *GetNarrationPromptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_prompt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNarrationPromptRequest.ProtoReflect.Descriptor instead.
func (*GetNarrationPromptRequest) Descriptor() ([]byte, []int) {
	return file_proto_prompt_proto_rawDescGZIP(), []int{1}
}

func (x *GetNarrationPromptRequest) GetLoreContext() string {
	if x != nil {
		return x.LoreContext
	}
	return ""
}

func (x *GetNarrationPromptRequest) GetGameState() string {
	if x != nil {
		return x.GameState
	}
	return ""
}

func (x *GetNarrationPromptRequest) GetUserInput() string {
	if x != nil {
		return x.UserInput
	}
	return ""
}

func (x *GetNarrationPromptRequest) GetChatHistory() []*Chat {
	if x != nil {
		return x.ChatHistory
	}
	return nil
}

type Prompt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prompt []*Chat `protobuf:"bytes,1,rep,name=prompt,proto3" json:"prompt,omitempty"`
}

func (x *Prompt) Reset() {
	*x = Prompt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_prompt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Prompt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prompt) ProtoMessage() {}

func (x *Prompt) ProtoReflect() protoreflect.Message {
	mi := &file_proto_prompt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Prompt.ProtoReflect.Descriptor instead.
func (*Prompt) Descriptor() ([]byte, []int) {
	return file_proto_prompt_proto_rawDescGZIP(), []int{2}
}

func (x *Prompt) GetPrompt() []*Chat {
	if x != nil {
		return x.Prompt
	}
	return nil
}

type GetSkillCheckResolutionPromptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckOutcome common.CheckOutcome `protobuf:"varint,1,opt,name=check_outcome,json=checkOutcome,proto3,enum=common.CheckOutcome" json:"check_outcome,omitempty"`
	UserInput    string              `protobuf:"bytes,2,opt,name=user_input,json=userInput,proto3" json:"user_input,omitempty"`
	ChatHistory  []*Chat             `protobuf:"bytes,3,rep,name=chat_history,json=chatHistory,proto3" json:"chat_history,omitempty"`
}

func (x *GetSkillCheckResolutionPromptRequest) Reset() {
	*x = GetSkillCheckResolutionPromptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_prompt_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSkillCheckResolutionPromptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSkillCheckResolutionPromptRequest) ProtoMessage() {}

func (x *GetSkillCheckResolutionPromptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_prompt_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSkillCheckResolutionPromptRequest.ProtoReflect.Descriptor instead.
func (*GetSkillCheckResolutionPromptRequest) Descriptor() ([]byte, []int) {
	return file_proto_prompt_proto_rawDescGZIP(), []int{3}
}

func (x *GetSkillCheckResolutionPromptRequest) GetCheckOutcome() common.CheckOutcome {
	if x != nil {
		return x.CheckOutcome
	}
	return common.CheckOutcome(0)
}

func (x *GetSkillCheckResolutionPromptRequest) GetUserInput() string {
	if x != nil {
		return x.UserInput
	}
	return ""
}

func (x *GetSkillCheckResolutionPromptRequest) GetChatHistory() []*Chat {
	if x != nil {
		return x.ChatHistory
	}
	return nil
}

type GetSanityNarrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentSanity CurrentSanity `protobuf:"varint,1,opt,name=current_sanity,json=currentSanity,proto3,enum=prompt.CurrentSanity" json:"current_sanity,omitempty"`
	UserInput     string        `protobuf:"bytes,2,opt,name=user_input,json=userInput,proto3" json:"user_input,omitempty"`
	LoreContent   string        `protobuf:"bytes,3,opt,name=lore_content,json=loreContent,proto3" json:"lore_content,omitempty"`
	ChatHistory   []*Chat       `protobuf:"bytes,4,rep,name=chat_history,json=chatHistory,proto3" json:"chat_history,omitempty"`
}

func (x *GetSanityNarrationRequest) Reset() {
	*x = GetSanityNarrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_prompt_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSanityNarrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSanityNarrationRequest) ProtoMessage() {}

func (x *GetSanityNarrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_prompt_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSanityNarrationRequest.ProtoReflect.Descriptor instead.
func (*GetSanityNarrationRequest) Descriptor() ([]byte, []int) {
	return file_proto_prompt_proto_rawDescGZIP(), []int{4}
}

func (x *GetSanityNarrationRequest) GetCurrentSanity() CurrentSanity {
	if x != nil {
		return x.CurrentSanity
	}
	return CurrentSanity_LOW
}

func (x *GetSanityNarrationRequest) GetUserInput() string {
	if x != nil {
		return x.UserInput
	}
	return ""
}

func (x *GetSanityNarrationRequest) GetLoreContent() string {
	if x != nil {
		return x.LoreContent
	}
	return ""
}

func (x *GetSanityNarrationRequest) GetChatHistory() []*Chat {
	if x != nil {
		return x.ChatHistory
	}
	return nil
}

var File_proto_prompt_proto protoreflect.FileDescriptor

var file_proto_prompt_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x1a, 0x12, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x42, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x20, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x22, 0xad, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x72, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x6f, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x6f, 0x72, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x2f, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x68, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x6d,
	0x70, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x74, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x22, 0x2e, 0x0a, 0x06, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x24,
	0x0a, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x06, 0x70, 0x72,
	0x6f, 0x6d, 0x70, 0x74, 0x22, 0xb1, 0x01, 0x0a, 0x24, 0x47, 0x65, 0x74, 0x53, 0x6b, 0x69, 0x6c,
	0x6c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a,
	0x0d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x0c, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x2f, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x74, 0x5f,
	0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x0b, 0x63, 0x68, 0x61,
	0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x22, 0xcc, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74,
	0x53, 0x61, 0x6e, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x72, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x61, 0x6e, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15,
	0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53,
	0x61, 0x6e, 0x69, 0x74, 0x79, 0x52, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x61,
	0x6e, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x6f, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x6f, 0x72, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x68,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70,
	0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x74,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2a, 0x2b, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12,
	0x08, 0x0a, 0x04, 0x55, 0x53, 0x45, 0x52, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x53, 0x53,
	0x49, 0x53, 0x54, 0x41, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x59, 0x53, 0x54,
	0x45, 0x4d, 0x10, 0x02, 0x2a, 0x2e, 0x0a, 0x0d, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53,
	0x61, 0x6e, 0x69, 0x74, 0x79, 0x12, 0x07, 0x0a, 0x03, 0x4c, 0x4f, 0x57, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x4d, 0x45, 0x44, 0x49, 0x55, 0x4d, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x49,
	0x47, 0x48, 0x10, 0x02, 0x32, 0x82, 0x02, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x72,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x21, 0x2e, 0x70,
	0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x72, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12,
	0x5d, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74,
	0x12, 0x2c, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6b, 0x69,
	0x6c, 0x6c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e,
	0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x49,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x61, 0x6e, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x72,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x61, 0x6e, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x72, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x6d,
	0x70, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x42, 0x0e, 0x5a, 0x0c, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_prompt_proto_rawDescOnce sync.Once
	file_proto_prompt_proto_rawDescData = file_proto_prompt_proto_rawDesc
)

func file_proto_prompt_proto_rawDescGZIP() []byte {
	file_proto_prompt_proto_rawDescOnce.Do(func() {
		file_proto_prompt_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_prompt_proto_rawDescData)
	})
	return file_proto_prompt_proto_rawDescData
}

var file_proto_prompt_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_prompt_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_prompt_proto_goTypes = []any{
	(Role)(0),                         // 0: prompt.Role
	(CurrentSanity)(0),                // 1: prompt.CurrentSanity
	(*Chat)(nil),                      // 2: prompt.Chat
	(*GetNarrationPromptRequest)(nil), // 3: prompt.GetNarrationPromptRequest
	(*Prompt)(nil),                    // 4: prompt.Prompt
	(*GetSkillCheckResolutionPromptRequest)(nil), // 5: prompt.GetSkillCheckResolutionPromptRequest
	(*GetSanityNarrationRequest)(nil),            // 6: prompt.GetSanityNarrationRequest
	(common.CheckOutcome)(0),                     // 7: common.CheckOutcome
}
var file_proto_prompt_proto_depIdxs = []int32{
	0,  // 0: prompt.Chat.role:type_name -> prompt.Role
	2,  // 1: prompt.GetNarrationPromptRequest.chat_history:type_name -> prompt.Chat
	2,  // 2: prompt.Prompt.prompt:type_name -> prompt.Chat
	7,  // 3: prompt.GetSkillCheckResolutionPromptRequest.check_outcome:type_name -> common.CheckOutcome
	2,  // 4: prompt.GetSkillCheckResolutionPromptRequest.chat_history:type_name -> prompt.Chat
	1,  // 5: prompt.GetSanityNarrationRequest.current_sanity:type_name -> prompt.CurrentSanity
	2,  // 6: prompt.GetSanityNarrationRequest.chat_history:type_name -> prompt.Chat
	3,  // 7: prompt.PromptService.GetNarrationPrompt:input_type -> prompt.GetNarrationPromptRequest
	5,  // 8: prompt.PromptService.GetSkillCheckResolutionPrompt:input_type -> prompt.GetSkillCheckResolutionPromptRequest
	6,  // 9: prompt.PromptService.GetInsanityNarration:input_type -> prompt.GetSanityNarrationRequest
	4,  // 10: prompt.PromptService.GetNarrationPrompt:output_type -> prompt.Prompt
	4,  // 11: prompt.PromptService.GetSkillCheckResolutionPrompt:output_type -> prompt.Prompt
	4,  // 12: prompt.PromptService.GetInsanityNarration:output_type -> prompt.Prompt
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_prompt_proto_init() }
func file_proto_prompt_proto_init() {
	if File_proto_prompt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_prompt_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_proto_prompt_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetNarrationPromptRequest); i {
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
		file_proto_prompt_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Prompt); i {
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
		file_proto_prompt_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetSkillCheckResolutionPromptRequest); i {
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
		file_proto_prompt_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetSanityNarrationRequest); i {
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
			RawDescriptor: file_proto_prompt_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_prompt_proto_goTypes,
		DependencyIndexes: file_proto_prompt_proto_depIdxs,
		EnumInfos:         file_proto_prompt_proto_enumTypes,
		MessageInfos:      file_proto_prompt_proto_msgTypes,
	}.Build()
	File_proto_prompt_proto = out.File
	file_proto_prompt_proto_rawDesc = nil
	file_proto_prompt_proto_goTypes = nil
	file_proto_prompt_proto_depIdxs = nil
}
