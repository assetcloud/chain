#!/usr/bin/env bash
# shellcheck disable=SC2128

MAIN_HTTP=""
IS_PARA=false
CASE_ERR=""

#color
RED='\033[1;31m'
GRE='\033[1;32m'
NOC='\033[0m'

# $2=0 means true, other false
echo_rst() {
    if [ "$2" -eq 0 ]; then
        echo -e "${GRE}$1 ok${NOC}"
    elif [ "$2" -eq 2 ]; then
        echo -e "${GRE}$1 not support${NOC}"
    else
        echo -e "${RED}$1 fail${NOC}"
        echo -e "${RED}$3 ${NOC}"
        CASE_ERR="err"
    fi
}

http_req() {
    #  echo "#$4 request: $1"
    body=$(curl -ksd "$1" "$2")
    #echo "#response: $body"
    ok=$(echo "$body" | jq -r "$3")
    [ "$ok" == true ]
    rst=$?
    echo_rst "$4" "$rst" "$body"
}

chain_lock() {
    http_req '{"method":"Chain.Lock","params":[]}' ${MAIN_HTTP} ".result.isOK" "$FUNCNAME"
}

chain_unlock() {
    http_req '{"method":"Chain.UnLock","params":[{"passwd":"1314fuzamei","timeout":0}]}' ${MAIN_HTTP} ".result.isOK" "$FUNCNAME"
}

chain_WalletTxList() {
    req='{"method":"Chain.WalletTxList", "params":[{"fromTx":"", "count":2, "direction":1}]}'
    http_req "$req" ${MAIN_HTTP} '(.error|not) and (.result.txDetails|length == 2)' "$FUNCNAME"
}

chain_ImportPrivkey() {
    req='{"method":"Chain.ImportPrivkey", "params":[{"privkey":"0x88b2fb90411935872f0501dd13345aba19b5fac9b00eb0dddd7df977d4d5477e", "label":"testimportkey"}]}'
    http_req "$req" ${MAIN_HTTP} '(.error|not) and (.result.label=="testimportkey") and (.result.acc.addr == "1D9xKRnLvV2zMtSxSx33ow1GF4pcbLcNRt")' "$FUNCNAME"
}

chain_DumpPrivkey() {
    req='{"method":"Chain.DumpPrivkey", "params":[{"data":"1D9xKRnLvV2zMtSxSx33ow1GF4pcbLcNRt"}]}'
    http_req "$req" ${MAIN_HTTP} '(.error|not) and (.result.data=="0x88b2fb90411935872f0501dd13345aba19b5fac9b00eb0dddd7df977d4d5477e")' "$FUNCNAME"
}

chain_DumpPrivkeysFile() {
    req='{"method":"Chain.DumpPrivkeysFile", "params":[{"fileName":"PrivkeysFile","passwd":"123456"}]}'
    http_req "$req" ${MAIN_HTTP} '(.error|not) and .result.isOK' "$FUNCNAME"
}

chain_ImportPrivkeysFile() {
    req='{"method":"Chain.ImportPrivkeysFile", "params":[{"fileName":"PrivkeysFile","passwd":"123456"}]}'
    http_req "$req" ${MAIN_HTTP} '(.error|not) and .result.isOK' "$FUNCNAME"
    # rm -rf ./PrivkeysFile
}

chain_SendToAddress() {
    req='{"method":"Chain.SendToAddress", "params":[{"from":"12qyocayNF7Lv6C9qW4avxs2E7U41fKSfv","to":"1D9xKRnLvV2zMtSxSx33ow1GF4pcbLcNRt", "amount":100000000, "note":"test\n"}]}'
    http_req "$req" ${MAIN_HTTP} '(.error|not) and (.result.hash|length==66)' "$FUNCNAME"
}

chain_SetTxFee() {
    http_req '{"method":"Chain.SetTxFee", "params":[{"amount":100000}]}' ${MAIN_HTTP} '(.error|not) and .result.isOK' "$FUNCNAME"
}

chain_SetLabl() {
    req='{"method":"Chain.SetLabl", "params":[{"addr":"1D9xKRnLvV2zMtSxSx33ow1GF4pcbLcNRt", "label":"updatetestimport"}]}'
    http_req "$req" ${MAIN_HTTP} '(.error|not) and (.result.label=="updatetestimport") and (.result.acc.addr == "1D9xKRnLvV2zMtSxSx33ow1GF4pcbLcNRt")' "$FUNCNAME"
}

