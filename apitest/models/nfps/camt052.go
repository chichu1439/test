package models

import "encoding/xml"

type FPSCamt052 struct {
	XMLName  xml.Name `xml:"fps:FpsMsg"`
	NbOfMsgs string   `xml:"fps:NbOfMsgs"`
	XmlHead  XmlHead  `xml:"fps:FpsPylds>fps:BizData>ah:AppHdr"`
	Document Camt052  `xml:"fps:FpsPylds>fps:BizData>doc:Document>doc:BkToCstmrAcctRpt"`
}

type Camt052 struct {
	MsgId      string `xml:"doc:GrpHdr>doc:MsgId"`
	CreDtTm    string `xml:"doc:GrpHdr>doc:CreDtTm"`
	RptMsgId   string `xml:"doc:Rpt>doc:Id"`
	RptCreDtTm string `xml:"doc:Rpt>doc:CreDtTm"`
	// Member identification of the account owner
	AcctId string `xml:"doc:Acct>doc:Id>doc:Othr>doc:Id"`
	// Currency of account balance
	Currency string `xml:"doc:Acct>doc:Ccy"`
	// Balance->Type->Code Or Proprietary->Code
	//The valid balance type code is
	//1. ITAV=Interim Available
	BalType string   `xml:"doc:Bal>doc:Tp>doc:CdOrPrtry>doc:Cd"`
	Amount  AmtField `xml:"doc:Bal>doc:Amt"`
	// Indicates whether the account balance is a credit or a debit entry.
	// A zero balance is considered to be a credit balance.
	CdtrDbtrInde string `xml:"doc:Bal>doc:CdtDbtInd"`
	// System time for current available balance
	BalDtTm string `xml:"doc:Bal>doc:Dt>doc:DtTm"`
}
