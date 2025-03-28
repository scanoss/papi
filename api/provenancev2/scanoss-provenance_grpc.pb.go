//
//SPDX-License-Identifier: MIT
//
//Copyright (c) 2023, SCANOSS
//
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in
//all copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
//THE SOFTWARE.

//*
// Provenance definition details

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.3
// source: scanoss/api/provenance/v2/scanoss-provenance.proto

package provenancev2

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

const (
	Provenance_Echo_FullMethodName                   = "/scanoss.api.provenance.v2.Provenance/Echo"
	Provenance_GetComponentProvenance_FullMethodName = "/scanoss.api.provenance.v2.Provenance/GetComponentProvenance"
)

// ProvenanceClient is the client API for Provenance service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProvenanceClient interface {
	// Standard echo
	Echo(ctx context.Context, in *commonv2.EchoRequest, opts ...grpc.CallOption) (*commonv2.EchoResponse, error)
	// Get Provenance countries associated with a list of PURLs
	GetComponentProvenance(ctx context.Context, in *commonv2.PurlRequest, opts ...grpc.CallOption) (*ProvenanceResponse, error)
}

type provenanceClient struct {
	cc grpc.ClientConnInterface
}

func NewProvenanceClient(cc grpc.ClientConnInterface) ProvenanceClient {
	return &provenanceClient{cc}
}

func (c *provenanceClient) Echo(ctx context.Context, in *commonv2.EchoRequest, opts ...grpc.CallOption) (*commonv2.EchoResponse, error) {
	out := new(commonv2.EchoResponse)
	err := c.cc.Invoke(ctx, Provenance_Echo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *provenanceClient) GetComponentProvenance(ctx context.Context, in *commonv2.PurlRequest, opts ...grpc.CallOption) (*ProvenanceResponse, error) {
	out := new(ProvenanceResponse)
	err := c.cc.Invoke(ctx, Provenance_GetComponentProvenance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProvenanceServer is the server API for Provenance service.
// All implementations must embed UnimplementedProvenanceServer
// for forward compatibility
type ProvenanceServer interface {
	// Standard echo
	Echo(context.Context, *commonv2.EchoRequest) (*commonv2.EchoResponse, error)
	// Get Provenance countries associated with a list of PURLs
	GetComponentProvenance(context.Context, *commonv2.PurlRequest) (*ProvenanceResponse, error)
	mustEmbedUnimplementedProvenanceServer()
}

// UnimplementedProvenanceServer must be embedded to have forward compatible implementations.
type UnimplementedProvenanceServer struct {
}

func (UnimplementedProvenanceServer) Echo(context.Context, *commonv2.EchoRequest) (*commonv2.EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedProvenanceServer) GetComponentProvenance(context.Context, *commonv2.PurlRequest) (*ProvenanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComponentProvenance not implemented")
}
func (UnimplementedProvenanceServer) mustEmbedUnimplementedProvenanceServer() {}

// UnsafeProvenanceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProvenanceServer will
// result in compilation errors.
type UnsafeProvenanceServer interface {
	mustEmbedUnimplementedProvenanceServer()
}

func RegisterProvenanceServer(s grpc.ServiceRegistrar, srv ProvenanceServer) {
	s.RegisterService(&Provenance_ServiceDesc, srv)
}

func _Provenance_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commonv2.EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvenanceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Provenance_Echo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvenanceServer).Echo(ctx, req.(*commonv2.EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provenance_GetComponentProvenance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commonv2.PurlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvenanceServer).GetComponentProvenance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Provenance_GetComponentProvenance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvenanceServer).GetComponentProvenance(ctx, req.(*commonv2.PurlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Provenance_ServiceDesc is the grpc.ServiceDesc for Provenance service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Provenance_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "scanoss.api.provenance.v2.Provenance",
	HandlerType: (*ProvenanceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Provenance_Echo_Handler,
		},
		{
			MethodName: "GetComponentProvenance",
			Handler:    _Provenance_GetComponentProvenance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scanoss/api/provenance/v2/scanoss-provenance.proto",
}
