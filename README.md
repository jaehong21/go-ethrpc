### Introduction
**This is a Project that is still Working on** <br />
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
https://ethereum.org/en/developers/docs/apis/json-rpc <br />
https://github.com/onrik/ethrpc
