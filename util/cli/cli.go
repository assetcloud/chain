// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cli

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/assetcloud/chain/common/log"
	"github.com/assetcloud/chain/common/version"
	"github.com/assetcloud/chain/pluginmgr"
	"github.com/assetcloud/chain/rpc/jsonclient"
	rpctypes "github.com/assetcloud/chain/rpc/types"
	"github.com/assetcloud/chain/system/dapp/commands"
	"github.com/assetcloud/chain/types"
	"github.com/spf13/cobra"
)

// Run :
func Run(RPCAddr, ParaName, name string) {
	// cli 命令只打印错误级别到控制台
	log.SetLogLevel("error")
	configPath := ""
	for i, arg := range os.Args[:] {
		if arg == "--conf" && i+1 <= len(os.Args)-1 { // --conf chain.toml 可以配置读入cli配置文件路径
			configPath = os.Args[i+1]
			break
		}
		if strings.HasPrefix(arg, "--conf=") { // --conf="chain.toml"
			configPath = strings.TrimPrefix(arg, "--conf=")
			break
		}
	}
	if configPath == "" {
		if name == "" {
			configPath = "chain.toml"
		} else {
			configPath = name + ".toml"
		}
	}

	exist, _ := pathExists(configPath)
	var chainCfg *types.ChainConfig
	if exist {
		chainCfg = types.NewChainConfig(types.ReadFile(configPath))
	} else {
		cfgstring := types.GetDefaultCfgstring()
		if ParaName != "" {
			cfgstring = strings.Replace(cfgstring, "Title=\"local\"", fmt.Sprintf("Title=\"%s\"", ParaName), 1)
			cfgstring = strings.Replace(cfgstring, "FixTime=false", "CoinSymbol=\"para\"", 1)
		}
		chainCfg = types.NewChainConfig(cfgstring)
	}

	types.SetCliSysParam(chainCfg.GetTitle(), chainCfg)

	rootCmd := &cobra.Command{
		Use:     chainCfg.GetTitle() + "-cli",
		Short:   chainCfg.GetTitle() + " client tools",
		Version: fmt.Sprintf("%s %s", version.GetVersion(), version.BuildTime),
	}

	closeCmd := &cobra.Command{
		Use:   "close",
		Short: "Close " + chainCfg.GetTitle(),
		Run: func(cmd *cobra.Command, args []string) {
			rpcLaddr, err := cmd.Flags().GetString("rpc_laddr")
			if err != nil {
				panic(err)
			}
			//		rpc, _ := jsonrpc.NewJSONClient(rpcLaddr)
			//		rpc.Call("Chain.CloseQueue", nil, nil)
			var res rpctypes.Reply
			ctx := jsonclient.NewRPCCtx(rpcLaddr, "Chain.CloseQueue", nil, &res)
			ctx.Run()
		},
	}

	rootCmd.AddCommand(
		commands.CertCmd(),
		commands.AccountCmd(),
		commands.BlockCmd(),
		commands.CoinsCmd(),
		commands.ExecCmd(),
		commands.MempoolCmd(),
		commands.NetCmd(),
		commands.SeedCmd(),
		commands.StatCmd(),
		commands.TxCmd(),
		commands.WalletCmd(),
		commands.VersionCmd(),
		commands.SystemCmd(),
		commands.OneStepSendCmd(),
		commands.OneStepSendCertTxCmd(),
		commands.BlacklistCmd(),
		closeCmd,
		commands.AssetCmd(),
		commands.NoneCmd(),
		commands.BtcScriptCmd(),
	)

	//test tls is enable
	RPCAddr = testTLS(RPCAddr)
	pluginmgr.AddCmd(rootCmd)
	log.SetLogLevel("error")
	chainCfg.S("RPCAddr", RPCAddr)
	chainCfg.S("ParaName", ParaName)
	rootCmd.PersistentFlags().String("rpc_laddr", chainCfg.GStr("RPCAddr"), "http url")
	rootCmd.PersistentFlags().String("paraName", chainCfg.GStr("ParaName"), "parachain")
	rootCmd.PersistentFlags().String("title", chainCfg.GetTitle(), "get title name")
	rootCmd.PersistentFlags().MarkHidden("title")
	rootCmd.PersistentFlags().String("conf", "", "cli config")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func testTLS(RPCAddr string) string {
	rpcaddr := RPCAddr
	if !strings.HasPrefix(rpcaddr, "http://") {
		return RPCAddr
	}
	// if http://
	if rpcaddr[len(rpcaddr)-1] != '/' {
		rpcaddr += "/"
	}
	rpcaddr += "test"
	/* #nosec */
	resp, err := http.Get(rpcaddr)
	if err != nil {
		return "https://" + RPCAddr[7:]
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		return RPCAddr
	}
	return "https://" + RPCAddr[7:]
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
