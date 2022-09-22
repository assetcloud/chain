package admin

import (
	"fmt"

	"github.com/assetcloud/chain/client"
	"github.com/assetcloud/chain/queue"
	rpcclient "github.com/assetcloud/chain/rpc/client"
	etypes "github.com/assetcloud/chain/rpc/ethrpc/types"
	"github.com/assetcloud/chain/types"
	ctypes "github.com/assetcloud/chain/types"
	"github.com/ethereum/go-ethereum/common"
)

type adminHandler struct {
	cli rpcclient.ChannelClient
	cfg *ctypes.Chain33Config
}

//NewAdminAPI create a admin api
func NewAdminAPI(cfg *ctypes.Chain33Config, c queue.Client, api client.QueueProtocolAPI) interface{} {
	p := &adminHandler{}
	p.cli.Init(c, api)
	p.cfg = cfg
	return p
}

//Peers admin_peers
func (p *adminHandler) Peers() ([]*etypes.Peer, error) {
	var in = types.P2PGetPeerReq{}
	reply, err := p.cli.PeerInfo(&in)
	if err != nil {
		return nil, err
	}

	var peerlist []*etypes.Peer
	for _, peer := range reply.Peers {
		var pr etypes.Peer
		pr.ID = peer.GetName()
		pr.Self = peer.GetSelf()
		pr.NetWork = &etypes.Network{
			RemoteAddress: fmt.Sprintf("%v:%v", peer.GetAddr(), peer.GetPort()),
		}

		pr.Protocols = &etypes.Protocols{
			EthProto: &etypes.EthProto{
				Version:    peer.GetVersion(),
				Difficulty: peer.GetHeader().Difficulty,
				Head:       common.Bytes2Hex(peer.GetHeader().GetHash()),
			},
		}
		if pr.Self {
			//ip4/ip/tcp/port/p2p/id
			pr.Encode = fmt.Sprintf("/ip4/%v/tcp/%v/p2p/%v", peer.GetAddr(), peer.GetPort(), pr.ID)
			pr.ListenAddr = fmt.Sprintf("%v:%v", peer.GetAddr(), peer.GetPort())
			pr.Ports = &etypes.Ports{
				Listener:  peer.GetPort(),
				Discovery: peer.GetPort(),
			}
		}

		peerlist = append(peerlist, &pr)

	}
	return peerlist, nil
}

//Datadir admin_datadir
func (p *adminHandler) Datadir() (string, error) {
	mcfg := p.cfg.GetModuleConfig()
	dbpath := mcfg.BlockChain.DbPath
	return dbpath, nil
}

//NodeInfo admin_nodeInfo
func (p *adminHandler) NodeInfo() (*etypes.Peer, error) {

	peers, err := p.Peers()
	if err != nil {
		return nil, err
	}

	for _, peer := range peers {
		if peer.Self {
			peer.Self = false //去掉self字段
			return peer, nil
		}
	}

	return nil, nil
}
