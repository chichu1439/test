// Package models will define request and response message struct
// Version: v0.0.1
package models

import (
	"encoding/xml"
)

type FP010028I struct {
	XMLName      xml.Name     `xml:"FpsMsg"`
	NbOfMsgs     string       `xml:"NbOfMsgs,omitempty" validate:"required,max=15"`
	XmlHead      XmlHead      `xml:"FpsPylds>BizData>AppHdr"`
	FP010028IDoc FP010028IDoc `xml:"FpsPylds>BizData>Document"`
}

type FP010028IDoc struct {
	MsgId              string   `xml:"FIToFICstmrCdtTrf>GrpHdr>MsgId,omitempty" validate:"required,max=35"`
	CreateDatetime     string   `xml:"FIToFICstmrCdtTrf>GrpHdr>CreDtTm,omitempty" validate:"required,max=35"`
	NumOfTrans         string   `xml:"FIToFICstmrCdtTrf>GrpHdr>NbOfTxs,omitempty" validate:"required,max=15"`
	StlmtMethod        string   `xml:"FIToFICstmrCdtTrf>GrpHdr>SttlmInf>SttlmMtd,omitempty" validate:"required,max=4"`
	ClrSysId           string   `xml:"FIToFICstmrCdtTrf>GrpHdr>SttlmInf>ClrSys>Prtry,omitempty" validate:"max=35"`
	InstructionId      string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>PmtId>InstrId,omitempty" validate:"max=35"`
	EndToEndId         string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>PmtId>EndToEndId,omitempty" validate:"required,max=35"`
	TransId            string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>PmtId>TxId,omitempty" validate:"required,max=35"`
	ClrSysRef          string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>PmtId>ClrSysRef,omitempty" validate:"max=35"`
	AcctVerifyOption   string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>PmtTpInf>LclInstrm>Prtry,omitempty" validate:"max=35"`
	CategoryPurpose    string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>PmtTpInf>CtgyPurp>Prtry,omitempty" validate:"max=35"`
	InterbankStlmtAmt  AmtField `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>IntrBkSttlmAmt,omitempty" validate:"required,max=24"`
	InterbankStlmtDate string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>IntrBkSttlmDt,omitempty" validate:"max=10"`
	CreditDatetime     string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>SttlmTmIndctn>CdtDtTm,omitempty" validate:"max=19"`
	InstructedAmt      AmtField `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>InstdAmt,omitempty" validate:"max=24"`
	ChargeBearerType   string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>ChrgBr,omitempty" validate:"required,max=4"`
	ChargeAmt          AmtField `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>ChrgsInf>Amt,omitempty" validate:"max=24"`
	ChargePartBic      string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>ChrgsInf>Agt>FinInstnId>BICFI,omitempty" validate:"max=11"`
	ChargeMemberId     string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>ChrgsInf>Agt>FinInstnId>ClrSysMmbId>MmbId,omitempty" validate:"max=35"`
	DrName             string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Dbtr>Nm,omitempty" validate:"max=140"`
	DrOrgBic           string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Dbtr>Id>OrgId>AnyBIC,omitempty" validate:"max=11"`
	DrId               string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Dbtr>Id>OrgId>Othr>Id,omitempty" validate:"max=35"`
	DrSchemeName       string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Dbtr>Id>OrgId>Othr>SchmeNm>Cd,omitempty" validate:"max=4"`
	DrMemberName       string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Dbtr>Id>OrgId>Othr>Issr,omitempty" validate:"max=35"`
	DrCustId           string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Dbtr>Id>PrvtId>Othr>Id,omitempty" validate:"max=35"`
	DrCustSchemeName   string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Dbtr>Id>PrvtId>Othr>SchmeNm>Cd,omitempty" validate:"max=4"`
	DrCustMemberName   string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Dbtr>Id>PrvtId>Othr>Issr,omitempty" validate:"max=35"`
	DrMobileNum        string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Dbtr>CtctDtls>MobNb,omitempty" validate:"max=35"`
	DrEmailAddr        string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Dbtr>CtctDtls>EmailAdr,omitempty" validate:"max=2048"`
	DrAcctNo           string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>DbtrAcct>Id>Othr>Id,omitempty" validate:"required,max=35"`
	DrAcctType         string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>DbtrAcct>Id>Othr>SchmeNm>Prtry,omitempty" validate:"max=4"`
	DrPartBic          string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>DbtrAgt>FinInstnId>BICFI,omitempty" validate:"max=11"`
	DrMemberId         string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>DbtrAgt>FinInstnId>ClrSysMmbId>MmbId,omitempty" validate:"required,max=35"`
	CrPartBic          string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>CdtrAgt>FinInstnId>BICFI,omitempty" validate:"max=11"`
	CrMemberId         string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>CdtrAgt>FinInstnId>ClrSysMmbId>MmbId,omitempty" validate:"required,max=35"`
	CrName             string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Cdtr>Nm,omitempty" validate:"max=140"`
	CrOrgBic           string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Cdtr>Id>OrgId>AnyBIC,omitempty" validate:"max=11"`
	CrId               string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Cdtr>Id>OrgId>Othr>Id,omitempty" validate:"max=35"`
	CrSchemeName       string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Cdtr>Id>OrgId>Othr>SchmeNm>Cd,omitempty" validate:"max=4"`
	CrMemberName       string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Cdtr>Id>OrgId>Othr>Issr,omitempty" validate:"max=35"`
	CrCustId           string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Cdtr>Id>PrvtId>Othr>Id,omitempty" validate:"max=35"`
	CrCustSchemeName   string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Cdtr>Id>PrvtId>Othr>SchmeNm>Cd,omitempty" validate:"max=4"`
	CrCustMemberName   string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Cdtr>Id>PrvtId>Othr>Issr,omitempty" validate:"max=35"`
	CrMobileNum        string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Cdtr>CtctDtls>MobNb,omitempty" validate:"max=35"`
	CrEmailAddr        string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Cdtr>CtctDtls>EmailAdr,omitempty" validate:"max=2048"`
	CrAcctNo           string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>CdtrAcct>Id>Othr>Id,omitempty" validate:"required,max=35"`
	CrAcctType         string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>CdtrAcct>Id>Othr>SchmeNm>Prtry,omitempty" validate:"max=4"`
	PaymentPurposeCode string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Purp>Cd,omitempty" validate:"max=4"`
	PaymentPurposeType string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>Purp>Prtry,omitempty" validate:"max=35"`
	Unstructured       string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>RmtInf>Ustrd,omitempty" validate:"max=140"`
}

type FP010028O struct {
}
