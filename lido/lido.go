package lido

import (
	"github.com/juju/xml"
	"github.com/verisart/cidoccrm/crm"
	"github.com/verisart/xsd/gml"
	"github.com/verisart/xsd/xsdt"
	"time"
)

const LocalRecordType = "local"
const URIType = "URI"
const Preferred = "preferred"
const Alternate = "alternate"
const RepositoryTitle = "Repository title"
const AlternateTitle = "Alternate title"

// Has two values "yes" or "no"
type AddedSearchTerm xsdt.String

// Convenience for converting a bool into preferred/alternate string
func ToPref(pref bool) xsdt.String {
	if pref {
		return Preferred
	} else {
		return Alternate
	}
}

// Convenience for converting lang string to xsdt.Language, no checking
func ToLang(lang string) xsdt.Language {
	return xsdt.Language(lang)
}

// Convenience for converting from string to xsdt.String
func ToXsdt(text string) xsdt.String {
	return xsdt.String(text)
}

type Lido struct {
	XMLName xml.Name `xml:"http://www.lido-schema.org lido"`

	UsesLido string `xml:"xmlns:lido,attr,omitempty"`

	// Attr{Name: xml.Name{"xmlns", "ns"}, Value: "http://example.com/ns"}

	// A unique lido record identification preferably composed of an
	// identifier for the contributor and a record identification in the
	// contributor's (local) system.
	LidoRecIDs []*Identifier `xml:"http://www.lido-schema.org lidoRecID"`

	// Definition: A unique, published identification of the described object /
	// work. May link to authority files maintained outside of the contributor's
	// documentation system or may be an identifier for the object published by
	// its repository, e.g. composed of an identifier for the repository and an
	// inventory number of the object.Preferably a dereferenceable URL.
	ObjectPublishedIDs []*Identifier `xml:"http://www.lido-schema.org objectPublishedID"`

	// Definition: Indicates the category of which this item is an instance,
	// preferably referring to CIDOC-CRM concept definitions. CIDOC-CRM concept
	// definitions are  given at http://www.cidoc-crm.org/crm-concepts/Data
	// values in the sub-element term may often be: Man-Made Object (with
	// conceptID "http://www.cidoc-crm.org/crm-concepts/E22"), Man-Made Feature
	// (http://www.cidoc-crm.org/crm-concepts/E25), Collection
	// (http://www.cidoc-crm.org/crm-concepts/E78).
	Category *Concept `xml:"http://www.lido-schema.org category"`

	// Holds the descriptive metadata of an object record. The attribute xml:lang
	// is mandatory and specifies the language of the descriptive metadata.For
	// fully multi-lingual resources, repeat this element once for each language
	// represented.If only a few data fields (e.g. title) are provided in more
	// than one language, the respective text elements may be repeated specifying
	// the lang attribute on the text level.
	DescriptiveMetadatas []*DescriptiveMetadata `xml:"http://www.lido-schema.org descriptiveMetadata"`

	// Holds the administrative metadata for an object / work record. The
	// attribute xml:lang is mandatory and specifies the language of the
	// administrative metadata.For fully multi-lingual resources, repeat this
	// element once for each language represented.If only a few data fields (e.g.
	// title, creditline) are provided in more than one language, the respective
	// text elements may be repeated specifying the lang attribute on the text
	// level.
	AdministrativeMetadatas []*AdministrativeMetadata `xml:"http://www.lido-schema.org administrativeMetadata"`

	// Indicates the format of the data source from which the data were migrated.
	// For each sub-element with data values then the related source data fields
	// can be referenced through the attributes encodinganalog and label.
	RelatedEncoding xsdt.String `xml:"http://www.lido-schema.org relatedencoding,attr,omitempty"`
}

// Append a record ID to a lido document. A record id is a unique record
// identification in the contributor's (local) system.
func (l *Lido) AppendRecID(recSource string, recType string, recID string) {
	l.LidoRecIDs = append(l.LidoRecIDs, &Identifier{
		Value:  ToXsdt(recID),
		Source: ToXsdt(recSource),
		Type:   ToXsdt(recType),
	})
}

// Sets the LIDO category to a category defined in the CIDOC CRM
func (l *Lido) SetCRMCategory(crmClass *crm.ConcreteClass) error {
	concept, err := NewCRMConcept(crmClass)

	if err != nil {
		return err
	}

	l.Category = concept
	return nil
}

// Gets or creates a descriptive metadata element for the lido file. These
// elements are only supposed to be repeated for language variants, so if a
// matching language already exists, that element is returned.
func (l *Lido) CreateDesc(lang string) *DescriptiveMetadata {
	xsdtLang := ToLang(lang)
	// Search for metadta description by name
	for _, metadata := range l.DescriptiveMetadatas {
		if metadata.Lang == xsdtLang {
			return metadata
		}
	}

	// Create it
	metadata := &DescriptiveMetadata{
		Lang: xsdtLang,
	}

	l.DescriptiveMetadatas = append(l.DescriptiveMetadatas, metadata)
	return metadata
}

// Identifier is a unique identifier for the concept. Preferably taken from and
// linking to a published controlled vocabulary. How to record: There is no
// controlled list of identifier types. Suggested values include, but are not
// limited to the following: doi (Digital Objects Identifier)guid (Globally
// unique identifier)hdl (Handle)isbn (International Standard Book Number)ismn
// (International Standard Music Number)isrc (International Standard Recording
// Code)issn (International Standard Serials Number)localpermalinkpurl
// (Persistent Uniform Resource Locator)url (Uniform Resource Locator)urn
// (Uniform Resource Name)
type Identifier struct {
	Value xsdt.String `xml:",chardata"`

	// Definition: Qualifies the value as a preferred or alternative variant.
	// How to record: Data values: preferred, alternate
	Pref xsdt.String `xml:"http://www.lido-schema.org pref,attr,omitempty"`

	// Source of the information given in the holding element.
	Source xsdt.String `xml:"http://www.lido-schema.org source,attr,omitempty"`

	// Definition: Qualifies the type of information given in the holding element.
	// How to record: Will generally have to be populated with a given value list.
	Type xsdt.String `xml:"http://www.lido-schema.org type,attr,omitempty"`

	// How to record: Elements with data values are accompanied by the attributes
	// encodinganalog and label to indicate the format of the data source from
	// which the data were migrated. The attribute encodinganalog refers to the
	// internal field label of the source database. The source format is indicated
	// in the attribute relatedencoding of the lidoWrap
	EncodingAnalog xsdt.String `xml:"http://www.lido-schema.org encodinganalog,attr,omitempty"`

	// How to record: Elements with data values are accompanied by the attributes
	// encodinganalog and label, to indicate the format of the data source from
	// which the data were migrated. The attribute label refers to the external
	// label of a data field at the visible user interface. The source format is
	// indicated in the attribute
	Label xsdt.String `xml:"http://www.lido-schema.org label,attr,omitempty"`
}

