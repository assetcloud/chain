// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.8
// +build go1.8

// Package main chain程序入口
package main

import (
	_ "github.com/assetcloud/chain/system"
	"github.com/assetcloud/chain/util/cli"
)

func main() {
	cli.RunChain("", "")
}
