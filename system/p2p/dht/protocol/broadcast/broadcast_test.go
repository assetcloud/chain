// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package broadcast

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	net "github.com/assetcloud/chain/system/p2p/dht/extension"
	"github.com/libp2p/go-libp2p"
	core "github.com/libp2p/go-libp2p/core"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/multiformats/go-multiaddr"

	"github.com/assetcloud/chain/client"
	commlog "github.com/assetcloud/chain/common/log"
	"github.com/assetcloud/chain/p2p"
	"github.com/assetcloud/chain/queue"
	prototypes "github.com/assetcloud/chain/system/p2p/dht/protocol"
	p2pty "github.com/assetcloud/chain/system/p2p/dht/types"
	"github.com/assetcloud/chain/types"
	"github.com/stretchr/testify/assert"
)

func init() {
	commlog.SetLogLevel("error")
}

func newHost(port int32) core.Host {
	priv, _, _ := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, rand.Reader)
	m, err := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port))
	if err != nil {
		return nil
	}

	host, err := libp2p.New(libp2p.ListenAddrs(m), libp2p.Identity(priv))

	if err != nil {
		panic(err)
	}

	return host
}

func newTestEnv(q queue.Queue) (*prototypes.P2PEnv, context.CancelFunc) {
	cfg := types.NewChainConfig(types.ReadFile("../../../../../cmd/chain/chain.test.toml"))
	q.SetConfig(cfg)
	go q.Start()

	mgr := p2p.NewP2PMgr(cfg)
	mgr.Client = q.Client()
	mgr.SysAPI, _ = client.New(mgr.Client, nil)

	subCfg := &p2pty.P2PSubConfig{}
	types.MustDecode(cfg.GetSubConfig().P2P[p2pty.DHTTypeName], subCfg)
	env := &prototypes.P2PEnv{
		ChainCfg:        cfg,
		QueueClient:     q.Client(),
		Host:            newHost(13902),
		ConnManager:     nil,
		PeerInfoManager: nil,
		P2PManager:      mgr,
		SubConfig:       subCfg,
		Ctx:             context.Background(),
	}
	ctx, cancel := context.WithCancel(context.Background())
	env.Ctx = ctx
	var err error
	env.Pubsub, err = net.NewPubSub(env.Ctx, env.Host, &p2pty.PubSubConfig{})
	if err != nil {
		panic("new pubsub err" + err.Error())
	}
	return env, cancel
}

func newTestProtocolWithQueue(q queue.Queue) (*broadcastProtocol, context.CancelFunc) {
	env, cancel := newTestEnv(q)
	prototypes.ClearEventHandler()
	p := &broadcastProtocol{syncStatus: true}
	p.init(env)
	return p, cancel
}

func newTestProtocol() (*broadcastProtocol, context.CancelFunc) {

	q := queue.New("test")
	return newTestProtocolWithQueue(q)
}

func TestBroadCastProtocol_InitProtocol(t *testing.T) {

	protocol, cancel := newTestProtocol()
	defer cancel()
	assert.Equal(t, defaultMinLtBlockSize*1024, protocol.cfg.MinLtBlockSize)
	assert.Equal(t, defaultLtBlockTimeout, int(protocol.cfg.LtBlockPendTimeout))
}

func TestBroadcastSend(t *testing.T) {
	protocol, cancel := newTestProtocol()
	defer cancel()
	msgs := []*queue.Message{
		protocol.QueueClient.NewMessage("p2p", types.EventTxBroadcast, &types.Transaction{}),
		protocol.QueueClient.NewMessage("p2p", types.EventBlockBroadcast, &types.Block{}),
		protocol.QueueClient.NewMessage("p2p", types.EventAddBlock, &types.Block{}),
		protocol.QueueClient.NewMessage("p2p", types.EventBlockBroadcast, &types.Block{Txs: []*types.Transaction{tx1, tx2}}),
	}
	protocol.cfg.MinLtBlockSize = 0
	for _, msg := range msgs {
		protocol.handleBroadcastSend(msg)
	}
	_, ok := protocol.txFilter.Get(hex.EncodeToString((&types.Transaction{}).Hash()))
	assert.True(t, ok)
	_, ok = protocol.blockFilter.Get(hex.EncodeToString((&types.Block{}).Hash(protocol.ChainCfg)))
	assert.True(t, ok)
	_, ok = protocol.blockFilter.Get(hex.EncodeToString((&types.Block{Txs: []*types.Transaction{tx1, tx2}}).Hash(protocol.ChainCfg)))
	assert.True(t, ok)
}

func TestBroadCastReceive(t *testing.T) {

	p, cancel := newTestProtocol()
	defer cancel()
	pid := p.Host.ID()
	peerTopic := p.getPeerTopic(pid)
	msgs := []subscribeMsg{
		{value: &types.Transaction{}, topic: psTxTopic},
		{value: &types.Block{}, topic: psBlockTopic},
		{value: &types.LightBlock{}, topic: psLtBlockTopic},
		{value: &types.PeerPubSubMsg{MsgID: blockReqMsgID, ProtoMsg: types.Encode(&types.ReqInt{})}, topic: peerTopic},
		{value: &types.PeerPubSubMsg{MsgID: blockRespMsgID}, topic: peerTopic},
		{value: &types.PeerPubSubMsg{}, topic: peerTopic},
	}
	for _, msg := range msgs {
		msg.receiveFrom = pid
		msg.publisher = pid
		p.handleBroadcastReceive(msg)
	}
	require.Equal(t, 0, p.ltB.pendBlockList.Len())
	require.Equal(t, 0, p.ltB.blockRequestList.Len())
}
