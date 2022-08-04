package client

import (
	. "ethrpc/utils"
	"math/big"
)

func (rpc *ethRpc) EthBlockNumber() (*big.Int, error) {
	var response string
	if err := rpc.call("eth_blockNumber", &response); err != nil {
		return new(big.Int), err
	}
	return ParseHexToBigInt(response, 0), nil
}
