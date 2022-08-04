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
	Id      int           `json:"id"`
	JsonRpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
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

type TxRequest struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Gas   string `json:"gas"`
	Value string `json:"value"`
	Data  string `json:"data"`
}

type TxResponse struct {
	Hash                 string `json:"hash"`
	Nonce                string `json:"nonce"`
	Id                   string `json:"chainId"`
	BlockHash            string `json:"blockHash"`
	BlockNumber          string `json:"blockNumber"`
	TransactionIndex     string `json:"transactionIndex"`
	From                 string `json:"from"`
	To                   string `json:"to"`
	Value                string `json:"value"`
	Gas                  string `json:"gas"`
	GasPrice             string `json:"gasPrice"`
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas"`
	Input                string `json:"input"`
}

type transaction struct {
	Hash                 string     `json:"hash"`
	Nonce                int        `json:"nonce"`
	Id                   int        `json:"chainId"`
	BlockHash            string     `json:"blockHash"`
	BlockNumber          int        `json:"blockNumber"`
	TransactionIndex     int        `json:"transactionIndex"`
	From                 string     `json:"from"`
	To                   string     `json:"to"`
	Value                big.Int    `json:"value"`
	Gas                  *big.Float `json:"gas"`
	GasPrice             *big.Float `json:"gasPrice"`
	MaxPriorityFeePerGas *big.Float `json:"maxPriorityFeePerGas"`
	Input                string     `json:"input"`
}

func (tx *TxResponse) UnmarshalJSON() *transaction {
	transaction := new(transaction)
	transaction.Hash = tx.Hash
	transaction.Nonce = ParseHexToInt(tx.Nonce, 0)
	// transaction.Id = ParseHexToInt(tx.Id, 0)
	transaction.Hash = tx.Hash
	transaction.To = tx.To
	transaction.From = tx.To
	transaction.BlockHash = tx.BlockHash
	transaction.BlockNumber = ParseHexToInt(tx.BlockNumber, 0)
	transaction.TransactionIndex = ParseHexToInt(tx.TransactionIndex, 0)
	transaction.Gas = ParseHexToBigFloat(tx.Gas, 9)
	transaction.GasPrice = ParseHexToBigFloat(tx.GasPrice, 9)
	// transaction.MaxPriorityFeePerGas = ParseHexToBigFloat(tx.MaxPriorityFeePerGas, 9)
	transaction.Input = tx.Input
	return transaction
}
