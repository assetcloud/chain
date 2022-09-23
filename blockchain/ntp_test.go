package blockchain

import (
	"testing"

	"github.com/assetcloud/chain/queue"
	"github.com/assetcloud/chain/types"
)

func TestCheckClockDrift(t *testing.T) {
	cfg := types.NewChainConfig(types.GetDefaultCfgstring())
	q := queue.New("channel")
	q.SetConfig(cfg)

	blockchain := &BlockChain{}
	blockchain.client = q.Client()
	blockchain.checkClockDrift()

	cfg.GetModuleConfig().NtpHosts = append(cfg.GetModuleConfig().NtpHosts, types.NtpHosts...)
	blockchain.checkClockDrift()
}