chain_GetPeerInfo() {
    if [ "$IS_PARA" == true ]; then
        echo_rst "$FUNCNAME" 2
    else
        resok='(.error|not) and (.result.peers|length >= 1) and (.result.peers[0] | [has("addr", "port", "name", "mempoolSize", "self", "header"), true] | unique | length == 1)'
        http_req '{"method":"Chain.GetPeerInfo", "params":[{}]}' ${MAIN_HTTP} "$resok" "$FUNCNAME"
    fi
}

chain_GetHeaders() {
    resok='(.error|not) and (.result.items|length == 2) and (.result.items[0] | [has("version","parentHash", "txHash", "stateHash", "height", "blockTime", "txCount", "hash", "difficulty"),true] | unique | length == 1 )'
    http_req '{"method":"Chain.GetHeaders", "params":[{"start":1, "end":2, "isDetail":true}]}' ${MAIN_HTTP} "$resok" "$FUNCNAME"
}

chain_GetLastMemPool() {
    if [ "$IS_PARA" == true ]; then
        echo_rst "$FUNCNAME" 2
    else
        http_req '{"method":"Chain.GetLastMemPool", "params":[{}]}' ${MAIN_HTTP} '(.error|not) and (.result.txs|length >= 0)' "$FUNCNAME"
    fi
}

chain_GetProperFee() {
    http_req '{"method":"Chain.GetProperFee", "params":[{}]}' ${MAIN_HTTP} '(.error|not) and (.result.properFee > 10000)' "$FUNCNAME"
}

chain_GetBlockOverview() {
    hash=$(curl -ksd '{"method":"Chain.GetHeaders", "params":[{"start":1, "end":1, "isDetail":true}]}' ${MAIN_HTTP} | jq '.result.items[0].hash')
    req='{"method":"Chain.GetBlockOverview", "params":[{"hash":'"$hash"'}]}'
    http_req "$req" ${MAIN_HTTP} '(.error|not) and (.result| [has("head", "txCount", "txHashes"), true]|unique|length == 1) and (.result.txCount == (.result.txHashes|length))' "$FUNCNAME"
}

chain_GetAddrOverview() {
    req='{"method":"Chain.GetAddrOverview", "params":[{"addr":"12qyocayNF7Lv6C9qW4avxs2E7U41fKSfv"}]}'
    http_req "$req" ${MAIN_HTTP} '(.error|not) and (.result|[has("reciver", "balance", "txCount"), true]|unique|length == 1)' "$FUNCNAME"
}

chain_SetPasswd() {
    http_req '{"method":"Chain.SetPasswd", "params":[{"oldPass":"1314fuzamei", "newPass":"1314fuzamei"}]}' ${MAIN_HTTP} '(.error|not) and .result.isOK' "$FUNCNAME"
}

chain_MergeBalance() {
    http_req '{"method":"Chain.MergeBalance", "params":[{"to":"12qyocayNF7Lv6C9qW4avxs2E7U41fKSfv"}]}' ${MAIN_HTTP} '(.error|not) and (.result.hashes|length > 0)' "$FUNCNAME"
}

chain_QueryTotalFee() {
    local height=1
    hash=$(curl -ksd '{"method":"Chain.GetBlockHash","params":[{"height":'$height'}]}' ${MAIN_HTTP} | jq -r ".result.hash")
    if [ -z "$hash" ]; then
        echo "hash is null"
        echo_rst "$FUNCNAME" 1
    fi
    prefixhash_base64=$(echo -n "TotalFeeKey:" | base64)
    blockhash_base64=$(echo -n "$hash" | cut -d " " -f 1 | xxd -r -p | base64)
    base64_hash="$prefixhash_base64$blockhash_base64"

    req='{"method":"Chain.QueryTotalFee","params":[{"keys":["'$base64_hash'"]}]}'
    http_req "$req" ${MAIN_HTTP} '(.result.txCount >= 0)' "$FUNCNAME"
}

chain_GetNetInfo() {
    if [ "$IS_PARA" == true ]; then
        echo_rst "$FUNCNAME" 2
    else
        http_req '{"method":"Chain.GetNetInfo", "params":[]}' ${MAIN_HTTP} '(.result.externalAddr| length > 0)' "$FUNCNAME"
    fi
}

