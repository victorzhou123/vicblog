// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.17.3
// source: tag.proto

package rpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TagName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TagName) Reset() {
	*x = TagName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagName) ProtoMessage() {}

func (x *TagName) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagName.ProtoReflect.Descriptor instead.
func (*TagName) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{0}
}

func (x *TagName) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type TagNames struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Names []*TagName `protobuf:"bytes,2,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *TagNames) Reset() {
	*x = TagNames{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagNames) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagNames) ProtoMessage() {}

func (x *TagNames) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagNames.ProtoReflect.Descriptor instead.
func (*TagNames) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{1}
}

func (x *TagNames) GetNames() []*TagName {
	if x != nil {
		return x.Names
	}
	return nil
}

type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt int64  `protobuf:"varint,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{2}
}

func (x *Tag) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Tag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tag) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type Tags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags []*Tag `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *Tags) Reset() {
	*x = Tags{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tags) ProtoMessage() {}

func (x *Tags) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tags.ProtoReflect.Descriptor instead.
func (*Tags) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{3}
}

func (x *Tags) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

type TagWithRelatedArticleAmount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag                  *Tag  `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	RelatedArticleAmount int64 `protobuf:"varint,2,opt,name=relatedArticleAmount,proto3" json:"relatedArticleAmount,omitempty"`
}

func (x *TagWithRelatedArticleAmount) Reset() {
	*x = TagWithRelatedArticleAmount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagWithRelatedArticleAmount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagWithRelatedArticleAmount) ProtoMessage() {}

func (x *TagWithRelatedArticleAmount) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagWithRelatedArticleAmount.ProtoReflect.Descriptor instead.
func (*TagWithRelatedArticleAmount) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{4}
}

func (x *TagWithRelatedArticleAmount) GetTag() *Tag {
	if x != nil {
		return x.Tag
	}
	return nil
}

func (x *TagWithRelatedArticleAmount) GetRelatedArticleAmount() int64 {
	if x != nil {
		return x.RelatedArticleAmount
	}
	return 0
}

type TagsWithRelatedArticleAmount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags []*TagWithRelatedArticleAmount `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *TagsWithRelatedArticleAmount) Reset() {
	*x = TagsWithRelatedArticleAmount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagsWithRelatedArticleAmount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagsWithRelatedArticleAmount) ProtoMessage() {}

func (x *TagsWithRelatedArticleAmount) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagsWithRelatedArticleAmount.ProtoReflect.Descriptor instead.
func (*TagsWithRelatedArticleAmount) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{5}
}

func (x *TagsWithRelatedArticleAmount) GetTags() []*TagWithRelatedArticleAmount {
	if x != nil {
		return x.Tags
	}
	return nil
}

type TagList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaginationStatus *PaginationStatus              `protobuf:"bytes,1,opt,name=paginationStatus,proto3" json:"paginationStatus,omitempty"`
	Tags             []*TagWithRelatedArticleAmount `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *TagList) Reset() {
	*x = TagList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagList) ProtoMessage() {}

func (x *TagList) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagList.ProtoReflect.Descriptor instead.
func (*TagList) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{6}
}

func (x *TagList) GetPaginationStatus() *PaginationStatus {
	if x != nil {
		return x.PaginationStatus
	}
	return nil
}

func (x *TagList) GetTags() []*TagWithRelatedArticleAmount {
	if x != nil {
		return x.Tags
	}
	return nil
}

type ArticleIdAndTagIds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArticleId *Id   `protobuf:"bytes,1,opt,name=articleId,proto3" json:"articleId,omitempty"`
	TagIds    []*Id `protobuf:"bytes,2,rep,name=tagIds,proto3" json:"tagIds,omitempty"`
}

func (x *ArticleIdAndTagIds) Reset() {
	*x = ArticleIdAndTagIds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleIdAndTagIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleIdAndTagIds) ProtoMessage() {}

func (x *ArticleIdAndTagIds) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleIdAndTagIds.ProtoReflect.Descriptor instead.
func (*ArticleIdAndTagIds) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{7}
}

func (x *ArticleIdAndTagIds) GetArticleId() *Id {
	if x != nil {
		return x.ArticleId
	}
	return nil
}

func (x *ArticleIdAndTagIds) GetTagIds() []*Id {
	if x != nil {
		return x.TagIds
	}
	return nil
}

