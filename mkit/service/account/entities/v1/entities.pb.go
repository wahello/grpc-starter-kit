// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.3
// source: mkit/service/account/entities/v1/entities.proto

package entitiesv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	types "github.com/infobloxopen/protoc-gen-gorm/types"
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

type Profile_GenderType int32

const (
	Profile_GENDER_TYPE_UNSPECIFIED Profile_GenderType = 0
	Profile_GENDER_TYPE_MALE        Profile_GenderType = 1
	Profile_GENDER_TYPE_FEMALE      Profile_GenderType = 2
)

// Enum value maps for Profile_GenderType.
var (
	Profile_GenderType_name = map[int32]string{
		0: "GENDER_TYPE_UNSPECIFIED",
		1: "GENDER_TYPE_MALE",
		2: "GENDER_TYPE_FEMALE",
	}
	Profile_GenderType_value = map[string]int32{
		"GENDER_TYPE_UNSPECIFIED": 0,
		"GENDER_TYPE_MALE":        1,
		"GENDER_TYPE_FEMALE":      2,
	}
)

func (x Profile_GenderType) Enum() *Profile_GenderType {
	p := new(Profile_GenderType)
	*p = x
	return p
}

func (x Profile_GenderType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Profile_GenderType) Descriptor() protoreflect.EnumDescriptor {
	return file_mkit_service_account_entities_v1_entities_proto_enumTypes[0].Descriptor()
}

func (Profile_GenderType) Type() protoreflect.EnumType {
	return &file_mkit_service_account_entities_v1_entities_proto_enumTypes[0]
}

