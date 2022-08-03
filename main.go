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
	price, _ := rpc.EthGasPrice()
	num, _ := rpc.EthGetTransactionByHash("0x5e876a1c4dd953d7fbef7950c7670823fb8d71132fd815c7d9c820f094396495")
	fmt.Println("price: ", price)
	fmt.Println(num)
}
