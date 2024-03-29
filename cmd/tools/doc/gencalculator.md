# calculator generate

基于 gendapp 自动生成合约命令，介绍合约的完整开发步骤

### 简介

calculator 合约支持在区块链上进行整数加减乘除交易操作，同时方便演示
开发，记录运算符参与运算的次数，并提供查询接口

### 编写合约 proto

```proto
syntax = "proto3";

package calculator;
// calculator 合约交易行为总类型
message CalculatorAction {
    oneof value {
        Add      add = 1;
        Subtract sub = 2;
        Multiply mul = 3;
        Divide   div = 4;
    }
    int32 ty = 5;
}

message Add {
    int32 summand = 1; //被加数
    int32 addend  = 2; //加数
}
message AddLog {
    int32 sum = 1; //和
}

message Subtract {
    int32 minuend    = 1; //被减数
    int32 subtrahend = 2; //减数
}
message SubLog {
    int32 remainder = 1; //差
}

message Multiply {
    int32 faciend    = 1; //被乘数
    int32 multiplier = 2; //乘数
}
message MultiplyLog {
    int32 product = 1; //积
}

message Divide {
    int32 dividend = 1; //被除数
    int32 divisor  = 2; //除数
}
message DivideLog {
    int32 quotient = 1; //商
    int32 remain   = 2; //余数
}

message ReqQueryCalcCount {
    string action = 1;
}
message ReplyQueryCalcCount {
    int32 count = 1;
}

service calculator {
    rpc QueryCalcCount(ReqQueryCalcCount) returns (ReplyQueryCalcCount) {}
}
```

主要有以下几个部分：

- 定义交易行为总结构，CalculatorAction，包含加减乘除
- 分别定义涉及的交易行为结构， Add，Sub 等
- 定义交易涉及到的日志结构，每种运算除均有对应结果日志
- 如果需要 grpc 服务，定义 service 结构，如本例增加了查询次数的 rpc
- 定义查询中涉及的 request，reply 结构

### 代码生成

##### 生成基本代码

