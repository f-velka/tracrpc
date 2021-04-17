package v1_1_8

import (
	"errors"
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

// WikiService is wiki API service.
type WikiService struct {
	rpc common.RpcClient
}

// PageOptions represents options sent on some WIKI APIs.
type PageOptions struct {
	PageName *string
	Version  *int
}

// PutPageParams represents params sent on wiki.putPage.
type PutPageParams struct {
	PageName   string
	Content    string
	Attributes PutPageAttributes
}

// PutPageAttributes represents attributes in putPageParams.
type PutPageAttributes struct {
	Readonly *bool   `xmlrpc:"readonly"`
	Author   *string `xmlrpc:"author"`
	Comment  *string `xmlrpc:"comment"`
}

// PageInfo represents the info returned by wiki.getRecentChanges, wiki.getPageInfo, wiki.getPageInfoVersion.
type PageInfo struct {
	Name         string    `xmlrpc:"name"`
	LastModified time.Time `xmlrpc:"lastModified"`
	Author       string    `xmlrpc:"author"`
	Version      int       `xmlrpc:"version"`
	Comment      string    `xmlrpc:"comment"`
}

// NewWikiService creates new SikiService instance.
func NewWikiService(rpc common.RpcClient) (*WikiService, error) {
	if rpc == nil {
		return nil, errors.New("rpc client cannot be nil")
	}

	return &WikiService{
		rpc: rpc,
	}, nil
}

// GetRecentChanges calls wiki.getRecentChanges.
func (w *WikiService) GetRecentChanges(since time.Time) ([]PageInfo, error) {
	reply := []PageInfo{}
	if err := w.rpc.Call(wiki_get_recent_changes, since, &reply); err != nil {
		return nil, err
	}

	return reply, nil
}

// GetRPCVersionSupported calls wiki.getRPCVersionSupported.
func (w *WikiService) GetRPCVersionSupported() (int, error) {
	var reply int
	if err := w.rpc.Call(wiki_get_rpc_version_supported, nil, &reply); err != nil {
		return 0, err
	}

	return reply, nil
}

// GetPage calls wiki.getPage
func (w *WikiService) GetPage(options *PageOptions) (string, error) {
	args := readPageOptions(options)

	var reply string
	if err := w.rpc.Call(wiki_get_page, args, &reply); err != nil {
		return "", err
	}

	return reply, nil
}

// GetPageVersion calls wiki.getPageVersion.
func (w *WikiService) GetPageVersion(options *PageOptions) (string, error) {
	args := readPageOptions(options)

	var reply string
	if err := w.rpc.Call(wiki_get_page_version, args, &reply); err != nil {
		return "", err
	}

	return reply, nil
}

// GetPageHTML calls wiki.getPageHTML.
func (w *WikiService) GetPageHTML(options *PageOptions) (string, error) {
	args := readPageOptions(options)

	var reply string
	if err := w.rpc.Call(wiki_get_page_html, args, &reply); err != nil {
		return "", err
	}

	return reply, nil
}

// GetPageHTMLVersion calls wiki.getPageHTMLVersion.
func (w *WikiService) GetPageHTMLVersion(options *PageOptions) (string, error) {
	args := readPageOptions(options)

	var reply string
	if err := w.rpc.Call(wiki_get_page_html_version, args, &reply); err != nil {
		return "", err
	}

	return reply, nil
}

// GetAllPages calls wiki.getAllPages.
func (w *WikiService) GetAllPages() ([]string, error) {
	var reply []string
	if err := w.rpc.Call(wiki_get_all_pages, nil, &reply); err != nil {
		return nil, err
	}

	return reply, nil
}

// GetPageInfo calls wiki.getPageInfo.
func (w *WikiService) GetPageInfo(options *PageOptions) (PageInfo, error) {
	args := readPageOptions(options)

	var reply PageInfo
	if err := w.rpc.Call(wiki_get_page_info, args, &reply); err != nil {
		return PageInfo{}, err
	}

	return reply, nil
}

// GetPageInfoVersion calls wiki.getPageInfoVersion.
func (w *WikiService) GetPageInfoVersion(options *PageOptions) (PageInfo, error) {
	args := readPageOptions(options)

	var reply PageInfo
	if err := w.rpc.Call(wiki_get_page_info_version, args, &reply); err != nil {
		return PageInfo{}, err
	}

	return reply, nil
}

// PutPage calls wiki.putPage
func (w *WikiService) PutPage() (bool, error) {
	r := true
	o := "user33333"
	cm := "thi is owsome rice."
	params := PutPageParams{
		PageName: "新ページ",
		Content:  "中身です。/ndddddd",
		Attributes: PutPageAttributes{
			Readonly: &r,
			Author:   &o,
			Comment:  &cm,
		},
	}

	args := []interface{}{params.PageName, params.Content, params.Attributes}
	var reply bool
	if err := w.rpc.Call(wiki_put_page, args, &reply); err != nil {
		return false, nil
	}

	return reply, nil
}

func readPageOptions(options *PageOptions) []interface{} {
	args := []interface{}{}
	if options.PageName != nil {
		args = append(args, *options.PageName)
	}
	if options.Version != nil {
		args = append(args, *options.Version)
	}

	return args
}
