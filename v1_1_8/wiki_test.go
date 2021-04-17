package v1_1_8

import (
	"reflect"
	"testing"
	"time"

	"github.com/f-velka/go-trac-rpc/common"
)

func TestNewWikiService(t *testing.T) {
	tests := []struct {
		name      string
		rpcClient common.RpcClient
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
			_, err := NewWikiService(tt.rpcClient)
			if err != nil {
				if tt.wantErr {
					return
				}
				t.Fatal(err)
			}
		})
	}
}

func TestGetRecentChanges(t *testing.T) {
	test := struct {
		reply    string
		expected []PageInfo
	}{
		`<?xml version='1.0'?>
<methodResponse>
<params>
<param>
<value><array><data>
<value><struct>
<member>
<name>comment</name>
<value><string>ouch</string></value>
</member>
<member>
<name>lastModified</name>
<value><dateTime.iso8601>18600324T00:00:00</dateTime.iso8601></value>
</member>
<member>
<name>version</name>
<value><int>1</int></value>
</member>
<member>
<name>name</name>
<value><string>sakuradamon</string></value>
</member>
<member>
<name>author</name>
<value><string>n_ii</string></value>
</member>
</struct></value>
<value><struct>
<member>
<name>comment</name>
<value><string>nobu</string></value>
</member>
<member>
<name>lastModified</name>
<value><dateTime.iso8601>18671109T00:00:00</dateTime.iso8601></value>
</member>
<member>
<name>version</name>
<value><int>2</int></value>
</member>
<member>
<name>name</name>
<value><string>edo</string></value>
</member>
<member>
<name>author</name>
<value><string>yoshi</string></value>
</member>
</struct></value>
</data></array></value>
</param>
</params>
</methodResponse>`,
		[]PageInfo{
			{
				Name:         "sakuradamon",
				LastModified: time.Date(1860, time.March, 24, 0, 0, 0, 0, time.UTC),
				Author:       "n_ii",
				Version:      1,
				Comment:      "ouch",
			},
			{
				Name:         "edo",
				LastModified: time.Date(1867, time.November, 9, 0, 0, 0, 0, time.UTC),
				Author:       "yoshi",
				Version:      2,
				Comment:      "nobu",
			},
		},
	}

	c := NewTestClient(wiki_get_recent_changes, test.reply)
	res, err := c.Wiki.GetRecentChanges(time.Now())
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(res, test.expected) {
		t.Fatalf("unexpected result. expected=%v, got=%v", test.expected, res)
	}
}

func TestGetRPCVersionSupported(t *testing.T) {
	test := struct {
		reply    string
		expected int
	}{
		`<?xml version='1.0'?>
<methodResponse>
<params>
<param>
<value><int>2</int></value>
</param>
</params>
</methodResponse>`,
		2,
	}

	c := NewTestClient(wiki_get_rpc_version_supported, test.reply)
	res, err := c.Wiki.GetRPCVersionSupported()
	if err != nil {
		t.Fatal(err)
	}
	if res != test.expected {
		t.Fatalf("unexpected result. expected=%d, got=%d", test.expected, res)
	}
}

func TestGetPage(t *testing.T) {

	test := struct {
		options  PageOptions
		reply    string
		expected string
	}{
		PageOptions{
			PageName: PStr("Shiga"),
			Version:  PInt(1),
		},
		`<?xml version='1.0'?>
<methodResponse>
<params>
<param>
<value><string>Shiga</string></value>
</param>
</params>
</methodResponse>`,

		"Shiga",
	}

	c := NewTestClient(wiki_get_page, test.reply)
	res, err := c.Wiki.GetPage(&test.options)
	if err != nil {
		t.Fatal(err)
	}
	if res != test.expected {
		t.Fatalf("unexpected result. expected=%s, got=%s", test.expected, res)
	}
}

func TestGetPageVersion(t *testing.T) {
	test := struct {
		options  PageOptions
		reply    string
		expected string
	}{
		PageOptions{
			PageName: PStr("Shiga"),
			Version:  PInt(1),
		},
		`<?xml version='1.0'?>
<methodResponse>
<params>
<param>
<value><string>Shiga</string></value>
</param>
</params>
</methodResponse>`,

		"Shiga",
	}

	c := NewTestClient(wiki_get_page_version, test.reply)
	res, err := c.Wiki.GetPageVersion(&test.options)
	if err != nil {
		t.Fatal(err)
	}
	if res != test.expected {
		t.Fatalf("unexpected result. expected=%s, got=%s", test.expected, res)
	}
}

func TestGetPageHtml(t *testing.T) {
	test := struct {
		options  PageOptions
		reply    string
		expected string
	}{
		PageOptions{
			PageName: PStr("this is a test page."),
			Version:  PInt(1),
		},
		`<?xml version='1.0'?>
<methodResponse>
<params>
<param>
<value><string>this is a test page.</string></value>
</param>
</params>
</methodResponse>`,

		"this is a test page.",
	}

	c := NewTestClient(wiki_get_page_html, test.reply)
	res, err := c.Wiki.GetPageHTML(&test.options)
	if err != nil {
		t.Fatal(err)
	}
	if res != test.expected {
		t.Fatalf("unexpected result. expected=%s, got=%s", test.expected, res)
	}
}

