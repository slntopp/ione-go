// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	structpb "google.golang.org/protobuf/types/known/structpb"
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
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	Update(ctx context.Context, in *ServicesProvider, opts ...grpc.CallOption) (*ServicesProvider, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*ServicesProvider, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Invoke(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*structpb.Struct, error)
	ListExtentions(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListExtentionsResponse, error)
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

func (c *servicesProvidersServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/nocloud.services_providers.ServicesProvidersService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesProvidersServiceClient) Update(ctx context.Context, in *ServicesProvider, opts ...grpc.CallOption) (*ServicesProvider, error) {
	out := new(ServicesProvider)
	err := c.cc.Invoke(ctx, "/nocloud.services_providers.ServicesProvidersService/Update", in, out, opts...)
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

func (c *servicesProvidersServiceClient) Invoke(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*structpb.Struct, error) {
	out := new(structpb.Struct)
	err := c.cc.Invoke(ctx, "/nocloud.services_providers.ServicesProvidersService/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesProvidersServiceClient) ListExtentions(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListExtentionsResponse, error) {
	out := new(ListExtentionsResponse)
	err := c.cc.Invoke(ctx, "/nocloud.services_providers.ServicesProvidersService/ListExtentions", in, out, opts...)
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
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	Update(context.Context, *ServicesProvider) (*ServicesProvider, error)
	Get(context.Context, *GetRequest) (*ServicesProvider, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Invoke(context.Context, *ActionRequest) (*structpb.Struct, error)
	ListExtentions(context.Context, *ListRequest) (*ListExtentionsResponse, error)
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
func (UnimplementedServicesProvidersServiceServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedServicesProvidersServiceServer) Update(context.Context, *ServicesProvider) (*ServicesProvider, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedServicesProvidersServiceServer) Get(context.Context, *GetRequest) (*ServicesProvider, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedServicesProvidersServiceServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedServicesProvidersServiceServer) Invoke(context.Context, *ActionRequest) (*structpb.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}
func (UnimplementedServicesProvidersServiceServer) ListExtentions(context.Context, *ListRequest) (*ListExtentionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListExtentions not implemented")
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

func _ServicesProvidersService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesProvidersServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nocloud.services_providers.ServicesProvidersService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesProvidersServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesProvidersService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServicesProvider)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesProvidersServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nocloud.services_providers.ServicesProvidersService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesProvidersServiceServer).Update(ctx, req.(*ServicesProvider))
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

func _ServicesProvidersService_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesProvidersServiceServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nocloud.services_providers.ServicesProvidersService/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesProvidersServiceServer).Invoke(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesProvidersService_ListExtentions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesProvidersServiceServer).ListExtentions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nocloud.services_providers.ServicesProvidersService/ListExtentions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesProvidersServiceServer).ListExtentions(ctx, req.(*ListRequest))
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
			MethodName: "Delete",
			Handler:    _ServicesProvidersService_Delete_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ServicesProvidersService_Update_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ServicesProvidersService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ServicesProvidersService_List_Handler,
		},
		{
			MethodName: "Invoke",
			Handler:    _ServicesProvidersService_Invoke_Handler,
		},
		{
			MethodName: "ListExtentions",
			Handler:    _ServicesProvidersService_ListExtentions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/services_providers/proto/services_providers.proto",
}

// ServicesProvidersExtentionsServiceClient is the client API for ServicesProvidersExtentionsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServicesProvidersExtentionsServiceClient interface {
	GetType(ctx context.Context, in *GetTypeRequest, opts ...grpc.CallOption) (*GetTypeResponse, error)
	Test(ctx context.Context, in *ServicesProvidersExtentionData, opts ...grpc.CallOption) (*GenericResponse, error)
	Register(ctx context.Context, in *ServicesProvidersExtentionData, opts ...grpc.CallOption) (*GenericResponse, error)
	Update(ctx context.Context, in *ServicesProvidersExtentionData, opts ...grpc.CallOption) (*GenericResponse, error)
	Unregister(ctx context.Context, in *ServicesProvidersExtentionData, opts ...grpc.CallOption) (*GenericResponse, error)
}

type servicesProvidersExtentionsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServicesProvidersExtentionsServiceClient(cc grpc.ClientConnInterface) ServicesProvidersExtentionsServiceClient {
	return &servicesProvidersExtentionsServiceClient{cc}
}

