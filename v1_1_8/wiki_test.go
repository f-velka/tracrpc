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
		expected []GetRecentChangesResult
	}{
		`<?xml version='1.0'?>
<methodResponse>
<params>
<param>
<value><array><data>
<value><struct>
<member>
<name>comment</name>
<value><string>comment1</string></value>
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
<value><string>name1</string></value>
</member>
<member>
<name>author</name>
<value><string>author1</string></value>
</member>
</struct></value>
<value><struct>
<member>
<name>comment</name>
<value><string>comment2</string></value>
</member>
<member>
<name>lastModified</name>
<value><dateTime.iso8601>19360226T00:00:00</dateTime.iso8601></value>
</member>
<member>
<name>version</name>
<value><int>2</int></value>
</member>
<member>
<name>name</name>
<value><string>name2</string></value>
</member>
<member>
<name>author</name>
<value><string>author2</string></value>
</member>
</struct></value>
</data></array></value>
</param>
</params>
</methodResponse>`,
		[]GetRecentChangesResult{
			{
				Name:         "name1",
				LastModified: time.Date(1860, time.March, 24, 0, 0, 0, 0, time.UTC),
				Author:       "author1",
				Version:      1,
				Comment:      "comment1",
			},
			{
				Name:         "name2",
				LastModified: time.Date(1936, time.February, 26, 0, 0, 0, 0, time.UTC),
				Author:       "author2",
				Version:      2,
				Comment:      "comment2",
			},
		},
	}

	c := NewTestClient(test.reply)
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

	c := NewTestClient(test.reply)
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
		name     string
		options  GetPageOptions
		reply    string
		expected string
	}{
		"OK",
		GetPageOptions{
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

	c := NewTestClient(test.reply)
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
		name     string
		options  GetPageOptions
		reply    string
		expected string
	}{
		"OK",
		GetPageOptions{
			PageName: PStr("近江"),
			Version:  PInt(1),
		},
		`<?xml version='1.0'?>
<methodResponse>
<params>
<param>
<value><string>近江</string></value>
</param>
</params>
</methodResponse>`,

		"近江",
	}

	c := NewTestClient(test.reply)
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
		name     string
		options  GetPageOptions
		reply    string
		expected string
	}{
		"OK",
		GetPageOptions{
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

	c := NewTestClient(test.reply)
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
		name     string
		options  GetPageOptions
		reply    string
		expected string
	}{
		"OK",
		GetPageOptions{
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

	c := NewTestClient(test.reply)
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

	c := NewTestClient(test.reply)
	res, err := c.Wiki.GetAllPages()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(res, test.expected) {
		t.Fatalf("unexpected result. expected=%s, got=%s", test.expected, res)
	}
}

func TestGetPageInfo(t *testing.T) {
	t.FailNow()
}

func TestGetPageInfoVersion(t *testing.T) {
	t.FailNow()
}

func TestPutPage(t *testing.T) {
	t.FailNow()
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

func TestReadGetPageOptions(t *testing.T) {
	pageName := "testPageName"
	version := 1
	tests := []struct {
		name     string
		pageName *string
		version  *int
		expected []interface{}
		wantErr  bool
	}{
		{
			"OK(without Version)",
			&pageName,
			nil,
			[]interface{}{pageName},
			false,
		},
		{
			"OK(with Version)",
			&pageName,
			&version,
			[]interface{}{pageName, version},
			false,
		},
		{
			"NG(no PageName)",
			nil,
			nil,
			nil,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			options := &GetPageOptions{
				PageName: tt.pageName,
				Version:  tt.version,
			}
			res, err := readGetPageOptions(options, "TestReadGetPageOptions")
			if err != nil {
				if tt.wantErr {
					return
				}
				t.Fatal(err)
			}
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("unexpected result. expected=%s, got=%s", tt.expected, res)
			}
		})
	}
}
