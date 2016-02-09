package mets

import (
	"encoding/xml"
	//xlink "github.com/verisart/go-prov/schema/xlink"
	//xsdt "github.com/verisart/go-prov/schema/xsdt"
	"github.com/verisart/xsd/xlink"
	"github.com/verisart/xsd/xsdt"
)

// METS Version 1.8 via http://www.loc.gov/standards/mets/version18/mets.xsd
type Mets struct {
	XMLName xml.Name `xml:"http://www.loc.gov/METS/ mets"`

	// ID (ID/O): This attribute uniquely identifies the element within the METS
	// document, and would allow the element to be referenced unambiguously from
	// another element or document via an IDREF or an XPTR. For more information
	// on using ID attributes for internal and external linking see Chapter 4 of
	// the METS Primer.
	ID xsdt.Id `xml:"ID,attr,omitempty"`

	//  OBJID (string/O): Is the primary identifier assigned to the METS object as
	// a whole. Although this attribute is not required, it is strongly
	// recommended. This identifier is used to tag the entire METS object to
	// external systems, in contrast with the ID identifier.
	ObjID xsdt.String `xml:"OBJID,attr,omitempty"`

	// LABEL (string/O): Is a simple title string used to identify the
	// object/entity being described in the METS document for the user.
	Label xsdt.String `xml:"LABEL,attr,omitempty"`

	//  TYPE (string/O): Specifies the class or type of the object, e.g.: book,
	// journal, stereograph, dataset, video, etc.
	Type xsdt.String `xml:"TYPE,attr,omitempty"`

	// PROFILE (string/O): Indicates to which of the registered profile(s) the
	// METS document conforms. For additional information about PROFILES see
	// Chapter 5 of the METS Primer.
	Profile xsdt.String `xml:"PROFILE,attr,omitempty"`

	// The mets header element <metsHdr> captures metadata about the METS
	// document itself, not the digital object the METS document encodes.
	// Although it records a more limited set of metadata, it is very similar in
	// function and purpose to the headers employed in other schema such as the
	// Text Encoding Initiative (TEI) or in the Encoded Archival Description (EAD).
	MetsHdr *MetsHdr `xml:"metsHdr,omitempty"`

	// The administrative metadata section <amdSec> contains the administrative
	// metadata pertaining to the digital object, its components and any original
	// source material from which the digital object is derived. The <amdSec> is
	// separated into four sub-sections that accommodate technical metadata
	// (techMD), intellectual property rights (rightsMD), analog/digital source
	// metadata (sourceMD), and digital provenance metadata (digiprovMD). Each of
	// these subsections can either wrap the metadata  (mdWrap) or reference it in
	// an external location (mdRef) or both. Multiple instances of the <amdSec>
	// element can occur within a METS document and multiple instances of its
	// subsections can occur in one <amdSec> element. This allows considerable
	// flexibility in the structuring of the administrative metadata. METS does
	// not define a vocabulary or syntax for encoding administrative metadata.
	// Administrative metadata can be expressed within the amdSec sub-elements
	// according to many current community defined standards, or locally produced
	// XML schemas.
	//AmdSecs []*TamdSecType `xml:"http://www.loc.gov/METS/ amdSec"`

	// The overall purpose of the content file section element <fileSec> is to
	// provide an inventory of and the location for the content files that
	// comprise the digital object being described in the METS document.
	FileSec *MetsFileSec `xml:"fileSec,omitempty"`

	// The structural map section <structMap> is the heart of a METS document. It
	// provides a means for organizing the digital content represented by the
	// <file> elements in the <fileSec> of the METS document into a coherent
	// hierarchical structure. Such a hierarchical structure can be presented to
	// users to facilitate their comprehension and navigation of the digital
	// content. It can further be applied to any purpose requiring an
	// understanding of the structural relationship of the content files or parts
	// of the content files. The organization may be specified to any level of
	// granularity (intellectual and or physical) that is desired. Since the
	// <structMap> element is repeatable, more than one organization can be
	// applied to the digital content represented by the METS document.  The
	// hierarchical structure specified by a <structMap> is encoded as a tree of
	// nested <div> elements. A <div> element may directly point to content via
	// child file pointer <fptr> elements (if the content is represented in the
	// <fileSec<) or child METS pointer <mptr> elements (if the content is
	// represented by an external METS document). The <fptr> element may point to
	// a single whole <file> element that manifests its parent <div<, or to part
	// of a <file> that manifests its <div<. It can also point to multiple files
	// or parts of files that must be played/displayed either in sequence or in
	// parallel to reveal its structural division. In addition to providing a
	// means for organizing content, the <structMap> provides a mechanism for
	// linking content at any hierarchical level with relevant descriptive and
	// administrative metadata.
	//StructMaps []*TstructMapType `xml:"http://www.loc.gov/METS/ structMap"`

	// The structural link section element <structLink> allows for the
	// specification of hyperlinks between the different components of a METS
	// structure that are delineated in a structural map. This element is a
	// container for a single, repeatable element, <smLink> which indicates a
	// hyperlink between two nodes in the structural map. The <structLink> section
	// in the METS document is identified using its XML ID attributes.
	//StructLink *TxsdMetsTypeSequenceStructLink `xml:"http://www.loc.gov/METS/ structLink"`

	// A behavior section element <behaviorSec> associates executable behaviors
	// with content in the METS document by means of a repeatable behavior
	// <behavior> element. This element has an interface definition <interfaceDef>
	// element that represents an abstract definition of the set of behaviors
	// represented by a particular behavior section. A <behavior> element also has
	// a <mechanism> element which is used to point to a module of executable code
	// that implements and runs the behavior defined by the interface definition.
	// The <behaviorSec> element, which is repeatable as well as nestable, can be
	// used to group individual behaviors within the structure of the METS
	// document. Such grouping can be useful for organizing families of behaviors
	// together or to indicate other relationships between particular behaviors.
	//BehaviorSecs []*TbehaviorSecType `xml:"http://www.loc.gov/METS/ behaviorSec"`

	//  A descriptive metadata section <dmdSec> records descriptive metadata
	// pertaining to the METS object as a whole or one of its components. The
	// <dmdSec> element conforms to same generic datatype as the <techMD>,
	// <rightsMD>, <sourceMD> and <digiprovMD> elements, and supports the same
	// sub-elements and attributes. A descriptive metadata element can either wrap
	// the metadata  (mdWrap) or reference it in an external location (mdRef) or
	// both.  METS allows multiple <dmdSec> elements; and descriptive metadata can
	// be associated with any METS element that supports a DMDID attribute.
	// Descriptive metadata can be expressed according to many current description
	// standards (i.e., MARC, MODS, Dublin Core, TEI Header, EAD, VRA, FGDC, DDI)
	// or a locally produced XML schema.
	//DmdSecs []*TmdSecType `xml:"http://www.loc.gov/METS/ dmdSec"`

}

