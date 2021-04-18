package v1_1_8

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

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
	expectedArgs       interface{}
}

func (c *rpcClientWithExpectedMethodName) Call(methodName string, args interface{}, reply interface{}) error {
	if methodName != c.expectedMethodName {
		return fmt.Errorf("unexpected method name. expected=%s, got=%s", c.expectedMethodName, methodName)
	}
	if !reflect.DeepEqual(args, c.expectedArgs) {
		return fmt.Errorf("unexpected args. expected=%v, got=%v", c.expectedArgs, args)
	}

	return c.rpcClient.Call(methodName, args, reply)
}

func newRpcClientWithExpectedValues(rpcClient tracrpc.RpcClient, expectedMethodName string, expectedArgs interface{}) tracrpc.RpcClient {
	return &rpcClientWithExpectedMethodName{
		rpcClient:          rpcClient,
		expectedMethodName: expectedMethodName,
		expectedArgs:       expectedArgs,
	}
}

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(expectedMethodName string, expectedArgs interface{}, reply string) *Client {
	c, _ := NewClient(
		"http://example.com",
		RoundTripFunc(func(_ *http.Request) *http.Response {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBufferString(reply)),
			}
		}),
	)
	c.Wiki.rpc = newRpcClientWithExpectedValues(
		c.Wiki.rpc,
		expectedMethodName,
		expectedArgs,
	)
	return c
}
