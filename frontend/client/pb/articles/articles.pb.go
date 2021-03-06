// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: pb/articles.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_articles_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_pb_articles_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_pb_articles_proto_rawDescGZIP(), []int{0}
}

func (x *ID) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type IDs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IDs []*ID `protobuf:"bytes,1,rep,name=IDs,proto3" json:"IDs,omitempty"`
}

func (x *IDs) Reset() {
	*x = IDs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_articles_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDs) ProtoMessage() {}

func (x *IDs) ProtoReflect() protoreflect.Message {
	mi := &file_pb_articles_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDs.ProtoReflect.Descriptor instead.
func (*IDs) Descriptor() ([]byte, []int) {
	return file_pb_articles_proto_rawDescGZIP(), []int{1}
}

func (x *IDs) GetIDs() []*ID {
	if x != nil {
		return x.IDs
	}
	return nil
}

type Count struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
}

func (x *Count) Reset() {
	*x = Count{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_articles_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Count) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Count) ProtoMessage() {}

func (x *Count) ProtoReflect() protoreflect.Message {
	mi := &file_pb_articles_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Count.ProtoReflect.Descriptor instead.
func (*Count) Descriptor() ([]byte, []int) {
	return file_pb_articles_proto_rawDescGZIP(), []int{2}
}

func (x *Count) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category string `protobuf:"bytes,1,opt,name=Category,proto3" json:"Category,omitempty"`
}

func (x *Category) Reset() {
	*x = Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_articles_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_pb_articles_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_pb_articles_proto_rawDescGZIP(), []int{3}
}

func (x *Category) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