type MetsHdr struct {

	//  CREATEDATE (dateTime/O): Records the date/time the METS document was
	// created.
	CreateDate xsdt.DateTime `xml:"CREATEDATE,attr,omitempty"`

	//  LASTMODDATE (dateTime/O): Is used to indicate the date/time the METS
	// document was last modified.
	LastModDate xsdt.DateTime `xml:"LASTMODDATE,attr,omitempty"`

	//  RECORDSTATUS (string/O): Specifies the status of the METS document. It is
	// used for internal processing purposes.
	RecordStatus xsdt.String `xml:"RECORDSTATUS,attr,omitempty"`

	//  ID (ID/O): This attribute uniquely identifies the element within the METS
	// document, and would allow the element to be referenced unambiguously from
	// another element or document via an IDREF or an XPTR. For more information
	// on using ID attributes for internal and external linking see Chapter 4 of
	// the METS Primer.
	ID xsdt.Id `xml:"ID,attr,omitempty"`

	//  agent:
	//  The agent element <agent> provides for various parties and their roles
	// with respect to the METS record to be documented.
	Agents []*MetsAgent `xml:"agent,omitempty"`

	//  The alternative record identifier element <altRecordID> allows one to use
	// alternative record identifier values for the digital object represented by
	// the METS document; the primary record identifier is stored in the OBJID
	// attribute in the root <mets> element.
	AltRecordIDs []*MetsAltRecordID `xml:"altRecordID,omitempty"`

	//  The metsDocument identifier element <metsDocumentID> allows a unique
	// identifier to be assigned to the METS document itself.  This may be
	// different from the OBJID attribute value in the root <mets> element, which
	// uniquely identifies the entire digital object represented by the METS
	// document.
	MetsDocumentID *MetsDocumentID `xml:"metsDocumentID,omitempty"`

	//  ADMID (IDREFS/O): Contains the ID attribute values of the <techMD>,
	// <sourceMD>, <rightsMD> and/or <digiprovMD> elements within the <amdSec> of
	// the METS document that contain administrative metadata pertaining to the
	// METS document itself.  For more information on using METS IDREFS and IDREF
	// type attributes for internal linking, see Chapter 4 of the METS Primer.
	AdmID xsdt.Idrefs `xml:"ADMID,attr,omitempty"`
}

