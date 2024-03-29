// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: protos/pages.proto

package protos

import (
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

// Export io description
type ContentType int32

const (
	ContentType_JSON     ContentType = 0
	ContentType_HTML     ContentType = 1
	ContentType_WIKITEXT ContentType = 2
)

// Enum value maps for ContentType.
var (
	ContentType_name = map[int32]string{
		0: "JSON",
		1: "HTML",
		2: "WIKITEXT",
	}
	ContentType_value = map[string]int32{
		"JSON":     0,
		"HTML":     1,
		"WIKITEXT": 2,
	}
)

func (x ContentType) Enum() *ContentType {
	p := new(ContentType)
	*p = x
	return p
}

func (x ContentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContentType) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_pages_proto_enumTypes[0].Descriptor()
}

func (ContentType) Type() protoreflect.EnumType {
	return &file_protos_pages_proto_enumTypes[0]
}

func (x ContentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContentType.Descriptor instead.
func (ContentType) EnumDescriptor() ([]byte, []int) {
	return file_protos_pages_proto_rawDescGZIP(), []int{0}
}

// Index io description
type IndexRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IndexRequest) Reset() {
	*x = IndexRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_pages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexRequest) ProtoMessage() {}

func (x *IndexRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_pages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexRequest.ProtoReflect.Descriptor instead.
func (*IndexRequest) Descriptor() ([]byte, []int) {
	return file_protos_pages_proto_rawDescGZIP(), []int{0}
}

type IndexResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total  int32 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Errors int32 `protobuf:"varint,2,opt,name=errors,proto3" json:"errors,omitempty"`
}

func (x *IndexResponse) Reset() {
	*x = IndexResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_pages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexResponse) ProtoMessage() {}

func (x *IndexResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_pages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexResponse.ProtoReflect.Descriptor instead.
func (*IndexResponse) Descriptor() ([]byte, []int) {
	return file_protos_pages_proto_rawDescGZIP(), []int{1}
}

func (x *IndexResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *IndexResponse) GetErrors() int32 {
	if x != nil {
		return x.Errors
	}
	return 0
}

// Scan io description
type FetchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workers int32  `protobuf:"varint,1,opt,name=workers,proto3" json:"workers,omitempty"`
	Batch   int32  `protobuf:"varint,2,opt,name=batch,proto3" json:"batch,omitempty"`
	DbName  string `protobuf:"bytes,3,opt,name=db_name,json=dbName,proto3" json:"db_name,omitempty"`
	Ns      int32  `protobuf:"varint,4,opt,name=ns,proto3" json:"ns,omitempty"`
	Failed  bool   `protobuf:"varint,5,opt,name=failed,proto3" json:"failed,omitempty"`
}

func (x *FetchRequest) Reset() {
	*x = FetchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_pages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchRequest) ProtoMessage() {}

func (x *FetchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_pages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchRequest.ProtoReflect.Descriptor instead.
func (*FetchRequest) Descriptor() ([]byte, []int) {
	return file_protos_pages_proto_rawDescGZIP(), []int{2}
}

func (x *FetchRequest) GetWorkers() int32 {
	if x != nil {
		return x.Workers
	}
	return 0
}

func (x *FetchRequest) GetBatch() int32 {
	if x != nil {
		return x.Batch
	}
	return 0
}

func (x *FetchRequest) GetDbName() string {
	if x != nil {
		return x.DbName
	}
	return ""
}

func (x *FetchRequest) GetNs() int32 {
	if x != nil {
		return x.Ns
	}
	return 0
}

func (x *FetchRequest) GetFailed() bool {
	if x != nil {
		return x.Failed
	}
	return false
}

type FetchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total     int32 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Errors    int32 `protobuf:"varint,2,opt,name=errors,proto3" json:"errors,omitempty"`
	Redirects int32 `protobuf:"varint,3,opt,name=redirects,proto3" json:"redirects,omitempty"`
}

func (x *FetchResponse) Reset() {
	*x = FetchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_pages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchResponse) ProtoMessage() {}

