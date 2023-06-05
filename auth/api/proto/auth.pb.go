// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: auth.proto

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

type RequestReqPQ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nonce     string `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	MessageId int32  `protobuf:"varint,2,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
}

func (x *RequestReqPQ) Reset() {
	*x = RequestReqPQ{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestReqPQ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestReqPQ) ProtoMessage() {}

func (x *RequestReqPQ) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestReqPQ.ProtoReflect.Descriptor instead.
func (*RequestReqPQ) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

func (x *RequestReqPQ) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *RequestReqPQ) GetMessageId() int32 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

type ResponseReqPQ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nonce       string `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ServerNonce string `protobuf:"bytes,2,opt,name=server_nonce,json=serverNonce,proto3" json:"server_nonce,omitempty"`
	MessageId   int32  `protobuf:"varint,3,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	P           int64  `protobuf:"varint,4,opt,name=p,proto3" json:"p,omitempty"`
	G           int64  `protobuf:"varint,5,opt,name=g,proto3" json:"g,omitempty"`
}

func (x *ResponseReqPQ) Reset() {
	*x = ResponseReqPQ{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseReqPQ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseReqPQ) ProtoMessage() {}

func (x *ResponseReqPQ) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseReqPQ.ProtoReflect.Descriptor instead.
func (*ResponseReqPQ) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseReqPQ) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *ResponseReqPQ) GetServerNonce() string {
	if x != nil {
		return x.ServerNonce
	}
	return ""
}

func (x *ResponseReqPQ) GetMessageId() int32 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

func (x *ResponseReqPQ) GetP() int64 {
	if x != nil {
		return x.P
	}
	return 0
}

func (x *ResponseReqPQ) GetG() int64 {
	if x != nil {
		return x.G
	}
	return 0
}

type RequestReqDHParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nonce       string `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ServerNonce string `protobuf:"bytes,2,opt,name=server_nonce,json=serverNonce,proto3" json:"server_nonce,omitempty"`
	MessageId   int32  `protobuf:"varint,3,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	A           int64  `protobuf:"varint,4,opt,name=a,proto3" json:"a,omitempty"`
}

func (x *RequestReqDHParams) Reset() {
	*x = RequestReqDHParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestReqDHParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestReqDHParams) ProtoMessage() {}

func (x *RequestReqDHParams) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestReqDHParams.ProtoReflect.Descriptor instead.
func (*RequestReqDHParams) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{2}
}

func (x *RequestReqDHParams) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *RequestReqDHParams) GetServerNonce() string {
	if x != nil {
		return x.ServerNonce
	}
	return ""
}

func (x *RequestReqDHParams) GetMessageId() int32 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

func (x *RequestReqDHParams) GetA() int64 {
	if x != nil {
		return x.A
	}
	return 0
}

type ResponseReqDHParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nonce       string `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ServerNonce string `protobuf:"bytes,2,opt,name=server_nonce,json=serverNonce,proto3" json:"server_nonce,omitempty"`
	MessageId   int32  `protobuf:"varint,3,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	B           int64  `protobuf:"varint,4,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *ResponseReqDHParams) Reset() {
	*x = ResponseReqDHParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseReqDHParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseReqDHParams) ProtoMessage() {}

func (x *ResponseReqDHParams) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseReqDHParams.ProtoReflect.Descriptor instead.
func (*ResponseReqDHParams) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseReqDHParams) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *ResponseReqDHParams) GetServerNonce() string {
	if x != nil {
		return x.ServerNonce
	}
	return ""
}

func (x *ResponseReqDHParams) GetMessageId() int32 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

func (x *ResponseReqDHParams) GetB() int64 {
	if x != nil {
		return x.B
	}
	return 0
}

var File_auth_proto protoreflect.FileDescriptor

var file_auth_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75,
	0x74, 0x68, 0x22, 0x43, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x50, 0x51, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x83, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x50, 0x51, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x6f, 0x6e,
	0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x12, 0x0c, 0x0a, 0x01, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x70, 0x12,
	0x0c, 0x0a, 0x01, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x67, 0x22, 0x7a, 0x0a,
	0x12, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x44, 0x48, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x61,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x61, 0x22, 0x7b, 0x0a, 0x13, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x44, 0x48, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x01, 0x62, 0x32, 0x7f, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x31,
	0x0a, 0x06, 0x72, 0x65, 0x71, 0x5f, 0x70, 0x71, 0x12, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x50, 0x51, 0x1a, 0x13, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x50,
	0x51, 0x12, 0x44, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x5f, 0x44, 0x48, 0x5f, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x44, 0x48, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x19, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x44,
	0x48, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x0c, 0x5a, 0x0a, 0x73, 0x72, 0x63, 0x2f, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_proto_rawDescOnce sync.Once
	file_auth_proto_rawDescData = file_auth_proto_rawDesc
)

func file_auth_proto_rawDescGZIP() []byte {
	file_auth_proto_rawDescOnce.Do(func() {
		file_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_proto_rawDescData)
	})
	return file_auth_proto_rawDescData
}

var file_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_auth_proto_goTypes = []interface{}{
	(*RequestReqPQ)(nil),        // 0: auth.RequestReqPQ
	(*ResponseReqPQ)(nil),       // 1: auth.ResponseReqPQ
	(*RequestReqDHParams)(nil),  // 2: auth.RequestReqDHParams
	(*ResponseReqDHParams)(nil), // 3: auth.ResponseReqDHParams
}
var file_auth_proto_depIdxs = []int32{
	0, // 0: auth.Auth.req_pq:input_type -> auth.RequestReqPQ
	2, // 1: auth.Auth.req_DH_params:input_type -> auth.RequestReqDHParams
	1, // 2: auth.Auth.req_pq:output_type -> auth.ResponseReqPQ
	3, // 3: auth.Auth.req_DH_params:output_type -> auth.ResponseReqDHParams
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_proto_init() }
func file_auth_proto_init() {
	if File_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestReqPQ); i {
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
		file_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseReqPQ); i {
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
		file_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestReqDHParams); i {
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
		file_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseReqDHParams); i {
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
			RawDescriptor: file_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_proto_goTypes,
		DependencyIndexes: file_auth_proto_depIdxs,
		MessageInfos:      file_auth_proto_msgTypes,
	}.Build()
	File_auth_proto = out.File
	file_auth_proto_rawDesc = nil
	file_auth_proto_goTypes = nil
	file_auth_proto_depIdxs = nil
}
