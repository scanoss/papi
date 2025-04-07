//
//SPDX-License-Identifier: MIT
//
//Copyright (c) 2024, SCANOSS
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.29.3
// source: scanoss/api/scanning/v2/scanoss-scanning.proto

package scanningv2

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

// High precision Folder Hashing scan request
type HFHRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Report best match (most hits) or not. Default: false
	BestMatch bool `protobuf:"varint,1,opt,name=best_match,json=bestMatch,proto3" json:"best_match,omitempty"`
	// Detection threshold (distance) for component selection. Default: 60
	Threshold int32 `protobuf:"varint,2,opt,name=threshold,proto3" json:"threshold,omitempty"`
	// Folder root node to be scanned
	Root *HFHRequest_Children `protobuf:"bytes,3,opt,name=root,proto3" json:"root,omitempty"`
}

func (x *HFHRequest) Reset() {
	*x = HFHRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanoss_api_scanning_v2_scanoss_scanning_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HFHRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HFHRequest) ProtoMessage() {}

func (x *HFHRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_scanning_v2_scanoss_scanning_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HFHRequest.ProtoReflect.Descriptor instead.
func (*HFHRequest) Descriptor() ([]byte, []int) {
	return file_scanoss_api_scanning_v2_scanoss_scanning_proto_rawDescGZIP(), []int{0}
}

func (x *HFHRequest) GetBestMatch() bool {
	if x != nil {
		return x.BestMatch
	}
	return false
}

func (x *HFHRequest) GetThreshold() int32 {
	if x != nil {
		return x.Threshold
	}
	return 0
}

func (x *HFHRequest) GetRoot() *HFHRequest_Children {
	if x != nil {
		return x.Root
	}
	return nil
}

// High precision Folder Hashing scan response
type HFHResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of folders containing unique components
	Results []*HFHResponse_Result    `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Status  *commonv2.StatusResponse `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *HFHResponse) Reset() {
	*x = HFHResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanoss_api_scanning_v2_scanoss_scanning_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HFHResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HFHResponse) ProtoMessage() {}