func (x Profile_GenderType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Profile_GenderType.Descriptor instead.
func (Profile_GenderType) EnumDescriptor() ([]byte, []int) {
	return file_mkit_service_account_entities_v1_entities_proto_rawDescGZIP(), []int{1, 0}
}

// User Entity
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        *types.UUID           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // primary key
	CreatedAt *timestamp.Timestamp  `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamp.Timestamp  `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt *timestamp.Timestamp  `protobuf:"bytes,4,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	Username  *wrappers.StringValue `protobuf:"bytes,7,opt,name=username,proto3" json:"username,omitempty"`
	FirstName string                `protobuf:"bytes,8,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string                `protobuf:"bytes,9,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email     string                `protobuf:"bytes,10,opt,name=email,proto3" json:"email,omitempty"`
	Profile   *Profile              `protobuf:"bytes,11,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mkit_service_account_entities_v1_entities_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_mkit_service_account_entities_v1_entities_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_mkit_service_account_entities_v1_entities_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() *types.UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *User) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *User) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *User) GetDeletedAt() *timestamp.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *User) GetUsername() *wrappers.StringValue {
	if x != nil {
		return x.Username
	}
	return nil
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetProfile() *Profile {
	if x != nil {
		return x.Profile
	}
	return nil
}

// Profile Entity
type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// gorm.types.UUID id                      = 1  [(gorm.field).tag = {
	// type: "uuid", primary_key: true, default: "uuid_generate_v4()" }]; primary
	// key
	Id        *types.UUID          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // primary key
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	Tz        string               `protobuf:"bytes,5,opt,name=tz,proto3" json:"tz,omitempty"` // *time.Location?
	Avatar    string               `protobuf:"bytes,6,opt,name=avatar,proto3" json:"avatar,omitempty"`
	// FIXME: https://github.com/jinzhu/gorm/issues/1978
	Gender         Profile_GenderType    `protobuf:"varint,7,opt,name=gender,proto3,enum=mkit.service.account.entities.v1.Profile_GenderType" json:"gender,omitempty"` //  GenderType `gorm:"not null;type:ENUM('M', 'F')"`
	Birthday       *timestamp.Timestamp  `protobuf:"bytes,8,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Age            uint32                `protobuf:"varint,9,opt,name=age,proto3" json:"age,omitempty"` // synthetic field
	PreferredTheme *wrappers.StringValue `protobuf:"bytes,10,opt,name=preferred_theme,json=preferredTheme,proto3" json:"preferred_theme,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mkit_service_account_entities_v1_entities_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_mkit_service_account_entities_v1_entities_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_mkit_service_account_entities_v1_entities_proto_rawDescGZIP(), []int{1}
}

func (x *Profile) GetId() *types.UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Profile) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Profile) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Profile) GetDeletedAt() *timestamp.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *Profile) GetTz() string {
	if x != nil {
		return x.Tz
	}
	return ""
}

func (x *Profile) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *Profile) GetGender() Profile_GenderType {
	if x != nil {
		return x.Gender
	}
	return Profile_GENDER_TYPE_UNSPECIFIED
}

func (x *Profile) GetBirthday() *timestamp.Timestamp {
	if x != nil {
		return x.Birthday
	}
	return nil
}

func (x *Profile) GetAge() uint32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Profile) GetPreferredTheme() *wrappers.StringValue {
	if x != nil {
		return x.PreferredTheme
	}
	return nil
}

var File_mkit_service_account_entities_v1_entities_proto protoreflect.FileDescriptor

var file_mkit_service_account_entities_v1_entities_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x6d, 0x6b, 0x69, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x20, 0x6d, 0x6b, 0x69, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x72, 0x6d, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f,
	0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x72, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x04, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x34, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x72, 0x6d,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x42, 0x12, 0xba, 0xb9, 0x19,
	0x0e, 0x0a, 0x0c, 0x12, 0x04, 0x75, 0x75, 0x69, 0x64, 0x28, 0x01, 0x30, 0x01, 0x40, 0x01, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x43, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x43, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02,
	0x40, 0x01, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x57, 0x0a,
	0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x1c, 0xba,
	0xb9, 0x19, 0x18, 0x0a, 0x16, 0x52, 0x14, 0x69, 0x64, 0x78, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x52, 0x09, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x44, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0a, 0xba, 0xb9, 0x19, 0x06, 0x0a, 0x04, 0x18, 0x64,
	0x40, 0x01, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x0a,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0b, 0xba, 0xb9, 0x19, 0x07, 0x0a, 0x05, 0x18, 0xff, 0x01, 0x40, 0x01, 0x52, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0xb9, 0x19,
	0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08,
	0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x4d, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x6d, 0x6b, 0x69, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x08, 0xba, 0xb9, 0x19,
	0x04, 0x1a, 0x02, 0x38, 0x01, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x3a, 0x09,
	0xf8, 0x42, 0x01, 0xba, 0xb9, 0x19, 0x02, 0x08, 0x01, 0x22, 0xa2, 0x05, 0x0a, 0x07, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x34, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x55,
	0x55, 0x49, 0x44, 0x42, 0x12, 0xba, 0xb9, 0x19, 0x0e, 0x0a, 0x0c, 0x12, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x28, 0x01, 0x30, 0x01, 0x40, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x43, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xba, 0xb9, 0x19,
	0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x43, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x5a, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x1f, 0xba, 0xb9, 0x19, 0x1b, 0x0a, 0x19, 0x52, 0x17, 0x69,
	0x64, 0x78, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x7a, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74,
	0x7a, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x4c, 0x0a, 0x06, 0x67, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x6d, 0x6b, 0x69, 0x74,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68,
	0x64, 0x61, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52,
	0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x18, 0x0a, 0x03, 0x61, 0x67, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x10, 0x01, 0x52, 0x03,
	0x61, 0x67, 0x65, 0x12, 0x45, 0x0a, 0x0f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64,
	0x5f, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x70, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x72, 0x65, 0x64, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x22, 0x57, 0x0a, 0x0a, 0x47, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x47, 0x45, 0x4e, 0x44,
	0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x47,
	0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x45, 0x4d, 0x41, 0x4c,
	0x45, 0x10, 0x02, 0x3a, 0x09, 0xf8, 0x42, 0x01, 0xba, 0xb9, 0x19, 0x02, 0x08, 0x01, 0x42, 0x75,
	0x0a, 0x20, 0x6d, 0x6b, 0x69, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x50, 0x01, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x78, 0x6d, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x65, 0x72, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x6d, 0x6b, 0x69, 0x74, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mkit_service_account_entities_v1_entities_proto_rawDescOnce sync.Once
	file_mkit_service_account_entities_v1_entities_proto_rawDescData = file_mkit_service_account_entities_v1_entities_proto_rawDesc
)

func file_mkit_service_account_entities_v1_entities_proto_rawDescGZIP() []byte {
	file_mkit_service_account_entities_v1_entities_proto_rawDescOnce.Do(func() {
		file_mkit_service_account_entities_v1_entities_proto_rawDescData = protoimpl.X.CompressGZIP(file_mkit_service_account_entities_v1_entities_proto_rawDescData)
	})
	return file_mkit_service_account_entities_v1_entities_proto_rawDescData
}

var file_mkit_service_account_entities_v1_entities_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mkit_service_account_entities_v1_entities_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mkit_service_account_entities_v1_entities_proto_goTypes = []interface{}{
	(Profile_GenderType)(0),      // 0: mkit.service.account.entities.v1.Profile.GenderType
	(*User)(nil),                 // 1: mkit.service.account.entities.v1.User
	(*Profile)(nil),              // 2: mkit.service.account.entities.v1.Profile
	(*types.UUID)(nil),           // 3: gorm.types.UUID
	(*timestamp.Timestamp)(nil),  // 4: google.protobuf.Timestamp
	(*wrappers.StringValue)(nil), // 5: google.protobuf.StringValue
}
var file_mkit_service_account_entities_v1_entities_proto_depIdxs = []int32{
	3,  // 0: mkit.service.account.entities.v1.User.id:type_name -> gorm.types.UUID
	4,  // 1: mkit.service.account.entities.v1.User.created_at:type_name -> google.protobuf.Timestamp
	4,  // 2: mkit.service.account.entities.v1.User.updated_at:type_name -> google.protobuf.Timestamp
	4,  // 3: mkit.service.account.entities.v1.User.deleted_at:type_name -> google.protobuf.Timestamp
	5,  // 4: mkit.service.account.entities.v1.User.username:type_name -> google.protobuf.StringValue
	2,  // 5: mkit.service.account.entities.v1.User.profile:type_name -> mkit.service.account.entities.v1.Profile
	3,  // 6: mkit.service.account.entities.v1.Profile.id:type_name -> gorm.types.UUID
	4,  // 7: mkit.service.account.entities.v1.Profile.created_at:type_name -> google.protobuf.Timestamp
	4,  // 8: mkit.service.account.entities.v1.Profile.updated_at:type_name -> google.protobuf.Timestamp
	4,  // 9: mkit.service.account.entities.v1.Profile.deleted_at:type_name -> google.protobuf.Timestamp
	0,  // 10: mkit.service.account.entities.v1.Profile.gender:type_name -> mkit.service.account.entities.v1.Profile.GenderType
	4,  // 11: mkit.service.account.entities.v1.Profile.birthday:type_name -> google.protobuf.Timestamp
	5,  // 12: mkit.service.account.entities.v1.Profile.preferred_theme:type_name -> google.protobuf.StringValue
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_mkit_service_account_entities_v1_entities_proto_init() }
func file_mkit_service_account_entities_v1_entities_proto_init() {
	if File_mkit_service_account_entities_v1_entities_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mkit_service_account_entities_v1_entities_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_mkit_service_account_entities_v1_entities_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile); i {
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
			RawDescriptor: file_mkit_service_account_entities_v1_entities_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mkit_service_account_entities_v1_entities_proto_goTypes,
		DependencyIndexes: file_mkit_service_account_entities_v1_entities_proto_depIdxs,
		EnumInfos:         file_mkit_service_account_entities_v1_entities_proto_enumTypes,
		MessageInfos:      file_mkit_service_account_entities_v1_entities_proto_msgTypes,
	}.Build()
	File_mkit_service_account_entities_v1_entities_proto = out.File
	file_mkit_service_account_entities_v1_entities_proto_rawDesc = nil
	file_mkit_service_account_entities_v1_entities_proto_goTypes = nil
	file_mkit_service_account_entities_v1_entities_proto_depIdxs = nil
}
