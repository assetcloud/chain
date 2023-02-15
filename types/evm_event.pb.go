// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.5.0
// source: evm_event.proto

package types

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

// 一条evm event log数据
type EVMLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic [][]byte `protobuf:"bytes,1,rep,name=topic,proto3" json:"topic,omitempty"`
	Data  []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *EVMLog) Reset() {
	*x = EVMLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evm_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EVMLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EVMLog) ProtoMessage() {}

func (x *EVMLog) ProtoReflect() protoreflect.Message {
	mi := &file_evm_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EVMLog.ProtoReflect.Descriptor instead.
func (*EVMLog) Descriptor() ([]byte, []int) {
	return file_evm_event_proto_rawDescGZIP(), []int{0}
}

func (x *EVMLog) GetTopic() [][]byte {
	if x != nil {
		return x.Topic
	}
	return nil
}

func (x *EVMLog) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// 多条evm event log数据
type EVMLogsPerTx struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Logs []*EVMLog `protobuf:"bytes,1,rep,name=logs,proto3" json:"logs,omitempty"`
}

func (x *EVMLogsPerTx) Reset() {
	*x = EVMLogsPerTx{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evm_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EVMLogsPerTx) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EVMLogsPerTx) ProtoMessage() {}

func (x *EVMLogsPerTx) ProtoReflect() protoreflect.Message {
	mi := &file_evm_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EVMLogsPerTx.ProtoReflect.Descriptor instead.
func (*EVMLogsPerTx) Descriptor() ([]byte, []int) {
	return file_evm_event_proto_rawDescGZIP(), []int{1}
}

func (x *EVMLogsPerTx) GetLogs() []*EVMLog {
	if x != nil {
		return x.Logs
	}
	return nil
}

type EVMTxAndLogs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tx        *Transaction  `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
	LogsPerTx *EVMLogsPerTx `protobuf:"bytes,2,opt,name=logsPerTx,proto3" json:"logsPerTx,omitempty"`
}

func (x *EVMTxAndLogs) Reset() {
	*x = EVMTxAndLogs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evm_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EVMTxAndLogs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EVMTxAndLogs) ProtoMessage() {}

func (x *EVMTxAndLogs) ProtoReflect() protoreflect.Message {
	mi := &file_evm_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EVMTxAndLogs.ProtoReflect.Descriptor instead.
func (*EVMTxAndLogs) Descriptor() ([]byte, []int) {
	return file_evm_event_proto_rawDescGZIP(), []int{2}
}

func (x *EVMTxAndLogs) GetTx() *Transaction {
	if x != nil {
		return x.Tx
	}
	return nil
}

func (x *EVMTxAndLogs) GetLogsPerTx() *EVMLogsPerTx {
	if x != nil {
		return x.LogsPerTx
	}
	return nil
}

//一个块中包含的多条evm event log数据
type EVMTxLogPerBlk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxAndLogs    []*EVMTxAndLogs `protobuf:"bytes,1,rep,name=txAndLogs,proto3" json:"txAndLogs,omitempty"`
	Height       int64           `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	BlockHash    []byte          `protobuf:"bytes,3,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	ParentHash   []byte          `protobuf:"bytes,4,opt,name=parentHash,proto3" json:"parentHash,omitempty"`
	PreviousHash []byte          `protobuf:"bytes,5,opt,name=previousHash,proto3" json:"previousHash,omitempty"`
	AddDelType   int32           `protobuf:"varint,6,opt,name=addDelType,proto3" json:"addDelType,omitempty"`
	SeqNum       int64           `protobuf:"varint,7,opt,name=seqNum,proto3" json:"seqNum,omitempty"`
}

func (x *EVMTxLogPerBlk) Reset() {
	*x = EVMTxLogPerBlk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evm_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EVMTxLogPerBlk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EVMTxLogPerBlk) ProtoMessage() {}

func (x *EVMTxLogPerBlk) ProtoReflect() protoreflect.Message {
	mi := &file_evm_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EVMTxLogPerBlk.ProtoReflect.Descriptor instead.
func (*EVMTxLogPerBlk) Descriptor() ([]byte, []int) {
	return file_evm_event_proto_rawDescGZIP(), []int{3}
}

func (x *EVMTxLogPerBlk) GetTxAndLogs() []*EVMTxAndLogs {
	if x != nil {
		return x.TxAndLogs
	}
	return nil
}

func (x *EVMTxLogPerBlk) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *EVMTxLogPerBlk) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

func (x *EVMTxLogPerBlk) GetParentHash() []byte {
	if x != nil {
		return x.ParentHash
	}
	return nil
}

func (x *EVMTxLogPerBlk) GetPreviousHash() []byte {
	if x != nil {
		return x.PreviousHash
	}
	return nil
}

func (x *EVMTxLogPerBlk) GetAddDelType() int32 {
	if x != nil {
		return x.AddDelType
	}
	return 0
}

func (x *EVMTxLogPerBlk) GetSeqNum() int64 {
	if x != nil {
		return x.SeqNum
	}
	return 0
}

//多个块中包含的多条evm event log数据
type EVMTxLogsInBlks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Logs4EVMPerBlk []*EVMTxLogPerBlk `protobuf:"bytes,1,rep,name=logs4EVMPerBlk,proto3" json:"logs4EVMPerBlk,omitempty"`
}

func (x *EVMTxLogsInBlks) Reset() {
	*x = EVMTxLogsInBlks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evm_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EVMTxLogsInBlks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EVMTxLogsInBlks) ProtoMessage() {}

func (x *EVMTxLogsInBlks) ProtoReflect() protoreflect.Message {
	mi := &file_evm_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EVMTxLogsInBlks.ProtoReflect.Descriptor instead.
func (*EVMTxLogsInBlks) Descriptor() ([]byte, []int) {
	return file_evm_event_proto_rawDescGZIP(), []int{4}
}

func (x *EVMTxLogsInBlks) GetLogs4EVMPerBlk() []*EVMTxLogPerBlk {
	if x != nil {
		return x.Logs4EVMPerBlk
	}
	return nil
}

// 创建/调用合约的请求结构
type EVMContractAction4Chain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 转账金额
	Amount uint64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	// 消耗限制，默认为Transaction.Fee
	GasLimit uint64 `protobuf:"varint,2,opt,name=gasLimit,proto3" json:"gasLimit,omitempty"`
	// gas价格，默认为1
	GasPrice uint32 `protobuf:"varint,3,opt,name=gasPrice,proto3" json:"gasPrice,omitempty"`
	// 合约数据
	Code []byte `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	//交易参数
	Para []byte `protobuf:"bytes,5,opt,name=para,proto3" json:"para,omitempty"`
	// 合约别名，方便识别
	Alias string `protobuf:"bytes,6,opt,name=alias,proto3" json:"alias,omitempty"`
	// 交易备注
	Note string `protobuf:"bytes,7,opt,name=note,proto3" json:"note,omitempty"`
	// 调用合约地址
	ContractAddr string `protobuf:"bytes,8,opt,name=contractAddr,proto3" json:"contractAddr,omitempty"`
}

