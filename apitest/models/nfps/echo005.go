package models

import "encoding/xml"

type FPSADMNo005 struct {
	XMLName  xml.Name `xml:"FpsMsg"`
	NumOfMsg string   `xml:"NbOfMsgs"`
	XMLHead  Head001  `xml:"FpsPylds>BizData>AppHdr"`
	Document Echo005  `xml:"FpsPylds>BizData>Document"`
}

type Echo005 struct {
	MsgId string `xml:"AdmnEchoReq>GrpHdr>MsgId,omitempty" validate:"required"`
	CreDtTm string `xml:"AdmnEchoReq>GrpHdr>CreDtTm,omitempty" validate:"required"`
	FnctnCd string `xml:"AdmnEchoReq>EchoTxInf>FnctnCd,omitempty" validate:"required"`
	InstrId string `xml:"AdmnEchoReq>EchoTxInf>InstrId,omitempty" validate:"required"`
	Instructing string `xml:"AdmnEchoReq>EchoTxInf>InstgAgt>FinInstnId>ClrSysMmbId>MmbId,omitempty" validate:"required"`
	Instructed string `xml:"AdmnEchoReq>EchoTxInf>InstdAgt>FinInstnId>ClrSysMmbId>MmbId,omitempty" validate:"required"`
}