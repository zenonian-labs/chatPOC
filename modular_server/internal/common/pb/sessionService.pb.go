// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.15.8
// source: sessionService.proto

package pb

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

type RegisterClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientID string              `protobuf:"bytes,1,opt,name=ClientID,proto3" json:"ClientID,omitempty"`
	Handler  *HandlerInformation `protobuf:"bytes,2,opt,name=handler,proto3" json:"handler,omitempty"`
}

func (x *RegisterClientRequest) Reset() {
	*x = RegisterClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sessionService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterClientRequest) ProtoMessage() {}

func (x *RegisterClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sessionService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterClientRequest.ProtoReflect.Descriptor instead.
func (*RegisterClientRequest) Descriptor() ([]byte, []int) {
	return file_sessionService_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterClientRequest) GetClientID() string {
	if x != nil {
		return x.ClientID
	}
	return ""
}

func (x *RegisterClientRequest) GetHandler() *HandlerInformation {
	if x != nil {
		return x.Handler
	}
	return nil
}

type RegisterClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error bool   `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Data  string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RegisterClientResponse) Reset() {
	*x = RegisterClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sessionService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterClientResponse) ProtoMessage() {}

func (x *RegisterClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sessionService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterClientResponse.ProtoReflect.Descriptor instead.
func (*RegisterClientResponse) Descriptor() ([]byte, []int) {
	return file_sessionService_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterClientResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *RegisterClientResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type UnRegisterClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientID       string `protobuf:"bytes,1,opt,name=ClientID,proto3" json:"ClientID,omitempty"`
	HandlerID      string `protobuf:"bytes,2,opt,name=HandlerID,proto3" json:"HandlerID,omitempty"`
	HandlerAddress string `protobuf:"bytes,3,opt,name=HandlerAddress,proto3" json:"HandlerAddress,omitempty"`
}

func (x *UnRegisterClientRequest) Reset() {
	*x = UnRegisterClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sessionService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnRegisterClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnRegisterClientRequest) ProtoMessage() {}

func (x *UnRegisterClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sessionService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnRegisterClientRequest.ProtoReflect.Descriptor instead.
func (*UnRegisterClientRequest) Descriptor() ([]byte, []int) {
	return file_sessionService_proto_rawDescGZIP(), []int{2}
}

func (x *UnRegisterClientRequest) GetClientID() string {
	if x != nil {
		return x.ClientID
	}
	return ""
}

func (x *UnRegisterClientRequest) GetHandlerID() string {
	if x != nil {
		return x.HandlerID
	}
	return ""
}

func (x *UnRegisterClientRequest) GetHandlerAddress() string {
	if x != nil {
		return x.HandlerAddress
	}
	return ""
}

type UnRegisterClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error bool   `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Data  string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UnRegisterClientResponse) Reset() {
	*x = UnRegisterClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sessionService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnRegisterClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnRegisterClientResponse) ProtoMessage() {}

