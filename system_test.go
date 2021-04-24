package tracrpc

import (
	"reflect"
	"testing"
)

func TestNewSystemService(t *testing.T) {
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
			_, err := newSystemService(tt.rpcClient)
			if err != nil {
				if tt.wantErr {
					return
				}
				t.Fatal(err)
			}
		})
	}
}

func TestListMethods(t *testing.T) {
	test := struct {
		reply    string
		expected []string
	}{
		`<?xml version='1.0'?>
<methodResponse>
<params>
<param>
<value><array><data>
<value><string>system.destroy</string></value>
<value><string>system.messUp</string></value>
</data></array></value>
</param>
</params>
</methodResponse>`,
		[]string{"system.destroy", "system.messUp"},
	}

	c := NewTestClient(system_list_methods, nil, test.reply)
	res, err := c.System.ListMethods()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(res, test.expected) {
		t.Fatalf("unexpected result. expected=%v, got=%v", test.expected, res)
	}
}

func TestMethodHelp(t *testing.T) {
	test := struct {
		methodName *string
		reply      string
		expected   string
	}{
		String("help.help"),
		`<?xml version='1.0'?>
<methodResponse>
<params>
<param>
<value><string>This is very helpful help.</string></value>
</param>
</params>
</methodResponse>`,
		"This is very helpful help.",
	}

	c := NewTestClient(system_method_help, packArgs(test.methodName), test.reply)
	res, err := c.System.MethodHelp(test.methodName)
	if err != nil {
		t.Fatal(err)
	}
	if res != test.expected {
		t.Fatalf("unexpected result. expected=%v, got=%v", test.expected, res)
	}
}

func TestMethodSignature(t *testing.T) {
	test := struct {
		methodName *string
		reply      string
		expected   []string
	}{
		String("wiki.getHelp"),
		`<?xml version='1.0'?>
<methodResponse>
<params>
<param>
<value><array><data>
<value><string>string,string</string></value>
<value><string>string,string,int</string></value>
</data></array></value>
</param>
</params>
</methodResponse>`,
		[]string{"string,string", "string,string,int"},
	}

	c := NewTestClient(system_method_signature, packArgs(test.methodName), test.reply)
	res, err := c.System.MethodSignature(test.methodName)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(res, test.expected) {
		t.Fatalf("unexpected result. expected=%v, got=%v", test.expected, res)
	}
}

func TestGetAPIVersion(t *testing.T) {
	test := struct {
		reply    string
		expected []int
	}{
		`<?xml version='1.0'?>
<methodResponse>
<params>
<param>
<value><array><data>
<value><int>1</int></value>
<value><int>10</int></value>
<value><int>821</int></value>
</data></array></value>
</param>
</params>
</methodResponse>`,
		[]int{1, 10, 821},
	}

	c := NewTestClient(system_get_API_version, nil, test.reply)
	res, err := c.System.GetAPIVersion()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(res, test.expected) {
		t.Fatalf("unexpected result. expected=%v, got=%v", test.expected, res)
	}
}
