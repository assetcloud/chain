
## 安装

```
$ sudo install chain-cli.bash  /usr/share/bash-completion/completions/chain-cli
```

不安装
```
. chain-cli.bash
```

```
# 重新开个窗口就有用了
$ ./chain/chain-cli 
account   bty       coins     evm       hashlock  mempool   privacy   retrieve  send      ticket    trade     version   
block     close     config    exec      help      net       relay     seed      stat      token     tx        wallet    
```

## 演示
```
linj@linj-TM1701:~$ ./chain/chain-cli 
account   bty       coins     evm       hashlock  mempool   privacy   retrieve  send      ticket    trade     version   
block     close     config    exec      help      net       relay     seed      stat      token     tx        wallet    
linj@linj-TM1701:~$ ./chain/chain-cli b
block  bty    
linj@linj-TM1701:~$ ./chain/chain-cli bty 
priv2priv  priv2pub   pub2priv   send       transfer   txgroup    withdraw   
linj@linj-TM1701:~$ ./chain/chain-cli bty t
transfer  txgroup   
linj@linj-TM1701:~$ ./chain/chain-cli bty transfer -
-a        --amount  -h        --help    -n        --note    --para    --rpc     -t        --to  
```


## 原理

 地址 https://confluence.33.cn/pages/viewpage.action?pageId=5967798