chain_GetFatalFailure() {
    http_req '{"method":"Chain.GetFatalFailure", "params":[]}' ${MAIN_HTTP} '(.error|not) and (.result | 0)' "$FUNCNAME"
}

chain_DecodeRawTransaction() {
    tx="0a05636f696e73122c18010a281080c2d72f222131477444795771577233553637656a7663776d333867396e7a6e7a434b58434b7120a08d0630a696c0b3f78dd9ec083a2131477444795771577233553637656a7663776d333867396e7a6e7a434b58434b71"
    http_req '{"method":"Chain.DecodeRawTransaction", "params":[{"txHex":"'$tx'"}]}' ${MAIN_HTTP} '(.result.txs[0].execer == "coins")' "$FUNCNAME"
}

chain_GetTimeStatus() {
    r1=$(curl -ksd '{"method":"Chain.GetTimeStatus","params":[]}' ${MAIN_HTTP} | jq -r ".result.localTime")
    if [ -z "$r1" ]; then
        curl -ksd '{"method":"Chain.GetTimeStatus","params":[]}' ${MAIN_HTTP}
    fi
    [ -n "$r1" ]
    echo_rst "$FUNCNAME" "$?"
}

chain_GetLastBlockSequence() {
    if [ "$IS_PARA" == true ]; then
        echo_rst "$FUNCNAME" 2
    else
        http_req '{"method":"Chain.GetLastBlockSequence","params":[]}' ${MAIN_HTTP} ".result >= 0" "$FUNCNAME"
    fi
}

chain_GetBlockSequences() {
    if [ "$IS_PARA" == true ]; then
        echo_rst "$FUNCNAME" 2
    else
        http_req '{"method":"Chain.GetBlockSequences","params":[{"start":1,"end":3,"isDetail":true}]}' ${MAIN_HTTP} ".result.blkseqInfos|length==3" "$FUNCNAME"
    fi
}

chain_GetBlockByHashes() {
    if [ "$IS_PARA" == true ]; then
        geneis=$(curl -ksd '{"method":"Chain.GetBlockHash", "params":[{"height":0}]}' "${MAIN_HTTP}" | jq -r '(.result.hash)')
        req='{"method":"Chain.GetBlockByHashes","params":[{"hashes":["'"${geneis}"'"]}]}'
        http_req "$req" ${MAIN_HTTP} '(.result.items[0].block.parentHash == "0x0000000000000000000000000000000000000000000000000000000000000000")' "$FUNCNAME"
    else
        hash0=$(curl -ksd '{"method":"Chain.GetBlockSequences","params":[{"start":1,"end":3,"isDetail":true}]}' ${MAIN_HTTP} | jq -r ".result.blkseqInfos[0].hash")
        hash1=$(curl -ksd '{"method":"Chain.GetBlockSequences","params":[{"start":1,"end":3,"isDetail":true}]}' ${MAIN_HTTP} | jq -r ".result.blkseqInfos[1].hash")
        hash2=$(curl -ksd '{"method":"Chain.GetBlockSequences","params":[{"start":1,"end":3,"isDetail":true}]}' ${MAIN_HTTP} | jq -r ".result.blkseqInfos[2].hash")

        req='{"method":"Chain.GetBlockByHashes","params":[{"hashes":["'"$hash1"'","'"$hash2"'"]}]}'
        resok='(.result.items[0].block.parentHash == "'"$hash0"'") and (.result.items[1].block.parentHash =="'"$hash1"'")'
        http_req "$req" ${MAIN_HTTP} "$resok" "$FUNCNAME"
    fi
}

chain_ConvertExectoAddr() {
    http_req '{"method":"Chain.ConvertExectoAddr","params":[{"execname":"coins"}]}' ${MAIN_HTTP} '(.result == "1GaHYpWmqAJsqRwrpoNcB8VvgKtSwjcHqt")' "$FUNCNAME"
}