func (x *HFHResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_scanning_v2_scanoss_scanning_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HFHResponse.ProtoReflect.Descriptor instead.
func (*HFHResponse) Descriptor() ([]byte, []int) {
	return file_scanoss_api_scanning_v2_scanoss_scanning_proto_rawDescGZIP(), []int{1}
}

func (x *HFHResponse) GetResults() []*HFHResponse_Result {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *HFHResponse) GetStatus() *commonv2.StatusResponse {
	if x != nil {
		return x.Status
	}
	return nil
}

// Child node from the folder structure
type HFHRequest_Children struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Folder path (can be actual or obfuscated)
	PathId string `protobuf:"bytes,1,opt,name=path_id,json=pathId,proto3" json:"path_id,omitempty"`
	// Proximity hash calculated from this nodes filenames (and their children).
	SimHashNames string `protobuf:"bytes,2,opt,name=sim_hash_names,json=simHashNames,proto3" json:"sim_hash_names,omitempty"`
	// Proximity hash calculated from this nodes file contents (and their children).
	SimHashContent string `protobuf:"bytes,3,opt,name=sim_hash_content,json=simHashContent,proto3" json:"sim_hash_content,omitempty"`
	// Sub-folders inside this child
	Children []*HFHRequest_Children `protobuf:"bytes,4,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *HFHRequest_Children) Reset() {
	*x = HFHRequest_Children{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanoss_api_scanning_v2_scanoss_scanning_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HFHRequest_Children) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HFHRequest_Children) ProtoMessage() {}

func (x *HFHRequest_Children) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_scanning_v2_scanoss_scanning_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HFHRequest_Children.ProtoReflect.Descriptor instead.
func (*HFHRequest_Children) Descriptor() ([]byte, []int) {
	return file_scanoss_api_scanning_v2_scanoss_scanning_proto_rawDescGZIP(), []int{0, 0}
}

func (x *HFHRequest_Children) GetPathId() string {
	if x != nil {
		return x.PathId
	}
	return ""
}

func (x *HFHRequest_Children) GetSimHashNames() string {
	if x != nil {
		return x.SimHashNames
	}
	return ""
}

func (x *HFHRequest_Children) GetSimHashContent() string {
	if x != nil {
		return x.SimHashContent
	}
	return ""
}

func (x *HFHRequest_Children) GetChildren() []*HFHRequest_Children {
	if x != nil {
		return x.Children
	}
	return nil
}

// Matched component details
type HFHResponse_Component struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Component PURL
	Purl string `protobuf:"bytes,1,opt,name=purl,proto3" json:"purl,omitempty"`
	// Component match version (could be multiple)
	Versions []string `protobuf:"bytes,2,rep,name=versions,proto3" json:"versions,omitempty"`
	// Result confidence percentage
	Confidence float32 `protobuf:"fixed32,3,opt,name=confidence,proto3" json:"confidence,omitempty"`
}

func (x *HFHResponse_Component) Reset() {
	*x = HFHResponse_Component{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanoss_api_scanning_v2_scanoss_scanning_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HFHResponse_Component) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HFHResponse_Component) ProtoMessage() {}

func (x *HFHResponse_Component) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_scanning_v2_scanoss_scanning_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HFHResponse_Component.ProtoReflect.Descriptor instead.
func (*HFHResponse_Component) Descriptor() ([]byte, []int) {
	return file_scanoss_api_scanning_v2_scanoss_scanning_proto_rawDescGZIP(), []int{1, 0}
}

func (x *HFHResponse_Component) GetPurl() string {
	if x != nil {
		return x.Purl
	}
	return ""
}

func (x *HFHResponse_Component) GetVersions() []string {
	if x != nil {
		return x.Versions
	}
	return nil
}

func (x *HFHResponse_Component) GetConfidence() float32 {
	if x != nil {
		return x.Confidence
	}
	return 0
}

// Result item, link a path with a list of components
type HFHResponse_Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Folder path (can be actual or obfuscated)
	PathId string `protobuf:"bytes,1,opt,name=path_id,json=pathId,proto3" json:"path_id,omitempty"`
	// List of matching components
	Components []*HFHResponse_Component `protobuf:"bytes,2,rep,name=components,proto3" json:"components,omitempty"`
}

func (x *HFHResponse_Result) Reset() {
	*x = HFHResponse_Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanoss_api_scanning_v2_scanoss_scanning_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HFHResponse_Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HFHResponse_Result) ProtoMessage() {}

func (x *HFHResponse_Result) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_scanning_v2_scanoss_scanning_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HFHResponse_Result.ProtoReflect.Descriptor instead.
func (*HFHResponse_Result) Descriptor() ([]byte, []int) {
	return file_scanoss_api_scanning_v2_scanoss_scanning_proto_rawDescGZIP(), []int{1, 1}
}

func (x *HFHResponse_Result) GetPathId() string {
	if x != nil {
		return x.PathId
	}
	return ""
}

func (x *HFHResponse_Result) GetComponents() []*HFHResponse_Component {
	if x != nil {
		return x.Components
	}
	return nil
}

var File_scanoss_api_scanning_v2_scanoss_scanning_proto protoreflect.FileDescriptor

var file_scanoss_api_scanning_v2_scanoss_scanning_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x63,
	0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73,
	0x73, 0x2d, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x17, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x63,
	0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x1a, 0x2a, 0x73, 0x63, 0x61, 0x6e, 0x6f,
	0x73, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32,
	0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xcb, 0x02, 0x0a, 0x0a, 0x48, 0x46, 0x48, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x65, 0x73, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x62, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x40, 0x0a,
	0x04, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x73, 0x63,
	0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x48, 0x46, 0x48, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x1a,
	0xbd, 0x01, 0x0a, 0x08, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x12, 0x17, 0x0a, 0x07,
	0x70, 0x61, 0x74, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x61, 0x74, 0x68, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x69, 0x6d, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73,
	0x69, 0x6d, 0x48, 0x61, 0x73, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x73,
	0x69, 0x6d, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x69, 0x6d, 0x48, 0x61, 0x73, 0x68, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x48, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65,
	0x6e, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x32, 0x2e, 0x48, 0x46, 0x48, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x68, 0x69,
	0x6c, 0x64, 0x72, 0x65, 0x6e, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22,
	0xe3, 0x02, 0x0a, 0x0b, 0x48, 0x46, 0x48, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x45, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x63, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x48, 0x46, 0x48, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x3d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x5b, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x75, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e,
	0x63, 0x65, 0x1a, 0x71, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x70, 0x61, 0x74, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x61, 0x74, 0x68, 0x49, 0x64, 0x12, 0x4e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x73, 0x63, 0x61, 0x6e,
	0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x32, 0x2e, 0x48, 0x46, 0x48, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x81, 0x02, 0x0a, 0x08, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x69,
	0x6e, 0x67, 0x12, 0x71, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x22, 0x2e, 0x73, 0x63, 0x61,
	0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x32, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23,
	0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x63,
	0x68, 0x6f, 0x3a, 0x01, 0x2a, 0x12, 0x81, 0x01, 0x0a, 0x0e, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x48, 0x61, 0x73, 0x68, 0x53, 0x63, 0x61, 0x6e, 0x12, 0x23, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f,
	0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x32, 0x2e, 0x48, 0x46, 0x48, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x63, 0x61, 0x6e,
	0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x48, 0x46, 0x48, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x19, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x68, 0x66,
	0x68, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x3a, 0x01, 0x2a, 0x42, 0x8a, 0x02, 0x5a, 0x31, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73,
	0x2f, 0x70, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x69,
	0x6e, 0x67, 0x76, 0x32, 0x3b, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x76, 0x32, 0x92,
	0x41, 0xd3, 0x01, 0x12, 0x6d, 0x0a, 0x18, 0x53, 0x43, 0x41, 0x4e, 0x4f, 0x53, 0x53, 0x20, 0x53,
	0x63, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22,
	0x4c, 0x0a, 0x10, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2d, 0x73, 0x63, 0x61, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x12, 0x23, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2f,
	0x73, 0x63, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x1a, 0x13, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x40, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x32, 0x03, 0x32,
	0x2e, 0x30, 0x2a, 0x01, 0x01, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x52, 0x3b, 0x0a, 0x03, 0x34, 0x30, 0x34,
	0x12, 0x34, 0x0a, 0x2a, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65,
	0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x64,
	0x6f, 0x65, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x2e, 0x12, 0x06,
	0x0a, 0x04, 0x9a, 0x02, 0x01, 0x07, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_scanoss_api_scanning_v2_scanoss_scanning_proto_rawDescOnce sync.Once
	file_scanoss_api_scanning_v2_scanoss_scanning_proto_rawDescData = file_scanoss_api_scanning_v2_scanoss_scanning_proto_rawDesc
)

func file_scanoss_api_scanning_v2_scanoss_scanning_proto_rawDescGZIP() []byte {
	file_scanoss_api_scanning_v2_scanoss_scanning_proto_rawDescOnce.Do(func() {
		file_scanoss_api_scanning_v2_scanoss_scanning_proto_rawDescData = protoimpl.X.CompressGZIP(file_scanoss_api_scanning_v2_scanoss_scanning_proto_rawDescData)
	})
	return file_scanoss_api_scanning_v2_scanoss_scanning_proto_rawDescData
}

var file_scanoss_api_scanning_v2_scanoss_scanning_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_scanoss_api_scanning_v2_scanoss_scanning_proto_goTypes = []interface{}{
	(*HFHRequest)(nil),              // 0: scanoss.api.scanning.v2.HFHRequest
	(*HFHResponse)(nil),             // 1: scanoss.api.scanning.v2.HFHResponse
	(*HFHRequest_Children)(nil),     // 2: scanoss.api.scanning.v2.HFHRequest.Children
	(*HFHResponse_Component)(nil),   // 3: scanoss.api.scanning.v2.HFHResponse.Component
	(*HFHResponse_Result)(nil),      // 4: scanoss.api.scanning.v2.HFHResponse.Result
	(*commonv2.StatusResponse)(nil), // 5: scanoss.api.common.v2.StatusResponse
	(*commonv2.EchoRequest)(nil),    // 6: scanoss.api.common.v2.EchoRequest
	(*commonv2.EchoResponse)(nil),   // 7: scanoss.api.common.v2.EchoResponse
}
var file_scanoss_api_scanning_v2_scanoss_scanning_proto_depIdxs = []int32{
	2, // 0: scanoss.api.scanning.v2.HFHRequest.root:type_name -> scanoss.api.scanning.v2.HFHRequest.Children
	4, // 1: scanoss.api.scanning.v2.HFHResponse.results:type_name -> scanoss.api.scanning.v2.HFHResponse.Result
	5, // 2: scanoss.api.scanning.v2.HFHResponse.status:type_name -> scanoss.api.common.v2.StatusResponse
	2, // 3: scanoss.api.scanning.v2.HFHRequest.Children.children:type_name -> scanoss.api.scanning.v2.HFHRequest.Children
	3, // 4: scanoss.api.scanning.v2.HFHResponse.Result.components:type_name -> scanoss.api.scanning.v2.HFHResponse.Component
	6, // 5: scanoss.api.scanning.v2.Scanning.Echo:input_type -> scanoss.api.common.v2.EchoRequest
	0, // 6: scanoss.api.scanning.v2.Scanning.FolderHashScan:input_type -> scanoss.api.scanning.v2.HFHRequest
	7, // 7: scanoss.api.scanning.v2.Scanning.Echo:output_type -> scanoss.api.common.v2.EchoResponse
	1, // 8: scanoss.api.scanning.v2.Scanning.FolderHashScan:output_type -> scanoss.api.scanning.v2.HFHResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_scanoss_api_scanning_v2_scanoss_scanning_proto_init() }
func file_scanoss_api_scanning_v2_scanoss_scanning_proto_init() {
	if File_scanoss_api_scanning_v2_scanoss_scanning_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scanoss_api_scanning_v2_scanoss_scanning_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HFHRequest); i {
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
		file_scanoss_api_scanning_v2_scanoss_scanning_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HFHResponse); i {
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
		file_scanoss_api_scanning_v2_scanoss_scanning_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HFHRequest_Children); i {
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
		file_scanoss_api_scanning_v2_scanoss_scanning_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HFHResponse_Component); i {
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
		file_scanoss_api_scanning_v2_scanoss_scanning_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HFHResponse_Result); i {
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
			RawDescriptor: file_scanoss_api_scanning_v2_scanoss_scanning_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_scanoss_api_scanning_v2_scanoss_scanning_proto_goTypes,
		DependencyIndexes: file_scanoss_api_scanning_v2_scanoss_scanning_proto_depIdxs,
		MessageInfos:      file_scanoss_api_scanning_v2_scanoss_scanning_proto_msgTypes,
	}.Build()
	File_scanoss_api_scanning_v2_scanoss_scanning_proto = out.File
	file_scanoss_api_scanning_v2_scanoss_scanning_proto_rawDesc = nil
	file_scanoss_api_scanning_v2_scanoss_scanning_proto_goTypes = nil
	file_scanoss_api_scanning_v2_scanoss_scanning_proto_depIdxs = nil
}
