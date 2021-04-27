//
//Define messages to create new plugin to TKS-Plugin service

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.5
// source: plugin.proto

package pbgo

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

// StartDeployPluginRequest is a request to deploy a Kubernetes plugin
type StartDeployPluginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// cluster_id is a unique id for the cluster
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// kubeconfig is a default kubeconfig for the cluster
	Kubeconfig string `protobuf:"bytes,2,opt,name=kubeconfig,proto3" json:"kubeconfig,omitempty"`
	// plugin is a plugin name defined by users
	Plugin string `protobuf:"bytes,3,opt,name=plugin,proto3" json:"plugin,omitempty"`
	// version is a plugin version
	Version string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *StartDeployPluginRequest) Reset() {
	*x = StartDeployPluginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartDeployPluginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartDeployPluginRequest) ProtoMessage() {}

func (x *StartDeployPluginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartDeployPluginRequest.ProtoReflect.Descriptor instead.
func (*StartDeployPluginRequest) Descriptor() ([]byte, []int) {
	return file_plugin_proto_rawDescGZIP(), []int{0}
}

func (x *StartDeployPluginRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *StartDeployPluginRequest) GetKubeconfig() string {
	if x != nil {
		return x.Kubeconfig
	}
	return ""
}

func (x *StartDeployPluginRequest) GetPlugin() string {
	if x != nil {
		return x.Plugin
	}
	return ""
}

func (x *StartDeployPluginRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

// StartDeployPluginResponse is a response to the StartDeployPluginRequest
//  from the plugin service.
type StartDeployPluginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// code is a response code.
	Code Code `protobuf:"varint,1,opt,name=code,proto3,enum=pbgo.Code" json:"code,omitempty"`
	// error is a detailed error message.
	Error *Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	// plugin_id is a global unique ID created by the plugin service.
	PluginId string `protobuf:"bytes,3,opt,name=plugin_id,json=pluginId,proto3" json:"plugin_id,omitempty"`
}

func (x *StartDeployPluginResponse) Reset() {
	*x = StartDeployPluginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartDeployPluginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartDeployPluginResponse) ProtoMessage() {}

func (x *StartDeployPluginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartDeployPluginResponse.ProtoReflect.Descriptor instead.
func (*StartDeployPluginResponse) Descriptor() ([]byte, []int) {
	return file_plugin_proto_rawDescGZIP(), []int{1}
}

func (x *StartDeployPluginResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_OK_UNSPECIFIED
}

func (x *StartDeployPluginResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *StartDeployPluginResponse) GetPluginId() string {
	if x != nil {
		return x.PluginId
	}
	return ""
}

// StopDeployPluginRequest is a request to stop deploy a Kubernetes plugin
type StopDeployPluginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// plugin_id is a global unique ID created by the plugin service.
	PluginId string `protobuf:"bytes,1,opt,name=plugin_id,json=pluginId,proto3" json:"plugin_id,omitempty"`
}

func (x *StopDeployPluginRequest) Reset() {
	*x = StopDeployPluginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopDeployPluginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopDeployPluginRequest) ProtoMessage() {}

func (x *StopDeployPluginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopDeployPluginRequest.ProtoReflect.Descriptor instead.
func (*StopDeployPluginRequest) Descriptor() ([]byte, []int) {
	return file_plugin_proto_rawDescGZIP(), []int{2}
}

func (x *StopDeployPluginRequest) GetPluginId() string {
	if x != nil {
		return x.PluginId
	}
	return ""
}

// GetPluginsRequest is a request to get information
//  about plugin that is deployed
type GetPluginsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// cluster_id is a unique id for the cluster
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
}

func (x *GetPluginsRequest) Reset() {
	*x = GetPluginsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPluginsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPluginsRequest) ProtoMessage() {}

func (x *GetPluginsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPluginsRequest.ProtoReflect.Descriptor instead.
func (*GetPluginsRequest) Descriptor() ([]byte, []int) {
	return file_plugin_proto_rawDescGZIP(), []int{3}
}

func (x *GetPluginsRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

// GetPluginsResponse is a response to the GetPluginsRequest
// from the plugin service.
type GetPluginsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// code is a response code.
	Code Code `protobuf:"varint,1,opt,name=code,proto3,enum=pbgo.Code" json:"code,omitempty"`
	// error is a detailed error message.
	Error *Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	// plugin_id is a global unique ID created by the plugin service.
	PluginId string `protobuf:"bytes,3,opt,name=plugin_id,json=pluginId,proto3" json:"plugin_id,omitempty"`
	// cluster_id is a unique id for the cluster
	ClusterId string `protobuf:"bytes,4,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// plugin is a plugin name defined by users
	Plugin string `protobuf:"bytes,5,opt,name=plugin,proto3" json:"plugin,omitempty"`
	// version is a plugin version
	Version string `protobuf:"bytes,6,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *GetPluginsResponse) Reset() {
	*x = GetPluginsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPluginsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPluginsResponse) ProtoMessage() {}

func (x *GetPluginsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPluginsResponse.ProtoReflect.Descriptor instead.
func (*GetPluginsResponse) Descriptor() ([]byte, []int) {
	return file_plugin_proto_rawDescGZIP(), []int{4}
}

func (x *GetPluginsResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_OK_UNSPECIFIED
}

func (x *GetPluginsResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *GetPluginsResponse) GetPluginId() string {
	if x != nil {
		return x.PluginId
	}
	return ""
}

func (x *GetPluginsResponse) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *GetPluginsResponse) GetPlugin() string {
	if x != nil {
		return x.Plugin
	}
	return ""
}

func (x *GetPluginsResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_plugin_proto protoreflect.FileDescriptor

var file_plugin_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x70, 0x62, 0x67, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x8b, 0x01, 0x0a, 0x18, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0x7b, 0x0a, 0x19, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x70, 0x62,
	0x67, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70,
	0x62, 0x67, 0x6f, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x36, 0x0a,
	0x17, 0x53, 0x74, 0x6f, 0x70, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x22, 0xc5, 0x01, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1e, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a,
	0x2e, 0x70, 0x62, 0x67, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x21, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x70, 0x62, 0x67, 0x6f, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x32, 0xf5, 0x01, 0x0a, 0x0d, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x1e, 0x2e, 0x70, 0x62, 0x67, 0x6f, 0x2e,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x62, 0x67, 0x6f, 0x2e,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x10, 0x53,
	0x74, 0x6f, 0x70, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12,
	0x1d, 0x2e, 0x70, 0x62, 0x67, 0x6f, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x70, 0x62, 0x67, 0x6f, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x73, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x67, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x70, 0x62, 0x67, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6b, 0x74, 0x65, 0x6c, 0x65, 0x63, 0x6f,
	0x6d, 0x2f, 0x74, 0x6b, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x67, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plugin_proto_rawDescOnce sync.Once
	file_plugin_proto_rawDescData = file_plugin_proto_rawDesc
)

func file_plugin_proto_rawDescGZIP() []byte {
	file_plugin_proto_rawDescOnce.Do(func() {
		file_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugin_proto_rawDescData)
	})
	return file_plugin_proto_rawDescData
}

var file_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_plugin_proto_goTypes = []interface{}{
	(*StartDeployPluginRequest)(nil),  // 0: pbgo.StartDeployPluginRequest
	(*StartDeployPluginResponse)(nil), // 1: pbgo.StartDeployPluginResponse
	(*StopDeployPluginRequest)(nil),   // 2: pbgo.StopDeployPluginRequest
	(*GetPluginsRequest)(nil),         // 3: pbgo.GetPluginsRequest
	(*GetPluginsResponse)(nil),        // 4: pbgo.GetPluginsResponse
	(Code)(0),                         // 5: pbgo.Code
	(*Error)(nil),                     // 6: pbgo.Error
	(*SimpleResponse)(nil),            // 7: pbgo.SimpleResponse
}
var file_plugin_proto_depIdxs = []int32{
	5, // 0: pbgo.StartDeployPluginResponse.code:type_name -> pbgo.Code
	6, // 1: pbgo.StartDeployPluginResponse.error:type_name -> pbgo.Error
	5, // 2: pbgo.GetPluginsResponse.code:type_name -> pbgo.Code
	6, // 3: pbgo.GetPluginsResponse.error:type_name -> pbgo.Error
	0, // 4: pbgo.PluginService.StartDeployPlugin:input_type -> pbgo.StartDeployPluginRequest
	2, // 5: pbgo.PluginService.StopDeployPlugin:input_type -> pbgo.StopDeployPluginRequest
	3, // 6: pbgo.PluginService.GetPlugins:input_type -> pbgo.GetPluginsRequest
	1, // 7: pbgo.PluginService.StartDeployPlugin:output_type -> pbgo.StartDeployPluginResponse
	7, // 8: pbgo.PluginService.StopDeployPlugin:output_type -> pbgo.SimpleResponse
	4, // 9: pbgo.PluginService.GetPlugins:output_type -> pbgo.GetPluginsResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_plugin_proto_init() }
func file_plugin_proto_init() {
	if File_plugin_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_plugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartDeployPluginRequest); i {
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
		file_plugin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartDeployPluginResponse); i {
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
		file_plugin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopDeployPluginRequest); i {
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
		file_plugin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPluginsRequest); i {
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
		file_plugin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPluginsResponse); i {
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
			RawDescriptor: file_plugin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_plugin_proto_goTypes,
		DependencyIndexes: file_plugin_proto_depIdxs,
		MessageInfos:      file_plugin_proto_msgTypes,
	}.Build()
	File_plugin_proto = out.File
	file_plugin_proto_rawDesc = nil
	file_plugin_proto_goTypes = nil
	file_plugin_proto_depIdxs = nil
}