func (c *servicesProvidersExtentionsServiceClient) GetType(ctx context.Context, in *GetTypeRequest, opts ...grpc.CallOption) (*GetTypeResponse, error) {
	out := new(GetTypeResponse)
	err := c.cc.Invoke(ctx, "/nocloud.services_providers.ServicesProvidersExtentionsService/GetType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesProvidersExtentionsServiceClient) Test(ctx context.Context, in *ServicesProvidersExtentionData, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/nocloud.services_providers.ServicesProvidersExtentionsService/Test", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesProvidersExtentionsServiceClient) Register(ctx context.Context, in *ServicesProvidersExtentionData, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/nocloud.services_providers.ServicesProvidersExtentionsService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesProvidersExtentionsServiceClient) Update(ctx context.Context, in *ServicesProvidersExtentionData, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/nocloud.services_providers.ServicesProvidersExtentionsService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesProvidersExtentionsServiceClient) Unregister(ctx context.Context, in *ServicesProvidersExtentionData, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/nocloud.services_providers.ServicesProvidersExtentionsService/Unregister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServicesProvidersExtentionsServiceServer is the server API for ServicesProvidersExtentionsService service.
// All implementations must embed UnimplementedServicesProvidersExtentionsServiceServer
// for forward compatibility
type ServicesProvidersExtentionsServiceServer interface {
	GetType(context.Context, *GetTypeRequest) (*GetTypeResponse, error)
	Test(context.Context, *ServicesProvidersExtentionData) (*GenericResponse, error)
	Register(context.Context, *ServicesProvidersExtentionData) (*GenericResponse, error)
	Update(context.Context, *ServicesProvidersExtentionData) (*GenericResponse, error)
	Unregister(context.Context, *ServicesProvidersExtentionData) (*GenericResponse, error)
	mustEmbedUnimplementedServicesProvidersExtentionsServiceServer()
}

// UnimplementedServicesProvidersExtentionsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServicesProvidersExtentionsServiceServer struct {
}

func (UnimplementedServicesProvidersExtentionsServiceServer) GetType(context.Context, *GetTypeRequest) (*GetTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetType not implemented")
}
func (UnimplementedServicesProvidersExtentionsServiceServer) Test(context.Context, *ServicesProvidersExtentionData) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test not implemented")
}
func (UnimplementedServicesProvidersExtentionsServiceServer) Register(context.Context, *ServicesProvidersExtentionData) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedServicesProvidersExtentionsServiceServer) Update(context.Context, *ServicesProvidersExtentionData) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedServicesProvidersExtentionsServiceServer) Unregister(context.Context, *ServicesProvidersExtentionData) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unregister not implemented")
}
func (UnimplementedServicesProvidersExtentionsServiceServer) mustEmbedUnimplementedServicesProvidersExtentionsServiceServer() {
}

// UnsafeServicesProvidersExtentionsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServicesProvidersExtentionsServiceServer will
// result in compilation errors.
type UnsafeServicesProvidersExtentionsServiceServer interface {
	mustEmbedUnimplementedServicesProvidersExtentionsServiceServer()
}

func RegisterServicesProvidersExtentionsServiceServer(s grpc.ServiceRegistrar, srv ServicesProvidersExtentionsServiceServer) {
	s.RegisterService(&ServicesProvidersExtentionsService_ServiceDesc, srv)
}

func _ServicesProvidersExtentionsService_GetType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesProvidersExtentionsServiceServer).GetType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nocloud.services_providers.ServicesProvidersExtentionsService/GetType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesProvidersExtentionsServiceServer).GetType(ctx, req.(*GetTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesProvidersExtentionsService_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServicesProvidersExtentionData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesProvidersExtentionsServiceServer).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nocloud.services_providers.ServicesProvidersExtentionsService/Test",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesProvidersExtentionsServiceServer).Test(ctx, req.(*ServicesProvidersExtentionData))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesProvidersExtentionsService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServicesProvidersExtentionData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesProvidersExtentionsServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nocloud.services_providers.ServicesProvidersExtentionsService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesProvidersExtentionsServiceServer).Register(ctx, req.(*ServicesProvidersExtentionData))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesProvidersExtentionsService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServicesProvidersExtentionData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesProvidersExtentionsServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nocloud.services_providers.ServicesProvidersExtentionsService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesProvidersExtentionsServiceServer).Update(ctx, req.(*ServicesProvidersExtentionData))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesProvidersExtentionsService_Unregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServicesProvidersExtentionData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesProvidersExtentionsServiceServer).Unregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nocloud.services_providers.ServicesProvidersExtentionsService/Unregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesProvidersExtentionsServiceServer).Unregister(ctx, req.(*ServicesProvidersExtentionData))
	}
	return interceptor(ctx, in, info, handler)
}

// ServicesProvidersExtentionsService_ServiceDesc is the grpc.ServiceDesc for ServicesProvidersExtentionsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServicesProvidersExtentionsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nocloud.services_providers.ServicesProvidersExtentionsService",
	HandlerType: (*ServicesProvidersExtentionsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetType",
			Handler:    _ServicesProvidersExtentionsService_GetType_Handler,
		},
		{
			MethodName: "Test",
			Handler:    _ServicesProvidersExtentionsService_Test_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _ServicesProvidersExtentionsService_Register_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ServicesProvidersExtentionsService_Update_Handler,
		},
		{
			MethodName: "Unregister",
			Handler:    _ServicesProvidersExtentionsService_Unregister_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/services_providers/proto/services_providers.proto",
}
