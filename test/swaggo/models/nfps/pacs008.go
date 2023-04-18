package models

import "encoding/xml"

type FPSPacs008 struct {
	XMLName  xml.Name `xml:"FpsMsg"`
	NumOfMsg string   `xml:"NbOfMsgs"`
	XMLHead  Head001  `xml:"FpsPylds>BizData>AppHdr"`
	Document Pacs008  `xml:"FpsPylds>BizData>Document"`
}

type Pacs008 struct {
	MsgId          string `xml:"FIToFICstmrCdtTrf>GrpHdr>MsgId,omitempty" validate:"required,max=35"`
	CreateDatetime string `xml:"FIToFICstmrCdtTrf>GrpHdr>CreDtTm,omitempty" validate:"required,max=35"`
	NumOfTrans     string `xml:"FIToFICstmrCdtTrf>GrpHdr>NbOfTxs,omitempty" validate:"required,max=15"`
	//TotInterbankStlmtAmt string   `xml:"FIToFICstmrCdtTrf>GrpHdr>TtlIntrBkSttlmAmt,omitempty" validate:"required,max=24"` //itmx add
	InterbankStlmtDate string   `xml:"FIToFICstmrCdtTrf>GrpHdr>IntrBkSttlmDt,omitempty" validate:"required,max=10"` //itmx add
	StlmtMethod        string   `xml:"FIToFICstmrCdtTrf>GrpHdr>SttlmInf>SttlmMtd,omitempty" validate:"in=[STLMT-METHOD],max=4"`
	ClrSysId           string   `xml:"FIToFICstmrCdtTrf>GrpHdr>SttlmInf>ClrSys>Prtry,omitempty" validate:"max=35"`
	ClrSysCode         string   `xml:"FIToFICstmrCdtTrf>GrpHdr>SttlmInf>ClrSys>Cd,omitempty" validate:"in=[ExternalCashClearingSystem1Code],max=35"` //itmx add
	InstructionId      string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>PmtId>InstrId,omitempty" validate:"max=35"`
	EndToEndId         string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>PmtId>EndToEndId,omitempty" validate:"required,max=35"`
	TransId            string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>PmtId>TxId,omitempty" validate:"required,max=35"`
	ClrSysRef          string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>PmtId>ClrSysRef,omitempty" validate:"max=35"`
	SvcLvlCode         string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>PmtTpInf>SvcLvl>Cd,omitempty" validate:"max=4"` //itmx add
	AcctVerifyOption   string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>PmtTpInf>LclInstrm>Prtry,omitempty" validate:"max=35"`
	CategoryPurpose    string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>PmtTpInf>CtgyPurp>Prtry,omitempty" validate:"max=35"`
	InterbankStlmtAmt  AmtField `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>IntrBkSttlmAmt,omitempty" validate:"required,max=24"`
	//InterbankStlmtDate   string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>IntrBkSttlmDt,omitempty" validate:"max=10"`         // removed
	CreditDatetime      string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>SttlmTmIndctn>CdtDtTm,omitempty" validate:"max=19"` // removed
	InstructedAmt       AmtField `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>InstdAmt,omitempty" validate:"max=24"`              // removed
	ChargeBearerType    string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>ChrgBr,omitempty" validate:"in=[CHARGE-BEARER],max=4"`
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
	Unstructured        string   `xml:"FIToFICstmrCdtTrf>CdtTrfTxInf>RmtInf>Ustrd,omitempty" validate:"max=140"`
	ProcessCode         string   `xml:"FIToFICstmrCdtTrf>SplmtryData>Envlp>AdditionalData>PrcCd,omitempty" validate:"required,max=6"` //itmx add
	MessageType         string   `xml:"FIToFICstmrCdtTrf>SplmtryData>Envlp>AdditionalData>MTI,omitempty" validate:"required,max=4"`   //itmx add
	MemberId            string   `xml:"FIToFICstmrCdtTrf>SplmtryData>Agent>FinInstnId>ClrSysMmbId>MmbId,omitempty" validate:"max=3"`  //itmx add
	MerchantType        string   `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>MchTp,omitempty" validate:"required,max=4"`               //itmx add
	TermType            string   `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>TTp,omitempty" validate:"required,max=2"`                 //itmx add
	OneTimeAuth         string   `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>Ota,omitempty" validate:"max=16"`                         //itmx add
	OrgInstructionId    string   `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>OrgnlInstrId,omitempty" validate:"required,max=35"`       //itmx add
	RecvType            string   `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>TpRecv,omitempty" validate:"required,max=1"`              //itmx add
	SendType            string   `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>TpSend,omitempty" validate:"required,max=1"`              //itmx add
	RecvTaxId           string   `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>RecvTaxId,omitempty" validate:"max=13"`                   //itmx add
	SendTaxId           string   `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>SendTaxId,omitempty" validate:"max=13"`                   //itmx add
	MerchantBillId      string   `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>MchBillId,omitempty" validate:"required,max=35"`          //itmx add
	BillDisNameTha      string   `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>BillDisNmTHA,omitempty" validate:"required,max=50"`       //itmx add
	BillDisNameEng      string   `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>BillDisNmENG,omitempty" validate:"required,max=50"`       //itmx add
	CustDisNameEng      string   `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>CustDisNm,omitempty" validate:"required,max=50"`          //itmx add
	BillRef1            string   `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>Ref1,omitempty" validate:"max=20"`                        //itmx add
	BillRef2            string   `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>Ref2,omitempty" validate:"max=20"`                        //itmx add
	BillRef3            string   `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>Ref3,omitempty" validate:"max=20"`                        //itmx add
	AddiNote            string   `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>AddiNote,omitempty" validate:"max=50"`                    //itmx add
	ReqExecDate         string   `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>ReqdExctnDt,omitempty" validate:"max=50"`                 //itmx add
	ProxyType           string   `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>PrxTp,omitempty" validate:"max=50"`                       //itmx add
	ProxyVal            string   `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>PrxVal,omitempty" validate:"max=50"`                      //itmx add
	LocalAmount         AmtField `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>LcAmt,omitempty" validate:"max=24"`                       //itmx add
	CountryCode         string   `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>CtryCd,omitempty" validate:"max=2"`                       //itmx add
	VatRate             string   `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>VatRate,omitempty" validate:"max=35"`                     //itmx add
	Vat                 string   `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>Vat,omitempty" validate:"max=35"`                         //itmx add
	IncomeType          string   `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>TpInc,omitempty" validate:"max=35"`                       //itmx add
	WithHoldTaxRate     string   `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>WhtTaxRate,omitempty" validate:"max=35"`                  //itmx add
	WithHoldTax         string   `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>WhtTax,omitempty" validate:"max=35"`                      //itmx add
	WithHoldTaxCond     string   `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>WhtTaxCon,omitempty" validate:"max=35"`                   //itmx add
	DrAgtSubId          string   `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>DbtrAgtSubId,omitempty" validate:"max=4"`                 //itmx add
	CrAgtSubId          string   `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>CdtrAgtSubId,omitempty" validate:"max=4"`                 //itmx add
	DrAddressLine       string   `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>Dbtr>AdrLine,omitempty" validate:"max=69"`                //itmx add
	DrBirthDate         string   `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>Dbtr>BirthDt,omitempty" validate:"max=10"`                //itmx add
	DrBirthCity         string   `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>Dbtr>CityOfBirth,omitempty" validate:"max=35"`            //itmx add
	DrBirthCountry      string   `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>Dbtr>CtryOfBirth,omitempty" validate:"max=35"`            //itmx add
	SendBicCode         string   `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>DbtrAgtBICFI,omitempty" validate:"max=35"`                //itmx add
	RecvBicCode         string   `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>CdtrAgtBICFI,omitempty" validate:"max=35"`                //itmx add
	PurposeCode         string   `xml:"FIToFICstmrCdtTrf>SplmtryData>OrgMsg>PurpCd,omitempty" validate:"max=4"`                       //itmx add
}
