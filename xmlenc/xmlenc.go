package xmlenc

import (
	"encoding/xml"
	"github.com/verisart/xsd/dsig"
	"github.com/verisart/xsd/xsdt"
)

type KeySizeType xsdt.Integer

// EncryptedData *TEncryptedDataType `xml:"http://www.w3.org/2001/04/xmlenc# EncryptedData"`

// EncryptedType is the abstract type from which EncryptedData and EncryptedKey
// are derived. While these two latter element types are very similar with
// respect to their content models, a syntactical distinction is useful to
// processing. Implementation MUST generate laxly schema valid [XML-schema]
// EncryptedData or EncryptedKey as specified by the subsequent schema
// declarations. (Note the laxly schema valid generation means that the content
// permitted by xsd:ANY need not be valid.) Implementations SHOULD create these
// XML structures (EncryptedType elements and their descendents/content) in
// Normalization Form C [NFC, NFC-Corrigendum].
type Encrypted struct {
	// MimeType is an optional (advisory) attribute which describes the media type
	// of the data which has been encrypted. The value of this attribute is a
	// string with values defined by [MIME]. For example, if the data that is
	// encrypted is a base64 encoded PNG, the transfer Encoding may be specified
	// as 'http://www.w3.org/2000/09/xmldsig#base64' and the MimeType as
	// 'image/png'. This attribute is purely advisory; no validation of the
	// MimeType information is required and it does not indicate the encryption
	// application must do any additional processing. Note, this information may
	// not be necessary if it is already bound to the identifier in the Type
	// attribute. For example, the Element and Content types defined in this
	// specification are always UTF-8 encoded text.
	MimeType xsdt.String `xml:"http://www.w3.org/2001/04/xmlenc# MimeType,attr,omitempty"`

	Encoding xsdt.AnyURI `xml:"http://www.w3.org/2001/04/xmlenc# Encoding,attr,omitempty"`

	// EncryptionMethod is an optional element that describes the encryption
	// algorithm applied to the cipher data. If the element is absent, the
	// encryption algorithm must be known by the recipient or the decryption will
	// fail.
	EncryptionMethod *EncryptionMethod `xml:"http://www.w3.org/2001/04/xmlenc# EncryptionMethod"`

	// ds:KeyInfo is an optional element, defined by [XML-DSIG], that carries
	// information about the key used to encrypt the data. Subsequent sections of
	// this specification define new elements that may appear as children of
	// ds:KeyInfo.
	KeyInfo *dsig.KeyInfo `xml:"http://www.w3.org/2000/09/xmldsig# KeyInfo"`

	// CipherData is a mandatory element that contains the CipherValue or
	// CipherReference with the encrypted data.
	CipherData *CipherData `xml:"http://www.w3.org/2001/04/xmlenc# CipherData"`

	// EncryptionProperties can contain additional information concerning the
	// generation of the EncryptedType (e.g., date/time stamp).
	EncryptionProperties *EncryptionProperties `xml:"http://www.w3.org/2001/04/xmlenc# EncryptionProperties"`

	// Id is an optional attribute providing for the standard method of assigning
	// a string id to the element within the document context.
	ID xsdt.Id `xml:"http://www.w3.org/2001/04/xmlenc# Id,attr"`

	// Type is an optional attribute identifying type information about the
	// plaintext form of the encrypted content. While optional, this specification
	// takes advantage of it for mandatory processing described in Processing
	// Rules: Decryption (section 4.2). If the EncryptedData element contains data
	// of Type 'element' or element 'content', and replaces that data in an XML
	// document context, it is strongly recommended the Type attribute be
	// provided. Without this information, the decryptor will be unable to
	// automatically restore the XML document to its original cleartext form.
	Type xsdt.AnyURI `xml:"http://www.w3.org/2001/04/xmlenc# Type,attr"`
}

// The EncryptedData element is the core element in the syntax. Not only does
// its CipherData child contain the encrypted data, but it's also the element
// that replaces the encrypted element, or serves as the new document root.
type EncryptedData struct {
	Encrypted
}

