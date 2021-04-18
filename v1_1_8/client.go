package v1_1_8

import (
	"net/http"

	"github.com/kolo/xmlrpc"
)

// Client replresents trac API client.
type Client struct {
	Wiki *WikiService
}

// base64String represents Base64-Encoded bytes.
type base64String = xmlrpc.Base64

// NewClient creates new Client
func NewClient(url string, transport http.RoundTripper) (*Client, error) {
	xmlrpcClient, err := xmlrpc.NewClient(url, transport)
	if err != nil {
		return nil, err
	}
	wiki, err := newWikiService(xmlrpcClient)
	if err != nil {
		return nil, err
	}

	return &Client{
		Wiki: wiki,
	}, nil
}
