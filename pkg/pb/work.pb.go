// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: work.proto

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

type FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" bson:"name"`
	Url  string `protobuf:"bytes,2,opt,name=url,proto3" json:"url" bson:"url"`
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_work_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_work_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_work_proto_rawDescGZIP(), []int{0}
}

func (x *FileInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileInfo) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type NoteInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid  string `protobuf:"bytes,1,opt,name=userid,proto3" json:"userid" bson:"userid"`
	Score   string `protobuf:"bytes,2,opt,name=score,proto3" json:"score" bson:"score"`
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content" bson:"content"`
	Status  string `protobuf:"bytes,4,opt,name=status,proto3" json:"status" bson:"status"`
	Value   string `protobuf:"bytes,5,opt,name=value,proto3" json:"value" bson:"value"`
}

func (x *NoteInfo) Reset() {
	*x = NoteInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_work_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoteInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteInfo) ProtoMessage() {}

func (x *NoteInfo) ProtoReflect() protoreflect.Message {
	mi := &file_work_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteInfo.ProtoReflect.Descriptor instead.
func (*NoteInfo) Descriptor() ([]byte, []int) {
	return file_work_proto_rawDescGZIP(), []int{1}
}

func (x *NoteInfo) GetUserid() string {
	if x != nil {
		return x.Userid
	}
	return ""
}

func (x *NoteInfo) GetScore() string {
	if x != nil {
		return x.Score
	}
	return ""
}

func (x *NoteInfo) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *NoteInfo) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *NoteInfo) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Work struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
	Userid        string      `protobuf:"bytes,2,opt,name=userid,proto3" json:"userid" bson:"userid"`
	Course        string      `protobuf:"bytes,3,opt,name=course,proto3" json:"course" bson:"course"`
	CourseTitle   string      `protobuf:"bytes,4,opt,name=course_title,json=courseTitle,proto3" json:"course_title" bson:"course_title"`
	Files         []*FileInfo `protobuf:"bytes,5,rep,name=files,proto3" json:"files" bson:"files"`
	Content       string      `protobuf:"bytes,6,opt,name=content,proto3" json:"content" bson:"content"`
	Note          string      `protobuf:"bytes,7,opt,name=note,proto3" json:"note" bson:"note"`
	Score         string      `protobuf:"bytes,8,opt,name=score,proto3" json:"score" bson:"score"`
	Type          int32       `protobuf:"varint,9,opt,name=type,proto3" json:"type" bson:"type"`
	Create        int64       `protobuf:"varint,10,opt,name=create,proto3" json:"create" bson:"create"`
	Update        int64       `protobuf:"varint,11,opt,name=update,proto3" json:"update" bson:"update"`
	Username      string      `protobuf:"bytes,12,opt,name=username,proto3" json:"username" bson:"username"`
	CourseCreater string      `protobuf:"bytes,13,opt,name=course_creater,json=courseCreater,proto3" json:"course_creater" bson:"course_creater"`
	Teacherid     string      `protobuf:"bytes,14,opt,name=teacherid,proto3" json:"teacherid" bson:"teacherid"`
	Noter         []string    `protobuf:"bytes,15,rep,name=noter,proto3" json:"noter" bson:"noter"`
	NoteInfo      []*NoteInfo `protobuf:"bytes,16,rep,name=note_info,json=noteInfo,proto3" json:"note_info" bson:"note_info"`
}

func (x *Work) Reset() {
	*x = Work{}
	if protoimpl.UnsafeEnabled {
		mi := &file_work_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Work) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Work) ProtoMessage() {}

func (x *Work) ProtoReflect() protoreflect.Message {
	mi := &file_work_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Work.ProtoReflect.Descriptor instead.
func (*Work) Descriptor() ([]byte, []int) {
	return file_work_proto_rawDescGZIP(), []int{2}
}

func (x *Work) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Work) GetUserid() string {
	if x != nil {
		return x.Userid
	}
	return ""
}

func (x *Work) GetCourse() string {
	if x != nil {
		return x.Course
	}
	return ""
}

func (x *Work) GetCourseTitle() string {
	if x != nil {
		return x.CourseTitle
	}
	return ""
}

func (x *Work) GetFiles() []*FileInfo {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *Work) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Work) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *Work) GetScore() string {
	if x != nil {
		return x.Score
	}
	return ""
}