func (x *EVMContractAction4Chain) Reset() {
	*x = EVMContractAction4Chain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evm_event_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EVMContractAction4Chain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EVMContractAction4Chain) ProtoMessage() {}

func (x *EVMContractAction4Chain) ProtoReflect() protoreflect.Message {
	mi := &file_evm_event_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EVMContractAction4Chain.ProtoReflect.Descriptor instead.
func (*EVMContractAction4Chain) Descriptor() ([]byte, []int) {
	return file_evm_event_proto_rawDescGZIP(), []int{5}
}

func (x *EVMContractAction4Chain) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *EVMContractAction4Chain) GetGasLimit() uint64 {
	if x != nil {
		return x.GasLimit
	}
	return 0
}

func (x *EVMContractAction4Chain) GetGasPrice() uint32 {
	if x != nil {
		return x.GasPrice
	}
	return 0
}

func (x *EVMContractAction4Chain) GetCode() []byte {
	if x != nil {
		return x.Code
	}
	return nil
}

func (x *EVMContractAction4Chain) GetPara() []byte {
	if x != nil {
		return x.Para
	}
	return nil
}

func (x *EVMContractAction4Chain) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *EVMContractAction4Chain) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *EVMContractAction4Chain) GetContractAddr() string {
	if x != nil {
		return x.ContractAddr
	}
	return ""
}

var File_evm_event_proto protoreflect.FileDescriptor

