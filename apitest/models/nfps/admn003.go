package models

import "encoding/xml"

type FPSSignOff003 struct {
	XMLName  xml.Name `xml:"FpsMsg"`
	NumOfMsg string   `xml:"NbOfMsgs"`
	XMLHead  Head001  `xml:"FpsPylds>BizData>AppHdr"`
	Document Admn003  `xml:"FpsPylds>BizData>Document"`
}

type Admn003 struct {
	MsgId       string `xml:"AdmnSignOffReq>GrpHdr>MsgId,omitempty" validate:"required,max=35"`
	CreDtTm     string `xml:"AdmnSignOffReq>GrpHdr>CreDtTm,omitempty" validate:"required"`
	InstrId     string `xml:"AdmnSignOffReq>SignOffReq>InstrId,omitempty" validate:"required"  validate:"required,max=35"`
	Instructing string `xml:"AdmnSignOffReq>SignOffReq>InstgAgt>FinInstnId>ClrSysMmbId>MmbId,omitempty" validate:"required"`
	Instructed  string `xml:"AdmnSignOffReq>SignOffReq>InstdAgt>FinInstnId>ClrSysMmbId>MmbId,omitempty"`
}

