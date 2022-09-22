package dht

import (
	_ "github.com/assetcloud/chain/system/p2p/dht/protocol/broadcast" //register init package
	_ "github.com/assetcloud/chain/system/p2p/dht/protocol/download"  //register init package
	_ "github.com/assetcloud/chain/system/p2p/dht/protocol/p2pstore"  //register init package
	_ "github.com/assetcloud/chain/system/p2p/dht/protocol/peer"      //register init package
)
