package mets

import (
	"encoding/xml"
	"strings"
	"testing"
)

var marshalTests = []struct {
	Value         interface{}
	ExpectXML     string
	MarshalOnly   bool
	UnmarshalOnly bool
}{
	{
		Value: &Mets{
			Label:   "ALabel",
			Profile: "AProfile",
			ID:      "AnID",
			ObjID:   "12345",
		},
		ExpectXML: `<mets xmlns="http://www.loc.gov/METS/"` +
			` ID="AnID" OBJID="12345" LABEL="ALabel" PROFILE="AProfile">` +
			`</mets>`,
	},
}

func TestMarshal(t *testing.T) {
	for idx, test := range marshalTests {
		if test.UnmarshalOnly {
			continue
		}
		data, err := xml.Marshal(test.Value)
		if err != nil {
			t.Errorf("#%d: Error: %s", idx, err)
			continue
		}

		if got, want := string(data), test.ExpectXML; got != want {
			if strings.Contains(want, "\n") {
				t.Errorf("#%d: marshal(%#v):\nHAVE:\n%s\nWANT:\n%s", idx, test.Value, got, want)
			} else {
				t.Errorf("#%d: marshal(%#v):\nhave %#q\nwant %#q", idx, test.Value, got, want)
			}
		}
	}
}