//  The alternative record identifier element <altRecordID> allows one to use
// alternative record identifier values for the digital object represented by
// the METS document; the primary record identifier is stored in the OBJID
// attribute in the root <mets> element.
type MetsDocumentID struct {

	//  ID (ID/O): This attribute uniquely identifies the element within the METS
	// document, and would allow the element to be referenced unambiguously from
	// another element or document via an IDREF or an XPTR. For more information
	// on using ID attributes for internal and external linking see Chapter 4 of
	// the METS Primer.
	ID xsdt.Id `xml:"ID,attr,omitempty"`

	//  TYPE (string/O): A description of the identifier type (e.g., OCLC record
	// number, LCCN, etc.).
	Type xsdt.String `xml:"TYPE,attr,omitempty"`
}

//  The alternative record identifier element <altRecordID> allows one to use
// alternative record identifier values for the digital object represented by
// the METS document; the primary record identifier is stored in the OBJID
// attribute in the root <mets> element.
type MetsAltRecordID struct {

	//  ID (ID/O): This attribute uniquely identifies the element within the METS
	// document, and would allow the element to be referenced unambiguously from
	// another element or document via an IDREF or an XPTR. For more information
	// on using ID attributes for internal and external linking see Chapter 4 of
	// the METS Primer.
	ID xsdt.Id `xml:"ID,attr,omitempty"`

	//  TYPE (string/O): A description of the identifier type (e.g., OCLC record
	// number, LCCN, etc.).
	Type xsdt.String `xml:"TYPE,attr,omitempty"`
}

type MetsOtherType xsdt.String

//  INDIVIDUAL | ORGANIZATION | OTHER
type MetsAgentType xsdt.String

// CREATOR | EDITOR | ARCHIVIST | PRESERVATION | DISSEMINATOR | CUSTODIAN |
// IPOWNER | OTHER
type MetsAgentRoleType xsdt.String

type MetsAgent struct {

	//  The <note> element can be used to record any additional information
	// regarding the agent's activities with respect to the METS document.
	Notes []xsdt.String `xml:"note,omitempty"`

	//  OTHERTYPE (string/O): Specifies the type of agent when the value OTHER is
	// indicated in the TYPE attribute.
	OtherType MetsOtherType `xml:"OTHERTYPE,attr,omitempty"`

	//  ID (ID/O): This attribute uniquely identifies the element within the METS
	// document, and would allow the element to be referenced unambiguously from
	// another element or document via an IDREF or an XPTR. For more information
	// on using ID attributes for internal and external linking see Chapter 4 of
	// the METS Primer.
	ID xsdt.Id `xml:"ID,attr,omitempty"`

	//  ROLE (string/R): Specifies the function of the agent with respect to the METS record. The allowed values are:
	//  CREATOR: The person(s) or institution(s) responsible for the METS document.
	//  EDITOR: The person(s) or institution(s) that prepares the metadata for encoding.
	//  ARCHIVIST: The person(s) or institution(s) responsible for the document/collection.
	//  PRESERVATION: The person(s) or institution(s) responsible for preservation functions.
	//  DISSEMINATOR: The person(s) or institution(s) responsible for dissemination functions.
	//  CUSTODIAN: The person(s) or institution(s) charged with the oversight of a document/collection.
	//  IPOWNER: Intellectual Property Owner: The person(s) or institution holding
	// copyright, trade or service marks or other intellectual property rights for
	// the object.
	//  OTHER: Use OTHER if none of the preceding values pertains and clarify the
	// type and location specifier being used in the OTHERROLE attribute (see below).
	Role MetsAgentRoleType `xml:"ROLE,attr"`

	//  OTHERROLE (string/O): Denotes a role not contained in the allowed values
	// set if OTHER is indicated in the ROLE attribute.
	OtherRole xsdt.String `xml:"OTHERROLE,attr,omitempty"`

	//  TYPE (string/O): is used to specify the type of AGENT. It must be one of the following values:
	//  INDIVIDUAL: Use if an individual has served as the agent.
	//  ORGANIZATION: Use if an institution, corporate body, association, non-profit enterprise, government, religious body, etc. has served as the agent.
	//  OTHER: Use OTHER if none of the preceding values pertain and clarify the type of agent specifier being used in the OTHERTYPE attribute
	Type MetsAgentType `xml:"TYPE,attr,omitempty"`

	//  The element <name> can be used to record the full name of the document agent.
	Name xsdt.String `xml:"name,omitempty"`
}

