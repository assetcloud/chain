## 比特币脚本集成

### 集成方案

比特币脚本集成主要涉及两个绑定

#### 交易绑定

由于比特币脚本内部仍然是基于比特币交易，锁定脚本和解锁脚本对应的是比特币交易，
无法约束到 chain 交易，需要完成交易绑定，相关步骤如下：

1. 根据锁定脚本生成解锁脚本时，需要构造临时的比特币交易 btctx
2. btctx 只包含一个输入 UTXO，且该 UTXO 的索引哈希为 chain 交易的哈希
3. 这样基于 btctx 生成的解锁脚本就和 chain 交易相互绑定，保证了 chain 交易不被篡改

#### 交易发送方绑定

锁定脚本和解锁脚本智能约束交易是否正确，
在 chain 交易中，需要在签名结构中指定公钥，即对应交易发送方的公钥，
需要通过脚本限定交易发送方的地址

1. 在比特币脚本签名验证中，签名的数据需要包含锁定脚本和解锁脚本
2. chain 交易签名公钥设为 Hash256(锁定脚本)，即交易的发送方地址由锁定脚本生成
3. 类似于比特币 P2SH，采用比特币脚本验证方式，交易的发起地址都是基于锁定脚本哈希计算
4. 在验证时，即需要验证公钥是否和锁定脚本一一对应，即绑定了 chain 交易的发送方地址

### 相关脚本

介绍相关脚本实现方案

#### 多重签名

多重签名脚本可以通过接口调用直接构建，M <Public Key 1> <Public Key 2> … <Public Key N> N CHECKMULTISIGVERIFY

设有地址 A，B，C，在 chain 中完成 3:2 多重签名大概步骤如下

1. 构造锁定脚本 S1， 2 PubkeyA PubkeyB PubkeyC 3 CHECKMULTISIG
2. 生成 chain 多重签名地址 X = Pubkey2Addr(Hash256(S1))
3. 在 chain 中，在地址 X 中存入需要多重签名的资产
4. 提取 X 资产需提供解锁脚本，即对应的多重签名，任意 ABC 中的两个私钥的数据签名

#### 钱包找回

1. 构造延时存证交易，即 payload 为实际需要延时的交易，链上执行后需要记录延时交易的开始时刻

2. 基于比特元 CSV（check sequence verify）操作码， 构造锁定脚本，即限定 utxo 在相对时长后能花费

3. CSV 在比特币中需要基于一个开始时间点，在集成到 chain 中，使用 1 中构造的 none 延时交易执行时间点作为 CSV 的开始时间，并进行相对时长验证

4. 钱包找回原始地址 A，备份地址 B，即地址 A 私钥丢失时，可以使用备份地址 B 的私钥控制资产，但需要延时验证

5. 构建锁定脚本 S， IF <A's Pubkey> CHECKSIGVERIFY ELSE <1 day> CHECKSEQUENCEVERIFY DROP <B's Pubkey> CHECKSIGVERIFY ENDIF

> 基于锁定脚本 S 生成脚本地址 X，将需要找回控制的资产转入 X, A 的私钥可以控制 X 的资产，B 的私钥转移 X 的资产时，需要满足延时时长一天

#### 延时转账

延时转账可支持，绝对时刻和相对时长两种

##### 绝对时刻

1. 基于 CLTV（check lock time verify）操作

2. 地址 A 和 B，A 需要向 B 进行延时一小时转账，A 直接根据当前区块时间或高度计算出延时截止时刻 T

3. 构建锁定脚本 S， IF <A's Pubkey> CHECKSIGVERIFY ELSE <T> CHECKLOCKTIMEVERIFY DROP <B's Pubkey> CHECKSIGVERIFY ENDIF

> 根据 S 生成中间地址 X，A 首先向 X 进行相应资产转账, A 的私钥可以一直控制 X 的资产，而 B 的私钥当且仅当在时刻 T 后能控制 X 的资产

##### 相对时长

基于 none 延时交易，类似于钱包找回操作

1. 构造 none 合约延时交易，即 payload 为实际需要延时的交易，链上执行后需要记录延时交易的开始时刻

2. 基于比特元 CSV（check sequence verify）操作码， 构造锁定脚本，即限定 utxo 在相对时长后能花费

3. CSV 在比特币中需要基于一个开始时间点，在集成到 chain 中，使用 1 中构造的 none 延时交易执行时间点作为 CSV 的开始时间，并进行相对时长验证

4. 延时转账发起方 A，接收方 B，A 向 B 进行延时转账

5. 构建锁定脚本 S， IF <A's Pubkey> CHECKSIGVERIFY ELSE <1 hour> CHECKSEQUENCEVERIFY DROP <B's Pubkey> CHECKSIGVERIFY ENDIF

> 基于锁定脚本 S 生成脚本地址 X，将需要延时转账的资产转入 X, A 的私钥可以控制 X 的资产，B 的私钥转移 X 的资产时，需要满足延时时长一个小时
