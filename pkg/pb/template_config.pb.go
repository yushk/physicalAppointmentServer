// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: template_config.proto

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

type Evaluate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" bson:"name"`
	Excellent string `protobuf:"bytes,3,opt,name=excellent,proto3" json:"excellent" bson:"excellent"`
	Good      string `protobuf:"bytes,4,opt,name=good,proto3" json:"good" bson:"good"`
	Pass      string `protobuf:"bytes,5,opt,name=pass,proto3" json:"pass" bson:"pass"`
	Flunk     string `protobuf:"bytes,6,opt,name=flunk,proto3" json:"flunk" bson:"flunk"`
}

func (x *Evaluate) Reset() {
	*x = Evaluate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Evaluate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Evaluate) ProtoMessage() {}

func (x *Evaluate) ProtoReflect() protoreflect.Message {
	mi := &file_template_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Evaluate.ProtoReflect.Descriptor instead.
func (*Evaluate) Descriptor() ([]byte, []int) {
	return file_template_config_proto_rawDescGZIP(), []int{0}
}

func (x *Evaluate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Evaluate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Evaluate) GetExcellent() string {
	if x != nil {
		return x.Excellent
	}
	return ""
}

func (x *Evaluate) GetGood() string {
	if x != nil {
		return x.Good
	}
	return ""
}

func (x *Evaluate) GetPass() string {
	if x != nil {
		return x.Pass
	}
	return ""
}

func (x *Evaluate) GetFlunk() string {
	if x != nil {
		return x.Flunk
	}
	return ""
}

type QuestionOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Score string `protobuf:"bytes,1,opt,name=score,proto3" json:"score" bson:"score"`
	Desc  string `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc" bson:"desc"`
	Id    string `protobuf:"bytes,3,opt,name=id,proto3" json:"id" bson:"id"`
}

func (x *QuestionOptions) Reset() {
	*x = QuestionOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestionOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionOptions) ProtoMessage() {}

func (x *QuestionOptions) ProtoReflect() protoreflect.Message {
	mi := &file_template_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionOptions.ProtoReflect.Descriptor instead.
func (*QuestionOptions) Descriptor() ([]byte, []int) {
	return file_template_config_proto_rawDescGZIP(), []int{1}
}

func (x *QuestionOptions) GetScore() string {
	if x != nil {
		return x.Score
	}
	return ""
}

func (x *QuestionOptions) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *QuestionOptions) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SportItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
	Label           string             `protobuf:"bytes,2,opt,name=label,proto3" json:"label" bson:"label"`
	First           string             `protobuf:"bytes,3,opt,name=first,proto3" json:"first" bson:"first"`
	Second          string             `protobuf:"bytes,4,opt,name=second,proto3" json:"second" bson:"second"`
	UploadFile      string             `protobuf:"bytes,5,opt,name=upload_file,json=uploadFile,proto3" json:"upload_file" bson:"upload_file"`
	FileType        string             `protobuf:"bytes,6,opt,name=file_type,json=fileType,proto3" json:"file_type" bson:"file_type"`
	Approver        string             `protobuf:"bytes,7,opt,name=approver,proto3" json:"approver" bson:"approver"`
	FileName        string             `protobuf:"bytes,8,opt,name=file_name,json=fileName,proto3" json:"file_name" bson:"file_name"`
	CreateTime      int64              `protobuf:"varint,9,opt,name=create_time,json=createTime,proto3" json:"create_time" bson:"create_time"`
	UpdateTime      int64              `protobuf:"varint,10,opt,name=update_time,json=updateTime,proto3" json:"update_time" bson:"update_time"`
	QuestionOptions []*QuestionOptions `protobuf:"bytes,11,rep,name=question_options,json=questionOptions,proto3" json:"question_options" bson:"question_options"`
	Version         string             `protobuf:"bytes,12,opt,name=version,proto3" json:"version" bson:"version"`
}

func (x *SportItem) Reset() {
	*x = SportItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SportItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SportItem) ProtoMessage() {}

func (x *SportItem) ProtoReflect() protoreflect.Message {
	mi := &file_template_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SportItem.ProtoReflect.Descriptor instead.
func (*SportItem) Descriptor() ([]byte, []int) {
	return file_template_config_proto_rawDescGZIP(), []int{2}
}

func (x *SportItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SportItem) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *SportItem) GetFirst() string {
	if x != nil {
		return x.First
	}
	return ""
}

func (x *SportItem) GetSecond() string {
	if x != nil {
		return x.Second
	}
	return ""
}

