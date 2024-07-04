// handlers.go
package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

const rpcURL = "https://polygon-rpc.com/"

type RPCRequest struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params,omitempty"`
	ID      int           `json:"id"`
}

type GetBlockNumberRequest struct {
	ID int `json:"id"`
}

type GetBlockByNumberRequest struct {
	BlockNumber string `json:"blockNumber"`
	ID          int    `json:"id"`
}

func getBlockNumber(client RPCClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqPayload GetBlockNumberRequest

		err := json.NewDecoder(r.Body).Decode(&reqPayload)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		reqBody := RPCRequest{
			Jsonrpc: "2.0",
			Method:  "eth_blockNumber",
			ID:      reqPayload.ID,
		}

		response, err := client.Call(reqBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}

func getBlockByNumber(client RPCClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params GetBlockByNumberRequest

		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		reqBody := RPCRequest{
			Jsonrpc: "2.0",
			Method:  "eth_getBlockByNumber",
			Params:  []interface{}{params.BlockNumber, true},
			ID:      params.ID,
		}

		response, err := client.Call(reqBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("ok"))
}

func callRPC(reqBody RPCRequest) ([]byte, error) {
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(rpcURL, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
