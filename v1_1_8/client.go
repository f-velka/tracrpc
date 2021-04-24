package v1_1_8

import (
	"net/http"
	"reflect"

	"github.com/kolo/xmlrpc"
)

// Client replresents trac API client.
type Client struct {
	Search *SearchService
	System *SystemService
	Wiki   *WikiService
}

// base64String represents Base64-Encoded bytes.
type base64String = xmlrpc.Base64

// NewClient creates new Client
func NewClient(url string, transport http.RoundTripper) (*Client, error) {
	xmlrpcClient, err := xmlrpc.NewClient(url, transport)
	if err != nil {
		return nil, err
	}
	search, err := newSearchService(xmlrpcClient)
	if err != nil {
		return nil, err
	}
	system, err := newSystemService(xmlrpcClient)
	if err != nil {
		return nil, err
	}
	wiki, err := newWikiService(xmlrpcClient)
	if err != nil {
		return nil, err
	}

	return &Client{
		Search: search,
		System: system,
		Wiki:   wiki,
	}, nil
}

// packArgs packs args into []interface{}.
// Args must be pointers.
func packArgs(args ...interface{}) []interface{} {
	packed := make([]interface{}, 0, len(args))
	for _, arg := range args {
		if reflect.TypeOf(arg).Kind() != reflect.Ptr {
			panic("args must be pointers.")
		}

		if !reflect.ValueOf(arg).IsNil() {
			packed = append(packed, arg)
		}
	}

	return packed
}
