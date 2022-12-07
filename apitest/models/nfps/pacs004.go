package models

import "encoding/xml"

type FPSPacs004 struct {
	XMLName  xml.Name `xml:"FpsMsg"`
	NumOfMsg string   `xml:"NbOfMsgs"`
	XMLHead  Head001  `xml:"FpsPylds>BizData>AppHdr"`
	Document Pacs004  `xml:"FpsPylds>BizData>Document"`
}

type Pacs004 struct {
	MsgId          string `xml:"PmtRtr>GrpHdr>MsgId" validate:"required,max=35"`
	CreateDatetime string `xml:"PmtRtr>GrpHdr>CreDtTm" validate:"required,max=25"`
	NumOfTrans     string `xml:"PmtRtr>GrpHdr>NbOfTxs" validate:"required,max=25"`
	StlmtMethod    string `xml:"PmtRtr>GrpHdr>SttlmInf>SttlmMtd" validate:"required,max=4"`
	ClrSysId       string `xml:"PmtRtr>GrpHdr>SttlmInf>ClrSys>Prtry" validate:"required,max=35"`
	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the returned transaction.
	// Usage: The instructing party is the party sending the return message and not the party that sent the original instruction that is being returned.
	RetId             string `xml:"PmtRtr>TxInf>RtrId" validate:"required,max=35"`
	OrigInstructionId string `xml:"PmtRtr>TxInf>OrgnlInstrId"`
	OrigEndToEndId    string `xml:"PmtRtr>TxInf>OrgnlEndToEndId" validate:"required,max=35"`
	OrigTransId       string `xml:"PmtRtr>TxInf>OrgnlTxId" validate:"required,max=35"`
	OrigClrSysRef     string `xml:"PmtRtr>TxInf>OrgnlClrSysRef" validate:"required,max=35"`
	// Amount of money moved between the instructing agent and the instructed agent in the returned transaction.
	RetInterbankStlmtAmt AmtField `xml:"PmtRtr>TxInf>RtrdIntrBkSttlmAmt" validate:"required,max=24"`
	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	//
	// Usage: the InterbankSettlementDate is the interbank settlement date of the return message, and not of the original instruction.
	InterbankStlmtDate     string   `xml:"PmtRtr>TxInf>IntrBkSttlmDt" validate:"required,max=10"`
	RetInstructedAmt       AmtField `xml:"PmtRtr>TxInf>RtrdInstdAmt"`
	ChargeBearerType       string   `xml:"PmtRtr>TxInf>ChrgBr"`
	ChargeAmt              AmtField `xml:"PmtRtr>TxInf>ChrgsInf>Amt"`
	ChargePartBic          string   `xml:"PmtRtr>TxInf>ChrgsInf>Agt>FinInstnId>BICFI"`
	ChargeMemberId         string   `xml:"PmtRtr>TxInf>ChrgsInf>Agt>FinInstnId>ClrSysMmbId>MmbId"`
	RetReasonCode          string   `xml:"PmtRtr>TxInf>RtrRsnInf>Rsn>Prtry" validate:"required"`
	RetReasonInfo          []string `xml:"PmtRtr>TxInf>RtrRsnInf>AddtlInf" validate:"required"`
	OrigInterbankStlmtAmt  AmtField `xml:"PmtRtr>TxInf>OrgnlTxRef>IntrBkSttlmAmt" validate:"required"`
	OrigInterbankStlmtDate string   `xml:"PmtRtr>TxInf>OrgnlTxRef>IntrBkSttlmDt" validate:"required"`
	OrigCategoryPurpose    string   `xml:"PmtRtr>TxInf>OrgnlTxRef>PmtTpInf>CtgyPurp>Prtry"`
	OrigMandateId          string   `xml:"PmtRtr>TxInf>OrgnlTxRef>MndtRltdInf>MndtId"`
	OrigUnstructured       string   `xml:"PmtRtr>TxInf>OrgnlTxRef>RmtInf>Ustrd"`
	OrigDrName             string   `xml:"PmtRtr>TxInf>OrgnlTxRef>Dbtr>Nm" validate:"required,max=140"`
	OrigDrOrgBic           string   `xml:"PmtRtr>TxInf>OrgnlTxRef>Dbtr>Id>OrgId>AnyBIC"`
	OrigDrId               string   `xml:"PmtRtr>TxInf>OrgnlTxRef>Dbtr>Id>OrgId>Othr>Id" validate:"max=35"`
	OrigDrSchemeName       string   `xml:"PmtRtr>TxInf>OrgnlTxRef>Dbtr>Id>OrgId>Othr>SchmeNm>Cd" validate:"max=4"`
	OrigDrMemberName       string   `xml:"PmtRtr>TxInf>OrgnlTxRef>Dbtr>Id>OrgId>Othr>Issr" validate:"max=35"`
	OrigDrCustId           string   `xml:"PmtRtr>TxInf>OrgnlTxRef>Dbtr>Id>PrvtId>Othr>Id" validate:"max=35"`
	OrigDrCustSchemeName   string   `xml:"PmtRtr>TxInf>OrgnlTxRef>Dbtr>Id>PrvtId>Othr>SchmeNm>Cd" validate:"max=4"`
	OrigDrCustMemberName   string   `xml:"PmtRtr>TxInf>OrgnlTxRef>Dbtr>Id>PrvtId>Othr>Issr" validate:"max=35"`
	OrigDrMobileNum        string   `xml:"PmtRtr>TxInf>OrgnlTxRef>Dbtr>CtctDtls>MobNb" validate:"max=35"`
	OrigDrEmailAddr        string   `xml:"PmtRtr>TxInf>OrgnlTxRef>Dbtr>CtctDtls>EmailAdr" validate:"max=2048"`
	OrigDrAcctNo           string   `xml:"PmtRtr>TxInf>OrgnlTxRef>DbtrAcct>Id>Othr>Id" validate:"required,max=35"`
	OrigDrAcctType         string   `xml:"PmtRtr>TxInf>OrgnlTxRef>DbtrAcct>Id>Othr>SchmeNm>Prtry" validate:"required,max=4"`
	OrigDrPartBic          string   `xml:"PmtRtr>TxInf>OrgnlTxRef>DbtrAgt>FinInstnId>BICFI" validate:"max=11"`
	OrigDrMemberId         string   `xml:"PmtRtr>TxInf>OrgnlTxRef>DbtrAgt>FinInstnId>ClrSysMmbId>MmbId" validate:"required,max=35"`
	OrigCrPartBic          string   `xml:"PmtRtr>TxInf>OrgnlTxRef>CdtrAgt>FinInstnId>BICFI" validate:"max=11"`
	OrigCrMemberId         string   `xml:"PmtRtr>TxInf>OrgnlTxRef>CdtrAgt>FinInstnId>ClrSysMmbId>MmbId" validate:"required,max=35"`
	OrigCrName             string   `xml:"PmtRtr>TxInf>OrgnlTxRef>Cdtr>Nm" validate:"required,max=140"`
	OrigCrOrgBic           string   `xml:"PmtRtr>TxInf>OrgnlTxRef>Cdtr>Id>OrgId>AnyBIC" validate:"max=11"`
	OrigCrId               string   `xml:"PmtRtr>TxInf>OrgnlTxRef>Cdtr>Id>OrgId>Othr>Id" validate:"max=35"`
	OrigCrSchemeName       string   `xml:"PmtRtr>TxInf>OrgnlTxRef>Cdtr>Id>OrgId>Othr>SchmeNm>Cd" validate:"max=4"`
	OrigCrMemberName       string   `xml:"PmtRtr>TxInf>OrgnlTxRef>Cdtr>Id>OrgId>Othr>Issr" validate:"max=35"`
	OrigCrCustId           string   `xml:"PmtRtr>TxInf>OrgnlTxRef>Cdtr>Id>PrvtId>Othr>Id" validate:"max=35"`
	OrigCrCustSchemeName   string   `xml:"PmtRtr>TxInf>OrgnlTxRef>Cdtr>Id>PrvtId>Othr>SchmeNm>Cd" validate:"max=4"`
	OrigCrCustMemberName   string   `xml:"PmtRtr>TxInf>OrgnlTxRef>Cdtr>Id>PrvtId>Othr>Issr" validate:"max=35"`
	OrigCrMobileNum        string   `xml:"PmtRtr>TxInf>OrgnlTxRef>Cdtr>CtctDtls>MobNb" validate:"max=35"`
	OrigCrEmailAddr        string   `xml:"PmtRtr>TxInf>OrgnlTxRef>Cdtr>CtctDtls>EmailAdr" validate:"max=2048"`
	OrigCrAcctNo           string   `xml:"PmtRtr>TxInf>OrgnlTxRef>CdtrAcct>Id>Othr>Id" validate:"required,max=35"`
	OrigCrAcctType         string   `xml:"PmtRtr>TxInf>OrgnlTxRef>CdtrAcct>Id>Othr>SchmeNm>Prtry" validate:"required,max=4"`
}
