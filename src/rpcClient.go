// rpcClient.go for unit test, we can use mockClient
package main

type RPCClient interface {
	Call(reqBody RPCRequest) ([]byte, error)
}

type DefaultRPCClient struct{}

func (c *DefaultRPCClient) Call(reqBody RPCRequest) ([]byte, error) {
	return callRPC(reqBody)
}