type MetsFileSec struct {

	//  A sequence of file group elements <fileGrp> can be used group the digital
	// files comprising the content of a METS object either into a flat
	// arrangement or, because each file group element can itself contain one or
	// more  file group elements,  into a nested (hierarchical) arrangement. In
	// the case where the content files are images of different formats and
	// resolutions, for example, one could group the image content files by format
	// and create a separate <fileGrp> for each image format/resolution such as:
	//  -- one <fileGrp> for the thumbnails of the images
	//  -- one <fileGrp> for the higher resolution JPEGs of the image
	//  -- one <fileGrp> for the master archival TIFFs of the images
	//  For a text resource with a variety of content file types one might group
	// the content files at the highest level by type,  and then use the <fileGrp>
	// element’s nesting capabilities to subdivide a <fileGrp> by format within
	// the type, such as:
	//  -- one <fileGrp> for all of the page images with nested <fileGrp> elements
	// for each image format/resolution (tiff, jpeg, gif)
	//  -- one <fileGrp> for a PDF version of all the pages of the document
	//  -- one <fileGrp> for  a TEI encoded XML version of the entire document or
	// each of its pages.
	//  A <fileGrp> may contain zero or more <fileGrp> elements and or <file> elements.
	FileGrps []*MetsFileGrp `xml:"fileGrp,omitempty"`

	//  ID (ID/O): This attribute uniquely identifies the element within the METS
	// document, and would allow the element to be referenced unambiguously from
	// another element or document via an IDREF or an XPTR. For more information
	// on using ID attributes for internal and external linking see Chapter 4 of
	// the METS Primer.
	ID xsdt.Id `xml:"ID,attr,omitempty,omitempty"`
}

type MetsFileGrp struct {
	//  ADMID (IDREFS/O): Contains the ID attribute values of the <techMD>,
	// <sourceMD>, <rightsMD> and/or <digiprovMD> elements within the <amdSec> of
	// the METS document that contain administrative metadata pertaining to the
	// METS document itself.  For more information on using METS IDREFS and IDREF
	// type attributes for internal linking, see Chapter 4 of the METS Primer.
	AdmID xsdt.Idrefs `xml:"ADMID,attr,omitempty"`

	//  USE (string/O): A tagging attribute to indicate the intended use of files
	// within this file group (e.g., master, reference, thumbnails for image
	// files). A USE attribute can be expressed at the<fileGrp> level, the <file>
	// level, the <FLocat> level and/or the <FContent> level.  A USE attribute
	// value at the <fileGrp> level should pertain to all of the files in the
	// <fileGrp>.  A USE attribute at the <file> level should pertain to all
	// copies of the file as represented by subsidiary <FLocat> and/or <FContent>
	// elements.  A USE attribute at the <FLocat> or <FContent> level pertains to
	// the particular copy of the file that is either referenced (<FLocat>) or
	// wrapped (<FContent>).
	Use xsdt.String `xml:"USE,attr,omitempty"`

	// Recursively nested file groups.
	FileGrps []*MetsFileGrp `xml:"fileGrp,omitempty"`

	//  The file element <file> provides access to the content files for the
	// digital object being described by the METS document. A <file> element may
	// contain one or more <FLocat> elements which provide pointers to a content
	// file and/or a <FContent> element which wraps an encoded version of the file.
	// Embedding files using <FContent> can be a valuable feature for exchanging
	// digital objects between repositories or for archiving versions of digital
	// objects for off-site storage. All <FLocat> and <FContent> elements should
	// identify and/or contain identical copies of a single file. The <file>
	// element is recursive, thus allowing sub-files or component files of a
	// larger file to be listed in the inventory. Alternatively, by using the
	// <stream> element, a smaller component of a file or of a related file can be
	// placed within a <file> element. Finally, by using the <transformFile>
	// element, it is possible to include within a <file> element a different
	// version of a file that has undergone a transformation for some reason, such
	// as format migration.
	Files []*MetsFile `xml:"file,omitempty"`

	//  ID (ID/O): This attribute uniquely identifies the element within the METS
	// document, and would allow the element to be referenced unambiguously from
	// another element or document via an IDREF or an XPTR. For more information
	// on using ID attributes for internal and external linking see Chapter 4 of
	// the METS Primer.
	ID xsdt.Id `xml:"ID,attr,omitempty"`

	//  VERSDATE (dateTime/O): An optional dateTime attribute specifying the date
	// this version/fileGrp of the digital object was created.
	VersDate xsdt.DateTime `xml:"VERSDATE,attr,omitempty"`
}

