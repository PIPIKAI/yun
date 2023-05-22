// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.1
// source: yun.proto

package pb

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

type DownloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Md5  string `protobuf:"bytes,1,opt,name=md5,proto3" json:"md5,omitempty"`
	Size int64  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *DownloadRequest) Reset() {
	*x = DownloadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yun_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadRequest) ProtoMessage() {}

func (x *DownloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yun_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadRequest.ProtoReflect.Descriptor instead.
func (*DownloadRequest) Descriptor() ([]byte, []int) {
	return file_yun_proto_rawDescGZIP(), []int{0}
}

func (x *DownloadRequest) GetMd5() string {
	if x != nil {
		return x.Md5
	}
	return ""
}

func (x *DownloadRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type DownloadReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *DownloadReply) Reset() {
	*x = DownloadReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yun_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadReply) ProtoMessage() {}

func (x *DownloadReply) ProtoReflect() protoreflect.Message {
	mi := &file_yun_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadReply.ProtoReflect.Descriptor instead.
func (*DownloadReply) Descriptor() ([]byte, []int) {
	return file_yun_proto_rawDescGZIP(), []int{1}
}

func (x *DownloadReply) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type FileMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size    int64  `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Format  string `protobuf:"bytes,3,opt,name=format,proto3" json:"format,omitempty"`
	ModTime int64  `protobuf:"varint,4,opt,name=ModTime,proto3" json:"ModTime,omitempty"`
	Md5     string `protobuf:"bytes,5,opt,name=md5,proto3" json:"md5,omitempty"`
}

func (x *FileMeta) Reset() {
	*x = FileMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yun_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileMeta) ProtoMessage() {}

func (x *FileMeta) ProtoReflect() protoreflect.Message {
	mi := &file_yun_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileMeta.ProtoReflect.Descriptor instead.
func (*FileMeta) Descriptor() ([]byte, []int) {
	return file_yun_proto_rawDescGZIP(), []int{2}
}

func (x *FileMeta) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *FileMeta) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileMeta) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

func (x *FileMeta) GetModTime() int64 {
	if x != nil {
		return x.ModTime
	}
	return 0
}

func (x *FileMeta) GetMd5() string {
	if x != nil {
		return x.Md5
	}
	return ""
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	Size     int64  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Content  []byte `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yun_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_yun_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_yun_proto_rawDescGZIP(), []int{3}
}

func (x *File) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *File) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *File) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type UploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Md5     string `protobuf:"bytes,1,opt,name=md5,proto3" json:"md5,omitempty"`
	RawData []byte `protobuf:"bytes,3,opt,name=rawData,proto3" json:"rawData,omitempty"`
}

func (x *UploadRequest) Reset() {
	*x = UploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yun_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRequest) ProtoMessage() {}

func (x *UploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yun_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRequest.ProtoReflect.Descriptor instead.
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return file_yun_proto_rawDescGZIP(), []int{4}
}

func (x *UploadRequest) GetMd5() string {
	if x != nil {
		return x.Md5
	}
	return ""
}

func (x *UploadRequest) GetRawData() []byte {
	if x != nil {
		return x.RawData
	}
	return nil
}

type UploadReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Md5  string `protobuf:"bytes,2,opt,name=md5,proto3" json:"md5,omitempty"` // bytes link= 1;
}

func (x *UploadReply) Reset() {
	*x = UploadReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yun_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadReply) ProtoMessage() {}

func (x *UploadReply) ProtoReflect() protoreflect.Message {
	mi := &file_yun_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadReply.ProtoReflect.Descriptor instead.
func (*UploadReply) Descriptor() ([]byte, []int) {
	return file_yun_proto_rawDescGZIP(), []int{5}
}

func (x *UploadReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *UploadReply) GetMd5() string {
	if x != nil {
		return x.Md5
	}
	return ""
}

type ManageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Md5    string `protobuf:"bytes,1,opt,name=md5,proto3" json:"md5,omitempty"`
	Method string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *ManageRequest) Reset() {
	*x = ManageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yun_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManageRequest) ProtoMessage() {}

