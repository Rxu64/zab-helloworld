// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: zab/zab.proto

package __

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

// ZabClient is the client API for Zab service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ZabClient interface {
	// Boradcasts
	Broadcast(ctx context.Context, in *Txn, opts ...grpc.CallOption) (*Ack, error)
}

type zabClient struct {
	cc grpc.ClientConnInterface
}

func NewZabClient(cc grpc.ClientConnInterface) ZabClient {
	return &zabClient{cc}
}

func (c *zabClient) Broadcast(ctx context.Context, in *Txn, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/zab.Zab/Broadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ZabServer is the server API for Zab service.
// All implementations must embed UnimplementedZabServer
// for forward compatibility
type ZabServer interface {
	// Boradcasts
	Broadcast(context.Context, *Txn) (*Ack, error)
	mustEmbedUnimplementedZabServer()
}

// UnimplementedZabServer must be embedded to have forward compatible implementations.
type UnimplementedZabServer struct {
}

func (UnimplementedZabServer) Broadcast(context.Context, *Txn) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
func (UnimplementedZabServer) mustEmbedUnimplementedZabServer() {}

// UnsafeZabServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ZabServer will
// result in compilation errors.
type UnsafeZabServer interface {
	mustEmbedUnimplementedZabServer()
}

func RegisterZabServer(s grpc.ServiceRegistrar, srv ZabServer) {
	s.RegisterService(&Zab_ServiceDesc, srv)
}

func _Zab_Broadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Txn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZabServer).Broadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zab.Zab/Broadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZabServer).Broadcast(ctx, req.(*Txn))
	}
	return interceptor(ctx, in, info, handler)
}

// Zab_ServiceDesc is the grpc.ServiceDesc for Zab service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Zab_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "zab.Zab",
	HandlerType: (*ZabServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Broadcast",
			Handler:    _Zab_Broadcast_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "zab/zab.proto",
}
