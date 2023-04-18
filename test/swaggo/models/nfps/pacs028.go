package models

import "encoding/xml"

type FPSPacs028 struct {
	XMLName  xml.Name `xml:"FpsMsg"`
	NbOfMsgs string   `xml:"NbOfMsgs,omitempty"` //  validate:"required,max=15"
	XmlHead  XmlHead  `xml:"FpsPylds>BizData>AppHdr"`
	Document Pacs028  `xml:"FpsPylds>BizData>Document>FIToFIPmtStsReq"`
}

type Pacs028 struct {
	MsgId          string `xml:"GrpHdr>MsgId,omitempty"`
	CreateDatetime string `xml:"GrpHdr>CreDtTm,omitempty"`
	// StatusRequestId string `xml:"TxInf>StsReqId,omitempty"`
	OriMessId string `xml:"TxInf>OrgnlGrpInf>OrgnlMsgId,omitempty"`
	// Message Name Identifier
	OriMessNameId string `xml:"TxInf>OrgnlGrpInf>OrgnlMsgNmId,omitempty"`
	// Original Transaction Identification
	OriTranId string `xml:"TxInf>OrgnlTxId,omitempty"`
	// NFPS Reference Number
	ClrSysRef string `xml:"TxInf>ClrSysRef,omitempty"`
	// Debtor
	// Business Identifier Code (BIC) of payer participant
	// DbtrAgtBICFI string `xml:"TxInf>OrgnlTxRef>DbtrAgt>FinInstnId>BICFI,omitempty"`
	// Clearing System Member Id
	DbtrAgtMmbId string `xml:"TxInf>OrgnlTxRef>DbtrAgt>FinInstnId>ClrSysMmbId>MmbId"`
	// Creditor
	// Business Identifier Code (BIC) of payer participant
	// CdtrAgtBICFI string `xml:"TxInf>OrgnlTxRef>CdtrAgt>FinInstnId>BICFI"`
	// Clearing System Member Id
	CdtrAgtMmbId string `xml:"TxInf>OrgnlTxRef>CdtrAgt>FinInstnId>ClrSysMmbId>MmbId"`
}
