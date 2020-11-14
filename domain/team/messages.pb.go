// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: domain/team/messages.proto

package team

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Create Team
type CreateTeam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TeamName string `protobuf:"bytes,2,opt,name=team_name,json=teamName,proto3" json:"team_name,omitempty"`
}

func (x *CreateTeam) Reset() {
	*x = CreateTeam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_team_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTeam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTeam) ProtoMessage() {}

func (x *CreateTeam) ProtoReflect() protoreflect.Message {
	mi := &file_domain_team_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTeam.ProtoReflect.Descriptor instead.
func (*CreateTeam) Descriptor() ([]byte, []int) {
	return file_domain_team_messages_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTeam) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateTeam) GetTeamName() string {
	if x != nil {
		return x.TeamName
	}
	return ""
}

type UserCreatedTeam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId      string `protobuf:"bytes,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	OwnerUserId string `protobuf:"bytes,2,opt,name=owner_user_id,json=ownerUserId,proto3" json:"owner_user_id,omitempty"`
	TeamName    string `protobuf:"bytes,3,opt,name=team_name,json=teamName,proto3" json:"team_name,omitempty"`
}

func (x *UserCreatedTeam) Reset() {
	*x = UserCreatedTeam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_team_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCreatedTeam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCreatedTeam) ProtoMessage() {}

func (x *UserCreatedTeam) ProtoReflect() protoreflect.Message {
	mi := &file_domain_team_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCreatedTeam.ProtoReflect.Descriptor instead.
func (*UserCreatedTeam) Descriptor() ([]byte, []int) {
	return file_domain_team_messages_proto_rawDescGZIP(), []int{1}
}

func (x *UserCreatedTeam) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

func (x *UserCreatedTeam) GetOwnerUserId() string {
	if x != nil {
		return x.OwnerUserId
	}
	return ""
}

func (x *UserCreatedTeam) GetTeamName() string {
	if x != nil {
		return x.TeamName
	}
	return ""
}

// Change Team Owner
type ChangeTeamOwner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId         string `protobuf:"bytes,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	NewOwnerUserId string `protobuf:"bytes,2,opt,name=new_owner_user_id,json=newOwnerUserId,proto3" json:"new_owner_user_id,omitempty"`
}

func (x *ChangeTeamOwner) Reset() {
	*x = ChangeTeamOwner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_team_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeTeamOwner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeTeamOwner) ProtoMessage() {}

func (x *ChangeTeamOwner) ProtoReflect() protoreflect.Message {
	mi := &file_domain_team_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeTeamOwner.ProtoReflect.Descriptor instead.
func (*ChangeTeamOwner) Descriptor() ([]byte, []int) {
	return file_domain_team_messages_proto_rawDescGZIP(), []int{2}
}

func (x *ChangeTeamOwner) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

func (x *ChangeTeamOwner) GetNewOwnerUserId() string {
	if x != nil {
		return x.NewOwnerUserId
	}
	return ""
}

type TeamOwnerChanged struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId      string `protobuf:"bytes,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	OwnerUserId string `protobuf:"bytes,2,opt,name=owner_user_id,json=ownerUserId,proto3" json:"owner_user_id,omitempty"`
}

func (x *TeamOwnerChanged) Reset() {
	*x = TeamOwnerChanged{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_team_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamOwnerChanged) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamOwnerChanged) ProtoMessage() {}

func (x *TeamOwnerChanged) ProtoReflect() protoreflect.Message {
	mi := &file_domain_team_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamOwnerChanged.ProtoReflect.Descriptor instead.
func (*TeamOwnerChanged) Descriptor() ([]byte, []int) {
	return file_domain_team_messages_proto_rawDescGZIP(), []int{3}
}

func (x *TeamOwnerChanged) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

func (x *TeamOwnerChanged) GetOwnerUserId() string {
	if x != nil {
		return x.OwnerUserId
	}
	return ""
}

// Remove User From Team
type RemoveUserFromTeam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId string `protobuf:"bytes,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *RemoveUserFromTeam) Reset() {
	*x = RemoveUserFromTeam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_team_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveUserFromTeam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserFromTeam) ProtoMessage() {}

func (x *RemoveUserFromTeam) ProtoReflect() protoreflect.Message {
	mi := &file_domain_team_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserFromTeam.ProtoReflect.Descriptor instead.
func (*RemoveUserFromTeam) Descriptor() ([]byte, []int) {
	return file_domain_team_messages_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveUserFromTeam) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

func (x *RemoveUserFromTeam) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UserRemovedFromTeam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId string `protobuf:"bytes,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserRemovedFromTeam) Reset() {
	*x = UserRemovedFromTeam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_team_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRemovedFromTeam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRemovedFromTeam) ProtoMessage() {}

func (x *UserRemovedFromTeam) ProtoReflect() protoreflect.Message {
	mi := &file_domain_team_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRemovedFromTeam.ProtoReflect.Descriptor instead.
func (*UserRemovedFromTeam) Descriptor() ([]byte, []int) {
	return file_domain_team_messages_proto_rawDescGZIP(), []int{5}
}

func (x *UserRemovedFromTeam) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

func (x *UserRemovedFromTeam) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// Add User to Team
type AddUserToTeam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId string `protobuf:"bytes,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *AddUserToTeam) Reset() {
	*x = AddUserToTeam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_team_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserToTeam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserToTeam) ProtoMessage() {}

