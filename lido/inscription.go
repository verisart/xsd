package lido

import (
	"github.com/verisart/xsd/xsdt"
)

type InscriptionsWrap struct {
	Inscriptions []*Inscription `xml:"inscriptions"`
}

type Inscription struct {
	// Wrapper for a description of the inscription, including description
	// identifer, descriptive note of the inscription and sources.
	InscriptionDescriptions []*DescriptiveNote `xml:"inscriptionDescription"`

	// Assigns a priority order for online presentation of the element. Has to be
	// a positive integer, with descending priority from 1 to x.
	SortOrder xsdt.Integer `xml:"sortorder,attr,omitempty"`

	Type xsdt.String `xml:"type,attr,omitempty"`

	// Transcription of the inscription. Repeat this element only for language
	// variants.
	InscriptionTranscriptions []*Text `xml:"inscriptionTranscription"`
}
