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
// Vulnerability definition details
//*

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: scanoss/api/vulnerabilities/v2/scanoss-vulnerabilities.proto

package vulnerabilitiesv2

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

// Vulnerability request data (JSON payload)
type VulnerabilityRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// JSON vulnerability request purls
	Purls         []*VulnerabilityRequest_Purls `protobuf:"bytes,1,rep,name=purls,proto3" json:"purls,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VulnerabilityRequest) Reset() {
	*x = VulnerabilityRequest{}
	mi := &file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VulnerabilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VulnerabilityRequest) ProtoMessage() {}

func (x *VulnerabilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VulnerabilityRequest.ProtoReflect.Descriptor instead.
func (*VulnerabilityRequest) Descriptor() ([]byte, []int) {
	return file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDescGZIP(), []int{0}
}

func (x *VulnerabilityRequest) GetPurls() []*VulnerabilityRequest_Purls {
	if x != nil {
		return x.Purls
	}
	return nil
}

// Vulnerability CPE response data (JSON payload)
type CpeResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Purl/CPE details
	Purls []*CpeResponse_Purls `protobuf:"bytes,1,rep,name=purls,proto3" json:"purls,omitempty"`
	// Response status
	Status        *commonv2.StatusResponse `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CpeResponse) Reset() {
	*x = CpeResponse{}
	mi := &file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CpeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CpeResponse) ProtoMessage() {}

func (x *CpeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CpeResponse.ProtoReflect.Descriptor instead.
func (*CpeResponse) Descriptor() ([]byte, []int) {
	return file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDescGZIP(), []int{1}
}

func (x *CpeResponse) GetPurls() []*CpeResponse_Purls {
	if x != nil {
		return x.Purls
	}
	return nil
}

func (x *CpeResponse) GetStatus() *commonv2.StatusResponse {
	if x != nil {
		return x.Status
	}
	return nil
}

// Vulnerability response data (JSON payload)
type VulnerabilityResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Vulnerability response details
	Purls []*VulnerabilityResponse_Purls `protobuf:"bytes,1,rep,name=purls,proto3" json:"purls,omitempty"`
	// Response status
	Status        *commonv2.StatusResponse `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VulnerabilityResponse) Reset() {
	*x = VulnerabilityResponse{}
	mi := &file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VulnerabilityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VulnerabilityResponse) ProtoMessage() {}

func (x *VulnerabilityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VulnerabilityResponse.ProtoReflect.Descriptor instead.
func (*VulnerabilityResponse) Descriptor() ([]byte, []int) {
	return file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDescGZIP(), []int{2}
}

func (x *VulnerabilityResponse) GetPurls() []*VulnerabilityResponse_Purls {
	if x != nil {
		return x.Purls
	}
	return nil
}

func (x *VulnerabilityResponse) GetStatus() *commonv2.StatusResponse {
	if x != nil {
		return x.Status
	}
	return nil
}

type VulnerabilityRequest_Purls struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Purl          string                 `protobuf:"bytes,1,opt,name=purl,proto3" json:"purl,omitempty"`
	Requirement   string                 `protobuf:"bytes,2,opt,name=requirement,proto3" json:"requirement,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VulnerabilityRequest_Purls) Reset() {
	*x = VulnerabilityRequest_Purls{}
	mi := &file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VulnerabilityRequest_Purls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VulnerabilityRequest_Purls) ProtoMessage() {}

func (x *VulnerabilityRequest_Purls) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VulnerabilityRequest_Purls.ProtoReflect.Descriptor instead.
func (*VulnerabilityRequest_Purls) Descriptor() ([]byte, []int) {
	return file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDescGZIP(), []int{0, 0}
}

func (x *VulnerabilityRequest_Purls) GetPurl() string {
	if x != nil {
		return x.Purl
	}
	return ""
}

func (x *VulnerabilityRequest_Purls) GetRequirement() string {
	if x != nil {
		return x.Requirement
	}
	return ""
}

