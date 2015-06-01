package lido

import (
	"github.com/verisart/xsd/xsdt"
)

type Appellation struct {
	// Appellations, e.g. titles, identifying phrases, or names given to an item,
	// but also name of a person or corporation, also place name etc. Repeat this
	// element only for language variants.
	Values []*AppellationValue `xml:"appellationValue"`

	// The source for the appellation, generally a published source.
	Sources []*Text `xml:"sourceAppellation"`
}

// Appellations, e.g. titles, identifying phrases, or names given to an item,
// but also name of a person or corporation, also place name etc.
// How to record: Repeat this element only for language variants.
type AppellationValue struct {
	Value xsdt.String `xml:",chardata"`

	// Appellation values are mainly there to store language variants.
	Lang xsdt.Language `xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty"`

	// Qualifies the value as a preferred or alternative variant. Data values:
	// preferred, alternate
	Pref xsdt.String `xml:"pref,attr,omitempty"`

	// How to record: Elements with data values are accompanied by the attributes
	// encodinganalog and label to indicate the format of the data source from
	// which the data were migrated. The attribute encodinganalog refers to the
	// internal field label of the source database. The source format is indicated
	// in the attribute relatedencoding of the lidoWrap
	EncodingAnalog xsdt.String `xml:"encodinganalog,attr,omitempty"`

	// Elements with data values are accompanied by the attributes
	// encodinganalog and label, to indicate the format of the data source from
	// which the data were migrated. The attribute label refers to the external
	// label of a data field at the visible user interface. The source format is
	// indicated in the attribute
	Label xsdt.String `xml:"label,attr,omitempty"`
}

func (ap *Appellation) Set(value string, lang string, pref bool) error {
	// First clear, then append
	ap.Values = ap.Values[:0]
	return ap.Append(value, lang, pref)
}

// Adds a an appellation value
func (ap *Appellation) Append(value string, lang string, pref bool) error {
	ap.Values = append(ap.Values, &AppellationValue{
		Value: ToXsdt(value),
		Lang:  ToLang(lang),
		Pref:  ToPref(pref),
	})

	return nil
}