var File_tag_proto protoreflect.FileDescriptor

var file_tag_proto_rawDesc = []byte{
	0x0a, 0x09, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x74, 0x61, 0x67,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x07, 0x54,
	0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2e, 0x0a, 0x08, 0x54, 0x61,
	0x67, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x54, 0x61, 0x67, 0x4e,
	0x61, 0x6d, 0x65, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x47, 0x0a, 0x03, 0x54, 0x61,
	0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x24, 0x0a, 0x04, 0x54, 0x61, 0x67, 0x73, 0x12, 0x1c, 0x0a, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x74, 0x61, 0x67, 0x2e,
	0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x6d, 0x0a, 0x1b, 0x54, 0x61, 0x67,
	0x57, 0x69, 0x74, 0x68, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x54, 0x61, 0x67, 0x52,
	0x03, 0x74, 0x61, 0x67, 0x12, 0x32, 0x0a, 0x14, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x14, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x54, 0x0a, 0x1c, 0x54, 0x61, 0x67, 0x73,
	0x57, 0x69, 0x74, 0x68, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x54, 0x61, 0x67,
	0x57, 0x69, 0x74, 0x68, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x85,
	0x01, 0x0a, 0x07, 0x54, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x10, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x10,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x34, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x74, 0x61, 0x67, 0x2e, 0x54, 0x61, 0x67, 0x57, 0x69, 0x74, 0x68, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x62, 0x0a, 0x12, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x49, 0x64, 0x41, 0x6e, 0x64, 0x54, 0x61, 0x67, 0x49, 0x64, 0x73, 0x12, 0x28, 0x0a, 0x09,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x64, 0x52, 0x09, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x06, 0x74, 0x61, 0x67, 0x49, 0x64, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x49, 0x64, 0x52, 0x06, 0x74, 0x61, 0x67, 0x49, 0x64, 0x73, 0x32, 0xde, 0x04, 0x0a, 0x0a, 0x54,
	0x61, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x41, 0x64, 0x64,
	0x54, 0x61, 0x67, 0x73, 0x12, 0x0d, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x54, 0x61, 0x67, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x39, 0x0a,
	0x13, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x42, 0x79, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0c, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x54,
	0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x61, 0x67, 0x73, 0x12, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x21, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x54, 0x61, 0x67, 0x73, 0x57,
	0x69, 0x74, 0x68, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x61, 0x67, 0x12, 0x0a, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x64, 0x1a, 0x09, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x54, 0x61, 0x67,
	0x73, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x54, 0x61, 0x67, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x0a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12,
	0x0a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x64, 0x1a, 0x0b, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x64, 0x73, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x20, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49,
	0x64, 0x73, 0x54, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x54, 0x61, 0x67, 0x49, 0x64, 0x12, 0x0a,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x64, 0x1a, 0x0b, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x64, 0x73, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x18, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x17, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x49, 0x64, 0x41, 0x6e, 0x64, 0x54, 0x61, 0x67, 0x49, 0x64, 0x73, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x19, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x0a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49,
	0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x38, 0x5a, 0x36, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x69, 0x63, 0x74, 0x6f, 0x72,
	0x7a, 0x68, 0x6f, 0x75, 0x31, 0x32, 0x33, 0x2f, 0x76, 0x69, 0x63, 0x62, 0x6c, 0x6f, 0x67, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tag_proto_rawDescOnce sync.Once
	file_tag_proto_rawDescData = file_tag_proto_rawDesc
)

func file_tag_proto_rawDescGZIP() []byte {
	file_tag_proto_rawDescOnce.Do(func() {
		file_tag_proto_rawDescData = protoimpl.X.CompressGZIP(file_tag_proto_rawDescData)
	})
	return file_tag_proto_rawDescData
}

