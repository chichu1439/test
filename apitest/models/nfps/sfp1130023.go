// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP113023I struct {
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

func (*FP113023I) GetServiceKey() string {
	return "FP113023"
}

type FP113023O struct {
	MsgId                 string `json:"msgId"`
	FromMemberId          string `json:"fromMemberId" `
	ToMemberId            string `json:"toMemberId" `
	TransStatus           string `json:"transStatus"`
	TransCode             string `json:"transCode"`
	TransRejectReason     string `json:"transRejectReason"`
	TransRejectReasonInfo string `json:"transRejectReasonInfo"`
	MessageContent        []byte `json:"messageContent"`
}
type FP113023Outgress struct {
	Body []byte `json:"body"`
}

func (*FP113023Outgress) GetServiceKey() string {
	return "FP213023"
}

type FP1130023OutgressResponse struct {
	Body []byte `json:"body"`
}
