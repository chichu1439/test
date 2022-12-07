package models

import "encoding/xml"

type FPSAdmi004 struct {
	XMLName  xml.Name `xml:"FpsMsg"`
	NumOfMsg string   `xml:"NbOfMsgs,omitempty" validate:"required,max=15"`
	XMLHead  Head001  `xml:"FpsPylds>BizData>AppHdr"`
	Document Admi004  `xml:"FpsPylds>BizData>Document"`
}

// Admi004
type Admi004 struct {
	EvtCd    string `xml:"SysEvtNtfctn>EvtInf>EvtCd"`
	EvtParam string `xml:"SysEvtNtfctn>EvtInf>EvtParam"`
	EvtDesc  string `xml:"SysEvtNtfctn>EvtInf>EvtDesc"`
	EvtTm    string `xml:"SysEvtNtfctn>EvtInf>EvtTm"`
}
