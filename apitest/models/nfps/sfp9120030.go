// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP912030I struct {
	MsgId     string `json:"msgId"`
	TransId   string `json:"transId"`
	ClrSysRef string `json:"clrSysRef"`
}

func (t *FP912030I) GetServiceKey() string {
	return "FP912030"
}

type FP912030O struct {
	Records []FP912030Record
}

type FP912030Record struct {
	MsgId                 string  `json:"msgId,omitempty"`
	MsgType               string  `json:"msgType,omitempty"`
	TransId               string  `json:"transId,omitempty"`
	InstructionId         string  `json:"instructionId,omitempty"`
	EndToEndId            string  `json:"endToEndId,omitempty"`
	TransNo               string  `json:"transNo,omitempty"`
	CreateDatetime        string  `json:"createDatetime,omitempty"`
	TransCurrency         string  `json:"transCurrency,omitempty"`
	TransAmt              float64 `json:"transAmt,omitempty"`
	TransType             string  `json:"transType,omitempty"`
	PaymentPurposeCode    string  `json:"paymentPurposeCode,omitempty"`
	ProcessCode           string  `json:"processCode,omitempty"`
	DebtorMemberId        string  `json:"debtorMemberId,omitempty"`
	DebtorMemberName      string  `json:"debtorMemberName,omitempty"`
	DebtorAcctId          string  `json:"debtorAcctId,omitempty"`
	DebtorAcctName        string  `json:"debtorAcctName,omitempty"`
	CreditorMemberId      string  `json:"creditorMemberId,omitempty"`
	CreditorMemberName    string  `json:"creditorMemberName,omitempty"`
	CreditorAcctId        string  `json:"creditorAcctId,omitempty"`
	CreditorAcctName      string  `json:"creditorAcctName,omitempty"`
	ProxyType             string  `json:"proxyType,omitempty"`
	StatusCode            string  `json:"statusCode,omitempty"`
	TransRef              string  `json:"transRef,omitempty"`
	TransDate             string  `json:"transDate,omitempty"`
	TransStatus           string  `json:"transStatus,omitempty"`
	TransRejectCode       string  `json:"transRejectCode,omitempty"`
	TransRejectReason     string  `json:"transRejectReason,omitempty"`
	TransRejectReasonInfo string  `json:"transRejectReasonInfo,omitempty"`
	SettleDate            string  `json:"settleDate,omitempty"`
	SettleTime            string  `json:"settleTime,omitempty"`
	SettleStatus          string  `json:"settleStatus,omitempty"`
	MessageContent        []byte  `json:"messageContent,omitempty"`
	FileId                string  `json:"fileId,omitempty"`
}
