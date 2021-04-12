// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pbgo

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// InfoServiceClient is the client API for InfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InfoServiceClient interface {
	// CreateMC create Multi cluster and return the id for the mc
	CreateMC(ctx context.Context, in *CreateMCRequest, opts ...grpc.CallOption) (*IDResponse, error)
	// GetMCIDs get all multicluster ids for the request
	GetMCIDs(ctx context.Context, in *GetIDsRequest, opts ...grpc.CallOption) (*IDsResponse, error)
	// GetMC get the multicluster for the requsted id.
	GetMC(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*GetMCResponse, error)
	// GetMCs get every multicusters
	GetMCs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetMCsResponse, error)
	// CreateCluster create cluster on the multicuster with tenent id
	CreateCluster(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*IDResponse, error)
	// GetClusterget cluster for the id of the cluster
	GetCluster(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*GetClusterResponse, error)
	// GetClusters get every clusters on the mutlcluster
	GetClusters(ctx context.Context, in *GetClustersRequest, opts ...grpc.CallOption) (*GetClustersResponse, error)
	// UpdateClusterStatus update Status of the Cluster
	UpdateClusterStatus(ctx context.Context, in *UpdateClusterStatusRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
	// ValidateLabelUniqueness check uniqueness of the label
	ValidateLabelUniqueness(ctx context.Context, in *ValidateLabelUniquenessRequest, opts ...grpc.CallOption) (*ValidateLabelUniquenessResponse, error)
	// GetAppID get an array of app IDs with knowlege
	GetAppID(ctx context.Context, in *GetAppRequest, opts ...grpc.CallOption) (*IDsResponse, error)
	// GetAllApps get app info in the cluster with the ID
	GetAllApps(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*GetAppsResponse, error)
	// GetApps get app info for the given contdition
	GetApps(ctx context.Context, in *GetAppsRequest, opts ...grpc.CallOption) (*GetAppResponse, error)
	// UpdateAppStatus update Status of the application
	UpdateAppStatus(ctx context.Context, in *UpdateAppStatusRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
	// UpdateEndpoints update endpoints by the application
	UpdateEndpoints(ctx context.Context, in *UpdateEndpointsRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
}

type infoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInfoServiceClient(cc grpc.ClientConnInterface) InfoServiceClient {
	return &infoServiceClient{cc}
}

func (c *infoServiceClient) CreateMC(ctx context.Context, in *CreateMCRequest, opts ...grpc.CallOption) (*IDResponse, error) {
	out := new(IDResponse)
	err := c.cc.Invoke(ctx, "/pbgo.InfoService/CreateMC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoServiceClient) GetMCIDs(ctx context.Context, in *GetIDsRequest, opts ...grpc.CallOption) (*IDsResponse, error) {
	out := new(IDsResponse)
	err := c.cc.Invoke(ctx, "/pbgo.InfoService/GetMCIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoServiceClient) GetMC(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*GetMCResponse, error) {
	out := new(GetMCResponse)
	err := c.cc.Invoke(ctx, "/pbgo.InfoService/GetMC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoServiceClient) GetMCs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetMCsResponse, error) {
	out := new(GetMCsResponse)
	err := c.cc.Invoke(ctx, "/pbgo.InfoService/GetMCs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoServiceClient) CreateCluster(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*IDResponse, error) {
	out := new(IDResponse)
	err := c.cc.Invoke(ctx, "/pbgo.InfoService/CreateCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoServiceClient) GetCluster(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*GetClusterResponse, error) {
	out := new(GetClusterResponse)
	err := c.cc.Invoke(ctx, "/pbgo.InfoService/GetCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoServiceClient) GetClusters(ctx context.Context, in *GetClustersRequest, opts ...grpc.CallOption) (*GetClustersResponse, error) {
	out := new(GetClustersResponse)
	err := c.cc.Invoke(ctx, "/pbgo.InfoService/GetClusters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoServiceClient) UpdateClusterStatus(ctx context.Context, in *UpdateClusterStatusRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := c.cc.Invoke(ctx, "/pbgo.InfoService/UpdateClusterStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoServiceClient) ValidateLabelUniqueness(ctx context.Context, in *ValidateLabelUniquenessRequest, opts ...grpc.CallOption) (*ValidateLabelUniquenessResponse, error) {
	out := new(ValidateLabelUniquenessResponse)
	err := c.cc.Invoke(ctx, "/pbgo.InfoService/ValidateLabelUniqueness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoServiceClient) GetAppID(ctx context.Context, in *GetAppRequest, opts ...grpc.CallOption) (*IDsResponse, error) {
	out := new(IDsResponse)
	err := c.cc.Invoke(ctx, "/pbgo.InfoService/GetAppID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoServiceClient) GetAllApps(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*GetAppsResponse, error) {
	out := new(GetAppsResponse)
	err := c.cc.Invoke(ctx, "/pbgo.InfoService/GetAllApps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoServiceClient) GetApps(ctx context.Context, in *GetAppsRequest, opts ...grpc.CallOption) (*GetAppResponse, error) {
	out := new(GetAppResponse)
	err := c.cc.Invoke(ctx, "/pbgo.InfoService/GetApps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoServiceClient) UpdateAppStatus(ctx context.Context, in *UpdateAppStatusRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := c.cc.Invoke(ctx, "/pbgo.InfoService/UpdateAppStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoServiceClient) UpdateEndpoints(ctx context.Context, in *UpdateEndpointsRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := c.cc.Invoke(ctx, "/pbgo.InfoService/UpdateEndpoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InfoServiceServer is the server API for InfoService service.
// All implementations must embed UnimplementedInfoServiceServer
// for forward compatibility
type InfoServiceServer interface {
	// CreateMC create Multi cluster and return the id for the mc
	CreateMC(context.Context, *CreateMCRequest) (*IDResponse, error)
	// GetMCIDs get all multicluster ids for the request
	GetMCIDs(context.Context, *GetIDsRequest) (*IDsResponse, error)
	// GetMC get the multicluster for the requsted id.
	GetMC(context.Context, *IDRequest) (*GetMCResponse, error)
	// GetMCs get every multicusters
	GetMCs(context.Context, *emptypb.Empty) (*GetMCsResponse, error)
	// CreateCluster create cluster on the multicuster with tenent id
	CreateCluster(context.Context, *CreateClusterRequest) (*IDResponse, error)
	// GetClusterget cluster for the id of the cluster
	GetCluster(context.Context, *GetClusterRequest) (*GetClusterResponse, error)
	// GetClusters get every clusters on the mutlcluster
	GetClusters(context.Context, *GetClustersRequest) (*GetClustersResponse, error)
	// UpdateClusterStatus update Status of the Cluster
	UpdateClusterStatus(context.Context, *UpdateClusterStatusRequest) (*SimpleResponse, error)
	// ValidateLabelUniqueness check uniqueness of the label
	ValidateLabelUniqueness(context.Context, *ValidateLabelUniquenessRequest) (*ValidateLabelUniquenessResponse, error)
	// GetAppID get an array of app IDs with knowlege
	GetAppID(context.Context, *GetAppRequest) (*IDsResponse, error)
	// GetAllApps get app info in the cluster with the ID
	GetAllApps(context.Context, *IDRequest) (*GetAppsResponse, error)
	// GetApps get app info for the given contdition
	GetApps(context.Context, *GetAppsRequest) (*GetAppResponse, error)
	// UpdateAppStatus update Status of the application
	UpdateAppStatus(context.Context, *UpdateAppStatusRequest) (*SimpleResponse, error)
	// UpdateEndpoints update endpoints by the application
	UpdateEndpoints(context.Context, *UpdateEndpointsRequest) (*SimpleResponse, error)
	mustEmbedUnimplementedInfoServiceServer()
}

// UnimplementedInfoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInfoServiceServer struct {
}

func (UnimplementedInfoServiceServer) CreateMC(context.Context, *CreateMCRequest) (*IDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMC not implemented")
}
func (UnimplementedInfoServiceServer) GetMCIDs(context.Context, *GetIDsRequest) (*IDsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMCIDs not implemented")
}
func (UnimplementedInfoServiceServer) GetMC(context.Context, *IDRequest) (*GetMCResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMC not implemented")
}
func (UnimplementedInfoServiceServer) GetMCs(context.Context, *emptypb.Empty) (*GetMCsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMCs not implemented")
}
func (UnimplementedInfoServiceServer) CreateCluster(context.Context, *CreateClusterRequest) (*IDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCluster not implemented")
}
func (UnimplementedInfoServiceServer) GetCluster(context.Context, *GetClusterRequest) (*GetClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCluster not implemented")
}
func (UnimplementedInfoServiceServer) GetClusters(context.Context, *GetClustersRequest) (*GetClustersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusters not implemented")
}
func (UnimplementedInfoServiceServer) UpdateClusterStatus(context.Context, *UpdateClusterStatusRequest) (*SimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClusterStatus not implemented")
}
func (UnimplementedInfoServiceServer) ValidateLabelUniqueness(context.Context, *ValidateLabelUniquenessRequest) (*ValidateLabelUniquenessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateLabelUniqueness not implemented")
}
func (UnimplementedInfoServiceServer) GetAppID(context.Context, *GetAppRequest) (*IDsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppID not implemented")
}
func (UnimplementedInfoServiceServer) GetAllApps(context.Context, *IDRequest) (*GetAppsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllApps not implemented")
}
func (UnimplementedInfoServiceServer) GetApps(context.Context, *GetAppsRequest) (*GetAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApps not implemented")
}
func (UnimplementedInfoServiceServer) UpdateAppStatus(context.Context, *UpdateAppStatusRequest) (*SimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppStatus not implemented")
}
func (UnimplementedInfoServiceServer) UpdateEndpoints(context.Context, *UpdateEndpointsRequest) (*SimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEndpoints not implemented")
}
func (UnimplementedInfoServiceServer) mustEmbedUnimplementedInfoServiceServer() {}

// UnsafeInfoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InfoServiceServer will
// result in compilation errors.
type UnsafeInfoServiceServer interface {
	mustEmbedUnimplementedInfoServiceServer()
}

func RegisterInfoServiceServer(s grpc.ServiceRegistrar, srv InfoServiceServer) {
	s.RegisterService(&InfoService_ServiceDesc, srv)
}

func _InfoService_CreateMC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServiceServer).CreateMC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbgo.InfoService/CreateMC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServiceServer).CreateMC(ctx, req.(*CreateMCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfoService_GetMCIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServiceServer).GetMCIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbgo.InfoService/GetMCIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServiceServer).GetMCIDs(ctx, req.(*GetIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfoService_GetMC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServiceServer).GetMC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbgo.InfoService/GetMC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServiceServer).GetMC(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfoService_GetMCs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServiceServer).GetMCs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbgo.InfoService/GetMCs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServiceServer).GetMCs(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfoService_CreateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServiceServer).CreateCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbgo.InfoService/CreateCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServiceServer).CreateCluster(ctx, req.(*CreateClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfoService_GetCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServiceServer).GetCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbgo.InfoService/GetCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServiceServer).GetCluster(ctx, req.(*GetClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfoService_GetClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServiceServer).GetClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbgo.InfoService/GetClusters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServiceServer).GetClusters(ctx, req.(*GetClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfoService_UpdateClusterStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClusterStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServiceServer).UpdateClusterStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbgo.InfoService/UpdateClusterStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServiceServer).UpdateClusterStatus(ctx, req.(*UpdateClusterStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfoService_ValidateLabelUniqueness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateLabelUniquenessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServiceServer).ValidateLabelUniqueness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbgo.InfoService/ValidateLabelUniqueness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServiceServer).ValidateLabelUniqueness(ctx, req.(*ValidateLabelUniquenessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfoService_GetAppID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServiceServer).GetAppID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbgo.InfoService/GetAppID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServiceServer).GetAppID(ctx, req.(*GetAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfoService_GetAllApps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServiceServer).GetAllApps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbgo.InfoService/GetAllApps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServiceServer).GetAllApps(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfoService_GetApps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServiceServer).GetApps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbgo.InfoService/GetApps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServiceServer).GetApps(ctx, req.(*GetAppsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfoService_UpdateAppStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServiceServer).UpdateAppStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbgo.InfoService/UpdateAppStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServiceServer).UpdateAppStatus(ctx, req.(*UpdateAppStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfoService_UpdateEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEndpointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServiceServer).UpdateEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbgo.InfoService/UpdateEndpoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServiceServer).UpdateEndpoints(ctx, req.(*UpdateEndpointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InfoService_ServiceDesc is the grpc.ServiceDesc for InfoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InfoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pbgo.InfoService",
	HandlerType: (*InfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMC",
			Handler:    _InfoService_CreateMC_Handler,
		},
		{
			MethodName: "GetMCIDs",
			Handler:    _InfoService_GetMCIDs_Handler,
		},
		{
			MethodName: "GetMC",
			Handler:    _InfoService_GetMC_Handler,
		},
		{
			MethodName: "GetMCs",
			Handler:    _InfoService_GetMCs_Handler,
		},
		{
			MethodName: "CreateCluster",
			Handler:    _InfoService_CreateCluster_Handler,
		},
		{
			MethodName: "GetCluster",
			Handler:    _InfoService_GetCluster_Handler,
		},
		{
			MethodName: "GetClusters",
			Handler:    _InfoService_GetClusters_Handler,
		},
		{
			MethodName: "UpdateClusterStatus",
			Handler:    _InfoService_UpdateClusterStatus_Handler,
		},
		{
			MethodName: "ValidateLabelUniqueness",
			Handler:    _InfoService_ValidateLabelUniqueness_Handler,
		},
		{
			MethodName: "GetAppID",
			Handler:    _InfoService_GetAppID_Handler,
		},
		{
			MethodName: "GetAllApps",
			Handler:    _InfoService_GetAllApps_Handler,
		},
		{
			MethodName: "GetApps",
			Handler:    _InfoService_GetApps_Handler,
		},
		{
			MethodName: "UpdateAppStatus",
			Handler:    _InfoService_UpdateAppStatus_Handler,
		},
		{
			MethodName: "UpdateEndpoints",
			Handler:    _InfoService_UpdateEndpoints_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "info.proto",
}
