package models

import "encoding/xml"

type FPSSignOn002 struct {
	XMLName  xml.Name `xml:"FpsMsg"`
	NumOfMsg string   `xml:"NbOfMsgs"`
	XMLHead  Head001  `xml:"FpsPylds>BizData>AppHdr"`
	Document Admn002  `xml:"FpsPylds>BizData>Document"`
}

type Admn002 struct {
	MsgId       string `xml:"AdmnSignOnResp>GrpHdr>MsgId,omitempty" validate:"required,max=35"`
	CreDtTm     string `xml:"AdmnSignOnResp>GrpHdr>CreDtTm,omitempty" validate:"required"`
	InstrId     string `xml:"AdmnSignOnResp>SignOnResp>OrgInstrId,omitempty" validate:"required"  validate:"required,max=35"`
	Instructing string `xml:"AdmnSignOnResp>SignOnResp>InstgAgt>FinInstnId>ClrSysMmbId>MmbId,omitempty" validate:"required"`
	Instructed  string `xml:"AdmnSignOnResp>SignOnResp>InstdAgt>FinInstnId>ClrSysMmbId>MmbId,omitempty"`
	Sts         string `xml:"AdmnSignOnResp>SignOnResp>Sts,omitempty"`
	Cd          string `xml:"AdmnSignOnResp>SignOnResp>StsRsnInf>Rsn>Cd,omitempty"`
	Prtry       string `xml:"AdmnSignOnResp>SignOnResp>StsRsnInf>Rsn>Prtry,omitempty" validate:"max=35"`
}
