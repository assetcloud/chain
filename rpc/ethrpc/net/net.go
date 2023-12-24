package net

import (
	"strconv"

	"github.com/assetcloud/chain/system/crypto/secp256k1eth"

	"github.com/assetcloud/chain/client"
	"github.com/assetcloud/chain/queue"
	rpcclient "github.com/assetcloud/chain/rpc/client"
	"github.com/assetcloud/chain/types"
	ctypes "github.com/assetcloud/chain/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type netHandler struct {
	cli rpcclient.ChannelClient
	cfg *ctypes.ChainConfig
}

// NewNetAPI create a net  api
func NewNetAPI(cfg *ctypes.ChainConfig, c queue.Client, api client.QueueProtocolAPI) interface{} {
	p := &netHandler{}
	p.cli.Init(c, api)
	p.cfg = cfg
	return p
}

// PeerCount net_peerCount
func (n *netHandler) PeerCount() (string, error) {

	var in = types.P2PGetPeerReq{}
	reply, err := n.cli.PeerInfo(&in)
	if err != nil {
		return "0x0", err
	}

	numPeers := len(reply.Peers)
	return hexutil.EncodeUint64(uint64(numPeers)), nil
}

// Listening net_listening
func (n *netHandler) Listening() (bool, error) {
	return true, nil
}

// Version net_version
func (n *netHandler) Version() (string, error) {
	return strconv.FormatInt(secp256k1eth.GetEvmChainID(), 10), nil
}