func (x *ManageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yun_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManageRequest.ProtoReflect.Descriptor instead.
func (*ManageRequest) Descriptor() ([]byte, []int) {
	return file_yun_proto_rawDescGZIP(), []int{6}
}

func (x *ManageRequest) GetMd5() string {
	if x != nil {
		return x.Md5
	}
	return ""
}

func (x *ManageRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

type ManageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ManageReply) Reset() {
	*x = ManageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yun_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManageReply) ProtoMessage() {}

func (x *ManageReply) ProtoReflect() protoreflect.Message {
	mi := &file_yun_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManageReply.ProtoReflect.Descriptor instead.
func (*ManageReply) Descriptor() ([]byte, []int) {
	return file_yun_proto_rawDescGZIP(), []int{7}
}

var File_yun_proto protoreflect.FileDescriptor

var file_yun_proto_rawDesc = []byte{
	0x0a, 0x09, 0x79, 0x75, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22,
	0x37, 0x0a, 0x0f, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x64, 0x35, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x64, 0x35, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x29, 0x0a, 0x0d, 0x44, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x22, 0x76, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x4d, 0x6f, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x4d, 0x6f, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x64, 0x35,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x64, 0x35, 0x22, 0x50, 0x0a, 0x04, 0x46,
	0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x3b, 0x0a,
	0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x64, 0x35, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x64, 0x35,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x72, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x22, 0x33, 0x0a, 0x0b, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x64, 0x35, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x64, 0x35, 0x22,
	0x39, 0x0a, 0x0d, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x64, 0x35, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x64, 0x35, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x0d, 0x0a, 0x0b, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x9f, 0x01, 0x0a, 0x07, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x11, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x06, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x12,
	0x11, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x08, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yun_proto_rawDescOnce sync.Once
	file_yun_proto_rawDescData = file_yun_proto_rawDesc
)

func file_yun_proto_rawDescGZIP() []byte {
	file_yun_proto_rawDescOnce.Do(func() {
		file_yun_proto_rawDescData = protoimpl.X.CompressGZIP(file_yun_proto_rawDescData)
	})
	return file_yun_proto_rawDescData
}

var file_yun_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_yun_proto_goTypes = []interface{}{
	(*DownloadRequest)(nil), // 0: pb.DownloadRequest
	(*DownloadReply)(nil),   // 1: pb.DownloadReply
	(*FileMeta)(nil),        // 2: pb.FileMeta
	(*File)(nil),            // 3: pb.File
	(*UploadRequest)(nil),   // 4: pb.UploadRequest
	(*UploadReply)(nil),     // 5: pb.UploadReply
	(*ManageRequest)(nil),   // 6: pb.ManageRequest
	(*ManageReply)(nil),     // 7: pb.ManageReply
}
var file_yun_proto_depIdxs = []int32{
	4, // 0: pb.Storage.Upload:input_type -> pb.UploadRequest
	6, // 1: pb.Storage.Manage:input_type -> pb.ManageRequest
	0, // 2: pb.Storage.Download:input_type -> pb.DownloadRequest
	5, // 3: pb.Storage.Upload:output_type -> pb.UploadReply
	7, // 4: pb.Storage.Manage:output_type -> pb.ManageReply
	1, // 5: pb.Storage.Download:output_type -> pb.DownloadReply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yun_proto_init() }
func file_yun_proto_init() {
	if File_yun_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yun_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadRequest); i {
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
		file_yun_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadReply); i {
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
		file_yun_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileMeta); i {
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
		file_yun_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_yun_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadRequest); i {
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
		file_yun_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadReply); i {
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
		file_yun_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManageRequest); i {
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
		file_yun_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManageReply); i {
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
			RawDescriptor: file_yun_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yun_proto_goTypes,
		DependencyIndexes: file_yun_proto_depIdxs,
		MessageInfos:      file_yun_proto_msgTypes,
	}.Build()
	File_yun_proto = out.File
	file_yun_proto_rawDesc = nil
	file_yun_proto_goTypes = nil
	file_yun_proto_depIdxs = nil
}
