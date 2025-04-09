//
//SPDX-License-Identifier: MIT
//
//Copyright (c) 2022, SCANOSS
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
// Dependency definition details
//*

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.29.3
// source: scanoss/api/dependencies/v2/scanoss-dependencies.proto

package dependenciesv2

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

// Dependency request data (JSON payload)
type DependencyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of dependency files
	Files []*DependencyRequest_Files `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	// Depth to go when searching dependencies
	Depth int32 `protobuf:"varint,2,opt,name=depth,proto3" json:"depth,omitempty"`
}

func (x *DependencyRequest) Reset() {
	*x = DependencyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DependencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DependencyRequest) ProtoMessage() {}

func (x *DependencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DependencyRequest.ProtoReflect.Descriptor instead.
func (*DependencyRequest) Descriptor() ([]byte, []int) {
	return file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_rawDescGZIP(), []int{0}
}

func (x *DependencyRequest) GetFiles() []*DependencyRequest_Files {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *DependencyRequest) GetDepth() int32 {
	if x != nil {
		return x.Depth
	}
	return 0
}

// Dependency response data (JSON payload)
type DependencyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Dependency response details
	Files []*DependencyResponse_Files `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	// Response status (required?)
	Status *commonv2.StatusResponse `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DependencyResponse) Reset() {
	*x = DependencyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DependencyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DependencyResponse) ProtoMessage() {}

func (x *DependencyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DependencyResponse.ProtoReflect.Descriptor instead.
func (*DependencyResponse) Descriptor() ([]byte, []int) {
	return file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_rawDescGZIP(), []int{1}
}

func (x *DependencyResponse) GetFiles() []*DependencyResponse_Files {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *DependencyResponse) GetStatus() *commonv2.StatusResponse {
	if x != nil {
		return x.Status
	}
	return nil
}

// List of Purls/requirements
type DependencyRequest_Purls struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Purl name
	Purl string `protobuf:"bytes,1,opt,name=purl,proto3" json:"purl,omitempty"`
	// requirement - optional
	Requirement string `protobuf:"bytes,2,opt,name=requirement,proto3" json:"requirement,omitempty"` //    string scope = 3;  // TODO what did we want scope for?
}

func (x *DependencyRequest_Purls) Reset() {
	*x = DependencyRequest_Purls{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DependencyRequest_Purls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DependencyRequest_Purls) ProtoMessage() {}

func (x *DependencyRequest_Purls) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DependencyRequest_Purls.ProtoReflect.Descriptor instead.
func (*DependencyRequest_Purls) Descriptor() ([]byte, []int) {
	return file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_rawDescGZIP(), []int{0, 0}
}

func (x *DependencyRequest_Purls) GetPurl() string {
	if x != nil {
		return x.Purl
	}
	return ""
}

func (x *DependencyRequest_Purls) GetRequirement() string {
	if x != nil {
		return x.Requirement
	}
	return ""
}

// List of dependency files
type DependencyRequest_Files struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// dependency filename
	File string `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	// List of purls
	Purls []*DependencyRequest_Purls `protobuf:"bytes,2,rep,name=purls,proto3" json:"purls,omitempty"`
}

func (x *DependencyRequest_Files) Reset() {
	*x = DependencyRequest_Files{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DependencyRequest_Files) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DependencyRequest_Files) ProtoMessage() {}

func (x *DependencyRequest_Files) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DependencyRequest_Files.ProtoReflect.Descriptor instead.
func (*DependencyRequest_Files) Descriptor() ([]byte, []int) {
	return file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_rawDescGZIP(), []int{0, 1}
}

func (x *DependencyRequest_Files) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

func (x *DependencyRequest_Files) GetPurls() []*DependencyRequest_Purls {
	if x != nil {
		return x.Purls
	}
	return nil
}

type DependencyResponse_Licenses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	SpdxId         string `protobuf:"bytes,2,opt,name=spdx_id,json=spdxId,proto3" json:"spdx_id,omitempty"`
	IsSpdxApproved bool   `protobuf:"varint,3,opt,name=is_spdx_approved,json=isSpdxApproved,proto3" json:"is_spdx_approved,omitempty"`
	Url            string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *DependencyResponse_Licenses) Reset() {
	*x = DependencyResponse_Licenses{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DependencyResponse_Licenses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DependencyResponse_Licenses) ProtoMessage() {}

func (x *DependencyResponse_Licenses) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DependencyResponse_Licenses.ProtoReflect.Descriptor instead.
func (*DependencyResponse_Licenses) Descriptor() ([]byte, []int) {
	return file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_rawDescGZIP(), []int{1, 0}
}

func (x *DependencyResponse_Licenses) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DependencyResponse_Licenses) GetSpdxId() string {
	if x != nil {
		return x.SpdxId
	}
	return ""
}

