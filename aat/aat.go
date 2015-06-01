package aat

import (
	//"encoding/xml"
	//xlink "github.com/verisart/go-prov/schema/xlink"
	"github.com/verisart/xsd/rdf"
	"github.com/verisart/xsd/rdfs"
	"github.com/verisart/xsd/xsdt"
)

// Place defined by administrative boundaries and conditions, including
// inhabited places, nations, and empires. Used in TGN only. Example: Burgundy
// region in France (TGN)
const AdminPlaceConceptURI = "http://vocab.getty.edu/ontology#AdminPlaceConcept"

// Biography of a ULAN agent (schema:Person|Organization). Depending on the
// agent, has these fields: - schema:description: one-line-biography -
// gvp:estStart - schema:birthPlace | foundationLocation, pointing to TGN -
// gvp:estEnd - schema:deathPlace | dissolutionLocation, pointing to TGN -
// schema:gender, pointing to AAT (male, female, other) - dct:contributor The
// fields of the preferred biogrpahy are also available in the agent.
const BiographyURI = "http://vocab.getty.edu/ontology#Biography"

// Proper concept. Used in AAT only; TGN & ULAN have their own, e.g.
// gvp:PhysPlaceConcept & gvp:PersonConcept. Used for indexing and cataloguing.
// Example: rhyta (AAT)
const ConceptURI = "http://vocab.getty.edu/ontology#Concept"

// One of the major divisions of a vocabulary. Example: Objects Facet (AAT),
// World (TGN), Artists (ULAN)
const FacetURI = "http://vocab.getty.edu/ontology#Facet"

// PTwo or more people who generally worked together to collectively create art.
// Not necessarily legally incorporated. A family of artists may be considered a
// "corporate body". Corresponds to crm:E74_Group, not its subclass
// crm:E40_Legal_Body. Example: 500356337 Albrecht Duerer Workshop (ULAN)
const GroupConceptURI = "http://vocab.getty.edu/ontology#GroupConcept"

// Guide Term: place holder to create a level in the hierarchy. Used in AAT and
// ULAN. Not used for indexing or cataloguing.. Example: <vessels for serving
// and consuming food> (AAT), ulan:500353455 <named animals> under facet
// Non-Artists (ULAN)
const GuideTermURI = "http://vocab.getty.edu/ontology#GuideTerm"

// Top of a hierarchy. Used in AAT only. Not used for indexing or cataloguing.
// Example: Containers (Hierarchy Name) (AAT)
const HierarchyURI = "http://vocab.getty.edu/ontology#Hierarchy"

// Obsolete subject: moved out of the publishable hierarchy, or merged to
// another (pointed by dct:isReplacedBy). Example: 300375205 "shranks" (AAT) was
// merged to 300039264 "schranks" (AAT), so aat:300375205 dct:isReplacedBy
// aat:300039264
const ObsoleteSubjectURI = "http://vocab.getty.edu/ontology#ObsoleteSubject"

// A single individual. Usually people engaged in the design or creation of art
// or architecture ("Artists"); but may also include donors, patrons, rulers,
// sitters, art historians, etc ("Non-Artists"). Example: 500115493 Duerer,
// Albrecht (ULAN)
const PersonConceptURI = "http://vocab.getty.edu/ontology#PersonConcept"

// Place that is both administrative and physical. Rarely used. Used in TGN
// only. Example: 6003090 Kiik-Koba (TGN)
const PhysAdminPlaceConcept = "http://vocab.getty.edu/ontology#PhysAdminPlaceConcept"

// Physical feature, defined by its physical characteristics on planet Earth,
// including mountains, rivers, and oceans. Used in TGN only. Example: Amazon River (TGN)
const PhysPlaceConcept = "http://vocab.getty.edu/ontology#PhysPlaceConcept"

// Defines a GVP subject or provides usage information. Has fields:
// -  dc:identifier: numeric ID, also used in the URL.
// - rdf:value: the note itself (as per SKOS Primer: 4.2 Advanced Documentation
// Features) with language tag
// - dct:language: matches the language tag
// - gvp:displayOrder, order of this note amongst others
// - gvp:historicFlag, vp:estStart, gvp:estEnd, rdfs:comment: Historic
// Information about applicability - dct:source or subproperties thereof
// - dct:contributor or subproperties thereof
const ScopeNoteURI = "http://vocab.getty.edu/ontology#ScopeNote"

// Node in a GVP vocabulary hierarchy.  May be Facet, GuideTerm, Hierarchy,
// Concept, AdminPlaceConcept, PhysPlaceConcept, PhysAdminPlaceConcept, or
// ObsoleteSubject. Implemented as skos:Concept or iso:ThesaurusArray and
// skos:Collection
const SubjectURI = "http://vocab.getty.edu/ontology#Subject"

// Unknown person representing a nationality/culture (the Unknown People by
// Culture facet). Example: 500355202 Unknown Bulgarian (modern) (ULAN)
const UnknownPersonConceptURI = "http://vocab.getty.edu/ontology#UnknownPersonConcept"

// METS Version 1.8 via http://www.loc.gov/standards/mets/version18/mets.xsd
type Term struct {
	//XMLName xml.Name `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# rdf"`

	// TODO: can there be multiple subjects
	Subject *GVPSubject `xml:"http://vocab.getty.edu/ontology# Subject"`

	Statements []*rdf.Statement `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# Statement"`
}

type GVPSubject struct {
	About xsdt.String `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# about,attr"`

	// The types of the subject. Could be concept, etc.
	Types []*rdf.Type `xml:"type"`

	// The types of the subject. Could be concept, etc.
	Labels []*rdfs.Label `xml:"label"`

	BroaderTerms []*rdf.ResourceAttr `xml:"http://vocab.getty.edu/ontology# broader"`

	BroaderPreferredTerms []*rdf.ResourceAttr `xml:"http://vocab.getty.edu/ontology# broaderPreferred"`

	MemberTerms []*rdf.ResourceAttr `xml:"http://www.w3.org/2004/02/skos/core# member"`

	NarrowerTerms []*rdf.ResourceAttr `xml:"http://www.w3.org/2004/02/skos/core# narrower"`
}

func (term *Term) IsConcept() bool {
	return term.IsType(ConceptURI)
}

func (term *Term) IsGuideTerm() bool {
	return term.IsType(GuideTermURI)
}

func (term *Term) IsType(typeURI xsdt.String) bool {
	if term.Subject != nil {
		for _, subjectType := range term.Subject.Types {
			if subjectType.Resource == typeURI {
				return true
			}
		}
	}

	return false
}

// TODO: SKOS:MEMBER seems to indicate subtopics for a guide term
// TODO: SKOS:NARROWER seems to indicate subtopics for a concept
