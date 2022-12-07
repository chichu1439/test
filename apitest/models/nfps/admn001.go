package models

import "encoding/xml"

type FPSSignOn001 struct {
	XMLName  xml.Name `xml:"FpsMsg"`
	NumOfMsg string   `xml:"NbOfMsgs"`
	XMLHead  Head001  `xml:"FpsPylds>BizData>AppHdr"`
	Document Admn001  `xml:"FpsPylds>BizData>Document"`
}

type Admn001 struct {
	MsgId       string `xml:"AdmnSignOnReq>GrpHdr>MsgId,omitempty" validate:"required,max=35"`
	CreDtTm     string `xml:"AdmnSignOnReq>GrpHdr>CreDtTm,omitempty" validate:"required"`
	InstrId     string `xml:"AdmnSignOnReq>SignOnReq>InstrId,omitempty" validate:"required"  validate:"required,max=35"`
	Instructing string `xml:"AdmnSignOnReq>SignOnReq>InstgAgt>FinInstnId>ClrSysMmbId>MmbId,omitempty" validate:"required"`
	Instructed  string `xml:"AdmnSignOnReq>SignOnReq>InstdAgt>FinInstnId>ClrSysMmbId>MmbId,omitempty"`
}