func (x *DependencyResponse_Licenses) GetIsSpdxApproved() bool {
	if x != nil {
		return x.IsSpdxApproved
	}
	return false
}

func (x *DependencyResponse_Licenses) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type DependencyResponse_Dependencies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Component string                         `protobuf:"bytes,1,opt,name=component,proto3" json:"component,omitempty"`
	Purl      string                         `protobuf:"bytes,2,opt,name=purl,proto3" json:"purl,omitempty"`
	Version   string                         `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Licenses  []*DependencyResponse_Licenses `protobuf:"bytes,4,rep,name=licenses,proto3" json:"licenses,omitempty"`
	Url       string                         `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	Comment   string                         `protobuf:"bytes,6,opt,name=comment,proto3" json:"comment,omitempty"` //    string scope = 7;
}

func (x *DependencyResponse_Dependencies) Reset() {
	*x = DependencyResponse_Dependencies{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DependencyResponse_Dependencies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DependencyResponse_Dependencies) ProtoMessage() {}

func (x *DependencyResponse_Dependencies) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DependencyResponse_Dependencies.ProtoReflect.Descriptor instead.
func (*DependencyResponse_Dependencies) Descriptor() ([]byte, []int) {
	return file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_rawDescGZIP(), []int{1, 1}
}

func (x *DependencyResponse_Dependencies) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *DependencyResponse_Dependencies) GetPurl() string {
	if x != nil {
		return x.Purl
	}
	return ""
}

func (x *DependencyResponse_Dependencies) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *DependencyResponse_Dependencies) GetLicenses() []*DependencyResponse_Licenses {
	if x != nil {
		return x.Licenses
	}
	return nil
}

func (x *DependencyResponse_Dependencies) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *DependencyResponse_Dependencies) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type DependencyResponse_Files struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File         string                             `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	Id           string                             `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Status       string                             `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Dependencies []*DependencyResponse_Dependencies `protobuf:"bytes,4,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
}

func (x *DependencyResponse_Files) Reset() {
	*x = DependencyResponse_Files{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DependencyResponse_Files) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DependencyResponse_Files) ProtoMessage() {}

func (x *DependencyResponse_Files) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DependencyResponse_Files.ProtoReflect.Descriptor instead.
func (*DependencyResponse_Files) Descriptor() ([]byte, []int) {
	return file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_rawDescGZIP(), []int{1, 2}
}

func (x *DependencyResponse_Files) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

func (x *DependencyResponse_Files) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DependencyResponse_Files) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *DependencyResponse_Files) GetDependencies() []*DependencyResponse_Dependencies {
	if x != nil {
		return x.Dependencies
	}
	return nil
}

var File_scanoss_api_dependencies_v2_scanoss_dependencies_proto protoreflect.FileDescriptor

var file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_rawDesc = []byte{
	0x0a, 0x36, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65,
	0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x63,
	0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2d, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69,
	0x65, 0x73, 0x2e, 0x76, 0x32, 0x1a, 0x2a, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x63, 0x61,
	0x6e, 0x6f, 0x73, 0x73, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77, 0x61, 0x67,
	0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x02,
	0x0a, 0x11, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x34, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x32,
	0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x64, 0x65, 0x70, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x64, 0x65, 0x70, 0x74, 0x68, 0x1a, 0x3d, 0x0a, 0x05, 0x50, 0x75, 0x72, 0x6c, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x75,
	0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x67, 0x0a, 0x05, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0x4a, 0x0a, 0x05, 0x70, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x34, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64,
	0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x44,
	0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x50, 0x75, 0x72, 0x6c, 0x73, 0x52, 0x05, 0x70, 0x75, 0x72, 0x6c, 0x73, 0x22, 0x9c, 0x05,
	0x0a, 0x12, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x76,
	0x32, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x12, 0x3d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x1a, 0x73, 0x0a, 0x08, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x73, 0x70, 0x64, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x70, 0x64, 0x78, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x73, 0x5f,
	0x73, 0x70, 0x64, 0x78, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x73, 0x53, 0x70, 0x64, 0x78, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x1a, 0xdc, 0x01, 0x0a, 0x0c, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64,
	0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x75, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x54, 0x0a, 0x08, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x2e,
	0x76, 0x32, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x52, 0x08,
	0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x1a, 0xa5, 0x01, 0x0a, 0x05, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x69,
	0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x60, 0x0a, 0x0c, 0x64, 0x65,
	0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x3c, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64,
	0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x44,
	0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x0c,
	0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x32, 0xa8, 0x02, 0x0a,
	0x0c, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x75, 0x0a,
	0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x22, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x63,
	0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x63, 0x61, 0x6e,
	0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x32, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f,
	0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x65, 0x63, 0x68,
	0x6f, 0x3a, 0x01, 0x2a, 0x12, 0xa0, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x65,
	0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x2e, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f,
	0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63,
	0x69, 0x65, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f,
	0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63,
	0x69, 0x65, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x26, 0x22, 0x21, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x64, 0x65, 0x70, 0x65, 0x6e,
	0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e,
	0x63, 0x69, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x9c, 0x02, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2f, 0x70,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e,
	0x63, 0x69, 0x65, 0x73, 0x76, 0x32, 0x3b, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63,
	0x69, 0x65, 0x73, 0x76, 0x32, 0x92, 0x41, 0xdd, 0x01, 0x12, 0x77, 0x0a, 0x1a, 0x53, 0x43, 0x41,
	0x4e, 0x4f, 0x53, 0x53, 0x20, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x20,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x54, 0x0a, 0x14, 0x73, 0x63, 0x61, 0x6e, 0x6f,
	0x73, 0x73, 0x2d, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12,
	0x27, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x65,
	0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x1a, 0x13, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72,
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
	file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_rawDescOnce sync.Once
	file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_rawDescData = file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_rawDesc
)