type Article struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Timestamp   int64  `protobuf:"varint,2,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	AID         string `protobuf:"bytes,3,opt,name=AID,proto3" json:"AID,omitempty"`
	Title       string `protobuf:"bytes,4,opt,name=Title,proto3" json:"Title,omitempty"`
	Category    string `protobuf:"bytes,5,opt,name=Category,proto3" json:"Category,omitempty"`
	Abstract    string `protobuf:"bytes,6,opt,name=Abstract,proto3" json:"Abstract,omitempty"`
	ArticleTags string `protobuf:"bytes,7,opt,name=ArticleTags,proto3" json:"ArticleTags,omitempty"`
	Authors     string `protobuf:"bytes,8,opt,name=Authors,proto3" json:"Authors,omitempty"`
	Language    string `protobuf:"bytes,9,opt,name=Language,proto3" json:"Language,omitempty"`
	Text        string `protobuf:"bytes,10,opt,name=Text,proto3" json:"Text,omitempty"`
	Image       string `protobuf:"bytes,11,opt,name=Image,proto3" json:"Image,omitempty"`
	Video       string `protobuf:"bytes,12,opt,name=Video,proto3" json:"Video,omitempty"`
}

func (x *Article) Reset() {
	*x = Article{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_articles_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Article) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Article) ProtoMessage() {}

func (x *Article) ProtoReflect() protoreflect.Message {
	mi := &file_pb_articles_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Article.ProtoReflect.Descriptor instead.
func (*Article) Descriptor() ([]byte, []int) {
	return file_pb_articles_proto_rawDescGZIP(), []int{4}
}

func (x *Article) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Article) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Article) GetAID() string {
	if x != nil {
		return x.AID
	}
	return ""
}

func (x *Article) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Article) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Article) GetAbstract() string {
	if x != nil {
		return x.Abstract
	}
	return ""
}

func (x *Article) GetArticleTags() string {
	if x != nil {
		return x.ArticleTags
	}
	return ""
}

func (x *Article) GetAuthors() string {
	if x != nil {
		return x.Authors
	}
	return ""
}

func (x *Article) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *Article) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Article) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Article) GetVideo() string {
	if x != nil {
		return x.Video
	}
	return ""
}

type Articles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Articles []*Article `protobuf:"bytes,1,rep,name=Articles,proto3" json:"Articles,omitempty"`
}

func (x *Articles) Reset() {
	*x = Articles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_articles_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Articles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Articles) ProtoMessage() {}

func (x *Articles) ProtoReflect() protoreflect.Message {
	mi := &file_pb_articles_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Articles.ProtoReflect.Descriptor instead.
func (*Articles) Descriptor() ([]byte, []int) {
	return file_pb_articles_proto_rawDescGZIP(), []int{5}
}

func (x *Articles) GetArticles() []*Article {
	if x != nil {
		return x.Articles
	}
	return nil
}

type InformationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InformationRequest) Reset() {
	*x = InformationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_articles_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InformationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InformationRequest) ProtoMessage() {}

func (x *InformationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_articles_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InformationRequest.ProtoReflect.Descriptor instead.
func (*InformationRequest) Descriptor() ([]byte, []int) {
	return file_pb_articles_proto_rawDescGZIP(), []int{6}
}

type InformationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IP   string `protobuf:"bytes,1,opt,name=IP,proto3" json:"IP,omitempty"`
	Host string `protobuf:"bytes,2,opt,name=Host,proto3" json:"Host,omitempty"`
}

func (x *InformationResponse) Reset() {
	*x = InformationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_articles_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InformationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InformationResponse) ProtoMessage() {}

func (x *InformationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_articles_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InformationResponse.ProtoReflect.Descriptor instead.
func (*InformationResponse) Descriptor() ([]byte, []int) {
	return file_pb_articles_proto_rawDescGZIP(), []int{7}
}

func (x *InformationResponse) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

func (x *InformationResponse) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

var File_pb_articles_proto protoreflect.FileDescriptor

var file_pb_articles_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x62, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x44, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x22, 0x1f, 0x0a,
	0x03, 0x49, 0x44, 0x73, 0x12, 0x18, 0x0a, 0x03, 0x49, 0x44, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x44, 0x52, 0x03, 0x49, 0x44, 0x73, 0x22, 0x1d,
	0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x26, 0x0a,
	0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0xaf, 0x02, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x10, 0x0a, 0x03, 0x41, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x41, 0x49,
	0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x41, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x61, 0x67, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x61, 0x67,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x4c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x22, 0x33, 0x0a, 0x08, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x08, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x52, 0x08, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x22, 0x14, 0x0a, 0x12,
	0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x39, 0x0a, 0x13, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x50, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x50, 0x12, 0x12, 0x0a, 0x04, 0x48, 0x6f, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x32, 0xb5, 0x02,
	0x0a, 0x0e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x47, 0x0a, 0x12, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x06, 0x2e,
	0x70, 0x62, 0x2e, 0x49, 0x44, 0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x12, 0x0c, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x42, 0x79, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x12, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x44, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x22, 0x00, 0x12, 0x23, 0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x1a, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x44, 0x22, 0x00, 0x12, 0x26, 0x0a,
	0x0b, 0x4e, 0x65, 0x77, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x12, 0x0c, 0x2e, 0x70,
	0x62, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x1a, 0x07, 0x2e, 0x70, 0x62, 0x2e,
	0x49, 0x44, 0x73, 0x22, 0x00, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x72, 0x6f, 0x75, 0x69, 0x6f, 0x75, 0x69, 0x2f, 0x74, 0x61, 0x67,
	0x65, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x73, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_articles_proto_rawDescOnce sync.Once
	file_pb_articles_proto_rawDescData = file_pb_articles_proto_rawDesc
)

func file_pb_articles_proto_rawDescGZIP() []byte {
	file_pb_articles_proto_rawDescOnce.Do(func() {
		file_pb_articles_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_articles_proto_rawDescData)
	})
	return file_pb_articles_proto_rawDescData
}

var file_pb_articles_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pb_articles_proto_goTypes = []interface{}{
	(*ID)(nil),                  // 0: pb.ID
	(*IDs)(nil),                 // 1: pb.IDs
	(*Count)(nil),               // 2: pb.Count
	(*Category)(nil),            // 3: pb.Category
	(*Article)(nil),             // 4: pb.Article
	(*Articles)(nil),            // 5: pb.Articles
	(*InformationRequest)(nil),  // 6: pb.InformationRequest
	(*InformationResponse)(nil), // 7: pb.InformationResponse
}
var file_pb_articles_proto_depIdxs = []int32{
	0, // 0: pb.IDs.IDs:type_name -> pb.ID
	4, // 1: pb.Articles.Articles:type_name -> pb.Article
	6, // 2: pb.ArticleService.ServiceInformation:input_type -> pb.InformationRequest
	0, // 3: pb.ArticleService.GetSingleArticle:input_type -> pb.ID
	3, // 4: pb.ArticleService.GetCategoryArticles:input_type -> pb.Category
	0, // 5: pb.ArticleService.GetArticlesByRegion:input_type -> pb.ID
	4, // 6: pb.ArticleService.NewArticle:input_type -> pb.Article
	5, // 7: pb.ArticleService.NewArticles:input_type -> pb.Articles
	7, // 8: pb.ArticleService.ServiceInformation:output_type -> pb.InformationResponse
	4, // 9: pb.ArticleService.GetSingleArticle:output_type -> pb.Article
	5, // 10: pb.ArticleService.GetCategoryArticles:output_type -> pb.Articles
	5, // 11: pb.ArticleService.GetArticlesByRegion:output_type -> pb.Articles
	0, // 12: pb.ArticleService.NewArticle:output_type -> pb.ID
	1, // 13: pb.ArticleService.NewArticles:output_type -> pb.IDs
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pb_articles_proto_init() }
func file_pb_articles_proto_init() {
	if File_pb_articles_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_articles_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
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
		file_pb_articles_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDs); i {
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
		file_pb_articles_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Count); i {
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
		file_pb_articles_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Category); i {
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
		file_pb_articles_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Article); i {
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
		file_pb_articles_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Articles); i {
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
		file_pb_articles_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InformationRequest); i {
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
		file_pb_articles_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InformationResponse); i {
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
			RawDescriptor: file_pb_articles_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_articles_proto_goTypes,
		DependencyIndexes: file_pb_articles_proto_depIdxs,
		MessageInfos:      file_pb_articles_proto_msgTypes,
	}.Build()
	File_pb_articles_proto = out.File
	file_pb_articles_proto_rawDesc = nil
	file_pb_articles_proto_goTypes = nil
	file_pb_articles_proto_depIdxs = nil
}