// Holds the descriptive metadata of an object record. The attribute xml:lang is
// mandatory and specifies the language of the descriptive metadata.For fully
// multi-lingual resources, repeat this element once for each language
// represented.If only a few data fields (e.g. title) are provided in more than
// one language, the respective text elements may be repeated specifying the
// lang attribute on the text level.
type DescriptiveMetadata struct {
	// Wrapper for data classifying the object / work. Includes all classifying
	// information about an object / work, such as: object / work type, style,
	// genre, form, age, sex, and phase, or by how holding organization structures
	// its collection (e.g. fine art, decorative art, prints and drawings, natural
	// science, numismatics, or local history).
	ObjectClass ObjectClassification `xml:"http://www.lido-schema.org objectClassificationWrap"`

	// A Wrapper for information that identifies the object.
	ObjectID ObjectIdentification `xml:"http://www.lido-schema.org objectIdentificationWrap"`

	// Wrapper for event sets.
	EventWrap *EventWrap `xml:"http://www.lido-schema.org eventWrap"`

	// Wrapper for infomation about related topics and works, collections, etc.
	// Notes: This includes visual contents and all associated entities the object
	// is about.
	ObjectRelationWrap *ObjectRelationWrap `xml:"http://www.lido-schema.org objectRelationWrap"`

	// Required language
	Lang xsdt.Language `xml:"http://www.w3.org/XML/1998/namespace lang,attr"`
}

// Were assuming that the language is already provided by the descriptive metadata
func (dm *DescriptiveMetadata) AppendAATWorkType(conceptType string, aatID string, term string) {
	dm.AppendTermWorkType("AAT", conceptType, aatID, term)
}

// Were assuming that the language is already provided by the descriptive metadata
func (dm *DescriptiveMetadata) AppendTermWorkType(termSource string, conceptType string, termID string, term string) {
	dm.ObjectClass.WorkType.Types = append(dm.ObjectClass.WorkType.Types,
		NewConceptClassification(NewTermConcept(termSource, conceptType, termID, term)))
}

type ObjectIdentification struct {
	TitleWrap TitleWrap `xml:"http://www.lido-schema.org titleWrap"`

	InscriptionsWrap *InscriptionsWrap `xml:"http://www.lido-schema.org inscriptionsWrap"`

	RepositoryWrap *RepositoryWrap `xml:"http://www.lido-schema.org repositoryWrap"`

	DisplayStateEditionWrap *DisplayStateEdition `xml:"http://www.lido-schema.org displayStateEditionWrap"`

	Description *ObjectDescription `xml:"http://www.lido-schema.org objectDescriptionWrap"`

	MeasurementsWrap *MeasurementsWrap `xml:"http://www.lido-schema.org objectMeasurementsWrap"`
}

type DisplayStateEdition struct {
	//  Definition: A description of the state of the object / work. Used
	// primarily for prints and other multiples Formulated according to rules.
	// For State, include state identification and known states, as appropriate.
	// Repeat this element only for language variants.
	DisplayStates []*Text `xml:"http://www.lido-schema.org displayState"`

	// A description of the edition of the object / work. Used primarily for
	// prints and other multiples. Formulated according to rules. For Edition,
	// include impression number, edition size, and edition number, or edition
	// name, as appropriate.Repeat this element only for language variants.
	DisplayEditions []*Text `xml:"http://www.lido-schema.org displayEdition"`

	// The published source of the state or edition information.
	SourceStateEditions []*Text `xml:"http://www.lido-schema.org sourceStateEdition"`
}

type ObjectDescription struct {
	Notes []*DescriptiveNote `xml:"http://www.lido-schema.org objectDescriptionSet"`
}

// Wrapper for infomation about related topics and works, collections, etc.
// Notes: This includes visual contents and all associated entities the object
// is about.
type ObjectRelationWrap struct {
	// A wrapper for Subject information. This may be the visual content (e.g. the
	// iconography of a painting) or what the object is about.
	SubjectWrap *SubjectWrap `xml:"http://www.lido-schema.org subjectWrap"`

	// A wrapper for Related Works information.
	RelatedWorksWrap *RelatedWorksWrap `xml:"http://www.lido-schema.org relatedWorksWrap"`
}

// A wrapper for Related Works information.
type RelatedWorksWrap struct {
	// A wrapper for a object / work, group, collection, or series that is
	// directly related to the object / work being recorded.
	RelatedWorkSets []*RelatedWorkSet `xml:"http://www.lido-schema.org relatedWorkSet"`
}

// A wrapper for a object / work, group, collection, or series that is directly
// related to the object / work being recorded.
type RelatedWorkSet struct {
	// A term describing the nature of the relationship between the object / work
	// at hand and the related entity. Example values: part of, larger context
	// for, model of, model for, study of, study forrendering of, copy of, related
	// to. Indicate a term characterizing the relationship from the perspective of
	// the currently described object / work towards the related object / work.
	// Notes: For implementation of the data, note that relationships are
	// conceptually reciprocal, but the Relationship Type is often different on
	// either side of the relationship (e.g., one work is part of a second work,
	// but from the point of view of the second record, the first work is the
	// larger context for the second work). Whether or not relationships are
	// physically reciprocal as implemented in systems is a local decision.
	RelatedWorkRelType *Concept `xml:"http://www.lido-schema.org relatedWorkRelType"`

	// Wrapper for the display and reference elements of a related object / work.
	RelatedWork *ObjectSet `xml:"http://www.lido-schema.org relatedWork"`

	// Assigns a priority order for online presentation of the element. Has to be
	// a positive integer, with descending priority from 1 to x.
	SortOrder xsdt.Integer `xml:"http://www.lido-schema.org sortorder,attr,omitempty"`
}

