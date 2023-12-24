// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found In the LICENSE file.

package timeline

import (
	"github.com/assetcloud/chain/queue"
	drivers "github.com/assetcloud/chain/system/mempool"
	"github.com/assetcloud/chain/types"
)

func init() {
	drivers.Reg("timeline", New)
}

// New 创建timeline cache 结构的 mempool
func New(cfg *types.Mempool, sub []byte) queue.Module {
	c := drivers.NewMempool(cfg)
	var subcfg drivers.SubConfig
	types.MustDecode(sub, &subcfg)
	if subcfg.PoolCacheSize == 0 {
		subcfg.PoolCacheSize = cfg.PoolCacheSize
	}
	if subcfg.ProperFee == 0 {
		subcfg.ProperFee = cfg.MinTxFeeRate
	}
	c.SetQueueCache(drivers.NewSimpleQueue(subcfg))
	return c
}
