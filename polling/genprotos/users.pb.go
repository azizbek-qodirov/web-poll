// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.1
// source: protos/users.proto

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

type GenderType int32

const (
	GenderType_MALE   GenderType = 0
	GenderType_FEMALE GenderType = 1
)

// Enum value maps for GenderType.
var (
	GenderType_name = map[int32]string{
		0: "MALE",
		1: "FEMALE",
	}
	GenderType_value = map[string]int32{
		"MALE":   0,
		"FEMALE": 1,
	}
)

func (x GenderType) Enum() *GenderType {
	p := new(GenderType)
	*p = x
	return p
}

func (x GenderType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GenderType) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_users_proto_enumTypes[0].Descriptor()
}

func (GenderType) Type() protoreflect.EnumType {
	return &file_protos_users_proto_enumTypes[0]
}

func (x GenderType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GenderType.Descriptor instead.
func (GenderType) EnumDescriptor() ([]byte, []int) {
	return file_protos_users_proto_rawDescGZIP(), []int{0}
}

type LevelType int32

const (
	LevelType_JUNIOR LevelType = 0
	LevelType_MIDDLE LevelType = 1
	LevelType_SENIOR LevelType = 2
)

// Enum value maps for LevelType.
var (
	LevelType_name = map[int32]string{
		0: "JUNIOR",
		1: "MIDDLE",
		2: "SENIOR",
	}
	LevelType_value = map[string]int32{
		"JUNIOR": 0,
		"MIDDLE": 1,
		"SENIOR": 2,
	}
)

func (x LevelType) Enum() *LevelType {
	p := new(LevelType)
	*p = x
	return p
}

func (x LevelType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LevelType) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_users_proto_enumTypes[1].Descriptor()
}

func (LevelType) Type() protoreflect.EnumType {
	return &file_protos_users_proto_enumTypes[1]
}

func (x LevelType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LevelType.Descriptor instead.
func (LevelType) EnumDescriptor() ([]byte, []int) {
	return file_protos_users_proto_rawDescGZIP(), []int{1}
}

type RegisterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name              string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Surname           string     `protobuf:"bytes,2,opt,name=surname,proto3" json:"surname,omitempty"`
	Gender            GenderType `protobuf:"varint,3,opt,name=gender,proto3,enum=polling.GenderType" json:"gender,omitempty"`
	Email             string     `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password          string     `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	PhoneNumber       string     `protobuf:"bytes,6,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	WorkingExperience int32      `protobuf:"varint,7,opt,name=working_experience,json=workingExperience,proto3" json:"working_experience,omitempty"`
	Level             LevelType  `protobuf:"varint,8,opt,name=level,proto3,enum=polling.LevelType" json:"level,omitempty"`
}

func (x *RegisterReq) Reset() {
	*x = RegisterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_users_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReq) ProtoMessage() {}

func (x *RegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_users_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReq.ProtoReflect.Descriptor instead.
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return file_protos_users_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterReq) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *RegisterReq) GetGender() GenderType {
	if x != nil {
		return x.Gender
	}
	return GenderType_MALE
}

func (x *RegisterReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterReq) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *RegisterReq) GetWorkingExperience() int32 {
	if x != nil {
		return x.WorkingExperience
	}
	return 0
}

func (x *RegisterReq) GetLevel() LevelType {
	if x != nil {
		return x.Level
	}
	return LevelType_JUNIOR
}

type RegisterReqForSwagger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name              string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Surname           string `protobuf:"bytes,2,opt,name=surname,proto3" json:"surname,omitempty"`
	Gender            string `protobuf:"bytes,3,opt,name=gender,proto3" json:"gender,omitempty"`
	Email             string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password          string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	PhoneNumber       string `protobuf:"bytes,6,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	WorkingExperience int32  `protobuf:"varint,7,opt,name=working_experience,json=workingExperience,proto3" json:"working_experience,omitempty"`
	Level             string `protobuf:"bytes,8,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *RegisterReqForSwagger) Reset() {
	*x = RegisterReqForSwagger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_users_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReqForSwagger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReqForSwagger) ProtoMessage() {}

func (x *RegisterReqForSwagger) ProtoReflect() protoreflect.Message {
	mi := &file_protos_users_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReqForSwagger.ProtoReflect.Descriptor instead.
func (*RegisterReqForSwagger) Descriptor() ([]byte, []int) {
	return file_protos_users_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterReqForSwagger) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterReqForSwagger) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *RegisterReqForSwagger) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *RegisterReqForSwagger) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterReqForSwagger) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterReqForSwagger) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *RegisterReqForSwagger) GetWorkingExperience() int32 {
	if x != nil {
		return x.WorkingExperience
	}
	return 0
}

func (x *RegisterReqForSwagger) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

type ConfirmUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"` // User's email to confirm
}

