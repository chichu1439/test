package models

import "encoding/xml"

type FPSCamt060 struct {
	XMLName  xml.Name `xml:"FpsMsg"`
	NbOfMsgs string   `xml:"NbOfMsgs"`
	XmlHead  XmlHead  `xml:"FpsPylds>BizData>AppHdr"`
	Document Camt060  `xml:"FpsPylds>BizData>Document>AcctRptgReq"`
}

type Camt060 struct {
	MsgId       string `xml:"GrpHdr>MsgId"`
	CreDtTm     string `xml:"GrpHdr>CreDtTm"`
	ReqdMsgNmId string `xml:"RptgReq>ReqdMsgNmId"`
	AcctId      string `xml:"RptgReq>Acct>Id>Othr>Id"`
	Currency    string `xml:"RptgReq>Acct>Ccy"` // [A-Z]{3,3}
	// Member identification of the sender participant or the IP of the sender participant
	MmbId string `xml:"RptgReq>AcctOwnr>Agt>FinInstnId>ClrSysMmbId>MmbId"`
}
