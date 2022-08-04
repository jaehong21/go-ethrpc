package client

import (
	. "ethrpc/utils"
	"math/big"
)

func (rpc *ethRpc) EthGetBalance(address, block string) (*big.Float, error) {
	var response string
	if err := rpc.call("eth_getBalance", &response, address, block); err != nil {
		return new(big.Float), err
	}
	num := ParseHexToBigFloat(response, 18)
	return num, nil
}