type MetsFile struct {
	//  The file content element <FContent> is used to identify a content file
	// contained internally within a METS document. The content file must be
	// either Base64 encoded and contained within the subsidiary <binData> wrapper
	// element, or consist of XML information and be contained within the
	// subsidiary <xmlData> wrapper element.
	FContent *MetsFileContent `xml:"FContent,omitempty"`

	//  A component byte stream element <stream> may be composed of one or more
	// subsidiary streams. An MPEG4 file, for example, might contain separate
	// audio and video streams, each of which is associated with technical
	// metadata. The repeatable <stream> element provides a mechanism to record
	// the existence of separate data streams within a particular file, and the
	// opportunity to associate <dmdSec> and <amdSec> with those subsidiary data
	// streams if desired.
	Streams []*MetsStream `xml:"stream,omitempty"`

	Files []*MetsFile `xml:"file,omitempty"`

	MetsFileCore

	//  GROUPID (string/O): An identifier that establishes a correspondence
	// between this file and files in other file groups. Typically, this will be
	// used to associate a master file in one file group with the derivative files
	// made from it in other file groups.
	Groupid xsdt.String `xml:"GROUPID,attr"`

	//  OWNERID (string/O): Used to provide a unique identifier (which could
	// include a URI) assigned to the file. This identifier may differ from the
	// URI used to retrieve the file.
	OwnerID xsdt.String `xml:"OWNERID,attr,omitempty"`

	//  The file location element <FLocat> provides a pointer to the location of a
	// content file. It uses the XLink reference syntax to provide linking
	// information indicating the actual location of the content file, along with
	// other attributes specifying additional linking information. NOTE: <FLocat>
	// is an empty element. The location of the resource pointed to MUST be stored
	// in the xlink:href attribute.
	FLocats []*MetsFLocat `xml:"FLocat"`

	//  ADMID (IDREFS/O): Contains the ID attribute values of the <techMD>,
	// <sourceMD>, <rightsMD> and/or <digiprovMD> elements within the <amdSec> of
	// the METS document that contain administrative metadata pertaining to the
	// METS document itself.  For more information on using METS IDREFS and IDREF
	// type attributes for internal linking, see Chapter 4 of the METS Primer.
	AdmID xsdt.Idrefs `xml:"ADMID,attr,omitempty"`

	//  DMDID (IDREFS/O): Contains the ID attribute values identifying the
	// <dmdSec>, elements in the METS document that contain or link to descriptive
	// metadata pertaining to the content file stream represented by the current
	// <stream> element.  For more information on using METS IDREFS and IDREF type
	// attributes for internal linking, see Chapter 4 of the METS Primer.
	DmdID xsdt.Idrefs `xml:"DMDID,attr,omitempty"`

	//  USE (string/O): A tagging attribute to indicate the intended use of the
	// specific copy of the file represented by the <FContent> element (e.g.,
	// service master, archive master). A USE attribute can be expressed at the
	// <fileGrp> level, the <file> level, the <FLocat> level and/or the <FContent>
	// level.  A USE attribute value at the <fileGrp> level should pertain to all
	// of the files in the <fileGrp>.  A USE attribute at the <file> level should
	// pertain to all copies of the file as represented by subsidiary <FLocat>
	// and/or <FContent> elements.  A USE attribute at the <FLocat> or <FContent>
	// level pertains to the particular copy of the file that is either referenced
	// (<FLocat>) or wrapped (<FContent>).
	Use xsdt.String `xml:"USE,attr,omitempty"`

	//  The transform file element <transformFile> provides a means to access any
	// subsidiary files listed below a <file> element by indicating the steps
	// required to "unpack" or transform the subsidiary files. This element is
	// repeatable and might provide a link to a <behavior> in the <behaviorSec>
	// that performs the transformation.
	TransformFiles []*MetsTransformFile `xml:"transformFile,omitempty"`

	//  SEQ (integer/O): Indicates the sequence of this <file> relative to the
	// others in its <fileGrp>.
	Seq xsdt.Int `xml:"SEQ,attr,omitempty"`

	//  BEGIN (string/O): An attribute that specifies the point in the parent
	// <file> where the current <stream> begins. It can be used in conjunction
	// with the END attribute as a means of defining the location of the stream
	// within its parent file. However, the BEGIN attribute can be used with or
	// without a companion END attribute. When no END attribute is specified,
	// the end of the parent file is assumed also to be the end point of the
	// stream. The BEGIN and END attributes can only be interpreted meaningfully
	// in conjunction with a BETYPE attribute, which specifies the kind of
	// beginning/ending point values that are being used.
	Begin xsdt.String `xml:"BEGIN,attr,omitempty"`

	//  END (string/O): An attribute that specifies the point in the parent <file>
	// where the <stream> ends. It can only be interpreted meaningfully in
	// conjunction with the BETYPE, which specifies the kind of ending point
	// values being used. Typically the END attribute would only appear in
	// conjunction with a BEGIN attribute.
	End xsdt.String `xml:"END,attr,omitempty"`

	//  BETYPE: Begin/End Type.
	//  BETYPE (string/O): An attribute that specifies the kind of BEGIN and/or
	// END values that are being used. Currently BYTE is the only valid value that
	// can be used in conjunction with nested <file> or <stream> elements.
	BEType MetsBEType `xml:"BETYPE,attr,omitempty"`

	//  ID (ID/R): This attribute uniquely identifies the element within the METS
	// document, and would allow the element to be referenced unambiguously from
	// another element or document via an IDREF or an XPTR. For more information
	// on using ID attributes for internal and external linking see Chapter 4 of
	// the METS Primer.
	ID xsdt.Id `xml:"ID,attr"`
}

