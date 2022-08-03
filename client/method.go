package client

import (
	"encoding/json"
	"ethrpc/utils"
	"math/big"
)

func (rpc *ethRpc) EthBlockNumber() (*big.Int, error) {
	var response string
	if err := rpc.call("eth_blockNumber", &response); err != nil {
		return new(big.Int), err
	}
	return utils.ParseHexToBigInt(response, 0), nil
}

func (rpc *ethRpc) EthGasPrice() (*big.Float, error) {
	var response string
	if err := rpc.call("eth_gasPrice", &response); err != nil {
		return new(big.Float), err
	}
	return utils.ParseHexToBigFloat(response, 9), nil
}

func (rpc *ethRpc) EthGetBalance(address, block string) (*big.Float, error) {
	var response string
	if err := rpc.call("eth_getBalance", &response, address, block); err != nil {
		return new(big.Float), err
	}
	num := utils.ParseHexToBigFloat(response, 18)
	return num, nil
}

func (rpc *ethRpc) EthGetTransactionByHash(hash string) (string, error) {
	tx, err := rpc.getTransaction("eth_getTransactionByHash", hash)
	if err != nil {
		return "", err
	}
	result, err := json.Marshal(tx)
	if err != nil {
		return "", err
	}
	return string(result), err
}
