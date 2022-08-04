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
	// rpc := client.NewEthRpc(os.Getenv("INFURA_NODE_URL"), false)
	rpc := client.NewEthRpc(os.Getenv("LOCAL_NODE_URL"), true)
	result, _ := rpc.EthGetTransactionByHash("0x810c26a9241cad4422643d2a406789114421fe43cc00fb0397793b859293b69b")
	fmt.Printf(result)
}
