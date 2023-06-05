// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: biz.proto

package golang

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

type UserSex int32

const (
	UserSex_UserSex_MALE   UserSex = 0
	UserSex_UserSex_FEMALE UserSex = 1
)

// Enum value maps for UserSex.
var (
	UserSex_name = map[int32]string{
		0: "UserSex_MALE",
		1: "UserSex_FEMALE",
	}
	UserSex_value = map[string]int32{
		"UserSex_MALE":   0,
		"UserSex_FEMALE": 1,
	}
)

func (x UserSex) Enum() *UserSex {
	p := new(UserSex)
	*p = x
	return p
}

func (x UserSex) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserSex) Descriptor() protoreflect.EnumDescriptor {
	return file_biz_proto_enumTypes[0].Descriptor()
}

func (UserSex) Type() protoreflect.EnumType {
	return &file_biz_proto_enumTypes[0]
}

func (x UserSex) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserSex.Descriptor instead.
func (UserSex) EnumDescriptor() ([]byte, []int) {
	return file_biz_proto_rawDescGZIP(), []int{0}
}

type RequestGetUsers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MessageId int32 `protobuf:"varint,2,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	AuthKey   int32 `protobuf:"varint,3,opt,name=auth_key,json=authKey,proto3" json:"auth_key,omitempty"`
}

func (x *RequestGetUsers) Reset() {
	*x = RequestGetUsers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biz_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestGetUsers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestGetUsers) ProtoMessage() {}

func (x *RequestGetUsers) ProtoReflect() protoreflect.Message {
	mi := &file_biz_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestGetUsers.ProtoReflect.Descriptor instead.
func (*RequestGetUsers) Descriptor() ([]byte, []int) {
	return file_biz_proto_rawDescGZIP(), []int{0}
}

func (x *RequestGetUsers) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RequestGetUsers) GetMessageId() int32 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

func (x *RequestGetUsers) GetAuthKey() int32 {
	if x != nil {
		return x.AuthKey
	}
	return 0
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Family string  `protobuf:"bytes,2,opt,name=family,proto3" json:"family,omitempty"`
	Id     int32   `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Age    int32   `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
	Sex    UserSex `protobuf:"varint,5,opt,name=sex,proto3,enum=auth.UserSex" json:"sex,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biz_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_biz_proto_msgTypes[1]
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
	return file_biz_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetFamily() string {
	if x != nil {
		return x.Family
	}
	return ""
}

func (x *User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *User) GetSex() UserSex {
	if x != nil {
		return x.Sex
	}
	return UserSex_UserSex_MALE
}

type RequestGetUsersWithSQLInjection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MessageId int32  `protobuf:"varint,2,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	AuthKey   int32  `protobuf:"varint,3,opt,name=auth_key,json=authKey,proto3" json:"auth_key,omitempty"`
}

func (x *RequestGetUsersWithSQLInjection) Reset() {
	*x = RequestGetUsersWithSQLInjection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biz_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestGetUsersWithSQLInjection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestGetUsersWithSQLInjection) ProtoMessage() {}

func (x *RequestGetUsersWithSQLInjection) ProtoReflect() protoreflect.Message {
	mi := &file_biz_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestGetUsersWithSQLInjection.ProtoReflect.Descriptor instead.
func (*RequestGetUsersWithSQLInjection) Descriptor() ([]byte, []int) {
	return file_biz_proto_rawDescGZIP(), []int{2}
}

func (x *RequestGetUsersWithSQLInjection) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RequestGetUsersWithSQLInjection) GetMessageId() int32 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

func (x *RequestGetUsersWithSQLInjection) GetAuthKey() int32 {
	if x != nil {
		return x.AuthKey
	}
	return 0
}

type ResponseGetUsers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users     []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	MessageId int32   `protobuf:"varint,2,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
}

func (x *ResponseGetUsers) Reset() {
	*x = ResponseGetUsers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biz_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseGetUsers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseGetUsers) ProtoMessage() {}