var file_evm_event_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x65, 0x76, 0x6d, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x06, 0x45,
	0x56, 0x4d, 0x4c, 0x6f, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0c, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x31, 0x0a, 0x0c, 0x45, 0x56, 0x4d, 0x4c, 0x6f, 0x67, 0x73, 0x50, 0x65, 0x72, 0x54, 0x78, 0x12,
	0x21, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x56, 0x4d, 0x4c, 0x6f, 0x67, 0x52, 0x04, 0x6c, 0x6f,
	0x67, 0x73, 0x22, 0x65, 0x0a, 0x0c, 0x45, 0x56, 0x4d, 0x54, 0x78, 0x41, 0x6e, 0x64, 0x4c, 0x6f,
	0x67, 0x73, 0x12, 0x22, 0x0a, 0x02, 0x74, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x02, 0x74, 0x78, 0x12, 0x31, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x73, 0x50, 0x65,
	0x72, 0x54, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x45, 0x56, 0x4d, 0x4c, 0x6f, 0x67, 0x73, 0x50, 0x65, 0x72, 0x54, 0x78, 0x52, 0x09,
	0x6c, 0x6f, 0x67, 0x73, 0x50, 0x65, 0x72, 0x54, 0x78, 0x22, 0xf5, 0x01, 0x0a, 0x0e, 0x45, 0x56,
	0x4d, 0x54, 0x78, 0x4c, 0x6f, 0x67, 0x50, 0x65, 0x72, 0x42, 0x6c, 0x6b, 0x12, 0x31, 0x0a, 0x09,
	0x74, 0x78, 0x41, 0x6e, 0x64, 0x4c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x56, 0x4d, 0x54, 0x78, 0x41, 0x6e, 0x64,
	0x4c, 0x6f, 0x67, 0x73, 0x52, 0x09, 0x74, 0x78, 0x41, 0x6e, 0x64, 0x4c, 0x6f, 0x67, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x48, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48,
	0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75,
	0x73, 0x48, 0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x70, 0x72, 0x65,
	0x76, 0x69, 0x6f, 0x75, 0x73, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x64, 0x64,
	0x44, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x61,
	0x64, 0x64, 0x44, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x71,
	0x4e, 0x75, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x65, 0x71, 0x4e, 0x75,
	0x6d, 0x22, 0x50, 0x0a, 0x0f, 0x45, 0x56, 0x4d, 0x54, 0x78, 0x4c, 0x6f, 0x67, 0x73, 0x49, 0x6e,
	0x42, 0x6c, 0x6b, 0x73, 0x12, 0x3d, 0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x73, 0x34, 0x45, 0x56, 0x4d,
	0x50, 0x65, 0x72, 0x42, 0x6c, 0x6b, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x56, 0x4d, 0x54, 0x78, 0x4c, 0x6f, 0x67, 0x50, 0x65, 0x72,
	0x42, 0x6c, 0x6b, 0x52, 0x0e, 0x6c, 0x6f, 0x67, 0x73, 0x34, 0x45, 0x56, 0x4d, 0x50, 0x65, 0x72,
	0x42, 0x6c, 0x6b, 0x22, 0xdf, 0x01, 0x0a, 0x17, 0x45, 0x56, 0x4d, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x34, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x61, 0x73, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x67, 0x61, 0x73, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x67, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x72, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x70, 0x61, 0x72, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x64, 0x64,
	0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x41, 0x64, 0x64, 0x72, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_evm_event_proto_rawDescOnce sync.Once
	file_evm_event_proto_rawDescData = file_evm_event_proto_rawDesc
)

func file_evm_event_proto_rawDescGZIP() []byte {
	file_evm_event_proto_rawDescOnce.Do(func() {
		file_evm_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_evm_event_proto_rawDescData)
	})
	return file_evm_event_proto_rawDescData
}

var file_evm_event_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_evm_event_proto_goTypes = []interface{}{
	(*EVMLog)(nil),                  // 0: types.EVMLog
	(*EVMLogsPerTx)(nil),            // 1: types.EVMLogsPerTx
	(*EVMTxAndLogs)(nil),            // 2: types.EVMTxAndLogs
	(*EVMTxLogPerBlk)(nil),          // 3: types.EVMTxLogPerBlk
	(*EVMTxLogsInBlks)(nil),         // 4: types.EVMTxLogsInBlks
	(*EVMContractAction4Chain)(nil), // 5: types.EVMContractAction4Chain
	(*Transaction)(nil),             // 6: types.Transaction
}
var file_evm_event_proto_depIdxs = []int32{
	0, // 0: types.EVMLogsPerTx.logs:type_name -> types.EVMLog
	6, // 1: types.EVMTxAndLogs.tx:type_name -> types.Transaction
	1, // 2: types.EVMTxAndLogs.logsPerTx:type_name -> types.EVMLogsPerTx
	2, // 3: types.EVMTxLogPerBlk.txAndLogs:type_name -> types.EVMTxAndLogs
	3, // 4: types.EVMTxLogsInBlks.logs4EVMPerBlk:type_name -> types.EVMTxLogPerBlk
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_evm_event_proto_init() }
func file_evm_event_proto_init() {
	if File_evm_event_proto != nil {
		return
	}
	file_transaction_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_evm_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EVMLog); i {
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
		file_evm_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EVMLogsPerTx); i {
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
		file_evm_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EVMTxAndLogs); i {
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
		file_evm_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EVMTxLogPerBlk); i {
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
		file_evm_event_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EVMTxLogsInBlks); i {
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
		file_evm_event_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EVMContractAction4Chain); i {
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
			RawDescriptor: file_evm_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_evm_event_proto_goTypes,
		DependencyIndexes: file_evm_event_proto_depIdxs,
		MessageInfos:      file_evm_event_proto_msgTypes,
	}.Build()
	File_evm_event_proto = out.File
	file_evm_event_proto_rawDesc = nil
	file_evm_event_proto_goTypes = nil
	file_evm_event_proto_depIdxs = nil
}
