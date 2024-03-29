// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.10
// source: rpc/themes/service.proto

package themes

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddThemesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddThemesRequest) Reset() {
	*x = AddThemesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_themes_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddThemesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddThemesRequest) ProtoMessage() {}

func (x *AddThemesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_themes_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddThemesRequest.ProtoReflect.Descriptor instead.
func (*AddThemesRequest) Descriptor() ([]byte, []int) {
	return file_rpc_themes_service_proto_rawDescGZIP(), []int{0}
}

type GetThemesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetThemesRequest) Reset() {
	*x = GetThemesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_themes_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetThemesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetThemesRequest) ProtoMessage() {}

func (x *GetThemesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_themes_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetThemesRequest.ProtoReflect.Descriptor instead.
func (*GetThemesRequest) Descriptor() ([]byte, []int) {
	return file_rpc_themes_service_proto_rawDescGZIP(), []int{1}
}

type GetThemesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ThemeData []*Theme `protobuf:"bytes,1,rep,name=themeData,proto3" json:"themeData,omitempty"`
}

func (x *GetThemesResponse) Reset() {
	*x = GetThemesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_themes_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetThemesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetThemesResponse) ProtoMessage() {}

func (x *GetThemesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_themes_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetThemesResponse.ProtoReflect.Descriptor instead.
func (*GetThemesResponse) Descriptor() ([]byte, []int) {
	return file_rpc_themes_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetThemesResponse) GetThemeData() []*Theme {
	if x != nil {
		return x.ThemeData
	}
	return nil
}

type GetThemeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetThemeRequest) Reset() {
	*x = GetThemeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_themes_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetThemeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetThemeRequest) ProtoMessage() {}

func (x *GetThemeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_themes_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetThemeRequest.ProtoReflect.Descriptor instead.
func (*GetThemeRequest) Descriptor() ([]byte, []int) {
	return file_rpc_themes_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetThemeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateThemeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Background string `protobuf:"bytes,2,opt,name=background,proto3" json:"background,omitempty"`
	Foreground string `protobuf:"bytes,3,opt,name=foreground,proto3" json:"foreground,omitempty"`
}

func (x *CreateThemeRequest) Reset() {
	*x = CreateThemeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_themes_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateThemeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateThemeRequest) ProtoMessage() {}

func (x *CreateThemeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_themes_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateThemeRequest.ProtoReflect.Descriptor instead.
func (*CreateThemeRequest) Descriptor() ([]byte, []int) {
	return file_rpc_themes_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateThemeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateThemeRequest) GetBackground() string {
	if x != nil {
		return x.Background
	}
	return ""
}

func (x *CreateThemeRequest) GetForeground() string {
	if x != nil {
		return x.Foreground
	}
	return ""
}

type GetThemeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ThemeData *Theme `protobuf:"bytes,1,opt,name=themeData,proto3" json:"themeData,omitempty"`
}

func (x *GetThemeResponse) Reset() {
	*x = GetThemeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_themes_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetThemeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetThemeResponse) ProtoMessage() {}

func (x *GetThemeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_themes_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetThemeResponse.ProtoReflect.Descriptor instead.
func (*GetThemeResponse) Descriptor() ([]byte, []int) {
	return file_rpc_themes_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetThemeResponse) GetThemeData() *Theme {
	if x != nil {
		return x.ThemeData
	}
	return nil
}

type Theme struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Background string `protobuf:"bytes,2,opt,name=Background,proto3" json:"Background,omitempty"`
	Foreground string `protobuf:"bytes,3,opt,name=Foreground,proto3" json:"Foreground,omitempty"`
}

func (x *Theme) Reset() {
	*x = Theme{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_themes_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Theme) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Theme) ProtoMessage() {}

func (x *Theme) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_themes_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Theme.ProtoReflect.Descriptor instead.
func (*Theme) Descriptor() ([]byte, []int) {
	return file_rpc_themes_service_proto_rawDescGZIP(), []int{6}
}

func (x *Theme) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Theme) GetBackground() string {
	if x != nil {
		return x.Background
	}
	return ""
}

func (x *Theme) GetForeground() string {
	if x != nil {
		return x.Foreground
	}
	return ""
}

var File_rpc_themes_service_proto protoreflect.FileDescriptor

