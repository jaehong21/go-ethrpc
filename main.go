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
	rpc := client.NewEthRpc(os.Getenv("INFURA_NODE_URL"), false)
	// rpc := client.NewEthRpc(os.Getenv("LOCAL_NODE_URL"), true)
	result, _ := rpc.EthBlockNumber()
	fmt.Println(result.Int64())
}
