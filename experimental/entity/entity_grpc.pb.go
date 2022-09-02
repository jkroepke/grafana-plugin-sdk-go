// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: entity.proto

package entity

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EntityStoreClient is the client API for EntityStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EntityStoreClient interface {
	GetEntity(ctx context.Context, in *GetEntityRequest, opts ...grpc.CallOption) (*EntityMessage, error)
	ListFolder(ctx context.Context, in *ListFolderRequest, opts ...grpc.CallOption) (*FolderListing, error)
	SaveEntity(ctx context.Context, in *SaveEntityRequest, opts ...grpc.CallOption) (*EntityMessage, error)
	DeleteEntity(ctx context.Context, in *DeleteEntityRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	GetEntityHistory(ctx context.Context, in *GetHistoryRequest, opts ...grpc.CallOption) (*EntityHistoryResponse, error)
	CreatePR(ctx context.Context, in *CreatePullRequest, opts ...grpc.CallOption) (*EntityMessage, error)
	ListKinds(ctx context.Context, in *ListKindsRequest, opts ...grpc.CallOption) (*ListKindsResponse, error)
	// Later...
	WatchEntity(ctx context.Context, in *GetEntityRequest, opts ...grpc.CallOption) (EntityStore_WatchEntityClient, error)
}

type entityStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewEntityStoreClient(cc grpc.ClientConnInterface) EntityStoreClient {
	return &entityStoreClient{cc}
}

func (c *entityStoreClient) GetEntity(ctx context.Context, in *GetEntityRequest, opts ...grpc.CallOption) (*EntityMessage, error) {
	out := new(EntityMessage)
	err := c.cc.Invoke(ctx, "/entity.EntityStore/GetEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityStoreClient) ListFolder(ctx context.Context, in *ListFolderRequest, opts ...grpc.CallOption) (*FolderListing, error) {
	out := new(FolderListing)
	err := c.cc.Invoke(ctx, "/entity.EntityStore/ListFolder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityStoreClient) SaveEntity(ctx context.Context, in *SaveEntityRequest, opts ...grpc.CallOption) (*EntityMessage, error) {
	out := new(EntityMessage)
	err := c.cc.Invoke(ctx, "/entity.EntityStore/SaveEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityStoreClient) DeleteEntity(ctx context.Context, in *DeleteEntityRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/entity.EntityStore/DeleteEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityStoreClient) GetEntityHistory(ctx context.Context, in *GetHistoryRequest, opts ...grpc.CallOption) (*EntityHistoryResponse, error) {
	out := new(EntityHistoryResponse)
	err := c.cc.Invoke(ctx, "/entity.EntityStore/GetEntityHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityStoreClient) CreatePR(ctx context.Context, in *CreatePullRequest, opts ...grpc.CallOption) (*EntityMessage, error) {
	out := new(EntityMessage)
	err := c.cc.Invoke(ctx, "/entity.EntityStore/CreatePR", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityStoreClient) ListKinds(ctx context.Context, in *ListKindsRequest, opts ...grpc.CallOption) (*ListKindsResponse, error) {
	out := new(ListKindsResponse)
	err := c.cc.Invoke(ctx, "/entity.EntityStore/ListKinds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityStoreClient) WatchEntity(ctx context.Context, in *GetEntityRequest, opts ...grpc.CallOption) (EntityStore_WatchEntityClient, error) {
	stream, err := c.cc.NewStream(ctx, &EntityStore_ServiceDesc.Streams[0], "/entity.EntityStore/WatchEntity", opts...)
	if err != nil {
		return nil, err
	}
	x := &entityStoreWatchEntityClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EntityStore_WatchEntityClient interface {
	Recv() (*EntityMessage, error)
	grpc.ClientStream
}

type entityStoreWatchEntityClient struct {
	grpc.ClientStream
}

func (x *entityStoreWatchEntityClient) Recv() (*EntityMessage, error) {
	m := new(EntityMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EntityStoreServer is the server API for EntityStore service.
// All implementations should embed UnimplementedEntityStoreServer
// for forward compatibility
type EntityStoreServer interface {
	GetEntity(context.Context, *GetEntityRequest) (*EntityMessage, error)
	ListFolder(context.Context, *ListFolderRequest) (*FolderListing, error)
	SaveEntity(context.Context, *SaveEntityRequest) (*EntityMessage, error)
	DeleteEntity(context.Context, *DeleteEntityRequest) (*DeleteResponse, error)
	GetEntityHistory(context.Context, *GetHistoryRequest) (*EntityHistoryResponse, error)
	CreatePR(context.Context, *CreatePullRequest) (*EntityMessage, error)
	ListKinds(context.Context, *ListKindsRequest) (*ListKindsResponse, error)
	// Later...
	WatchEntity(*GetEntityRequest, EntityStore_WatchEntityServer) error
}

// UnimplementedEntityStoreServer should be embedded to have forward compatible implementations.
type UnimplementedEntityStoreServer struct {
}

func (UnimplementedEntityStoreServer) GetEntity(context.Context, *GetEntityRequest) (*EntityMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEntity not implemented")
}
func (UnimplementedEntityStoreServer) ListFolder(context.Context, *ListFolderRequest) (*FolderListing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFolder not implemented")
}
func (UnimplementedEntityStoreServer) SaveEntity(context.Context, *SaveEntityRequest) (*EntityMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveEntity not implemented")
}
func (UnimplementedEntityStoreServer) DeleteEntity(context.Context, *DeleteEntityRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEntity not implemented")
}
func (UnimplementedEntityStoreServer) GetEntityHistory(context.Context, *GetHistoryRequest) (*EntityHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEntityHistory not implemented")
}
func (UnimplementedEntityStoreServer) CreatePR(context.Context, *CreatePullRequest) (*EntityMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePR not implemented")
}
func (UnimplementedEntityStoreServer) ListKinds(context.Context, *ListKindsRequest) (*ListKindsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKinds not implemented")
}
func (UnimplementedEntityStoreServer) WatchEntity(*GetEntityRequest, EntityStore_WatchEntityServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchEntity not implemented")
}

// UnsafeEntityStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EntityStoreServer will
// result in compilation errors.
type UnsafeEntityStoreServer interface {
	mustEmbedUnimplementedEntityStoreServer()
}

func RegisterEntityStoreServer(s grpc.ServiceRegistrar, srv EntityStoreServer) {
	s.RegisterService(&EntityStore_ServiceDesc, srv)
}

func _EntityStore_GetEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityStoreServer).GetEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.EntityStore/GetEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityStoreServer).GetEntity(ctx, req.(*GetEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntityStore_ListFolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityStoreServer).ListFolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.EntityStore/ListFolder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityStoreServer).ListFolder(ctx, req.(*ListFolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntityStore_SaveEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityStoreServer).SaveEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.EntityStore/SaveEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityStoreServer).SaveEntity(ctx, req.(*SaveEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntityStore_DeleteEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityStoreServer).DeleteEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.EntityStore/DeleteEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityStoreServer).DeleteEntity(ctx, req.(*DeleteEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntityStore_GetEntityHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityStoreServer).GetEntityHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.EntityStore/GetEntityHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityStoreServer).GetEntityHistory(ctx, req.(*GetHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntityStore_CreatePR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePullRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityStoreServer).CreatePR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.EntityStore/CreatePR",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityStoreServer).CreatePR(ctx, req.(*CreatePullRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntityStore_ListKinds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListKindsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityStoreServer).ListKinds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.EntityStore/ListKinds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityStoreServer).ListKinds(ctx, req.(*ListKindsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntityStore_WatchEntity_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetEntityRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EntityStoreServer).WatchEntity(m, &entityStoreWatchEntityServer{stream})
}

type EntityStore_WatchEntityServer interface {
	Send(*EntityMessage) error
	grpc.ServerStream
}

type entityStoreWatchEntityServer struct {
	grpc.ServerStream
}

func (x *entityStoreWatchEntityServer) Send(m *EntityMessage) error {
	return x.ServerStream.SendMsg(m)
}

// EntityStore_ServiceDesc is the grpc.ServiceDesc for EntityStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EntityStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "entity.EntityStore",
	HandlerType: (*EntityStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEntity",
			Handler:    _EntityStore_GetEntity_Handler,
		},
		{
			MethodName: "ListFolder",
			Handler:    _EntityStore_ListFolder_Handler,
		},
		{
			MethodName: "SaveEntity",
			Handler:    _EntityStore_SaveEntity_Handler,
		},
		{
			MethodName: "DeleteEntity",
			Handler:    _EntityStore_DeleteEntity_Handler,
		},
		{
			MethodName: "GetEntityHistory",
			Handler:    _EntityStore_GetEntityHistory_Handler,
		},
		{
			MethodName: "CreatePR",
			Handler:    _EntityStore_CreatePR_Handler,
		},
		{
			MethodName: "ListKinds",
			Handler:    _EntityStore_ListKinds_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchEntity",
			Handler:       _EntityStore_WatchEntity_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "entity.proto",
}
