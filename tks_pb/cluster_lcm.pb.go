//
//Define messages for TKS-ClusterLCM service.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: cluster_lcm.proto

package tks_pb

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

// CreateClusterRequest is a request to create a Kubernetes cluster
type CreateClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// contract_id is a contract ID.
	ContractId string `protobuf:"bytes,1,opt,name=contract_id,json=contractId,proto3" json:"contract_id,omitempty"`
	// csp_id is an ID of CSPInfo.
	CspId string `protobuf:"bytes,2,opt,name=csp_id,json=cspId,proto3" json:"csp_id,omitempty"`
	// cluster_name is a cluster name defined by users
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// conf is a cluster configuration
	Conf *ClusterConf `protobuf:"bytes,4,opt,name=conf,proto3" json:"conf,omitempty"`
}

func (x *CreateClusterRequest) Reset() {
	*x = CreateClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_lcm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClusterRequest) ProtoMessage() {}

func (x *CreateClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_lcm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClusterRequest.ProtoReflect.Descriptor instead.
func (*CreateClusterRequest) Descriptor() ([]byte, []int) {
	return file_cluster_lcm_proto_rawDescGZIP(), []int{0}
}

func (x *CreateClusterRequest) GetContractId() string {
	if x != nil {
		return x.ContractId
	}
	return ""
}

func (x *CreateClusterRequest) GetCspId() string {
	if x != nil {
		return x.CspId
	}
	return ""
}

func (x *CreateClusterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateClusterRequest) GetConf() *ClusterConf {
	if x != nil {
		return x.Conf
	}
	return nil
}

// ScaleClusterRequest is a request to scale the number of nodes on the cluster
type ScaleClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// cluster_id is Global Unique ID
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// master_replicas is the number of master node
	MasterReplicas int32 `protobuf:"varint,2,opt,name=master_replicas,json=masterReplicas,proto3" json:"master_replicas,omitempty"`
	// worker_replicas is the number of worker node
	WorkerReplicas int32 `protobuf:"varint,3,opt,name=worker_replicas,json=workerReplicas,proto3" json:"worker_replicas,omitempty"`
}

func (x *ScaleClusterRequest) Reset() {
	*x = ScaleClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_lcm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScaleClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScaleClusterRequest) ProtoMessage() {}

func (x *ScaleClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_lcm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScaleClusterRequest.ProtoReflect.Descriptor instead.
func (*ScaleClusterRequest) Descriptor() ([]byte, []int) {
	return file_cluster_lcm_proto_rawDescGZIP(), []int{1}
}

func (x *ScaleClusterRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *ScaleClusterRequest) GetMasterReplicas() int32 {
	if x != nil {
		return x.MasterReplicas
	}
	return 0
}

func (x *ScaleClusterRequest) GetWorkerReplicas() int32 {
	if x != nil {
		return x.WorkerReplicas
	}
	return 0
}

// InstallAppGroupsRequest requests to install application groups on cluster
type InstallAppGroupsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// apps are applications to install into target cluster.
	AppGroups []*AppGroup `protobuf:"bytes,1,rep,name=app_groups,json=appGroups,proto3" json:"app_groups,omitempty"`
}

func (x *InstallAppGroupsRequest) Reset() {
	*x = InstallAppGroupsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_lcm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstallAppGroupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallAppGroupsRequest) ProtoMessage() {}

func (x *InstallAppGroupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_lcm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallAppGroupsRequest.ProtoReflect.Descriptor instead.
func (*InstallAppGroupsRequest) Descriptor() ([]byte, []int) {
	return file_cluster_lcm_proto_rawDescGZIP(), []int{2}
}

func (x *InstallAppGroupsRequest) GetAppGroups() []*AppGroup {
	if x != nil {
		return x.AppGroups
	}
	return nil
}

// UninstallAppGroupsRequest is a request to delete applications on the Cluster
type UninstallAppGroupsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// cluster_id is Global Unique ID
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// app_group_ids is a array of app_group_id.
	AppGroupIds []string `protobuf:"bytes,2,rep,name=app_group_ids,json=appGroupIds,proto3" json:"app_group_ids,omitempty"`
}

func (x *UninstallAppGroupsRequest) Reset() {
	*x = UninstallAppGroupsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_lcm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UninstallAppGroupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UninstallAppGroupsRequest) ProtoMessage() {}

