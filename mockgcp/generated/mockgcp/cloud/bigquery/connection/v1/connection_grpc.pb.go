// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: mockgcp/cloud/bigquery/connection/v1/connection.proto

package connectionpb

import (
	iampb "cloud.google.com/go/iam/apiv1/iampb"
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ConnectionServiceClient is the client API for ConnectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectionServiceClient interface {
	// Creates a new connection.
	CreateConnection(ctx context.Context, in *CreateConnectionRequest, opts ...grpc.CallOption) (*Connection, error)
	// Returns specified connection.
	GetConnection(ctx context.Context, in *GetConnectionRequest, opts ...grpc.CallOption) (*Connection, error)
	// Returns a list of connections in the given project.
	ListConnections(ctx context.Context, in *ListConnectionsRequest, opts ...grpc.CallOption) (*ListConnectionsResponse, error)
	// Updates the specified connection. For security reasons, also resets
	// credential if connection properties are in the update field mask.
	UpdateConnection(ctx context.Context, in *UpdateConnectionRequest, opts ...grpc.CallOption) (*Connection, error)
	// Deletes connection and associated credential.
	DeleteConnection(ctx context.Context, in *DeleteConnectionRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Gets the access control policy for a resource.
	// Returns an empty policy if the resource exists and does not have a policy
	// set.
	GetIamPolicy(ctx context.Context, in *iampb.GetIamPolicyRequest, opts ...grpc.CallOption) (*iampb.Policy, error)
	// Sets the access control policy on the specified resource. Replaces any
	// existing policy.
	//
	// Can return `NOT_FOUND`, `INVALID_ARGUMENT`, and `PERMISSION_DENIED` errors.
	SetIamPolicy(ctx context.Context, in *iampb.SetIamPolicyRequest, opts ...grpc.CallOption) (*iampb.Policy, error)
	// Returns permissions that a caller has on the specified resource.
	// If the resource does not exist, this will return an empty set of
	// permissions, not a `NOT_FOUND` error.
	//
	// Note: This operation is designed to be used for building permission-aware
	// UIs and command-line tools, not for authorization checking. This operation
	// may "fail open" without warning.
	TestIamPermissions(ctx context.Context, in *iampb.TestIamPermissionsRequest, opts ...grpc.CallOption) (*iampb.TestIamPermissionsResponse, error)
}

type connectionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectionServiceClient(cc grpc.ClientConnInterface) ConnectionServiceClient {
	return &connectionServiceClient{cc}
}

func (c *connectionServiceClient) CreateConnection(ctx context.Context, in *CreateConnectionRequest, opts ...grpc.CallOption) (*Connection, error) {
	out := new(Connection)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.bigquery.connection.v1.ConnectionService/CreateConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) GetConnection(ctx context.Context, in *GetConnectionRequest, opts ...grpc.CallOption) (*Connection, error) {
	out := new(Connection)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.bigquery.connection.v1.ConnectionService/GetConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) ListConnections(ctx context.Context, in *ListConnectionsRequest, opts ...grpc.CallOption) (*ListConnectionsResponse, error) {
	out := new(ListConnectionsResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.bigquery.connection.v1.ConnectionService/ListConnections", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) UpdateConnection(ctx context.Context, in *UpdateConnectionRequest, opts ...grpc.CallOption) (*Connection, error) {
	out := new(Connection)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.bigquery.connection.v1.ConnectionService/UpdateConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) DeleteConnection(ctx context.Context, in *DeleteConnectionRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.bigquery.connection.v1.ConnectionService/DeleteConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) GetIamPolicy(ctx context.Context, in *iampb.GetIamPolicyRequest, opts ...grpc.CallOption) (*iampb.Policy, error) {
	out := new(iampb.Policy)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.bigquery.connection.v1.ConnectionService/GetIamPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) SetIamPolicy(ctx context.Context, in *iampb.SetIamPolicyRequest, opts ...grpc.CallOption) (*iampb.Policy, error) {
	out := new(iampb.Policy)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.bigquery.connection.v1.ConnectionService/SetIamPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) TestIamPermissions(ctx context.Context, in *iampb.TestIamPermissionsRequest, opts ...grpc.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	out := new(iampb.TestIamPermissionsResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.bigquery.connection.v1.ConnectionService/TestIamPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectionServiceServer is the server API for ConnectionService service.
// All implementations must embed UnimplementedConnectionServiceServer
// for forward compatibility
type ConnectionServiceServer interface {
	// Creates a new connection.
	CreateConnection(context.Context, *CreateConnectionRequest) (*Connection, error)
	// Returns specified connection.
	GetConnection(context.Context, *GetConnectionRequest) (*Connection, error)
	// Returns a list of connections in the given project.
	ListConnections(context.Context, *ListConnectionsRequest) (*ListConnectionsResponse, error)
	// Updates the specified connection. For security reasons, also resets
	// credential if connection properties are in the update field mask.
	UpdateConnection(context.Context, *UpdateConnectionRequest) (*Connection, error)
	// Deletes connection and associated credential.
	DeleteConnection(context.Context, *DeleteConnectionRequest) (*empty.Empty, error)
	// Gets the access control policy for a resource.
	// Returns an empty policy if the resource exists and does not have a policy
	// set.
	GetIamPolicy(context.Context, *iampb.GetIamPolicyRequest) (*iampb.Policy, error)
	// Sets the access control policy on the specified resource. Replaces any
	// existing policy.
	//
	// Can return `NOT_FOUND`, `INVALID_ARGUMENT`, and `PERMISSION_DENIED` errors.
	SetIamPolicy(context.Context, *iampb.SetIamPolicyRequest) (*iampb.Policy, error)
	// Returns permissions that a caller has on the specified resource.
	// If the resource does not exist, this will return an empty set of
	// permissions, not a `NOT_FOUND` error.
	//
	// Note: This operation is designed to be used for building permission-aware
	// UIs and command-line tools, not for authorization checking. This operation
	// may "fail open" without warning.
	TestIamPermissions(context.Context, *iampb.TestIamPermissionsRequest) (*iampb.TestIamPermissionsResponse, error)
	mustEmbedUnimplementedConnectionServiceServer()
}

// UnimplementedConnectionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConnectionServiceServer struct {
}

func (UnimplementedConnectionServiceServer) CreateConnection(context.Context, *CreateConnectionRequest) (*Connection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConnection not implemented")
}
func (UnimplementedConnectionServiceServer) GetConnection(context.Context, *GetConnectionRequest) (*Connection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnection not implemented")
}
func (UnimplementedConnectionServiceServer) ListConnections(context.Context, *ListConnectionsRequest) (*ListConnectionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListConnections not implemented")
}
func (UnimplementedConnectionServiceServer) UpdateConnection(context.Context, *UpdateConnectionRequest) (*Connection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConnection not implemented")
}
func (UnimplementedConnectionServiceServer) DeleteConnection(context.Context, *DeleteConnectionRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConnection not implemented")
}
func (UnimplementedConnectionServiceServer) GetIamPolicy(context.Context, *iampb.GetIamPolicyRequest) (*iampb.Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIamPolicy not implemented")
}
func (UnimplementedConnectionServiceServer) SetIamPolicy(context.Context, *iampb.SetIamPolicyRequest) (*iampb.Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetIamPolicy not implemented")
}
func (UnimplementedConnectionServiceServer) TestIamPermissions(context.Context, *iampb.TestIamPermissionsRequest) (*iampb.TestIamPermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestIamPermissions not implemented")
}
func (UnimplementedConnectionServiceServer) mustEmbedUnimplementedConnectionServiceServer() {}

// UnsafeConnectionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectionServiceServer will
// result in compilation errors.
type UnsafeConnectionServiceServer interface {
	mustEmbedUnimplementedConnectionServiceServer()
}

func RegisterConnectionServiceServer(s grpc.ServiceRegistrar, srv ConnectionServiceServer) {
	s.RegisterService(&ConnectionService_ServiceDesc, srv)
}

func _ConnectionService_CreateConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).CreateConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.bigquery.connection.v1.ConnectionService/CreateConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).CreateConnection(ctx, req.(*CreateConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_GetConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).GetConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.bigquery.connection.v1.ConnectionService/GetConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).GetConnection(ctx, req.(*GetConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_ListConnections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListConnectionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).ListConnections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.bigquery.connection.v1.ConnectionService/ListConnections",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).ListConnections(ctx, req.(*ListConnectionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_UpdateConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).UpdateConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.bigquery.connection.v1.ConnectionService/UpdateConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).UpdateConnection(ctx, req.(*UpdateConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_DeleteConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).DeleteConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.bigquery.connection.v1.ConnectionService/DeleteConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).DeleteConnection(ctx, req.(*DeleteConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_GetIamPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(iampb.GetIamPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).GetIamPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.bigquery.connection.v1.ConnectionService/GetIamPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).GetIamPolicy(ctx, req.(*iampb.GetIamPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_SetIamPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(iampb.SetIamPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).SetIamPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.bigquery.connection.v1.ConnectionService/SetIamPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).SetIamPolicy(ctx, req.(*iampb.SetIamPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_TestIamPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(iampb.TestIamPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).TestIamPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.bigquery.connection.v1.ConnectionService/TestIamPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).TestIamPermissions(ctx, req.(*iampb.TestIamPermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConnectionService_ServiceDesc is the grpc.ServiceDesc for ConnectionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConnectionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mockgcp.cloud.bigquery.connection.v1.ConnectionService",
	HandlerType: (*ConnectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateConnection",
			Handler:    _ConnectionService_CreateConnection_Handler,
		},
		{
			MethodName: "GetConnection",
			Handler:    _ConnectionService_GetConnection_Handler,
		},
		{
			MethodName: "ListConnections",
			Handler:    _ConnectionService_ListConnections_Handler,
		},
		{
			MethodName: "UpdateConnection",
			Handler:    _ConnectionService_UpdateConnection_Handler,
		},
		{
			MethodName: "DeleteConnection",
			Handler:    _ConnectionService_DeleteConnection_Handler,
		},
		{
			MethodName: "GetIamPolicy",
			Handler:    _ConnectionService_GetIamPolicy_Handler,
		},
		{
			MethodName: "SetIamPolicy",
			Handler:    _ConnectionService_SetIamPolicy_Handler,
		},
		{
			MethodName: "TestIamPermissions",
			Handler:    _ConnectionService_TestIamPermissions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mockgcp/cloud/bigquery/connection/v1/connection.proto",
}