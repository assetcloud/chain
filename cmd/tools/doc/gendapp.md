# chain gendapp

根据定义的合约 protobuf 原型文件，自动生成 chain dapp 基本代码

### 编译

```
//本地存在chain代码，该步骤可省略
$ go get github.com/assetcloud/chain
//编译chain tools
$ go build -i -o $GOPATH/bin/chain-tool github.com/assetcloud/chain/cmd/tools
```

### 使用

```
//查看命令使用方法
$ chain-tool gendapp --help
Usage:
  tools gendapp [flags]

Flags:
  -h, --help            help for gendapp
  -n, --name string     dapp name
  -o, --output string   go package for output (default github.com/assetcloud/plugin/plugin/dapp/)
  -p, --proto string    dapp protobuf file path
```

- -n 指定合约名字，不能含有空格和特殊字符
- -p 指定合约的 protobuf 文件
- -o 生成代码的输出目录路径，此处是 go 包路径，及相对于$GOPATH/src的路径，
默认为官方项目路径（$GOPATH/src/github.com/assetcloud/plugin/plugin/dapp/)

举例:

```
// 默认路径生成名为demo的合约代码
$ chain-tool gendapp -n demo -p ./demo.proto

// 指定输出包路径
$ chain-tool gendapp -n demo -p ./demo.proto -o github.com/assetcloud/chain/plugin/dapp/

```

### proto 规范

- 定义合约交易行为结构，采用**oneof value**形式，且名称必须为**NameAction**格式，
  如 demo 合约，定义 echo 和 hello 两种交易行为

```proto
message DemoAction {
    oneof value {
        DemoHello hello = 1;
        DemoEcho  echo  = 2;
    }
    int32 ty = 3;
}
```

- package name 设为 types，适配后续生成目录结构

```proto
package types;
```

- 定义 service，直接以合约名作为名称

```proto
service demo {
}
```

### 代码

##### 目录结构，以 demo 合约为例

```
demo
├── cmd             //包含官方ci集成相关脚本
│   ├── build.sh
│   └── Makefile
├── commands        //合约客户端模块
│   └── commands.go
├── executor        //执行器模块
│   ├── demo.go
│   ├── exec_del_local.go
│   ├── exec.go
│   ├── exec_local.go
│   └── kv.go
├── plugin.go
├── proto           //proto文件及生成pb.go命令
│   ├── create_protobuf.sh
│   ├── demo.proto
│   └── Makefile
├── rpc             //rpc模块
│   ├── rpc.go
│   └── types.go
└── types           //类型模块
    └── demo.go

```

##### 生成 pb.go 文件

pb.go 文件基于 protobuf 提供的 proto-gen-go 插件生成，这里 protobuf 的版本必须和 chain 引用的保持一致，
具体可以查看 chain 项目 go.mod 文件，github.com/golang/protobuf 库的版本

```
//进入到上述proto目录执行相关脚本，将会在types目录下生成对应pb.go文件
$ cd proto && make
```

##### 后续开发

在生成代码基础上，需要实现交易创建，执行，及所需 rpc 服务<br/>
初次开发可以参考官方简单计算器合约
[开发步骤](https://github.com/assetcloud/chain/blob/master/cmd/tools/doc/gencalculator.md)