func (x *ConfirmUserReq) Reset() {
	*x = ConfirmUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_users_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmUserReq) ProtoMessage() {}

func (x *ConfirmUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_users_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmUserReq.ProtoReflect.Descriptor instead.
func (*ConfirmUserReq) Descriptor() ([]byte, []int) {
	return file_protos_users_proto_rawDescGZIP(), []int{2}
}

func (x *ConfirmUserReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type GetProfileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"` // User's email to retrieve the profile
}

func (x *GetProfileReq) Reset() {
	*x = GetProfileReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_users_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileReq) ProtoMessage() {}

func (x *GetProfileReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_users_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileReq.ProtoReflect.Descriptor instead.
func (*GetProfileReq) Descriptor() ([]byte, []int) {
	return file_protos_users_proto_rawDescGZIP(), []int{3}
}

func (x *GetProfileReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type GetProfileResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                       // User's unique identifier
	Email       string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`                                 // User's email address
	Password    string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`                           // User's password (consider if you want to expose this)
	Role        string `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`                                   // User's role
	IsConfirmed bool   `protobuf:"varint,5,opt,name=is_confirmed,json=isConfirmed,proto3" json:"is_confirmed,omitempty"` // Indicates if the user is confirmed
}

func (x *GetProfileResp) Reset() {
	*x = GetProfileResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_users_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfileResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileResp) ProtoMessage() {}

func (x *GetProfileResp) ProtoReflect() protoreflect.Message {
	mi := &file_protos_users_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileResp.ProtoReflect.Descriptor instead.
func (*GetProfileResp) Descriptor() ([]byte, []int) {
	return file_protos_users_proto_rawDescGZIP(), []int{4}
}

func (x *GetProfileResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetProfileResp) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetProfileResp) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *GetProfileResp) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *GetProfileResp) GetIsConfirmed() bool {
	if x != nil {
		return x.IsConfirmed
	}
	return false
}

type UpdatePasswordReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email       string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`                                // User's email
	NewPassword string `protobuf:"bytes,2,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"` // New password to set
}

func (x *UpdatePasswordReq) Reset() {
	*x = UpdatePasswordReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_users_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePasswordReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePasswordReq) ProtoMessage() {}

func (x *UpdatePasswordReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_users_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePasswordReq.ProtoReflect.Descriptor instead.
func (*UpdatePasswordReq) Descriptor() ([]byte, []int) {
	return file_protos_users_proto_rawDescGZIP(), []int{5}
}

func (x *UpdatePasswordReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdatePasswordReq) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

type IsEmailExistsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"` // User's email to check
}

func (x *IsEmailExistsReq) Reset() {
	*x = IsEmailExistsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_users_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsEmailExistsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsEmailExistsReq) ProtoMessage() {}