type MetsTransformType xsdt.String

type MetsTransformFile struct {
	//  TRANSFORMBEHAVIOR (string/O): An IDREF to a behavior element for this
	// transformation.
	TransformBehavior xsdt.Idref `xml:"TRANSFORMBEHAVIOR,attr,omitempty"`

	//XsdtAnyType

	//  TRANSFORMORDER (postive-integer/R): The order in which the instructions
	// must be followed in order to unpack or transform the container file.
	TransformOrder xsdt.PositiveInteger `xml:"TRANSFORMORDER,attr"`

	//  ID (ID/O): This attribute uniquely identifies the element within the METS
	// document, and would allow the element to be referenced unambiguously from
	// another element or document via an IDREF or an XPTR. For more information
	// on using ID attributes for internal and external linking see Chapter 4 of
	// the METS Primer.
	ID xsdt.Id `xml:"ID,attr"`

	//  TRANSFORMTYPE (string/R): Is used to indicate the type of transformation
	// needed to render content of a file accessible. This may include unpacking a
	// file into subsidiary files/streams. The controlled value constraints for
	// this XML string include “decompression” and “decryption”. Decompression is
	// defined as the action of reversing data compression, i.e., the process of
	// encoding information using fewer bits than an unencoded representation
	// would use by means of specific encoding schemas. Decryption is defined as
	// the process of restoring data that has been obscured to make it unreadable
	// without special knowledge (encrypted data) to its original form.
	Transformtype MetsTransformType `xml:"TRANSFORMTYPE,attr"`

	//  TRANSFORM-ALGORITHM (string/R): Specifies the decompression or decryption
	// routine used to access the contents of the file. Algorithms for compression
	// can be either loss-less or lossy.
	TransformAlgorithm xsdt.String `xml:"TRANSFORMALGORITHM,attr"`

	//  TRANSFORMKEY (string/O): A key to be used with the transform algorithm for
	// accessing the file’s contents.
	TransformKey xsdt.String `xml:"TRANSFORMKEY,attr,omitempty"`
}

type MetsXMLFileContent struct {
}

