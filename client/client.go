package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func NewEthRpc(url string, debug bool) *ethRpc {
	rpc := &ethRpc{
		url:    url,
		client: http.DefaultClient,
		Debug:  debug,
	}
	return rpc
}

func (rpc *ethRpc) Call(method string, params ...interface{}) (json.RawMessage, error) {
	request := ethRequest{
		Id:      1,
		JsonRpc: "2.0",
		Method:  method,
		Params:  params,
	}

	if len(params) == 0 {
		// prevent params: null in data
		request.Params = make([]interface{}, 0)
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	response, err := rpc.client.Post(rpc.url, "application/json", bytes.NewBuffer(body))
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if rpc.Debug {
		fmt.Println(fmt.Sprintf("%s\nRequest: %s\nResponse: %s\n", method, body, data))
	}

	resp := new(ethResponse)
	if err := json.Unmarshal(data, resp); err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, *resp.Error
	}

	return resp.Result, nil
}

func (rpc *ethRpc) RawCall(method string, params ...interface{}) (json.RawMessage, error) {
	return rpc.Call(method, params...)
}

func (rpc *ethRpc) call(method string, target interface{}, params ...interface{}) error {
	result, err := rpc.Call(method, params...)
	if err != nil {
		return err
	}
	return json.Unmarshal(result, target)
}
