// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: personal.proto

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

type Personal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
	Files    []*FileInfo `protobuf:"bytes,2,rep,name=files,proto3" json:"files" bson:"files"`
	Type     string      `protobuf:"bytes,3,opt,name=type,proto3" json:"type" bson:"type"`
	Userid   string      `protobuf:"bytes,4,opt,name=userid,proto3" json:"userid" bson:"userid"`
	Address  string      `protobuf:"bytes,5,opt,name=address,proto3" json:"address" bson:"address"`
	Name     string      `protobuf:"bytes,6,opt,name=name,proto3" json:"name" bson:"name"`
	ClubName string      `protobuf:"bytes,7,opt,name=club_name,json=clubName,proto3" json:"club_name" bson:"club_name"`
}

func (x *Personal) Reset() {
	*x = Personal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Personal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Personal) ProtoMessage() {}

func (x *Personal) ProtoReflect() protoreflect.Message {
	mi := &file_personal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Personal.ProtoReflect.Descriptor instead.
func (*Personal) Descriptor() ([]byte, []int) {
	return file_personal_proto_rawDescGZIP(), []int{0}
}

func (x *Personal) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Personal) GetFiles() []*FileInfo {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *Personal) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Personal) GetUserid() string {
	if x != nil {
		return x.Userid
	}
	return ""
}

func (x *Personal) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Personal) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Personal) GetClubName() string {
	if x != nil {
		return x.ClubName
	}
	return ""
}

type CreatePersonalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Personal `protobuf:"bytes,1,opt,name=data,proto3" json:"data" bson:"data"`
}

func (x *CreatePersonalRequest) Reset() {
	*x = CreatePersonalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePersonalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePersonalRequest) ProtoMessage() {}

