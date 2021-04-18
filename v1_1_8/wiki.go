package v1_1_8

import (
	"encoding/base64"
	"errors"
	"time"

	tracrpc "github.com/f-velka/go-trac-rpc"
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
	wiki_get_attachment            string = "wiki.getAttachment"
	wiki_put_attachment            string = "wiki.putAttachment"
	wiki_put_attachment_ex         string = "wiki.putAttachmentEx"
	wiki_delete_page               string = "wiki.deletePage"
	wiki_delete_attachment         string = "wiki.deleteAttachment"
	wiki_list_links                string = "wiki.listLinks"
	wiki_wiki_to_html              string = "wiki.wikiToHtml"
)

// WikiService represents wiki API service.
type WikiService struct {
	rpc tracrpc.RpcClient
}

// PutPageAttributes represents attributes of wiki.putPage.
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

// newWikiService creates new WikiService instance.
func newWikiService(rpc tracrpc.RpcClient) (*WikiService, error) {
	if rpc == nil {
		return nil, errors.New("rpc client cannot be nil")
	}

	return &WikiService{
		rpc: rpc,
	}, nil
}

// GetRecentChanges calls wiki.getRecentChanges.
func (w *WikiService) GetRecentChanges(since *time.Time) ([]PageInfo, error) {
	args := packArgs(since)
	reply := []PageInfo{}
	if err := w.rpc.Call(wiki_get_recent_changes, args, &reply); err != nil {
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
func (w *WikiService) GetPage(pagename *string, version *int) (string, error) {
	args := packArgs(pagename, version)
	var reply string
	if err := w.rpc.Call(wiki_get_page, args, &reply); err != nil {
		return "", err
	}

	return reply, nil
}

// GetPageVersion calls wiki.getPageVersion.
func (w *WikiService) GetPageVersion(pagename *string, version *int) (string, error) {
	args := packArgs(pagename, version)
	var reply string
	if err := w.rpc.Call(wiki_get_page_version, args, &reply); err != nil {
		return "", err
	}

	return reply, nil
}

// GetPageHTML calls wiki.getPageHTML.
func (w *WikiService) GetPageHTML(pagename *string, version *int) (string, error) {
	args := packArgs(pagename, version)
	var reply string
	if err := w.rpc.Call(wiki_get_page_html, args, &reply); err != nil {
		return "", err
	}

	return reply, nil
}

// GetPageHTMLVersion calls wiki.getPageHTMLVersion.
func (w *WikiService) GetPageHTMLVersion(pagename *string, version *int) (string, error) {
	args := packArgs(pagename, version)
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
func (w *WikiService) GetPageInfo(pagename *string, version *int) (PageInfo, error) {
	args := packArgs(pagename, version)
	var reply PageInfo
	if err := w.rpc.Call(wiki_get_page_info, args, &reply); err != nil {
		return PageInfo{}, err
	}

	return reply, nil
}

// GetPageInfoVersion calls wiki.getPageInfoVersion.
func (w *WikiService) GetPageInfoVersion(pagename *string, version *int) (PageInfo, error) {
	args := packArgs(pagename, version)
	var reply PageInfo
	if err := w.rpc.Call(wiki_get_page_info_version, args, &reply); err != nil {
		return PageInfo{}, err
	}

	return reply, nil
}

// PutPage calls wiki.putPage.
func (w *WikiService) PutPage(pagename *string, content *string, attributes PutPageAttributes) (bool, error) {
	args := packArgs(pagename, content, &attributes)
	var reply bool
	if err := w.rpc.Call(wiki_put_page, args, &reply); err != nil {
		return false, nil
	}

	return reply, nil
}

// ListAttachments calls wiki.listAttachments.
func (w *WikiService) ListAttachments(pagename *string) ([]string, error) {
	args := packArgs(pagename)
	var reply []string
	if err := w.rpc.Call(wiki_list_attachments, args, &reply); err != nil {
		return nil, err
	}

	return reply, nil
}

// GetAttachment calls wiki.getAttachment.
func (w *WikiService) GetAttachment(path *string) ([]byte, error) {
	args := packArgs(path)
	var replyBase64 string
	if err := w.rpc.Call(wiki_get_attachment, args, &replyBase64); err != nil {
		return nil, err
	}

	reply, err := base64.StdEncoding.DecodeString(replyBase64)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

// PutAttachment calls wiki.putAttachment.
func (w *WikiService) PutAttachment(path *string, data []byte) (bool, error) {
	encData := base64String(base64.StdEncoding.EncodeToString(data))
	args := packArgs(path, &encData)
	var reply bool
	if err := w.rpc.Call(wiki_put_attachment, args, &reply); err != nil {
		return false, err
	}

	return reply, nil
}

// PutAttachmentEx calls wiki.putAttachmentEx.
// NOTE: This API returns the filename of the created attachment, not boolean as described in the reference.
func (w *WikiService) PutAttachmentEx(pagename *string, filename *string, description *string, data []byte, replace *bool) (string, error) {
	encData := base64String(base64.StdEncoding.EncodeToString(data))
	args := packArgs(pagename, filename, description, &encData, replace)
	var reply string
	if err := w.rpc.Call(wiki_put_attachment_ex, args, &reply); err != nil {
		return "", err
	}

	return reply, nil
}

// DeletePage calls wiki.deletePage.
func (w *WikiService) DeletePage(pagename *string, version *int) (bool, error) {
	args := packArgs(pagename, version)
	var reply bool
	if err := w.rpc.Call(wiki_delete_page, args, &reply); err != nil {
		return false, err
	}

	return reply, nil
}

// DeletePage calls wiki.deleteAttachment.
func (w *WikiService) DeleteAttachment(path *string) (bool, error) {
	args := packArgs(path)
	var reply bool
	if err := w.rpc.Call(wiki_delete_attachment, args, &reply); err != nil {
		return false, err
	}

	return reply, nil
}

// ListLinks calls wiki.listLinks.
// Unfortunately, this API is not implemented yet.
func (w *WikiService) ListLinks(pagename *string) (bool, error) {
	return false, errors.New("wiki.listLinks is not implemented")
	// args := packArgs(pagename)
	// var reply bool
	// if err := w.rpc.Call(wiki_list_links, args, &reply); err != nil {
	// 	return false, err
	// }

	// return reply, nil
}

// WikiToHtml calls wiki.wikiToHtml.
func (w *WikiService) WikiToHtml(text *string) (string, error) {
	args := packArgs(text)
	var reply string
	if err := w.rpc.Call(wiki_wiki_to_html, args, &reply); err != nil {
		return "", err
	}

	return reply, nil
}
