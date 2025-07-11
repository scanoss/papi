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
// Component definition details
//*

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: scanoss/api/components/v2/scanoss-components.proto

package componentsv2

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	commonv2 "github.com/scanoss/papi/api/commonv2"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Component Search request data (JSON payload)
type CompSearchRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Component to search for (vendor/component/purl). Search is exclusive (will override vendor/component)
	Search string `protobuf:"bytes,1,opt,name=search,proto3" json:"search,omitempty"`
	// Vendor to search for. Vendor can be combined with Component
	Vendor string `protobuf:"bytes,2,opt,name=vendor,proto3" json:"vendor,omitempty"`
	// Component to search for. Component can be combined with Vendor
	Component string `protobuf:"bytes,3,opt,name=component,proto3" json:"component,omitempty"`
	// Type of package (purl type) to search for. i.e. github, maven, golang, npm, all - default github
	Package string `protobuf:"bytes,4,opt,name=package,proto3" json:"package,omitempty"`
	// Number of matches to return - default 50
	Limit int32 `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	// Offset to submit to return the next (limit) of component matches
	Offset        int32 `protobuf:"varint,7,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CompSearchRequest) Reset() {
	*x = CompSearchRequest{}
	mi := &file_scanoss_api_components_v2_scanoss_components_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompSearchRequest) ProtoMessage() {}

func (x *CompSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_components_v2_scanoss_components_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompSearchRequest.ProtoReflect.Descriptor instead.
func (*CompSearchRequest) Descriptor() ([]byte, []int) {
	return file_scanoss_api_components_v2_scanoss_components_proto_rawDescGZIP(), []int{0}
}

func (x *CompSearchRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *CompSearchRequest) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *CompSearchRequest) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *CompSearchRequest) GetPackage() string {
	if x != nil {
		return x.Package
	}
	return ""
}

