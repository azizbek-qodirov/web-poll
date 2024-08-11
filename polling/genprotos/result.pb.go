// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.1
// source: protos/result.proto

package genprotos

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

type CreateResultReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PollId string `protobuf:"bytes,2,opt,name=poll_id,json=pollId,proto3" json:"poll_id,omitempty"`
}

func (x *CreateResultReq) Reset() {
	*x = CreateResultReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_result_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResultReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResultReq) ProtoMessage() {}

func (x *CreateResultReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_result_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResultReq.ProtoReflect.Descriptor instead.
func (*CreateResultReq) Descriptor() ([]byte, []int) {
	return file_protos_result_proto_rawDescGZIP(), []int{0}
}

func (x *CreateResultReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateResultReq) GetPollId() string {
	if x != nil {
		return x.PollId
	}
	return ""
}

type CreateResultRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResultId string `protobuf:"bytes,1,opt,name=result_id,json=resultId,proto3" json:"result_id,omitempty"`
}

func (x *CreateResultRes) Reset() {
	*x = CreateResultRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_result_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResultRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResultRes) ProtoMessage() {}

func (x *CreateResultRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_result_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResultRes.ProtoReflect.Descriptor instead.
func (*CreateResultRes) Descriptor() ([]byte, []int) {
	return file_protos_result_proto_rawDescGZIP(), []int{1}
}

func (x *CreateResultRes) GetResultId() string {
	if x != nil {
		return x.ResultId
	}
	return ""
}

type SavePollAnswerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResultId   string `protobuf:"bytes,1,opt,name=result_id,json=resultId,proto3" json:"result_id,omitempty"`
	QuestionId string `protobuf:"bytes,2,opt,name=question_id,json=questionId,proto3" json:"question_id,omitempty"`
	Answer     int32  `protobuf:"varint,3,opt,name=answer,proto3" json:"answer,omitempty"`
}

func (x *SavePollAnswerReq) Reset() {
	*x = SavePollAnswerReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_result_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SavePollAnswerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavePollAnswerReq) ProtoMessage() {}

func (x *SavePollAnswerReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_result_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavePollAnswerReq.ProtoReflect.Descriptor instead.
func (*SavePollAnswerReq) Descriptor() ([]byte, []int) {
	return file_protos_result_proto_rawDescGZIP(), []int{2}
}

func (x *SavePollAnswerReq) GetResultId() string {
	if x != nil {
		return x.ResultId
	}
	return ""
}

func (x *SavePollAnswerReq) GetQuestionId() string {
	if x != nil {
		return x.QuestionId
	}
	return ""
}

func (x *SavePollAnswerReq) GetAnswer() int32 {
	if x != nil {
		return x.Answer
	}
	return 0
}

type GetPoll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PollId string `protobuf:"bytes,1,opt,name=poll_id,json=pollId,proto3" json:"poll_id,omitempty"`
}

func (x *GetPoll) Reset() {
	*x = GetPoll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_result_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPoll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPoll) ProtoMessage() {}

