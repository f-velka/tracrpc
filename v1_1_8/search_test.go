package v1_1_8

import (
	"reflect"
	"testing"
	"time"

	tracrpc "github.com/f-velka/go-trac-rpc"
)

func TestNewSearchService(t *testing.T) {
	tests := []struct {
		name      string
		rpcClient tracrpc.RpcClient
		wantErr   bool
	}{
		{
			name:      "OK",
			rpcClient: &RpcClientMock{},
			wantErr:   false,
		},
		{
			name:      "NG",
			rpcClient: nil,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := newSearchService(tt.rpcClient)
			if err != nil {
				if tt.wantErr {
					return
				}
				t.Fatal(err)
			}
		})
	}
}

func TestGetSearchFilters(t *testing.T) {
	test := struct {
		reply    string
		expected []SearchFilter
	}{
		`<?xml version='1.0'?>
<methodResponse>
<params>
<param>
<value><array><data>
<value><array><data>
<value><string>changeset</string></value>
<value><string>Changesets</string></value>
</data></array></value>
<value><array><data>
<value><string>milestone</string></value>
<value><string>Milestones</string></value>
</data></array></value>
</data></array></value>
</param>
</params>
</methodResponse>`,
		[]SearchFilter{
			{
				Name:        "changeset",
				Description: "Changesets",
			},
			{
				Name:        "milestone",
				Description: "Milestones",
			},
		},
	}

	c := NewTestClient(search_get_search_filters, nil, test.reply)
	res, err := c.Search.GetSearchFilters()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(res, test.expected) {
		t.Fatalf("unexpected result. expected=%v, got=%v", test.expected, res)
	}
}

func TestPerformSearch(t *testing.T) {
	test := struct {
		query       *string
		filterNames []string
		reply       string
		expected    []SearchResult
	}{
		tracrpc.String("myself"),
		[]string{"environment", "others"},
		`<?xml version='1.0'?>
<methodResponse>
<params>
<param>
<value><array><data>
<value><array><data>
<value><string>http://example.com</string></value>
<value><string>monday</string></value>
<value><dateTime.iso8601>20210401T00:00:00</dateTime.iso8601></value>
<value><string>sunday</string></value>
<value><string>sigh</string></value>
</data></array></value>
<value><array><data>
<value><nil/></value>
<value><string></string></value>
<value><dateTime.iso8601></dateTime.iso8601></value>
<value><string></string></value>
<value><nil/></value>
</data></array></value>
</data></array></value>
</param>
</params>
</methodResponse>`,
		[]SearchResult{
			{
				Href:    "http://example.com",
				Title:   "monday",
				Date:    time.Date(2021, time.April, 1, 0, 0, 0, 0, time.UTC),
				Author:  "sunday",
				Excerpt: "sigh",
			},
			{},
		},
	}

	c := NewTestClient(search_perform_search, packArgs(test.query, &test.filterNames), test.reply)
	res, err := c.Search.PerformSearch(test.query, test.filterNames)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(res, test.expected) {
		t.Fatalf("unexpected result. expected=%v, got=%v", test.expected, res)
	}
}
