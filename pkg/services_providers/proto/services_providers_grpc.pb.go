// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// ServicesProvidersServiceClient is the client API for ServicesProvidersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServicesProvidersServiceClient interface {
	Test(ctx context.Context, in *ServicesProvider, opts ...grpc.CallOption) (*TestResponse, error)
	Create(ctx context.Context, in *ServicesProvider, opts ...grpc.CallOption) (*ServicesProvider, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*ServicesProvider, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type servicesProvidersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServicesProvidersServiceClient(cc grpc.ClientConnInterface) ServicesProvidersServiceClient {
	return &servicesProvidersServiceClient{cc}
}

func (c *servicesProvidersServiceClient) Test(ctx context.Context, in *ServicesProvider, opts ...grpc.CallOption) (*TestResponse, error) {
	out := new(TestResponse)
	err := c.cc.Invoke(ctx, "/nocloud.services_providers.ServicesProvidersService/Test", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesProvidersServiceClient) Create(ctx context.Context, in *ServicesProvider, opts ...grpc.CallOption) (*ServicesProvider, error) {
	out := new(ServicesProvider)
	err := c.cc.Invoke(ctx, "/nocloud.services_providers.ServicesProvidersService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesProvidersServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*ServicesProvider, error) {
	out := new(ServicesProvider)
	err := c.cc.Invoke(ctx, "/nocloud.services_providers.ServicesProvidersService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesProvidersServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/nocloud.services_providers.ServicesProvidersService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServicesProvidersServiceServer is the server API for ServicesProvidersService service.
// All implementations must embed UnimplementedServicesProvidersServiceServer
// for forward compatibility
type ServicesProvidersServiceServer interface {
	Test(context.Context, *ServicesProvider) (*TestResponse, error)
	Create(context.Context, *ServicesProvider) (*ServicesProvider, error)
	Get(context.Context, *GetRequest) (*ServicesProvider, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	mustEmbedUnimplementedServicesProvidersServiceServer()
}

// UnimplementedServicesProvidersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServicesProvidersServiceServer struct {
}

func (UnimplementedServicesProvidersServiceServer) Test(context.Context, *ServicesProvider) (*TestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test not implemented")
}
func (UnimplementedServicesProvidersServiceServer) Create(context.Context, *ServicesProvider) (*ServicesProvider, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedServicesProvidersServiceServer) Get(context.Context, *GetRequest) (*ServicesProvider, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedServicesProvidersServiceServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedServicesProvidersServiceServer) mustEmbedUnimplementedServicesProvidersServiceServer() {
}

// UnsafeServicesProvidersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServicesProvidersServiceServer will
// result in compilation errors.
type UnsafeServicesProvidersServiceServer interface {
	mustEmbedUnimplementedServicesProvidersServiceServer()
}

func RegisterServicesProvidersServiceServer(s grpc.ServiceRegistrar, srv ServicesProvidersServiceServer) {
	s.RegisterService(&ServicesProvidersService_ServiceDesc, srv)
}

func _ServicesProvidersService_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServicesProvider)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesProvidersServiceServer).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nocloud.services_providers.ServicesProvidersService/Test",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesProvidersServiceServer).Test(ctx, req.(*ServicesProvider))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesProvidersService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServicesProvider)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesProvidersServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nocloud.services_providers.ServicesProvidersService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesProvidersServiceServer).Create(ctx, req.(*ServicesProvider))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesProvidersService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesProvidersServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nocloud.services_providers.ServicesProvidersService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesProvidersServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesProvidersService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesProvidersServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nocloud.services_providers.ServicesProvidersService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesProvidersServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServicesProvidersService_ServiceDesc is the grpc.ServiceDesc for ServicesProvidersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServicesProvidersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nocloud.services_providers.ServicesProvidersService",
	HandlerType: (*ServicesProvidersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Test",
			Handler:    _ServicesProvidersService_Test_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ServicesProvidersService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ServicesProvidersService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ServicesProvidersService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/services_providers/proto/services_providers.proto",
}
