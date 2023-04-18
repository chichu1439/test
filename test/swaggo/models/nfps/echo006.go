package models

import "encoding/xml"

type FPSEcho006 struct {
	XMLName  xml.Name `xml:"fps:FpsMsg"`
	NumOfMsg string   `xml:"fps:NbOfMsgs"`
	XMLHead  Head001  `xml:"fps:FpsPylds>fps:BizData>ah:AppHdr"`
	Document Echo006  `xml:"FpsPylds>BizData>Document"`
}

type Echo006 struct {
	MsgId        string `xml:"AdmnEchoResp>GrpHdr>MsgId,omitempty" validate:"required"`
	CreDtTm      string `xml:"AdmnEchoResp>GrpHdr>CreDtTm,omitempty" validate:"required"`
	Instructing  string `xml:"AdmnEchoResp>EchoResponse>InstgAgt>FinInstnId>ClrSysMmbId>MmbId,omitempty" validate:"required"`
	Instruted    string `xml:"AdmnEchoResp>EchoResponse>InstdAgt>FinInstnId>ClrSysMmbId>MmbId,omitempty" validate:"required"`
	OrgnlInstrId string `xml:"AdmnEchoResp>OrgnlInstrId,omitempty" validate:"required"`
	FnctnCd      string `xml:"AdmnEchoResp>FnctnCd,omitempty" validate:"required"`
	TxSts        string `xml:"AdmnEchoResp>TxSts,omitempty" validate:"required"`
}
