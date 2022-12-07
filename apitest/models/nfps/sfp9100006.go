// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP910006I struct {
	MsgId                 string `json:"msgId" `
	EndToEndId            string `json:"endToEndId" `
	TransId               string `json:"transId" `
	FromMemberId          string `json:"fromMemberId" `
	ToMemberId            string `json:"toMemberId" `
	CreateDatetime        string `json:"createDatetime" `
	NumOfTrans            string `json:"numOfTrans" `
	DrMemberId            string `json:"drMemberId" `
	CrMemberId            string `json:"crMemberId" `
	InstructedAmt         string `json:"instructedAmt" `
	InstructedCcy         string `json:"instructedCcy" `
	HandleStatusCode      string `json:"handleStatusCode" `
	ProcessCode           string `json:"processCode" `
	StlmtDate             string `json:"stlmtDate" `
	StlmtTime             string `json:"stlmtTime" `
	StlmtStatus           string `json:"stlmtStatus" `
	AcctVerifyType        string `json:"acctVerifyType" `
	ReconcileFlag         string `json:"reconcileFlag" `
	RejectReason          string `json:"rejectReason" `
	RejectReasonInfo      string `json:"rejectReasonInfo" `
	TransStatus           string `json:"transStatus" `
	TransRejectReason     string `json:"transRejectReason" `
	TransRejectReasonInfo string `json:"transRejectReasonInfo" `
	AcceptDatetime        string `json:"acceptDatetime" `
}

func (*FP910006I) GetServiceKey() string {
	return "FP910006"
}

type FP910006O struct {
	Status string `json:"status"`
}