// A wrapper for Subject information. This may be the visual content (e.g. the
// iconography of a painting) or what the object is about.
type SubjectWrap struct {
	// Wrapper for display and index elements for one set of subject information.
	// If an object / work has multiple parts or otherwise has separate, multiple
	// subjects, repeat this element and use Extent Subject in the Subject
	// element. This element may also be repeated to distinguish between subjects
	// that reflect what an object / work is *of* (description and identification)
	// from what it is *about* (interpretation).
	SubjectSets []*SubjectSet `xml:"http://www.lido-schema.org subjectSet"`
}

// Wrapper for display and index elements for one set of subject information.
// If an object / work has multiple parts or otherwise has separate, multiple
// subjects, repeat this element and use Extent Subject in the Subject element.
// This element may also be repeated to distinguish between subjects that
// reflect what an object / work is *of* (description and identification) from
// what it is *about* (interpretation).
type SubjectSet struct {
	// A free-text description of the subject matter represented by/in the object
	// / work, corresponding to the following subject element Repeat this element
	// only for language variants.
	DisplaySubjects []*Text `xml:"http://www.lido-schema.org displaySubject"`

	// Contains sub-elements for a structured subject description. These identify,
	// describe, and/or interpret what is depicted in and by an object / work or
	// what it is about.
	Subject *Subject `xml:"http://www.lido-schema.org subject"`

	// Assigns a priority order for online presentation of the element. Has to be
	// a positive integer, with descending priority from 1 to x.
	SortOrder xsdt.Integer `xml:"http://www.lido-schema.org sortorder,attr,omitempty"`
}

type Subject struct {
	//	Definition: A place depicted in or by an object / work, or what it is about, provided as display and index elements.
	SubjectPlaces []*PlaceSet `xml:"http://www.lido-schema.org subjectPlace"`

	//	Definition: An object - e.g. a building or a work of art depicted in or by an object / work, or what it is about, provided as display and index elements.
	SubjectObjects []*ThingPresent `xml:"http://www.lido-schema.org subjectObject"`

	Type xsdt.String `xml:"http://www.lido-schema.org type,attr,omitempty"`

	//	Definition: When there are multiple subjects, a term indicating the part of the object / work to which these subject terms apply.
	//	How to record: Example values: recto, verso, side A, side B, main panel, and predella.Repeat this element only for language variants.
	ExtentSubjects []*Text `xml:"http://www.lido-schema.org extentSubject"`

	//	Definition: Provides references to concepts related to the subject of the described object / work.
	//	How to record: May include iconography, themes from literature, or generic terms describing the material world, or topics (e.g., concepts, themes, or issues). However, references to people, dates, events, places, objects are indicated in the the respective sub-elements Subject Actor Set, Subject Date Set, Subject Event Set, Subject Place Set, and Subject Object Set.Preferably taken from a published controlled vocabulary.
	SubjectConcepts []*ConceptElement `xml:"http://www.lido-schema.org subjectConcept"`

	//	Definition: A person, group, or institution depicted in or by an object / work, or what it is about, provided as display and index elements.
	SubjectActors []*SubjectActor `xml:"http://www.lido-schema.org subjectActor"`

	//	Definition: A time specification depicted in or by an object / work, or what it is about, provided as display and index elements.
	SubjectDates []*DateSpan `xml:"http://www.lido-schema.org subjectDate"`

	//	Definition: An event depicted in or by an object / work, or what it is about, provided as display and index elements.
	SubjectEvents []*EventElement `xml:"http://www.lido-schema.org subjectEvent"`
}

type SubjectActor struct {
	// Display element for one actor, corresponding to the following actor
	// element. May include name, brief biographical information of the named
	// actor, presented in a syntax suitable for display to the end-user. If there
	// is no known actor, make a reference to the presumed culture or nationality
	// of the unknown actor. May be concatenated from the respective Actor element.
	// The name should be in natural order, if possible, although inverted order
	// is acceptable. Include nationality and life dates. For unknown actors, use
	// e.g.: "unknown," "unknown Chinese," "Chinese," or "unknown 15th century
	// Chinese." Repeat this element only for language variants.
	DisplayActors []*Text `xml:"http://www.lido-schema.org displayActor"`

	// Describes and identifies an actor, i.e. a person, corporation, family or
	// group, containing structured sub-elements for indexing and identification
	// references.
	Actor *Actor `xml:"http://www.lido-schema.org actor"`

	// Assigns a priority order for online presentation of the element. Has to be
	// a positive integer, with descending priority from 1 to x.
	SortOrder xsdt.Integer `xml:"http://www.lido-schema.org sortorder,attr,omitempty"`
}

type ObjectClassification struct {
	// A wrapper for Object/Work Types.
	WorkType ObjectWorkTypeWrap `xml:"http://www.lido-schema.org objectWorkTypeWrap"`

	// A wrapper for any classification used to categorize an object / work by
	// grouping it together with others on the basis of similar characteristics.
	ClassificationWrap *ClassificationWrap `xml:"http://www.lido-schema.org classificationWrap"`
}

type ObjectWorkTypeWrap struct {
	Types []*ClassificationElement `xml:"http://www.lido-schema.org objectWorkType"`
}

type ClassificationWrap struct {
	Classifications []*ClassificationElement `xml:"http://www.lido-schema.org classification"`
}

type LegalBodyRef struct {
	//	Definition: Unambiguous identification of the institution or person referred to as legal body.
	LegalBodyIDs []*Identifier `xml:"http://www.lido-schema.org legalBodyID"`

	//	Definition: Appellation of the institution or person.
	LegalBodyNames []*Appellation `xml:"http://www.lido-schema.org legalBodyName"`

	//	Definition: Weblink of the institution or person referred to as legal body.
	LegalBodyWeblinks []*WebResource `xml:"http://www.lido-schema.org legalBodyWeblink"`

	// Assigns a priority order for online presentation of the element. Has to be
	// a positive integer, with descending priority from 1 to x.
	SortOrder xsdt.Integer `xml:"http://www.lido-schema.org sortorder,attr,omitempty"`

	Type xsdt.String `xml:"http://www.lido-schema.org type,attr,omitempty"`
}

type AdministrativeMetadata struct {
	// Wrapper for rights information about the object / work described.
	// Notes: Rights information for the record and for resources is recorded in
	// the respective rights elements recordRights and rightsResource.
	RightsWorkWrap *RightsWorkWrap `xml:"http://www.lido-schema.org rightsWorkWrap"`

	// A wrapper for information about the record that contains the cataloguing
	// information. Note that this section does not refer to any object or
	// resource information, but only to the source record.
	RecordWrap *RecordWrap `xml:"http://www.lido-schema.org recordWrap"`

	// A wrapper for resources that are surrogates for an object / work, including
	// digital images, videos or audio files that represent it in an online
	// service.
	ResourceWrap *ResourceWrap `xml:"http://www.lido-schema.org resourceWrap"`

	Lang xsdt.Language `xml:"http://www.w3.org/XML/1998/namespace lang,attr"`
}