chain_GetExecBalance() {
    local height=6802
    statehash=$(curl -ksd '{"method":"Chain.GetBlocks","params":[{"start":'$height',"end":'$height',"isDetail":false}]}' ${MAIN_HTTP} | jq -r ".result.items[0].block.stateHash")
    state_base64=$(echo -n "$statehash" | cut -d " " -f 1 | xxd -r -p | base64)
    addr="12qyocayNF7Lv6C9qW4avxs2E7U41fKSfv"
    addr_base64=$(echo -n "$addr" | base64)

    req='{"method":"Chain.GetExecBalance","params":[{"symbol":"bty","stateHash":"'$state_base64'","addr":"'$addr_base64'","execer":"coins","count":100}]}'
    http_req "$req" ${MAIN_HTTP} "(.error|not)" "$FUNCNAME"
}

chain_AddPushSubscribe() {
    http_req '{"method":"Chain.AddPushSubscribe","params":[{"name":"test","url":"http://test","encode":"json"}]}' ${MAIN_HTTP} '(.result.isOk == true)' "$FUNCNAME"
}

chain_ListPushes() {
    http_req '{"method":"Chain.ListPushes","params":[]}' ${MAIN_HTTP} ' (.result.pushes[0].name == "test")' "$FUNCNAME"
}

chain_GetPushSeqLastNum() {
    http_req '{"method":"Chain.GetPushSeqLastNum","params":[{"data":"test-another"}]}' ${MAIN_HTTP} '(.result.data == -1)' "$FUNCNAME"
}

chain_GetCoinSymbol() {
    symbol="bty"
    if [ "$IS_PARA" == true ]; then
        symbol="para"
    fi

    resok='(.result.data == "'"$symbol"'")'
    http_req '{"method":"Chain.GetCoinSymbol","params":[]}' ${MAIN_HTTP} "$resok" "$FUNCNAME"
}

chain_GetChainID() {
    resok='(.result.chainID == 0)'
    http_req '{"method":"Chain.GetChainID","params":[]}' ${MAIN_HTTP} "$resok" "$FUNCNAME"
}

chain_GetHexTxByHash() {
    #先获取一笔交易
    reHash=$(curl -ksd '{"jsonrpc":"2.0","id":2,"method":"Chain.GetTxByAddr","params":[{"addr":"14KEKbYtKKQm4wMthSK9J4La4nAiidGozt","flag":0,"count":1,"direction":0,"height":-1,"index":0}]}' -H 'content-type:text/plain;' ${MAIN_HTTP} | jq -r '.result.txInfos[0].hash')
    #查询交易
    req='{"method":"Chain.GetHexTxByHash","params":[{"hash":"'"$reHash"'","upgrade":false}]}'
    http_req "$req" ${MAIN_HTTP} '(.error|not) and (.result != null)' "$FUNCNAME"
}

chain_QueryTransaction() {
    #先获取一笔交易
    reHash=$(curl -ksd '{"jsonrpc":"2.0","id":2,"method":"Chain.GetTxByAddr","params":[{"addr":"14KEKbYtKKQm4wMthSK9J4La4nAiidGozt","flag":0,"count":1,"direction":0,"height":-1,"index":0}]}' -H 'content-type:text/plain;' ${MAIN_HTTP} | jq -r '.result.txInfos[0].hash')
    #查询交易
    req='{"method":"Chain.QueryTransaction","params":[{"hash":"'"$reHash"'","upgrade":false}]}'
    http_req "$req" ${MAIN_HTTP} '(.error|not) and (.result.receipt.tyName == "ExecOk") and (.result.height >= 0) and (.result.index >= 0) and (.result.amount >= 0)' "$FUNCNAME"
}

chain_GetBlocks() {
    http_req '{"method":"Chain.GetBlocks","params":[{"start":1,"end":2}]}' ${MAIN_HTTP} '(.result.items[1].block.height == 2)' "$FUNCNAME"
}

chain_GetLastHeader() {
    resok='(.error|not) and (.result.height >= 0) and (.result | [has("version","parentHash", "txHash", "stateHash", "height", "blockTime", "txCount", "hash", "difficulty"),true] | unique | length == 1)'
    http_req '{"method":"Chain.GetLastHeader","params":[{}]}' ${MAIN_HTTP} "$resok" "$FUNCNAME"
}

chain_GetTxByAddr() {
    req='{"method":"Chain.GetTxByAddr","params":[{"addr":"14KEKbYtKKQm4wMthSK9J4La4nAiidGozt","flag":0,"count":1,"direction":0,"height":-1,"index":0}]}'
    resok='(.error|not) and (.result.txInfos[0].index >= 0) and (.result.txInfos[0] | [has("hash", "height", "index", "assets"),true] | unique | length == 1)'
    http_req "$req" ${MAIN_HTTP} "$resok" "$FUNCNAME"
}