func (x *UninstallAppGroupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_lcm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UninstallAppGroupsRequest.ProtoReflect.Descriptor instead.
func (*UninstallAppGroupsRequest) Descriptor() ([]byte, []int) {
	return file_cluster_lcm_proto_rawDescGZIP(), []int{3}
}

func (x *UninstallAppGroupsRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *UninstallAppGroupsRequest) GetAppGroupIds() []string {
	if x != nil {
		return x.AppGroupIds
	}
	return nil
}

var File_cluster_lcm_proto protoreflect.FileDescriptor

var file_cluster_lcm_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6c, 0x63, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x6b, 0x73, 0x5f, 0x70, 0x62, 0x1a, 0x0c, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x01, 0x0a, 0x14, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x73, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x73, 0x70, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27,
	0x0a, 0x04, 0x63, 0x6f, 0x6e, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74,
	0x6b, 0x73, 0x5f, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x52, 0x04, 0x63, 0x6f, 0x6e, 0x66, 0x22, 0x86, 0x01, 0x0a, 0x13, 0x53, 0x63, 0x61, 0x6c,
	0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27,
	0x0a, 0x0f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x77, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73,
	0x22, 0x4a, 0x0a, 0x17, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x41, 0x70, 0x70, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x0a, 0x61,
	0x70, 0x70, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x74, 0x6b, 0x73, 0x5f, 0x70, 0x62, 0x2e, 0x41, 0x70, 0x70, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x09, 0x61, 0x70, 0x70, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x5e, 0x0a, 0x19,
	0x55, 0x6e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x41, 0x70, 0x70, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x61, 0x70, 0x70, 0x5f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0b, 0x61, 0x70, 0x70, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x73, 0x32, 0xf9, 0x02, 0x0a,
	0x11, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x63, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x43, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x74, 0x6b, 0x73, 0x5f, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x74, 0x6b, 0x73, 0x5f, 0x70, 0x62, 0x2e, 0x49, 0x44, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0c, 0x53, 0x63, 0x61, 0x6c, 0x65,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x74, 0x6b, 0x73, 0x5f, 0x70, 0x62,
	0x2e, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x74, 0x6b, 0x73, 0x5f, 0x70, 0x62, 0x2e, 0x53, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c,
	0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x11, 0x2e, 0x74, 0x6b, 0x73, 0x5f, 0x70, 0x62, 0x2e, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x74, 0x6b, 0x73, 0x5f, 0x70, 0x62, 0x2e, 0x53, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x10,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x41, 0x70, 0x70, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x12, 0x1f, 0x2e, 0x74, 0x6b, 0x73, 0x5f, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x41, 0x70, 0x70, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x74, 0x6b, 0x73, 0x5f, 0x70, 0x62, 0x2e, 0x49, 0x44, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x12, 0x55, 0x6e, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x41, 0x70, 0x70, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x21,
	0x2e, 0x74, 0x6b, 0x73, 0x5f, 0x70, 0x62, 0x2e, 0x55, 0x6e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x41, 0x70, 0x70, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x74, 0x6b, 0x73, 0x5f, 0x70, 0x62, 0x2e, 0x49, 0x44, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x64, 0x65, 0x76, 0x2f, 0x74, 0x6b, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x6b,
	0x73, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cluster_lcm_proto_rawDescOnce sync.Once
	file_cluster_lcm_proto_rawDescData = file_cluster_lcm_proto_rawDesc
)

func file_cluster_lcm_proto_rawDescGZIP() []byte {
	file_cluster_lcm_proto_rawDescOnce.Do(func() {
		file_cluster_lcm_proto_rawDescData = protoimpl.X.CompressGZIP(file_cluster_lcm_proto_rawDescData)
	})
	return file_cluster_lcm_proto_rawDescData
}

var file_cluster_lcm_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cluster_lcm_proto_goTypes = []interface{}{
	(*CreateClusterRequest)(nil),      // 0: tks_pb.CreateClusterRequest
	(*ScaleClusterRequest)(nil),       // 1: tks_pb.ScaleClusterRequest
	(*InstallAppGroupsRequest)(nil),   // 2: tks_pb.InstallAppGroupsRequest
	(*UninstallAppGroupsRequest)(nil), // 3: tks_pb.UninstallAppGroupsRequest
	(*ClusterConf)(nil),               // 4: tks_pb.ClusterConf
	(*AppGroup)(nil),                  // 5: tks_pb.AppGroup
	(*IDRequest)(nil),                 // 6: tks_pb.IDRequest
	(*IDResponse)(nil),                // 7: tks_pb.IDResponse
	(*SimpleResponse)(nil),            // 8: tks_pb.SimpleResponse
	(*IDsResponse)(nil),               // 9: tks_pb.IDsResponse
}
var file_cluster_lcm_proto_depIdxs = []int32{
	4, // 0: tks_pb.CreateClusterRequest.conf:type_name -> tks_pb.ClusterConf
	5, // 1: tks_pb.InstallAppGroupsRequest.app_groups:type_name -> tks_pb.AppGroup
	0, // 2: tks_pb.ClusterLcmService.CreateCluster:input_type -> tks_pb.CreateClusterRequest
	1, // 3: tks_pb.ClusterLcmService.ScaleCluster:input_type -> tks_pb.ScaleClusterRequest
	6, // 4: tks_pb.ClusterLcmService.DeleteCluster:input_type -> tks_pb.IDRequest
	2, // 5: tks_pb.ClusterLcmService.InstallAppGroups:input_type -> tks_pb.InstallAppGroupsRequest
	3, // 6: tks_pb.ClusterLcmService.UninstallAppGroups:input_type -> tks_pb.UninstallAppGroupsRequest
	7, // 7: tks_pb.ClusterLcmService.CreateCluster:output_type -> tks_pb.IDResponse
	8, // 8: tks_pb.ClusterLcmService.ScaleCluster:output_type -> tks_pb.SimpleResponse
	8, // 9: tks_pb.ClusterLcmService.DeleteCluster:output_type -> tks_pb.SimpleResponse
	9, // 10: tks_pb.ClusterLcmService.InstallAppGroups:output_type -> tks_pb.IDsResponse
	9, // 11: tks_pb.ClusterLcmService.UninstallAppGroups:output_type -> tks_pb.IDsResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cluster_lcm_proto_init() }
func file_cluster_lcm_proto_init() {
	if File_cluster_lcm_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cluster_lcm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClusterRequest); i {
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
		file_cluster_lcm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScaleClusterRequest); i {
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
		file_cluster_lcm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstallAppGroupsRequest); i {
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
		file_cluster_lcm_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UninstallAppGroupsRequest); i {
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
			RawDescriptor: file_cluster_lcm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cluster_lcm_proto_goTypes,
		DependencyIndexes: file_cluster_lcm_proto_depIdxs,
		MessageInfos:      file_cluster_lcm_proto_msgTypes,
	}.Build()
	File_cluster_lcm_proto = out.File
	file_cluster_lcm_proto_rawDesc = nil
	file_cluster_lcm_proto_goTypes = nil
	file_cluster_lcm_proto_depIdxs = nil
}
