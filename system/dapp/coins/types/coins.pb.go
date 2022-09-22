// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.9.1
// source: coins.proto

package types

import (
	reflect "reflect"
	sync "sync"

	types "github.com/assetcloud/chain/types"
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

// message for execs.coins
type CoinsAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*CoinsAction_Transfer
	//	*CoinsAction_Withdraw
	//	*CoinsAction_Genesis
	//	*CoinsAction_TransferToExec
	Value isCoinsAction_Value `protobuf_oneof:"value"`
	Ty    int32               `protobuf:"varint,3,opt,name=ty,proto3" json:"ty,omitempty"`
}

func (x *CoinsAction) Reset() {
	*x = CoinsAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coins_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoinsAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoinsAction) ProtoMessage() {}

func (x *CoinsAction) ProtoReflect() protoreflect.Message {
	mi := &file_coins_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoinsAction.ProtoReflect.Descriptor instead.
func (*CoinsAction) Descriptor() ([]byte, []int) {
	return file_coins_proto_rawDescGZIP(), []int{0}
}

func (m *CoinsAction) GetValue() isCoinsAction_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *CoinsAction) GetTransfer() *types.AssetsTransfer {
	if x, ok := x.GetValue().(*CoinsAction_Transfer); ok {
		return x.Transfer
	}
	return nil
}

func (x *CoinsAction) GetWithdraw() *types.AssetsWithdraw {
	if x, ok := x.GetValue().(*CoinsAction_Withdraw); ok {
		return x.Withdraw
	}
	return nil
}

func (x *CoinsAction) GetGenesis() *types.AssetsGenesis {
	if x, ok := x.GetValue().(*CoinsAction_Genesis); ok {
		return x.Genesis
	}
	return nil
}

func (x *CoinsAction) GetTransferToExec() *types.AssetsTransferToExec {
	if x, ok := x.GetValue().(*CoinsAction_TransferToExec); ok {
		return x.TransferToExec
	}
	return nil
}

func (x *CoinsAction) GetTy() int32 {
	if x != nil {
		return x.Ty
	}
	return 0
}

type isCoinsAction_Value interface {
	isCoinsAction_Value()
}

type CoinsAction_Transfer struct {
	Transfer *types.AssetsTransfer `protobuf:"bytes,1,opt,name=transfer,proto3,oneof"`
}

type CoinsAction_Withdraw struct {
	Withdraw *types.AssetsWithdraw `protobuf:"bytes,4,opt,name=withdraw,proto3,oneof"`
}

type CoinsAction_Genesis struct {
	Genesis *types.AssetsGenesis `protobuf:"bytes,2,opt,name=genesis,proto3,oneof"`
}

type CoinsAction_TransferToExec struct {
	TransferToExec *types.AssetsTransferToExec `protobuf:"bytes,5,opt,name=transferToExec,proto3,oneof"`
}

func (*CoinsAction_Transfer) isCoinsAction_Value() {}

func (*CoinsAction_Withdraw) isCoinsAction_Value() {}

func (*CoinsAction_Genesis) isCoinsAction_Value() {}

func (*CoinsAction_TransferToExec) isCoinsAction_Value() {}

var File_coins_proto protoreflect.FileDescriptor

var file_coins_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x1a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x02, 0x0a, 0x0b, 0x43, 0x6f, 0x69, 0x6e,
	0x73, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x48, 0x00, 0x52, 0x08, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x08,
	0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x57, 0x69, 0x74,
	0x68, 0x64, 0x72, 0x61, 0x77, 0x48, 0x00, 0x52, 0x08, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61,
	0x77, 0x12, 0x30, 0x0a, 0x07, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x73, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x48, 0x00, 0x52, 0x07, 0x67, 0x65, 0x6e, 0x65,
	0x73, 0x69, 0x73, 0x12, 0x45, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x54,
	0x6f, 0x45, 0x78, 0x65, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x54, 0x6f, 0x45, 0x78, 0x65, 0x63, 0x48, 0x00, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x54, 0x6f, 0x45, 0x78, 0x65, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x74, 0x79, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coins_proto_rawDescOnce sync.Once
	file_coins_proto_rawDescData = file_coins_proto_rawDesc
)

func file_coins_proto_rawDescGZIP() []byte {
	file_coins_proto_rawDescOnce.Do(func() {
		file_coins_proto_rawDescData = protoimpl.X.CompressGZIP(file_coins_proto_rawDescData)
	})
	return file_coins_proto_rawDescData
}

var file_coins_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_coins_proto_goTypes = []interface{}{
	(*CoinsAction)(nil),                // 0: types.CoinsAction
	(*types.AssetsTransfer)(nil),       // 1: types.AssetsTransfer
	(*types.AssetsWithdraw)(nil),       // 2: types.AssetsWithdraw
	(*types.AssetsGenesis)(nil),        // 3: types.AssetsGenesis
	(*types.AssetsTransferToExec)(nil), // 4: types.AssetsTransferToExec
}
var file_coins_proto_depIdxs = []int32{
	1, // 0: types.CoinsAction.transfer:type_name -> types.AssetsTransfer
	2, // 1: types.CoinsAction.withdraw:type_name -> types.AssetsWithdraw
	3, // 2: types.CoinsAction.genesis:type_name -> types.AssetsGenesis
	4, // 3: types.CoinsAction.transferToExec:type_name -> types.AssetsTransferToExec
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_coins_proto_init() }
func file_coins_proto_init() {
	if File_coins_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coins_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoinsAction); i {
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
	file_coins_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*CoinsAction_Transfer)(nil),
		(*CoinsAction_Withdraw)(nil),
		(*CoinsAction_Genesis)(nil),
		(*CoinsAction_TransferToExec)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_coins_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_coins_proto_goTypes,
		DependencyIndexes: file_coins_proto_depIdxs,
		MessageInfos:      file_coins_proto_msgTypes,
	}.Build()
	File_coins_proto = out.File
	file_coins_proto_rawDesc = nil
	file_coins_proto_goTypes = nil
	file_coins_proto_depIdxs = nil
}