func file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_rawDescGZIP() []byte {
	file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_rawDescOnce.Do(func() {
		file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_rawDescData = protoimpl.X.CompressGZIP(file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_rawDescData)
	})
	return file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_rawDescData
}

var file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_goTypes = []interface{}{
	(*DependencyRequest)(nil),               // 0: scanoss.api.dependencies.v2.DependencyRequest
	(*DependencyResponse)(nil),              // 1: scanoss.api.dependencies.v2.DependencyResponse
	(*DependencyRequest_Purls)(nil),         // 2: scanoss.api.dependencies.v2.DependencyRequest.Purls
	(*DependencyRequest_Files)(nil),         // 3: scanoss.api.dependencies.v2.DependencyRequest.Files
	(*DependencyResponse_Licenses)(nil),     // 4: scanoss.api.dependencies.v2.DependencyResponse.Licenses
	(*DependencyResponse_Dependencies)(nil), // 5: scanoss.api.dependencies.v2.DependencyResponse.Dependencies
	(*DependencyResponse_Files)(nil),        // 6: scanoss.api.dependencies.v2.DependencyResponse.Files
	(*commonv2.StatusResponse)(nil),         // 7: scanoss.api.common.v2.StatusResponse
	(*commonv2.EchoRequest)(nil),            // 8: scanoss.api.common.v2.EchoRequest
	(*commonv2.EchoResponse)(nil),           // 9: scanoss.api.common.v2.EchoResponse
}
var file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_depIdxs = []int32{
	3, // 0: scanoss.api.dependencies.v2.DependencyRequest.files:type_name -> scanoss.api.dependencies.v2.DependencyRequest.Files
	6, // 1: scanoss.api.dependencies.v2.DependencyResponse.files:type_name -> scanoss.api.dependencies.v2.DependencyResponse.Files
	7, // 2: scanoss.api.dependencies.v2.DependencyResponse.status:type_name -> scanoss.api.common.v2.StatusResponse
	2, // 3: scanoss.api.dependencies.v2.DependencyRequest.Files.purls:type_name -> scanoss.api.dependencies.v2.DependencyRequest.Purls
	4, // 4: scanoss.api.dependencies.v2.DependencyResponse.Dependencies.licenses:type_name -> scanoss.api.dependencies.v2.DependencyResponse.Licenses
	5, // 5: scanoss.api.dependencies.v2.DependencyResponse.Files.dependencies:type_name -> scanoss.api.dependencies.v2.DependencyResponse.Dependencies
	8, // 6: scanoss.api.dependencies.v2.Dependencies.Echo:input_type -> scanoss.api.common.v2.EchoRequest
	0, // 7: scanoss.api.dependencies.v2.Dependencies.GetDependencies:input_type -> scanoss.api.dependencies.v2.DependencyRequest
	9, // 8: scanoss.api.dependencies.v2.Dependencies.Echo:output_type -> scanoss.api.common.v2.EchoResponse
	1, // 9: scanoss.api.dependencies.v2.Dependencies.GetDependencies:output_type -> scanoss.api.dependencies.v2.DependencyResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_init() }
func file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_init() {
	if File_scanoss_api_dependencies_v2_scanoss_dependencies_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DependencyRequest); i {
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
		file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DependencyResponse); i {
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
		file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DependencyRequest_Purls); i {
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
		file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DependencyRequest_Files); i {
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
		file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DependencyResponse_Licenses); i {
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
		file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DependencyResponse_Dependencies); i {
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
		file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DependencyResponse_Files); i {
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
			RawDescriptor: file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_goTypes,
		DependencyIndexes: file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_depIdxs,
		MessageInfos:      file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_msgTypes,
	}.Build()
	File_scanoss_api_dependencies_v2_scanoss_dependencies_proto = out.File
	file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_rawDesc = nil
	file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_goTypes = nil
	file_scanoss_api_dependencies_v2_scanoss_dependencies_proto_depIdxs = nil
}
