package client

import (
	"encoding/json"
	"fmt"
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
