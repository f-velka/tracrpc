package v1_1_8

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	tracrpc "github.com/f-velka/go-trac-rpc"
)

const (
	search_get_search_filters string = "search.getSearchFilters"
	search_perform_search     string = "search.performSearch"
)

// SearchService represents search API service.
type SearchService struct {
	rpc tracrpc.RpcClient
}

// SearchFilter represents the info returned by search.getSearchFilters.
type SearchFilter struct {
	Name        string
	Description string
}

// SearchResult represents the result returned by search.performSearch.
type SearchResult struct {
	Href    string
	Title   string
	Date    time.Time
	Author  string
	Excerpt string
}

// newSearchService creates new SearchService instance.
func newSearchService(rpc tracrpc.RpcClient) (*SearchService, error) {
	if rpc == nil {
		return nil, errors.New("rpc client cannot be nil")
	}

	return &SearchService{
		rpc: rpc,
	}, nil
}

// GetSearchFilters calls search.getSearchFilters.
func (s *SearchService) GetSearchFilters() ([]SearchFilter, error) {
	var rawReply [][]string
	if err := s.rpc.Call(search_get_search_filters, nil, &rawReply); err != nil {
		return nil, err
	}

	reply := make([]SearchFilter, 0, len(rawReply))
	for _, elem := range rawReply {
		if len(elem) != 2 {
			return nil, fmt.Errorf("%s: unexpected reply form. got=%v", search_get_search_filters, elem)
		}
		reply = append(reply, SearchFilter{
			Name:        elem[0],
			Description: elem[1],
		})
	}

	return reply, nil
}

// PerformSearch calls search.performSearch.
func (s *SearchService) PerformSearch(query *string, filterNames []string) ([]SearchResult, error) {
	args := packArgs(query, &filterNames)
	var rawReply [][]interface{}
	if err := s.rpc.Call(search_perform_search, args, &rawReply); err != nil {
		return nil, err
	}

	reply := make([]SearchResult, 0, len(rawReply))
	for _, elem := range rawReply {
		if len(elem) != 5 {
			return nil, fmt.Errorf("%s: unexpected reply form. got=%v", search_perform_search, elem)
		}
		var result SearchResult
		if isNil(elem[0]) {
			// do nothing
		} else if href, ok := elem[0].(string); ok {
			result.Href = href
		} else {
			return nil, fmt.Errorf("%s: unexpected Href type. got=%v", search_perform_search, elem[0])
		}
		if isNil(elem[1]) {
			// do nothing
		} else if title, ok := elem[1].(string); ok {
			result.Title = title
		} else {
			return nil, fmt.Errorf("%s: unexpected Title type. got=%v", search_perform_search, elem[1])
		}
		if isNil(elem[2]) {
			// do nothing
		} else if date, ok := elem[2].(time.Time); ok {
			result.Date = date
		} else {
			return nil, fmt.Errorf("%s: unexpected Date type. got=%v", search_perform_search, elem[2])
		}
		if isNil(elem[3]) {
			// do nothing
		} else if author, ok := elem[3].(string); ok {
			result.Author = author
		} else {
			return nil, fmt.Errorf("%s: unexpected Author type. got=%v", search_perform_search, elem[3])
		}
		if isNil(elem[4]) {
			// do nothing
		} else if excerpt, ok := elem[4].(string); ok {
			result.Excerpt = excerpt
		} else {
			return nil, fmt.Errorf("%s: unexpected Excerpt type. got=%v", search_perform_search, elem[4])
		}
		reply = append(reply, result)
	}

	return reply, nil
}

// isNil checks interface{} is nil or not.
func isNil(v interface{}) bool {
	if v == nil {
		return true
	}

	switch reflect.TypeOf(v).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(v).IsNil()
	}

	return false
}
