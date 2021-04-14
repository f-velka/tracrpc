package v1_1_8

import (
	"net/http"
	"time"

	"github.com/f-velka/go-trac-rpc/common"
	"github.com/kolo/xmlrpc"
)

const (
	DEFAULT_TIMEOUT time.Duration = 10 * time.Second
)

type Client struct {
	rpc  common.RpcClient
	Wiki *WikiService
}

func NewClient(url string, transport http.RoundTripper) (*Client, error) {
	xmlrpcClient, err := xmlrpc.NewClient(url, transport)
	if err != nil {
		return nil, err
	}
	wiki, err := NewWikiService(xmlrpcClient)
	if err != nil {
		return nil, err
	}

	return &Client{
		rpc:  xmlrpcClient,
		Wiki: wiki,
	}, nil
}