type RightsWorkWrap struct {
	RightsWorkSets []*Rights `xml:"http://www.lido-schema.org rightsWorkSet"`
}

type RecordWrap struct {
	// A unique record identification in the contributor's (local) system.
	RecordIDs []*Identifier `xml:"http://www.lido-schema.org recordID"`

	// Term establishing whether the record represents an individual item or a
	// collection, series, or group of works. Mandatory. Example values: item,
	// collection, series, group, volume, fonds. Preferably taken from a published
	// controlled value list.
	RecordType *Concept `xml:"http://www.lido-schema.org recordType"`

	// The source of information in this record, generally the repository or other
	// institution.
	RecordSources []*LegalBodyRef `xml:"http://www.lido-schema.org recordSource"`

	// Information about rights regarding the content provided in this LIDO
	// record.
	RecordRights []*Rights `xml:"http://www.lido-schema.org recordRights"`

	// Wrapper for metadata information about this record.
	RecordInfoSets []*RecordInfo `xml:"http://www.lido-schema.org recordInfoSet"`
}

type RecordInfo struct {

	// Unique ID of the metadata. Record Info ID has the same definition as Record
	// ID but out of the context of original local system, such as a persistent
	// identifier or an oai identifier (e.g., oai1:getty.edu:paintings/00001234
	// attribute type= oai).
	RecordInfoIDs []*Identifier `xml:"http://www.lido-schema.org recordInfoID"`

	// Link of the metadata, e.g., to the object data sheet (not the same as link
	// of the object).
	RecordInfoLinks []*WebResource `xml:"http://www.lido-schema.org recordInfoLink"`

	// Creation date or date modified of the metadata record. Format will vary
	// depending upon implementation.
	RecordMetadataDates []*Note `xml:"http://www.lido-schema.org recordMetadataDate"`

	Type xsdt.String `xml:"http://www.lido-schema.org type,attr,omitempty"`
}

type ResourceWrap struct {
	//	Definition: Contains sub-elements for a structured resource description.
	//	Notes: Provides identification of a surrogate of the object / work including digital images, slides, transparencies, photographs, audio, video and moving images, but excluding items that are considered object / works in their own right. For such as drawings, prints, paintings, or photographs considered art, and other works that themselves contain representations of other works, use Related Works and/or Subjects.
	ResourceSets []*ResourceSet `xml:"http://www.lido-schema.org resourceSet"`
}

type ResourceSet struct {
	// Information about rights regarding the image or other resource. Use this
	// sub-element if the holder of the reproduction rights for the image/resource
	// differs from the holder of rights for the work. See also Rights Work above.
	// (E.g., the work rights are " National Museum of African Art, Smithsonian
	// Instituition (Washing DC), " but the image rights are "Photo Frank Khoury.")
	RightsResources []*Rights `xml:"http://www.lido-schema.org rightsResource"`

	// A digital representation of a resource for online presentation. Repeat this
	// element set for variants representing the same resource, e.g. different
	// sizes of the same image, or a thumbnail representing an audio or video file
	// and the digital audio or video file itself.
	ResourceRepresentations []*ResourceRep `xml:"http://www.lido-schema.org resourceRepresentation"`

	// The generic identification of the medium of the image or other resource.
	// Preferably using a controlled published value list. Example values: digital
	// image, photograph, slide, videotape, X-ray photograph, negative.
	ResourceType *Concept `xml:"http://www.lido-schema.org resourceType"`

	// A date or range of dates associated with the creation or production of the
	// original resource, e.g. the image or recording.
	// Notes: This is not necessarily the same as the date of production of the
	// digital resource (e.g. a digitization of a negative is usually made years
	// after the image was captured on film). Format will vary depending upon
	// implementation.
	ResourceDateTaken *DateSet `xml:"http://www.lido-schema.org resourceDateTaken"`

	// Identification of the agency, individual, or repository from which the
	// image or other resource was obtained. Include this sub-element when the
	// source of the image/resource differs from the source named in Record Source.
	ResourceSources []*LegalBodyRef `xml:"http://www.lido-schema.org resourceSource"`

	// The unique numeric or alphanumeric identification of the original (digital
	// or analogue) resource.
	ResourceID *Identifier `xml:"http://www.lido-schema.org resourceID"`

	// The relationship of the resource to the object / work being described.
	// Example values: conservation image, documentary image, contextual image,
	// historical image, reconstruction, and installation image
	ResourceRelTypes []*Concept `xml:"http://www.lido-schema.org resourceRelType"`

	// The specific vantage point or perspective of the view.
	ResourcePerspectives []*Concept `xml:"http://www.lido-schema.org resourcePerspective"`

	// A description of the spatial, chronological, or contextual aspects of the
	// object / work as captured in this particular resource.
	ResourceDescriptions []*Note `xml:"http://www.lido-schema.org resourceDescription"`

	// Assigns a priority order for online presentation of the element. Has to be
	// a positive integer, with descending priority from 1 to x.
	SortOrder xsdt.Integer `xml:"http://www.lido-schema.org sortorder,attr,omitempty"`
}

type Rights struct {
	//	Definition: The specific type of right being recorded.
	//	How to record: For example: copyright, publication right, data protection right, trademark.Preferably taken from a published controlled value list.
	RightsTypes []*Concept `xml:"http://www.lido-schema.org rightsType"`

	//	Definition: The date on which a right is or was current.
	RightsDate *DateSpan `xml:"http://www.lido-schema.org rightsDate"`

	//	Definition: The holder of the right.
	RightsHolders []*LegalBodyRef `xml:"http://www.lido-schema.org rightsHolder"`

	//	Definition: Acknowledgement of the rights associated with the physical and/or digital object as requested.
	//	How to record: Repeat this element only for language variants.
	CreditLines []*Text `xml:"http://www.lido-schema.org creditLine"`

	// Assigns a priority order for online presentation of the element. Has to be
	// a positive integer, with descending priority from 1 to x.
	SortOrder xsdt.Integer `xml:"http://www.lido-schema.org sortorder,attr,omitempty"`
}