func (x *AddUserToTeam) ProtoReflect() protoreflect.Message {
	mi := &file_domain_team_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserToTeam.ProtoReflect.Descriptor instead.
func (*AddUserToTeam) Descriptor() ([]byte, []int) {
	return file_domain_team_messages_proto_rawDescGZIP(), []int{6}
}

func (x *AddUserToTeam) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

func (x *AddUserToTeam) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UserAddedToTeam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId string `protobuf:"bytes,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserAddedToTeam) Reset() {
	*x = UserAddedToTeam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_team_messages_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAddedToTeam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAddedToTeam) ProtoMessage() {}

func (x *UserAddedToTeam) ProtoReflect() protoreflect.Message {
	mi := &file_domain_team_messages_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAddedToTeam.ProtoReflect.Descriptor instead.
func (*UserAddedToTeam) Descriptor() ([]byte, []int) {
	return file_domain_team_messages_proto_rawDescGZIP(), []int{7}
}

func (x *UserAddedToTeam) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

func (x *UserAddedToTeam) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_domain_team_messages_proto protoreflect.FileDescriptor

var file_domain_team_messages_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x65,
	0x61, 0x6d, 0x22, 0x42, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x61,
	0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65,
	0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x6b, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d,
	0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x55, 0x0a, 0x0f, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x65, 0x61,
	0x6d, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12,
	0x29, 0x0a, 0x11, 0x6e, 0x65, 0x77, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x65, 0x77, 0x4f,
	0x77, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x10, 0x54, 0x65,
	0x61, 0x6d, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x46, 0x0a, 0x12, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x65, 0x61,
	0x6d, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65,
	0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x61,
	0x6d, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x0d,
	0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x17, 0x0a,
	0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x43, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x64, 0x65, 0x64, 0x54, 0x6f, 0x54, 0x65,
	0x61, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x65, 0x78, 0x72, 0x75, 0x64, 0x64, 0x2f, 0x6c, 0x62, 0x2d, 0x74,
	0x65, 0x61, 0x6d, 0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x74, 0x65, 0x61, 0x6d,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_domain_team_messages_proto_rawDescOnce sync.Once
	file_domain_team_messages_proto_rawDescData = file_domain_team_messages_proto_rawDesc
)

func file_domain_team_messages_proto_rawDescGZIP() []byte {
	file_domain_team_messages_proto_rawDescOnce.Do(func() {
		file_domain_team_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_domain_team_messages_proto_rawDescData)
	})
	return file_domain_team_messages_proto_rawDescData
}

var file_domain_team_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_domain_team_messages_proto_goTypes = []interface{}{
	(*CreateTeam)(nil),          // 0: team.CreateTeam
	(*UserCreatedTeam)(nil),     // 1: team.UserCreatedTeam
	(*ChangeTeamOwner)(nil),     // 2: team.ChangeTeamOwner
	(*TeamOwnerChanged)(nil),    // 3: team.TeamOwnerChanged
	(*RemoveUserFromTeam)(nil),  // 4: team.RemoveUserFromTeam
	(*UserRemovedFromTeam)(nil), // 5: team.UserRemovedFromTeam
	(*AddUserToTeam)(nil),       // 6: team.AddUserToTeam
	(*UserAddedToTeam)(nil),     // 7: team.UserAddedToTeam
}
var file_domain_team_messages_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_domain_team_messages_proto_init() }
func file_domain_team_messages_proto_init() {
	if File_domain_team_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_domain_team_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTeam); i {
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
		file_domain_team_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCreatedTeam); i {
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
		file_domain_team_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeTeamOwner); i {
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
		file_domain_team_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamOwnerChanged); i {
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
		file_domain_team_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveUserFromTeam); i {
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
		file_domain_team_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRemovedFromTeam); i {
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
		file_domain_team_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserToTeam); i {
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
		file_domain_team_messages_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAddedToTeam); i {
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
			RawDescriptor: file_domain_team_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_domain_team_messages_proto_goTypes,
		DependencyIndexes: file_domain_team_messages_proto_depIdxs,
		MessageInfos:      file_domain_team_messages_proto_msgTypes,
	}.Build()
	File_domain_team_messages_proto = out.File
	file_domain_team_messages_proto_rawDesc = nil
	file_domain_team_messages_proto_goTypes = nil
	file_domain_team_messages_proto_depIdxs = nil
}
