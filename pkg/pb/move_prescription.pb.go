// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: move_prescription.proto

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

type MovePrescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
	Userid string `protobuf:"bytes,2,opt,name=userid,proto3" json:"userid" bson:"userid"`
}

func (x *MovePrescription) Reset() {
	*x = MovePrescription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_move_prescription_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovePrescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovePrescription) ProtoMessage() {}

func (x *MovePrescription) ProtoReflect() protoreflect.Message {
	mi := &file_move_prescription_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovePrescription.ProtoReflect.Descriptor instead.
func (*MovePrescription) Descriptor() ([]byte, []int) {
	return file_move_prescription_proto_rawDescGZIP(), []int{0}
}

func (x *MovePrescription) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MovePrescription) GetUserid() string {
	if x != nil {
		return x.Userid
	}
	return ""
}

type CreateMovePrescriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *MovePrescription `protobuf:"bytes,1,opt,name=data,proto3" json:"data" bson:"data"`
}

func (x *CreateMovePrescriptionRequest) Reset() {
	*x = CreateMovePrescriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_move_prescription_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMovePrescriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMovePrescriptionRequest) ProtoMessage() {}

func (x *CreateMovePrescriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_move_prescription_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMovePrescriptionRequest.ProtoReflect.Descriptor instead.
func (*CreateMovePrescriptionRequest) Descriptor() ([]byte, []int) {
	return file_move_prescription_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMovePrescriptionRequest) GetData() *MovePrescription {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetMovePrescriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
}

func (x *GetMovePrescriptionRequest) Reset() {
	*x = GetMovePrescriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_move_prescription_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMovePrescriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovePrescriptionRequest) ProtoMessage() {}

func (x *GetMovePrescriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_move_prescription_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovePrescriptionRequest.ProtoReflect.Descriptor instead.
func (*GetMovePrescriptionRequest) Descriptor() ([]byte, []int) {
	return file_move_prescription_proto_rawDescGZIP(), []int{2}
}

func (x *GetMovePrescriptionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateMovePrescriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
	Data *MovePrescription `protobuf:"bytes,2,opt,name=data,proto3" json:"data" bson:"data"`
}

func (x *UpdateMovePrescriptionRequest) Reset() {
	*x = UpdateMovePrescriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_move_prescription_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMovePrescriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMovePrescriptionRequest) ProtoMessage() {}

func (x *UpdateMovePrescriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_move_prescription_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMovePrescriptionRequest.ProtoReflect.Descriptor instead.
func (*UpdateMovePrescriptionRequest) Descriptor() ([]byte, []int) {
	return file_move_prescription_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateMovePrescriptionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateMovePrescriptionRequest) GetData() *MovePrescription {
	if x != nil {
		return x.Data
	}
	return nil
}

type DeleteMovePrescriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
}

func (x *DeleteMovePrescriptionRequest) Reset() {
	*x = DeleteMovePrescriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_move_prescription_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMovePrescriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMovePrescriptionRequest) ProtoMessage() {}

func (x *DeleteMovePrescriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_move_prescription_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMovePrescriptionRequest.ProtoReflect.Descriptor instead.
func (*DeleteMovePrescriptionRequest) Descriptor() ([]byte, []int) {
	return file_move_prescription_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteMovePrescriptionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetMovePrescriptionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query" bson:"query"`
	Skip  int64  `protobuf:"varint,2,opt,name=skip,proto3" json:"skip" bson:"skip"`
	Limit int64  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit" bson:"limit"`
}

func (x *GetMovePrescriptionsRequest) Reset() {
	*x = GetMovePrescriptionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_move_prescription_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMovePrescriptionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovePrescriptionsRequest) ProtoMessage() {}

func (x *GetMovePrescriptionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_move_prescription_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovePrescriptionsRequest.ProtoReflect.Descriptor instead.
func (*GetMovePrescriptionsRequest) Descriptor() ([]byte, []int) {
	return file_move_prescription_proto_rawDescGZIP(), []int{5}
}

func (x *GetMovePrescriptionsRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *GetMovePrescriptionsRequest) GetSkip() int64 {
	if x != nil {
		return x.Skip
	}
	return 0
}

func (x *GetMovePrescriptionsRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetMovePrescriptionsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount int64               `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count" bson:"total_count"`
	Items      []*MovePrescription `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *GetMovePrescriptionsReply) Reset() {
	*x = GetMovePrescriptionsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_move_prescription_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMovePrescriptionsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovePrescriptionsReply) ProtoMessage() {}

func (x *GetMovePrescriptionsReply) ProtoReflect() protoreflect.Message {
	mi := &file_move_prescription_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovePrescriptionsReply.ProtoReflect.Descriptor instead.
func (*GetMovePrescriptionsReply) Descriptor() ([]byte, []int) {
	return file_move_prescription_proto_rawDescGZIP(), []int{6}
}

func (x *GetMovePrescriptionsReply) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *GetMovePrescriptionsReply) GetItems() []*MovePrescription {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_move_prescription_proto protoreflect.FileDescriptor

var file_move_prescription_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x3a, 0x0a,
	0x10, 0x4d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x22, 0x49, 0x0a, 0x1d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x6f,
	0x76, 0x65, 0x50, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x2c, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x65, 0x50,
	0x72, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x59, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x65,
	0x50, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2f, 0x0a,
	0x1d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5d,
	0x0a, 0x1b, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6b, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x73, 0x6b, 0x69, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x68, 0x0a,
	0x19, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x2e,
	0x4d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xc1, 0x03, 0x0a, 0x17, 0x4d, 0x6f, 0x76, 0x65,
	0x50, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x12, 0x53, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76,
	0x65, 0x50, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d,
	0x6f, 0x76, 0x65, 0x50, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76,
	0x65, 0x50, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x50, 0x72,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x16,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4d,
	0x6f, 0x76, 0x65, 0x50, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x00, 0x12, 0x58, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x4d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x75, 0x73, 0x68, 0x6b, 0x2f,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_move_prescription_proto_rawDescOnce sync.Once
	file_move_prescription_proto_rawDescData = file_move_prescription_proto_rawDesc
)

func file_move_prescription_proto_rawDescGZIP() []byte {
	file_move_prescription_proto_rawDescOnce.Do(func() {
		file_move_prescription_proto_rawDescData = protoimpl.X.CompressGZIP(file_move_prescription_proto_rawDescData)
	})
	return file_move_prescription_proto_rawDescData
}

var file_move_prescription_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_move_prescription_proto_goTypes = []interface{}{
	(*MovePrescription)(nil),              // 0: pb.MovePrescription
	(*CreateMovePrescriptionRequest)(nil), // 1: pb.CreateMovePrescriptionRequest
	(*GetMovePrescriptionRequest)(nil),    // 2: pb.GetMovePrescriptionRequest
	(*UpdateMovePrescriptionRequest)(nil), // 3: pb.UpdateMovePrescriptionRequest
	(*DeleteMovePrescriptionRequest)(nil), // 4: pb.DeleteMovePrescriptionRequest
	(*GetMovePrescriptionsRequest)(nil),   // 5: pb.GetMovePrescriptionsRequest
	(*GetMovePrescriptionsReply)(nil),     // 6: pb.GetMovePrescriptionsReply
}
var file_move_prescription_proto_depIdxs = []int32{
	0, // 0: pb.CreateMovePrescriptionRequest.data:type_name -> pb.MovePrescription
	0, // 1: pb.UpdateMovePrescriptionRequest.data:type_name -> pb.MovePrescription
	0, // 2: pb.GetMovePrescriptionsReply.items:type_name -> pb.MovePrescription
	1, // 3: pb.MovePrescriptionManager.CreateMovePrescription:input_type -> pb.CreateMovePrescriptionRequest
	2, // 4: pb.MovePrescriptionManager.GetMovePrescription:input_type -> pb.GetMovePrescriptionRequest
	3, // 5: pb.MovePrescriptionManager.UpdateMovePrescription:input_type -> pb.UpdateMovePrescriptionRequest
	4, // 6: pb.MovePrescriptionManager.DeleteMovePrescription:input_type -> pb.DeleteMovePrescriptionRequest
	5, // 7: pb.MovePrescriptionManager.GetMovePrescriptions:input_type -> pb.GetMovePrescriptionsRequest
	0, // 8: pb.MovePrescriptionManager.CreateMovePrescription:output_type -> pb.MovePrescription
	0, // 9: pb.MovePrescriptionManager.GetMovePrescription:output_type -> pb.MovePrescription
	0, // 10: pb.MovePrescriptionManager.UpdateMovePrescription:output_type -> pb.MovePrescription
	0, // 11: pb.MovePrescriptionManager.DeleteMovePrescription:output_type -> pb.MovePrescription
	6, // 12: pb.MovePrescriptionManager.GetMovePrescriptions:output_type -> pb.GetMovePrescriptionsReply
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_move_prescription_proto_init() }
func file_move_prescription_proto_init() {
	if File_move_prescription_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_move_prescription_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovePrescription); i {
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
		file_move_prescription_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMovePrescriptionRequest); i {
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
		file_move_prescription_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMovePrescriptionRequest); i {
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
		file_move_prescription_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMovePrescriptionRequest); i {
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
		file_move_prescription_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMovePrescriptionRequest); i {
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
		file_move_prescription_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMovePrescriptionsRequest); i {
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
		file_move_prescription_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMovePrescriptionsReply); i {
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
			RawDescriptor: file_move_prescription_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_move_prescription_proto_goTypes,
		DependencyIndexes: file_move_prescription_proto_depIdxs,
		MessageInfos:      file_move_prescription_proto_msgTypes,
	}.Build()
	File_move_prescription_proto = out.File
	file_move_prescription_proto_rawDesc = nil
	file_move_prescription_proto_goTypes = nil
	file_move_prescription_proto_depIdxs = nil
}
