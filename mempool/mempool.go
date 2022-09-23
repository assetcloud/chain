package mempool

import (
	"github.com/assetcloud/chain/queue"
	"github.com/assetcloud/chain/system/mempool"
	"github.com/assetcloud/chain/types"
)

// New new mempool queue module
func New(cfg *types.ChainConfig) queue.Module {
	mcfg := cfg.GetModuleConfig().Mempool
	sub := cfg.GetSubConfig().Mempool
	con, err := mempool.Load(mcfg.Name)
	if err != nil {
		panic("Unsupported mempool type:" + mcfg.Name + " " + err.Error())
	}
	subcfg, ok := sub[mcfg.Name]
	if !ok {
		subcfg = nil
	}
	obj := con(mcfg, subcfg)
	return obj
}
