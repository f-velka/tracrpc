package tracrpc

import (
	"encoding/base64"
	"reflect"
	"testing"
	"time"
)

func TestNewWikiService(t *testing.T) {
	tests := []struct {
		name      string
		rpcClient RpcClient
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
			_, err := newWikiService(tt.rpcClient)
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
		since    *time.Time
		reply    string
		expected []PageInfo
	}{
		Time(time.Now()),
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

	c := NewTestClient(wiki_get_recent_changes, packArgs(test.since), test.reply)
	res, err := c.Wiki.GetRecentChanges(test.since)
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

	c := NewTestClient(wiki_get_rpc_version_supported, nil, test.reply)
	res, err := c.Wiki.GetRPCVersionSupported()
	if err != nil {
		t.Fatal(err)
	}
	if res != test.expected {
		t.Fatalf("unexpected result. expected=%v, got=%v", test.expected, res)
	}
}

func TestGetPage(t *testing.T) {
	test := struct {
		pagename *string
		version  *int
		reply    string
		expected string
	}{
		String("shiga"),
		Int(1),
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

	c := NewTestClient(wiki_get_page, packArgs(test.pagename, test.version), test.reply)
	res, err := c.Wiki.GetPage(test.pagename, test.version)
	if err != nil {
		t.Fatal(err)
	}
	if res != test.expected {
		t.Fatalf("unexpected result. expected=%v, got=%v", test.expected, res)
	}
}

func TestGetPageVersion(t *testing.T) {
	test := struct {
		pagename *string
		version  *int
		reply    string
		expected string
	}{
		String("shiga"),
		Int(1),
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

	c := NewTestClient(wiki_get_page_version, packArgs(test.pagename, test.version), test.reply)
	res, err := c.Wiki.GetPageVersion(test.pagename, test.version)
	if err != nil {
		t.Fatal(err)
	}
	if res != test.expected {
		t.Fatalf("unexpected result. expected=%v, got=%v", test.expected, res)
	}
}

func TestGetPageHtml(t *testing.T) {
	test := struct {
		pagename *string
		version  *int
		reply    string
		expected string
	}{
		String("shiga"),
		Int(1),
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

	c := NewTestClient(wiki_get_page_html, packArgs(test.pagename, test.version), test.reply)
	res, err := c.Wiki.GetPageHTML(test.pagename, test.version)
	if err != nil {
		t.Fatal(err)
	}
	if res != test.expected {
		t.Fatalf("unexpected result. expected=%v, got=%v", test.expected, res)
	}
}

func TestGetPageHtmlVersion(t *testing.T) {
	test := struct {
		pagename *string
		version  *int
		reply    string
		expected string
	}{
		String("shiga"),
		Int(1),
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

	c := NewTestClient(wiki_get_page_html_version, packArgs(test.pagename, test.version), test.reply)
	res, err := c.Wiki.GetPageHTMLVersion(test.pagename, test.version)
	if err != nil {
		t.Fatal(err)
	}
	if res != test.expected {
		t.Fatalf("unexpected result. expected=%v, got=%v", test.expected, res)
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

	c := NewTestClient(wiki_get_all_pages, nil, test.reply)
	res, err := c.Wiki.GetAllPages()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(res, test.expected) {
		t.Fatalf("unexpected result. expected=%v, got=%v", test.expected, res)
	}
}

func TestGetPageInfo(t *testing.T) {
	test := struct {
		pagename *string
		version  *int
		reply    string
		expected PageInfo
	}{
		String("shiga"),
		Int(1),
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

	c := NewTestClient(wiki_get_page_info, packArgs(test.pagename, test.version), test.reply)
	res, err := c.Wiki.GetPageInfo(test.pagename, test.version)
	if err != nil {
		t.Fatal(err)
	}
	if res != test.expected {
		t.Fatalf("unexpected result. expected=%v, got=%v", test.expected, res)
	}
}

func TestGetPageInfoVersion(t *testing.T) {
	test := struct {
		pagename *string
		version  *int
		reply    string
		expected PageInfo
	}{
		String("shiga"),
		Int(1),
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

	c := NewTestClient(wiki_get_page_info_version, packArgs(test.pagename, test.version), test.reply)
	res, err := c.Wiki.GetPageInfoVersion(test.pagename, test.version)
	if err != nil {
		t.Fatal(err)
	}
	if res != test.expected {
		t.Fatalf("unexpected result. expected=%v, got=%v", test.expected, res)
	}
}

func TestPutPage(t *testing.T) {
	test := struct {
		pagename   *string
		content    *string
		attributes PutPageAttributes
		reply      string
		expected   bool
	}{
		String("shiga"),
		String("content"),
		PutPageAttributes{
			Readonly: Bool(true),
			Author:   String("murasakishikibu"),
			Comment:  String("comment"),
		},
		`<?xml version='1.0'?>
<methodResponse>
<params>
<param>
<value><boolean>1</boolean></value>
</param>
</params>
</methodResponse>`,
		true,
	}

	c := NewTestClient(wiki_put_page, packArgs(test.pagename, test.content, &test.attributes), test.reply)
	res, err := c.Wiki.PutPage(test.pagename, test.content, test.attributes)
	if err != nil {
		t.Fatal(err)
	}
	if res != test.expected {
		t.Fatalf("unexpected result. expected=%v, got=%v", test.expected, res)
	}
}

func TestListAttachments(t *testing.T) {
	test := struct {
		pagename *string
		reply    string
		expected []string
	}{
		String("WikiTest"),
		`<?xml version='1.0'?>
<methodResponse>
<params>
<param>
<value><array><data>
<value><string>WikiTest/Otsu.txt</string></value>
<value><string>WikiTest/Kyoto.txt</string></value>
</data></array></value>
</param>
</params>
</methodResponse>`,
		[]string{"WikiTest/Otsu.txt", "WikiTest/Kyoto.txt"},
	}

	c := NewTestClient(wiki_list_attachments, packArgs(test.pagename), test.reply)
	res, err := c.Wiki.ListAttachments(test.pagename)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(res, test.expected) {
		t.Fatalf("unexpected result. expected=%v, got=%v", test.expected, res)
	}
}

func TestGetAttachment(t *testing.T) {
	expected, _ := base64.StdEncoding.DecodeString("滋賀")
	test := struct {
		path     *string
		reply    string
		expected []byte
	}{
		String("WikiTest/Otsu.txt"),
		`<?xml version='1.0'?>
<methodResponse>
<params>
<param>
<value><base64>5ruL6LOA</base64></value>
</param>
</params>
</methodResponse>`,
		expected,
	}

	c := NewTestClient(wiki_get_attachment, packArgs(test.path), test.reply)
	res, err := c.Wiki.GetAttachment(test.path)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(res, test.expected) {
		t.Fatalf("unexpected result. expected=%v, got=%v", test.expected, res)
	}
}

func TestPutAttachment(t *testing.T) {
	test := struct {
		path     *string
		data     []byte
		reply    string
		expected bool
	}{
		String("WikiTest/Shiga.txt"),
		[]byte{},
		`<?xml version='1.0'?>
<methodResponse>
<params>
<param>
<value><boolean>1</boolean></value>
</param>
</params>
</methodResponse>`,
		true,
	}

	encData := base64String(test.data)
	c := NewTestClient(wiki_put_attachment, packArgs(test.path, &encData), test.reply)
	res, err := c.Wiki.PutAttachment(test.path, test.data)
	if err != nil {
		t.Fatal(err)
	}
	if res != test.expected {
		t.Fatalf("unexpected result. expected=%v, got=%v", test.expected, res)
	}
}

func TestPutAttachmentEx(t *testing.T) {
	test := struct {
		pagename    *string
		filename    *string
		description *string
		data        []byte
		replace     *bool
		reply       string
		expected    string
	}{
		String("WikiTest"),
		String("Shiga.txt"),
		String("test desc"),
		[]byte{},
		Bool(true),
		`<?xml version='1.0'?>
<methodResponse>
<params>
<param>
<value><string>Shiga.txt</string></value>
</param>
</params>
</methodResponse>`,
		"Shiga.txt",
	}

	encData := base64String(test.data)
	c := NewTestClient(wiki_put_attachment_ex, packArgs(test.pagename, test.filename, test.description, &encData, test.replace), test.reply)
	res, err := c.Wiki.PutAttachmentEx(test.pagename, test.filename, test.description, test.data, test.replace)
	if err != nil {
		t.Fatal(err)
	}
	if res != test.expected {
		t.Fatalf("unexpected result. expected=%v, got=%v", test.expected, res)
	}
}

func TestDeletePage(t *testing.T) {
	test := struct {
		pagename *string
		version  *int
		reply    string
		expected bool
	}{
		String("WikiTest"),
		Int(1),
		`<?xml version='1.0'?>
<methodResponse>
<params>
<param>
<value><boolean>1</boolean></value>
</param>
</params>
</methodResponse>`,
		true,
	}

	c := NewTestClient(wiki_delete_page, packArgs(test.pagename, test.version), test.reply)
	res, err := c.Wiki.DeletePage(test.pagename, test.version)
	if err != nil {
		t.Fatal(err)
	}
	if res != test.expected {
		t.Fatalf("unexpected result. expected=%v, got=%v", test.expected, res)
	}
}

func TestDeleteAttachment(t *testing.T) {
	test := struct {
		path     *string
		reply    string
		expected bool
	}{
		String("WikiTest/Shiga.txt"),
		`<?xml version='1.0'?>
<methodResponse>
<params>
<param>
<value><boolean>1</boolean></value>
</param>
</params>
</methodResponse>`,
		true,
	}

	c := NewTestClient(wiki_delete_attachment, packArgs(test.path), test.reply)
	res, err := c.Wiki.DeleteAttachment(test.path)
	if err != nil {
		t.Fatal(err)
	}
	if res != test.expected {
		t.Fatalf("unexpected result. expected=%v, got=%v", test.expected, res)
	}
}

func TestListLinks(t *testing.T) {
	// this API is not implemented.
}

func TestWikiToHtml(t *testing.T) {
	test := struct {
		text     *string
		reply    string
		expected string
	}{
		String("Test"),
		`<?xml version='1.0'?>
<methodResponse>
<params>
<param>
<value><string>&lt;p&gt;
Test
&lt;/p&gt;
</string></value>
</param>
</params>
</methodResponse>`,
		`<p>
Test
</p>
`,
	}

	c := NewTestClient(wiki_wiki_to_html, packArgs(test.text), test.reply)
	res, err := c.Wiki.WikiToHtml(test.text)
	if err != nil {
		t.Fatal(err)
	}
	if res != test.expected {
		t.Fatalf("unexpected result. expected=%v, got=%v", test.expected, res)
	}
}