func TestGetPageHtmlVersion(t *testing.T) {
	test := struct {
		options  PageOptions
		reply    string
		expected string
	}{
		PageOptions{
			PageName: PStr("this is a test page."),
			Version:  PInt(1),
		},
		`<?xml version='1.0'?>
<methodResponse>
<params>
<param>
<value><string>this is a test page.</string></value>
</param>
</params>
</methodResponse>`,

		"this is a test page.",
	}

	c := NewTestClient(wiki_get_page_html_version, test.reply)
	res, err := c.Wiki.GetPageHTMLVersion(&test.options)
	if err != nil {
		t.Fatal(err)
	}
	if res != test.expected {
		t.Fatalf("unexpected result. expected=%s, got=%s", test.expected, res)
	}
}

func TestGetAllPages(t *testing.T) {
	test := struct {
		reply    string
		expected []string
	}{
		`<?xml version='1.0'?>
<methodResponse>
<params>
<param>
<value><array><data>
<value><string>Biwako</string></value>
<value><string>Kasumigaura</string></value>
</data></array></value>
</param>
</params>
</methodResponse>`,
		[]string{"Biwako", "Kasumigaura"},
	}

	c := NewTestClient(wiki_get_all_pages, test.reply)
	res, err := c.Wiki.GetAllPages()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(res, test.expected) {
		t.Fatalf("unexpected result. expected=%s, got=%s", test.expected, res)
	}
}

func TestGetPageInfo(t *testing.T) {
	test := struct {
		options  PageOptions
		reply    string
		expected PageInfo
	}{
		PageOptions{
			PageName: PStr("sakuradamon"),
			Version:  PInt(1),
		},
		`<?xml version='1.0'?>
<methodResponse>
<params>
<param>
<value><struct>
<member>
<name>comment</name>
<value><string>ouch</string></value>
</member>
<member>
<name>lastModified</name>
<value><dateTime.iso8601>18600324T00:00:00</dateTime.iso8601></value>
</member>
<member>
<name>version</name>
<value><int>1</int></value>
</member>
<member>
<name>name</name>
<value><string>sakuradamon</string></value>
</member>
<member>
<name>author</name>
<value><string>n_ii</string></value>
</member>
</struct></value>
</param>
</params>
</methodResponse>`,
		PageInfo{
			Name:         "sakuradamon",
			LastModified: time.Date(1860, time.March, 24, 0, 0, 0, 0, time.UTC),
			Author:       "n_ii",
			Version:      1,
			Comment:      "ouch",
		},
	}
	c := NewTestClient(wiki_get_page_info, test.reply)
	res, err := c.Wiki.GetPageInfo(&test.options)
	if err != nil {
		t.Fatal(err)
	}
	if res != test.expected {
		t.Fatalf("unexpected result. expected=%v, got=%v", test.expected, res)
	}
}

func TestGetPageInfoVersion(t *testing.T) {
	test := struct {
		options  PageOptions
		reply    string
		expected PageInfo
	}{
		PageOptions{
			PageName: PStr("sakuradamon"),
			Version:  PInt(1),
		},
		`<?xml version='1.0'?>
<methodResponse>
<params>
<param>
<value><struct>
<member>
<name>comment</name>
<value><string>ouch</string></value>
</member>
<member>
<name>lastModified</name>
<value><dateTime.iso8601>18600324T00:00:00</dateTime.iso8601></value>
</member>
<member>
<name>version</name>
<value><int>1</int></value>
</member>
<member>
<name>name</name>
<value><string>sakuradamon</string></value>
</member>
<member>
<name>author</name>
<value><string>n_ii</string></value>
</member>
</struct></value>
</param>
</params>
</methodResponse>`,
		PageInfo{
			Name:         "sakuradamon",
			LastModified: time.Date(1860, time.March, 24, 0, 0, 0, 0, time.UTC),
			Author:       "n_ii",
			Version:      1,
			Comment:      "ouch",
		},
	}
	c := NewTestClient(wiki_get_page_info_version, test.reply)
	res, err := c.Wiki.GetPageInfoVersion(&test.options)
	if err != nil {
		t.Fatal(err)
	}
	if res != test.expected {
		t.Fatalf("unexpected result. expected=%v, got=%v", test.expected, res)
	}
}

func TestPutPage(t *testing.T) {
	// test := struct{

	// }

	// c := NewTestClient(wiki_put_page, test.reply)
	// res, err := c.Wiki.PutPage(&test.options)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.FailNow()
}

func TestListAttachments(t *testing.T) {
	t.FailNow()
}

func TestGetAttachment(t *testing.T) {
	t.FailNow()
}

func TestPutAttachment(t *testing.T) {
	t.FailNow()
}

func TestPutAttachmentEx(t *testing.T) {
	t.FailNow()
}

func TestDeletePage(t *testing.T) {
	t.FailNow()
}

func TestDeleteAttachment(t *testing.T) {
	t.FailNow()
}

func TestListLinks(t *testing.T) {
	t.FailNow()
}

func TestWikiToHtml(t *testing.T) {
	t.FailNow()
}