func (x *FetchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_pages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchResponse.ProtoReflect.Descriptor instead.
func (*FetchResponse) Descriptor() ([]byte, []int) {
	return file_protos_pages_proto_rawDescGZIP(), []int{3}
}

func (x *FetchResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *FetchResponse) GetErrors() int32 {
	if x != nil {
		return x.Errors
	}
	return 0
}

func (x *FetchResponse) GetRedirects() int32 {
	if x != nil {
		return x.Redirects
	}
	return 0
}

type ExportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DbName string `protobuf:"bytes,1,opt,name=db_name,json=dbName,proto3" json:"db_name,omitempty"`
	// Deprecated: Do not use.
	ContentType ContentType `protobuf:"varint,2,opt,name=content_type,json=contentType,proto3,enum=pages.ContentType" json:"content_type,omitempty"`
	Workers     int32       `protobuf:"varint,3,opt,name=workers,proto3" json:"workers,omitempty"`
	Ns          int32       `protobuf:"varint,4,opt,name=ns,proto3" json:"ns,omitempty"`
}

func (x *ExportRequest) Reset() {
	*x = ExportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_pages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportRequest) ProtoMessage() {}

func (x *ExportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_pages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportRequest.ProtoReflect.Descriptor instead.
func (*ExportRequest) Descriptor() ([]byte, []int) {
	return file_protos_pages_proto_rawDescGZIP(), []int{4}
}

func (x *ExportRequest) GetDbName() string {
	if x != nil {
		return x.DbName
	}
	return ""
}

// Deprecated: Do not use.
func (x *ExportRequest) GetContentType() ContentType {
	if x != nil {
		return x.ContentType
	}
	return ContentType_JSON
}

func (x *ExportRequest) GetWorkers() int32 {
	if x != nil {
		return x.Workers
	}
	return 0
}

func (x *ExportRequest) GetNs() int32 {
	if x != nil {
		return x.Ns
	}
	return 0
}

type ExportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total  int32 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Errors int32 `protobuf:"varint,2,opt,name=errors,proto3" json:"errors,omitempty"`
}

func (x *ExportResponse) Reset() {
	*x = ExportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_pages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportResponse) ProtoMessage() {}

func (x *ExportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_pages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportResponse.ProtoReflect.Descriptor instead.
func (*ExportResponse) Descriptor() ([]byte, []int) {
	return file_protos_pages_proto_rawDescGZIP(), []int{5}
}

func (x *ExportResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ExportResponse) GetErrors() int32 {
	if x != nil {
		return x.Errors
	}
	return 0
}

// Copy io description
type CopyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workers int32    `protobuf:"varint,1,opt,name=workers,proto3" json:"workers,omitempty"`
	DbNames []string `protobuf:"bytes,2,rep,name=db_names,json=dbNames,proto3" json:"db_names,omitempty"`
	Ns      int32    `protobuf:"varint,3,opt,name=ns,proto3" json:"ns,omitempty"`
}

func (x *CopyRequest) Reset() {
	*x = CopyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_pages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CopyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CopyRequest) ProtoMessage() {}

func (x *CopyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_pages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CopyRequest.ProtoReflect.Descriptor instead.
func (*CopyRequest) Descriptor() ([]byte, []int) {
	return file_protos_pages_proto_rawDescGZIP(), []int{6}
}

func (x *CopyRequest) GetWorkers() int32 {
	if x != nil {
		return x.Workers
	}
	return 0
}

func (x *CopyRequest) GetDbNames() []string {
	if x != nil {
		return x.DbNames
	}
	return nil
}

func (x *CopyRequest) GetNs() int32 {
	if x != nil {
		return x.Ns
	}
	return 0
}

type CopyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total  int32 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Errors int32 `protobuf:"varint,2,opt,name=errors,proto3" json:"errors,omitempty"`
}

func (x *CopyResponse) Reset() {
	*x = CopyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_pages_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CopyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CopyResponse) ProtoMessage() {}