var file_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_tag_proto_goTypes = []any{
	(*TagName)(nil),                      // 0: tag.TagName
	(*TagNames)(nil),                     // 1: tag.TagNames
	(*Tag)(nil),                          // 2: tag.Tag
	(*Tags)(nil),                         // 3: tag.Tags
	(*TagWithRelatedArticleAmount)(nil),  // 4: tag.TagWithRelatedArticleAmount
	(*TagsWithRelatedArticleAmount)(nil), // 5: tag.TagsWithRelatedArticleAmount
	(*TagList)(nil),                      // 6: tag.TagList
	(*ArticleIdAndTagIds)(nil),           // 7: tag.ArticleIdAndTagIds
	(*PaginationStatus)(nil),             // 8: common.PaginationStatus
	(*Id)(nil),                           // 9: common.Id
	(*Pagination)(nil),                   // 10: common.Pagination
	(*Amount)(nil),                       // 11: common.Amount
	(*emptypb.Empty)(nil),                // 12: google.protobuf.Empty
	(*Ids)(nil),                          // 13: common.Ids
}
var file_tag_proto_depIdxs = []int32{
	0,  // 0: tag.TagNames.names:type_name -> tag.TagName
	2,  // 1: tag.Tags.tags:type_name -> tag.Tag
	2,  // 2: tag.TagWithRelatedArticleAmount.tag:type_name -> tag.Tag
	4,  // 3: tag.TagsWithRelatedArticleAmount.tags:type_name -> tag.TagWithRelatedArticleAmount
	8,  // 4: tag.TagList.paginationStatus:type_name -> common.PaginationStatus
	4,  // 5: tag.TagList.tags:type_name -> tag.TagWithRelatedArticleAmount
	9,  // 6: tag.ArticleIdAndTagIds.articleId:type_name -> common.Id
	9,  // 7: tag.ArticleIdAndTagIds.tagIds:type_name -> common.Id
	1,  // 8: tag.TagService.AddTags:input_type -> tag.TagNames
	10, // 9: tag.TagService.ListTagByPagination:input_type -> common.Pagination
	11, // 10: tag.TagService.ListTags:input_type -> common.Amount
	9,  // 11: tag.TagService.GetArticleTag:input_type -> common.Id
	12, // 12: tag.TagService.GetTotalNumberOfTags:input_type -> google.protobuf.Empty
	9,  // 13: tag.TagService.Delete:input_type -> common.Id
	9,  // 14: tag.TagService.GetRelationWithArticle:input_type -> common.Id
	9,  // 15: tag.TagService.GetRelatedArticleIdsThroughTagId:input_type -> common.Id
	7,  // 16: tag.TagService.BuildRelationWithArticle:input_type -> tag.ArticleIdAndTagIds
	9,  // 17: tag.TagService.RemoveRelationWithArticle:input_type -> common.Id
	12, // 18: tag.TagService.AddTags:output_type -> google.protobuf.Empty
	6,  // 19: tag.TagService.ListTagByPagination:output_type -> tag.TagList
	5,  // 20: tag.TagService.ListTags:output_type -> tag.TagsWithRelatedArticleAmount
	3,  // 21: tag.TagService.GetArticleTag:output_type -> tag.Tags
	11, // 22: tag.TagService.GetTotalNumberOfTags:output_type -> common.Amount
	12, // 23: tag.TagService.Delete:output_type -> google.protobuf.Empty
	13, // 24: tag.TagService.GetRelationWithArticle:output_type -> common.Ids
	13, // 25: tag.TagService.GetRelatedArticleIdsThroughTagId:output_type -> common.Ids
	12, // 26: tag.TagService.BuildRelationWithArticle:output_type -> google.protobuf.Empty
	12, // 27: tag.TagService.RemoveRelationWithArticle:output_type -> google.protobuf.Empty
	18, // [18:28] is the sub-list for method output_type
	8,  // [8:18] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_tag_proto_init() }
func file_tag_proto_init() {
	if File_tag_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tag_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TagName); i {
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
		file_tag_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TagNames); i {
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
		file_tag_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Tag); i {
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
		file_tag_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Tags); i {
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
		file_tag_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*TagWithRelatedArticleAmount); i {
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
		file_tag_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*TagsWithRelatedArticleAmount); i {
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
		file_tag_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*TagList); i {
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
		file_tag_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*ArticleIdAndTagIds); i {
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
			RawDescriptor: file_tag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tag_proto_goTypes,
		DependencyIndexes: file_tag_proto_depIdxs,
		MessageInfos:      file_tag_proto_msgTypes,
	}.Build()
	File_tag_proto = out.File
	file_tag_proto_rawDesc = nil
	file_tag_proto_goTypes = nil
	file_tag_proto_depIdxs = nil
}
