package peer

import (
	"testing"

	"github.com/assetcloud/chain/queue"
	"github.com/assetcloud/chain/system/p2p/dht/protocol"
	"github.com/assetcloud/chain/types"
	"github.com/stretchr/testify/require"
)

func Test_statistic(t *testing.T) {
	q := queue.New("test")
	p, cancel := initEnv(t, q)
	defer cancel()
	require.Equal(t, false, p.checkDone())
	remotePid := p.Host.Network().Conns()[0].RemotePeer()
	stream, err := p.Host.NewStream(p.Ctx, remotePid, statisticalInfo)
	require.Nil(t, err)
	defer protocol.CloseStream(stream)
	var resp types.Statistical
	err = protocol.ReadStream(&resp, stream)
	require.Nil(t, err)
	require.Equal(t, int32(0), resp.Nodeinfo.Outbounds)
}
