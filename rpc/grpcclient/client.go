package grpcclient

import (
	"fmt"
	"sync"
	"time"

	"github.com/assetcloud/chain/common/log/log15"

	"github.com/assetcloud/chain/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// paraChainGrpcRecSize 平行链receive最大100M
const paraChainGrpcRecSize = 100 * 1024 * 1024

var mu sync.RWMutex

var defaultClient types.ChainClient

// GetDefaultMainClient get default client
func GetDefaultMainClient() types.ChainClient {
	mu.RLock()
	defer mu.RUnlock()
	return defaultClient
}

//NewMainChainClient 创建一个平行链的 主链 grpc chain 客户端
func NewMainChainClient(cfg *types.ChainConfig, grpcaddr string) (types.ChainClient, error) {
	mu.Lock()
	defer mu.Unlock()
	if grpcaddr == "" && defaultClient != nil {
		return defaultClient, nil
	}
	serverAddr := cfg.GetModuleConfig().RPC.ParaChain.MainChainGrpcAddr
	if grpcaddr != "" {
		serverAddr = grpcaddr
	}

	kp := keepalive.ClientParameters{
		Time:                time.Second * 5,
		Timeout:             time.Second * 20,
		PermitWithoutStream: true,
	}

	var conn *grpc.ClientConn
	var err error
	log15.Error("NewMainChainClient start+++++++++++++++++++++++++++++++++")
	if cfg.GetModuleConfig().RPC.ParaChain.UseGrpcLBSync {
		conn, err = grpc.Dial(NewSyncURL(serverAddr), grpc.WithInsecure(),
			grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(paraChainGrpcRecSize)),
			grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, SyncLbName)),
			grpc.WithKeepaliveParams(kp))
	} else {
		conn, err = grpc.Dial(NewMultipleURL(serverAddr), grpc.WithInsecure(),
			grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(paraChainGrpcRecSize)),
			grpc.WithKeepaliveParams(kp))
	}
	if err != nil {
		return nil, err
	}
	grpcClient := types.NewChainClient(conn)
	if grpcaddr == "" {
		defaultClient = grpcClient
	}
	return grpcClient, nil
}