chain_GetTxByHashes() {
    req='{"method":"Chain.GetTxByHashes","params":[{"hashes":["0x8040109d3859827d0f0c80ce91cc4ec80c496c45250f5e5755064b6da60842ab","0x501b910fd85d13d1ab7d776bce41a462f27c4bfeceb561dc47f0a11b10f452e4"]}]}'
    http_req "$req" ${MAIN_HTTP} '(.error|not) and (.result.txs|length == 2)' "$FUNCNAME"
}

chain_GetMempool() {
    if [ "$IS_PARA" == true ]; then
        echo_rst "$FUNCNAME" 2
    else
        http_req '{"method":"Chain.GetMempool","params":[{}]}' ${MAIN_HTTP} '(.error|not) and (.result.txs|length >= 0)' "$FUNCNAME"
    fi
}

chain_GetAccountsV2() {
    http_req '{"method":"Chain.GetAccountsV2","params":[{}]}' ${MAIN_HTTP} '(.error|not) and (.result.wallets|length >= 0)' "$FUNCNAME"
}

chain_GetAccounts() {
    http_req '{"method":"Chain.GetAccounts","params":[{}]}' ${MAIN_HTTP} '(.error|not) and (.result.wallets|length >= 0)' "$FUNCNAME"
}

chain_NewAccount() {
    http_req '{"method":"Chain.NewAccount","params":[{"label":"test169"}]}' ${MAIN_HTTP} '(.error|not) and (.result.label == "test169") and (.result.acc | [has("addr"),true] | unique | length == 1)' "$FUNCNAME"
}

# hyb
chain_CreateRawTransaction() {
    local to="1EDDghAtgBsamrNEtNmYdQzC1QEhLkr87t"
    local amount=10000000
    tx=$(curl -ksd '{"method":"Chain.CreateRawTransaction","params":[{"to":"'$to'","amount":'$amount'}]}' ${MAIN_HTTP} | jq -r ".result")

    req='{"method":"Chain.DecodeRawTransaction","params":[{"txHex":"'"$tx"'"}]}'
    resok='(.result.txs[0].payload.transfer.amount == "'$amount'") and (.result.txs[0].to == "'$to'")'
    http_req "$req" ${MAIN_HTTP} "$resok" "$FUNCNAME"
}

chain_CreateTransaction() {
    local to="1EDDghAtgBsamrNEtNmYdQzC1QEhLkr87t"
    local amount=10000000
    local exec=""

    if [ "$IS_PARA" == true ]; then
        exec="user.p.para.coins"
    else
        exec="coins"
    fi

    tx=$(curl -ksd '{"method":"Chain.CreateTransaction","params":[{"execer":"'$exec'","actionName":"Transfer","payload":{"to":"'$to'", "amount":'$amount'}}]}' ${MAIN_HTTP} | jq -r ".result")

    req='{"method":"Chain.DecodeRawTransaction","params":[{"txHex":"'"$tx"'"}]}'
    resok='(.result.txs[0].payload.transfer.amount == "'$amount'") and (.result.txs[0].payload.transfer.to == "'$to'")'
    http_req "$req" ${MAIN_HTTP} "$resok" "$FUNCNAME"
}

chain_ReWriteRawTx() {
    local fee=1000000
    local tx1="0a05636f696e73122d18010a291080ade20422223145444467684174674273616d724e45744e6d5964517a43315145684c6b7238377420a08d0630f6db93c0e0d3f1ff5e3a223145444467684174674273616d724e45744e6d5964517a43315145684c6b72383774"
    tx=$(curl -ksd '{"method":"Chain.ReWriteRawTx","params":[{"expire":"120s","fee":'$fee',"tx":"'$tx1'"}]}' ${MAIN_HTTP} | jq -r ".result")

    req='{"method":"Chain.DecodeRawTransaction","params":[{"txHex":"'"$tx"'"}]}'
    http_req "$req" ${MAIN_HTTP} '(.error|not) and (.result.txs[0].execer == "coins") and (.result.txs[0].to == "1EDDghAtgBsamrNEtNmYdQzC1QEhLkr87t") and (.result.txs[0].fee == '$fee')' "$FUNCNAME"
}

