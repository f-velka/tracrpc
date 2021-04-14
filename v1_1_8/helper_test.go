package v1_1_8

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func PStr(str string) *string {
	return &str
}

func PInt(i int) *int {
	return &i
}

type RpcClientMock struct {
}

func (c *RpcClientMock) Call(methodName string, args interface{}, reply interface{}) error {
	return nil
}

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(reply string) *Client {
	c, _ := NewClient(
		"http://example.com",
		RoundTripFunc(func(_ *http.Request) *http.Response {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBufferString(reply)),
			}
		}),
	)
	return c
}