func (x *CopyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_pages_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CopyResponse.ProtoReflect.Descriptor instead.
func (*CopyResponse) Descriptor() ([]byte, []int) {
	return file_protos_pages_proto_rawDescGZIP(), []int{7}
}

func (x *CopyResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *CopyResponse) GetErrors() int32 {
	if x != nil {
		return x.Errors
	}
	return 0
}

var File_protos_pages_proto protoreflect.FileDescriptor

var file_protos_pages_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x22, 0x0e, 0x0a, 0x0c, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3d, 0x0a, 0x0d, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x7f, 0x0a, 0x0c, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x77, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x61, 0x74, 0x63, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x62, 0x61, 0x74, 0x63, 0x68, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x62,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x62, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x22, 0x5b, 0x0a, 0x0d, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72,
	0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x73, 0x22, 0x8d, 0x01, 0x0a, 0x0d, 0x45, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x62,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x62, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x02, 0x18,
	0x01, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x6e, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x6e, 0x73, 0x22, 0x3e, 0x0a, 0x0e, 0x45, 0x78, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x52, 0x0a, 0x0b, 0x43, 0x6f, 0x70, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x73, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x64, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02,
	0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x6e, 0x73, 0x22, 0x3c, 0x0a, 0x0c,
	0x43, 0x6f, 0x70, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2a, 0x2f, 0x0a, 0x0b, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x53, 0x4f,
	0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x54, 0x4d, 0x4c, 0x10, 0x01, 0x12, 0x0c, 0x0a,
	0x08, 0x57, 0x49, 0x4b, 0x49, 0x54, 0x45, 0x58, 0x54, 0x10, 0x02, 0x32, 0xd7, 0x01, 0x0a, 0x05,
	0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x13,
	0x2e, 0x70, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x12, 0x13, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a,
	0x06, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x70, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x43, 0x6f, 0x70, 0x79, 0x12, 0x12, 0x2e, 0x70,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x70, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x70, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1b, 0x5a, 0x19, 0x6f, 0x6b, 0x61, 0x70, 0x69, 0x2d, 0x64,
	0x61, 0x74, 0x61, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_pages_proto_rawDescOnce sync.Once
	file_protos_pages_proto_rawDescData = file_protos_pages_proto_rawDesc
)

func file_protos_pages_proto_rawDescGZIP() []byte {
	file_protos_pages_proto_rawDescOnce.Do(func() {
		file_protos_pages_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_pages_proto_rawDescData)
	})
	return file_protos_pages_proto_rawDescData
}

var file_protos_pages_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_pages_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protos_pages_proto_goTypes = []interface{}{
	(ContentType)(0),       // 0: pages.ContentType
	(*IndexRequest)(nil),   // 1: pages.IndexRequest
	(*IndexResponse)(nil),  // 2: pages.IndexResponse
	(*FetchRequest)(nil),   // 3: pages.FetchRequest
	(*FetchResponse)(nil),  // 4: pages.FetchResponse
	(*ExportRequest)(nil),  // 5: pages.ExportRequest
	(*ExportResponse)(nil), // 6: pages.ExportResponse
	(*CopyRequest)(nil),    // 7: pages.CopyRequest
	(*CopyResponse)(nil),   // 8: pages.CopyResponse
}
var file_protos_pages_proto_depIdxs = []int32{
	0, // 0: pages.ExportRequest.content_type:type_name -> pages.ContentType
	1, // 1: pages.Pages.Index:input_type -> pages.IndexRequest
	3, // 2: pages.Pages.Fetch:input_type -> pages.FetchRequest
	5, // 3: pages.Pages.Export:input_type -> pages.ExportRequest
	7, // 4: pages.Pages.Copy:input_type -> pages.CopyRequest
	2, // 5: pages.Pages.Index:output_type -> pages.IndexResponse
	4, // 6: pages.Pages.Fetch:output_type -> pages.FetchResponse
	6, // 7: pages.Pages.Export:output_type -> pages.ExportResponse
	8, // 8: pages.Pages.Copy:output_type -> pages.CopyResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_pages_proto_init() }
func file_protos_pages_proto_init() {
	if File_protos_pages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_pages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexRequest); i {
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
		file_protos_pages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexResponse); i {
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
		file_protos_pages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchRequest); i {
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
		file_protos_pages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchResponse); i {
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
		file_protos_pages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportRequest); i {
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
		file_protos_pages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportResponse); i {
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
		file_protos_pages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CopyRequest); i {
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
		file_protos_pages_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CopyResponse); i {
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
			RawDescriptor: file_protos_pages_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_pages_proto_goTypes,
		DependencyIndexes: file_protos_pages_proto_depIdxs,
		EnumInfos:         file_protos_pages_proto_enumTypes,
		MessageInfos:      file_protos_pages_proto_msgTypes,
	}.Build()
	File_protos_pages_proto = out.File
	file_protos_pages_proto_rawDesc = nil
	file_protos_pages_proto_goTypes = nil
	file_protos_pages_proto_depIdxs = nil
}
