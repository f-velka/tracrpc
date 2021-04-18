package tracrpc

type RpcClient interface {
	Call(methodName string, args interface{}, reply interface{}) error
}