func (x *CreatePersonalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_personal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePersonalRequest.ProtoReflect.Descriptor instead.
func (*CreatePersonalRequest) Descriptor() ([]byte, []int) {
	return file_personal_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePersonalRequest) GetData() *Personal {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetPersonalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
}

func (x *GetPersonalRequest) Reset() {
	*x = GetPersonalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersonalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonalRequest) ProtoMessage() {}

func (x *GetPersonalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_personal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonalRequest.ProtoReflect.Descriptor instead.
func (*GetPersonalRequest) Descriptor() ([]byte, []int) {
	return file_personal_proto_rawDescGZIP(), []int{2}
}

func (x *GetPersonalRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdatePersonalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
	Data *Personal `protobuf:"bytes,2,opt,name=data,proto3" json:"data" bson:"data"`
}

func (x *UpdatePersonalRequest) Reset() {
	*x = UpdatePersonalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personal_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePersonalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePersonalRequest) ProtoMessage() {}

func (x *UpdatePersonalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_personal_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePersonalRequest.ProtoReflect.Descriptor instead.
func (*UpdatePersonalRequest) Descriptor() ([]byte, []int) {
	return file_personal_proto_rawDescGZIP(), []int{3}
}

func (x *UpdatePersonalRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdatePersonalRequest) GetData() *Personal {
	if x != nil {
		return x.Data
	}
	return nil
}

type DeletePersonalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
}

func (x *DeletePersonalRequest) Reset() {
	*x = DeletePersonalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personal_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePersonalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePersonalRequest) ProtoMessage() {}

func (x *DeletePersonalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_personal_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePersonalRequest.ProtoReflect.Descriptor instead.
func (*DeletePersonalRequest) Descriptor() ([]byte, []int) {
	return file_personal_proto_rawDescGZIP(), []int{4}
}

func (x *DeletePersonalRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetPersonalsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query" bson:"query"`
	Skip  int64  `protobuf:"varint,2,opt,name=skip,proto3" json:"skip" bson:"skip"`
	Limit int64  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit" bson:"limit"`
}

func (x *GetPersonalsRequest) Reset() {
	*x = GetPersonalsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personal_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersonalsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonalsRequest) ProtoMessage() {}

func (x *GetPersonalsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_personal_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonalsRequest.ProtoReflect.Descriptor instead.
func (*GetPersonalsRequest) Descriptor() ([]byte, []int) {
	return file_personal_proto_rawDescGZIP(), []int{5}
}

func (x *GetPersonalsRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *GetPersonalsRequest) GetSkip() int64 {
	if x != nil {
		return x.Skip
	}
	return 0
}

func (x *GetPersonalsRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetPersonalsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount int64       `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count" bson:"total_count"`
	Items      []*Personal `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *GetPersonalsReply) Reset() {
	*x = GetPersonalsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personal_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersonalsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonalsReply) ProtoMessage() {}

func (x *GetPersonalsReply) ProtoReflect() protoreflect.Message {
	mi := &file_personal_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonalsReply.ProtoReflect.Descriptor instead.
func (*GetPersonalsReply) Descriptor() ([]byte, []int) {
	return file_personal_proto_rawDescGZIP(), []int{6}
}

func (x *GetPersonalsReply) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *GetPersonalsReply) GetItems() []*Personal {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_personal_proto protoreflect.FileDescriptor

var file_personal_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb5, 0x01, 0x0a, 0x08, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a,
	0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70,
	0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x6c, 0x75, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x6c, 0x75, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x39, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x49, 0x0a, 0x15, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x20, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x55, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6b,
	0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x6b, 0x69, 0x70, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x22, 0x58, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xc1,
	0x02, 0x0a, 0x0f, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x12, 0x3b, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x22, 0x00, 0x12,
	0x35, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x16,
	0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x61, 0x6c, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61,
	0x6c, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x22, 0x00,
	0x12, 0x40, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x73,
	0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61,
	0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x79, 0x75, 0x73, 0x68, 0x6b, 0x2f, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x3b,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_personal_proto_rawDescOnce sync.Once
	file_personal_proto_rawDescData = file_personal_proto_rawDesc
)

func file_personal_proto_rawDescGZIP() []byte {
	file_personal_proto_rawDescOnce.Do(func() {
		file_personal_proto_rawDescData = protoimpl.X.CompressGZIP(file_personal_proto_rawDescData)
	})
	return file_personal_proto_rawDescData
}

var file_personal_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_personal_proto_goTypes = []interface{}{
	(*Personal)(nil),              // 0: pb.Personal
	(*CreatePersonalRequest)(nil), // 1: pb.CreatePersonalRequest
	(*GetPersonalRequest)(nil),    // 2: pb.GetPersonalRequest
	(*UpdatePersonalRequest)(nil), // 3: pb.UpdatePersonalRequest
	(*DeletePersonalRequest)(nil), // 4: pb.DeletePersonalRequest
	(*GetPersonalsRequest)(nil),   // 5: pb.GetPersonalsRequest
	(*GetPersonalsReply)(nil),     // 6: pb.GetPersonalsReply
	(*FileInfo)(nil),              // 7: pb.FileInfo
}
var file_personal_proto_depIdxs = []int32{
	7, // 0: pb.Personal.files:type_name -> pb.FileInfo
	0, // 1: pb.CreatePersonalRequest.data:type_name -> pb.Personal
	0, // 2: pb.UpdatePersonalRequest.data:type_name -> pb.Personal
	0, // 3: pb.GetPersonalsReply.items:type_name -> pb.Personal
	1, // 4: pb.PersonalManager.CreatePersonal:input_type -> pb.CreatePersonalRequest
	2, // 5: pb.PersonalManager.GetPersonal:input_type -> pb.GetPersonalRequest
	3, // 6: pb.PersonalManager.UpdatePersonal:input_type -> pb.UpdatePersonalRequest
	4, // 7: pb.PersonalManager.DeletePersonal:input_type -> pb.DeletePersonalRequest
	5, // 8: pb.PersonalManager.GetPersonals:input_type -> pb.GetPersonalsRequest
	0, // 9: pb.PersonalManager.CreatePersonal:output_type -> pb.Personal
	0, // 10: pb.PersonalManager.GetPersonal:output_type -> pb.Personal
	0, // 11: pb.PersonalManager.UpdatePersonal:output_type -> pb.Personal
	0, // 12: pb.PersonalManager.DeletePersonal:output_type -> pb.Personal
	6, // 13: pb.PersonalManager.GetPersonals:output_type -> pb.GetPersonalsReply
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_personal_proto_init() }
func file_personal_proto_init() {
	if File_personal_proto != nil {
		return
	}
	file_work_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_personal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Personal); i {
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
		file_personal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePersonalRequest); i {
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
		file_personal_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersonalRequest); i {
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
		file_personal_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePersonalRequest); i {
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
		file_personal_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePersonalRequest); i {
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
		file_personal_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersonalsRequest); i {
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
		file_personal_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersonalsReply); i {
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
			RawDescriptor: file_personal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_personal_proto_goTypes,
		DependencyIndexes: file_personal_proto_depIdxs,
		MessageInfos:      file_personal_proto_msgTypes,
	}.Build()
	File_personal_proto = out.File
	file_personal_proto_rawDesc = nil
	file_personal_proto_goTypes = nil
	file_personal_proto_depIdxs = nil
}