//
//SPDX-License-Identifier: MIT
//
//Copyright (c) 2021, SCANOSS
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
//Scanning definition details
//*

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: scanoss/api/scanning/v2/scanoss-scanning.proto

package scanningv2

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
	Scanning_Echo_FullMethodName = "/scanoss.api.scanning.v2.Scanning/Echo"
)

// ScanningClient is the client API for Scanning service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScanningClient interface {
	// Standard echo
	Echo(ctx context.Context, in *commonv2.EchoRequest, opts ...grpc.CallOption) (*commonv2.EchoResponse, error)
}

type scanningClient struct {
	cc grpc.ClientConnInterface
}

func NewScanningClient(cc grpc.ClientConnInterface) ScanningClient {
	return &scanningClient{cc}
}

func (c *scanningClient) Echo(ctx context.Context, in *commonv2.EchoRequest, opts ...grpc.CallOption) (*commonv2.EchoResponse, error) {
	out := new(commonv2.EchoResponse)
	err := c.cc.Invoke(ctx, Scanning_Echo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScanningServer is the server API for Scanning service.
// All implementations must embed UnimplementedScanningServer
// for forward compatibility
type ScanningServer interface {
	// Standard echo
	Echo(context.Context, *commonv2.EchoRequest) (*commonv2.EchoResponse, error)
	mustEmbedUnimplementedScanningServer()
}

// UnimplementedScanningServer must be embedded to have forward compatible implementations.
type UnimplementedScanningServer struct {
}

func (UnimplementedScanningServer) Echo(context.Context, *commonv2.EchoRequest) (*commonv2.EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedScanningServer) mustEmbedUnimplementedScanningServer() {}

// UnsafeScanningServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScanningServer will
// result in compilation errors.
type UnsafeScanningServer interface {
	mustEmbedUnimplementedScanningServer()
}

func RegisterScanningServer(s grpc.ServiceRegistrar, srv ScanningServer) {
	s.RegisterService(&Scanning_ServiceDesc, srv)
}

func _Scanning_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commonv2.EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScanningServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Scanning_Echo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScanningServer).Echo(ctx, req.(*commonv2.EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Scanning_ServiceDesc is the grpc.ServiceDesc for Scanning service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Scanning_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "scanoss.api.scanning.v2.Scanning",
	HandlerType: (*ScanningServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Scanning_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scanoss/api/scanning/v2/scanoss-scanning.proto",
}