func (x *IsEmailExistsReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_users_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsEmailExistsReq.ProtoReflect.Descriptor instead.
func (*IsEmailExistsReq) Descriptor() ([]byte, []int) {
	return file_protos_users_proto_rawDescGZIP(), []int{6}
}

func (x *IsEmailExistsReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type IsEmailExistsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exists bool `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"` // Indicates if the email exists
}

func (x *IsEmailExistsResp) Reset() {
	*x = IsEmailExistsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_users_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsEmailExistsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsEmailExistsResp) ProtoMessage() {}

func (x *IsEmailExistsResp) ProtoReflect() protoreflect.Message {
	mi := &file_protos_users_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsEmailExistsResp.ProtoReflect.Descriptor instead.
func (*IsEmailExistsResp) Descriptor() ([]byte, []int) {
	return file_protos_users_proto_rawDescGZIP(), []int{7}
}

func (x *IsEmailExistsResp) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

type GetProfileByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // User's unique identifier
}

func (x *GetProfileByIdReq) Reset() {
	*x = GetProfileByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_users_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfileByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileByIdReq) ProtoMessage() {}

func (x *GetProfileByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_users_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileByIdReq.ProtoReflect.Descriptor instead.
func (*GetProfileByIdReq) Descriptor() ([]byte, []int) {
	return file_protos_users_proto_rawDescGZIP(), []int{8}
}

func (x *GetProfileByIdReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetProfileByIdResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`       // User's unique identifier
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"` // User's email address
	Role  string `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`   // User's role
}

func (x *GetProfileByIdResp) Reset() {
	*x = GetProfileByIdResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_users_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfileByIdResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileByIdResp) ProtoMessage() {}

func (x *GetProfileByIdResp) ProtoReflect() protoreflect.Message {
	mi := &file_protos_users_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileByIdResp.ProtoReflect.Descriptor instead.
func (*GetProfileByIdResp) Descriptor() ([]byte, []int) {
	return file_protos_users_proto_rawDescGZIP(), []int{9}
}

func (x *GetProfileByIdResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetProfileByIdResp) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetProfileByIdResp) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

var File_protos_users_proto protoreflect.FileDescriptor

var file_protos_users_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x1a, 0x11, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x6f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x96, 0x02, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b,
	0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13,
	0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x2d, 0x0a, 0x12, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x77, 0x6f,
	0x72, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12,
	0x28, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12,
	0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0xf7, 0x01, 0x0a, 0x15, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x46, 0x6f, 0x72, 0x53, 0x77, 0x61, 0x67,
	0x67, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2d,
	0x0a, 0x12, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x77, 0x6f, 0x72, 0x6b,
	0x69, 0x6e, 0x67, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x22, 0x26, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x25, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x22, 0x89, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x69,
	0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x22, 0x4c,
	0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x77,
	0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x28, 0x0a, 0x10,
	0x49, 0x73, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x2b, 0x0a, 0x11, 0x49, 0x73, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x65,
	0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x78, 0x69,
	0x73, 0x74, 0x73, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x2a, 0x22, 0x0a, 0x0a, 0x47, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x2a, 0x2f, 0x0a, 0x09,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x4a, 0x55, 0x4e,
	0x49, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x49, 0x44, 0x44, 0x4c, 0x45, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x45, 0x4e, 0x49, 0x4f, 0x52, 0x10, 0x02, 0x32, 0xfa, 0x02,
	0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a,
	0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x70, 0x6f, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a,
	0x0d, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x35,
	0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x2e,
	0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x3a, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x12, 0x16, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x3b, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x1a, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x1a,
	0x0d, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x46,
	0x0a, 0x0d, 0x49, 0x73, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12,
	0x19, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x49, 0x73, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x70, 0x6f, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x49, 0x73, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x45, 0x78, 0x69, 0x73,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x42, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49,
	0x44, 0x12, 0x1a, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e,
	0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0c, 0x5a, 0x0a, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_users_proto_rawDescOnce sync.Once
	file_protos_users_proto_rawDescData = file_protos_users_proto_rawDesc
)

func file_protos_users_proto_rawDescGZIP() []byte {
	file_protos_users_proto_rawDescOnce.Do(func() {
		file_protos_users_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_users_proto_rawDescData)
	})
	return file_protos_users_proto_rawDescData
}

