package models

import "encoding/xml"

type FPSPrxy003 struct {
	XMLName  xml.Name `xml:"FpsMsg"`
	NumOfMsg string   `xml:"NbOfMsgs"`
	XMLHead  Head001  `xml:"FpsPylds>BizData>AppHdr"`
	Document Prxy003  `xml:"FpsPylds>BizData>Document"`
}

type Prxy003 struct {
	MsgId             string   `xml:"PrxyLookup>GrpHdr>MsgId,omitempty" validate:"required,max=35"`
	CreateDatetime    string   `xml:"PrxyLookup>GrpHdr>CreDtTm,omitempty" validate:"required,max=35"`
	MmbId             string   `xml:"PrxyLookup>GrpHdr>MsgSndr>Agt>FinInstnId>ClrSysMmbId>MmbId,omitempty" validate:"required,max=3"`
	InstructionId     string   `xml:"PrxyLookup>CdtTrfTxInf>PmtId>InstrId,omitempty" validate:"max=35"`
	EndToEndId        string   `xml:"PrxyLookup>CdtTrfTxInf>PmtId>EndToEndId,omitempty" validate:"required,max=35"`
	InterbankStlmtAmt AmtField `xml:"PrxyLookup>CdtTrfTxInf>IntrBkSttlmAmt,omitempty" validate:"required,max=18"`
	DrName            string   `xml:"PrxyLookup>CdtTrfTxInf>Dbtr>Nm,omitempty" validate:"max=140"`
	DrId              string   `xml:"PrxyLookup>CdtTrfTxInf>DbtrAcct>Id>Othr>Id,omitempty" validate:"max=34"`
	DrPrtry           string   `xml:"PrxyLookup>CdtTrfTxInf>DbtrAcct>Id>Othr>SchmeNm>Prtry,omitempty" validate:"max=35"`
	LookupProxyId     string   `xml:"PrxyLookup>Lookup>PrxyOnly>Id,omitempty" validate:"max=35"`
	LookupProxyType   string   `xml:"PrxyLookup>Lookup>PrxyOnly>PrxyRtrvl>Tp,omitempty" validate:"max=12"`
	LookupProxyVal    string   `xml:"PrxyLookup>Lookup>PrxyOnly>PrxyRtrvl>Val,omitempty" validate:"max=128"`
	LookupAccId       string   `xml:"PrxyLookup>Lookup>AccOnly>Id,omitempty" validate:"max=35"`
	LookupAccDisNm    string   `xml:"PrxyLookup>Lookup>AccOnly>DsplNm,omitempty" validate:"max=140"`
	LookupAccAgtMmbId string   `xml:"PrxyLookup>Lookup>AccOnly>Agt>FinInstnId>ClrSysMmbId>MmbId,omitempty" validate:"max=11"`
	LookupAccAcId     string   `xml:"PrxyLookup>Lookup>AccOnly>Acct>Id>Othr>Id,omitempty" validate:"max=34"`
	LookupAccAcNm     string   `xml:"PrxyLookup>Lookup>AccOnly>Acct>Nm,omitempty" validate:"max=140"`
	AdditionalData    string   `xml:"PrxyLookup>SplmtryData>Envlp>AdditionalData,omitempty" validate:"max=400"`
}
