// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/assetcloud/chain/common"
	"github.com/assetcloud/chain/common/address"
	"github.com/assetcloud/chain/rpc/jsonclient"
	rpctypes "github.com/assetcloud/chain/rpc/types"
	"github.com/assetcloud/chain/types"
	"github.com/assetcloud/chain/util"
	"github.com/assetcloud/chain/util/testnode"
	"github.com/stretchr/testify/assert"

	_ "github.com/assetcloud/chain/system"
)

func getRPCClient(t *testing.T, mocker *testnode.ChainMock) *jsonclient.JSONClient {
	jrpcClient := mocker.GetJSONC()
	assert.NotNil(t, jrpcClient)
	return jrpcClient
}

func TestErrLog(t *testing.T) {
	// 启动RPCmocker
	mocker := testnode.New("--free--", nil)
	defer mocker.Close()
	cfg := mocker.GetClient().GetConfig()
	mocker.Listen()
	jrpcClient := getRPCClient(t, mocker)
	gen := mocker.GetGenesisKey()
	//发送交易到区块链
	addr1, key1 := util.Genaddress()
	addr2, _ := util.Genaddress()
	tx1 := util.CreateCoinsTx(cfg, gen, addr1, 1*types.DefaultCoinPrecision)
	mocker.GetAPI().SendTx(tx1)
	mocker.WaitHeight(1)

	tx11 := util.CreateCoinsTx(cfg, key1, addr2, 6*int64(1e7))
	reply, err := mocker.GetAPI().SendTx(tx11)
	assert.Nil(t, err)
	assert.Equal(t, reply.GetMsg(), tx11.Hash())
	tx12 := util.CreateCoinsTx(cfg, key1, addr2, 6*int64(1e7))
	reply, err = mocker.GetAPI().SendTx(tx12)
	assert.Nil(t, err)
	assert.Equal(t, reply.GetMsg(), tx12.Hash())
	mocker.WaitTx(reply.GetMsg())
	var testResult rpctypes.TransactionDetail
	req := rpctypes.QueryParm{
		Hash: common.ToHex(tx12.Hash()),
	}
	//query transaction
	err = jrpcClient.Call("Chain.QueryTransaction", req, &testResult)
	assert.Nil(t, err)
	assert.Equal(t, string(testResult.Receipt.Logs[0].Log), `"ErrNoBalance"`)
	assert.Equal(t, "0.6000", testResult.Tx.AmountFmt)
	assert.Equal(t, int64(6e7), testResult.Tx.Amount)
}

func getTx(t *testing.T, hex string) *types.Transaction {
	data, err := common.FromHex(hex)
	assert.Nil(t, err)
	var tx types.Transaction
	err = types.Decode(data, &tx)
	assert.Nil(t, err)
	return &tx
}

func TestSendToExec(t *testing.T) {
	mocker := testnode.New("--free--", nil)
	defer mocker.Close()
	mocker.Listen()
	jrpcClient := getRPCClient(t, mocker)
	//1. 调用createrawtransaction 创建交易
	req := &rpctypes.CreateTx{
		To:          address.ExecAddress("user.f3d"),
		Amount:      10,
		Fee:         1,
		Note:        "12312",
		IsWithdraw:  false,
		IsToken:     false,
		TokenSymbol: "",
		ExecName:    "user.f3d",
	}
	var res string
	err := jrpcClient.Call("Chain.CreateRawTransaction", req, &res)
	assert.Nil(t, err)
	gen := mocker.GetGenesisKey()
	tx := getTx(t, res)
	tx.Sign(types.SECP256K1, gen)
	reply, err := mocker.GetAPI().SendTx(tx)
	assert.Nil(t, err)
	_, err = mocker.WaitTx(reply.GetMsg())
	assert.Nil(t, err)
	block := mocker.GetLastBlock()
	balance := mocker.GetExecAccount(block.StateHash, "user.f3d", mocker.GetGenesisAddress()).Balance
	assert.Equal(t, int64(10), balance)
}

func TestGetAllExecBalance(t *testing.T) {
	mocker := testnode.New("--free--", nil)
	defer mocker.Close()
	mocker.Listen()
	jrpcClient := getRPCClient(t, mocker)

	addr := "38BRY193Wvy9MkdqMjmuaYeUHnJaFjUxMP"
	req := types.ReqAddr{Addr: addr}
	var res rpctypes.AllExecBalance
	err := jrpcClient.Call("Chain.GetAllExecBalance", &req, &res)
	assert.Nil(t, err)
	assert.Equal(t, addr, res.Addr)
	assert.Nil(t, res.ExecAccount)
	assert.Equal(t, 0, len(res.ExecAccount))
}

