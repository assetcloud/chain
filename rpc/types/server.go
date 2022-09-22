// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"net/rpc"

	"github.com/assetcloud/chain/account"
	"github.com/assetcloud/chain/client"
	"github.com/assetcloud/chain/queue"
	"github.com/assetcloud/chain/types"
	"google.golang.org/grpc"
)

// RPCServer interface
type RPCServer interface {
	GetQueueClient() queue.Client
	GRPC() *grpc.Server
	JRPC() *rpc.Server
}

// ChannelClient interface
type ChannelClient struct {
	client.QueueProtocolAPI
	accountdb *account.DB
	grpc      interface{}
	jrpc      interface{}
}

// Init init function
func (c *ChannelClient) Init(name string, s RPCServer, jrpc, grpc interface{}) {
	if c.QueueProtocolAPI == nil {
		c.QueueProtocolAPI, _ = client.New(s.GetQueueClient(), nil)
	}
	if jrpc != nil {
		s.JRPC().RegisterName(name, jrpc)
	}
	c.grpc = grpc
	c.jrpc = jrpc
	types.AssertConfig(c.QueueProtocolAPI)
	c.accountdb = account.NewCoinsAccount(c.GetConfig())
}

// GetCoinsAccountDB  return accountdb
func (c *ChannelClient) GetCoinsAccountDB() *account.DB {
	return c.accountdb
}
