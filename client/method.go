package client

import (
	"math/big"
)

func (rpc *ethRpc) EthBlockNumber() (*big.Int, error) {
	var response string
	if err := rpc.call("eth_blockNumber", &response); err != nil {
		return new(big.Int), err
	}
	return ParseHexToBigInt(response, 0), nil
}

func (rpc *ethRpc) EthGasPrice() (*big.Float, error) {
	var response string
	if err := rpc.call("eth_gasPrice", &response); err != nil {
		return new(big.Float), err
	}
	return ParseHexToBigFloat(response, 9), nil
}

func (rpc *ethRpc) EthGetBalance(address, block string) (*big.Float, error) {
	var response string
	if err := rpc.call("eth_getBalance", &response, address, block); err != nil {
		return new(big.Float), err
	}
	num := ParseHexToBigFloat(response, 18)
	return num, nil
}
