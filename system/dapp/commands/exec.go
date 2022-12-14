// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package commands

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	commandtypes "github.com/assetcloud/chain/system/dapp/commands/types"
	"github.com/pkg/errors"

	"github.com/assetcloud/chain/common/address"
	"github.com/assetcloud/chain/types"
	"github.com/spf13/cobra"
)

// ExecCmd exec command
func ExecCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "exec",
		Short: "Executor operation",
		Args:  cobra.MinimumNArgs(1),
	}

	cmd.AddCommand(
		GetExecAddrCmd(),
		UserDataCmd(),
	)

	return cmd
}

// GetExecAddrCmd  get address of an execer
func GetExecAddrCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "addr",
		Short: "Get address of executor",
		Run:   getAddrByExec,
	}
	addGetAddrFlags(cmd)
	return cmd
}

func addGetAddrFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("exec", "e", "", getExecuterNameString())
	cmd.MarkFlagRequired("exec")
}

func getAddrByExec(cmd *cobra.Command, args []string) {
	execer, _ := cmd.Flags().GetString("exec")
	if ok := types.IsAllowExecName([]byte(execer), []byte(execer)); !ok {
		fmt.Println(types.ErrExecNameNotAllow.Error())
		return
	}
	rpcLaddr, _ := cmd.Flags().GetString("rpc_laddr")
	cfg, err := commandtypes.GetChainConfig(rpcLaddr)
	if err != nil {
		fmt.Fprintln(os.Stderr, errors.Wrapf(err, "GetChainConfig"))
		return
	}

	execAddr, err := commandtypes.GetExecAddr(execer, cfg.DefaultAddressID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println(execAddr)
}

// UserDataCmd  create user data
func UserDataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "userdata",
		Short: "Write data to user defined executor",
		Run:   addUserData,
	}
	addUserDataFlags(cmd)
	return cmd
}
func addUserDataFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("exec", "e", "", `executor name must contain prefix "user"`)
	cmd.MarkFlagRequired("exec")
	cmd.Flags().StringP("topic", "t", "", `add #topic# before data, if empty add nothing`)
	cmd.Flags().StringP("data", "d", "", `payload data that users want to write`)
	cmd.MarkFlagRequired("data")
}

func addUserData(cmd *cobra.Command, args []string) {
	title, _ := cmd.Flags().GetString("title")
	cfg := types.GetCliSysParam(title)
	execer, err := cmd.Flags().GetString("exec")
	if err != nil {
		fmt.Println(err)
		return
	}
	if !strings.HasPrefix(execer, "user.") {
		fmt.Println(`user defined executor should start with "user."`)
		return
	}
	if len(execer) > 50 {
		fmt.Println("executor name too long")
		return
	}
	addrResult := address.ExecAddress(execer)
	topic, _ := cmd.Flags().GetString("topic")
	data, _ := cmd.Flags().GetString("data")
	if topic != "" {
		data = "#" + topic + "#" + data
	}
	if data == "" {
		fmt.Println("write empty data")
		return
	}
	tx := &types.Transaction{
		Execer:  []byte(execer),
		Payload: []byte(data),
		To:      addrResult,
	}
	tx.Fee, err = tx.GetRealFee(cfg.GetMinTxFeeRate())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	tx.Nonce = random.Int63()
	tx.ChainID = cfg.GetChainID()
	//tx.Sign(int32(wallet.SignType), privKey)
	txHex := types.Encode(tx)
	fmt.Println(hex.EncodeToString(txHex))
}
