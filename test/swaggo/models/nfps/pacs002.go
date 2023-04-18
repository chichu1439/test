package models

import "encoding/xml"

type FPSPacs002 struct {
	XMLName  xml.Name `xml:"FpsMsg"`
	NumOfMsg string   `xml:"NbOfMsgs"`
	XMLHead  Head001  `xml:"FpsPylds>BizData>AppHdr"`
	Document Pacs002  `xml:"FpsPylds>BizData>Document"`
}

type Pacs002 struct {
	// Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MsgId string `xml:"FIToFIPmtStsRpt>GrpHdr>MsgId,omitempty"`
	// Date and time at which the message was created.
	CreateDatetime string `xml:"FIToFIPmtStsRpt>GrpHdr>CreDtTm,omitempty"`
	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OriginalMsgId string `xml:"FIToFIPmtStsRpt>OrgnlGrpInfAndSts>OrgnlMsgId,omitempty"`
	// Specifies the original message name identifier to which the message refers.
	OriginalMsgNameId      string `xml:"FIToFIPmtStsRpt>OrgnlGrpInfAndSts>OrgnlMsgNmId,omitempty"`
	OriginalCreateDatetime string `xml:"FIToFIPmtStsRpt>OrgnlGrpInfAndSts>OrgnlCreDtTm,omitempty"` //itmx add
	OriginalNumOfTrans     string `xml:"FIToFIPmtStsRpt>OrgnlGrpInfAndSts>OrgnlNbOfTxs,omitempty"` //itmx add
	// Specifies the status of a group of transactions.
	OriginalGroupStatus string `xml:"FIToFIPmtStsRpt>OrgnlGrpInfAndSts>GrpSts,omitempty"` // removed
	// Provides detailed information on the status reason.
	// Specifies the reason for the status report.
	OriginalRejectReason string `xml:"FIToFIPmtStsRpt>OrgnlGrpInfAndSts>StsRsnInf>Rsn>Prtry,omitempty"` // removed
	// Additional information can be used for several purposes such as the reporting of repaired information.
	OriginalRejectReasonInfo string `xml:"FIToFIPmtStsRpt>OrgnlGrpInfAndSts>StsRsnInf>AddtlInf,omitempty"` // removed
	OriginalInstructionId    string `xml:"FIToFIPmtStsRpt>TxInfAndSts>OrgnlInstrId,omitempty"`             //itmx add
	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndId string `xml:"FIToFIPmtStsRpt>TxInfAndSts>OrgnlEndToEndId,omitempty"`
	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OriginalTransId string `xml:"FIToFIPmtStsRpt>TxInfAndSts>OrgnlTxId,omitempty"`
	// Specifies the status of a transaction, in a coded form.
	TransStatus string `xml:"FIToFIPmtStsRpt>TxInfAndSts>TxSts,omitempty"`
	// Provides detailed information on the status reason.Non-ISO code
	TransRejectReason string `xml:"FIToFIPmtStsRpt>TxInfAndSts>StsRsnInf>Rsn>Prtry,omitempty"`
	// Provides detailed information on the status reason. base ISO
	TransRejectReason2 string `xml:"FIToFIPmtStsRpt>TxInfAndSts>StsRsnInf>Rsn>Cd,omitempty"`
	// Further details on the status reason.
	// Additional information can be used for several purposes such as the reporting of repaired information.
	TransRejectReasonInfo string `xml:"FIToFIPmtStsRpt>TxInfAndSts>StsRsnInf>AddtlInf,omitempty"` // removed
	// Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent. This means that the account servicing agent has received the payment order and has applied checks such as authorisation, availability of funds.
	AcceptDatetime string `xml:"FIToFIPmtStsRpt>TxInfAndSts>AccptncDtTm,omitempty"`
	// Unique reference, as assigned by a clearing system, to unambiguously identify the instruction.
	ClrSysRef        string `xml:"FIToFIPmtStsRpt>TxInfAndSts>ClrSysRef,omitempty"`
	InstructingAgent string `xml:"FIToFIPmtStsRpt>TxInfAndSts>InstgAgt>FinInstnId>ClrSysMmbId>MmbId,omitempty"` //itmx add
	InstructedAgent  string `xml:"FIToFIPmtStsRpt>TxInfAndSts>InstdAgt>FinInstnId>ClrSysMmbId>MmbId,omitempty"` //itmx add

	// Key elements used to identify the original transaction that is being referred to.
	// Amount of money moved between the instructing agent and the instructed agent.
	InterbankStlmtAmt AmtField `xml:"FIToFIPmtStsRpt>TxInfAndSts>OrgnlTxRef>IntrBkSttlmAmt,omitempty"`
	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankStlmtDate string   `xml:"FIToFIPmtStsRpt>TxInfAndSts>OrgnlTxRef>IntrBkSttlmDt,omitempty"`
	ProcessCode        string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>PrcCd,omitempty" validate:"required,max=6"`                     //itmx add
	MessageType        string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>MTI,omitempty" validate:"required,max=4"`                       //itmx add
	MemberId           string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>Agent>FinInstnId>ClrSysMmbId>MmbId,omitempty" validate:"max=3"` //itmx add
	MerchantType       string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>MchTp,omitempty" validate:"required,max=4"`              //itmx add
	TermType           string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>TTp,omitempty" validate:"required,max=2"`                //itmx add
	OneTimeAuth        string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>Ota,omitempty" validate:"max=16"`                        //itmx add
	OrgInstructionId   string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>OrgnlInstrId,omitempty" validate:"required,max=1"`       //itmx add
	RecvType           string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>TpRecv,omitempty" validate:"required,max=1"`             //itmx add
	SendType           string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>TpSend,omitempty" validate:"required,max=1"`             //itmx add
	RecvTaxId          string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>RecvTaxId,omitempty" validate:"max=13"`                  //itmx add
	SendTaxId          string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>SendTaxId,omitempty" validate:"max=13"`                  //itmx add
	MerchantBillId     string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>MchBillId,omitempty" validate:"required,max=35"`         //itmx add
	BillDisNameTha     string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>BillDisNmTHA,omitempty" validate:"required,max=50"`      //itmx add
	BillDisNameEng     string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>BillDisNmENG,omitempty" validate:"required,max=50"`      //itmx add
	CustDisNameEng     string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>CustDisNm,omitempty" validate:"required,max=50"`         //itmx add
	BillRef1           string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>Ref1,omitempty" validate:"max=20"`                       //itmx add
	BillRef2           string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>Ref2,omitempty" validate:"max=20"`                       //itmx add
	BillRef3           string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>Ref3,omitempty" validate:"max=20"`                       //itmx add
	AddiNote           string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>AddiNote,omitempty" validate:"max=50"`                   //itmx add
	ReqExecDate        string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>ReqdExctnDt,omitempty" validate:"max=50"`                //itmx add
	ProxyType          string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>PrxTp,omitempty" validate:"max=50"`                      //itmx add
	ProxyVal           string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>PrxVal,omitempty" validate:"max=50"`                     //itmx add
	LocalAmount        AmtField `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>LcAmt,omitempty" validate:"max=24"`                      //itmx add
	CountryCode        string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>CtryCd,omitempty" validate:"max=2"`                      //itmx add
	VatRate            string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>VatRate,omitempty" validate:"max=35"`                    //itmx add
	Vat                string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>Vat,omitempty" validate:"max=35"`                        //itmx add
	IncomeType         string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>TpInc,omitempty" validate:"max=35"`                      //itmx add
	WithHoldTaxRate    string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>WhtTaxRate,omitempty" validate:"max=35"`                 //itmx add
	WithHoldTax        string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>WhtTax,omitempty" validate:"max=35"`                     //itmx add
	WithHoldTaxCond    string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>WhtTaxCon,omitempty" validate:"max=35"`                  //itmx add
	SendIdType         string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>SendIdTp,omitempty" validate:"max=4"`                    //itmx add
	RecvIdType         string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>RecvIdTp,omitempty" validate:"max=4"`                    //itmx add
	RecvProxyType      string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>RecvPrxTp,omitempty" validate:"max=4"`                   //itmx add
	RecvProxyVal       string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>RecvPrxVal,omitempty" validate:"max=128"`                //itmx add
	SubSwitchId        string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>SubId,omitempty" validate:"max=3"`                       //itmx add
	DrAddressLine      string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>Dbtr>AdrLine,omitempty" validate:"max=69"`               //itmx add
	DrBirthDate        string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>Dbtr>BirthDt,omitempty" validate:"max=10"`               //itmx add
	DrBirthCity        string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>Dbtr>CityOfBirth,omitempty" validate:"max=35"`           //itmx add
	DrBirthCountry     string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>Dbtr>CtryOfBirth,omitempty" validate:"max=35"`           //itmx add
	SendBicCode        string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>DbtrAgtBICFI,omitempty" validate:"max=35"`               //itmx add
	RecvBicCode        string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>CdtrAgtBICFI,omitempty" validate:"max=35"`               //itmx add
	PurposeCode        string   `xml:"FIToFIPmtStsRpt>SplmtryData>Envlp>AdditionalData>OrgMsg>PurpCd,omitempty" validate:"max=4"`                      //itmx add
}