> 使用 chain-tool，工具使用参考[文档](https://github.com/assetcloud/chain/blob/master/cmd/tools/doc/gendapp.md)

```
//本例默认将calculator生成至官方plugin项目dapp目录下
$ cd $GOPATH/src/github.com/assetcloud/chain/cmd/tools && go build -o tool
$ ./tool gendapp -n calculator -p doc/calculator.proto
$ cd $GOPATH/src/github.com/assetcloud/plugin/plugin/dapp/calculator && ls
```

##### 生成 pb.go 文件

pb.go 文件基于 protobuf 提供的 proto-gen-go 插件生成，这里 protobuf 的版本必须和 chain 引用的保持一致，
具体可以查看 chain 项目 go.mod 文件，github.com/golang/protobuf 库的版本

```
//进入生成合约的目录
$ cd $GOPATH/src/github.com/assetcloud/plugin/plugin/dapp/calculator
//执行脚本生成calculator.pb.go
$ cd proto && make
```

### 后续开发

以下将以模块为顺序，依次介绍

#### types 类型模块

此目录统一归纳合约类型相关的代码

##### 交易的 action 和 log(types/calculator.go)

> 每一种交易通常有交易请求(action），交易执行回执(log)，
> 目前框架要求合约开发者自定义 aciton 和 log 的 id 及 name，
> 已经自动生成了这些常量，可以根据需要修改

```go
// action类型id和name，可以自定义修改
const (
	TyAddAction= iota + 100
	TySubAction
	TyMulAction
	TyDivAction

	NameAddAction = "Add"
	NameSubAction = "Sub"
	NameMulAction = "Mul"
	NameDivAction = "Div"
)

// log类型id值
const (
	TyUnknownLog = iota + 100
	TyAddLog
	TySubLog
	TyMulLog
	TyDivLog
)
```

> 开发者还需要提供 name 和 id 的映射结构，其中 actionMap 已自动生成,
> 交易 log 结构由开发者自由定义，这里 logMap 需要将对应结构按格式填充，
> 如本例中加减乘除都有对应的 log 类型（也可以采用一个通用结构对应多个交易回执），依次按照格式填入即可

```go

    //定义action的name和id
	actionMap = map[string]int32{
		NameAddAction: TyAddAction,
		NameSubAction: TySubAction,
		NameMulAction: TyMulAction,
		NameDivAction: TyDivAction,
	}
	//定义log的id和具体log类型及名称，填入具体自定义log类型
	logMap = map[int64]*types.LogInfo{
		TyAddLog: {Ty:reflect.TypeOf(AddLog{}), Name: "AddLog"},
		TySubLog: {Ty:reflect.TypeOf(SubLog{}), Name: "SubLog"},
		TyMulLog: {Ty:reflect.TypeOf(MultiplyLog{}), Name: "MultiplyLog"},
		TyDivLog: {Ty:reflect.TypeOf(DivideLog{}), Name: "DivideLog"},
	}
```

#### executor 执行模块

此目录归纳了交易执行逻辑实现代码

##### 实现 CheckTx 接口(executor/calculator.go)

> CheckTx 即检查交易合法性，隶属于框架 Driver 接口，将在交易执行前被框架调用，
> 本例简单实现除法非零检测

```go
func (*calculator) CheckTx(tx *types.Transaction, index int) error {

    action := &calculatortypes.CalculatorAction{}
	err := types.Decode(tx.GetPayload(), action)
	if err != nil {
		elog.Error("CheckTx", "DecodeActionErr", err)
		return types.ErrDecode
	}
	//这里只做除法除数零值检查
	if action.Ty == calculatortypes.TyDivAction {
		div, ok := action.Value.(*calculatortypes.CalculatorAction_Div)
		if !ok {
			return types.ErrTypeAsset
		}
		if div.Div.Divisor == 0 {	//除数不能为零
			elog.Error("CheckTx", "Err", "ZeroDivisor")
			return types.ErrInvalidParam
		}
	}
	return nil
}
```

##### KV 常量(executor/kv.go)

> 目前合约进行存取框架 KV 数据库(stateDB 或 localDB)时，
> 其 Key 的前缀必须满足框架要求规范，已经以常量形式自动生成在代码中，
> 开发者在构造数据 key 时，需要以此为前缀

```
var (
	//KeyPrefixStateDB state db key必须前缀
	KeyPrefixStateDB = "mavl-calculator-"
	//KeyPrefixLocalDB local db的key必须前缀
	KeyPrefixLocalDB = "LODB-calculator-"
)
```

##### 实现 Exec 类接口(executor/exec.go)

> Exec 类接口是交易链上执行的函数，实现交易执行的业务逻辑，
> 数据上链也是此部分完成(生成 stateDB KV 对)，以及生成交易日志，以 Add 交易为例

```go
func (c *calculator) Exec_Add(payload *ptypes.Add, tx *types.Transaction, index int) (*types.Receipt, error) {
	var receipt *types.Receipt
	sum := payload.Addend + payload.Summand
	addLog := &ptypes.AddLog{Sum: sum}
	logs := []*types.ReceiptLog{{Ty:ptypes.TyAddLog, Log: types.Encode(addLog)}}
	key := fmt.Sprintf("%s-%s-formula", KeyPrefixStateDB, tx.Hash())
	val := fmt.Sprintf("%d+%d=%d", payload.Summand, payload.Addend, sum)
	receipt = &types.Receipt{
		Ty: types.ExecOk,
		KV: []*types.KeyValue{{Key:[]byte(key), Value:[]byte(val)}},
		Logs: logs,
	}
	return receipt, nil
}
```

##### 实现 ExecLocal 类接口(executor/exec_local.go)

> ExecLocal 类接口是交易执行成功后本地执行，
> 主要目的是将辅助性数据进行 localDB 存取,方便前端查询，
> 以 Add 为例，在 localDB 中存入加法运算的次数，在函数最后需要调用 addAutoRollBack 接口，以适配框架 localdb 自动回滚功能

```go
func (c *calculator) ExecLocal_Add(payload *ptypes.Add, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	var dbSet *types.LocalDBSet
	var countInfo calculatortypes.ReplyQueryCalcCount
	localKey := []byte(fmt.Sprintf("%s-CalcCount-Add", KeyPrefixLocalDB))
	oldVal, err := c.GetLocalDB().Get(localKey)
	//此处需要注意，目前db接口，获取key未找到记录，返回空时候也带一个notFound错误，需要特殊处理，而不是直接返回错误
	if err != nil && err != types.ErrNotFound{
		return nil, err
	}
	err = types.Decode(oldVal, &countInfo)
	if err != nil {
		elog.Error("execLocalAdd", "DecodeErr", err)
		return nil, types.ErrDecode
	}
	countInfo.Count++
	dbSet = &types.LocalDBSet{KV: []*types.KeyValue{{Key:localKey, Value:types.Encode(&countInfo)}}}
	//封装kv，适配框架自动回滚，这部分代码已经自动生成
    return c.addAutoRollBack(tx, dbSet.KV), nil
}
```

##### 实现 ExecDelLocal 类接口(executor/exec_del_local.go)

> ExecDelLocal 类接口可以理解为 ExecLocal 的逆过程，在区块回退时候被调用，生成代码已支持自动回滚，无需实现

##### 实现 Query 类接口(executor/query.go)

> Query 类接口主要实现查询相关业务逻辑，如访问合约数据库，
> Query 类接口需要满足框架规范(固定格式函数名称和签名)，才能被框架注册和使用，
> 具体调用方法将在 rpc 模块介绍，本例实现查询运算符计算次数的接口

```go
//函数名称，Query_+实际方法名格式，返回值为protobuf Message结构
func (c *calculator) Query_CalcCount(in *ptypes.ReqQueryCalcCount) (types.Message, error) {

	var countInfo ptypes.ReplyQueryCalcCount
	localKey := []byte(fmt.Sprintf("%s-CalcCount-%s", KeyPrefixLocalDB, in.Action))
	oldVal, err := c.GetLocalDB().Get(localKey)
	if err != nil && err != types.ErrNotFound{
		return nil, err
	}
	err = types.Decode(oldVal, &countInfo)
	if err != nil {
		elog.Error("execLocalAdd", "DecodeErr", err)
		return nil, err
	}
	return &countInfo, nil
}
```

#### rpc 模块

此目录归纳了 rpc 相关类型和具体调用服务端实现的代码

##### 类型(rpc/types.go)

> 定义了 rpc 相关结构和初始化，此部分代码已经自动生成

```go
// 实现grpc的service接口
type channelClient struct { //实现grpc接口的类
	rpctypes.ChannelClient
}
// Jrpc 实现json rpc调用实例
type Jrpc struct {  //实现json rpc接口的类
	cli *channelClient
}
```

##### grpc 接口(rpc/rpc.go)

> grpc 即实现 proto 文件中 service 声明的 rpc 接口，本例中即查询计算次数的 rpc。
> 此处通过框架 Query 接口，间接调用之前实现的 Query_CalcCount 接口

```go
func (c *channelClient)QueryCalcCount(ctx context.Context, in *ptypes.ReqQueryCalcCount) (*ptypes.ReplyQueryCalcCount, error) {

	msg, err :=  c.Query(ptypes.CalculatorX, "CalcCount", in)
	if err != nil {
		return nil, err
	}
	if reply, ok := msg.(*ptypes.ReplyQueryCalcCount); ok {
		return reply, nil
	}
	return nil, types.ErrTypeAsset
}
```

##### json rpc 接口

> json rpc 主要给前端相关平台产品调用，本例为查询计算次数接口

```go
func (j *Jrpc)QueryCalcCount(in *ptypes.ReqQueryCalcCount, result *interface{}) error {

    //此处直接调用内部的grpc接口
	reply, err := j.cli.QueryCalcCount(context.Background(), in)
	if err != nil {
		return err
	}
	*result = *reply
	return nil
}
```

##### rpc 说明

> 对于构造交易和 query 类接口可以通过 chain 框架的 rpc 去调用，
> 分别是 Chain.CreateTransaction 和 Chain.Query，上述代码只是示例如何开发 rpc 接口，
> 实际使用中，只需要实现 query 接口，并通过框架 rpc 调用，也可以根据需求封装 rpc 接口，在 commands 模块将会介绍如何调用框架 rpc

#### commands 命令行模块

如果需要支持命令行交互式访问区块节点，开发者需要实现具体合约的命令，
框架的命令行基于 cobra 开源库

##### import 路径(commands/commands.go)

> 涉及框架基础库使用，包括相关类型和网络组件

```go
import (
	"github.com/assetcloud/chain/rpc/jsonclient"
	"github.com/assetcloud/chain/types"
	"github.com/spf13/cobra"

	rpctypes "github.com/assetcloud/chain/rpc/types"
	calculatortypes "github.com/assetcloud/plugin/plugin/dapp/calculator/types"
)
```

##### 创建交易命令(commands/commands.go)

> 前端输入相关参数，调用 rpc 实现创建原始交易的功能

```go
func createAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "add",
		Short:"create add calc tx",
		Run: createAdd,
	}
	cmd.Flags().Int32P("summand", "s", 0, "summand integer number")
	cmd.Flags().Int32P("addend", "a", 0, "addend integer number")
	cmd.MarkFlagRequired("summand")
	cmd.MarkFlagRequired("addend")
	return cmd
}


func createAdd(cmd *cobra.Command, args []string) {
	rpcLaddr, _ := cmd.Flags().GetString("rpc_laddr")
	summand, _ := cmd.Flags().GetInt32("summand")
	addend, _ := cmd.Flags().GetInt32("addend")

	req := ptypes.Add{
		Summand: summand,
		Addend:  addend,
	}
	chainReq := rpctypes.CreateTxIn{
		Execer:     ptypes.CalculatorX,
		ActionName: ptypes.NameAddAction,
		Payload:    types.MustPBToJSON(&req),
	}
	var res string
	//调用框架CreateTransaction接口构建原始交易
	ctx := jsonclient.NewRPCCtx(rpcLaddr, "Chain.CreateTransaction", chainReq, &res)
	ctx.RunWithoutMarshal()
}
```

##### 查询计算次数(commands/commands.go)

```go
func queryCalcCountCmd() *cobra.Command {

 	cmd := &cobra.Command{
 		Use:   "query_count",
 		Short: "query calculator count",
 		Run:   queryCalcCount,
 	}
 	cmd.Flags().StringP("action", "a", "", "calc action name[Add | Sub | Mul | Div]")
 	cmd.MarkFlagRequired("action")

 	return cmd
 }

 func queryCalcCount(cmd *cobra.Command, args []string) {

 	rpcLaddr, _ := cmd.Flags().GetString("rpc_laddr")
 	action, _ := cmd.Flags().GetString("action")
 	req := ptypes.ReqQueryCalcCount{
 		Action: action,
 	}
 	chainReq := &rpctypes.Query4Jrpc{
 		Execer:   ptypes.CalculatorX,
 		FuncName: "CalcCount",
 		Payload:  types.MustPBToJSON(&req),
 	}
 	var res interface{}
 	res = &calculatortypes.ReplyQueryCalcCount{}
 	//调用框架Query rpc接口, 通过框架调用，需要指定query对应的函数名称，具体参数见Query4Jrpc结构
 	ctx := jsonclient.NewRPCCtx(rpcLaddr, "Chain.Query", chainReq, &res)
 	//调用合约内部rpc接口, 注意合约自定义的rpc接口是以合约名称作为rpc服务，这里为calculator
 	//ctx := jsonclient.NewRPCCtx(rpcLaddr, "calculator.QueryCalcCount", req, &res)
 	ctx.Run()
 }
```

##### 添加到主命令(commands/commands.go)

```go
func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "calculator",
		Short: "calculator command",
		Args:  cobra.MinimumNArgs(1),
	}
	cmd.AddCommand(
		//add sub command
		createAddCmd(),
		queryCalcCountCmd(),
	)
	return cmd
}
```

#### 合约集成

新增合约需要显示初始化

##### 初始化（dapp/init/init.go)

> 需要在此文件 import 目录，新增 calculator 包导入

```go
import (
 	_ "github.com/assetcloud/plugin/plugin/dapp/calculator" //init calculator
)
```

##### 编译

> 直接通过官方 makefile 文件

```
$ cd $GOPATH/src/github.com/assetcloud/plugin && make
```

#### 测试

##### 单元测试

为合约代码增加必要的单元测试，提高测试覆盖

##### 集成测试

编译后可以运行节点，进行钱包相关配置，即可发送合约交易进行功能性测试，本例相关命令行

```bash
# 通过curl方式调用rpc接口构建Add原始交易
curl -kd '{"method":"Chain.CreateTransaction", "params":[{"execer":"calculator", "actionName":"Add", "payload":{"summand":1,"addend":1}}]}' http://localhost:8801
# 通过chain-cli构建Add原始交易
./chain-cli calculator add -a 1 -s 1

# queryCount接口类似
curl -kd '{"method":"calculator.QueryCalcCount", "params":[{"action":"Add"}]}' http://localhost:8801
./chain-cli calculator query_count -a Add
```

#### 进阶

##### 计算器

基于 [本例代码](https://github.com/bysomeone/plugin/tree/dapp-example-calculator) 实现减法等交易行为

##### 其他例子

官方 [plugin 项目](https://github.com/assetcloud/plugin) 提供了丰富的插件，可以参考学习