// The EncryptedKey element is used to transport encryption keys from the
// originator to a known recipient(s). It may be used as a stand-alone XML
// document, be placed within an application document, or appear inside an
// EncryptedData element as a child of a ds:KeyInfo element. The key value is
// always encrypted to the recipient(s). When EncryptedKey is decrypted the
// resulting octets are made available to the EncryptionMethod algorithm without
// any additional processing.
type EncryptedKey struct {
	Encrypted

	// ReferenceList is an element that contains pointers from a key value of an
	// EncryptedKey to items encrypted by that key value (EncryptedData or
	// EncryptedKey elements).
	ReferenceList *ReferenceList `xml:"http://www.w3.org/2001/04/xmlenc# ReferenceList"`

	// The CarriedKeyName element is used to identify the encrypted key value
	// which may be referenced by the KeyName element in ds:KeyInfo. (Since ID
	// attribute values must be unique to a document,CarriedKeyName can indicate
	// that several EncryptedKey structures contain the same key value encrypted
	// for different recipients.)
	CarriedKeyName xsdt.String `xml:"http://www.w3.org/2001/04/xmlenc# CarriedKeyName"`

	// Recipient is an optional attribute that contains a hint as to which
	// recipient this encrypted key value is intended for. Its contents are
	// application dependent.
	Recipient xsdt.String `xml:"http://www.w3.org/2001/04/xmlenc# Recipient,attr"`
}

// EncryptionMethod is an optional element that describes the encryption
// algorithm applied to the cipher data. If the element is absent, the
// encryption algorithm must be known by the recipient or the decryption will
// fail.
type EncryptionMethod struct {
	KeySize KeySizeType `xml:"http://www.w3.org/2001/04/xmlenc# KeySize"`

	// The RSA-OAEP algorithm (section 5.4.2) uses the ds:DigestMethod and
	// OAEPparams elements.
	OAEPparams xsdt.Base64Binary `xml:"http://www.w3.org/2001/04/xmlenc# OAEPparams"`

	Algorithm xsdt.AnyURI `xml:"http://www.w3.org/2001/04/xmlenc# Algorithm,attr"`

	CData string `xml:",chardata"`
}

// The CipherData is a mandatory element that provides the encrypted data. It
// must either contain the encrypted octet sequence as base64 encoded text of
// the CipherValue element, or provide a reference to an external location
// containing the encrypted octet sequence via the CipherReference element.
type CipherData struct {
	// If CipherValue is not supplied directly, the CipherReference identifies a
	// source which, when processed, yields the encrypted octet sequence.
	CipherReference *CipherReference `xml:"http://www.w3.org/2001/04/xmlenc# CipherReference"`

	// The CipherData's CipherValue is an octet sequence that is processed
	// (serialized, encrypted, and encoded) by a referring encrypted object's
	// EncryptionMethod. (Note, an EncryptedKey's EncryptionMethod is the
	// algorithm used to encrypt these octets and does not speak about what type
	// of octets they are.)
	CipherValue xsdt.Base64Binary `xml:"http://www.w3.org/2001/04/xmlenc# CipherValue"`
}

// If CipherValue is not supplied directly, the CipherReference identifies a
// source which, when processed, yields the encrypted octet sequence.
type CipherReference struct {
	Transforms *dsig.Transforms `xml:"http://www.w3.org/2001/04/xmlenc# Transforms"`

	Uri xsdt.AnyURI `xml:"http://www.w3.org/2001/04/xmlenc# URI,attr"`
}

type EncryptionProperties struct {
	EncryptionProperties []*EncryptionProperty `xml:"http://www.w3.org/2001/04/xmlenc# EncryptionProperty"`

	ID xsdt.Id `xml:"http://www.w3.org/2001/04/xmlenc# Id,attr"`
}

type EncryptionProperty struct {
	Target xsdt.AnyURI `xml:"http://www.w3.org/2001/04/xmlenc# Target,attr"`

	ID xsdt.Id `xml:"http://www.w3.org/2001/04/xmlenc# Id,attr"`

	CData string `xml:",chardata"`
}

// ReferenceList is an element that contains pointers from a key value of an
// EncryptedKey to items encrypted by that key value (EncryptedData or
// EncryptedKey elements).
type ReferenceList struct {
	// DataReference elements are used to refer to EncryptedData elements that
	// were encrypted using the key defined in the enclosing EncryptedKey element.
	// Multiple DataReference elements can occur if multiple EncryptedData
	// elements exist that are encrypted by the same key.
	DataReferences []*Reference `xml:"http://www.w3.org/2001/04/xmlenc# DataReference"`

	// KeyReference elements are used to refer to EncryptedKey elements that were
	// encrypted using the key defined in the enclosing EncryptedKey element.
	// Multiple KeyReference elements can occur if multiple EncryptedKey elements
	// exist that are encrypted by the same key.
	KeyReferences []*Reference `xml:"http://www.w3.org/2001/04/xmlenc# KeyReference"`
}

type Reference struct {
	URI xsdt.AnyURI `xml:"http://www.w3.org/2001/04/xmlenc# URI,attr"`
}
