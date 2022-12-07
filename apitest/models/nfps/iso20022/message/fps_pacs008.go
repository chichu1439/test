package message

import (
	"encoding/xml"

	"apitest/models/iso20022/head_v01"
	"apitest/models/iso20022/pacs_v08"
	"apitest/util"
)

type Pacs008BizData struct {
	XmlHead *head_v01.BusinessApplicationHeaderV01 `xml:"AppHdr"`
	XmlDoc  FpsPacs008Doc                          `xml:"Document"`
}

func (p Pacs008BizData) Validate() error {
	return util.Validate(&p)
}

type FpsPacs008Doc struct {
	XMLName xml.Name   `xml:"Document"`
	Attrs   []xml.Attr `xml:",any,attr,omitempty" json:",omitempty"`
	XmlDoc  *pacs_v08.FIToFICustomerCreditTransferV08
}

func (f FpsPacs008Doc) Validate() error {
	return util.Validate(&f)
}

type FpsPacs008 struct {
	XMLName xml.Name       `xml:"FpsMsg"`
	Attrs   []xml.Attr     `xml:",any,attr,omitempty" json:",omitempty"`
	BizData Pacs008BizData `xml:"BizData"`
}

func (f FpsPacs008) Validate() error {
	return util.Validate(&f)
}

type FpsPacs008Compact struct {
	XMLName  xml.Name   `xml:"FpsMsg"`
	Attrs    []xml.Attr `xml:",any,attr,omitempty" json:",omitempty"`
	XMLHead  Head001    `xml:"BizData>AppHdr"`
	Document Pacs008    `xml:"BizData>Document"`
}

type Head001 struct {
	FromMemberId   string `xml:"Fr>FIId>FinInstnId>ClrSysMmbId>MmbId,omitempty" validate:"required,max=35"`
	FromBranchId   string `xml:"Fr>BrnchId>Id,omitempty" validate:"max=35"` //itmx add
	ToMemberId     string `xml:"To>FIId>FinInstnId>ClrSysMmbId>MmbId,omitempty" validate:"required,max=35"`
	ToBranchId     string `xml:"To>BrnchId>Id,omitempty" validate:"max=35"` //itmx add
	BizMsgId       string `xml:"BizMsgIdr,omitempty" validate:"required,max=35"`
	MsgDefId       string `xml:"MsgDefIdr,omitempty" validate:"required,max=35"`
	BizSvc         string `xml:"BizSvc,omitempty" validate:"required,max=35"`
	CreateDatetime string `xml:"CreDt,omitempty" validate:"required,max=35"`
	CopyDup        string `xml:"CpyDplct,omitempty" validate:"max=35"` //itmx add
}

type AmtField struct {
	Amt string `xml:",chardata" validate:"omitempty,numeric"`
	Ccy string `xml:"Ccy,attr"`
}

