// handlers_test.go
package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockRPCClient struct{}

func (m *MockRPCClient) Call(reqBody RPCRequest) ([]byte, error) {
	mockResponse := `{"jsonrpc":"2.0","id":1,"result":"0x1b4"}`
	return []byte(mockResponse), nil
}

func TestGetBlockNumber(t *testing.T) {
	mockClient := &MockRPCClient{}

	payload := GetBlockNumberRequest{ID: 2}
	jsonPayload, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", "/getBlockNumber", bytes.NewBuffer(jsonPayload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := getBlockNumber(mockClient)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"jsonrpc":"2.0","id":1,"result":"0x1b4"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestGetBlockByNumber(t *testing.T) {
	mockClient := &MockRPCClient{}

	payload := GetBlockByNumberRequest{BlockNumber: "0x1b4", ID: 2}
	jsonPayload, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", "/getBlockByNumber", bytes.NewBuffer(jsonPayload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := getBlockByNumber(mockClient)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"jsonrpc":"2.0","id":1,"result":"0x1b4"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestHealthCheck(t *testing.T) {
	req, err := http.NewRequest("GET", "/health_check", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(healthCheck)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "ok"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
