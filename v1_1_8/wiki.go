package v1_1_8

import (
	"errors"
	"time"

	"github.com/f-velka/go-trac-rpc/common"
)

type WikiService struct {
	rpc common.RpcClient
}

func NewWikiService(rpc common.RpcClient) *WikiService {
	return &WikiService{
		rpc: rpc,
	}
}

type GetRecentChangesResult struct {
	Name         string    `xmlrpc:"name"`
	Author       string    `xmlrpc:"author"`
	Version      int       `xmlrpc:"version"`
	LastModified time.Time `xmlrpc:"lastModified"`
	Comment      string    `xmlrpc:"comment"`
}

func (w *WikiService) GetRecentChanges(since time.Time) ([]GetRecentChangesResult, error) {
	reply := []GetRecentChangesResult{}
	err := w.rpc.Call(
		"wiki.getRecentChanges",
		since,
		&reply,
	)
	if err != nil {
		return nil, err
	}

	return reply, err
}

func (w *WikiService) GetRPCVersionSupported() (int, error) {
	var reply int
	err := w.rpc.Call(
		"wiki.getRPCVersionSupported",
		nil,
		&reply,
	)
	if err != nil {
		return 0, err
	}

	return reply, err
}

func (w *WikiService) GetAllPages() ([]string, error) {
	var reply []string
	err := w.rpc.Call(
		"wiki.getAllPages",
		nil,
		&reply,
	)
	if err != nil {
		return nil, err
	}

	return reply, err
}

type GetPageOption struct {
	PageName *string
	Version  *int
}

func (w *WikiService) GetPage(option *GetPageOption) (string, error) {
	// args := make([]interface{}, 2)
	if option.PageName == nil {
		return "", errors.New("wiki.getPage: PageName cannot be null")
	}
	args := []interface{}{*option.PageName}
	if option.Version != nil {
		args = append(args, *option.Version)
	}

	var reply string
	err := w.rpc.Call(
		"wiki.getPage",
		args,
		&reply,
	)
	if err != nil {
		return "", err
	}

	return reply, nil
}
