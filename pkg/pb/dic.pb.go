// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: dic.proto

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

type SportType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
	Label    string `protobuf:"bytes,2,opt,name=label,proto3" json:"label" bson:"label"`
	Code     string `protobuf:"bytes,3,opt,name=code,proto3" json:"code" bson:"code"`
	Level    string `protobuf:"bytes,4,opt,name=level,proto3" json:"level" bson:"level"`
	ParentId string `protobuf:"bytes,5,opt,name=parent_id,json=parentId,proto3" json:"parent_id" bson:"parent_id"`
}

func (x *SportType) Reset() {
	*x = SportType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SportType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SportType) ProtoMessage() {}

func (x *SportType) ProtoReflect() protoreflect.Message {
	mi := &file_dic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SportType.ProtoReflect.Descriptor instead.
func (*SportType) Descriptor() ([]byte, []int) {
	return file_dic_proto_rawDescGZIP(), []int{0}
}

func (x *SportType) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SportType) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *SportType) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *SportType) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *SportType) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

type CreateSportTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *SportType `protobuf:"bytes,1,opt,name=data,proto3" json:"data" bson:"data"`
}

func (x *CreateSportTypeRequest) Reset() {
	*x = CreateSportTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSportTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSportTypeRequest) ProtoMessage() {}