type CpeResponse_Purls struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Purl          string                 `protobuf:"bytes,1,opt,name=purl,proto3" json:"purl,omitempty"`
	Cpes          []string               `protobuf:"bytes,2,rep,name=cpes,proto3" json:"cpes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CpeResponse_Purls) Reset() {
	*x = CpeResponse_Purls{}
	mi := &file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CpeResponse_Purls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CpeResponse_Purls) ProtoMessage() {}

func (x *CpeResponse_Purls) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CpeResponse_Purls.ProtoReflect.Descriptor instead.
func (*CpeResponse_Purls) Descriptor() ([]byte, []int) {
	return file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CpeResponse_Purls) GetPurl() string {
	if x != nil {
		return x.Purl
	}
	return ""
}

func (x *CpeResponse_Purls) GetCpes() []string {
	if x != nil {
		return x.Cpes
	}
	return nil
}

type VulnerabilityResponse_Vulnerabilities struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Cve           string                 `protobuf:"bytes,2,opt,name=cve,proto3" json:"cve,omitempty"`
	Url           string                 `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Summary       string                 `protobuf:"bytes,4,opt,name=summary,proto3" json:"summary,omitempty"`
	Severity      string                 `protobuf:"bytes,5,opt,name=severity,proto3" json:"severity,omitempty"`
	Published     string                 `protobuf:"bytes,6,opt,name=published,proto3" json:"published,omitempty"`
	Modified      string                 `protobuf:"bytes,7,opt,name=modified,proto3" json:"modified,omitempty"`
	Source        string                 `protobuf:"bytes,8,opt,name=source,proto3" json:"source,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VulnerabilityResponse_Vulnerabilities) Reset() {
	*x = VulnerabilityResponse_Vulnerabilities{}
	mi := &file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VulnerabilityResponse_Vulnerabilities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VulnerabilityResponse_Vulnerabilities) ProtoMessage() {}

func (x *VulnerabilityResponse_Vulnerabilities) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VulnerabilityResponse_Vulnerabilities.ProtoReflect.Descriptor instead.
func (*VulnerabilityResponse_Vulnerabilities) Descriptor() ([]byte, []int) {
	return file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDescGZIP(), []int{2, 0}
}

func (x *VulnerabilityResponse_Vulnerabilities) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *VulnerabilityResponse_Vulnerabilities) GetCve() string {
	if x != nil {
		return x.Cve
	}
	return ""
}

func (x *VulnerabilityResponse_Vulnerabilities) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *VulnerabilityResponse_Vulnerabilities) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *VulnerabilityResponse_Vulnerabilities) GetSeverity() string {
	if x != nil {
		return x.Severity
	}
	return ""
}

func (x *VulnerabilityResponse_Vulnerabilities) GetPublished() string {
	if x != nil {
		return x.Published
	}
	return ""
}

func (x *VulnerabilityResponse_Vulnerabilities) GetModified() string {
	if x != nil {
		return x.Modified
	}
	return ""
}

func (x *VulnerabilityResponse_Vulnerabilities) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

type VulnerabilityResponse_Purls struct {
	state           protoimpl.MessageState                   `protogen:"open.v1"`
	Purl            string                                   `protobuf:"bytes,1,opt,name=purl,proto3" json:"purl,omitempty"`
	Vulnerabilities []*VulnerabilityResponse_Vulnerabilities `protobuf:"bytes,2,rep,name=vulnerabilities,proto3" json:"vulnerabilities,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *VulnerabilityResponse_Purls) Reset() {
	*x = VulnerabilityResponse_Purls{}
	mi := &file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VulnerabilityResponse_Purls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VulnerabilityResponse_Purls) ProtoMessage() {}

func (x *VulnerabilityResponse_Purls) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VulnerabilityResponse_Purls.ProtoReflect.Descriptor instead.
func (*VulnerabilityResponse_Purls) Descriptor() ([]byte, []int) {
	return file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDescGZIP(), []int{2, 1}
}

func (x *VulnerabilityResponse_Purls) GetPurl() string {
	if x != nil {
		return x.Purl
	}
	return ""
}

func (x *VulnerabilityResponse_Purls) GetVulnerabilities() []*VulnerabilityResponse_Vulnerabilities {
	if x != nil {
		return x.Vulnerabilities
	}
	return nil
}

var File_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto protoreflect.FileDescriptor

