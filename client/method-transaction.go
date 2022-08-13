package client

import (
	"encoding/json"
	. "ethrpc/utils"
)

func (rpc *ethRpc) getTransaction(method string, params ...interface{}) (*transaction, error) {
	response := new(TxResponse)
	err := rpc.call(method, response, params...)
	return response.UnmarshalJSON(), err
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

func (rpc *ethRpc) EthSendTransaction(transaction TxRequest) (string, error) {
	var hash string
	if err := rpc.call("eth_sendTransaction", &hash, transaction); err != nil {
		return "", err
	}
	return hash, nil
}

func (rpc *ethRpc) EthGetTransactionCount(address, block string) (int, error) {
	var response string
	if err := rpc.call("eth_getTransactionCount", &response, address, block); err != nil {
		return 0, err
	}
	return ParseHexToInt(response, 0), nil
}
