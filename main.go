package main

import (
	"ethrpc/client"
	"fmt"
)

func main() {
	rpc := client.NewEthRpc("https://mainnet.infura.io/v3/5d46f79631a84babbe5a1253307d7145", true)
	// num, _ := rpc.EthGetBalance("0xae48484cF810f05C3761664d144ccd4299103545", "latest")
	num, _ := rpc.EthGasPrice()
	fmt.Println(num)
}
