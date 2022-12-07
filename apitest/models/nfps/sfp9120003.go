// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP912003I struct {
	MsgId                 string  `json:"msgId,omitempty" validate:"required"`
	MsgType               string  `json:"msgType,omitempty" validate:"required"`
	TransId               string  `json:"transId,omitempty" validate:"required"`
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
	CreditorMemberId      string  `json:"creditorMemberId,omitempty"`
	CreditorMemberName    string  `json:"creditorMemberName,omitempty"`
	StatusCode            string  `json:"statusCode,omitempty"`
	TransRef              string  `json:"transRef,omitempty"`
	TransDate             string  `json:"transDate,omitempty"`
	TransStatus           string  `json:"transStatus,omitempty"`
	TransRejectCode       string  `json:"transRejectCode,omitempty"`
	TransRejectReason     string  `json:"transRejectReason,omitempty"`
	TransRejectReasonInfo string  `json:"transRejectReasonInfo,omitempty"`
	SettleDate            string  `json:"settleDate,omitempty"`
	SettleTime            string  `json:"settleTime,omitempty"`
	MessageContent        []byte  `json:"messageContent,omitempty"`
	FileId                string  `json:"FileId,omitempty"`
	RecordType            string  `json:"recordType,omitempty"`
}

func (*FP912003I) GetServiceKey() string {
	return "FP912003"
}

type FP912003O struct {
	Status string
}