type ResourceRep struct {
	//	Definition: A url reference in the worldwide web environment.
	LinkResource *LinkResource `xml:"http://www.lido-schema.org linkResource"`

	//	Definition: Any technical measurement information needed for online presentation of the resource.
	//	How to record: For images provide width and height of the digital image, for audio or video resources provide duration, bit rate, frame size, and if necessary TC-IN, TC-OUT.
	ResourceMeasurementsSets []*MeasurementsSet `xml:"http://www.lido-schema.org resourceMeasurementsSet"`

	Type xsdt.String `xml:"http://www.lido-schema.org type,attr,omitempty"`
}

type LinkResource struct {
	WebResource

	//	Definition: Codec information about the digital resource.
	CodecResource xsdt.String `xml:"http://www.lido-schema.org codecResource,attr"`
}

type EventWrap struct {
	Events []*EventElement `xml:"http://www.lido-schema.org eventSet"`
}

func (ew *EventWrap) AppendEvent(event *Event) {
	element := &EventElement{
		Event: event,
	}

	ew.Events = append(ew.Events, element)
}

type WorkID struct {
	XsdtString xsdt.String `xml:",chardata"`
	// How to record: Elements with data values are accompanied by the attributes
	// encodinganalog and label to indicate the format of the data source from
	// which the data were migrated. The attribute encodinganalog refers to the
	// internal field label of the source database. The source format is indicated
	// in the attribute relatedencoding of the lidoWrap
	EncodingAnalog xsdt.String `xml:"http://www.lido-schema.org encodinganalog,attr,omitempty"`

	// How to record: Elements with data values are accompanied by the attributes
	// encodinganalog and label, to indicate the format of the data source from
	// which the data were migrated. The attribute label refers to the external
	// label of a data field at the visible user interface. The source format is
	// indicated in the attribute
	Label xsdt.String `xml:"http://www.lido-schema.org label,attr,omitempty"`

	// Assigns a priority order for online presentation of the element. Has to be
	// a positive integer, with descending priority from 1 to x.
	SortOrder xsdt.Integer `xml:"http://www.lido-schema.org sortorder,attr,omitempty"`

	// Specification of the date, e.g. if it is an exact or an estimated earliest
	// date. Data values may be: exactDate, estimatedDate.
	Type xsdt.String `xml:"http://www.lido-schema.org type,attr,omitempty"`
}

type EventElement struct {
	// Display element for an event, corresponding to the following event element.
	// How to record: Repeat this element only for language variants.
	DisplayEvents []*Text `xml:"http://www.lido-schema.org displayEvent"`

	// Identifying, descriptive and indexing information for the events in which
	// the object participated or was present at, e.g. creation, excavation,
	// collection, and use. All information related to the creation of an object:
	// creator, cutlural context, creation date, creation place, the material and
	// techniques used are recorded here, qualified by the event type “creation”.
	Event *Event `xml:"http://www.lido-schema.org event"`

	// Assigns a priority order for online presentation of the element. Has to be
	// a positive integer, with descending priority from 1 to x.
	SortOrder xsdt.Integer `xml:"http://www.lido-schema.org sortorder,attr,omitempty"`
}

// Simple text element with encodinganalog and label attribute
type Text struct {
	Value xsdt.String `xml:",chardata"`

	Lang xsdt.Language `xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty"`

	// How to record: Elements with data values are accompanied by the attributes
	// encodinganalog and label to indicate the format of the data source from
	// which the data were migrated. The attribute encodinganalog refers to the
	// internal field label of the source database. The source format is indicated
	// in the attribute relatedencoding of the lidoWrap
	EncodingAnalog xsdt.String `xml:"http://www.lido-schema.org encodinganalog,attr,omitempty"`

	// How to record: Elements with data values are accompanied by the attributes
	// encodinganalog and label, to indicate the format of the data source from
	// which the data were migrated. The attribute label refers to the external
	// label of a data field at the visible user interface. The source format is
	// indicated in the attribute
	Label xsdt.String `xml:"http://www.lido-schema.org label,attr,omitempty"`
}

type Event struct {
	//	Definition: A unique identifier for the event.
	//	How to record: Preferably taken from and linking to a published resource describing the event.
	EventIDs []*Identifier `xml:"http://www.lido-schema.org eventID"`

	//	Definition: The nature of the event associated with an object / work.
	//	How to record: Controlled. Recommended: Defined list of subclasses of CRM entity E5 Event.Basic event types as recorded in sub-element term include: Acquisition, Collecting, Commisioning, Creation, Designing, Destruction, Event (non-specified), Excavation, Exhibition, Finding, Loss, Modification, Move, Part addition, Part removal, Performance, Planning, Production, Provenance, Publication, Restoration, Transformation, Type assignment, Type creation, Use.
	EventTypes []*Concept `xml:"http://www.lido-schema.org eventType"`

	//	Definition: Date specification of the event.
	Date *DateSet `xml:"http://www.lido-schema.org eventDate"`

	//	Definition: Place specification of the event.
	EventPlaces []*EventPlace `xml:"http://www.lido-schema.org eventPlace"`

	// The method by which the event is carried out. Preferably taken from a
	// published controlled vocabulary.
	// Notes: Used e.g. for SPECTRUM Units of Information
	// "field collection method", "acquisition method".
	EventMethods []*ConceptElement `xml:"http://www.lido-schema.org eventMethod"`

	// References another object that was present at this same event.
	ThingPresents []*ThingPresent `xml:"http://www.lido-schema.org thingPresent"`

	// Wrapper for a description of the event, including description identifer,
	// descriptive note of the event and its sources. If there is more than one
	// descriptive note, repeat this element.
	EventDescriptionSets []*DescriptiveNote `xml:"http://www.lido-schema.org eventDescriptionSet"`

	// An appellation for the event, e.g. a title, identifying phrase, or name
	// given to it.
	EventNames []*Appellation `xml:"http://www.lido-schema.org eventName"`

	// Wrapper for display and index elements for an actor with role information
	// (participating or being present in the event). For multiple actors repeat
	// the element.
	EventActors []*EventActor `xml:"http://www.lido-schema.org eventActor"`

	//	Definition: References an event which is linked in some way to this event, e.g. a field trip within which this object was collected.
	RelatedEvents []*RelatedEvent `xml:"http://www.lido-schema.org relatedEventSet"`

	// The role played within this event by the described entity. Preferably taken
	// from a published controlled vocabulary.
	RoleInEvents []*Concept `xml:"http://www.lido-schema.org roleInEvent"`

	// Name of a culture, cultural context, people, or also a nationality.
	// Preferably using a controlled vocabuarly.
	Cultures []*ConceptElement `xml:"http://www.lido-schema.org culture"`

	// A period in which the event happened. Preferably taken from a published
	// controlled vocabulary. Repeat this element only for indicating an earliest
	// and latest period delimiting the event.
	// Notes: Period concepts have delimiting character in time and space.
	PeriodNames []*ClassificationElement `xml:"http://www.lido-schema.org periodName"`

	// Indicates the substances or materials used within the event (e.g. the
	// creation of an object / work), as well as any implements, production or
	// manufacturing techniques, processes, or methods incorporated. Will be used
	// most often within a production event, but also others such as excavation,
	// restoration, etc.
	EventMaterialsTechs []*EventMaterialsTech `xml:"http://www.lido-schema.org eventMaterialsTech"`

	// NOTE, below here is a modification for Verisart's internal uses, please
	// ignore and do not use should be no side effects
	MeasurementsWrap *MeasurementsWrap `xml:"http://www.lido-schema.org objectMeasurementsWrap"`
}

