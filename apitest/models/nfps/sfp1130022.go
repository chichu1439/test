// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP113022I struct {
	MsgId              string  `json:"msgId" `
	FromMemberId       string  `json:"fromMemberId" `
	ToMemberId         string  `json:"toMemberId" `
	TransId            string  `json:"transId"`
	InstructionId      string  `json:"instructionId"`
	EndToEndId         string  `json:"endToEndId"`
	CreateDatetime     string  `json:"createDatetime"`
	TransCurrency      string  `json:"transCurrency"`
	TransAmt           float64 `json:"transAmt"`
	TransDate          string  `json:"transDate"`
	TransType          string  `json:"transType"`
	AcctVerifyOption   string  `json:"acctVerifyOption"`
	PaymentPurposeCode string  `json:"paymentPurposeCode"`
	ProcessCode        string  `json:"processCode"`
	DebtorMemberId     string  `json:"debtorMemberId"`
	DebtorMemberName   string  `json:"debtorMemberName"`
	DebtorAcctId       string  `json:"debtorAcctId"`
	DebtorAcctName     string  `json:"debtorAcctName"`
	DebtorAcctType     string  `json:"debtorAcctType"`
	CreditorMemberId   string  `json:"creditorMemberId"`
	CreditorMemberName string  `json:"creditorMemberName"`
	CreditorAcctId     string  `json:"creditorAcctId"`
	CreditorAcctName   string  `json:"creditorAcctName"`
	CreditorAcctType   string  `json:"creditorAcctType"`
	MessageContent     []byte  `json:"messageContent"`
}

func (*FP113022I) GetServiceKey() string {
	return "FP113022"
}

type FP113022O struct {
	MsgId                 string `json:"msgId"`
	FromMemberId          string `json:"fromMemberId" `
	ToMemberId            string `json:"toMemberId" `
	TransStatus           string `json:"transStatus"`
	TransCode             string `json:"transCode"`
	TransRejectReason     string `json:"transRejectReason"`
	TransRejectReasonInfo string `json:"transRejectReasonInfo"`
	MessageContent        []byte `json:"messageContent"`
}
