package v1_1_8

import (
	"errors"
	"fmt"
	"time"

	"github.com/f-velka/go-trac-rpc/common"
)

const (
	wiki_get_recent_changes        string = "wiki.getRecentChanges"
	wiki_get_rpc_version_supported string = "wiki.getRPCVersionSupported"
	wiki_get_page                  string = "wiki.getPage"
	wiki_get_page_version          string = "wiki.getPageVersion"
	wiki_get_page_html             string = "wiki.getPageHTML"
	wiki_get_page_html_version     string = "wiki.getPageHTMLVersion"
	wiki_get_all_pages             string = "wiki.getAllPages"
	wiki_get_page_info             string = "wiki.getPageInfo"
	wiki_get_page_info_version     string = "wiki.getPageInfoVersion"
	wiki_put_page                  string = "wiki.putPage"
	wiki_list_attachments          string = "wiki.listAttachments"
	wiki_get_appachments           string = "wiki.getAttachment"
	wiki_put_attachment            string = "wiki.putAttachment"
	wiki_put_attachment_ex         string = "wiki.putAttachmentEx"
	wiki_delete_page               string = "wiki.deletePage"
	wiki_delete_attachment         string = "wiki.deleteAttachment"
	wiki_list_links                string = "wiki.listLinks"
	wiki_wiki_to_html              string = "wiki.wikiToHtml"
)

type WikiService struct {
	rpc common.RpcClient
}

func NewWikiService(rpc common.RpcClient) (*WikiService, error) {
	if rpc == nil {
		return nil, errors.New("rpc client cannot be nil")
	}

	return &WikiService{
		rpc: rpc,
	}, nil
}

type GetPageOptions struct {
	PageName *string
	Version  *int
}

type GetRecentChangesResult struct {
	Name         string    `xmlrpc:"name"`
	LastModified time.Time `xmlrpc:"lastModified"`
	Author       string    `xmlrpc:"author"`
	Version      int       `xmlrpc:"version"`
	Comment      string    `xmlrpc:"comment"`
}

func (w *WikiService) GetRecentChanges(since time.Time) ([]GetRecentChangesResult, error) {
	reply := []GetRecentChangesResult{}
	if err := w.rpc.Call(wiki_get_recent_changes, since, &reply); err != nil {
		return nil, err
	}

	return reply, nil
}

func (w *WikiService) GetRPCVersionSupported() (int, error) {
	var reply int
	if err := w.rpc.Call(wiki_get_rpc_version_supported, nil, &reply); err != nil {
		return 0, err
	}

	return reply, nil
}

func (w *WikiService) GetPage(options *GetPageOptions) (string, error) {
	args, err := readGetPageOptions(options, wiki_get_page)
	if err != nil {
		return "", err
	}

	var reply string
	if err := w.rpc.Call(wiki_get_page, args, &reply); err != nil {
		return "", err
	}

	return reply, nil
}

func (w *WikiService) GetPageVersion(options *GetPageOptions) (string, error) {
	args, err := readGetPageOptions(options, wiki_get_page_version)
	if err != nil {
		return "", err
	}

	var reply string
	if err := w.rpc.Call(wiki_get_page_version, args, &reply); err != nil {
		return "", err
	}

	return reply, nil
}

func (w *WikiService) GetPageHTML(options *GetPageOptions) (string, error) {
	args, err := readGetPageOptions(options, wiki_get_page_html)
	if err != nil {
		return "", err
	}

	var reply string
	if err := w.rpc.Call(wiki_get_page_html, args, &reply); err != nil {
		return "", err
	}

	return reply, nil
}

func (w *WikiService) GetPageHTMLVersion(options *GetPageOptions) (string, error) {
	args, err := readGetPageOptions(options, wiki_get_page_html_version)
	if err != nil {
		return "", err
	}

	var reply string
	if err := w.rpc.Call(wiki_get_page_html_version, args, &reply); err != nil {
		return "", err
	}

	return reply, nil
}

func (w *WikiService) GetAllPages() ([]string, error) {
	var reply []string
	if err := w.rpc.Call(wiki_get_all_pages, nil, &reply); err != nil {
		return nil, err
	}

	return reply, nil
}

func readGetPageOptions(options *GetPageOptions, methodName string) ([]interface{}, error) {
	if options.PageName == nil {
		return nil, fmt.Errorf("%s: PageName cannot be nil", methodName)
	}

	args := []interface{}{*options.PageName}
	if options.Version != nil {
		args = append(args, *options.Version)
	}

	return args, nil
}
