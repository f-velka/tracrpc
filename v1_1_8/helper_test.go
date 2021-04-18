package v1_1_8

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	tracrpc "github.com/f-velka/go-trac-rpc"
)

type RpcClientMock struct {
}

func (c *RpcClientMock) Call(methodName string, args interface{}, reply interface{}) error {
	return nil
}

type rpcClientWithExpectedMethodName struct {
	rpcClient          tracrpc.RpcClient
	expectedMethodName string
}

func (c *rpcClientWithExpectedMethodName) Call(methodName string, args interface{}, reply interface{}) error {
	if methodName != c.expectedMethodName {
		return fmt.Errorf("called method name is unexpected. expected=%s, got=%s", c.expectedMethodName, methodName)
	}

	return c.rpcClient.Call(methodName, args, reply)
}

func newRpcClientWithExpectedMethodName(rpcClient tracrpc.RpcClient, expectedMethodName string) tracrpc.RpcClient {
	return &rpcClientWithExpectedMethodName{
		rpcClient:          rpcClient,
		expectedMethodName: expectedMethodName,
	}
}

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(expectedMethodName string, reply string) *Client {
	c, _ := NewClient(
		"http://example.com",
		RoundTripFunc(func(_ *http.Request) *http.Response {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBufferString(reply)),
			}
		}),
	)
	c.Wiki.rpc = newRpcClientWithExpectedMethodName(c.Wiki.rpc, expectedMethodName)
	return c
}