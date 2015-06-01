package dsig

import (
	"encoding/xml"
	"github.com/verisart/xsd/xsdt"
)

type KeyInfo struct {
	RetrievalMethods []*RetrievalMethod `xml:"http://www.w3.org/2000/09/xmldsig# RetrievalMethod"`

	PGPDatas []*TPGPData `xml:"http://www.w3.org/2000/09/xmldsig# PGPData"`

	SPKIDatas []*SPKIData `xml:"http://www.w3.org/2000/09/xmldsig# SPKIData"`

	MgmtDatas []xsdt.String `xml:"http://www.w3.org/2000/09/xmldsig# MgmtData"`

	CData string `xml:",chardata"`

	KeyNames []xsdt.String `xml:"http://www.w3.org/2000/09/xmldsig# KeyName"`

	KeyValues []*KeyValue `xml:"http://www.w3.org/2000/09/xmldsig# KeyValue"`

	X509Datas []*X509Data `xml:"http://www.w3.org/2000/09/xmldsig# X509Data"`

	Id xsdt.Id `xml:"http://www.w3.org/2000/09/xmldsig# Id,attr"`
}

type RetrievalMethod struct {
	Transforms *Transforms `xml:"http://www.w3.org/2000/09/xmldsig# Transforms"`

	Uri xsdt.AnyURI `xml:"http://www.w3.org/2000/09/xmldsig# URI,attr"`

	Type xsdt.AnyURI `xml:"http://www.w3.org/2000/09/xmldsig# Type,attr"`
}

type Transforms struct {
	Transforms []*Transform `xml:"http://www.w3.org/2000/09/xmldsig# Transform"`
}

type Transform struct {
	CData string `xml:",chardata"`

	XPaths []xsdt.String `xml:"http://www.w3.org/2000/09/xmldsig# XPath"`

	Algorithm xsdt.AnyURI `xml:"http://www.w3.org/2000/09/xmldsig# Algorithm,attr"`
}

type PGPData struct {
	PGPKeyPacket xsdt.Base64Binary `xml:"http://www.w3.org/2000/09/xmldsig# PGPKeyPacket"`

	PGPKeyID xsdt.Base64Binary `xml:"http://www.w3.org/2000/09/xmldsig# PGPKeyID"`
}

type SPKIData struct {
	SPKISexps []xsdt.Base64Binary `xml:"http://www.w3.org/2000/09/xmldsig# SPKISexp"`
}

type KeyValue struct {
	DSAKeyValue *DSAKeyValue `xml:"http://www.w3.org/2000/09/xmldsig# DSAKeyValue"`

	RSAKeyValue *RSAKeyValue `xml:"http://www.w3.org/2000/09/xmldsig# RSAKeyValue"`

	CData string `xml:",chardata"`
}

type DSAKeyValue struct {
	// TODO:
	/*
	  XsdGoPkgHasElem_YsequenceDSAKeyValueTypeschema_Y_TCryptoBinary_

	  XsdGoPkgHasElem_JsequenceDSAKeyValueTypeschema_J_TCryptoBinary_

	  XsdGoPkgHasElem_PsequencesequenceDSAKeyValueTypeschema_P_TCryptoBinary_

	  XsdGoPkgHasElem_QsequencesequenceDSAKeyValueTypeschema_Q_TCryptoBinary_

	  XsdGoPkgHasElem_SeedsequencesequenceDSAKeyValueTypeschema_Seed_TCryptoBinary_

	  XsdGoPkgHasElem_PgenCountersequencesequenceDSAKeyValueTypeschema_PgenCounter_TCryptoBinary_

	  XsdGoPkgHasElem_GsequenceDSAKeyValueTypeschema_G_TCryptoBinary_
	*/
}

type RSAKeyValue struct {
	// TODO:
	/*
	  XsdGoPkgHasElem_ModulussequenceRSAKeyValueTypeschema_Modulus_TCryptoBinary_

	  XsdGoPkgHasElem_ExponentsequenceRSAKeyValueTypeschema_Exponent_TCryptoBinary_
	*/
}

type X509Data struct {
	// TODO:
	/*
	  XsdGoPkgHasElem_X509IssuerSerialchoicesequenceX509DataTypeschema_X509IssuerSerial_Tx509IssuerSerialType_

	  XsdGoPkgHasElem_X509SKIchoicesequenceX509DataTypeschema_X509Ski_XsdtBase64Binary_

	  XsdGoPkgHasElem_X509SubjectNamechoicesequenceX509DataTypeschema_X509SubjectName_XsdtString_

	  XsdGoPkgHasElem_X509CertificatechoicesequenceX509DataTypeschema_X509Certificate_XsdtBase64Binary_

	  XsdGoPkgHasElem_X509CRLchoicesequenceX509DataTypeschema_X509Crl_XsdtBase64Binary_
	*/
}