func (x *UnRegisterClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sessionService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnRegisterClientResponse.ProtoReflect.Descriptor instead.
func (*UnRegisterClientResponse) Descriptor() ([]byte, []int) {
	return file_sessionService_proto_rawDescGZIP(), []int{3}
}

func (x *UnRegisterClientResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *UnRegisterClientResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type GetWSHandlerWithClientIDsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientID string `protobuf:"bytes,1,opt,name=clientID,proto3" json:"clientID,omitempty"`
}

func (x *GetWSHandlerWithClientIDsRequest) Reset() {
	*x = GetWSHandlerWithClientIDsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sessionService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWSHandlerWithClientIDsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWSHandlerWithClientIDsRequest) ProtoMessage() {}

func (x *GetWSHandlerWithClientIDsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sessionService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWSHandlerWithClientIDsRequest.ProtoReflect.Descriptor instead.
func (*GetWSHandlerWithClientIDsRequest) Descriptor() ([]byte, []int) {
	return file_sessionService_proto_rawDescGZIP(), []int{4}
}

func (x *GetWSHandlerWithClientIDsRequest) GetClientID() string {
	if x != nil {
		return x.ClientID
	}
	return ""
}

type GetWSHandlerWithClientIDsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error    bool                  `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Msg      string                `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Handlers []*HandlerInformation `protobuf:"bytes,3,rep,name=handlers,proto3" json:"handlers,omitempty"`
}

func (x *GetWSHandlerWithClientIDsResponse) Reset() {
	*x = GetWSHandlerWithClientIDsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sessionService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWSHandlerWithClientIDsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWSHandlerWithClientIDsResponse) ProtoMessage() {}

func (x *GetWSHandlerWithClientIDsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sessionService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWSHandlerWithClientIDsResponse.ProtoReflect.Descriptor instead.
func (*GetWSHandlerWithClientIDsResponse) Descriptor() ([]byte, []int) {
	return file_sessionService_proto_rawDescGZIP(), []int{5}
}

func (x *GetWSHandlerWithClientIDsResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *GetWSHandlerWithClientIDsResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetWSHandlerWithClientIDsResponse) GetHandlers() []*HandlerInformation {
	if x != nil {
		return x.Handlers
	}
	return nil
}

var File_sessionService_proto protoreflect.FileDescriptor

var file_sessionService_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x15, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12,
	0x33, 0x0a, 0x07, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x68, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x72, 0x22, 0x42, 0x0a, 0x16, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x7b, 0x0a, 0x17, 0x55, 0x6e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12,
	0x1c, 0x0a, 0x09, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x49, 0x44, 0x12, 0x26, 0x0a,
	0x0e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x44, 0x0a, 0x18, 0x55, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3e, 0x0a, 0x20, 0x47,
	0x65, 0x74, 0x57, 0x53, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x82, 0x01, 0x0a, 0x21,
	0x47, 0x65, 0x74, 0x57, 0x53, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x35, 0x0a, 0x08, 0x68, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73,
	0x32, 0xa3, 0x02, 0x0a, 0x0e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x53, 0x0a, 0x10, 0x55, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55,
	0x6e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55,
	0x6e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6d, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x57, 0x53,
	0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x44, 0x12, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x57,
	0x53, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x44, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x53, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sessionService_proto_rawDescOnce sync.Once
	file_sessionService_proto_rawDescData = file_sessionService_proto_rawDesc
)

func file_sessionService_proto_rawDescGZIP() []byte {
	file_sessionService_proto_rawDescOnce.Do(func() {
		file_sessionService_proto_rawDescData = protoimpl.X.CompressGZIP(file_sessionService_proto_rawDescData)
	})
	return file_sessionService_proto_rawDescData
}

var file_sessionService_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_sessionService_proto_goTypes = []interface{}{
	(*RegisterClientRequest)(nil),             // 0: proto.RegisterClientRequest
	(*RegisterClientResponse)(nil),            // 1: proto.RegisterClientResponse
	(*UnRegisterClientRequest)(nil),           // 2: proto.UnRegisterClientRequest
	(*UnRegisterClientResponse)(nil),          // 3: proto.UnRegisterClientResponse
	(*GetWSHandlerWithClientIDsRequest)(nil),  // 4: proto.GetWSHandlerWithClientIDsRequest
	(*GetWSHandlerWithClientIDsResponse)(nil), // 5: proto.GetWSHandlerWithClientIDsResponse
	(*HandlerInformation)(nil),                // 6: proto.HandlerInformation
}
var file_sessionService_proto_depIdxs = []int32{
	6, // 0: proto.RegisterClientRequest.handler:type_name -> proto.HandlerInformation
	6, // 1: proto.GetWSHandlerWithClientIDsResponse.handlers:type_name -> proto.HandlerInformation
	0, // 2: proto.SessionService.RegisterClient:input_type -> proto.RegisterClientRequest
	2, // 3: proto.SessionService.UnRegisterClient:input_type -> proto.UnRegisterClientRequest
	4, // 4: proto.SessionService.GetWSHandlerWithClientID:input_type -> proto.GetWSHandlerWithClientIDsRequest
	1, // 5: proto.SessionService.RegisterClient:output_type -> proto.RegisterClientResponse
	3, // 6: proto.SessionService.UnRegisterClient:output_type -> proto.UnRegisterClientResponse
	5, // 7: proto.SessionService.GetWSHandlerWithClientID:output_type -> proto.GetWSHandlerWithClientIDsResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_sessionService_proto_init() }
func file_sessionService_proto_init() {
	if File_sessionService_proto != nil {
		return
	}
	file_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sessionService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterClientRequest); i {
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
		file_sessionService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterClientResponse); i {
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
		file_sessionService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnRegisterClientRequest); i {
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
		file_sessionService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnRegisterClientResponse); i {
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
		file_sessionService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWSHandlerWithClientIDsRequest); i {
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
		file_sessionService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWSHandlerWithClientIDsResponse); i {
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
			RawDescriptor: file_sessionService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sessionService_proto_goTypes,
		DependencyIndexes: file_sessionService_proto_depIdxs,
		MessageInfos:      file_sessionService_proto_msgTypes,
	}.Build()
	File_sessionService_proto = out.File
	file_sessionService_proto_rawDesc = nil
	file_sessionService_proto_goTypes = nil
	file_sessionService_proto_depIdxs = nil
}