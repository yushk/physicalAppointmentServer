// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: file.proto

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

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" bson:"name"`
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path" bson:"path"`
	File []byte `protobuf:"bytes,3,opt,name=file,proto3" json:"file" bson:"file"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[0]
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
	return file_file_proto_rawDescGZIP(), []int{0}
}

func (x *File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *File) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *File) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

type UploadFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name 文件名
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" bson:"name"`
	// Path 文件路径
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path" bson:"path"`
	// File 二进制文件流
	File []byte `protobuf:"bytes,3,opt,name=file,proto3" json:"file" bson:"file"`
}

func (x *UploadFileRequest) Reset() {
	*x = UploadFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileRequest) ProtoMessage() {}

func (x *UploadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileRequest.ProtoReflect.Descriptor instead.
func (*UploadFileRequest) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{1}
}

func (x *UploadFileRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UploadFileRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *UploadFileRequest) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

type UploadFileReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
}

func (x *UploadFileReply) Reset() {
	*x = UploadFileReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileReply) ProtoMessage() {}

func (x *UploadFileReply) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileReply.ProtoReflect.Descriptor instead.
func (*UploadFileReply) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{2}
}

func (x *UploadFileReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DownloadFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
}

func (x *DownloadFileRequest) Reset() {
	*x = DownloadFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadFileRequest) ProtoMessage() {}

func (x *DownloadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadFileRequest.ProtoReflect.Descriptor instead.
func (*DownloadFileRequest) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{3}
}

func (x *DownloadFileRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DownloadFileReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File []byte `protobuf:"bytes,1,opt,name=file,proto3" json:"file" bson:"file"`
}

func (x *DownloadFileReply) Reset() {
	*x = DownloadFileReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadFileReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadFileReply) ProtoMessage() {}

func (x *DownloadFileReply) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadFileReply.ProtoReflect.Descriptor instead.
func (*DownloadFileReply) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{4}
}

func (x *DownloadFileReply) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

type GetClubFileZipRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Club string `protobuf:"bytes,1,opt,name=club,proto3" json:"club" bson:"club"`
}

func (x *GetClubFileZipRequest) Reset() {
	*x = GetClubFileZipRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClubFileZipRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClubFileZipRequest) ProtoMessage() {}

func (x *GetClubFileZipRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClubFileZipRequest.ProtoReflect.Descriptor instead.
func (*GetClubFileZipRequest) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{5}
}

func (x *GetClubFileZipRequest) GetClub() string {
	if x != nil {
		return x.Club
	}
	return ""
}

type GetClubFileZipReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files []*File `protobuf:"bytes,1,rep,name=files,proto3" json:"files" bson:"files"`
}

func (x *GetClubFileZipReply) Reset() {
	*x = GetClubFileZipReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClubFileZipReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClubFileZipReply) ProtoMessage() {}

func (x *GetClubFileZipReply) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClubFileZipReply.ProtoReflect.Descriptor instead.
func (*GetClubFileZipReply) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{6}
}

func (x *GetClubFileZipReply) GetFiles() []*File {
	if x != nil {
		return x.Files
	}
	return nil
}

type DeleteFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
}

func (x *DeleteFileRequest) Reset() {
	*x = DeleteFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFileRequest) ProtoMessage() {}

func (x *DeleteFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFileRequest.ProtoReflect.Descriptor instead.
func (*DeleteFileRequest) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteFileRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteFileReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteFileReply) Reset() {
	*x = DeleteFileReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFileReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFileReply) ProtoMessage() {}

func (x *DeleteFileReply) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFileReply.ProtoReflect.Descriptor instead.
func (*DeleteFileReply) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{8}
}

var File_file_proto protoreflect.FileDescriptor