type MetsFileContent struct {

	//  A binary data wrapper element <binData> is used to contain a Base64 encoded file.
	BinData xsdt.Base64Binary `xml:"binData,omitempty"`

	//  An xml data wrapper element <xmlData> is used to contain  an XML encoded file. The content of an <xmlData> element can be in any namespace or in no namespace. As permitted by the XML Schema Standard, the processContents attribute value for the metadata in an <xmlData> element is set to “lax”. Therefore, if the source schema and its location are identified by means of an xsi:schemaLocation attribute, then an XML processor will validate the elements for which it can find declarations. If a source schema is not identified, or cannot be found at the specified schemaLocation, then an XML validator will check for well-formedness, but otherwise skip over the elements appearing in the <xmlData> element.
	XmlData *MetsXMLFileContent `xml:"xmlData,omitempty"`

	//  ID (ID/O): This attribute uniquely identifies the element within the METS
	// document, and would allow the element to be referenced unambiguously from
	// another element or document via an IDREF or an XPTR. For more information
	// on using ID attributes for internal and external linking see Chapter 4 of
	// the METS Primer.
	ID xsdt.Id `xml:"ID,attr,omitempty"`

	//  USE (string/O): A tagging attribute to indicate the intended use of the
	// specific copy of the file represented by the <FContent> element (e.g.,
	// service master, archive master). A USE attribute can be expressed at the
	// <fileGrp> level, the <file> level, the <FLocat> level and/or the <FContent>
	// level.  A USE attribute value at the <fileGrp> level should pertain to all
	// of the files in the <fileGrp>.  A USE attribute at the <file> level should
	// pertain to all copies of the file as represented by subsidiary <FLocat>
	// and/or <FContent> elements.  A USE attribute at the <FLocat> or <FContent>
	// level pertains to the particular copy of the file that is either referenced
	// (<FLocat>) or wrapped (<FContent>).
	Use xsdt.String `xml:"USE,attr,omitempty"`
}

// Currently BYTE is the only valid value that can be used in conjunction with
// nested <file> or <stream> elements.
// BYTE | IDREF | SMIL | MIDI | SMPTE-25 | SMPTE-24 | SMPTE-DF30 | SMPTE-NDF30 | SMPTE-DF29.97 | SMPTE-NDF29.97 | TIME | TCF | XPTR
type MetsBEType xsdt.String

type MetsStream struct {

	//  BEGIN (string/O): An attribute that specifies the point in the parent
	// <file> where the current <stream> begins. It can be used in conjunction
	// with the END attribute as a means of defining the location of the stream
	// within its parent file. However, the BEGIN attribute can be used with or
	// without a companion END attribute. When no END attribute is specified,
	// the end of the parent file is assumed also to be the end point of the
	// stream. The BEGIN and END attributes can only be interpreted meaningfully
	// in conjunction with a BETYPE attribute, which specifies the kind of
	// beginning/ending point values that are being used.
	Begin xsdt.String `xml:"BEGIN,attr,omitempty"`

	//  END (string/O): An attribute that specifies the point in the parent <file>
	// where the <stream> ends. It can only be interpreted meaningfully in
	// conjunction with the BETYPE, which specifies the kind of ending point
	// values being used. Typically the END attribute would only appear in
	// conjunction with a BEGIN attribute.
	End xsdt.String `xml:"END,attr,omitempty"`

	//  BETYPE: Begin/End Type.
	//  BETYPE (string/O): An attribute that specifies the kind of BEGIN and/or
	// END values that are being used. Currently BYTE is the only valid value that
	// can be used in conjunction with nested <file> or <stream> elements.
	BEType MetsBEType `xml:"BETYPE,attr,omitempty"`

	//  ID (ID/O): This attribute uniquely identifies the element within the METS
	// document, and would allow the element to be referenced unambiguously from
	// another element or document via an IDREF or an XPTR. For more information
	// on using ID attributes for internal and external linking see Chapter 4 of
	// the METS Primer.
	ID xsdt.Id `xml:"ID,attr,omitempty"`

	//  ADMID (IDREFS/O): Contains the ID attribute values of the <techMD>,
	// <sourceMD>, <rightsMD> and/or <digiprovMD> elements within the <amdSec> of
	// the METS document that contain administrative metadata pertaining to the
	// METS document itself.  For more information on using METS IDREFS and IDREF
	// type attributes for internal linking, see Chapter 4 of the METS Primer.
	AdmID xsdt.Idrefs `xml:"ADMID,attr,omitempty"`

	// TODO: not sure if this is for the xml content?
	//XsdtAnyType

	//  DMDID (IDREFS/O): Contains the ID attribute values identifying the
	// <dmdSec>, elements in the METS document that contain or link to descriptive
	// metadata pertaining to the content file stream represented by the current
	// <stream> element.  For more information on using METS IDREFS and IDREF type
	// attributes for internal linking, see Chapter 4 of the METS Primer.
	DmdID xsdt.Idrefs `xml:"DMDID,attr,omitempty"`

	//  streamType (string/O): The IANA MIME media type for the bytestream.
	StreamType xsdt.String `xml:"streamType,attr,omitempty"`

	//  OWNERID (string/O): Used to provide a unique identifier (which could
	// include a URI) assigned to the file. This identifier may differ from the
	// URI used to retrieve the file.
	OwnerID xsdt.String `xml:"OWNERID,attr,omitempty"`
}

