package models

import "github.com/shopspring/decimal"

type AL000056I struct {
	TransDate            string              `json:"transDate" description:"transaction date" validate:"required" `
	LoanReceiptNum       string              `json:"loanReceiptNum" description:"Loan receipt number" vaidate:"required"`
	AccountingAccountNum string              `json:"accountingAccountNum" description:"Accounting account number" validate:"required"`
	Currency             string              `json:"currency" description:"Currency"validate:"required" description:"Currency" `
	LoanAmount           decimal.Decimal     `json:"loanAmount" description:"Loan amount" validate:"required"`
	MatureDate           string              `json:"matureDate" description:"Mature date" `
	LoanPeriod           int                 `json:"loanPeriod" description:"Loan period" validate:"required"`
	RepaymentMethod      string              `json:"repaymentMethod" description:"Repayment method:01-Fixed Installment;02-Fixed Principal" validate:"required"`
	RepaymentCycle       string              `json:"repaymentCycle" description:"Repayment cycle"`
	RepayDay             string              `json:"repayDay" description:"Repay day"`
	InterestRate         decimal.Decimal     `json:"interestRate" description:"Interest rate"`
	ListInterest         []MicroLoanInterest `json:"listInterest" description:"List interest"`
}

type AL000056O struct {
}

func (*AL000056I) GetServiceKey() string {
	return "AL000056"
}