var file_file_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0x42, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x66, 0x69, 0x6c, 0x65, 0x22, 0x4f, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x21, 0x0a, 0x0f, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x27, 0x0a, 0x11, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x2b, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43,
	0x6c, 0x75, 0x62, 0x46, 0x69, 0x6c, 0x65, 0x5a, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6c, 0x75, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6c, 0x75, 0x62, 0x22, 0x35, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x62,
	0x46, 0x69, 0x6c, 0x65, 0x5a, 0x69, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1e, 0x0a, 0x05,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x23, 0x0a, 0x11,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x11, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x32, 0x93, 0x02, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69,
	0x6c, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x28, 0x01, 0x12, 0x42, 0x0a, 0x0c, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69,
	0x6c, 0x65, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62,
	0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x30, 0x01, 0x12, 0x46, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75,
	0x62, 0x46, 0x69, 0x6c, 0x65, 0x5a, 0x69, 0x70, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x6c, 0x75, 0x62, 0x46, 0x69, 0x6c, 0x65, 0x5a, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x62,
	0x46, 0x69, 0x6c, 0x65, 0x5a, 0x69, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3a,
	0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x15, 0x2e, 0x70,
	0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x75, 0x73, 0x68, 0x6b, 0x2f, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_file_proto_rawDescOnce sync.Once
	file_file_proto_rawDescData = file_file_proto_rawDesc
)

func file_file_proto_rawDescGZIP() []byte {
	file_file_proto_rawDescOnce.Do(func() {
		file_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_file_proto_rawDescData)
	})
	return file_file_proto_rawDescData
}

var file_file_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_file_proto_goTypes = []interface{}{
	(*File)(nil),                  // 0: pb.File
	(*UploadFileRequest)(nil),     // 1: pb.UploadFileRequest
	(*UploadFileReply)(nil),       // 2: pb.UploadFileReply
	(*DownloadFileRequest)(nil),   // 3: pb.DownloadFileRequest
	(*DownloadFileReply)(nil),     // 4: pb.DownloadFileReply
	(*GetClubFileZipRequest)(nil), // 5: pb.GetClubFileZipRequest
	(*GetClubFileZipReply)(nil),   // 6: pb.GetClubFileZipReply
	(*DeleteFileRequest)(nil),     // 7: pb.DeleteFileRequest
	(*DeleteFileReply)(nil),       // 8: pb.DeleteFileReply
}
var file_file_proto_depIdxs = []int32{
	0, // 0: pb.GetClubFileZipReply.files:type_name -> pb.File
	1, // 1: pb.FileManager.UploadFile:input_type -> pb.UploadFileRequest
	3, // 2: pb.FileManager.DownloadFile:input_type -> pb.DownloadFileRequest
	5, // 3: pb.FileManager.GetClubFileZip:input_type -> pb.GetClubFileZipRequest
	7, // 4: pb.FileManager.DeleteFile:input_type -> pb.DeleteFileRequest
	2, // 5: pb.FileManager.UploadFile:output_type -> pb.UploadFileReply
	4, // 6: pb.FileManager.DownloadFile:output_type -> pb.DownloadFileReply
	6, // 7: pb.FileManager.GetClubFileZip:output_type -> pb.GetClubFileZipReply
	8, // 8: pb.FileManager.DeleteFile:output_type -> pb.DeleteFileReply
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_file_proto_init() }
func file_file_proto_init() {
	if File_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileRequest); i {
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
		file_file_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileReply); i {
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
		file_file_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadFileRequest); i {
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
		file_file_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadFileReply); i {
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
		file_file_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClubFileZipRequest); i {
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
		file_file_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClubFileZipReply); i {
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
		file_file_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFileRequest); i {
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
		file_file_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFileReply); i {
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
			RawDescriptor: file_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_file_proto_goTypes,
		DependencyIndexes: file_file_proto_depIdxs,
		MessageInfos:      file_file_proto_msgTypes,
	}.Build()
	File_file_proto = out.File
	file_file_proto_rawDesc = nil
	file_file_proto_goTypes = nil
	file_file_proto_depIdxs = nil
}