func (x *Work) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Work) GetCreate() int64 {
	if x != nil {
		return x.Create
	}
	return 0
}

func (x *Work) GetUpdate() int64 {
	if x != nil {
		return x.Update
	}
	return 0
}

func (x *Work) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Work) GetCourseCreater() string {
	if x != nil {
		return x.CourseCreater
	}
	return ""
}

func (x *Work) GetTeacherid() string {
	if x != nil {
		return x.Teacherid
	}
	return ""
}

func (x *Work) GetNoter() []string {
	if x != nil {
		return x.Noter
	}
	return nil
}

func (x *Work) GetNoteInfo() []*NoteInfo {
	if x != nil {
		return x.NoteInfo
	}
	return nil
}

type CreateWorkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Work `protobuf:"bytes,1,opt,name=data,proto3" json:"data" bson:"data"`
}

func (x *CreateWorkRequest) Reset() {
	*x = CreateWorkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_work_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWorkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWorkRequest) ProtoMessage() {}

func (x *CreateWorkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_work_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWorkRequest.ProtoReflect.Descriptor instead.
func (*CreateWorkRequest) Descriptor() ([]byte, []int) {
	return file_work_proto_rawDescGZIP(), []int{3}
}

func (x *CreateWorkRequest) GetData() *Work {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetWorkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
}

func (x *GetWorkRequest) Reset() {
	*x = GetWorkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_work_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWorkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkRequest) ProtoMessage() {}

func (x *GetWorkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_work_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkRequest.ProtoReflect.Descriptor instead.
func (*GetWorkRequest) Descriptor() ([]byte, []int) {
	return file_work_proto_rawDescGZIP(), []int{4}
}

func (x *GetWorkRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateWorkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
	Data *Work  `protobuf:"bytes,2,opt,name=data,proto3" json:"data" bson:"data"`
}

func (x *UpdateWorkRequest) Reset() {
	*x = UpdateWorkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_work_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateWorkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWorkRequest) ProtoMessage() {}

func (x *UpdateWorkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_work_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWorkRequest.ProtoReflect.Descriptor instead.
func (*UpdateWorkRequest) Descriptor() ([]byte, []int) {
	return file_work_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateWorkRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateWorkRequest) GetData() *Work {
	if x != nil {
		return x.Data
	}
	return nil
}

type DeleteWorkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
}

func (x *DeleteWorkRequest) Reset() {
	*x = DeleteWorkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_work_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWorkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWorkRequest) ProtoMessage() {}

func (x *DeleteWorkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_work_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWorkRequest.ProtoReflect.Descriptor instead.
func (*DeleteWorkRequest) Descriptor() ([]byte, []int) {
	return file_work_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteWorkRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetWorksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sort  string `protobuf:"bytes,1,opt,name=sort,proto3" json:"sort" bson:"sort"`
	Query string `protobuf:"bytes,2,opt,name=query,proto3" json:"query" bson:"query"`
	Skip  int64  `protobuf:"varint,3,opt,name=skip,proto3" json:"skip" bson:"skip"`
	Limit int64  `protobuf:"varint,4,opt,name=limit,proto3" json:"limit" bson:"limit"`
}

func (x *GetWorksRequest) Reset() {
	*x = GetWorksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_work_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWorksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorksRequest) ProtoMessage() {}

func (x *GetWorksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_work_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorksRequest.ProtoReflect.Descriptor instead.
func (*GetWorksRequest) Descriptor() ([]byte, []int) {
	return file_work_proto_rawDescGZIP(), []int{7}
}

func (x *GetWorksRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *GetWorksRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *GetWorksRequest) GetSkip() int64 {
	if x != nil {
		return x.Skip
	}
	return 0
}

func (x *GetWorksRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetWorksReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount int64   `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count" bson:"total_count"`
	Items      []*Work `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *GetWorksReply) Reset() {
	*x = GetWorksReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_work_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWorksReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorksReply) ProtoMessage() {}

func (x *GetWorksReply) ProtoReflect() protoreflect.Message {
	mi := &file_work_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorksReply.ProtoReflect.Descriptor instead.
func (*GetWorksReply) Descriptor() ([]byte, []int) {
	return file_work_proto_rawDescGZIP(), []int{8}
}

func (x *GetWorksReply) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *GetWorksReply) GetItems() []*Work {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_work_proto protoreflect.FileDescriptor

var file_work_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0x30, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x22, 0x80, 0x01, 0x0a, 0x08, 0x4e, 0x6f, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xb7, 0x03, 0x0a, 0x04, 0x57, 0x6f, 0x72, 0x6b, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x22, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x6f, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x69, 0x64, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x72, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e,
	0x6f, 0x74, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x09, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x74,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x6e, 0x6f, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x31, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x41, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f,
	0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x65, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73,
	0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6b, 0x69,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x6b, 0x69, 0x70, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x22, 0x50, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0x81, 0x02, 0x0a, 0x0b, 0x57, 0x6f, 0x72, 0x6b, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57,
	0x6f, 0x72, 0x6b, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57,
	0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e,
	0x57, 0x6f, 0x72, 0x6b, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72,
	0x6b, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x22,
	0x00, 0x12, 0x2f, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x12,
	0x15, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x57, 0x6f, 0x72, 0x6b,
	0x22, 0x00, 0x12, 0x2f, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b,
	0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x12,
	0x13, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x75, 0x73, 0x68, 0x6b, 0x2f, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_work_proto_rawDescOnce sync.Once
	file_work_proto_rawDescData = file_work_proto_rawDesc
)

func file_work_proto_rawDescGZIP() []byte {
	file_work_proto_rawDescOnce.Do(func() {
		file_work_proto_rawDescData = protoimpl.X.CompressGZIP(file_work_proto_rawDescData)
	})
	return file_work_proto_rawDescData
}

var file_work_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_work_proto_goTypes = []interface{}{
	(*FileInfo)(nil),          // 0: pb.FileInfo
	(*NoteInfo)(nil),          // 1: pb.NoteInfo
	(*Work)(nil),              // 2: pb.Work
	(*CreateWorkRequest)(nil), // 3: pb.CreateWorkRequest
	(*GetWorkRequest)(nil),    // 4: pb.GetWorkRequest
	(*UpdateWorkRequest)(nil), // 5: pb.UpdateWorkRequest
	(*DeleteWorkRequest)(nil), // 6: pb.DeleteWorkRequest
	(*GetWorksRequest)(nil),   // 7: pb.GetWorksRequest
	(*GetWorksReply)(nil),     // 8: pb.GetWorksReply
}
var file_work_proto_depIdxs = []int32{
	0,  // 0: pb.Work.files:type_name -> pb.FileInfo
	1,  // 1: pb.Work.note_info:type_name -> pb.NoteInfo
	2,  // 2: pb.CreateWorkRequest.data:type_name -> pb.Work
	2,  // 3: pb.UpdateWorkRequest.data:type_name -> pb.Work
	2,  // 4: pb.GetWorksReply.items:type_name -> pb.Work
	3,  // 5: pb.WorkManager.CreateWork:input_type -> pb.CreateWorkRequest
	4,  // 6: pb.WorkManager.GetWork:input_type -> pb.GetWorkRequest
	5,  // 7: pb.WorkManager.UpdateWork:input_type -> pb.UpdateWorkRequest
	6,  // 8: pb.WorkManager.DeleteWork:input_type -> pb.DeleteWorkRequest
	7,  // 9: pb.WorkManager.GetWorks:input_type -> pb.GetWorksRequest
	2,  // 10: pb.WorkManager.CreateWork:output_type -> pb.Work
	2,  // 11: pb.WorkManager.GetWork:output_type -> pb.Work
	2,  // 12: pb.WorkManager.UpdateWork:output_type -> pb.Work
	2,  // 13: pb.WorkManager.DeleteWork:output_type -> pb.Work
	8,  // 14: pb.WorkManager.GetWorks:output_type -> pb.GetWorksReply
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_work_proto_init() }
func file_work_proto_init() {
	if File_work_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_work_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfo); i {
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
		file_work_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoteInfo); i {
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
		file_work_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Work); i {
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
		file_work_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWorkRequest); i {
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
		file_work_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWorkRequest); i {
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
		file_work_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateWorkRequest); i {
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
		file_work_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWorkRequest); i {
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
		file_work_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWorksRequest); i {
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
		file_work_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWorksReply); i {
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
			RawDescriptor: file_work_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_work_proto_goTypes,
		DependencyIndexes: file_work_proto_depIdxs,
		MessageInfos:      file_work_proto_msgTypes,
	}.Build()
	File_work_proto = out.File
	file_work_proto_rawDesc = nil
	file_work_proto_goTypes = nil
	file_work_proto_depIdxs = nil
}