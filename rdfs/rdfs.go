package rdfs

import (
	//"encoding/xml"
	//xlink "github.com/verisart/go-prov/schema/xlink"
	"github.com/verisart/xsd/xsdt"
)

type Label struct {
	XsdtString xsdt.String `xml:",chardata"`

	Lang xsdt.Language `xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty"`
}
