package lido

import (
	"github.com/verisart/xsd/xsdt"
)

type InscriptionsWrap struct {
	Inscriptions []*Inscription `xml:"http://www.lido-schema.org inscriptions"`
}

type Inscription struct {
	// Wrapper for a description of the inscription, including description
	// identifer, descriptive note of the inscription and sources.
	InscriptionDescriptions []*DescriptiveNote `xml:"http://www.lido-schema.org inscriptionDescription"`

	// Assigns a priority order for online presentation of the element. Has to be
	// a positive integer, with descending priority from 1 to x.
	SortOrder xsdt.Integer `xml:"http://www.lido-schema.org sortorder,attr,omitempty"`

	Type xsdt.String `xml:"http://www.lido-schema.org type,attr,omitempty"`

	// Transcription of the inscription. Repeat this element only for language
	// variants.
	InscriptionTranscriptions []*Text `xml:"http://www.lido-schema.org inscriptionTranscription"`
}
