package grpcclient

import (
	"fmt"
	"sync"
	"time"

	"github.com/assetcloud/chain/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// paraChainGrpcRecSize 平行链receive最大100M
const paraChainGrpcRecSize = 100 * 1024 * 1024

var mu sync.Mutex

var defaultClient types.ChainClient

//NewMainChainClient 创建一个平行链的 主链 grpc chain 客户端
func NewMainChainClient(cfg *types.ChainConfig, grpcaddr string) (types.ChainClient, error) {
	mu.Lock()
	defer mu.Unlock()
	if grpcaddr == "" && defaultClient != nil {
		return defaultClient, nil
	}
	paraRemoteGrpcClient := types.Conf(cfg, "config.consensus.sub.para").GStr("ParaRemoteGrpcClient")
	if grpcaddr != "" {
		paraRemoteGrpcClient = grpcaddr
	}
	if paraRemoteGrpcClient == "" {
		paraRemoteGrpcClient = "127.0.0.1:8802"
	}
	kp := keepalive.ClientParameters{
		Time:                time.Second * 5,
		Timeout:             time.Second * 20,
		PermitWithoutStream: true,
	}

	var conn *grpc.ClientConn
	var err error
	useLBSync := types.Conf(cfg, "config.consensus.sub.para").IsEnable("useGrpcLBSync")
	if useLBSync {
		conn, err = grpc.Dial(NewSyncURL(paraRemoteGrpcClient), grpc.WithInsecure(),
			grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(paraChainGrpcRecSize)),
			grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, SyncLbName)),
			grpc.WithKeepaliveParams(kp))
	} else {
		conn, err = grpc.Dial(NewMultipleURL(paraRemoteGrpcClient), grpc.WithInsecure(),
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
