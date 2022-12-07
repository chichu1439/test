// Package models will define request and response message struct
// Version: v0.0.1
package models

import "encoding/xml"

type FP110006I struct {
	MsgId           string   `json:"msgId,omitempty"`
	FrMmbId         string   `json:"frMmbId,omitempty"`
	ToMmbId         string   `json:"toMmbId,omitempty"`
	CreDtTm         string   `json:"creDtTm,omitempty"`
	OrgnlMsgId      string   `json:"orgnlMsgId,omitempty"`
	OrgnlMsgNmId    string   `json:"orgnlMsgNmId,omitempty"`
	GrpStatus       string   `json:"grpStatus,omitempty"`
	RjctRsn         string   `json:"rjctRsn,omitempty"`
	RjctRsnInf      string   `json:"rjctRsnInf,omitempty"`
	OrgnlEndToEndId string   `json:"orgnlEndToEndId,omitempty"`
	OrgnlTxId       string   `json:"orgnlTxId,omitempty"`
	TxStatus        string   `json:"txStatus,omitempty"`
	TxRjctRsn       string   `json:"txRjctRsn,omitempty"`
	TxRjctRsnInf    string   `json:"txRjctRsnInf,omitempty"`
	AccptncDtTm     string   `json:"accptncDtTm,omitempty"`
	SttlmDate       string   `json:"sttlmDate,omitempty"`
	SttlmTime       string   `json:"sttlmTime,omitempty"`
	SttlmStatus     string   `json:"sttlmStatus,omitempty"`
	ClrSysRef       string   `json:"clrSysRef,omitempty"`
	IntrBkSttlmAmt  AmtField `json:"intrBkSttlmAmt,omitempty"`
	IntrBkSttlmDt   string   `json:"intrBkSttlmDt,omitempty"`
	ServiceType     string   `json:"serviceType,omitempty"`
	TranNo          string   `json:"tranNo,omitempty"`
	MsgPaticipate   string   `json:"msgPaticipate,omitempty"` // 报文参与者
	MsgService      string   `json:"msgService"`
	ParticipantUrl  string   `json:"participantUrl,omitempty"`
}

func (*FP110006I) GetServiceKey() string {
	return "FP110006"
}

type FP110006O struct {
	XMLName      xml.Name     `xml:"FpsMsg"`
	NbOfMsgs     string       `xml:"NbOfMsgs,omitempty" validate:"required,max=15"`
	XmlHead      XmlHead      `xml:"FpsPylds>BizData>AppHdr"`
	FP110006ODoc FP110006ODoc `xml:"FpsPylds>BizData>Document"`
}

type FP110006ODoc struct {
	MsgId           string   `xml:"FIToFIPmtStsRpt>GrpHdr>MsgId,omitempty"`
	CreDtTm         string   `xml:"FIToFIPmtStsRpt>GrpHdr>CreDtTm,omitempty"`
	OrgnlMsgId      string   `xml:"FIToFIPmtStsRpt>OrgnlGrpInfAndSts>OrgnlMsgId,omitempty"`
	OrgnlMsgNmId    string   `xml:"FIToFIPmtStsRpt>OrgnlGrpInfAndSts>OrgnlMsgNmId,omitempty"`
	GrpStatus       string   `xml:"FIToFIPmtStsRpt>OrgnlGrpInfAndSts>GrpSts,omitempty"`
	RjctRsn         string   `xml:"FIToFIPmtStsRpt>OrgnlGrpInfAndSts>StsRsnInf>Rsn>Prtry,omitempty"`
	RjctRsnInf      string   `xml:"FIToFIPmtStsRpt>OrgnlGrpInfAndSts>StsRsnInf>AddtlInf,omitempty"`
	OrgnlEndToEndId string   `xml:"FIToFIPmtStsRpt>TxInfAndSts>OrgnlEndToEndId,omitempty"`
	OrgnlTxId       string   `xml:"FIToFIPmtStsRpt>TxInfAndSts>OrgnlTxId,omitempty"`
	TxStatus        string   `xml:"FIToFIPmtStsRpt>TxInfAndSts>TxSts,omitempty"`
	TxRjctRsn       string   `xml:"FIToFIPmtStsRpt>TxInfAndSts>StsRsnInf>Rsn>Prtry,omitempty"`
	TxRjctRsnInf    string   `xml:"FIToFIPmtStsRpt>TxInfAndSts>StsRsnInf>AddtlInf,omitempty"`
	AccptncDtTm     string   `xml:"FIToFIPmtStsRpt>TxInfAndSts>AccptncDtTm,omitempty"`
	ClrSysRef       string   `xml:"FIToFIPmtStsRpt>TxInfAndSts>ClrSysRef,omitempty"`
	IntrBkSttlmAmt  AmtField `xml:"FIToFIPmtStsRpt>TxInfAndSts>OrgnlTxRef>IntrBkSttlmAmt,omitempty"`
	IntrBkSttlmDt   string   `xml:"FIToFIPmtStsRpt>TxInfAndSts>OrgnlTxRef>IntrBkSttlmDt,omitempty"`
}

type FP110006Outgress struct {
	Body []byte `json:"body"`
}

func (*FP110006Outgress) GetServiceKey() string {
	return "FP210006"
}

type FP110006OutgressResponse struct {
	Body []byte `json:"body"`
}