type MetsChecksumType xsdt.String

type MetsFileCore struct {

	// MIMETYPE (string/O): The IANA MIME media type for the associated file or
	// wrapped content. Some values for this attribute can be found on the IANA
	// website.
	MimeType xsdt.String `xml:"MIMETYPE,attr,omitempty"`

	// SIZE (long/O): Specifies the size in bytes of the associated file or
	// wrapped content.
	Size xsdt.Long `xml:"SIZE,attr,omitempty"`

	// CREATED (dateTime/O): Specifies the date and time of creation for the
	// associated file or wrapped content.
	Created xsdt.DateTime `xml:"CREATED,attr,omitempty"`

	// CHECKSUM (string/O): Provides a checksum value for the associated file or
	// wrapped content.
	Checksum xsdt.String `xml:"CHECKSUM,attr,omitempty"`

	// CHECKSUMTYPE (enumerated string/O): Specifies the checksum algorithm used
	// to produce the value contained in the CHECKSUM attribute.  CHECKSUMTYPE
	// must contain one of the following values:
	//  Adler-32
	//  CRC32
	//  HAVAL
	//  MD5
	//  MNP
	//  SHA-1
	//  SHA-256
	//  SHA-384
	//  SHA-512
	//  TIGER
	//  WHIRLPOOL
	ChecksumType MetsChecksumType `xml:"CHECKSUMTYPE,attr,omitempty"`
}

type MetsFLocat struct {
	MetsLocation

	xlink.SimpleLink

	//  ID (ID/O): This attribute uniquely identifies the element within the METS
	// document, and would allow the element to be referenced unambiguously from
	// another element or document via an IDREF or an XPTR. For more information
	// on using ID attributes for internal and external linking see Chapter 4 of
	// the METS Primer.
	ID xsdt.Id `xml:"ID,attr,omitempty"`

	//  USE (string/O): A tagging attribute to indicate the intended use of the
	// specific copy of the file represented by the <FContent> element (e.g.,
	// service master, archive master). A USE attribute can be expressed at the
	// <fileGrp> level, the <file> level, the <FLocat> level and/or the <FContent>
	// level.  A USE attribute value at the <fileGrp> level should pertain to all
	// of the files in the <fileGrp>.  A USE attribute at the <file> level should
	// pertain to all copies of the file as represented by subsidiary <FLocat>
	// and/or <FContent> elements.  A USE attribute at the <FLocat> or <FContent>
	// level pertains to the particular copy of the file that is either referenced
	// (<FLocat>) or wrapped (<FContent>).
	Use xsdt.String `xml:"USE,attr,omitempty"`
}

//  LOCTYPE (string/R): Specifies the locator type used in the xlink:href attribute. Valid values for LOCTYPE are:
//  ARK
//  URN
//  URL
//  PURL
//  HANDLE
//  DOI
//  OTHER
type MetsLocationLocType xsdt.String

type MetsLocation struct {
	//  LOCTYPE (string/R): Specifies the locator type used in the xlink:href
	// attribute.
	LocType MetsLocationLocType `xml:"LOCTYPE,attr"`

	//  OTHERLOCTYPE (string/O): Specifies the locator type when the value OTHER
	// is used in the LOCTYPE attribute. Although optional, it is strongly
	// recommended when OTHER is used.
	OtherLocType xsdt.String `xml:"OTHERLOCTYPE,attr,omitempty"`
}
