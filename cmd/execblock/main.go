// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// package main 主要是执行已经同步好的区块链的某个区块
package main

import (
	"flag"
	"fmt"
	_ "net/http/pprof"
	"os/user"
	"path/filepath"

	"github.com/assetcloud/chain/blockchain"
	"github.com/assetcloud/chain/client"
	clog "github.com/assetcloud/chain/common/log"
	log "github.com/assetcloud/chain/common/log/log15"
	"github.com/assetcloud/chain/executor"
	"github.com/assetcloud/chain/queue"
	"github.com/assetcloud/chain/store"
	"github.com/assetcloud/chain/types"
	"github.com/assetcloud/chain/util"
)

var height = flag.Int64("height", 1, "exec block height")
var datadir = flag.String("datadir", "", "data dir of chain, include logs and datas")
var configPath = flag.String("f", "chain.toml", "configfile")

func resetDatadir(cfg *types.Config, datadir string) {
	// Check in case of paths like "/something/~/something/"
	if datadir[:2] == "~/" {
		usr, _ := user.Current()
		dir := usr.HomeDir
		datadir = filepath.Join(dir, datadir[2:])
	}
	log.Info("current user data dir is ", "dir", datadir)
	cfg.Log.LogFile = filepath.Join(datadir, cfg.Log.LogFile)
	cfg.BlockChain.DbPath = filepath.Join(datadir, cfg.BlockChain.DbPath)
	cfg.P2P.DbPath = filepath.Join(datadir, cfg.P2P.DbPath)
	cfg.Wallet.DbPath = filepath.Join(datadir, cfg.Wallet.DbPath)
	cfg.Store.DbPath = filepath.Join(datadir, cfg.Store.DbPath)
}

func initEnv() (queue.Queue, queue.Module, queue.Module) {
	cfg := types.NewChainConfig(types.ReadFile(*configPath))
	mcfg := cfg.GetModuleConfig()
	if *datadir != "" {
		resetDatadir(mcfg, *datadir)
	}
	mcfg.Consensus.Minerstart = false
	var q = queue.New("channel")
	q.SetConfig(cfg)
	chain := blockchain.New(cfg)
	chain.SetQueueClient(q.Client())
	exec := executor.New(cfg)
	exec.SetQueueClient(q.Client())
	cfg.SetMinFee(0)
	s := store.New(cfg)
	s.SetQueueClient(q.Client())
	return q, chain, s
}

func main() {
	clog.SetLogLevel("info")
	flag.Parse()
	q, chain, s := initEnv()
	defer s.Close()
	defer chain.Close()
	defer q.Close()
	qclient, err := client.New(q.Client(), nil)
	if err != nil {
		panic(err)
	}
	req := &types.ReqBlocks{Start: *height - 1, End: *height}
	blocks, err := qclient.GetBlocks(req)
	if err != nil {
		panic(err)
	}
	log.Info("execblock", "block height", *height)
	prevState := blocks.Items[0].Block.StateHash
	block := blocks.Items[1].Block
	receipt, err := util.ExecTx(q.Client(), prevState, block)
	if err != nil {
		panic(err)
	}
	for i, r := range receipt.GetReceipts() {
		println("=======================")
		println("tx index ", i)
		for j, kv := range r.GetKV() {
			fmt.Println("\tKV:", j, kv)
		}
		for k, l := range r.GetLogs() {
			logType := types.LoadLog(block.Txs[i].Execer, int64(l.Ty))
			lTy := "unkownType"
			var logIns interface{}
			if logType != nil {
				logIns, err = logType.Decode(l.GetLog())
				if err != nil {
					panic(err)
				}
				lTy = logType.Name()
			}
			fmt.Printf("\tLog:%d %s->%v\n", k, lTy, logIns)
		}
	}
}