type Pacs008 struct {
	MsgId               string   `xml:"FIToFICstmrCdtTrf>GrpHdr>MsgId,omitempty" validate:"required,max=35"`
	CreateDatetime      string   `xml:"FIToFICstmrCdtTrf>GrpHdr>CreDtTm,omitempty" validate:"required,max=35"`
	NumOfTrans          string   `xml:"FIToFICstmrCdtTrf>GrpHdr>NbOfTxs,omitempty" validate:"required,max=15"`
	InstructionId       string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>PmtId>InstrId,omitempty" validate:"max=35"`
	EndToEndId          string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>PmtId>EndToEndId,omitempty" validate:"required,max=35"`
	TransId             string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>PmtId>TxId,omitempty" validate:"required,max=35"`
	AcctVerifyOption    string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>PmtTpInf>LclInstrm>Prtry,omitempty" validate:"max=35"`
	InterbankStlmtAmt   AmtField `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>IntrBkSttlmAmt,omitempty" validate:"required,max=24"`
	CreditDatetime      string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>SttlmTmIndctn>CdtDtTm,omitempty" validate:"max=19"` // removed
	InstructedAmt       AmtField `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>InstdAmt,omitempty" validate:"max=24"`              // removed
	ChargeBearerType    string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>ChrgBr,omitempty" validate:"max=4"`
	ChargeAmt           AmtField `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>ChrgsInf>Amt,omitempty" validate:"max=24"`                              // removed
	ChargePartBic       string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>ChrgsInf>Agt>FinInstnId>BICFI,omitempty" validate:"max=11"`             // removed
	ChargeMemberId      string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>ChrgsInf>Agt>FinInstnId>ClrSysMmbId>MmbId,omitempty" validate:"max=35"` // removed
	InstructingMemberId string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>InstgAgt>Agt>FinInstnId>ClrSysMmbId>MmbId,omitempty" validate:"max=35"`
	InstructedMemberId  string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>InstdAgt>Agt>FinInstnId>ClrSysMmbId>MmbId,omitempty" validate:"max=35"`
	DrName              string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Dbtr>Nm,omitempty" validate:"max=140"`
	DrOrgBic            string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Dbtr>Id>OrgId>AnyBIC,omitempty" validate:"max=11"` // removed
	DrId                string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Dbtr>Id>OrgId>Othr>Id,omitempty" validate:"max=35"`
	DrSchemeName        string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Dbtr>Id>OrgId>Othr>SchmeNm>Cd,omitempty" validate:"max=4"`  // removed
	DrMemberName        string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Dbtr>Id>OrgId>Othr>Issr,omitempty" validate:"max=35"`       // removed
	DrCustId            string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Dbtr>Id>PrvtId>Othr>Id,omitempty" validate:"max=35"`        // removed
	DrCustSchemeName    string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Dbtr>Id>PrvtId>Othr>SchmeNm>Cd,omitempty" validate:"max=4"` // removed
	DrCustMemberName    string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Dbtr>Id>PrvtId>Othr>Issr,omitempty" validate:"max=35"`      // removed
	DrMobileNum         string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Dbtr>CtctDtls>MobNb,omitempty" validate:"max=35"`           // removed
	DrEmailAddr         string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Dbtr>CtctDtls>EmailAdr,omitempty" validate:"max=2048"`      // removed
	DrAcctNo            string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>DbtrAcct>Id>Othr>Id,omitempty" validate:"required,max=35"`
	DrAcctType          string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>DbtrAcct>Id>Othr>SchmeNm>Prtry,omitempty" validate:"max=4"` // removed
	DrPartBic           string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>DbtrAgt>FinInstnId>BICFI,omitempty" validate:"max=11"`      // removed
	DrMemberId          string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>DbtrAgt>FinInstnId>ClrSysMmbId>MmbId,omitempty" validate:"required,max=35"`
	CrPartBic           string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>CdtrAgt>FinInstnId>BICFI,omitempty" validate:"max=11"` // removed
	CrMemberId          string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>CdtrAgt>FinInstnId>ClrSysMmbId>MmbId,omitempty" validate:"required,max=35"`
	CrName              string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Cdtr>Nm,omitempty" validate:"max=140"`
	CrOrgBic            string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Cdtr>Id>OrgId>AnyBIC,omitempty" validate:"max=11"`          // removed
	CrId                string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Cdtr>Id>OrgId>Othr>Id,omitempty" validate:"max=35"`         // removed
	CrSchemeName        string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Cdtr>Id>OrgId>Othr>SchmeNm>Cd,omitempty" validate:"max=4"`  // removed
	CrMemberName        string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Cdtr>Id>OrgId>Othr>Issr,omitempty" validate:"max=35"`       // removed
	CrCustId            string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Cdtr>Id>PrvtId>Othr>Id,omitempty" validate:"max=35"`        // removed
	CrCustSchemeName    string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Cdtr>Id>PrvtId>Othr>SchmeNm>Cd,omitempty" validate:"max=4"` // removed
	CrCustMemberName    string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Cdtr>Id>PrvtId>Othr>Issr,omitempty" validate:"max=35"`      // removed
	CrMobileNum         string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Cdtr>CtctDtls>MobNb,omitempty" validate:"max=35"`           // removed
	CrEmailAddr         string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Cdtr>CtctDtls>EmailAdr,omitempty" validate:"max=2048"`      // removed
	CrAcctNo            string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>CdtrAcct>Id>Othr>Id,omitempty" validate:"required,max=35"`
	CrAcctType          string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>CdtrAcct>Id>Othr>SchmeNm>Prtry,omitempty" validate:"max=4"` // removed
	PaymentPurposeCode  string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Purp>Cd,omitempty" validate:"max=4"`
	PaymentPurposeType  string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Purp>Prtry,omitempty" validate:"max=35"`
}
