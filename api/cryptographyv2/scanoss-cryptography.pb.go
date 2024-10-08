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
// Cryptography definition details
//*

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.27.3
// source: scanoss/api/cryptography/v2/scanoss-cryptography.proto

package cryptographyv2

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	commonv2 "github.com/scanoss/papi/api/commonv2"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//
// Cryptography Algorithm response data (JSON payload)
type AlgorithmResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Cryptography details
	Purls []*AlgorithmResponse_Purls `protobuf:"bytes,1,rep,name=purls,proto3" json:"purls,omitempty"`
	// Response status
	Status *commonv2.StatusResponse `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *AlgorithmResponse) Reset() {
	*x = AlgorithmResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlgorithmResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlgorithmResponse) ProtoMessage() {}

func (x *AlgorithmResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlgorithmResponse.ProtoReflect.Descriptor instead.
func (*AlgorithmResponse) Descriptor() ([]byte, []int) {
	return file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_rawDescGZIP(), []int{0}
}

func (x *AlgorithmResponse) GetPurls() []*AlgorithmResponse_Purls {
	if x != nil {
		return x.Purls
	}
	return nil
}

func (x *AlgorithmResponse) GetStatus() *commonv2.StatusResponse {
	if x != nil {
		return x.Status
	}
	return nil
}

//
// Cryptography Algorithm response data for a given range or criteria (JSON payload)
type AlgorithmsInRangeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Cryptography details
	Purls []*AlgorithmsInRangeResponse_Purl `protobuf:"bytes,1,rep,name=purls,proto3" json:"purls,omitempty"`
	// Response status
	Status *commonv2.StatusResponse `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *AlgorithmsInRangeResponse) Reset() {
	*x = AlgorithmsInRangeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlgorithmsInRangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlgorithmsInRangeResponse) ProtoMessage() {}

func (x *AlgorithmsInRangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlgorithmsInRangeResponse.ProtoReflect.Descriptor instead.
func (*AlgorithmsInRangeResponse) Descriptor() ([]byte, []int) {
	return file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_rawDescGZIP(), []int{1}
}

func (x *AlgorithmsInRangeResponse) GetPurls() []*AlgorithmsInRangeResponse_Purl {
	if x != nil {
		return x.Purls
	}
	return nil
}

func (x *AlgorithmsInRangeResponse) GetStatus() *commonv2.StatusResponse {
	if x != nil {
		return x.Status
	}
	return nil
}

type AlgorithmResponse_Algorithms struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Algorithm string `protobuf:"bytes,1,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	Strength  string `protobuf:"bytes,2,opt,name=strength,proto3" json:"strength,omitempty"`
}

func (x *AlgorithmResponse_Algorithms) Reset() {
	*x = AlgorithmResponse_Algorithms{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlgorithmResponse_Algorithms) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlgorithmResponse_Algorithms) ProtoMessage() {}

func (x *AlgorithmResponse_Algorithms) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlgorithmResponse_Algorithms.ProtoReflect.Descriptor instead.
func (*AlgorithmResponse_Algorithms) Descriptor() ([]byte, []int) {
	return file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AlgorithmResponse_Algorithms) GetAlgorithm() string {
	if x != nil {
		return x.Algorithm
	}
	return ""
}

func (x *AlgorithmResponse_Algorithms) GetStrength() string {
	if x != nil {
		return x.Strength
	}
	return ""
}

type AlgorithmResponse_Purls struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Purl       string                          `protobuf:"bytes,1,opt,name=purl,proto3" json:"purl,omitempty"`
	Version    string                          `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Algorithms []*AlgorithmResponse_Algorithms `protobuf:"bytes,3,rep,name=algorithms,proto3" json:"algorithms,omitempty"`
}

