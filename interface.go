package tracrpc

// RpcClient represents a client for RPC.
type RpcClient interface {
	Call(methodName string, args interface{}, reply interface{}) error
}
