// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.1
// source: pb/yun.proto

package pb

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

const (
	Storage_HeartBeat_FullMethodName = "/pb.Storage/HeartBeat"
	Storage_Upload_FullMethodName    = "/pb.Storage/Upload"
	Storage_Manage_FullMethodName    = "/pb.Storage/Manage"
)

// StorageClient is the client API for Storage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageClient interface {
	// Sends a greeting
	HeartBeat(ctx context.Context, in *HeartBeatRequest, opts ...grpc.CallOption) (*HeartBeatReply, error)
	Upload(ctx context.Context, opts ...grpc.CallOption) (Storage_UploadClient, error)
	Manage(ctx context.Context, in *ManageRequest, opts ...grpc.CallOption) (*ManageReply, error)
}

type storageClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageClient(cc grpc.ClientConnInterface) StorageClient {
	return &storageClient{cc}
}

func (c *storageClient) HeartBeat(ctx context.Context, in *HeartBeatRequest, opts ...grpc.CallOption) (*HeartBeatReply, error) {
	out := new(HeartBeatReply)
	err := c.cc.Invoke(ctx, Storage_HeartBeat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Upload(ctx context.Context, opts ...grpc.CallOption) (Storage_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &Storage_ServiceDesc.Streams[0], Storage_Upload_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &storageUploadClient{stream}
	return x, nil
}

type Storage_UploadClient interface {
	Send(*UploadRequest) error
	CloseAndRecv() (*UploadReply, error)
	grpc.ClientStream
}

type storageUploadClient struct {
	grpc.ClientStream
}

func (x *storageUploadClient) Send(m *UploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *storageUploadClient) CloseAndRecv() (*UploadReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *storageClient) Manage(ctx context.Context, in *ManageRequest, opts ...grpc.CallOption) (*ManageReply, error) {
	out := new(ManageReply)
	err := c.cc.Invoke(ctx, Storage_Manage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageServer is the server API for Storage service.
// All implementations must embed UnimplementedStorageServer
// for forward compatibility
type StorageServer interface {
	// Sends a greeting
	HeartBeat(context.Context, *HeartBeatRequest) (*HeartBeatReply, error)
	Upload(Storage_UploadServer) error
	Manage(context.Context, *ManageRequest) (*ManageReply, error)
	mustEmbedUnimplementedStorageServer()
}

// UnimplementedStorageServer must be embedded to have forward compatible implementations.
type UnimplementedStorageServer struct {
}

func (UnimplementedStorageServer) HeartBeat(context.Context, *HeartBeatRequest) (*HeartBeatReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeartBeat not implemented")
}
func (UnimplementedStorageServer) Upload(Storage_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedStorageServer) Manage(context.Context, *ManageRequest) (*ManageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Manage not implemented")
}
func (UnimplementedStorageServer) mustEmbedUnimplementedStorageServer() {}

// UnsafeStorageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageServer will
// result in compilation errors.
type UnsafeStorageServer interface {
	mustEmbedUnimplementedStorageServer()
}

func RegisterStorageServer(s grpc.ServiceRegistrar, srv StorageServer) {
	s.RegisterService(&Storage_ServiceDesc, srv)
}

func _Storage_HeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartBeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).HeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_HeartBeat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).HeartBeat(ctx, req.(*HeartBeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StorageServer).Upload(&storageUploadServer{stream})
}

type Storage_UploadServer interface {
	SendAndClose(*UploadReply) error
	Recv() (*UploadRequest, error)
	grpc.ServerStream
}

type storageUploadServer struct {
	grpc.ServerStream
}

func (x *storageUploadServer) SendAndClose(m *UploadReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *storageUploadServer) Recv() (*UploadRequest, error) {
	m := new(UploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Storage_Manage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Manage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_Manage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Manage(ctx, req.(*ManageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Storage_ServiceDesc is the grpc.ServiceDesc for Storage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Storage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Storage",
	HandlerType: (*StorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HeartBeat",
			Handler:    _Storage_HeartBeat_Handler,
		},
		{
			MethodName: "Manage",
			Handler:    _Storage_Manage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _Storage_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "pb/yun.proto",
}
