# 交易签名

## 签名类型

> 即签名结构中 Ty 字段, 以下统称为 signID

```go
type Signature struct {

      	Ty     int32  `protobuf:"varint,1,opt,name=ty,proto3" json:"ty,omitempty"`
      	Pubkey []byte `protobuf:"bytes,2,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
      	Signature []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
 }
```

#### 底层设计

> 最初设计中, signID 即 cryptoID, 在引入多地址格式兼容(#1181)后, signID 集成了 cryptoID 和 addressID 含义

- addressID, 即签名公钥转换为地址时采用的地址类型 ID, 在 signID 中占 3 位, 即低位第 13~15 位
- cryptoID, 即签名对应的加密算法插件类型 ID, 在 signID 中, 除 addressID 以外的其余位
- 默认 addressID 为 0, 此时 cryptoID 和 signID 值依然相等

#### 集成算法

> 通过位运算集成(兼容最初的 signID 逻辑, cryptoID 只能作为低位)

- signID = (addressID << 12) | cryptoID
- addressID = 0x00007000 & signID >> 12
- cryptoID = signID & 0xffff8fff

#### 相关接口

> signID, cryptoID, addressID 转换接口

- EncodeSignID, 基于 cryptoID 和 addressID, 计算 signID
- ExtractAddressID, 基于 signID, 提取 addressID(主要用于交易的 fromAddr 计算)
- ExtractCryptoID, 基于 signID, 提取 cryptoID
- rpc, Chain.GetChainConfig, 获取节点默认的地址 ID 配置

#### 相关文档

- 加密模块介绍, chain/common/crypto/README.md
- 地址模块介绍, chain/common/address/README.md