var file_rpc_themes_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x65, 0x6e, 0x63,
	0x68, 0x61, 0x2e, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x12, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x54, 0x68, 0x65,
	0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x47,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x6e, 0x63, 0x68, 0x61, 0x2e,
	0x74, 0x68, 0x65, 0x6d, 0x65, 0x73, 0x2e, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x52, 0x09, 0x74, 0x68,
	0x65, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x25, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x68,
	0x65, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x68,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x61, 0x63, 0x6b,
	0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x61,
	0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x6f, 0x72, 0x65,
	0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x6f,
	0x72, 0x65, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x22, 0x46, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54,
	0x68, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x09,
	0x74, 0x68, 0x65, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x73, 0x65, 0x6e, 0x63, 0x68, 0x61, 0x2e, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x73, 0x2e,
	0x54, 0x68, 0x65, 0x6d, 0x65, 0x52, 0x09, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x5b, 0x0a, 0x05, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x46, 0x6f, 0x72, 0x65, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x46, 0x6f, 0x72, 0x65, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x32, 0xbe, 0x02,
	0x0a, 0x06, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x73, 0x12, 0x4b, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54,
	0x68, 0x65, 0x6d, 0x65, 0x12, 0x1e, 0x2e, 0x73, 0x65, 0x6e, 0x63, 0x68, 0x61, 0x2e, 0x74, 0x68,
	0x65, 0x6d, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x6e, 0x63, 0x68, 0x61, 0x2e, 0x74, 0x68,
	0x65, 0x6d, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x68, 0x65, 0x6d, 0x65, 0x12, 0x21, 0x2e, 0x73, 0x65, 0x6e, 0x63, 0x68, 0x61, 0x2e, 0x74, 0x68,
	0x65, 0x6d, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x68, 0x65, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x6e, 0x63, 0x68, 0x61,
	0x2e, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x68, 0x65, 0x6d, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x54,
	0x68, 0x65, 0x6d, 0x65, 0x73, 0x12, 0x1f, 0x2e, 0x73, 0x65, 0x6e, 0x63, 0x68, 0x61, 0x2e, 0x74,
	0x68, 0x65, 0x6d, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4e,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x73, 0x12, 0x1f, 0x2e, 0x73, 0x65,
	0x6e, 0x63, 0x68, 0x61, 0x2e, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x68, 0x65, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73,
	0x65, 0x6e, 0x63, 0x68, 0x61, 0x2e, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x68, 0x65, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d,
	0x5a, 0x0b, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_themes_service_proto_rawDescOnce sync.Once
	file_rpc_themes_service_proto_rawDescData = file_rpc_themes_service_proto_rawDesc
)

func file_rpc_themes_service_proto_rawDescGZIP() []byte {
	file_rpc_themes_service_proto_rawDescOnce.Do(func() {
		file_rpc_themes_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_themes_service_proto_rawDescData)
	})
	return file_rpc_themes_service_proto_rawDescData
}

var file_rpc_themes_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_rpc_themes_service_proto_goTypes = []interface{}{
	(*AddThemesRequest)(nil),   // 0: sencha.themes.AddThemesRequest
	(*GetThemesRequest)(nil),   // 1: sencha.themes.GetThemesRequest
	(*GetThemesResponse)(nil),  // 2: sencha.themes.GetThemesResponse
	(*GetThemeRequest)(nil),    // 3: sencha.themes.GetThemeRequest
	(*CreateThemeRequest)(nil), // 4: sencha.themes.CreateThemeRequest
	(*GetThemeResponse)(nil),   // 5: sencha.themes.GetThemeResponse
	(*Theme)(nil),              // 6: sencha.themes.Theme
	(*emptypb.Empty)(nil),      // 7: google.protobuf.Empty
}
var file_rpc_themes_service_proto_depIdxs = []int32{
	6, // 0: sencha.themes.GetThemesResponse.themeData:type_name -> sencha.themes.Theme
	6, // 1: sencha.themes.GetThemeResponse.themeData:type_name -> sencha.themes.Theme
	3, // 2: sencha.themes.Themes.GetTheme:input_type -> sencha.themes.GetThemeRequest
	4, // 3: sencha.themes.Themes.CreateTheme:input_type -> sencha.themes.CreateThemeRequest
	0, // 4: sencha.themes.Themes.AddThemes:input_type -> sencha.themes.AddThemesRequest
	1, // 5: sencha.themes.Themes.GetThemes:input_type -> sencha.themes.GetThemesRequest
	5, // 6: sencha.themes.Themes.GetTheme:output_type -> sencha.themes.GetThemeResponse
	5, // 7: sencha.themes.Themes.CreateTheme:output_type -> sencha.themes.GetThemeResponse
	7, // 8: sencha.themes.Themes.AddThemes:output_type -> google.protobuf.Empty
	2, // 9: sencha.themes.Themes.GetThemes:output_type -> sencha.themes.GetThemesResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rpc_themes_service_proto_init() }
func file_rpc_themes_service_proto_init() {
	if File_rpc_themes_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_themes_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddThemesRequest); i {
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
		file_rpc_themes_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetThemesRequest); i {
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
		file_rpc_themes_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetThemesResponse); i {
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
		file_rpc_themes_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetThemeRequest); i {
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
		file_rpc_themes_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateThemeRequest); i {
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
		file_rpc_themes_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetThemeResponse); i {
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
		file_rpc_themes_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Theme); i {
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
			RawDescriptor: file_rpc_themes_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_themes_service_proto_goTypes,
		DependencyIndexes: file_rpc_themes_service_proto_depIdxs,
		MessageInfos:      file_rpc_themes_service_proto_msgTypes,
	}.Build()
	File_rpc_themes_service_proto = out.File
	file_rpc_themes_service_proto_rawDesc = nil
	file_rpc_themes_service_proto_goTypes = nil
	file_rpc_themes_service_proto_depIdxs = nil
}
