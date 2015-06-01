package lido

import (
	"github.com/verisart/xsd/xsdt"
)

// Wrapper for Object name / Title information.
type TitleWrap struct {
	Titles []*Title `xml:"titleSet"`
}

// Wrapper for one title or object name and its source information.
// Mandatory. If there is no specific title, provide an object name in the
// appellation value. If there is more than one title, repeat the Title Set
// element. For objects from natural, technical, cultural history e.g. the
// object name given here and the object type, recorded in the object / work
// type element are often identical.
type Title struct {
	Appellation

	// Assigns a priority order for online presentation of the element. Has to be
	// a positive integer, with descending priority from 1 to x.
	SortOrder xsdt.Integer `xml:"sortorder,attr,omitempty"`

	// Type can be used to specify alternate or preferred i.e. 'Repository Title'
	// or 'Alternate Title'
	Type xsdt.String `xml:"type,attr,omitempty"`
}

func NewTitle(value string, lang string, pref bool, titleType string) (*Title, error) {
	title := &Title{
		Type: ToXsdt(titleType),
	}

	err := title.Set(value, lang, pref)

	if err != nil {
		return nil, err
	}

	return title, nil
}

// Adds a an appellation value to a title. The method is passed the value of
// the title, the language code, and whether this is the preferred variant.
func (title *Title) Set(value string, lang string, pref bool) error {
	return title.Appellation.Set(value, lang, pref)
}

// Adds a an appellation value to a title. The method is passed the value of
// the title, the language code, and whether this is the preferred variant.
func (title *Title) Append(value string, lang string, pref bool) error {
	return title.Appellation.Append(value, lang, pref)
}

// Convenience method for appending a title to a title wrap element.
func (wrap *TitleWrap) Append(title *Title) {
	wrap.Titles = append(wrap.Titles, title)
}
