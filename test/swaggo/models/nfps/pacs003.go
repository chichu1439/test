package models

import "encoding/xml"

type FPSPacs003 struct {
	XMLName  xml.Name `xml:"FpsMsg"`
	NumOfMsg string   `xml:"NbOfMsgs"`
	XMLHead  Head001  `xml:"FpsPylds>BizData>AppHdr"`
	Document Pacs003  `xml:"FpsPylds>BizData>Document"`
}

type Pacs003 struct {
	MsgId              string   `xml:"FIToFICstmrDrctDbt>GrpHdr>MsgId,omitempty" validate:"required,max=35"`
	CreateDatetime     string   `xml:"FIToFICstmrDrctDbt>GrpHdr>CreDtTm,omitempty" validate:"required,max=25"`
	NumOfTrans         string   `xml:"FIToFICstmrDrctDbt>GrpHdr>NbOfTxs,omitempty" validate:"required,max=15"`
	StlmtMethod        string   `xml:"FIToFICstmrDrctDbt>GrpHdr>SttlmInf>SttlmMtd,omitempty" validate:"required,max=4"`
	ClrSysId           string   `xml:"FIToFICstmrDrctDbt>GrpHdr>SttlmInf>ClrSys>Prtry,omitempty" validate:"required,max=35"`
	InstructionId      string   `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>PmtId>InstrId,omitempty" validate:"max=35"`
	EndToEndId         string   `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>PmtId>EndToEndId,omitempty" validate:"required,max=35"`
	TransId            string   `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>PmtId>TxId,omitempty" validate:"required,max=35"`
	ClrSysRef          string   `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>PmtId>ClrSysRef,omitempty" validate:"max=35"`
	CategoryPurpose    string   `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>PmtTpInf>CtgyPurp>Prtry,omitempty" validate:"max=35"`
	InterbankStlmtAmt  AmtField `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>IntrBkSttlmAmt,omitempty" validate:"required,max=24"`
	InterbankStlmtDate string   `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>IntrBkSttlmDt,omitempty" validate:"required,max=10"`
	InstructedAmt      AmtField `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>InstdAmt,omitempty" validate:"max=24"`
	ChargeBearerType   string   `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>ChrgBr,omitempty" validate:"required,max=4"`
	ChargeAmt          AmtField `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>ChrgsInf>Amt,omitempty" validate:"max=24"`
	ChargePartBic      string   `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>ChrgsInf>Agt>FinInstnId>BICFI,omitempty" validate:"max=11"`
	ChargeMemberId     string   `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>ChrgsInf>Agt>FinInstnId>ClrSysMmbId>MmbId,omitempty" validate:"max=35"`
	// Unique identification, as assigned by the creditor, to unambiguously identify the mandate.
	MandateId          string `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>DrctDbtTx>MndtRltdInf>MndtId,omitempty" validate:"max=35"`
	CrName             string `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>Cdtr>Nm,omitempty" validate:"required,max=140"`
	CrOrgBic           string `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>Cdtr>Id>OrgId>AnyBIC,omitempty" validate:"max=11"`
	CrId               string `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>Cdtr>Id>OrgId>Othr>Id,omitempty" validate:"max=35"`
	CrSchemeName       string `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>Cdtr>Id>OrgId>Othr>SchmeNm>Cd,omitempty" validate:"max=4"`
	CrMemberName       string `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>Cdtr>Id>OrgId>Othr>Issr,omitempty" validate:"max=35"`
	CrCustId           string `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>Cdtr>Id>PrvtId>Othr>Id,omitempty" validate:"max=35"`
	CrCustSchemeName   string `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>Cdtr>Id>PrvtId>Othr>SchmeNm>Cd,omitempty" validate:"max=4"`
	CrCustMemberName   string `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>Cdtr>Id>PrvtId>Othr>Issr,omitempty" validate:"max=35"`
	CrMobileNum        string `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>Cdtr>CtctDtls>MobNb,omitempty" validate:"max=35"`
	CrEmailAddr        string `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>Cdtr>CtctDtls>EmailAdr,omitempty" validate:"max=2048"`
	CrAcctNo           string `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>CdtrAcct>Id>Othr>Id,omitempty" validate:"required,max=35"`
	CrAcctType         string `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>CdtrAcct>Id>Othr>SchmeNm>Prtry,omitempty" validate:"required,max=4"`
	CrPartBic          string `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>CdtrAgt>FinInstnId>BICFI,omitempty" validate:"max=11"`
	CrMemberId         string `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>CdtrAgt>FinInstnId>ClrSysMmbId>MmbId,omitempty" validate:"required,max=35"`
	DrName             string `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>Dbtr>Nm,omitempty" validate:"required,max=140"`
	DrOrgBic           string `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>Dbtr>Id>OrgId>AnyBIC,omitempty" validate:"max=11"`
	DrId               string `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>Dbtr>Id>OrgId>Othr>Id,omitempty" validate:"max=35"`
	DrSchemeName       string `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>Dbtr>Id>OrgId>Othr>SchmeNm>Cd,omitempty" validate:"max=4"`
	DrMemberName       string `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>Dbtr>Id>OrgId>Othr>Issr,omitempty" validate:"max=35"`
	DrCustId           string `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>Dbtr>Id>PrvtId>Othr>Id,omitempty" validate:"max=35"`
	DrCustSchemeName   string `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>Dbtr>Id>PrvtId>Othr>SchmeNm>Cd,omitempty" validate:"max=4"`
	DrCustMemberName   string `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>Dbtr>Id>PrvtId>Othr>Issr,omitempty" validate:"max=35"`
	DrMobileNum        string `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>Dbtr>CtctDtls>MobNb,omitempty" validate:"max=35"`
	DrEmailAddr        string `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>Dbtr>CtctDtls>EmailAdr,omitempty" validate:"max=2048"`
	DrAcctNo           string `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>DbtrAcct>Id>Othr>Id,omitempty" validate:"required,max=35"`
	DrAcctType         string `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>DbtrAcct>Id>Othr>SchmeNm>Prtry,omitempty" validate:"required,max=4"`
	DrPartBic          string `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>DbtrAgt>FinInstnId>BICFI,omitempty" validate:"max=11"`
	DrMemberId         string `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>DbtrAgt>FinInstnId>ClrSysMmbId>MmbId,omitempty" validate:"required,max=35"`
	PaymentPurposeCode string `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>Purp>Cd,omitempty" validate:"max=4"`
	PaymentPurposeType string `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>Purp>Prtry,omitempty" validate:"max=35"`
	// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in an unstructured form.
	Unstructured string `xml:"FIToFICstmrDrctDbt>DrctDbtTxInf>RmtInf>Ustrd,omitempty" validate:"max=140"`
}