var file_protos_users_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_protos_users_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_protos_users_proto_goTypes = []any{
	(GenderType)(0),               // 0: polling.GenderType
	(LevelType)(0),                // 1: polling.LevelType
	(*RegisterReq)(nil),           // 2: polling.RegisterReq
	(*RegisterReqForSwagger)(nil), // 3: polling.RegisterReqForSwagger
	(*ConfirmUserReq)(nil),        // 4: polling.ConfirmUserReq
	(*GetProfileReq)(nil),         // 5: polling.GetProfileReq
	(*GetProfileResp)(nil),        // 6: polling.GetProfileResp
	(*UpdatePasswordReq)(nil),     // 7: polling.UpdatePasswordReq
	(*IsEmailExistsReq)(nil),      // 8: polling.IsEmailExistsReq
	(*IsEmailExistsResp)(nil),     // 9: polling.IsEmailExistsResp
	(*GetProfileByIdReq)(nil),     // 10: polling.GetProfileByIdReq
	(*GetProfileByIdResp)(nil),    // 11: polling.GetProfileByIdResp
	(*Void)(nil),                  // 12: polling.Void
}
var file_protos_users_proto_depIdxs = []int32{
	0,  // 0: polling.RegisterReq.gender:type_name -> polling.GenderType
	1,  // 1: polling.RegisterReq.level:type_name -> polling.LevelType
	2,  // 2: polling.UserService.Register:input_type -> polling.RegisterReq
	4,  // 3: polling.UserService.ConfirmUser:input_type -> polling.ConfirmUserReq
	5,  // 4: polling.UserService.Profile:input_type -> polling.GetProfileReq
	7,  // 5: polling.UserService.UpdatePassword:input_type -> polling.UpdatePasswordReq
	8,  // 6: polling.UserService.IsEmailExists:input_type -> polling.IsEmailExistsReq
	10, // 7: polling.UserService.GetByID:input_type -> polling.GetProfileByIdReq
	12, // 8: polling.UserService.Register:output_type -> polling.Void
	12, // 9: polling.UserService.ConfirmUser:output_type -> polling.Void
	6,  // 10: polling.UserService.Profile:output_type -> polling.GetProfileResp
	12, // 11: polling.UserService.UpdatePassword:output_type -> polling.Void
	9,  // 12: polling.UserService.IsEmailExists:output_type -> polling.IsEmailExistsResp
	11, // 13: polling.UserService.GetByID:output_type -> polling.GetProfileByIdResp
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_protos_users_proto_init() }
func file_protos_users_proto_init() {
	if File_protos_users_proto != nil {
		return
	}
	file_protos_void_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_users_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RegisterReq); i {
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
		file_protos_users_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RegisterReqForSwagger); i {
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
		file_protos_users_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ConfirmUserReq); i {
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
		file_protos_users_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetProfileReq); i {
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
		file_protos_users_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetProfileResp); i {
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
		file_protos_users_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UpdatePasswordReq); i {
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
		file_protos_users_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*IsEmailExistsReq); i {
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
		file_protos_users_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*IsEmailExistsResp); i {
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
		file_protos_users_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*GetProfileByIdReq); i {
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
		file_protos_users_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*GetProfileByIdResp); i {
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
			RawDescriptor: file_protos_users_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_users_proto_goTypes,
		DependencyIndexes: file_protos_users_proto_depIdxs,
		EnumInfos:         file_protos_users_proto_enumTypes,
		MessageInfos:      file_protos_users_proto_msgTypes,
	}.Build()
	File_protos_users_proto = out.File
	file_protos_users_proto_rawDesc = nil
	file_protos_users_proto_goTypes = nil
	file_protos_users_proto_depIdxs = nil
}