func (x *CompSearchRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *CompSearchRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type CompStatistic struct {
	state            protoimpl.MessageState    `protogen:"open.v1"`
	TotalSourceFiles int32                     `protobuf:"varint,1,opt,name=total_source_files,json=totalSourceFiles,proto3" json:"total_source_files,omitempty"`
	TotalLines       int32                     `protobuf:"varint,2,opt,name=total_lines,json=totalLines,proto3" json:"total_lines,omitempty"`
	TotalBlankLines  int32                     `protobuf:"varint,3,opt,name=total_blank_lines,json=totalBlankLines,proto3" json:"total_blank_lines,omitempty"`
	Languages        []*CompStatistic_Language `protobuf:"bytes,4,rep,name=languages,proto3" json:"languages,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *CompStatistic) Reset() {
	*x = CompStatistic{}
	mi := &file_scanoss_api_components_v2_scanoss_components_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompStatistic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompStatistic) ProtoMessage() {}

func (x *CompStatistic) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_components_v2_scanoss_components_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompStatistic.ProtoReflect.Descriptor instead.
func (*CompStatistic) Descriptor() ([]byte, []int) {
	return file_scanoss_api_components_v2_scanoss_components_proto_rawDescGZIP(), []int{1}
}

func (x *CompStatistic) GetTotalSourceFiles() int32 {
	if x != nil {
		return x.TotalSourceFiles
	}
	return 0
}

func (x *CompStatistic) GetTotalLines() int32 {
	if x != nil {
		return x.TotalLines
	}
	return 0
}

func (x *CompStatistic) GetTotalBlankLines() int32 {
	if x != nil {
		return x.TotalBlankLines
	}
	return 0
}

func (x *CompStatistic) GetLanguages() []*CompStatistic_Language {
	if x != nil {
		return x.Languages
	}
	return nil
}

// Component Statistics response data (JSON payload)
type CompStatisticResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Component statistic details
	Purls []*CompStatisticResponse_Purls `protobuf:"bytes,1,rep,name=purls,proto3" json:"purls,omitempty"`
	// Response status
	Status        *commonv2.StatusResponse `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CompStatisticResponse) Reset() {
	*x = CompStatisticResponse{}
	mi := &file_scanoss_api_components_v2_scanoss_components_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompStatisticResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompStatisticResponse) ProtoMessage() {}

func (x *CompStatisticResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_components_v2_scanoss_components_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompStatisticResponse.ProtoReflect.Descriptor instead.
func (*CompStatisticResponse) Descriptor() ([]byte, []int) {
	return file_scanoss_api_components_v2_scanoss_components_proto_rawDescGZIP(), []int{2}
}

func (x *CompStatisticResponse) GetPurls() []*CompStatisticResponse_Purls {
	if x != nil {
		return x.Purls
	}
	return nil
}

func (x *CompStatisticResponse) GetStatus() *commonv2.StatusResponse {
	if x != nil {
		return x.Status
	}
	return nil
}

// Component Search response data (JSON payload)
type CompSearchResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Component response details
	Components []*CompSearchResponse_Component `protobuf:"bytes,1,rep,name=components,proto3" json:"components,omitempty"`
	// Response status (required?)
	Status        *commonv2.StatusResponse `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CompSearchResponse) Reset() {
	*x = CompSearchResponse{}
	mi := &file_scanoss_api_components_v2_scanoss_components_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompSearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompSearchResponse) ProtoMessage() {}

func (x *CompSearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_components_v2_scanoss_components_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompSearchResponse.ProtoReflect.Descriptor instead.
func (*CompSearchResponse) Descriptor() ([]byte, []int) {
	return file_scanoss_api_components_v2_scanoss_components_proto_rawDescGZIP(), []int{3}
}

func (x *CompSearchResponse) GetComponents() []*CompSearchResponse_Component {
	if x != nil {
		return x.Components
	}
	return nil
}

func (x *CompSearchResponse) GetStatus() *commonv2.StatusResponse {
	if x != nil {
		return x.Status
	}
	return nil
}

// Component Version request data (JSON payload)
type CompVersionRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Component to search for (purl)
	Purl string `protobuf:"bytes,1,opt,name=purl,proto3" json:"purl,omitempty"`
	// Number of versions to return - default 50
	Limit         int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CompVersionRequest) Reset() {
	*x = CompVersionRequest{}
	mi := &file_scanoss_api_components_v2_scanoss_components_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompVersionRequest) ProtoMessage() {}

func (x *CompVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_components_v2_scanoss_components_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompVersionRequest.ProtoReflect.Descriptor instead.
func (*CompVersionRequest) Descriptor() ([]byte, []int) {
	return file_scanoss_api_components_v2_scanoss_components_proto_rawDescGZIP(), []int{4}
}

func (x *CompVersionRequest) GetPurl() string {
	if x != nil {
		return x.Purl
	}
	return ""
}

func (x *CompVersionRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

// Component Search response data (JSON payload)
type CompVersionResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Component response details
	Component *CompVersionResponse_Component `protobuf:"bytes,1,opt,name=component,proto3" json:"component,omitempty"`
	// Response status
	Status        *commonv2.StatusResponse `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CompVersionResponse) Reset() {
	*x = CompVersionResponse{}
	mi := &file_scanoss_api_components_v2_scanoss_components_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompVersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompVersionResponse) ProtoMessage() {}

func (x *CompVersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_components_v2_scanoss_components_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompVersionResponse.ProtoReflect.Descriptor instead.
func (*CompVersionResponse) Descriptor() ([]byte, []int) {
	return file_scanoss_api_components_v2_scanoss_components_proto_rawDescGZIP(), []int{5}
}

func (x *CompVersionResponse) GetComponent() *CompVersionResponse_Component {
	if x != nil {
		return x.Component
	}
	return nil
}

func (x *CompVersionResponse) GetStatus() *commonv2.StatusResponse {
	if x != nil {
		return x.Status
	}
	return nil
}

type CompStatistic_Language struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Files         int32                  `protobuf:"varint,2,opt,name=files,proto3" json:"files,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CompStatistic_Language) Reset() {
	*x = CompStatistic_Language{}
	mi := &file_scanoss_api_components_v2_scanoss_components_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompStatistic_Language) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompStatistic_Language) ProtoMessage() {}

func (x *CompStatistic_Language) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_components_v2_scanoss_components_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompStatistic_Language.ProtoReflect.Descriptor instead.
func (*CompStatistic_Language) Descriptor() ([]byte, []int) {
	return file_scanoss_api_components_v2_scanoss_components_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CompStatistic_Language) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CompStatistic_Language) GetFiles() int32 {
	if x != nil {
		return x.Files
	}
	return 0
}

type CompStatisticResponse_Purls struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Purl          string                 `protobuf:"bytes,1,opt,name=purl,proto3" json:"purl,omitempty"`
	Version       string                 `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Statistics    *CompStatistic         `protobuf:"bytes,3,opt,name=statistics,proto3" json:"statistics,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CompStatisticResponse_Purls) Reset() {
	*x = CompStatisticResponse_Purls{}
	mi := &file_scanoss_api_components_v2_scanoss_components_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompStatisticResponse_Purls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompStatisticResponse_Purls) ProtoMessage() {}

func (x *CompStatisticResponse_Purls) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_components_v2_scanoss_components_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompStatisticResponse_Purls.ProtoReflect.Descriptor instead.
func (*CompStatisticResponse_Purls) Descriptor() ([]byte, []int) {
	return file_scanoss_api_components_v2_scanoss_components_proto_rawDescGZIP(), []int{2, 0}
}

func (x *CompStatisticResponse_Purls) GetPurl() string {
	if x != nil {
		return x.Purl
	}
	return ""
}

func (x *CompStatisticResponse_Purls) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *CompStatisticResponse_Purls) GetStatistics() *CompStatistic {
	if x != nil {
		return x.Statistics
	}
	return nil
}

// Component details
type CompSearchResponse_Component struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Component     string                 `protobuf:"bytes,1,opt,name=component,proto3" json:"component,omitempty"`
	Purl          string                 `protobuf:"bytes,2,opt,name=purl,proto3" json:"purl,omitempty"`
	Url           string                 `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CompSearchResponse_Component) Reset() {
	*x = CompSearchResponse_Component{}
	mi := &file_scanoss_api_components_v2_scanoss_components_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompSearchResponse_Component) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompSearchResponse_Component) ProtoMessage() {}