func (x *ResponseGetUsers) ProtoReflect() protoreflect.Message {
	mi := &file_biz_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseGetUsers.ProtoReflect.Descriptor instead.
func (*ResponseGetUsers) Descriptor() ([]byte, []int) {
	return file_biz_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseGetUsers) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *ResponseGetUsers) GetMessageId() int32 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

var File_biz_proto protoreflect.FileDescriptor

var file_biz_proto_rawDesc = []byte{
	0x0a, 0x09, 0x62, 0x69, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74,
	0x68, 0x22, 0x64, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x61, 0x75, 0x74, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x61, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x22, 0x75, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a,
	0x03, 0x73, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x78, 0x52, 0x03, 0x73, 0x65, 0x78, 0x22, 0x74,
	0x0a, 0x1f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x57, 0x69, 0x74, 0x68, 0x53, 0x51, 0x4c, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x75, 0x74,
	0x68, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x75, 0x74,
	0x68, 0x4b, 0x65, 0x79, 0x22, 0x53, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x20, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x2a, 0x2f, 0x0a, 0x07, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x65, 0x78, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x78, 0x5f,
	0x4d, 0x41, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x78, 0x5f, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x32, 0xa0, 0x01, 0x0a, 0x03, 0x42,
	0x69, 0x7a, 0x12, 0x3a, 0x0a, 0x09, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12,
	0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x1a, 0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x5d,
	0x0a, 0x1c, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x77, 0x69, 0x74, 0x68,
	0x5f, 0x73, 0x71, 0x6c, 0x5f, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x57, 0x69, 0x74, 0x68, 0x53, 0x51, 0x4c, 0x49, 0x6e, 0x6a, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x0c, 0x5a,
	0x0a, 0x73, 0x72, 0x63, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_biz_proto_rawDescOnce sync.Once
	file_biz_proto_rawDescData = file_biz_proto_rawDesc
)

func file_biz_proto_rawDescGZIP() []byte {
	file_biz_proto_rawDescOnce.Do(func() {
		file_biz_proto_rawDescData = protoimpl.X.CompressGZIP(file_biz_proto_rawDescData)
	})
	return file_biz_proto_rawDescData
}

var file_biz_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_biz_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_biz_proto_goTypes = []interface{}{
	(UserSex)(0),                            // 0: auth.UserSex
	(*RequestGetUsers)(nil),                 // 1: auth.RequestGetUsers
	(*User)(nil),                            // 2: auth.User
	(*RequestGetUsersWithSQLInjection)(nil), // 3: auth.RequestGetUsersWithSQLInjection
	(*ResponseGetUsers)(nil),                // 4: auth.ResponseGetUsers
}
var file_biz_proto_depIdxs = []int32{
	0, // 0: auth.User.sex:type_name -> auth.UserSex
	2, // 1: auth.ResponseGetUsers.users:type_name -> auth.User
	1, // 2: auth.Biz.get_users:input_type -> auth.RequestGetUsers
	3, // 3: auth.Biz.get_users_with_sql_injection:input_type -> auth.RequestGetUsersWithSQLInjection
	4, // 4: auth.Biz.get_users:output_type -> auth.ResponseGetUsers
	4, // 5: auth.Biz.get_users_with_sql_injection:output_type -> auth.ResponseGetUsers
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_biz_proto_init() }
func file_biz_proto_init() {
	if File_biz_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_biz_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestGetUsers); i {
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
		file_biz_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_biz_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestGetUsersWithSQLInjection); i {
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
		file_biz_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseGetUsers); i {
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
			RawDescriptor: file_biz_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_biz_proto_goTypes,
		DependencyIndexes: file_biz_proto_depIdxs,
		EnumInfos:         file_biz_proto_enumTypes,
		MessageInfos:      file_biz_proto_msgTypes,
	}.Build()
	File_biz_proto = out.File
	file_biz_proto_rawDesc = nil
	file_biz_proto_goTypes = nil
	file_biz_proto_depIdxs = nil
}