const file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDesc = "" +
	"\n" +
	"<scanoss/api/vulnerabilities/v2/scanoss-vulnerabilities.proto\x12\x1escanoss.api.vulnerabilities.v2\x1a*scanoss/api/common/v2/scanoss-common.proto\x1a\x1cgoogle/api/annotations.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\xa7\x01\n" +
	"\x14VulnerabilityRequest\x12P\n" +
	"\x05purls\x18\x01 \x03(\v2:.scanoss.api.vulnerabilities.v2.VulnerabilityRequest.PurlsR\x05purls\x1a=\n" +
	"\x05Purls\x12\x12\n" +
	"\x04purl\x18\x01 \x01(\tR\x04purl\x12 \n" +
	"\vrequirement\x18\x02 \x01(\tR\vrequirement\"\xc6\x01\n" +
	"\vCpeResponse\x12G\n" +
	"\x05purls\x18\x01 \x03(\v21.scanoss.api.vulnerabilities.v2.CpeResponse.PurlsR\x05purls\x12=\n" +
	"\x06status\x18\x02 \x01(\v2%.scanoss.api.common.v2.StatusResponseR\x06status\x1a/\n" +
	"\x05Purls\x12\x12\n" +
	"\x04purl\x18\x01 \x01(\tR\x04purl\x12\x12\n" +
	"\x04cpes\x18\x02 \x03(\tR\x04cpes\"\x88\x04\n" +
	"\x15VulnerabilityResponse\x12Q\n" +
	"\x05purls\x18\x01 \x03(\v2;.scanoss.api.vulnerabilities.v2.VulnerabilityResponse.PurlsR\x05purls\x12=\n" +
	"\x06status\x18\x02 \x01(\v2%.scanoss.api.common.v2.StatusResponseR\x06status\x1a\xcd\x01\n" +
	"\x0fVulnerabilities\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x10\n" +
	"\x03cve\x18\x02 \x01(\tR\x03cve\x12\x10\n" +
	"\x03url\x18\x03 \x01(\tR\x03url\x12\x18\n" +
	"\asummary\x18\x04 \x01(\tR\asummary\x12\x1a\n" +
	"\bseverity\x18\x05 \x01(\tR\bseverity\x12\x1c\n" +
	"\tpublished\x18\x06 \x01(\tR\tpublished\x12\x1a\n" +
	"\bmodified\x18\a \x01(\tR\bmodified\x12\x16\n" +
	"\x06source\x18\b \x01(\tR\x06source\x1a\x8c\x01\n" +
	"\x05Purls\x12\x12\n" +
	"\x04purl\x18\x01 \x01(\tR\x04purl\x12o\n" +
	"\x0fvulnerabilities\x18\x02 \x03(\v2E.scanoss.api.vulnerabilities.v2.VulnerabilityResponse.VulnerabilitiesR\x0fvulnerabilities2\xdb\x03\n" +
	"\x0fVulnerabilities\x12x\n" +
	"\x04Echo\x12\".scanoss.api.common.v2.EchoRequest\x1a#.scanoss.api.common.v2.EchoResponse\"'\x82\xd3\xe4\x93\x02!:\x01*\"\x1c/api/v2/vulnerabilities/echo\x12\x95\x01\n" +
	"\aGetCpes\x124.scanoss.api.vulnerabilities.v2.VulnerabilityRequest\x1a+.scanoss.api.vulnerabilities.v2.CpeResponse\"'\x82\xd3\xe4\x93\x02!:\x01*\"\x1c/api/v2/vulnerabilities/cpes\x12\xb5\x01\n" +
	"\x12GetVulnerabilities\x124.scanoss.api.vulnerabilities.v2.VulnerabilityRequest\x1a5.scanoss.api.vulnerabilities.v2.VulnerabilityResponse\"2\x82\xd3\xe4\x93\x02,:\x01*\"'/api/v2/vulnerabilities/vulnerabilitiesB\xac\x02\x92A\xe7\x01\x12\x80\x01\n" +
	"\x1dSCANOSS Vulnerability Service\"Z\n" +
	"\x17scanoss-vulnerabilities\x12*https://github.com/scanoss/vulnerabilities\x1a\x13support@scanoss.com2\x032.0*\x01\x012\x10application/json:\x10application/jsonR;\n" +
	"\x03404\x124\n" +
	"*Returned when the resource does not exist.\x12\x06\n" +
	"\x04\x9a\x02\x01\aZ?github.com/scanoss/papi/api/vulnerabilitiesv2;vulnerabilitiesv2b\x06proto3"

