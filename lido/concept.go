package lido

import (
	"github.com/verisart/cidoccrm"
	"github.com/verisart/xsd/xsdt"
)

// Set for identifiers and terms of a concept. A concept describes a conceptual
// resource. Concepts are organized in concept schemes like thesauri,
// classification schemes, taxonomies, subject-heading systems, or any other
// type of structured controlled vocabulary. See also SKOS specifications at
// http://www.w3.org/2004/02/skos/
type Concept struct {

	// A unique identifier for the concept. Preferably taken from and linking to
	// a published controlled vocabulary.
	ConceptIDs []*Identifier `xml:"conceptID"`

	// A name for the referred concept, used for indexing.
	Terms []*Term `xml:"term"`
}

// A name for a concept / term, usually from a controlled vocabulary.
type Term struct {
	Value xsdt.String `xml:",chardata"`

	Lang xsdt.Language `xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty"`

	// Qualifies the value as a preferred or alternative variant. Data values:
	// preferred, alternate
	Pref xsdt.String `xml:"pref,attr,omitempty"`

	//  How to record: Has the two values: "yes" or "no". ”yes” indicates, that
	// the term is an additional term which is derived from an underlying
	// controlled vocabulary (eg. synonym, generic term, superordinate term) and
	// should be used only for retrieval."no" is default.
	AddedSearchTerm AddedSearchTerm `xml:"addedSearchTerm,attr,omitempty"`

	// How to record: Elements with data values are accompanied by the attributes
	// encodinganalog and label to indicate the format of the data source from
	// which the data were migrated. The attribute encodinganalog refers to the
	// internal field label of the source database. The source format is indicated
	// in the attribute relatedencoding of the lidoWrap
	EncodingAnalog xsdt.String `xml:"encodinganalog,attr,omitempty"`

	// How to record: Elements with data values are accompanied by the attributes
	// encodinganalog and label, to indicate the format of the data source from
	// which the data were migrated. The attribute label refers to the external
	// label of a data field at the visible user interface. The source format is
	// indicated in the attribute
	Label xsdt.String `xml:"label,attr,omitempty"`
}

func NewConcept(conceptID *Identifier, term *Term) *Concept {
	return &Concept{
		ConceptIDs: []*Identifier{conceptID},
		Terms:      []*Term{term},
	}
}

func NewURIConcept(uri string, term string, termLang string) *Concept {
	return NewConcept(
		&Identifier{
			Value: ToXsdt(uri),
			Type:  URIType,
		},
		&Term{
			Value: ToXsdt(term),
			Lang:  ToLang(termLang),
		})
}

func NewCRMConcept(crmID string) (*Concept, error) {
	name, err := cidoccrm.ClassName(crmID)

	if err != nil {
		return nil, err
	}

	return NewURIConcept(cidoccrm.FormatURI(crmID), name, "en"), nil
}

func NewTermConcept(source string, conceptType string, termID string, term string) *Concept {
	return NewConcept(
		&Identifier{
			Value:  ToXsdt(termID),
			Source: ToXsdt(source),
			Type:   ToXsdt(conceptType),
		},
		&Term{
			Value: ToXsdt(term),
		})
}

func NewAATConcept(conceptType string, aatID string, term string) *Concept {
	return NewTermConcept("AAT", conceptType, aatID, term)
}
