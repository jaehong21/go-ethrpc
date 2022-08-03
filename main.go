package main

import (
	"ethrpc/client"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	rpc := client.NewEthRpc(os.Getenv("INFURA_NODE_URL"), true)
	// num, _ := rpc.EthGetBalance("0xae48484cF810f05C3761664d144ccd4299103545", "latest")
	num, _ := rpc.EthGasPrice()
	fmt.Println(num)
}