// Sets the LIDO category to a category defined in the CIDOC CRM
func (e *Event) AppendCRMType(crmClass *crm.ConcreteClass) error {
	concept, err := NewCRMConcept(crmClass)

	if err != nil {
		return err
	}

	e.EventTypes = append(e.EventTypes, concept)
	return nil
}

func (e *Event) SetDate(min time.Time, max time.Time) {
	e.Date = &DateSet{
		Date: &DateSpan{
			EarliestDate: &Date{
				Value: ToXsdt(min.In(time.UTC).Format("2006-01-02T15:04:05Z0700")),
			},
			LatestDate: &Date{
				Value: ToXsdt(max.In(time.UTC).Format("2006-01-02T15:04:05Z0700")),
			},
		},
	}
}

type DateSet struct {
	// Display element for a date specification, corresponding to the following
	// date element. It is a concise description of the date, presented in a
	// syntax suitable for display to the end-user and including any necessary
	// indications of uncertainty, ambiguity, and nuance.Repeat this element only
	// for language variants.
	DisplayDates []*Text `xml:"http://www.lido-schema.org displayDate"`

	// Contains a date specification by providing a set of years as earliest and
	// latest date delimiting the respective span of time.This may be a period or
	// a set of years in the proleptic Gregorian calendar delimiting the span of
	// time. If it is an exact date, possibly with time, repeat the same date (and
	// time) in earliest and latest date.
	Date *DateSpan `xml:"http://www.lido-schema.org date"`
}

type DateSpan struct {
	//	Definition: A year or exact date that broadly delimits the beginning of an implied date span.
	//	How to record: General format: YYYY[-MM[-DD]]Format is according to ISO 8601. This may include date and time specification.
	EarliestDate *Date `xml:"http://www.lido-schema.org earliestDate"`

	//	Definition: A year or exact date that broadly delimits the end of an implied date span.
	//	How to record: General format: YYYY[-MM[-DD]]Format is according to ISO 8601. This may include date and time specification.
	LatestDate *Date `xml:"http://www.lido-schema.org latestDate"`
}

// A year or exact date that broadly delimits the beginning of an implied date
// span. General format: YYYY[-MM[-DD]]Format is according to ISO 8601. This may
// include date and time specification.
type Date struct {
	Value xsdt.String `xml:",chardata"`

	// Source of the information given in the holding element.
	Source xsdt.String `xml:"http://www.lido-schema.org source,attr,omitempty"`

	// Specification of the date, e.g. if it is an exact or an estimated earliest
	// date. Data values may be: exactDate, estimatedDate.
	Type xsdt.String `xml:"http://www.lido-schema.org type,attr,omitempty"`

	// How to record: Elements with data values are accompanied by the attributes
	// encodinganalog and label to indicate the format of the data source from
	// which the data were migrated. The attribute encodinganalog refers to the
	// internal field label of the source database. The source format is indicated
	// in the attribute relatedencoding of the lidoWrap
	EncodingAnalog xsdt.String `xml:"http://www.lido-schema.org encodinganalog,attr,omitempty"`

	// How to record: Elements with data values are accompanied by the attributes
	// encodinganalog and label, to indicate the format of the data source from
	// which the data were migrated. The attribute label refers to the external
	// label of a data field at the visible user interface. The source format is
	// indicated in the attribute
	Label xsdt.String `xml:"http://www.lido-schema.org label,attr,omitempty"`
}

type EventPlace struct {
	PlaceSet

	//	How to record: Data values may be: moveFrom, moveTo, alternative.
	Type xsdt.String `xml:"http://www.lido-schema.org type,attr,omitempty"`
}

type PlaceSet struct {
	//	Definition: Display element for a place specification, corresponding to the following place element.
	//	How to record: Repeat this element only for language variants.
	DisplayPlaces []*Text `xml:"http://www.lido-schema.org displayPlace"`

	// Contains structured identifying and indexing information for a geographical
	// entity.
	Place *Place `xml:"http://www.lido-schema.org place"`

	// Assigns a priority order for online presentation of the element. Has to be
	// a positive integer, with descending priority from 1 to x.
	SortOrder xsdt.Integer `xml:"http://www.lido-schema.org sortorder,attr,omitempty"`
}

type Place struct {
	//	Definition: Allows for indexing larger geographical entities.
	PartOfPlaces []*Place `xml:"http://www.lido-schema.org partOfPlace"`

	// A classification of the place, e.g. by geological complex, stratigraphic
	// unit or habitat type.
	PlaceClassifications []*PlaceClassification `xml:"http://www.lido-schema.org placeClassification"`

	//Data values can include: Gemeinde, Kreis, Bundesland, Staat, Herzogtum,
	// city, county, country, civil parish
	PoliticalEntity xsdt.String `xml:"http://www.lido-schema.org politicalEntity,attr"`

	//	Definition: Data values can include: Naturraum, Landschaft, natural environment, landscape
	GeographicalEntity xsdt.String `xml:"http://www.lido-schema.org geographicalEntity,attr"`

	//	Definition: A unique identifier for the place.
	//	How to record: Preferably taken from a published authority file.
	PlaceIDs []*Identifier `xml:"http://www.lido-schema.org placeID"`

	// The name of the geographic place. If there are different names of the same
	// place, e.g. today's and historical names, repeat this element.
	NamePlaceSets []*Appellation `xml:"http://www.lido-schema.org namePlaceSet"`

	// Georeferences of the place using the GML specification. Repeat this element
	// only for language variants.
	// Notes: For further documentation on GML refer to
	// http://www.opengis.net/gml/.
	GMLs []*GML `xml:"gml"`
}