func (x *CompSearchResponse_Component) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_components_v2_scanoss_components_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompSearchResponse_Component.ProtoReflect.Descriptor instead.
func (*CompSearchResponse_Component) Descriptor() ([]byte, []int) {
	return file_scanoss_api_components_v2_scanoss_components_proto_rawDescGZIP(), []int{3, 0}
}

func (x *CompSearchResponse_Component) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *CompSearchResponse_Component) GetPurl() string {
	if x != nil {
		return x.Purl
	}
	return ""
}

func (x *CompSearchResponse_Component) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// License details
type CompVersionResponse_License struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Name           string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	SpdxId         string                 `protobuf:"bytes,2,opt,name=spdx_id,json=spdxId,proto3" json:"spdx_id,omitempty"`
	IsSpdxApproved bool                   `protobuf:"varint,3,opt,name=is_spdx_approved,json=isSpdxApproved,proto3" json:"is_spdx_approved,omitempty"`
	Url            string                 `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *CompVersionResponse_License) Reset() {
	*x = CompVersionResponse_License{}
	mi := &file_scanoss_api_components_v2_scanoss_components_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompVersionResponse_License) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompVersionResponse_License) ProtoMessage() {}

func (x *CompVersionResponse_License) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_components_v2_scanoss_components_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompVersionResponse_License.ProtoReflect.Descriptor instead.
func (*CompVersionResponse_License) Descriptor() ([]byte, []int) {
	return file_scanoss_api_components_v2_scanoss_components_proto_rawDescGZIP(), []int{5, 0}
}

func (x *CompVersionResponse_License) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CompVersionResponse_License) GetSpdxId() string {
	if x != nil {
		return x.SpdxId
	}
	return ""
}

func (x *CompVersionResponse_License) GetIsSpdxApproved() bool {
	if x != nil {
		return x.IsSpdxApproved
	}
	return false
}

func (x *CompVersionResponse_License) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// Version details (including license)
type CompVersionResponse_Version struct {
	state         protoimpl.MessageState         `protogen:"open.v1"`
	Version       string                         `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Licenses      []*CompVersionResponse_License `protobuf:"bytes,4,rep,name=licenses,proto3" json:"licenses,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CompVersionResponse_Version) Reset() {
	*x = CompVersionResponse_Version{}
	mi := &file_scanoss_api_components_v2_scanoss_components_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompVersionResponse_Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompVersionResponse_Version) ProtoMessage() {}

func (x *CompVersionResponse_Version) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_components_v2_scanoss_components_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompVersionResponse_Version.ProtoReflect.Descriptor instead.
func (*CompVersionResponse_Version) Descriptor() ([]byte, []int) {
	return file_scanoss_api_components_v2_scanoss_components_proto_rawDescGZIP(), []int{5, 1}
}

func (x *CompVersionResponse_Version) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *CompVersionResponse_Version) GetLicenses() []*CompVersionResponse_License {
	if x != nil {
		return x.Licenses
	}
	return nil
}

// Component details (including versions)
type CompVersionResponse_Component struct {
	state         protoimpl.MessageState         `protogen:"open.v1"`
	Component     string                         `protobuf:"bytes,1,opt,name=component,proto3" json:"component,omitempty"`
	Purl          string                         `protobuf:"bytes,2,opt,name=purl,proto3" json:"purl,omitempty"`
	Url           string                         `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Versions      []*CompVersionResponse_Version `protobuf:"bytes,4,rep,name=versions,proto3" json:"versions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CompVersionResponse_Component) Reset() {
	*x = CompVersionResponse_Component{}
	mi := &file_scanoss_api_components_v2_scanoss_components_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompVersionResponse_Component) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompVersionResponse_Component) ProtoMessage() {}

func (x *CompVersionResponse_Component) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_components_v2_scanoss_components_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompVersionResponse_Component.ProtoReflect.Descriptor instead.
func (*CompVersionResponse_Component) Descriptor() ([]byte, []int) {
	return file_scanoss_api_components_v2_scanoss_components_proto_rawDescGZIP(), []int{5, 2}
}

func (x *CompVersionResponse_Component) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *CompVersionResponse_Component) GetPurl() string {
	if x != nil {
		return x.Purl
	}
	return ""
}

func (x *CompVersionResponse_Component) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CompVersionResponse_Component) GetVersions() []*CompVersionResponse_Version {
	if x != nil {
		return x.Versions
	}
	return nil
}

var File_scanoss_api_components_v2_scanoss_components_proto protoreflect.FileDescriptor

const file_scanoss_api_components_v2_scanoss_components_proto_rawDesc = "" +
	"\n" +
	"2scanoss/api/components/v2/scanoss-components.proto\x12\x19scanoss.api.components.v2\x1a*scanoss/api/common/v2/scanoss-common.proto\x1a\x1cgoogle/api/annotations.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\xa9\x01\n" +
	"\x11CompSearchRequest\x12\x16\n" +
	"\x06search\x18\x01 \x01(\tR\x06search\x12\x16\n" +
	"\x06vendor\x18\x02 \x01(\tR\x06vendor\x12\x1c\n" +
	"\tcomponent\x18\x03 \x01(\tR\tcomponent\x12\x18\n" +
	"\apackage\x18\x04 \x01(\tR\apackage\x12\x14\n" +
	"\x05limit\x18\x06 \x01(\x05R\x05limit\x12\x16\n" +
	"\x06offset\x18\a \x01(\x05R\x06offset\"\x91\x02\n" +
	"\rCompStatistic\x12,\n" +
	"\x12total_source_files\x18\x01 \x01(\x05R\x10totalSourceFiles\x12\x1f\n" +
	"\vtotal_lines\x18\x02 \x01(\x05R\n" +
	"totalLines\x12*\n" +
	"\x11total_blank_lines\x18\x03 \x01(\x05R\x0ftotalBlankLines\x12O\n" +
	"\tlanguages\x18\x04 \x03(\v21.scanoss.api.components.v2.CompStatistic.LanguageR\tlanguages\x1a4\n" +
	"\bLanguage\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x14\n" +
	"\x05files\x18\x02 \x01(\x05R\x05files\"\xa5\x02\n" +
	"\x15CompStatisticResponse\x12L\n" +
	"\x05purls\x18\x01 \x03(\v26.scanoss.api.components.v2.CompStatisticResponse.PurlsR\x05purls\x12=\n" +
	"\x06status\x18\x02 \x01(\v2%.scanoss.api.common.v2.StatusResponseR\x06status\x1a\x7f\n" +
	"\x05Purls\x12\x12\n" +
	"\x04purl\x18\x01 \x01(\tR\x04purl\x12\x18\n" +
	"\aversion\x18\x02 \x01(\tR\aversion\x12H\n" +
	"\n" +
	"statistics\x18\x03 \x01(\v2(.scanoss.api.components.v2.CompStatisticR\n" +
	"statistics\"\xfd\x01\n" +
	"\x12CompSearchResponse\x12W\n" +
	"\n" +
	"components\x18\x01 \x03(\v27.scanoss.api.components.v2.CompSearchResponse.ComponentR\n" +
	"components\x12=\n" +
	"\x06status\x18\x02 \x01(\v2%.scanoss.api.common.v2.StatusResponseR\x06status\x1aO\n" +
	"\tComponent\x12\x1c\n" +
	"\tcomponent\x18\x01 \x01(\tR\tcomponent\x12\x12\n" +
	"\x04purl\x18\x02 \x01(\tR\x04purl\x12\x10\n" +
	"\x03url\x18\x03 \x01(\tR\x03url\">\n" +
	"\x12CompVersionRequest\x12\x12\n" +
	"\x04purl\x18\x01 \x01(\tR\x04purl\x12\x14\n" +
	"\x05limit\x18\x02 \x01(\x05R\x05limit\"\xbf\x04\n" +
	"\x13CompVersionResponse\x12V\n" +
	"\tcomponent\x18\x01 \x01(\v28.scanoss.api.components.v2.CompVersionResponse.ComponentR\tcomponent\x12=\n" +
	"\x06status\x18\x02 \x01(\v2%.scanoss.api.common.v2.StatusResponseR\x06status\x1ar\n" +
	"\aLicense\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x17\n" +
	"\aspdx_id\x18\x02 \x01(\tR\x06spdxId\x12(\n" +
	"\x10is_spdx_approved\x18\x03 \x01(\bR\x0eisSpdxApproved\x12\x10\n" +
	"\x03url\x18\x04 \x01(\tR\x03url\x1aw\n" +
	"\aVersion\x12\x18\n" +
	"\aversion\x18\x01 \x01(\tR\aversion\x12R\n" +
	"\blicenses\x18\x04 \x03(\v26.scanoss.api.components.v2.CompVersionResponse.LicenseR\blicenses\x1a\xa3\x01\n" +
	"\tComponent\x12\x1c\n" +
	"\tcomponent\x18\x01 \x01(\tR\tcomponent\x12\x12\n" +
	"\x04purl\x18\x02 \x01(\tR\x04purl\x12\x10\n" +
	"\x03url\x18\x03 \x01(\tR\x03url\x12R\n" +
	"\bversions\x18\x04 \x03(\v26.scanoss.api.components.v2.CompVersionResponse.VersionR\bversions2\xd4\x04\n" +
	"\n" +
	"Components\x12s\n" +
	"\x04Echo\x12\".scanoss.api.common.v2.EchoRequest\x1a#.scanoss.api.common.v2.EchoResponse\"\"\x82\xd3\xe4\x93\x02\x1c:\x01*\"\x17/api/v2/components/echo\x12\x95\x01\n" +
	"\x10SearchComponents\x12,.scanoss.api.components.v2.CompSearchRequest\x1a-.scanoss.api.components.v2.CompSearchResponse\"$\x82\xd3\xe4\x93\x02\x1e:\x01*\"\x19/api/v2/components/search\x12\x9d\x01\n" +
	"\x14GetComponentVersions\x12-.scanoss.api.components.v2.CompVersionRequest\x1a..scanoss.api.components.v2.CompVersionResponse\"&\x82\xd3\xe4\x93\x02 :\x01*\"\x1b/api/v2/components/versions\x12\x98\x01\n" +
	"\x16GetComponentStatistics\x12\".scanoss.api.common.v2.PurlRequest\x1a0.scanoss.api.components.v2.CompStatisticResponse\"(\x82\xd3\xe4\x93\x02\":\x01*\"\x1d/api/v2/components/statisticsB\x9a\x03\x92A\xdf\x02\x12\x9d\x01\n" +
	"\x1aSCANOSS Components Service\x12(Provides component intelligence services\"P\n" +
	"\x12scanoss-components\x12%https://github.com/scanoss/components\x1a\x13support@scanoss.com2\x032.0\x1a\x0fapi.scanoss.com*\x02\x02\x012\x10application/json:\x10application/jsonR;\n" +
	"\x03404\x124\n" +
	"*Returned when the resource does not exist.\x12\x06\n" +
	"\x04\x9a\x02\x01\aZ8\n" +
	"6\n" +
	"\aapi_key\x12+\b\x02\x12\x1aAPI key for authentication\x1a\tx-api-key \x02b\r\n" +
	"\v\n" +
	"\aapi_key\x12\x00Z5github.com/scanoss/papi/api/componentsv2;componentsv2b\x06proto3"

var (
	file_scanoss_api_components_v2_scanoss_components_proto_rawDescOnce sync.Once
	file_scanoss_api_components_v2_scanoss_components_proto_rawDescData []byte
)

func file_scanoss_api_components_v2_scanoss_components_proto_rawDescGZIP() []byte {
	file_scanoss_api_components_v2_scanoss_components_proto_rawDescOnce.Do(func() {
		file_scanoss_api_components_v2_scanoss_components_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_scanoss_api_components_v2_scanoss_components_proto_rawDesc), len(file_scanoss_api_components_v2_scanoss_components_proto_rawDesc)))
	})
	return file_scanoss_api_components_v2_scanoss_components_proto_rawDescData
}

var file_scanoss_api_components_v2_scanoss_components_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_scanoss_api_components_v2_scanoss_components_proto_goTypes = []any{
	(*CompSearchRequest)(nil),             // 0: scanoss.api.components.v2.CompSearchRequest
	(*CompStatistic)(nil),                 // 1: scanoss.api.components.v2.CompStatistic
	(*CompStatisticResponse)(nil),         // 2: scanoss.api.components.v2.CompStatisticResponse
	(*CompSearchResponse)(nil),            // 3: scanoss.api.components.v2.CompSearchResponse
	(*CompVersionRequest)(nil),            // 4: scanoss.api.components.v2.CompVersionRequest
	(*CompVersionResponse)(nil),           // 5: scanoss.api.components.v2.CompVersionResponse
	(*CompStatistic_Language)(nil),        // 6: scanoss.api.components.v2.CompStatistic.Language
	(*CompStatisticResponse_Purls)(nil),   // 7: scanoss.api.components.v2.CompStatisticResponse.Purls
	(*CompSearchResponse_Component)(nil),  // 8: scanoss.api.components.v2.CompSearchResponse.Component
	(*CompVersionResponse_License)(nil),   // 9: scanoss.api.components.v2.CompVersionResponse.License
	(*CompVersionResponse_Version)(nil),   // 10: scanoss.api.components.v2.CompVersionResponse.Version
	(*CompVersionResponse_Component)(nil), // 11: scanoss.api.components.v2.CompVersionResponse.Component
	(*commonv2.StatusResponse)(nil),       // 12: scanoss.api.common.v2.StatusResponse
	(*commonv2.EchoRequest)(nil),          // 13: scanoss.api.common.v2.EchoRequest
	(*commonv2.PurlRequest)(nil),          // 14: scanoss.api.common.v2.PurlRequest
	(*commonv2.EchoResponse)(nil),         // 15: scanoss.api.common.v2.EchoResponse
}
var file_scanoss_api_components_v2_scanoss_components_proto_depIdxs = []int32{
	6,  // 0: scanoss.api.components.v2.CompStatistic.languages:type_name -> scanoss.api.components.v2.CompStatistic.Language
	7,  // 1: scanoss.api.components.v2.CompStatisticResponse.purls:type_name -> scanoss.api.components.v2.CompStatisticResponse.Purls
	12, // 2: scanoss.api.components.v2.CompStatisticResponse.status:type_name -> scanoss.api.common.v2.StatusResponse
	8,  // 3: scanoss.api.components.v2.CompSearchResponse.components:type_name -> scanoss.api.components.v2.CompSearchResponse.Component
	12, // 4: scanoss.api.components.v2.CompSearchResponse.status:type_name -> scanoss.api.common.v2.StatusResponse
	11, // 5: scanoss.api.components.v2.CompVersionResponse.component:type_name -> scanoss.api.components.v2.CompVersionResponse.Component
	12, // 6: scanoss.api.components.v2.CompVersionResponse.status:type_name -> scanoss.api.common.v2.StatusResponse
	1,  // 7: scanoss.api.components.v2.CompStatisticResponse.Purls.statistics:type_name -> scanoss.api.components.v2.CompStatistic
	9,  // 8: scanoss.api.components.v2.CompVersionResponse.Version.licenses:type_name -> scanoss.api.components.v2.CompVersionResponse.License
	10, // 9: scanoss.api.components.v2.CompVersionResponse.Component.versions:type_name -> scanoss.api.components.v2.CompVersionResponse.Version
	13, // 10: scanoss.api.components.v2.Components.Echo:input_type -> scanoss.api.common.v2.EchoRequest
	0,  // 11: scanoss.api.components.v2.Components.SearchComponents:input_type -> scanoss.api.components.v2.CompSearchRequest
	4,  // 12: scanoss.api.components.v2.Components.GetComponentVersions:input_type -> scanoss.api.components.v2.CompVersionRequest
	14, // 13: scanoss.api.components.v2.Components.GetComponentStatistics:input_type -> scanoss.api.common.v2.PurlRequest
	15, // 14: scanoss.api.components.v2.Components.Echo:output_type -> scanoss.api.common.v2.EchoResponse
	3,  // 15: scanoss.api.components.v2.Components.SearchComponents:output_type -> scanoss.api.components.v2.CompSearchResponse
	5,  // 16: scanoss.api.components.v2.Components.GetComponentVersions:output_type -> scanoss.api.components.v2.CompVersionResponse
	2,  // 17: scanoss.api.components.v2.Components.GetComponentStatistics:output_type -> scanoss.api.components.v2.CompStatisticResponse
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_scanoss_api_components_v2_scanoss_components_proto_init() }
func file_scanoss_api_components_v2_scanoss_components_proto_init() {
	if File_scanoss_api_components_v2_scanoss_components_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_scanoss_api_components_v2_scanoss_components_proto_rawDesc), len(file_scanoss_api_components_v2_scanoss_components_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_scanoss_api_components_v2_scanoss_components_proto_goTypes,
		DependencyIndexes: file_scanoss_api_components_v2_scanoss_components_proto_depIdxs,
		MessageInfos:      file_scanoss_api_components_v2_scanoss_components_proto_msgTypes,
	}.Build()
	File_scanoss_api_components_v2_scanoss_components_proto = out.File
	file_scanoss_api_components_v2_scanoss_components_proto_goTypes = nil
	file_scanoss_api_components_v2_scanoss_components_proto_depIdxs = nil
}