func TestCreateTransactionUserWrite(t *testing.T) {
	mocker := testnode.New("--free--", nil)
	defer mocker.Close()
	mocker.Listen()
	jrpcClient := getRPCClient(t, mocker)
	req := &rpctypes.CreateTxIn{
		Execer:     "user.write",
		ActionName: "write",
		Payload:    []byte(`{"key":"value"}`),
	}
	var res string
	err := jrpcClient.Call("Chain.CreateTransaction", req, &res)
	assert.Nil(t, err)
	tx := getTx(t, res)
	assert.NotNil(t, tx)
	fmt.Println(string(tx.Payload))
	assert.Nil(t, err)
	assert.Equal(t, `{"key":"value"}`, string(tx.Payload))
}

func TestExprieCreateNoBalanceTransaction(t *testing.T) {
	mocker := testnode.New("--free--", nil)
	defer mocker.Close()
	mocker.Listen()
	jrpcClient := getRPCClient(t, mocker)
	req := &rpctypes.CreateTxIn{
		Execer:     "user.write",
		ActionName: "write",
		Payload:    []byte(`{"key":"value"}`),
	}
	var res string
	err := jrpcClient.Call("Chain.CreateTransaction", req, &res)
	assert.Nil(t, err)
	gen := mocker.GetGenesisKey().Bytes()
	req2 := &types.NoBalanceTx{
		TxHex:   res,
		Privkey: common.ToHex(gen),
		Expire:  "300s",
	}
	var groupres string
	err = jrpcClient.Call("Chain.CreateNoBalanceTransaction", req2, &groupres)
	assert.Nil(t, err)

	txByteData, err := common.FromHex(groupres)
	assert.Nil(t, err)
	var tx types.Transaction
	err = types.Decode(txByteData, &tx)
	assert.Nil(t, err)
	txgroup, err := tx.GetTxGroup()
	assert.Nil(t, err)
	assert.True(t, txgroup.GetTxs()[0].GetExpire() > 0)
}

func TestExprieSignRawTx(t *testing.T) {
	mocker := testnode.New("--free--", nil)
	cfg := mocker.GetClient().GetConfig()
	defer mocker.Close()
	mocker.Listen()
	jrpcClient := getRPCClient(t, mocker)
	req := &rpctypes.CreateTxIn{
		Execer:     "user.write",
		ActionName: "write",
		Payload:    []byte(`{"key":"value"}`),
	}
	var res string
	err := jrpcClient.Call("Chain.CreateTransaction", req, &res)
	assert.Nil(t, err)

	txNone := &types.Transaction{Execer: []byte(cfg.ExecName(types.NoneX)), Payload: []byte("no-fee-transaction")}
	txNone.To = address.ExecAddress(string(txNone.Execer))
	txNone, err = types.FormatTx(cfg, cfg.ExecName(types.NoneX), txNone)
	assert.NoError(t, err)
	assert.Nil(t, err)
	gen := mocker.GetGenesisKey().Bytes()
	req2 := &types.CreateTransactionGroup{
		Txs: []string{hex.EncodeToString(types.Encode(txNone)), res},
	}
	var groupres string
	err = jrpcClient.Call("Chain.CreateRawTxGroup", req2, &groupres)
	assert.Nil(t, err)

	txByteData, err := common.FromHex(groupres)
	assert.Nil(t, err)
	var tx types.Transaction
	err = types.Decode(txByteData, &tx)
	assert.Nil(t, err)
	req3 := &types.ReqSignRawTx{
		TxHex:   common.ToHex(types.Encode(&tx)),
		Privkey: common.ToHex(gen),
		Expire:  "300s",
	}
	var signgrouptx string
	err = jrpcClient.Call("Chain.SignRawTx", req3, &signgrouptx)
	assert.Nil(t, err)

	txByteData, err = common.FromHex(signgrouptx)
	assert.Nil(t, err)
	var tx2 types.Transaction
	err = types.Decode(txByteData, &tx2)
	assert.Nil(t, err)
	txgroup2, err := tx2.GetTxGroup()
	assert.Nil(t, err)
	assert.True(t, txgroup2.GetTxs()[0].GetExpire() > 0)
}