var (
	file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDescOnce sync.Once
	file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDescData []byte
)

func file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDescGZIP() []byte {
	file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDescOnce.Do(func() {
		file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDesc), len(file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDesc)))
	})
	return file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDescData
}

var file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_goTypes = []any{
	(*VulnerabilityRequest)(nil),                  // 0: scanoss.api.vulnerabilities.v2.VulnerabilityRequest
	(*CpeResponse)(nil),                           // 1: scanoss.api.vulnerabilities.v2.CpeResponse
	(*VulnerabilityResponse)(nil),                 // 2: scanoss.api.vulnerabilities.v2.VulnerabilityResponse
	(*VulnerabilityRequest_Purls)(nil),            // 3: scanoss.api.vulnerabilities.v2.VulnerabilityRequest.Purls
	(*CpeResponse_Purls)(nil),                     // 4: scanoss.api.vulnerabilities.v2.CpeResponse.Purls
	(*VulnerabilityResponse_Vulnerabilities)(nil), // 5: scanoss.api.vulnerabilities.v2.VulnerabilityResponse.Vulnerabilities
	(*VulnerabilityResponse_Purls)(nil),           // 6: scanoss.api.vulnerabilities.v2.VulnerabilityResponse.Purls
	(*commonv2.StatusResponse)(nil),               // 7: scanoss.api.common.v2.StatusResponse
	(*commonv2.EchoRequest)(nil),                  // 8: scanoss.api.common.v2.EchoRequest
	(*commonv2.EchoResponse)(nil),                 // 9: scanoss.api.common.v2.EchoResponse
}
var file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_depIdxs = []int32{
	3, // 0: scanoss.api.vulnerabilities.v2.VulnerabilityRequest.purls:type_name -> scanoss.api.vulnerabilities.v2.VulnerabilityRequest.Purls
	4, // 1: scanoss.api.vulnerabilities.v2.CpeResponse.purls:type_name -> scanoss.api.vulnerabilities.v2.CpeResponse.Purls
	7, // 2: scanoss.api.vulnerabilities.v2.CpeResponse.status:type_name -> scanoss.api.common.v2.StatusResponse
	6, // 3: scanoss.api.vulnerabilities.v2.VulnerabilityResponse.purls:type_name -> scanoss.api.vulnerabilities.v2.VulnerabilityResponse.Purls
	7, // 4: scanoss.api.vulnerabilities.v2.VulnerabilityResponse.status:type_name -> scanoss.api.common.v2.StatusResponse
	5, // 5: scanoss.api.vulnerabilities.v2.VulnerabilityResponse.Purls.vulnerabilities:type_name -> scanoss.api.vulnerabilities.v2.VulnerabilityResponse.Vulnerabilities
	8, // 6: scanoss.api.vulnerabilities.v2.Vulnerabilities.Echo:input_type -> scanoss.api.common.v2.EchoRequest
	0, // 7: scanoss.api.vulnerabilities.v2.Vulnerabilities.GetCpes:input_type -> scanoss.api.vulnerabilities.v2.VulnerabilityRequest
	0, // 8: scanoss.api.vulnerabilities.v2.Vulnerabilities.GetVulnerabilities:input_type -> scanoss.api.vulnerabilities.v2.VulnerabilityRequest
	9, // 9: scanoss.api.vulnerabilities.v2.Vulnerabilities.Echo:output_type -> scanoss.api.common.v2.EchoResponse
	1, // 10: scanoss.api.vulnerabilities.v2.Vulnerabilities.GetCpes:output_type -> scanoss.api.vulnerabilities.v2.CpeResponse
	2, // 11: scanoss.api.vulnerabilities.v2.Vulnerabilities.GetVulnerabilities:output_type -> scanoss.api.vulnerabilities.v2.VulnerabilityResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_init() }
func file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_init() {
	if File_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDesc), len(file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_goTypes,
		DependencyIndexes: file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_depIdxs,
		MessageInfos:      file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes,
	}.Build()
	File_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto = out.File
	file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_goTypes = nil
	file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_depIdxs = nil
}
