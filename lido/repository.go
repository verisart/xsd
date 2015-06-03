package lido

import (
	"github.com/verisart/xsd/xsdt"
)

type RepositoryWrap struct {
	Repositories []*Repository `xml:"http://www.lido-schema.org repositorySet"`
}

type Repository struct {
	// Assigns a priority order for online presentation of the element. Has to be
	// a positive integer, with descending priority from 1 to x.
	SortOrder xsdt.Integer `xml:"http://www.lido-schema.org sortorder,attr,omitempty"`

	//  Definition: Qualifies the repository as a former or the current repository.
	//  How to record: Data values: current, former
	Type xsdt.String `xml:"http://www.lido-schema.org type,attr,omitempty"`

	//  Definition: Unambiguous identification, designation and weblink of the institution of custody.
	RepositoryName *LegalBodyRef `xml:"http://www.lido-schema.org repositoryName"`

	//  Definition: An unambiguous numeric or alphanumeric identification number, assigned to the object by the institution of custody.
	WorkIDs []*WorkID `xml:"http://www.lido-schema.org workID"`

	//  Definition: Location of the object, especially relevant for architecture and archaeological sites.
	RepositoryLocation *Place `xml:"http://www.lido-schema.org repositoryLocation"`
}
