package xlink

import (
	xsdt "github.com/verisart/xsd/xsdt"
)

type XLinkHrefAttr struct {
	Href xsdt.AnyURI `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
}

type XLinkRoleAttr struct {
	Role xsdt.String `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
}

type XLinkArcRoleAttr struct {
	Arcrole xsdt.String `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
}

type XLinkTitleAttr struct {
	Title xsdt.String `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
}

type XLinkShowType xsdt.String

//	Since TxsdShow is just a simple String type, this merely returns the current string value.
func (me XLinkShowType) String() string { return xsdt.String(me).String() }

//	This convenience method just performs a simple type conversion to TxsdShow's alias type xsdt.String.
func (me XLinkShowType) ToXsdtString() xsdt.String { return xsdt.String(me) }

//	Returns true if the value of this enumerated TxsdShow is "new".
func (me XLinkShowType) IsNew() bool { return me.String() == "new" }

//	Returns true if the value of this enumerated TxsdShow is "replace".
func (me XLinkShowType) IsReplace() bool { return me.String() == "replace" }

//	Returns true if the value of this enumerated TxsdShow is "embed".
func (me XLinkShowType) IsEmbed() bool { return me.String() == "embed" }

//	Returns true if the value of this enumerated TxsdShow is "other".
func (me XLinkShowType) IsOther() bool { return me.String() == "other" }

//	Returns true if the value of this enumerated TxsdShow is "none".
func (me XLinkShowType) IsNone() bool { return me.String() == "none" }

//	Since TxsdShow is just a simple String type, this merely sets the current value from the specified string.
func (me *XLinkShowType) Set(s string) { (*xsdt.String)(me).Set(s) }

type XLinkShowAttr struct {
	Show XLinkShowType `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
}

type XLinkActuateType xsdt.String

//	This convenience method just performs a simple type conversion to TxsdActuate's alias type xsdt.String.
func (me XLinkActuateType) ToXsdtString() xsdt.String { return xsdt.String(me) }

//	Returns true if the value of this enumerated TxsdActuate is "onLoad".
func (me XLinkActuateType) IsOnLoad() bool { return me.String() == "onLoad" }

//	Returns true if the value of this enumerated TxsdActuate is "onRequest".
func (me XLinkActuateType) IsOnRequest() bool { return me.String() == "onRequest" }

//	Returns true if the value of this enumerated TxsdActuate is "other".
func (me XLinkActuateType) IsOther() bool { return me.String() == "other" }

//	Returns true if the value of this enumerated TxsdActuate is "none".
func (me XLinkActuateType) IsNone() bool { return me.String() == "none" }

//	Since TxsdActuate is just a simple String type, this merely sets the current value from the specified string.
func (me *XLinkActuateType) Set(s string) { (*xsdt.String)(me).Set(s) }

//	Since TxsdActuate is just a simple String type, this merely returns the current string value.
func (me XLinkActuateType) String() string { return xsdt.String(me).String() }

type XLinkActuateAttr struct {
	Actuate XLinkActuateType `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type XLinkLabelAttr struct {
	Label xsdt.String `xml:"http://www.w3.org/1999/xlink label,attr,omitempty"`
}

type XLinkFromAttr struct {
	From xsdt.String `xml:"http://www.w3.org/1999/xlink from,attr,omitempty"`
}

type XLinkToAttr struct {
	To xsdt.String `xml:"http://www.w3.org/1999/xlink to,attr,omitempty"`
}

type XLinkTypeAttr struct {
	Type xsdt.String `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
}

//	Returns the fixed value for Type -- "simple"
func (me XLinkTypeAttr) TypeFixed() xsdt.String {
	return xsdt.String("simple")
}

type SimpleLink struct {
	XLinkRoleAttr

	XLinkArcRoleAttr

	XLinkTitleAttr

	XLinkShowAttr

	XLinkActuateAttr

	XLinkTypeAttr

	XLinkHrefAttr
}

type XLinkTypeExtendedAttr struct {
	Type xsdt.String `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
}

//	Returns the fixed value for Type -- "extended"
func (me XLinkTypeExtendedAttr) TypeFixed() xsdt.String {
	return xsdt.String("extended")
}

type ExtendedLink struct {
	XLinkTitleAttr

	XLinkTypeExtendedAttr

	XLinkRoleAttr
}

type XLinkTypeLocatorAttr struct {
	Type xsdt.String `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
}

//	Returns the fixed value for Type -- "locator"
func (me XLinkTypeLocatorAttr) TypeFixed() xsdt.String {
	return xsdt.String("locator")
}

type LocatorLink struct {
	XLinkTypeLocatorAttr

	XLinkHrefAttr

	XLinkRoleAttr

	XLinkTitleAttr

	XLinkLabelAttr
}

type XLinkTypeArcAttr struct {
	Type xsdt.String `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
}

//	Returns the fixed value for Type -- "arc"
func (me XLinkTypeArcAttr) TypeFixed() xsdt.String { return xsdt.String("arc") }

type ArcLink struct {
	XLinkActuateAttr

	XLinkFromAttr

	XLinkToAttr

	XLinkTypeArcAttr

	XLinkArcRoleAttr

	XLinkTitleAttr

	XLinkShowAttr
}

type XLinkTypeResourceAttr struct {
	Type xsdt.String `xml:"http://www.w3.org/1999/xlink type,attr"`
}

//	Returns the fixed value for Type -- "resource"
func (me XLinkTypeResourceAttr) TypeFixed() xsdt.String {
	return xsdt.String("resource")
}

type ResourceLink struct {
	XLinkTypeResourceAttr

	XLinkRoleAttr

	XLinkTitleAttr

	XLinkLabelAttr
}

type XLinkTypeTitleAttr struct {
	Type xsdt.String `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
}

//	Returns the fixed value for Type -- "title"
func (me XLinkTypeTitleAttr) TypeFixed() xsdt.String { return xsdt.String("title") }

type TitleLink struct {
	XLinkTypeTitleAttr
}

type XLinkTypeNoneAttr struct {
	Type xsdt.String `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
}

//	Returns the fixed value for Type -- "none"
func (me XLinkTypeNoneAttr) TypeFixed() xsdt.String { return xsdt.String("none") }

type EmptyLink struct {
	XLinkTypeNoneAttr
}
