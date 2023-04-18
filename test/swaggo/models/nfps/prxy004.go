package models

import "encoding/xml"

type FPSPrxy004 struct {
	XMLName  xml.Name `xml:"FpsMsg"`
	NumOfMsg string   `xml:"NbOfMsgs"`
	XMLHead  Head001  `xml:"FpsPylds>BizData>AppHdr"`
	Document Prxy004  `xml:"FpsPylds>BizData>Document"`
}

type Prxy004 struct {
	MsgId                  string   `xml:"PrxyLookUpRsp>GrpHdr>MsgId,omitempty" validate:"required,max=35"`
	CreateDatetime         string   `xml:"PrxyLookUpRsp>GrpHdr>CreDtTm,omitempty" validate:"required,max=35"`
	MmbId                  string   `xml:"PrxyLookUpRsp>GrpHdr>MsgRcpt>Agt>ClrSysMmbId>MmbId,omitempty" validate:"required,max=3"`
	OriginalInstructionId  string   `xml:"PrxyLookUpRsp>OrgnlTxRef>OrgnlInstrId,omitempty" validate:"max=35"`
	OriginalEndToEndId     string   `xml:"PrxyLookUpRsp>OrgnlTxRef>OrgnlEndToEndId,omitempty" validate:"required,max=35"`
	OriginalMsgId          string   `xml:"PrxyLookUpRsp>OrgnlGrpInf>OrgnlMsgId,omitempty" validate:"required,max=35"`
	OriginalMsgNameId      string   `xml:"PrxyLookUpRsp>OrgnlGrpInf>OrgnlMsgNmId,omitempty" validate:"required,max=35"`
	OriginalCreateDatetime string   `xml:"PrxyLookUpRsp>OrgnlGrpInf>OrgnlCreDtTm,omitempty" validate:"required,max=35"`
	InterbankStlmtAmt      AmtField `xml:"PrxyLookUpRsp>CdtTrfTxInf>IntrBkSttlmAmt,omitempty" validate:"required,max=18"`
	OriginalProxyId        string   `xml:"PrxyLookUpRsp>OrgnlLkUpVal>OrgnlId,omitempty" validate:"max=35"`
	OriginalProxyType      string   `xml:"PrxyLookUpRsp>OrgnlLkUpVal>OrgnlPrxyRtrvl>Tp,omitempty" validate:"max=12"`
	OriginalProxyVal       string   `xml:"PrxyLookUpRsp>OrgnlLkUpVal>OrgnlPrxyRtrvl>Val,omitempty" validate:"max=128"`
	//Ccy                    string   `xml:"PrxyLookUpRsp>CdtTrfTxInf>Ccy,omitempty" validate:"required,max=3"`
	//DrName                 string   `xml:"PrxyLookUpRsp>CdtTrfTxInf>Dbtr>Nm,omitempty" validate:"max=140"`
	//DrId                   string   `xml:"PrxyLookUpRsp>CdtTrfTxInf>DbtrAcct>Id>Othr>Id,omitempty" validate:"max=34"`
	//DrPrtry                string   `xml:"PrxyLookUpRsp>CdtTrfTxInf>DbtrAcct>Id>Othr>SchmeNm>Prtry,omitempty" validate:"max=35"`
	PrxRspSts         string `xml:"PrxyLookUpRsp>RgstRsp>PrxRspSts,omitempty" validate:"max=4"`
	PrxRspRsnPrtry    string `xml:"PrxyLookUpRsp>RgstRsp>StsRsnInfz>Rsn>Prtry,omitempty" validate:"max=35"`
	PrxRspAddtInf     string `xml:"PrxyLookUpRsp>RgstRsp>StsRsnInfz>AddtInf,omitempty" validate:"max=105"`
	PrxRspProxyType   string `xml:"PrxyLookUpRsp>RgstRsp>Prxy>Tp,omitempty" validate:"max=5"`
	PrxRspProxyVal    string `xml:"PrxyLookUpRsp>RgstRsp>Prxy>Val,omitempty" validate:"max=30"`
	PrxRspAccId       string `xml:"PrxyLookUpRsp>RgstRsp>Rgstr>RgstrId,omitempty" validate:"max=35"`
	PrxRspAccDisNm    string `xml:"PrxyLookUpRsp>RgstRsp>Rgstr>DsplNm,omitempty" validate:"max=50"`
	PrxRspAccAgtMmbId string `xml:"PrxyLookUpRsp>RgstRsp>Rgstr>Agt>FinInstnId>ClrSysMmbId>MmbId,omitempty" validate:"max=11"`
	PrxRspAccAcId     string `xml:"PrxyLookUpRsp>RgstRsp>Rgstr>Acct>Id>Othr>Id,omitempty" validate:"max=34"`
	PrxRspAccAcNm     string `xml:"PrxyLookUpRsp>RgstRsp>Rgstr>Acct>Nm,omitempty" validate:"max=140"`
	MchTp             string `xml:"PrxyLookUpRsp>SplmtryData>Envlp>AdditionalData>MchTp,omitempty" validate:"max=40"`
}
