// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package componentsv2

import (
	context "context"
	commonv2 "github.com/scanoss/papi/api/commonv2"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ComponentsClient is the client API for Components service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComponentsClient interface {
	// Standard echo
	Echo(ctx context.Context, in *commonv2.EchoRequest, opts ...grpc.CallOption) (*commonv2.EchoResponse, error)
	// Get dependency details
	SearchComponents(ctx context.Context, in *CompSearchRequest, opts ...grpc.CallOption) (*CompSearchResponse, error)
	GetComponentVersions(ctx context.Context, in *CompSearchRequest, opts ...grpc.CallOption) (*CompSearchResponse, error)
}

type componentsClient struct {
	cc grpc.ClientConnInterface
}

func NewComponentsClient(cc grpc.ClientConnInterface) ComponentsClient {
	return &componentsClient{cc}
}

func (c *componentsClient) Echo(ctx context.Context, in *commonv2.EchoRequest, opts ...grpc.CallOption) (*commonv2.EchoResponse, error) {
	out := new(commonv2.EchoResponse)
	err := c.cc.Invoke(ctx, "/scanoss.api.components.v2.Components/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *componentsClient) SearchComponents(ctx context.Context, in *CompSearchRequest, opts ...grpc.CallOption) (*CompSearchResponse, error) {
	out := new(CompSearchResponse)
	err := c.cc.Invoke(ctx, "/scanoss.api.components.v2.Components/SearchComponents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *componentsClient) GetComponentVersions(ctx context.Context, in *CompSearchRequest, opts ...grpc.CallOption) (*CompSearchResponse, error) {
	out := new(CompSearchResponse)
	err := c.cc.Invoke(ctx, "/scanoss.api.components.v2.Components/GetComponentVersions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComponentsServer is the server API for Components service.
// All implementations must embed UnimplementedComponentsServer
// for forward compatibility
type ComponentsServer interface {
	// Standard echo
	Echo(context.Context, *commonv2.EchoRequest) (*commonv2.EchoResponse, error)
	// Get dependency details
	SearchComponents(context.Context, *CompSearchRequest) (*CompSearchResponse, error)
	GetComponentVersions(context.Context, *CompSearchRequest) (*CompSearchResponse, error)
	mustEmbedUnimplementedComponentsServer()
}

// UnimplementedComponentsServer must be embedded to have forward compatible implementations.
type UnimplementedComponentsServer struct {
}

func (UnimplementedComponentsServer) Echo(context.Context, *commonv2.EchoRequest) (*commonv2.EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedComponentsServer) SearchComponents(context.Context, *CompSearchRequest) (*CompSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchComponents not implemented")
}
func (UnimplementedComponentsServer) GetComponentVersions(context.Context, *CompSearchRequest) (*CompSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComponentVersions not implemented")
}
func (UnimplementedComponentsServer) mustEmbedUnimplementedComponentsServer() {}

// UnsafeComponentsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComponentsServer will
// result in compilation errors.
type UnsafeComponentsServer interface {
	mustEmbedUnimplementedComponentsServer()
}

func RegisterComponentsServer(s grpc.ServiceRegistrar, srv ComponentsServer) {
	s.RegisterService(&Components_ServiceDesc, srv)
}

func _Components_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commonv2.EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComponentsServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scanoss.api.components.v2.Components/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComponentsServer).Echo(ctx, req.(*commonv2.EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Components_SearchComponents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComponentsServer).SearchComponents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scanoss.api.components.v2.Components/SearchComponents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComponentsServer).SearchComponents(ctx, req.(*CompSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Components_GetComponentVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComponentsServer).GetComponentVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scanoss.api.components.v2.Components/GetComponentVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComponentsServer).GetComponentVersions(ctx, req.(*CompSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Components_ServiceDesc is the grpc.ServiceDesc for Components service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Components_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "scanoss.api.components.v2.Components",
	HandlerType: (*ComponentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Components_Echo_Handler,
		},
		{
			MethodName: "SearchComponents",
			Handler:    _Components_SearchComponents_Handler,
		},
		{
			MethodName: "GetComponentVersions",
			Handler:    _Components_GetComponentVersions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scanoss/api/components/v2/scanoss-components.proto",
}