func (x *GetPoll) ProtoReflect() protoreflect.Message {
	mi := &file_protos_result_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPoll.ProtoReflect.Descriptor instead.
func (*GetPoll) Descriptor() ([]byte, []int) {
	return file_protos_result_proto_rawDescGZIP(), []int{3}
}

func (x *GetPoll) GetPollId() string {
	if x != nil {
		return x.PollId
	}
	return ""
}

type IncomingResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PollId  string            `protobuf:"bytes,1,opt,name=poll_id,json=pollId,proto3" json:"poll_id,omitempty"`
	UserId  string            `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Answers []*IncomingAnswer `protobuf:"bytes,3,rep,name=answers,proto3" json:"answers,omitempty"`
}

func (x *IncomingResult) Reset() {
	*x = IncomingResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_result_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncomingResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncomingResult) ProtoMessage() {}

func (x *IncomingResult) ProtoReflect() protoreflect.Message {
	mi := &file_protos_result_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncomingResult.ProtoReflect.Descriptor instead.
func (*IncomingResult) Descriptor() ([]byte, []int) {
	return file_protos_result_proto_rawDescGZIP(), []int{4}
}

func (x *IncomingResult) GetPollId() string {
	if x != nil {
		return x.PollId
	}
	return ""
}

func (x *IncomingResult) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *IncomingResult) GetAnswers() []*IncomingAnswer {
	if x != nil {
		return x.Answers
	}
	return nil
}

type IncomingAnswer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuestionNum int32 `protobuf:"varint,1,opt,name=question_num,json=questionNum,proto3" json:"question_num,omitempty"`
	AnswerPoint int32 `protobuf:"varint,2,opt,name=answer_point,json=answerPoint,proto3" json:"answer_point,omitempty"`
}

func (x *IncomingAnswer) Reset() {
	*x = IncomingAnswer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_result_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncomingAnswer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncomingAnswer) ProtoMessage() {}

func (x *IncomingAnswer) ProtoReflect() protoreflect.Message {
	mi := &file_protos_result_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncomingAnswer.ProtoReflect.Descriptor instead.
func (*IncomingAnswer) Descriptor() ([]byte, []int) {
	return file_protos_result_proto_rawDescGZIP(), []int{5}
}

func (x *IncomingAnswer) GetQuestionNum() int32 {
	if x != nil {
		return x.QuestionNum
	}
	return 0
}

func (x *IncomingAnswer) GetAnswerPoint() int32 {
	if x != nil {
		return x.AnswerPoint
	}
	return 0
}

type ExcelResultsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*ResultRes `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *ExcelResultsRes) Reset() {
	*x = ExcelResultsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_result_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExcelResultsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExcelResultsRes) ProtoMessage() {}

func (x *ExcelResultsRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_result_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExcelResultsRes.ProtoReflect.Descriptor instead.
func (*ExcelResultsRes) Descriptor() ([]byte, []int) {
	return file_protos_result_proto_rawDescGZIP(), []int{6}
}

func (x *ExcelResultsRes) GetResults() []*ResultRes {
	if x != nil {
		return x.Results
	}
	return nil
}

type ResultRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User    *UserExcelResp    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	PollNum int32             `protobuf:"varint,2,opt,name=poll_num,json=pollNum,proto3" json:"poll_num,omitempty"`
	Answers []*IncomingAnswer `protobuf:"bytes,3,rep,name=answers,proto3" json:"answers,omitempty"`
}

func (x *ResultRes) Reset() {
	*x = ResultRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_result_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultRes) ProtoMessage() {}

func (x *ResultRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_result_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultRes.ProtoReflect.Descriptor instead.
func (*ResultRes) Descriptor() ([]byte, []int) {
	return file_protos_result_proto_rawDescGZIP(), []int{7}
}

func (x *ResultRes) GetUser() *UserExcelResp {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *ResultRes) GetPollNum() int32 {
	if x != nil {
		return x.PollNum
	}
	return 0
}

func (x *ResultRes) GetAnswers() []*IncomingAnswer {
	if x != nil {
		return x.Answers
	}
	return nil
}

var File_protos_result_proto protoreflect.FileDescriptor

var file_protos_result_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x1a, 0x12,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x6f, 0x69, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x6c, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x0f, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x64, 0x22, 0x69, 0x0a, 0x11, 0x53, 0x61,
	0x76, 0x65, 0x50, 0x6f, 0x6c, 0x6c, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12,
	0x1b, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x22, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x6c,
	0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x6c, 0x49, 0x64, 0x22, 0x75, 0x0a, 0x0e, 0x49, 0x6e, 0x63,
	0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70,
	0x6f, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f,
	0x6c, 0x6c, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x31, 0x0a,
	0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e,
	0x67, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73,
	0x22, 0x56, 0x0a, 0x0e, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x41, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e,
	0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x5f,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x61, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x3f, 0x0a, 0x0f, 0x45, 0x78, 0x63, 0x65,
	0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x07, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70,
	0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73,
	0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x85, 0x01, 0x0a, 0x09, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6f, 0x6c, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x6f, 0x6c, 0x6c, 0x4e, 0x75, 0x6d, 0x12, 0x31,
	0x0a, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x69,
	0x6e, 0x67, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x73, 0x32, 0xce, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e,
	0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x12, 0x3b, 0x0a, 0x0e, 0x53, 0x61, 0x76, 0x65, 0x50,
	0x6f, 0x6c, 0x6c, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x70, 0x6f, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x50, 0x6f, 0x6c, 0x6c, 0x41, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e,
	0x56, 0x6f, 0x69, 0x64, 0x12, 0x3c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x49, 0x6e, 0x45, 0x78, 0x63, 0x65, 0x6c, 0x12, 0x0d, 0x2e, 0x70, 0x6f, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x18, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2e, 0x45, 0x78, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x42, 0x0c, 0x5a, 0x0a, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_result_proto_rawDescOnce sync.Once
	file_protos_result_proto_rawDescData = file_protos_result_proto_rawDesc
)

