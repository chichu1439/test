package models

import "encoding/xml"

type FPSCamt014 struct {
	XMLName  xml.Name       `xml:"Envelope"`
	XmlHead  FPSCamt014Head `xml:"AppHdr"`
	Document Camt014        `xml:"Document>RtrMmb"`
}
type FPSCamt014Head struct {
	FromMemberId   string `xml:"Fr>FIId>FinInstnId>Othr>Id,omitempty" validate:"required,max=35"`
	ToMemberId     string `xml:"To>FIId>FinInstnId>Othr>Id,omitempty" validate:"required,max=35"`
	BizMsgId       string `xml:"BizMsgIdr,omitempty" validate:"required,max=35"`
	MsgDefId       string `xml:"MsgDefIdr,omitempty" validate:"required,max=35"`
	CreateDatetime string `xml:"CreDt,omitempty" validate:"required,max=35"`
}
type Camt014 struct {
	MsgId    string     `xml:"MsgHdr>MsgId" json:"MsgId,omitempty"`
	CreDtTm  string     `xml:"MsgHdr>CreDtTm" json:"CreDtTm,omitempty"`
	RptOrErr []RptOrErr `xml:"RptOrErr>Rpt" json:"RptOrErr"`
}
type RptOrErr struct {
	MmbId    string        `xml:"MmbId>ClrSysMmbId>MmbId" json:"MmbId,omitempty"`
	Name     string        `xml:"MmbOrErr>Mmb>Nm" json:"Name,omitempty"`
	Type     string        `xml:"MmbOrErr>Mmb>Tp>Cd" json:"Type,omitempty"`
	Status   string        `xml:"MmbOrErr>Mmb>Sts>Cd" json:"Sts,omitempty"`
	Acctount []Camt014Acct `xml:"MmbOrErr>Mmb>Acct" json:"Acct,omitempty"`
}

type Camt014Acct struct {
	Id       string `xml:"Id>Othr>Id" json:"StNm,omitempty"`
	Currency string `xml:"Ccy" json:"EnNm,omitempty"`
}
