package client

import (
	"encoding/json"
	. "ethrpc/utils"
	"fmt"
	"math/big"
	"net/http"
)

type ethRpc struct {
	url    string
	client *http.Client
	Debug  bool
}

type ethRequest struct {
	Id      int      `json:"id"`
	JsonRpc string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
}

type ethResponse struct {
	Id      int             `json:"id"`
	JsonRpc string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
	Error   *ethError       `json:"error"`
}

type ethError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// duck-typed for error
func (err ethError) Error() string {
	return fmt.Sprintf("Error %d (%s)", err.Code, err.Message)
}

type txResponse struct {
	Hash                 string
	Nonce                string
	Id                   string
	BlockHash            string
	BlockNumber          string
	TransactionIndex     string
	From                 string
	To                   string
	Gas                  string
	GasPrice             string
	MaxPriorityFeePerGas string
	Input                string
}

type transaction struct {
	Hash                 string
	Nonce                int
	Id                   int
	BlockHash            string
	BlockNumber          int
	TransactionIndex     int
	From                 string
	To                   string
	Value                big.Int
	Gas                  *big.Float
	GasPrice             *big.Float
	MaxPriorityFeePerGas *big.Float
	Input                string
}

func (tx *txResponse) UnmarshalJson() *transaction {
	transaction := new(transaction)
	transaction.Hash = tx.Hash
	transaction.Nonce = ParseInt(tx.Nonce)
	transaction.Id = ParseInt(tx.Id)
	transaction.Hash = tx.Hash
	transaction.To = tx.To
	transaction.From = tx.To
	transaction.BlockHash = tx.BlockHash
	transaction.BlockNumber = ParseHexToInt(tx.BlockNumber, 0)
	transaction.TransactionIndex = ParseHexToInt(tx.TransactionIndex, 0)
	transaction.Gas = ParseHexToBigFloat(tx.Gas, 9)
	transaction.GasPrice = ParseHexToBigFloat(tx.GasPrice, 9)
	transaction.MaxPriorityFeePerGas = ParseHexToBigFloat(tx.MaxPriorityFeePerGas, 9)
	transaction.Input = tx.Input
	return transaction
}
