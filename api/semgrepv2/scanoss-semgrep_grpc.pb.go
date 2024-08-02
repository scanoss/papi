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

//**
// semgrep definition details
//*

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.3
// source: scanoss/api/semgrep/v2/scanoss-semgrep.proto

package semgrepv2

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
	Semgrep_Echo_FullMethodName      = "/scanoss.api.semgrep.v2.Semgrep/Echo"
	Semgrep_GetIssues_FullMethodName = "/scanoss.api.semgrep.v2.Semgrep/GetIssues"
)

// SemgrepClient is the client API for Semgrep service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SemgrepClient interface {
	// Standard echo
	Echo(ctx context.Context, in *commonv2.EchoRequest, opts ...grpc.CallOption) (*commonv2.EchoResponse, error)
	// Get Potential issues  associated with a list of PURLs
	GetIssues(ctx context.Context, in *commonv2.PurlRequest, opts ...grpc.CallOption) (*SemgrepResponse, error)
}

type semgrepClient struct {
	cc grpc.ClientConnInterface
}

func NewSemgrepClient(cc grpc.ClientConnInterface) SemgrepClient {
	return &semgrepClient{cc}
}

func (c *semgrepClient) Echo(ctx context.Context, in *commonv2.EchoRequest, opts ...grpc.CallOption) (*commonv2.EchoResponse, error) {
	out := new(commonv2.EchoResponse)
	err := c.cc.Invoke(ctx, Semgrep_Echo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *semgrepClient) GetIssues(ctx context.Context, in *commonv2.PurlRequest, opts ...grpc.CallOption) (*SemgrepResponse, error) {
	out := new(SemgrepResponse)
	err := c.cc.Invoke(ctx, Semgrep_GetIssues_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SemgrepServer is the server API for Semgrep service.
// All implementations must embed UnimplementedSemgrepServer
// for forward compatibility
type SemgrepServer interface {
	// Standard echo
	Echo(context.Context, *commonv2.EchoRequest) (*commonv2.EchoResponse, error)
	// Get Potential issues  associated with a list of PURLs
	GetIssues(context.Context, *commonv2.PurlRequest) (*SemgrepResponse, error)
	mustEmbedUnimplementedSemgrepServer()
}

// UnimplementedSemgrepServer must be embedded to have forward compatible implementations.
type UnimplementedSemgrepServer struct {
}

func (UnimplementedSemgrepServer) Echo(context.Context, *commonv2.EchoRequest) (*commonv2.EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedSemgrepServer) GetIssues(context.Context, *commonv2.PurlRequest) (*SemgrepResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIssues not implemented")
}
func (UnimplementedSemgrepServer) mustEmbedUnimplementedSemgrepServer() {}

// UnsafeSemgrepServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SemgrepServer will
// result in compilation errors.
type UnsafeSemgrepServer interface {
	mustEmbedUnimplementedSemgrepServer()
}

func RegisterSemgrepServer(s grpc.ServiceRegistrar, srv SemgrepServer) {
	s.RegisterService(&Semgrep_ServiceDesc, srv)
}

func _Semgrep_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commonv2.EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SemgrepServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Semgrep_Echo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SemgrepServer).Echo(ctx, req.(*commonv2.EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Semgrep_GetIssues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commonv2.PurlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SemgrepServer).GetIssues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Semgrep_GetIssues_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SemgrepServer).GetIssues(ctx, req.(*commonv2.PurlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Semgrep_ServiceDesc is the grpc.ServiceDesc for Semgrep service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Semgrep_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "scanoss.api.semgrep.v2.Semgrep",
	HandlerType: (*SemgrepServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Semgrep_Echo_Handler,
		},
		{
			MethodName: "GetIssues",
			Handler:    _Semgrep_GetIssues_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scanoss/api/semgrep/v2/scanoss-semgrep.proto",
}
