// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// protoc --go_out=plugins=grpc:./ ./sign.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.0
// source: sign.proto

package script

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LockScript   []byte `protobuf:"bytes,1,opt,name=lockScript,proto3" json:"lockScript,omitempty"`
	UnlockScript []byte `protobuf:"bytes,2,opt,name=unlockScript,proto3" json:"unlockScript,omitempty"`
	LockTime     int64  `protobuf:"varint,3,opt,name=lockTime,proto3" json:"lockTime,omitempty"`
	UtxoSequence int64  `protobuf:"varint,4,opt,name=utxoSequence,proto3" json:"utxoSequence,omitempty"`
}

func (x *Signature) Reset() {
	*x = Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sign_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature) ProtoMessage() {}

func (x *Signature) ProtoReflect() protoreflect.Message {
	mi := &file_sign_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature.ProtoReflect.Descriptor instead.
func (*Signature) Descriptor() ([]byte, []int) {
	return file_sign_proto_rawDescGZIP(), []int{0}
}

func (x *Signature) GetLockScript() []byte {
	if x != nil {
		return x.LockScript
	}
	return nil
}

func (x *Signature) GetUnlockScript() []byte {
	if x != nil {
		return x.UnlockScript
	}
	return nil
}

func (x *Signature) GetLockTime() int64 {
	if x != nil {
		return x.LockTime
	}
	return 0
}

func (x *Signature) GetUtxoSequence() int64 {
	if x != nil {
		return x.UtxoSequence
	}
	return 0
}

var File_sign_proto protoreflect.FileDescriptor

var file_sign_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x62, 0x74,
	0x63, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x22, 0x8f, 0x01, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x6b, 0x53,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x53,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x75, 0x6e, 0x6c,
	0x6f, 0x63, 0x6b, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x6b, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x6f, 0x63,
	0x6b, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x75, 0x74, 0x78, 0x6f, 0x53, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x75, 0x74, 0x78,
	0x6f, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2e, 0x2f,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sign_proto_rawDescOnce sync.Once
	file_sign_proto_rawDescData = file_sign_proto_rawDesc
)

func file_sign_proto_rawDescGZIP() []byte {
	file_sign_proto_rawDescOnce.Do(func() {
		file_sign_proto_rawDescData = protoimpl.X.CompressGZIP(file_sign_proto_rawDescData)
	})
	return file_sign_proto_rawDescData
}

var file_sign_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sign_proto_goTypes = []interface{}{
	(*Signature)(nil), // 0: btcscript.Signature
}
var file_sign_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sign_proto_init() }
func file_sign_proto_init() {
	if File_sign_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sign_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signature); i {
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
			RawDescriptor: file_sign_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sign_proto_goTypes,
		DependencyIndexes: file_sign_proto_depIdxs,
		MessageInfos:      file_sign_proto_msgTypes,
	}.Build()
	File_sign_proto = out.File
	file_sign_proto_rawDesc = nil
	file_sign_proto_goTypes = nil
	file_sign_proto_depIdxs = nil
}