func (x *AlgorithmResponse_Purls) Reset() {
	*x = AlgorithmResponse_Purls{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlgorithmResponse_Purls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlgorithmResponse_Purls) ProtoMessage() {}

func (x *AlgorithmResponse_Purls) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlgorithmResponse_Purls.ProtoReflect.Descriptor instead.
func (*AlgorithmResponse_Purls) Descriptor() ([]byte, []int) {
	return file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_rawDescGZIP(), []int{0, 1}
}

func (x *AlgorithmResponse_Purls) GetPurl() string {
	if x != nil {
		return x.Purl
	}
	return ""
}

func (x *AlgorithmResponse_Purls) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *AlgorithmResponse_Purls) GetAlgorithms() []*AlgorithmResponse_Algorithms {
	if x != nil {
		return x.Algorithms
	}
	return nil
}

type AlgorithmsInRangeResponse_Algorithm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Algorithm string `protobuf:"bytes,1,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	Strength  string `protobuf:"bytes,2,opt,name=strength,proto3" json:"strength,omitempty"`
}

func (x *AlgorithmsInRangeResponse_Algorithm) Reset() {
	*x = AlgorithmsInRangeResponse_Algorithm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlgorithmsInRangeResponse_Algorithm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlgorithmsInRangeResponse_Algorithm) ProtoMessage() {}

func (x *AlgorithmsInRangeResponse_Algorithm) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlgorithmsInRangeResponse_Algorithm.ProtoReflect.Descriptor instead.
func (*AlgorithmsInRangeResponse_Algorithm) Descriptor() ([]byte, []int) {
	return file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_rawDescGZIP(), []int{1, 0}
}

func (x *AlgorithmsInRangeResponse_Algorithm) GetAlgorithm() string {
	if x != nil {
		return x.Algorithm
	}
	return ""
}

func (x *AlgorithmsInRangeResponse_Algorithm) GetStrength() string {
	if x != nil {
		return x.Strength
	}
	return ""
}

type AlgorithmsInRangeResponse_Purl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Purl       string                                 `protobuf:"bytes,1,opt,name=purl,proto3" json:"purl,omitempty"`
	Versions   []string                               `protobuf:"bytes,2,rep,name=versions,proto3" json:"versions,omitempty"`
	Algorithms []*AlgorithmsInRangeResponse_Algorithm `protobuf:"bytes,3,rep,name=algorithms,proto3" json:"algorithms,omitempty"`
}

func (x *AlgorithmsInRangeResponse_Purl) Reset() {
	*x = AlgorithmsInRangeResponse_Purl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlgorithmsInRangeResponse_Purl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlgorithmsInRangeResponse_Purl) ProtoMessage() {}

func (x *AlgorithmsInRangeResponse_Purl) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlgorithmsInRangeResponse_Purl.ProtoReflect.Descriptor instead.
func (*AlgorithmsInRangeResponse_Purl) Descriptor() ([]byte, []int) {
	return file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_rawDescGZIP(), []int{1, 1}
}

func (x *AlgorithmsInRangeResponse_Purl) GetPurl() string {
	if x != nil {
		return x.Purl
	}
	return ""
}

func (x *AlgorithmsInRangeResponse_Purl) GetVersions() []string {
	if x != nil {
		return x.Versions
	}
	return nil
}

func (x *AlgorithmsInRangeResponse_Purl) GetAlgorithms() []*AlgorithmsInRangeResponse_Algorithm {
	if x != nil {
		return x.Algorithms
	}
	return nil
}

var File_scanoss_api_cryptography_v2_scanoss_cryptography_proto protoreflect.FileDescriptor

var file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_rawDesc = []byte{
	0x0a, 0x36, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x63,
	0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2d, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x79, 0x2e, 0x76, 0x32, 0x1a, 0x2a, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x63, 0x61,
	0x6e, 0x6f, 0x73, 0x73, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77, 0x61, 0x67,
	0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf9, 0x02,
	0x0a, 0x11, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x05, 0x70, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x34, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x2e, 0x76, 0x32,
	0x2e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x50, 0x75, 0x72, 0x6c, 0x73, 0x52, 0x05, 0x70, 0x75, 0x72, 0x6c, 0x73, 0x12,
	0x3d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x46,
	0x0a, 0x0a, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74,
	0x72, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74,
	0x72, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x1a, 0x90, 0x01, 0x0a, 0x05, 0x50, 0x75, 0x72, 0x6c, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x75, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x59,
	0x0a, 0x0a, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x39, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x2e, 0x76, 0x32,
	0x2e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x73, 0x52, 0x0a, 0x61,
	0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x73, 0x22, 0x8f, 0x03, 0x0a, 0x19, 0x41, 0x6c,
	0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x73, 0x49, 0x6e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x05, 0x70, 0x75, 0x72, 0x6c, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x79, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x73, 0x49,
	0x6e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50,
	0x75, 0x72, 0x6c, 0x52, 0x05, 0x70, 0x75, 0x72, 0x6c, 0x73, 0x12, 0x3d, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x63, 0x61,
	0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x45, 0x0a, 0x09, 0x41, 0x6c, 0x67,
	0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69,
	0x74, 0x68, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72,
	0x69, 0x74, 0x68, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x1a, 0x98, 0x01, 0x0a, 0x04, 0x50, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x75, 0x72, 0x6c, 0x12, 0x1a, 0x0a,
	0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x60, 0x0a, 0x0a, 0x61, 0x6c, 0x67,
	0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x40, 0x2e,
	0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x6c, 0x67, 0x6f,
	0x72, 0x69, 0x74, 0x68, 0x6d, 0x73, 0x49, 0x6e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x52,
	0x0a, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x73, 0x32, 0xbf, 0x03, 0x0a, 0x0c,
	0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x12, 0x75, 0x0a, 0x04,
	0x45, 0x63, 0x68, 0x6f, 0x12, 0x22, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x63, 0x68,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f,
	0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x32,
	0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x22, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x32, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x2f, 0x65,
	0x63, 0x68, 0x6f, 0x12, 0x8f, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x67, 0x6f, 0x72,
	0x69, 0x74, 0x68, 0x6d, 0x73, 0x12, 0x22, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x75,
	0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x73, 0x63, 0x61, 0x6e,
	0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x24, 0x3a, 0x01, 0x2a, 0x22, 0x1f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x2f, 0x61, 0x6c, 0x67, 0x6f, 0x72,
	0x69, 0x74, 0x68, 0x6d, 0x73, 0x12, 0xa5, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x67,
	0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x73, 0x49, 0x6e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x22,
	0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x75, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x36, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x2e, 0x76, 0x32,
	0x2e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x73, 0x49, 0x6e, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x31, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x2b, 0x3a, 0x01, 0x2a, 0x22, 0x26, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x2f, 0x61, 0x6c, 0x67, 0x6f,
	0x72, 0x69, 0x74, 0x68, 0x6d, 0x73, 0x49, 0x6e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x42, 0x9e, 0x02,
	0x92, 0x41, 0xdf, 0x01, 0x12, 0x79, 0x0a, 0x1c, 0x53, 0x43, 0x41, 0x4e, 0x4f, 0x53, 0x53, 0x20,
	0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x20, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x22, 0x54, 0x0a, 0x14, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2d,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x12, 0x27, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2f, 0x63, 0x72, 0x70, 0x79, 0x74, 0x6f, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x79, 0x1a, 0x13, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x40, 0x73,
	0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x32, 0x03, 0x32, 0x2e, 0x30, 0x2a,
	0x01, 0x01, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x52, 0x3b, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x34, 0x0a,
	0x2a, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x64, 0x6f, 0x65, 0x73,
	0x20, 0x6e, 0x6f, 0x74, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x2e, 0x12, 0x06, 0x0a, 0x04, 0x9a,
	0x02, 0x01, 0x07, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2f, 0x70, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x76, 0x32, 0x3b,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x76, 0x32, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_rawDescOnce sync.Once
	file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_rawDescData = file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_rawDesc
)

func file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_rawDescGZIP() []byte {
	file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_rawDescOnce.Do(func() {
		file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_rawDescData = protoimpl.X.CompressGZIP(file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_rawDescData)
	})
	return file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_rawDescData
}

var file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_goTypes = []interface{}{
	(*AlgorithmResponse)(nil),                   // 0: scanoss.api.cryptography.v2.AlgorithmResponse
	(*AlgorithmsInRangeResponse)(nil),           // 1: scanoss.api.cryptography.v2.AlgorithmsInRangeResponse
	(*AlgorithmResponse_Algorithms)(nil),        // 2: scanoss.api.cryptography.v2.AlgorithmResponse.Algorithms
	(*AlgorithmResponse_Purls)(nil),             // 3: scanoss.api.cryptography.v2.AlgorithmResponse.Purls
	(*AlgorithmsInRangeResponse_Algorithm)(nil), // 4: scanoss.api.cryptography.v2.AlgorithmsInRangeResponse.Algorithm
	(*AlgorithmsInRangeResponse_Purl)(nil),      // 5: scanoss.api.cryptography.v2.AlgorithmsInRangeResponse.Purl
	(*commonv2.StatusResponse)(nil),             // 6: scanoss.api.common.v2.StatusResponse
	(*commonv2.EchoRequest)(nil),                // 7: scanoss.api.common.v2.EchoRequest
	(*commonv2.PurlRequest)(nil),                // 8: scanoss.api.common.v2.PurlRequest
	(*commonv2.EchoResponse)(nil),               // 9: scanoss.api.common.v2.EchoResponse
}
var file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_depIdxs = []int32{
	3, // 0: scanoss.api.cryptography.v2.AlgorithmResponse.purls:type_name -> scanoss.api.cryptography.v2.AlgorithmResponse.Purls
	6, // 1: scanoss.api.cryptography.v2.AlgorithmResponse.status:type_name -> scanoss.api.common.v2.StatusResponse
	5, // 2: scanoss.api.cryptography.v2.AlgorithmsInRangeResponse.purls:type_name -> scanoss.api.cryptography.v2.AlgorithmsInRangeResponse.Purl
	6, // 3: scanoss.api.cryptography.v2.AlgorithmsInRangeResponse.status:type_name -> scanoss.api.common.v2.StatusResponse
	2, // 4: scanoss.api.cryptography.v2.AlgorithmResponse.Purls.algorithms:type_name -> scanoss.api.cryptography.v2.AlgorithmResponse.Algorithms
	4, // 5: scanoss.api.cryptography.v2.AlgorithmsInRangeResponse.Purl.algorithms:type_name -> scanoss.api.cryptography.v2.AlgorithmsInRangeResponse.Algorithm
	7, // 6: scanoss.api.cryptography.v2.Cryptography.Echo:input_type -> scanoss.api.common.v2.EchoRequest
	8, // 7: scanoss.api.cryptography.v2.Cryptography.GetAlgorithms:input_type -> scanoss.api.common.v2.PurlRequest
	8, // 8: scanoss.api.cryptography.v2.Cryptography.GetAlgorithmsInRange:input_type -> scanoss.api.common.v2.PurlRequest
	9, // 9: scanoss.api.cryptography.v2.Cryptography.Echo:output_type -> scanoss.api.common.v2.EchoResponse
	0, // 10: scanoss.api.cryptography.v2.Cryptography.GetAlgorithms:output_type -> scanoss.api.cryptography.v2.AlgorithmResponse
	1, // 11: scanoss.api.cryptography.v2.Cryptography.GetAlgorithmsInRange:output_type -> scanoss.api.cryptography.v2.AlgorithmsInRangeResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_init() }
func file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_init() {
	if File_scanoss_api_cryptography_v2_scanoss_cryptography_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlgorithmResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlgorithmsInRangeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlgorithmResponse_Algorithms); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlgorithmResponse_Purls); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlgorithmsInRangeResponse_Algorithm); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlgorithmsInRangeResponse_Purl); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_goTypes,
		DependencyIndexes: file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_depIdxs,
		MessageInfos:      file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes,
	}.Build()
	File_scanoss_api_cryptography_v2_scanoss_cryptography_proto = out.File
	file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_rawDesc = nil
	file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_goTypes = nil
	file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_depIdxs = nil
}
