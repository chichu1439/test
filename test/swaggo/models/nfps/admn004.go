package models

import "encoding/xml"

type FPSSignOff004 struct {
	XMLName  xml.Name `xml:"FpsMsg"`
	NumOfMsg string   `xml:"NbOfMsgs"`
	XMLHead  Head001  `xml:"FpsPylds>BizData>AppHdr"`
	Document Admn004  `xml:"FpsPylds>BizData>Document"`
}

type Admn004 struct {
	MsgId       string `xml:"AdmnSignOffResp>GrpHdr>MsgId,omitempty" validate:"required,max=35"`
	CreDtTm     string `xml:"AdmnSignOffResp>GrpHdr>CreDtTm,omitempty" validate:"required"`
	InstrId     string `xml:"AdmnSignOffResp>SignOffResp>OrgnlInstrId,omitempty" validate:"required"  validate:"required,max=35"`
	Instructing string `xml:"AdmnSignOffResp>SignOffResp>InstgAgt>FinInstnId>ClrSysMmbId>MmbId,omitempty" validate:"required"`
	Instructed  string `xml:"AdmnSignOffResp>SignOffResp>InstdAgt>FinInstnId>ClrSysMmbId>MmbId,omitempty"`
	Sts         string `xml:"AdmnSignOffResp>SignOffResp>Sts,omitempty"`
	Cd          string `xml:"AdmnSignOffResp>SignOffResp>StsRsnInf>Rsn>Cd,omitempty"`
	Prtry       string `xml:"AdmnSignOffResp>SignOffResp>StsRsnInf>Rsn>Prtry,omitempty" validate:"max=35"`
}
