//Version: v1.0.0.0
package models

type IN000004I struct {
	InternalAccount string `json:"internalAccount" description:"Internal account" validate:"required"`
	TranDate        string `json:"tranDate" description:"Tran date"`
	OnAccountNumber string `json:"onAccountNumber" description:"On account number"`
	OnCancelFlag    string `json:"onCancelFlag" description:"On cancel flag"`
	FlowNo          string `json:"flowNo" description:"Flow no"`
	PageNo          int    `json:"pageNo" description:"Page no"`
	PageRecCount    int    `json:"pageRecCount" description:"Page rec count"`
}

type IN000004O struct {
	InternalAccount    string          `json:"internalAccount" description:"Internal account"`
	Title              string          `json:"title" description:"Title"`
	DebCrtFlag         string          `json:"debCrtFlag" description:"Deb crt flag"`
	TranAmount         float64         `json:"tranAmount" description:"Tran amount"`
	OnAccountNumber    string          `json:"onAccountNumber" description:"On account number"`
	TranDate           string          `json:"tranDate" description:"Tran date"`
	FlowNo             string          `json:"flowNo" description:"Flow no"`
	ReverseCnlFlag     string          `json:"reverseCnlFlag" description:"Reverse cnl flag"`
	OnCancelFlag       string          `json:"onCancelFlag" description:"On cancel flag"`
	OrgTransactionDate string          `json:"orgTransactionDate" description:"Org transaction date"`
	OrgGlobalBizSeqNo  string          `json:"orgGlobalBizSeqNo" description:"Org global biz seq no"`
	CanceledAmount     float64         `json:"canceledAmount" description:"Canceled amount"`
	OnAccountType      string          `json:"onAccountType" description:"On account type"`
	Item               []IN000004OItem `json:"item" description:"Item"`
}

type IN000004OItem struct {
	TranDate           string  `json:"tranDate" description:"Tran date"`
	TranTime           string  `json:"tranTime" description:"Tran time"`
	TranAmount         float64 `json:"tranAmount" description:"Tran amount"`
	OnAccountNumber    string  `json:"onAccountNumber" description:"On account number"`
	OrgChannelType     string  `json:"orgChannelType" description:"Org channel type"`
	Flow               string  `json:"flow" description:"Flow"`
	ReverseCnlFlag     string  `json:"reverseCnlFlag" description:"Reverse cnl flag"`
	OrgTransactionDate string  `json:"orgTransactionDate" description:"Org transaction date"`
	OrgGlobalBizSeqNo  string  `json:"orgGlobalBizSeqNo" description:"Org global biz seq no"`
	OnCancelMark       string  `json:"onCancelMark" description:"On cancel mark"`
}

func (*IN000004I) GetServiceKey() string {
	return "sin0000004"
}
