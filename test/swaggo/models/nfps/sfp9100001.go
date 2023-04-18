// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP910001I struct {
	MsgId              string `json:"msgId" validate:"required"`
	EndToEndId         string `json:"endToEndId" validate:"required"`
	TransId            string `json:"transId" validate:"required"`
	FromMemberId       string `json:"fromMemberId" validate:"required"`
	ToMemberId         string `json:"toMemberId" validate:"required"`
	CreateDatetime     string `json:"createDatetime" validate:"required"`
	NumOfTrans         string `json:"numOfTrans" validate:"required"`
	DrMemberId         string `json:"drMemberId" validate:"required"`
	CrMemberId         string `json:"crMemberId" validate:"required"`
	InterbankStlmtAmt  string `json:"interbankStlmtAmt" validate:"omitempty,numeric"`
	InterbankStlmtCcy  string `json:"interbankStlmtCcy"`
	InterbankStlmtDate string `json:"interbankStlmtDate" validate:"max=10"`
	InstructedAmt      string `json:"instructedAmt" validate:"omitempty,numeric"`
	InstructedCcy      string `json:"instructedCcy" `
}

type FP910001O struct {
	MsgId                 string `json:"msgId" validate:"required"`
	EndToEndId            string `json:"endToEndId" validate:"required"`
	TransId               string `json:"transId" validate:"required"`
	FromMemberId          string `json:"fromMemberId" validate:"required"`
	ToMemberId            string `json:"toMemberId" validate:"required"`
	CreateDatetime        string `json:"createDatetime" validate:"required"`
	NumOfTrans            string `json:"numOfTrans" validate:"required"`
	DrMemberId            string `json:"drMemberId" validate:"required"`
	CrMemberId            string `json:"crMemberId" validate:"required"`
	InterbankStlmtAmt     string `json:"interbankStlmtAmt" validate:"omitempty,numeric"`
	InterbankStlmtCcy     string `json:"interbankStlmtCcy"`
	InterbankStlmtDate    string `json:"interbankStlmtDate" validate:"max=10"`
	InstructedAmt         string `json:"instructedAmt" `
	InstructedCcy         string `json:"instructedCcy" `
	HandleStatusCode      string `json:"handleStatusCode" `
	ProcessCode           string `json:"processCode" `
	StlmtDate             string `json:"stlmtDate" `
	StlmtTime             string `json:"stlmtTime" `
	StlmtStatus           string `json:"stlmtStatus" `
	AcctVerifyType        string `json:"acctVerifyType" `
	ReconcileFlag         string `json:"reconcileFlag"`
	RejectReason          string `json:"rejectReason"`
	RejectReasonInfo      string `json:"rejectReasonInfo"`
	TransStatus           string `json:"transStatus"`
	TransRejectReason     string `json:"transRejectReason"`
	TransRejectReasonInfo string `json:"transRejectReasonInfo"`
	AcceptDatetime        string `json:"acceptDatetime"`
}

func (*FP910001I) GetServiceKey() string {
	return "FP910001"
}