func (x *SportItem) GetUploadFile() string {
	if x != nil {
		return x.UploadFile
	}
	return ""
}

func (x *SportItem) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

func (x *SportItem) GetApprover() string {
	if x != nil {
		return x.Approver
	}
	return ""
}

func (x *SportItem) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *SportItem) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *SportItem) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *SportItem) GetQuestionOptions() []*QuestionOptions {
	if x != nil {
		return x.QuestionOptions
	}
	return nil
}

func (x *SportItem) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type CreateSportItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *SportItem `protobuf:"bytes,1,opt,name=data,proto3" json:"data" bson:"data"`
}

func (x *CreateSportItemRequest) Reset() {
	*x = CreateSportItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSportItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSportItemRequest) ProtoMessage() {}

func (x *CreateSportItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_template_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSportItemRequest.ProtoReflect.Descriptor instead.
func (*CreateSportItemRequest) Descriptor() ([]byte, []int) {
	return file_template_config_proto_rawDescGZIP(), []int{3}
}

func (x *CreateSportItemRequest) GetData() *SportItem {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetSportItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
}

func (x *GetSportItemRequest) Reset() {
	*x = GetSportItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSportItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSportItemRequest) ProtoMessage() {}

func (x *GetSportItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_template_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSportItemRequest.ProtoReflect.Descriptor instead.
func (*GetSportItemRequest) Descriptor() ([]byte, []int) {
	return file_template_config_proto_rawDescGZIP(), []int{4}
}

func (x *GetSportItemRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateSportItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
	Data *SportItem `protobuf:"bytes,2,opt,name=data,proto3" json:"data" bson:"data"`
}

func (x *UpdateSportItemRequest) Reset() {
	*x = UpdateSportItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSportItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSportItemRequest) ProtoMessage() {}

func (x *UpdateSportItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_template_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSportItemRequest.ProtoReflect.Descriptor instead.
func (*UpdateSportItemRequest) Descriptor() ([]byte, []int) {
	return file_template_config_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateSportItemRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateSportItemRequest) GetData() *SportItem {
	if x != nil {
		return x.Data
	}
	return nil
}

type DeleteSportItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
}

func (x *DeleteSportItemRequest) Reset() {
	*x = DeleteSportItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_config_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSportItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSportItemRequest) ProtoMessage() {}

func (x *DeleteSportItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_template_config_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSportItemRequest.ProtoReflect.Descriptor instead.
func (*DeleteSportItemRequest) Descriptor() ([]byte, []int) {
	return file_template_config_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteSportItemRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetSportItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query" bson:"query"`
	Skip  int64  `protobuf:"varint,2,opt,name=skip,proto3" json:"skip" bson:"skip"`
	Limit int64  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit" bson:"limit"`
}

func (x *GetSportItemsRequest) Reset() {
	*x = GetSportItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_config_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSportItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSportItemsRequest) ProtoMessage() {}

func (x *GetSportItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_template_config_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSportItemsRequest.ProtoReflect.Descriptor instead.
func (*GetSportItemsRequest) Descriptor() ([]byte, []int) {
	return file_template_config_proto_rawDescGZIP(), []int{7}
}

func (x *GetSportItemsRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *GetSportItemsRequest) GetSkip() int64 {
	if x != nil {
		return x.Skip
	}
	return 0
}

func (x *GetSportItemsRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetSportItemsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount int64        `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count" bson:"total_count"`
	Items      []*SportItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *GetSportItemsReply) Reset() {
	*x = GetSportItemsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_config_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSportItemsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSportItemsReply) ProtoMessage() {}

func (x *GetSportItemsReply) ProtoReflect() protoreflect.Message {
	mi := &file_template_config_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSportItemsReply.ProtoReflect.Descriptor instead.
func (*GetSportItemsReply) Descriptor() ([]byte, []int) {
	return file_template_config_proto_rawDescGZIP(), []int{8}
}

func (x *GetSportItemsReply) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *GetSportItemsReply) GetItems() []*SportItem {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_template_config_proto protoreflect.FileDescriptor

var file_template_config_proto_rawDesc = []byte{
	0x0a, 0x15, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x8a, 0x01, 0x0a, 0x08,
	0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x65, 0x78, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x65, 0x78, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x6f,
	0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x67, 0x6f, 0x6f, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x75, 0x6e, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x66, 0x6c, 0x75, 0x6e, 0x6b, 0x22, 0x4b, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xf2, 0x02, 0x0a, 0x09, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65,
	0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65,
	0x72, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x3e, 0x0a, 0x10, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x2e,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x0f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x3b, 0x0a, 0x16, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x70,
	0x6f, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4b,
	0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x70, 0x6f, 0x72,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x28, 0x0a, 0x16, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x56, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x70, 0x6f, 0x72,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6b, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x73, 0x6b, 0x69, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x5a, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xd1, 0x02, 0x0a, 0x10, 0x53, 0x70,
	0x6f, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x3e,
	0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x70, 0x6f,
	0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x00, 0x12, 0x38,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x17,
	0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x70, 0x6f,
	0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1a, 0x2e, 0x70, 0x62,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x70, 0x6f,
	0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1a, 0x2e, 0x70, 0x62,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x70, 0x6f,
	0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53,
	0x70, 0x6f, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x70, 0x6f, 0x72,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x2e, 0x5a,
	0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x75, 0x73, 0x68,
	0x6b, 0x2f, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_template_config_proto_rawDescOnce sync.Once
	file_template_config_proto_rawDescData = file_template_config_proto_rawDesc
)

func file_template_config_proto_rawDescGZIP() []byte {
	file_template_config_proto_rawDescOnce.Do(func() {
		file_template_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_template_config_proto_rawDescData)
	})
	return file_template_config_proto_rawDescData
}

var file_template_config_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_template_config_proto_goTypes = []interface{}{
	(*Evaluate)(nil),               // 0: pb.Evaluate
	(*QuestionOptions)(nil),        // 1: pb.QuestionOptions
	(*SportItem)(nil),              // 2: pb.SportItem
	(*CreateSportItemRequest)(nil), // 3: pb.CreateSportItemRequest
	(*GetSportItemRequest)(nil),    // 4: pb.GetSportItemRequest
	(*UpdateSportItemRequest)(nil), // 5: pb.UpdateSportItemRequest
	(*DeleteSportItemRequest)(nil), // 6: pb.DeleteSportItemRequest
	(*GetSportItemsRequest)(nil),   // 7: pb.GetSportItemsRequest
	(*GetSportItemsReply)(nil),     // 8: pb.GetSportItemsReply
}
var file_template_config_proto_depIdxs = []int32{
	1, // 0: pb.SportItem.question_options:type_name -> pb.QuestionOptions
	2, // 1: pb.CreateSportItemRequest.data:type_name -> pb.SportItem
	2, // 2: pb.UpdateSportItemRequest.data:type_name -> pb.SportItem
	2, // 3: pb.GetSportItemsReply.items:type_name -> pb.SportItem
	3, // 4: pb.SportItemManager.CreateSportItem:input_type -> pb.CreateSportItemRequest
	4, // 5: pb.SportItemManager.GetSportItem:input_type -> pb.GetSportItemRequest
	5, // 6: pb.SportItemManager.UpdateSportItem:input_type -> pb.UpdateSportItemRequest
	6, // 7: pb.SportItemManager.DeleteSportItem:input_type -> pb.DeleteSportItemRequest
	7, // 8: pb.SportItemManager.GetSportItems:input_type -> pb.GetSportItemsRequest
	2, // 9: pb.SportItemManager.CreateSportItem:output_type -> pb.SportItem
	2, // 10: pb.SportItemManager.GetSportItem:output_type -> pb.SportItem
	2, // 11: pb.SportItemManager.UpdateSportItem:output_type -> pb.SportItem
	2, // 12: pb.SportItemManager.DeleteSportItem:output_type -> pb.SportItem
	8, // 13: pb.SportItemManager.GetSportItems:output_type -> pb.GetSportItemsReply
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_template_config_proto_init() }
func file_template_config_proto_init() {
	if File_template_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_template_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Evaluate); i {
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
		file_template_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestionOptions); i {
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
		file_template_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SportItem); i {
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
		file_template_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSportItemRequest); i {
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
		file_template_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSportItemRequest); i {
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
		file_template_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSportItemRequest); i {
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
		file_template_config_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSportItemRequest); i {
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
		file_template_config_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSportItemsRequest); i {
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
		file_template_config_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSportItemsReply); i {
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
			RawDescriptor: file_template_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_template_config_proto_goTypes,
		DependencyIndexes: file_template_config_proto_depIdxs,
		MessageInfos:      file_template_config_proto_msgTypes,
	}.Build()
	File_template_config_proto = out.File
	file_template_config_proto_rawDesc = nil
	file_template_config_proto_goTypes = nil
	file_template_config_proto_depIdxs = nil
}
