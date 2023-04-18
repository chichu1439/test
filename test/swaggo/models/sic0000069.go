package models

import "github.com/shopspring/decimal"

type IC000069Request struct {
	CustomerId      string            `json:"customerId" description:"Customer id" validate:"required"`
	ProductId       string            `json:"productId" description:"Product id" validate:"required"`
	LoanAmount      decimal.Decimal   `json:"loanAmount" description:"Loan Amount" validate:"required"`
	PeriodNum       int               `json:"periodNum" description:"Period Num"`
	RepayCycle      string            `json:"repayCycle" description:"Repay Cycle" validate:"required"`
	RepayMethod     string            `json:"repayMethod" description:"Repay Method" validate:"required"`
	GracePeriod     int               `json:"gracePeriod" description:"Grace Period"`
	LoanPurpose     string            `json:"loanPurpose" description:"Loan Purpose"`
	AppDate         string            `json:"appDate" description:"App Date"`
	Maker           string            `json:"maker" description:"Maker"`
	MakerComment    string            `json:"makerComment" description:"Maker Comment"`
	LoanLinkAccount []LinkAccountInfo `json:"loanLinkAccount" description:"Loan link account"`
}
type IC000069Response struct {
	LoanReceiptNum       string          `json:"loanReceiptNum"`
	AccountingAccountNum string          `json:"accountingAccountNum"`
	RepaymentPlanList    []RepaymentPlan `json:"repaymentPlanList" description:"Repayment plan"`
}

type LinkAccountInfo struct {
	CreditApplySerialNum string `json:"creditApplySerialNum"`
	CustomerId           string `json:"customerId"`
	LinkAccountUseType   string `json:"linkAccountUseType"`
	LocalBankFlag        string `json:"localBankFlag"`
	BankNum              string `json:"bankNum"`
	BankName             string `json:"bankName"`
	AccountNum           string `json:"accountNum"`
	AccountName          string `json:"accountName"`
	ExpireMonth          string `json:"expireMonth"`
	ExpireYear           string `json:"expireYear"`
	SecurityCode         string `json:"securityCode"`
	UseStatus            string `json:"useStatus"  description:"1:have been used"`
}

func (*IC000069Request) GetServiceKey() string {
	return "IC000069"
}