func (x *CreateSportTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSportTypeRequest.ProtoReflect.Descriptor instead.
func (*CreateSportTypeRequest) Descriptor() ([]byte, []int) {
	return file_dic_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSportTypeRequest) GetData() *SportType {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetSportTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
}

func (x *GetSportTypeRequest) Reset() {
	*x = GetSportTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSportTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSportTypeRequest) ProtoMessage() {}

func (x *GetSportTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSportTypeRequest.ProtoReflect.Descriptor instead.
func (*GetSportTypeRequest) Descriptor() ([]byte, []int) {
	return file_dic_proto_rawDescGZIP(), []int{2}
}

func (x *GetSportTypeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateSportTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
	Data *SportType `protobuf:"bytes,2,opt,name=data,proto3" json:"data" bson:"data"`
}

func (x *UpdateSportTypeRequest) Reset() {
	*x = UpdateSportTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dic_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSportTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSportTypeRequest) ProtoMessage() {}

func (x *UpdateSportTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dic_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSportTypeRequest.ProtoReflect.Descriptor instead.
func (*UpdateSportTypeRequest) Descriptor() ([]byte, []int) {
	return file_dic_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateSportTypeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateSportTypeRequest) GetData() *SportType {
	if x != nil {
		return x.Data
	}
	return nil
}

type DeleteSportTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
}

func (x *DeleteSportTypeRequest) Reset() {
	*x = DeleteSportTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dic_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSportTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSportTypeRequest) ProtoMessage() {}

func (x *DeleteSportTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dic_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSportTypeRequest.ProtoReflect.Descriptor instead.
func (*DeleteSportTypeRequest) Descriptor() ([]byte, []int) {
	return file_dic_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteSportTypeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetSportTypesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query" bson:"query"`
	Skip  int64  `protobuf:"varint,2,opt,name=skip,proto3" json:"skip" bson:"skip"`
	Limit int64  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit" bson:"limit"`
}

func (x *GetSportTypesRequest) Reset() {
	*x = GetSportTypesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dic_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSportTypesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSportTypesRequest) ProtoMessage() {}

func (x *GetSportTypesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dic_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSportTypesRequest.ProtoReflect.Descriptor instead.
func (*GetSportTypesRequest) Descriptor() ([]byte, []int) {
	return file_dic_proto_rawDescGZIP(), []int{5}
}

func (x *GetSportTypesRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *GetSportTypesRequest) GetSkip() int64 {
	if x != nil {
		return x.Skip
	}
	return 0
}

func (x *GetSportTypesRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetSportTypesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount int64        `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count" bson:"total_count"`
	Items      []*SportType `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *GetSportTypesReply) Reset() {
	*x = GetSportTypesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dic_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSportTypesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSportTypesReply) ProtoMessage() {}

func (x *GetSportTypesReply) ProtoReflect() protoreflect.Message {
	mi := &file_dic_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSportTypesReply.ProtoReflect.Descriptor instead.
func (*GetSportTypesReply) Descriptor() ([]byte, []int) {
	return file_dic_proto_rawDescGZIP(), []int{6}
}

func (x *GetSportTypesReply) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *GetSportTypesReply) GetItems() []*SportType {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_dic_proto protoreflect.FileDescriptor

var file_dic_proto_rawDesc = []byte{
	0x0a, 0x09, 0x64, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22,
	0x78, 0x0a, 0x09, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x3b, 0x0a, 0x16, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x70, 0x6f,
	0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4b, 0x0a,
	0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x70, 0x6f, 0x72, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x28, 0x0a, 0x16, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x56, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x70, 0x6f, 0x72, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6b, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x73, 0x6b, 0x69, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x5a, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xd1, 0x02, 0x0a, 0x10, 0x53, 0x70, 0x6f,
	0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x3e, 0x0a,
	0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x70, 0x6f, 0x72,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x2e,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x70, 0x6f, 0x72,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x70, 0x6f, 0x72,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x70, 0x6f, 0x72,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x70,
	0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x70, 0x6f, 0x72, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x75, 0x73, 0x68, 0x6b,
	0x2f, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dic_proto_rawDescOnce sync.Once
	file_dic_proto_rawDescData = file_dic_proto_rawDesc
)

func file_dic_proto_rawDescGZIP() []byte {
	file_dic_proto_rawDescOnce.Do(func() {
		file_dic_proto_rawDescData = protoimpl.X.CompressGZIP(file_dic_proto_rawDescData)
	})
	return file_dic_proto_rawDescData
}

var file_dic_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_dic_proto_goTypes = []interface{}{
	(*SportType)(nil),              // 0: pb.SportType
	(*CreateSportTypeRequest)(nil), // 1: pb.CreateSportTypeRequest
	(*GetSportTypeRequest)(nil),    // 2: pb.GetSportTypeRequest
	(*UpdateSportTypeRequest)(nil), // 3: pb.UpdateSportTypeRequest
	(*DeleteSportTypeRequest)(nil), // 4: pb.DeleteSportTypeRequest
	(*GetSportTypesRequest)(nil),   // 5: pb.GetSportTypesRequest
	(*GetSportTypesReply)(nil),     // 6: pb.GetSportTypesReply
}
var file_dic_proto_depIdxs = []int32{
	0, // 0: pb.CreateSportTypeRequest.data:type_name -> pb.SportType
	0, // 1: pb.UpdateSportTypeRequest.data:type_name -> pb.SportType
	0, // 2: pb.GetSportTypesReply.items:type_name -> pb.SportType
	1, // 3: pb.SportTypeManager.CreateSportType:input_type -> pb.CreateSportTypeRequest
	2, // 4: pb.SportTypeManager.GetSportType:input_type -> pb.GetSportTypeRequest
	3, // 5: pb.SportTypeManager.UpdateSportType:input_type -> pb.UpdateSportTypeRequest
	4, // 6: pb.SportTypeManager.DeleteSportType:input_type -> pb.DeleteSportTypeRequest
	5, // 7: pb.SportTypeManager.GetSportTypes:input_type -> pb.GetSportTypesRequest
	0, // 8: pb.SportTypeManager.CreateSportType:output_type -> pb.SportType
	0, // 9: pb.SportTypeManager.GetSportType:output_type -> pb.SportType
	0, // 10: pb.SportTypeManager.UpdateSportType:output_type -> pb.SportType
	0, // 11: pb.SportTypeManager.DeleteSportType:output_type -> pb.SportType
	6, // 12: pb.SportTypeManager.GetSportTypes:output_type -> pb.GetSportTypesReply
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_dic_proto_init() }
func file_dic_proto_init() {
	if File_dic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SportType); i {
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
		file_dic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSportTypeRequest); i {
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
		file_dic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSportTypeRequest); i {
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
		file_dic_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSportTypeRequest); i {
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
		file_dic_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSportTypeRequest); i {
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
		file_dic_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSportTypesRequest); i {
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
		file_dic_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSportTypesReply); i {
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
			RawDescriptor: file_dic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dic_proto_goTypes,
		DependencyIndexes: file_dic_proto_depIdxs,
		MessageInfos:      file_dic_proto_msgTypes,
	}.Build()
	File_dic_proto = out.File
	file_dic_proto_rawDesc = nil
	file_dic_proto_goTypes = nil
	file_dic_proto_depIdxs = nil
}