func file_protos_result_proto_rawDescGZIP() []byte {
	file_protos_result_proto_rawDescOnce.Do(func() {
		file_protos_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_result_proto_rawDescData)
	})
	return file_protos_result_proto_rawDescData
}

var file_protos_result_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protos_result_proto_goTypes = []any{
	(*CreateResultReq)(nil),   // 0: polling.CreateResultReq
	(*CreateResultRes)(nil),   // 1: polling.CreateResultRes
	(*SavePollAnswerReq)(nil), // 2: polling.SavePollAnswerReq
	(*GetPoll)(nil),           // 3: polling.GetPoll
	(*IncomingResult)(nil),    // 4: polling.IncomingResult
	(*IncomingAnswer)(nil),    // 5: polling.IncomingAnswer
	(*ExcelResultsRes)(nil),   // 6: polling.ExcelResultsRes
	(*ResultRes)(nil),         // 7: polling.ResultRes
	(*UserExcelResp)(nil),     // 8: polling.UserExcelResp
	(*Void)(nil),              // 9: polling.Void
}
var file_protos_result_proto_depIdxs = []int32{
	5, // 0: polling.IncomingResult.answers:type_name -> polling.IncomingAnswer
	7, // 1: polling.ExcelResultsRes.results:type_name -> polling.ResultRes
	8, // 2: polling.ResultRes.user:type_name -> polling.UserExcelResp
	5, // 3: polling.ResultRes.answers:type_name -> polling.IncomingAnswer
	0, // 4: polling.ResultService.CreateResult:input_type -> polling.CreateResultReq
	2, // 5: polling.ResultService.SavePollAnswer:input_type -> polling.SavePollAnswerReq
	9, // 6: polling.ResultService.GetResultsInExcel:input_type -> polling.Void
	1, // 7: polling.ResultService.CreateResult:output_type -> polling.CreateResultRes
	9, // 8: polling.ResultService.SavePollAnswer:output_type -> polling.Void
	6, // 9: polling.ResultService.GetResultsInExcel:output_type -> polling.ExcelResultsRes
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_protos_result_proto_init() }
func file_protos_result_proto_init() {
	if File_protos_result_proto != nil {
		return
	}
	file_protos_users_proto_init()
	file_protos_void_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_result_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateResultReq); i {
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
		file_protos_result_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateResultRes); i {
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
		file_protos_result_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SavePollAnswerReq); i {
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
		file_protos_result_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetPoll); i {
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
		file_protos_result_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*IncomingResult); i {
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
		file_protos_result_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*IncomingAnswer); i {
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
		file_protos_result_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ExcelResultsRes); i {
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
		file_protos_result_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*ResultRes); i {
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
			RawDescriptor: file_protos_result_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_result_proto_goTypes,
		DependencyIndexes: file_protos_result_proto_depIdxs,
		MessageInfos:      file_protos_result_proto_msgTypes,
	}.Build()
	File_protos_result_proto = out.File
	file_protos_result_proto_rawDesc = nil
	file_protos_result_proto_goTypes = nil
	file_protos_result_proto_depIdxs = nil
}
