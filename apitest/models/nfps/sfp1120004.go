// Package models will define request and response message struct
// Version: v0.0.1
package models

import "encoding/xml"

type FP112004I struct {
	MsgId           string   `json:"msgId,omitempty"`
	MsgService      string   `json:"msgService"`
	FrMmbId         string   `json:"frMmbId,omitempty"`
	ToMmbId         string   `json:"toMmbId,omitempty"`
	OrgnlMsgId      string   `json:"orgnlMsgId,omitempty"`
	AssgnrClrSysId  string   `json:"assgnrClrSysId,omitempty"`
	AssgneClrSysId  string   `json:"assgneClrSysId,omitempty"`
	CreDtTm         string   `json:"creDtTm,omitempty"`
	OrgnlInstrId    string   `json:"orgnlInstrId,omitempty"`
	OrgnlEndToEndId string   `json:"orgnlEndToEndId,omitempty"`
	OrgnlTxId       string   `json:"orgnlTxId,omitempty"`
	OrgnlClrSysRef  string   `json:"orgnlClrSysRef,omitempty"`
	RjctRsn         string   `json:"rjctRsn,omitempty"`
	RjctRsnInf      string   `json:"rjctRsnInf,omitempty"`
	TranNo          string   `json:"tranNo,omitempty"`
	MsgPaticipate   string   `json:"msgPaticipate,omitempty"` // 报文参与者
	IntrBkSttlmAmt  AmtField `json:"intrBkSttlmAmt,omitempty"`
	IntrBkSttlmDt   string   `json:"intrBkSttlmDt,omitempty"`
	SttlmDate       string   `json:"sttlmDate,omitempty"`
	SttlmTime       string   `json:"sttlmTime,omitempty"`
	SttlmStatus     string   `json:"sttlmStatus,omitempty"`
	TxStatus        string   `json:"txStatus,omitempty"`
	TxRjctRsn       string   `json:"txRjctRsn,omitempty"`
	TxRjctRsnInf    string   `json:"txRjctRsnInf,omitempty"`
	ServiceType     string   `json:"serviceType,omitempty"`
	ParticipantUrl  string   `json:"participantUrl,omitempty"`
	MessageContent  []byte   `json:"messageContent"`
}

func (*FP112004I) GetServiceKey() string {
	return "FP112004"
}

type FP112004Outgress struct {
	Body []byte `json:"body"`
}

type FP112004OutgressResponse struct {
	Body []byte `json:"body"`
}

func (*FP112004Outgress) GetServiceKey() string {
	return "FP212004"
}

type FP112004ServiceResult struct {
	ServiceResult string `json:"serviceResult,omitempty"`
}

type FP112004O struct {
	XMLName       xml.Name     `xml:"FpsMsg"`
	NbOfMsgs      string       `xml:"NbOfMsgs,omitempty" validate:"required,max=15"`
	XmlHead       XmlHead      `xml:"FpsPylds>BizData>AppHdr"`
	FP112004OODoc FP112004ODoc `xml:"FpsPylds>BizData>Document"`
}

type FP112004ODoc struct {
	MsgId           string `xml:"FIToFIPmtCxlReq>Assgnmt>Id,omitempty"`
	AssgnrClrSysId  string `xml:"FIToFIPmtCxlReq>Assgnmt>Assgnr>Agt>FinInstnId>ClrSysMmbId>MmbId,omitempty"`
	AssgneClrSysId  string `xml:"FIToFIPmtCxlReq>Assgnmt>Assgne>Agt>FinInstnId>ClrSysMmbId>MmbId,omitempty"`
	CreDtTm         string `xml:"FIToFIPmtCxlReq>Assgnmt>CreDtTm,omitempty"`
	OrgnlInstrId    string `xml:"FIToFIPmtCxlReq>Undrlyg>TxInf>OrgnlInstrId,omitempty"`
	OrgnlEndToEndId string `xml:"FIToFIPmtCxlReq>Undrlyg>TxInf>OrgnlEndToEndId,omitempty"`
	OrgnlTxId       string `xml:"FIToFIPmtCxlReq>Undrlyg>TxInf>OrgnlTxId,omitempty"`
	OrgnlClrSysRef  string `xml:"FIToFIPmtCxlReq>Undrlyg>TxInf>OrgnlClrSysRef,omitempty"`
	RjctRsn         string `xml:"FIToFIPmtCxlReq>Undrlyg>TxInf>CxlRsnInf>Rsn>Prtry,omitempty"`
	RjctRsnInf      string `xml:"FIToFIPmtCxlReq>Undrlyg>TxInf>CxlRsnInf>Rsn>AddtlInf,omitempty"`
}