chain_CreateRawTxGroup() {
    local to="1DNaSDRG9RD19s59meAoeN4a2F6RH97fSo"
    local exec="user.write"
    local groupCount=2
    tx1="0a0a757365722e7772697465121d236d642368616b6468676f7177656a6872676f716a676f6a71776c6a6720a08d0630a0b7b1b1dda2f4c5743a2231444e615344524739524431397335396d65416f654e34613246365248393766536f"
    tx2="0a0a757365722e7772697465121d236d642368616b6468676f7177656a6872676f716a676f6a71776c6a6720a08d0630c5838f94e2f49acb4b3a2231444e615344524739524431397335396d65416f654e34613246365248393766536f"
    tx=$(curl -ksd '{"method":"Chain.CreateRawTxGroup","params":[{"txs":["'$tx1'","'$tx2'"]}]}' ${MAIN_HTTP} | jq -r ".result")

    req='{"method":"Chain.DecodeRawTransaction","params":[{"txHex":"'"$tx"'"}]}'
    resok='(.error|not) and (.result.txs[0].execer == "'$exec'") and (.result.txs[0].to == "'$to'") and (.result.txs[0].groupCount == '$groupCount') and (.result.txs[1].execer == "'$exec'") and (.result.txs[1].to == "'$to'") and (.result.txs[1].groupCount == '$groupCount')'
    http_req "$req" ${MAIN_HTTP} "$resok" "$FUNCNAME"
}

chain_SignRawTx() {
    local fee=1000000
    local privkey="CC38546E9E659D15E6B4893F0AB32A06D103931A8230B0BDE71459D2B27D6944"

    tx1="0a05636f696e73122d18010a291080ade20422223145444467684174674273616d724e45744e6d5964517a43315145684c6b7238377420a08d0628e1ddcae60530f6db93c0e0d3f1ff5e3a223145444467684174674273616d724e45744e6d5964517a43315145684c6b72383774"
    tx=$(curl -ksd '{"method":"Chain.SignRawTx","params":[{"expire":"120s","fee":'$fee',"privkey":"'$privkey'","txHex":"'$tx1'"}]}' ${MAIN_HTTP} | jq -r ".result")

    req='{"method":"Chain.DecodeRawTransaction","params":[{"txHex":"'"$tx"'"}]}'
    resok='(.error|not) and (.result.txs[0].execer == "coins") and (.result.txs[0].to == "1EDDghAtgBsamrNEtNmYdQzC1QEhLkr87t") and (.result.txs[0].fee == '$fee') and (.result.txs[0].from == "14KEKbYtKKQm4wMthSK9J4La4nAiidGozt")'
    http_req "$req" ${MAIN_HTTP} "$resok" "$FUNCNAME"
}

chain_SendTransaction() {
    local fee=1000000
    local exec="coins"
    local to="1EDDghAtgBsamrNEtNmYdQzC1QEhLkr87t"
    local privkey="CC38546E9E659D15E6B4893F0AB32A06D103931A8230B0BDE71459D2B27D6944"

    tx1="0a05636f696e73122d18010a291080ade20422223145444467684174674273616d724e45744e6d5964517a43315145684c6b7238377420a08d0628e1ddcae60530f6db93c0e0d3f1ff5e3a223145444467684174674273616d724e45744e6d5964517a43315145684c6b72383774"
    tx=$(curl -ksd '{"method":"Chain.SignRawTx","params":[{"expire":"120s","fee":'$fee',"privkey":"'$privkey'","txHex":"'$tx1'"}]}' ${MAIN_HTTP} | jq -r ".result")

    req='{"method":"Chain.SendTransaction","params":[{"data":"'"$tx"'"}]}'
    http_req "$req" ${MAIN_HTTP} '(.error|not) and (.result != null)' "$FUNCNAME"
}

