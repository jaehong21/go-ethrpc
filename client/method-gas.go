package client

import (
	. "ethrpc/utils"
	"math/big"
)

func (rpc *ethRpc) EthGasPrice() (*big.Float, error) {
	var response string
	if err := rpc.call("eth_gasPrice", &response); err != nil {
		return new(big.Float), err
	}
	return ParseHexToBigFloat(response, 9), nil
}
