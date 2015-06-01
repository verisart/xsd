package rdf

import (
	//"encoding/xml"
	//xlink "github.com/verisart/go-prov/schema/xlink"
	"github.com/verisart/xsd/xsdt"
)

type ResourceAttr struct {
	Resource xsdt.String `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# resource,attr"`
}

type Type ResourceAttr
type Subject ResourceAttr
type Object ResourceAttr
type Predicate ResourceAttr

type Statement struct {
	Subject   *Subject   `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# subject"`
	Predicate *Predicate `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# predicate"`
	Object    *Object    `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# object"`
}
