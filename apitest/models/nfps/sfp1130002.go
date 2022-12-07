// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP113002Request struct {
	MsgId                 string  `json:"msgId" `
	FromMemberId          string  `json:"fromMemberId" `
	ToMemberId            string  `json:"toMemberId" `
	TransId               string  `json:"transId"`
	InstructionId         string  `json:"instructionId"`
	EndToEndId            string  `json:"endToEndId"`
	CreateDatetime        string  `json:"createDatetime"`
	TransCurrency         string  `json:"transCurrency"`
	TransAmt              float64 `json:"transAmt"`
	TransDate             string  `json:"transDate"`
	TransType             string  `json:"transType"`
	AcctVerifyOption      string  `json:"acctVerifyOption"`
	PaymentPurposeCode    string  `json:"paymentPurposeCode"`
	ProcessCode           string  `json:"processCode"`
	DebtorMemberId        string  `json:"debtorMemberId"`
	DebtorMemberPath      string  `json:"debtorMemberPath"`
	DebtorMemberSu        string  `json:"debtorMemberSu"`
	DebtorMemberContact   string  `json:"debtorMemberContact"`
	DebtorMemberName      string  `json:"debtorMemberName"`
	DebtorAcctId          string  `json:"debtorAcctId"`
	DebtorAcctName        string  `json:"debtorAcctName"`
	DebtorAcctType        string  `json:"debtorAcctType"`
	CreditorMemberId      string  `json:"creditorMemberId"`
	CreditorMemberSu      string  `json:"creditorMemberSu"`
	CreditorMemberPath    string  `json:"creditorMemberPath"`
	CreditorMemberContact string  `json:"creditorMemberContact"`
	CreditorMemberName    string  `json:"creditorMemberName"`
	CreditorAcctId        string  `json:"creditorAcctId"`
	CreditorAcctName      string  `json:"creditorAcctName"`
	CreditorAcctType      string  `json:"creditorAcctType"`
	MessageContent        []byte  `json:"messageContent"`
}

func (*FP113002Request) GetServiceKey() string {
	return "FP113002"
}

type FP113002Response struct {
	MsgId                 string `json:"msgId"`
	FromMemberId          string `json:"fromMemberId" `
	ToMemberId            string `json:"toMemberId" `
	ProcessCode           string `json:"processCode"`
	SettleDate            string `json:"settleDate"`
	SettleTime            string `json:"settleTime"`
	SettleStatus          string `json:"settleStatus"`
	AcctVerifyType        string `json:"acctVerifyType"`
	ReconcileFlag         string `json:"reconcileFlag"`
	RejectReason          string `json:"rejectReason"`
	RejectReasonInfo      string `json:"rejectReasonInfo"`
	TransStatus           string `json:"transStatus"`
	TransRejectReason     string `json:"transRejectReason"`
	TransRejectReasonInfo string `json:"transRejectReasonInfo"`
	AcceptDatetime        string `json:"acceptDatetime"`
}
