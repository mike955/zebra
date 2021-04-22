// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package age

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

// AgeServiceClient is the client API for AgeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgeServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type ageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgeServiceClient(cc grpc.ClientConnInterface) AgeServiceClient {
	return &ageServiceClient{cc}
}

func (c *ageServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/age.AgeService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgeServiceServer is the server API for AgeService service.
// All implementations must embed UnimplementedAgeServiceServer
// for forward compatibility
type AgeServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	mustEmbedUnimplementedAgeServiceServer()
}

// UnimplementedAgeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAgeServiceServer struct {
}

func (UnimplementedAgeServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAgeServiceServer) mustEmbedUnimplementedAgeServiceServer() {}

// UnsafeAgeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgeServiceServer will
// result in compilation errors.
type UnsafeAgeServiceServer interface {
	mustEmbedUnimplementedAgeServiceServer()
}

func RegisterAgeServiceServer(s grpc.ServiceRegistrar, srv AgeServiceServer) {
	s.RegisterService(&AgeService_ServiceDesc, srv)
}

func _AgeService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgeServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/age.AgeService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgeServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AgeService_ServiceDesc is the grpc.ServiceDesc for AgeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "age.AgeService",
	HandlerType: (*AgeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _AgeService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "age.proto",
}
