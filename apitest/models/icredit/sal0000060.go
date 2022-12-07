package models

import "github.com/shopspring/decimal"

type AL000060I struct {
	Id        string `json:"id" description:"Id" validate:"required"`
	HashValue string `json:"hashValue" description:"Hash value" validate:"required"`
}

type AL000060O struct {
	LoanReceiptNum     string          `json:"loanReceiptNum" description:"Loan receipt number"`
	RepayTotalAmount   decimal.Decimal `json:"repayTotalAmount" description:"Repay total amount"`
	RepayPrinciple     decimal.Decimal `json:"repayPrinciple" description:"Repay principle"`
	RepayInterest      decimal.Decimal `json:"repayInterest" description:"Repay interest"`
	RepayFee           decimal.Decimal `json:"repayFee" description:"Repay fee"`
	RemainBalance      decimal.Decimal `json:"remainBalance" description:"Remain balance"`
	AccumulateInterest decimal.Decimal `json:"accumulateInterest" description:"Accumulate interest"`
	Currency           string          `json:"currency" description:"Currency"`
	PaymentList        []PaymentList   `json:"paymentList" description:"Payment list"`
}

type PaymentList struct {
	Period          int             `json:"period" description:"Period"`
	RepayDate       string          `json:"repayDate" description:"Repay date"`
	PlanRepayAmount decimal.Decimal `json:"PlanRepayAmount" description:"Plan repay amount"`
	ActualPayAmount decimal.Decimal `json:"actualPayAmount" description:"Actual pay amount"`
	RemainPrincipal decimal.Decimal `json:"remainPrincipal" description:"Remain principal"`
	RemainInterest  decimal.Decimal `json:"remainInterest" description:"Remain interest"`
}

func (*AL000060I) GetServiceKey() string {
	return "AL000060"
}