// A classification of the place, e.g. by geological complex, stratigraphic unit
// or habitat type.
type PlaceClassification struct {
	Concept

	Type xsdt.String `xml:"http://www.lido-schema.org type,attr,omitempty"`
}

// Specifies the GML instantiation for georeferences. Notes: For documentation
// on GML refer to http://www.opengis.net/gml/.
type GML struct {
	Lang xsdt.Language `xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty"`

	LineStrings []*gml.LineString `xml:"http://www.opengis.net/gml LineString"`

	Polygons []*gml.Polygon `xml:"http://www.opengis.net/gml Polygon"`

	Points []*gml.Point `xml:"http://www.opengis.net/gml Point"`
}

type ConceptElement struct {
	Concept

	// Assigns a priority order for online presentation of the element. Has to be
	// a positive integer, with descending priority from 1 to x.
	SortOrder xsdt.Integer `xml:"http://www.lido-schema.org sortorder,attr,omitempty"`
}

type ClassificationElement struct {
	Concept

	Type xsdt.String `xml:"http://www.lido-schema.org type,attr,omitempty"`

	// Assigns a priority order for online presentation of the element. Has to be
	// a positive integer, with descending priority from 1 to x.
	SortOrder xsdt.Integer `xml:"http://www.lido-schema.org sortorder,attr,omitempty"`
}

func NewConceptClassification(concept *Concept) *ClassificationElement {
	return &ClassificationElement{
		Concept: *concept,
	}
}

type ThingPresent struct {
	ObjectSet

	// Assigns a priority order for online presentation of the element. Has to be
	// a positive integer, with descending priority from 1 to x.
	SortOrder xsdt.Integer `xml:"http://www.lido-schema.org sortorder,attr,omitempty"`
}

type ObjectSet struct {
	//	Definition: A free-text description of the object, corresponding to the following object element
	//	How to record: Repeat this element only for language variants.
	DisplayObjects []*Text `xml:"http://www.lido-schema.org displayObject"`

	//	Definition: Contains identifying information and links to another object.
	Object *Object `xml:"http://www.lido-schema.org object"`
}

type Object struct {
	// A URL-Reference to a description of the object / work in the worldwide web
	// environment.
	ObjectWebResources []*WebResource `xml:"http://www.lido-schema.org objectWebResource"`

	//	Definition: Unique identifier of the referenced object / work.
	ObjectIDs []*Identifier `xml:"http://www.lido-schema.org objectID"`

	//	Definition: A descriptive identification of the object / work that will be meaningful to end-users, including some or all of the following information, as necessary for clarity and if known: title, object/work type, important actor, date and/or place information, potentially location of the object / work.
	//	How to record: The information should ideally be generated from fields/elements in the related record.
	ObjectNotes []*Note `xml:"http://www.lido-schema.org objectNote"`
}

type WebResource struct {
	XsdtString xsdt.String `xml:"http://www.lido-schema.org ,chardata"`

	//	Definition: Indicates the internet media type, e.g. the file format of the given web resource.
	//	How to record: Data values should be taken from the official IANA list (see http://www.iana.org/assignments/media-types/). Includes: text/html, text/xml, image/jpeg, audio/mpeg, video/mpeg, application/pdf.
	FormatResource xsdt.String `xml:"http://www.lido-schema.org formatResource,attr"`

	// How to record: Elements with data values are accompanied by the attributes
	// encodinganalog and label to indicate the format of the data source from
	// which the data were migrated. The attribute encodinganalog refers to the
	// internal field label of the source database. The source format is indicated
	// in the attribute relatedencoding of the lidoWrap
	EncodingAnalog xsdt.String `xml:"http://www.lido-schema.org encodinganalog,attr,omitempty"`

	// How to record: Elements with data values are accompanied by the attributes
	// encodinganalog and label, to indicate the format of the data source from
	// which the data were migrated. The attribute label refers to the external
	// label of a data field at the visible user interface. The source format is
	// indicated in the attribute
	Label xsdt.String `xml:"http://www.lido-schema.org label,attr,omitempty"`

	Lang xsdt.Language `xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty"`

	// Qualifies the value as a preferred or alternative variant. Data values:
	// preferred, alternate
	Pref xsdt.String `xml:"http://www.lido-schema.org pref,attr,omitempty"`
}

// A descriptive identification of the object / work that will be meaningful to
// end-users, including some or all of the following information, as necessary
// for clarity and if known: title, object/work type, important actor, date
// and/or place information, potentially location of the object / work.
// The information should ideally be generated from fields/elements in the
// related record.
type Note struct {
	Text

	// Assigns a priority order for online presentation of the element. Has to be
	// a positive integer, with descending priority from 1 to x.
	SortOrder xsdt.Integer `xml:"http://www.lido-schema.org sortorder,attr,omitempty"`

	// Source of the information given in the holding element.
	Source xsdt.String `xml:"http://www.lido-schema.org source,attr,omitempty"`

	Type xsdt.String `xml:"http://www.lido-schema.org type,attr,omitempty"`
}

type DescriptiveNote struct {
	Type xsdt.String `xml:"http://www.lido-schema.org type,attr,omitempty"`

	// Assigns a priority order for online presentation of the element. Has to be
	// a positive integer, with descending priority from 1 to x.
	SortOrder xsdt.Integer `xml:"http://www.lido-schema.org sortorder,attr,omitempty"`

	//	Definition: Identifier for an external resource describing the entity.
	//	Notes: The referenced resource may be any kind of document, preferably web-accessible.
	IDs []*Identifier `xml:"http://www.lido-schema.org descriptiveNoteID"`

	//	Definition: Usually a relatively brief essay-like text that describes the entity.
	//	How to record: Repeat this element only for language variants.
	Values []*Text `xml:"http://www.lido-schema.org descriptiveNoteValue"`

	//	DeTefinition: The source for the descriptive note, generally a published source.
	Sources []*Text `xml:"http://www.lido-schema.org sourceDescriptiveNote"`
}

type EventActor struct {
	ActorInRole

	// Assigns a priority order for online presentation of the element. Has to be
	// a positive integer, with descending priority from 1 to x.
	SortOrder xsdt.Integer `xml:"http://www.lido-schema.org sortorder,attr,omitempty"`
}

