### Introduction
Golang implementation on ethereum network using json-rpc

### Docs
~~~
func main() {
	rpc := client.NewEthRpc("http://localhost:7545", true)
	num, _ := rpc.EthBlockNumber()
	fmt.Println(num)
}
~~~

### Reference
https://github.com/onrik/ethrpc