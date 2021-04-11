package v1_1_8

import (
	"reflect"
	"testing"
)

func pStr(str string) *string {
	return &str
}

func pInt(i int) *int {
	return &i
}

func TestGetRPCVersionSupported_OK(t *testing.T) {
	expected := 2
	reply := `
<?xml version='1.0'?>
<methodResponse>
<params>
<param>
<value><int>2</int></value>
</param>
</params>
</methodResponse>`

	c := NewTestClient(reply)
	res, err := c.Wiki.GetRPCVersionSupported()
	if err != nil {
		t.Error(err)
		return
	}
	if res != expected {
		t.Errorf("unexpected result. expected=%d, got=%d", expected, res)
		return
	}
}

func TestGetPage(t *testing.T) {

	tests := []struct {
		name     string
		option   GetPageOption
		reply    string
		expected string
		ok       bool
	}{
		{
			"Success(one arg)",
			GetPageOption{
				PageName: pStr("Shiga"),
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
			true},
		{
			"Success(two args)",
			GetPageOption{
				PageName: pStr("Shiga"),
				Version:  pInt(1),
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
			true},
		{
			"Fail(no required arg)",
			GetPageOption{
				Version: pInt(1),
			},
			"",
			"",
			false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewTestClient(tt.reply)
			res, err := c.Wiki.GetPage(&tt.option)
			if err != nil {
				if !tt.ok {
					return
				}
				t.Error(err)
			}
			if res != tt.expected {
				t.Errorf("unexpected result. expected=%s, got=%s", tt.expected, res)
				return
			}
		})
	}
}

func TestGetAllPages_OK(t *testing.T) {
	expected := []string{"Biwako", "Kasumigaura"}
	reply := `
<?xml version='1.0'?>
<methodResponse>
<params>
<param>
<value><array><data>
<value><string>Biwako</string></value>
<value><string>Kasumigaura</string></value>
</data></array></value>
</param>
</params>
</methodResponse>`

	c := NewTestClient(reply)
	res, err := c.Wiki.GetAllPages()
	if err != nil {
		t.Error(err)
		return
	}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("unexpected result. expected=%s, got=%s", expected, res)
		return
	}
}