type ActorInRoleSet struct {
	// Display element for an actor coupled with its specific role, corresponding
	// to the following actor element. May include name, brief biographical
	// information, and roles (if necessary) of the named actor, presented in a
	// syntax suitable for display to the end-user and including any necessary
	// indications of uncertainty, ambiguity, and nuance. If there is no known
	// actor, make a reference to the presumed culture or nationality of the
	// unknown actor.May be concatenated from the respective Actor element. The
	// name should be in natural order, if possible, although inverted order is
	// acceptable. Include nationality and life dates. For unknown actors, use
	// e.g.: "unknown," "unknown Chinese," "Chinese," or "unknown 15th century
	// Chinese."Repeat this element only for language variants.
	DisplayActorInRoles []*Text `xml:"http://www.lido-schema.org displayActorInRole"`

	//	Definition: Describes an actor with role and (if necessary) attributions in a structured way, consisting of the sub-elements actor, its role, attribution and extent.
	ActorInRole *ActorInRole `xml:"http://www.lido-schema.org actorInRole"`
}

type ActorInRole struct {
	// Contains structured identifying and indexing actor information.
	Actor *Actor `xml:"http://www.lido-schema.org actor"`

	// Role of the Actor in the event. Preferably taken from a published
	// controlled vocabulary.
	RoleActors []*ConceptElement `xml:"http://www.lido-schema.org roleActor"`

	// A qualifier used when the attribution is uncertain, is in dispute, when
	// there is more than one actor, when there is a former attribution, or when
	// the attribution otherwise requires explanation. Example values: attributed
	// to, studio of, workshop of, atelier of, office of, assistant of, associate
	// of, pupil of, follower of, school of, circle of, style of, after copyist
	// of, manner of...
	AttributionQualifierActors []*Text `xml:"http://www.lido-schema.org attributionQualifierActor"`

	// Extent of the actor's participation in the event, if there are several
	// actors. Example values: design, execution, with additions by, figures,
	// renovation by, predella, embroidery, cast by, printed by, ...
	ExtentActors []*Text `xml:"http://www.lido-schema.org extentActor"`
}

// In some cases the actor will be encrypted such as events.
type Actor struct {
	XMLName xml.Name `xml:"http://www.lido-schema.org actor"`
	// A unique identifier for the actor. Preferably taken from a published
	// authority file.
	ActorIDs []*Identifier `xml:"http://www.lido-schema.org actorID"`

	// A wrapper for name elements. if there exists more than one name for a
	// single actor, repeat Name Actor Set. Indicates names, appellations, or
	// other identifiers assigned to an individual, group of people, firm or other
	// corporate body, or other entity.
	NameActorSets []*Appellation `xml:"http://www.lido-schema.org nameActorSet"`

	// National or cultural affiliation of the person or corporate body.
	// Preferably taken from a published controlled vocabulary.
	NationalityActors []*ConceptElement `xml:"http://www.lido-schema.org nationalityActor"`

	// The lifespan of the person or the existence of the corporate body or group.
	// For individuals, record birth date as earliest and death date as latest
	// date, estimated where necessary. For a corporate body or group, record the
	// dates of founding and dissolution.Although this is not a mandatory field
	// the use of birth date and death date is strongly recommended for unambigous
	// identification of individuals. The type attribute of earliest and latest
	// date may specify for indiviudals, if birth and death dates or if dates of
	// activity are recorded. Data values for type attribute may include:
	// birthDate, deathDate, estimatedDate.
	VitalDatesActor *DateSpan `xml:"http://www.lido-schema.org vitalDatesActor"`

	// The sex of the individual. Data values: male, female, unknown, not
	// applicable.Repeat this element for language variants only.
	// Notes: Not applicable for corporate bodies.
	GenderActors []*Text `xml:"http://www.lido-schema.org genderActor"`

	// Indicates if the actor is an individual, a group of individuals, a family
	// or a corporation (firm or other corporate body). Data values: person,
	// group, family, corporation.
	Type xsdt.String `xml:"http://www.lido-schema.org type,attr,omitempty"`
}

type RelatedEvent struct {
	// Display and index elements for the event related to the event being recorded.
	RelatedEvent *EventElement `xml:"http://www.lido-schema.org relatedEvent"`

	// A term describing the nature of the relationship between the described
	// event and the related event. Example values: part of, influence of,
	// related to.Indicate a term characterizing the relationship from the
	// perspective of the currently described event towards the related event.
	// Preferably taken from a published controlled vocabulary.
	// Notes: For implementation of the data, note that relationships are
	// conceptually reciprocal, but the Relationship Type is often different on
	// either side of the relationship.
	RelatedEventRelType *ConceptElement `xml:"http://www.lido-schema.org relatedEventRelType"`

	// Assigns a priority order for online presentation of the element. Has to be
	// a positive integer, with descending priority from 1 to x.
	SortOrder xsdt.Integer `xml:"http://www.lido-schema.org sortorder,attr,omitempty"`
}

type EventMaterialsTech struct {
	// Display element for materials/technique, corresponding to the following
	// materialsTech element. It is presented in a syntax suitable for display to
	// the end-user and including any necessary indications of uncertainty,
	// ambiguity, and nuance.Repeat this element only for language variants.
	DisplayMaterialsTechs []*Text `xml:"http://www.lido-schema.org displayMaterialsTech"`

	// Materials and techniques data used for indexing.
	MaterialsTech *MaterialsTech `xml:"http://www.lido-schema.org materialsTech"`

	// Assigns a priority order for online presentation of the element. Has to be
	// a positive integer, with descending priority from 1 to x.
	SortOrder xsdt.Integer `xml:"http://www.lido-schema.org sortorder,attr,omitempty"`
}

type MaterialsTech struct {
	XMLName xml.Name `xml:"http://www.lido-schema.org materialsTech"`
	// A concept to index materials and/or technique. Preferably taken from a
	// published controlled vocabulary.
	TermMaterialsTechs []*ClassificationElement `xml:"http://www.lido-schema.org termMaterialsTech"`

	// An explanation of the part of the object / work to which the corresponding
	// materials or technique are applicable; included when necessary for clarity.
	ExtentMaterialsTechs []*Text `xml:"http://www.lido-schema.org extentMaterialsTech"`

	// The source of the information about materials and technique, often used
	//when citing a published source of watermarks.
	SourceMaterialsTechs []*Text `xml:"http://www.lido-schema.org sourceMaterialsTech"`
}