chain_CreateNoBalanceTransaction() {
    local to="1EDDghAtgBsamrNEtNmYdQzC1QEhLkr87t"
    local txHex=""
    local exec=""
    local coinexec=""

    if [ "$IS_PARA" == true ]; then
        exec="user.p.para.none"
        coinexec="user.p.para.coins"
        txHex="0a11757365722e702e706172612e636f696e73122d18010a291080ade20422223145444467684174674273616d724e45744e6d5964517a43315145684c6b7238377420a08d0630e6cbfbf1a7bafcb8263a2231415662506538776f524a7a7072507a4575707735554262433259507331344a4354"
    else
        exec="none"
        coinexec="coins"
        txHex="0a05636f696e73122d18010a291080ade20422223145444467684174674273616d724e45744e6d5964517a43315145684c6b7238377420a08d0630a1938af2e88e97fb0d3a223145444467684174674273616d724e45744e6d5964517a43315145684c6b72383774"
    fi

    tx=$(curl -ksd '{"method":"Chain.CreateNoBalanceTransaction","params":[{"txHex":"'$txHex'"}]}' ${MAIN_HTTP} | jq -r ".result")
    req='{"method":"Chain.DecodeRawTransaction","params":[{"txHex":"'"$tx"'"}]}'
    resok='(.error|not) and (.result.txs[0].execer == "'$exec'") and (.result.txs[0].groupCount == 2) and (.result.txs[1].execer == "'$coinexec'") and (.result.txs[1].groupCount == 2) and (.result.txs[1].to == "'$to'")'
    http_req "$req" ${MAIN_HTTP} "$resok" "$FUNCNAME"
}

chain_GetBlockHash() {
    http_req '{"method":"Chain.GetBlockHash","params":[{"height":1}]}' ${MAIN_HTTP} '(.error|not) and (.result| has("hash"))' "$FUNCNAME"
}

chain_GenSeed() {
    http_req '{"method":"Chain.GenSeed", "params":[{"lang":0}]}' ${MAIN_HTTP} '(.error|not) and (.result| has("seed"))' "$FUNCNAME"
    seed=$(curl -ksd '{"method":"Chain.GenSeed", "params":[{"lang":0}]}' ${MAIN_HTTP} | jq -r ".result.seed")
}

chain_SaveSeed() {
    req='{"method":"Chain.SaveSeed", "params":[{"seed":"'"$seed"'", "passwd": "1314fuzamei"}]}'
    http_req "$req" ${MAIN_HTTP} '(.error|not) and (.result| has("isOK"))' "$FUNCNAME"
}

chain_GetSeed() {
    http_req '{"method":"Chain.GetSeed", "params":[{"passwd": "1314fuzamei"}]}' ${MAIN_HTTP} '(.error|not) and (.result| has("seed"))' "$FUNCNAME"
}

chain_testSeed() {
    seed=""
    chain_GenSeed
    chain_SaveSeed
    chain_GetSeed
}

chain_GetWalletStatus() {
    http_req '{"method":"Chain.GetWalletStatus","params":[{}]}' ${MAIN_HTTP} '(.error|not) and (.result| [has("isWalletLock", "isAutoMining", "isHasSeed", "isTicketLock"), true] | unique | length == 1)' "$FUNCNAME"
}

chain_GetBalance() {
    http_req '{"method":"Chain.GetBalance","params":[{"addresses" : ["14KEKbYtKKQm4wMthSK9J4La4nAiidGozt"], "execer" : "coins"}]}' ${MAIN_HTTP} '(.error|not) and (.result[0] | [has("balance", "frozen"), true] | unique | length == 1)' "$FUNCNAME"
}

chain_GetAllExecBalance() {
    resok='(.error|not) and (.result| [has("addr", "execAccount"), true] | unique | length == 1) and (.result.execAccount | [map(has("execer", "account")), true] | flatten | unique | length == 1) and ([.result.execAccount[].account] | [map(has("balance", "frozen")), true] | flatten | unique | length == 1)'
    http_req '{"method":"Chain.GetAllExecBalance", "params":[{"addr" : "14KEKbYtKKQm4wMthSK9J4La4nAiidGozt"}]}' ${MAIN_HTTP} "$resok" "$FUNCNAME"
}

chain_ExecWallet() {
    req='{"method":"Chain.ExecWallet", "params":[{"funcName" : "NewAccountByIndex", "payload" : {"data" : 100000009}, "stateHash" : "", "execer" : "wallet" }]}'
    http_req "$req" ${MAIN_HTTP} '(.error|not) and (.result | has("data"))' "$FUNCNAME"
}

