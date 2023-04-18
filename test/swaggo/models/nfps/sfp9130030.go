// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP913030Request struct {
	MsgId   string `json:"msgId"`
	TransId string `json:"transId" validate:"required"`
}

func (t *FP913030Request) GetServiceKey() string {
	return "FP913030"
}

type FP913030Response struct {
	Records []FP913030Record
}

type FP913030Record struct {
	MsgId                 string  `json:"msgId,omitempty" validate:"required"`
	MsgType               string  `json:"msgType,omitempty" validate:"required"`
	TransId               string  `json:"transId,omitempty" validate:"required"`
	InstructionId         string  `json:"instructionId,omitempty"`
	EndToEndId            string  `json:"endToEndId,omitempty"`
	TransNo               string  `json:"transNo,omitempty"`
	ClrSysRef             string  `json:"clrSysRef,omitempty"`
	CreateDatetime        string  `json:"createDatetime,omitempty"`
	TransCurrency         string  `json:"transCurrency,omitempty"`
	TransAmt              float64 `json:"transAmt,omitempty"`
	TransType             string  `json:"transType,omitempty" validate:"required"`
	PaymentPurposeCode    string  `json:"paymentPurposeCode,omitempty"`
	ProcessCode           string  `json:"processCode,omitempty"`
	DebtorMemberId        string  `json:"debtorMemberId,omitempty" validate:"required"`
	DebtorMemberName      string  `json:"debtorMemberName,omitempty" validate:"required"`
	DebtorAcctId          string  `json:"debtorAcctId,omitempty"`
	DebtorAcctName        string  `json:"debtorAcctName,omitempty"`
	DebtorAcctType        string  `json:"debtorAcctType,omitempty"`
	CreditorMemberId      string  `json:"creditorMemberId,omitempty" validate:"required"`
	CreditorMemberName    string  `json:"creditorMemberName,omitempty" validate:"required"`
	CreditorAcctId        string  `json:"creditorAcctId,omitempty"`
	CreditorAcctName      string  `json:"creditorAcctName,omitempty"`
	CreditorAcctType      string  `json:"creditorAcctType,omitempty"`
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
}
