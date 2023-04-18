// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP110015I struct {
	FromMemberId     string  `json:"fromMemberId"`
	ToMemberId       string  `json:"toMemberId"`
	PartClearingCode string  `json:"partClearingCode"`
	MsgId            string  `json:"msgId"`
	CreDtTm          string  `json:"creDtTm"`
	NoticeId         string  `json:"noticeId"`
	NoticeTm         string  `json:"noticeTm"`
	AcctId           string  `json:"acctId"`
	AcctTpCd         string  `json:"acctTpCd"`
	NtryCcy          string  `json:"ntryCcy"`
	NtryAmt          float64 `json:"ntryAmt"`
	NtryCdtDbtInd    string  `json:"ntryCdtDbtInd"`
	NtrySts          string  `json:"ntrySts"`
	PrtryCd          string  `json:"prtryCd"`
	EndToEndId       string  `json:"endToEndId"`
	TxId             string  `json:"txId"`
	ClrSysRef        string  `json:"clrSysRef"`
	NtryDtlsAmt      float64 `json:"ntryDtlsAmt"`
	CdtDbtInd        string  `json:"cdtDbtInd"`
	DbtrAgtNmbId     string  `json:"dbtrAgtNmbId"`
	CdtrAgtNmbId     string  `json:"cdtrAgtNmbId"`
	IntrBkSttlmDt    string  `json:"intrBkSttlmDt"`
	TxDtTm           string  `json:"txDtTm"`
	ReSend           string  `json:"reSend"` //Y-重发
	MsgService       string  `json:"msgService"`
}

type FP110015O struct {
}

func (*FP110015I) GetServiceKey() string {
	return "FP110015"
}

type FP110015OutgressRequest struct {
	Body []byte `json:"body"`
}

type FP110015OutgressResponse struct {
	Body []byte `json:"body"`
}

func (*FP110015OutgressRequest) GetServiceKey() string {
	return "FP210015"
}