chain_Query() {
    http_req '{"method":"Chain.Query", "params":[{ "execer":"coins", "funcName": "GetTxsByAddr", "payload" : {"addr" : "1KSBd17H7ZK8iT37aJztFB22XGwsPTdwE4"}}]}' ${MAIN_HTTP} '(. | has("result"))' "$FUNCNAME"
}

chain_Version() {
    http_req '{"method":"Chain.Version", "params":[{}]}' ${MAIN_HTTP} '(.error|not) and (.result)' "$FUNCNAME"
}

chain_GetTotalCoins() {
    http_req '{"method":"Chain.GetTotalCoins", "params":[{"symbol" : "bty", "stateHash":"", "startKey":"", "count":2, "execer":"coins"}]}' ${MAIN_HTTP} '(.error|not) and (.result| has("count"))' "$FUNCNAME"
}

chain_IsSync() {
    http_req '{"method":"Chain.IsSync", "params":[{}]}' ${MAIN_HTTP} '(.error|not) and (. | has("result"))' "$FUNCNAME"
}

chain_IsNtpClockSync() {
    http_req '{"method":"Chain.IsNtpClockSync", "params":[{}]}' ${MAIN_HTTP} '(.error|not) and (. | has("result"))' "$FUNCNAME"
}

chain_GetServerTime() {
    http_req '{"method":"Chain.GetServerTime", "params":[{}]}' ${MAIN_HTTP} '(.error|not) and (.result.currentTimestamp>0)' "$FUNCNAME"
}

chain_GetCryptoList() {
    http_req '{"method":"Chain.GetCryptoList", "params":[{}]}' ${MAIN_HTTP} '(.error|not) and (.result.cryptos[]|select(.name=="none").typeID==10)' "$FUNCNAME"
}

run_testcases() {
    #    set -x
    set +e
    IS_PARA=$(echo '"'"${1}"'"' | jq '.|contains("8901")')
    echo "ipara=$IS_PARA"

    chain_lock
    chain_unlock

    chain_WalletTxList "$1"
    chain_ImportPrivkey "$1"
    chain_DumpPrivkey "$1"
    chain_DumpPrivkeysFile "$1"
    chain_ImportPrivkeysFile "$1"
    chain_SendToAddress "$1"
    chain_SetTxFee "$1"
    chain_SetLabl "$1"
    chain_GetPeerInfo "$1"
    chain_GetHeaders "$1"
    chain_GetLastMemPool "$1"
    chain_GetProperFee "$1"
    chain_GetBlockOverview "$1"
    chain_GetAddrOverview "$1"

    chain_QueryTotalFee
    chain_GetNetInfo
    chain_GetFatalFailure
    chain_DecodeRawTransaction
    chain_GetTimeStatus
    chain_GetLastBlockSequence
    chain_GetBlockSequences
    chain_GetBlockByHashes
    chain_ConvertExectoAddr
    chain_GetExecBalance
    chain_AddPushSubscribe
    chain_ListPushes
    chain_GetPushSeqLastNum
    chain_GetCoinSymbol
    chain_GetChainID

    chain_GetHexTxByHash
    chain_QueryTransaction
    chain_GetBlocks
    chain_GetLastHeader
    chain_GetTxByAddr
    chain_GetTxByHashes
    chain_GetMempool
    chain_GetAccountsV2
    chain_GetAccounts
    chain_NewAccount

    chain_CreateRawTransaction
    chain_CreateTransaction
    chain_ReWriteRawTx
    chain_CreateRawTxGroup
    chain_SignRawTx
    chain_SendTransaction
    chain_CreateNoBalanceTransaction

    chain_GetBlockHash
    chain_testSeed
    chain_GetWalletStatus
    chain_GetBalance
    chain_GetAllExecBalance
    chain_ExecWallet
    chain_Query
    chain_Version
    chain_GetTotalCoins
    chain_IsSync
    chain_IsNtpClockSync

    #这两个测试放在最后
    chain_SetPasswd "$1"
    chain_MergeBalance "$1"
    set -e
}

function system_test_rpc() {
    MAIN_HTTP=$1
    echo "=========== # system rpc test ============="
    echo "ip=$1"

    run_testcases "$1"

    if [ -n "$CASE_ERR" ]; then
        echo -e "${RED}======system rpc test fail=======${NOC}"
        exit 1
    else
        echo -e "${GRE}======system rpc test pass=======${NOC}"
    fi
}
