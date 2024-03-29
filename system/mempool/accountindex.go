package mempool

import (
	"github.com/assetcloud/chain/common/listmap"
	"github.com/assetcloud/chain/types"
)

// AccountTxIndex 账户和交易索引
type AccountTxIndex struct {
	maxperaccount int
	accMap        map[string]*listmap.ListMap
}

// NewAccountTxIndex 创建一个新的索引
func NewAccountTxIndex(maxperaccount int) *AccountTxIndex {
	return &AccountTxIndex{
		maxperaccount: maxperaccount,
		accMap:        make(map[string]*listmap.ListMap),
	}
}

// TxNumOfAccount 返回账户在Mempool中交易数量
func (cache *AccountTxIndex) TxNumOfAccount(addr string) int {
	if _, ok := cache.accMap[addr]; ok {
		return cache.accMap[addr].Size()
	}
	return 0
}

// GetAccTxs 用来获取对应账户地址（列表）中的全部交易详细信息
func (cache *AccountTxIndex) GetAccTxs(addrs *types.ReqAddrs) *types.TransactionDetails {
	res := &types.TransactionDetails{}
	for _, addr := range addrs.Addrs {
		if value, ok := cache.accMap[addr]; ok {
			value.Walk(func(val interface{}) bool {
				v := val.(*types.Transaction)
				txAmount, err := v.Amount()
				if err != nil {
					txAmount = 0
				}
				res.Txs = append(res.Txs,
					&types.TransactionDetail{
						Tx:         v,
						Amount:     txAmount,
						Fromaddr:   addr,
						ActionName: v.ActionName(),
					})
				return true
			})
		}
	}
	return res
}

// Remove 根据交易哈希删除对应账户的对应交易
func (cache *AccountTxIndex) Remove(tx *types.Transaction, txHash string) {
	addr := tx.From()
	if lm, ok := cache.accMap[addr]; ok {
		lm.Remove(txHash)
		if lm.Size() == 0 {
			delete(cache.accMap, addr)
		}
	}
}

// Push push transaction to AccountTxIndex
func (cache *AccountTxIndex) Push(tx *types.Transaction, txHash string) error {
	addr := tx.From()
	_, ok := cache.accMap[addr]
	if !ok {
		cache.accMap[addr] = listmap.New()
	}
	if cache.accMap[addr].Size() >= cache.maxperaccount {
		return types.ErrManyTx
	}
	cache.accMap[addr].Push(txHash, tx)
	return nil
}

// CanPush 是否可以push 进 account index
func (cache *AccountTxIndex) CanPush(tx *types.Transaction) bool {
	if item, ok := cache.accMap[tx.From()]; ok {
		return item.Size() < cache.maxperaccount
	}
	return true
}
