// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package main chain-cli程序入口
package main

import (

	// 这一步是必需的，目的时让插件源码有机会进行匿名注册
	"github.com/assetcloud/chain/cmd/cli/buildflags"
	_ "github.com/assetcloud/chain/system"
	"github.com/assetcloud/chain/util/cli"
)

func main() {
	if buildflags.RPCAddr == "" {
		buildflags.RPCAddr = "http://localhost:8801"
	}
	cli.Run(buildflags.RPCAddr, buildflags.ParaName, "")
}
