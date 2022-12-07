package models

import "github.com/shopspring/decimal"

type IC000072Request struct {
	CustomerId      string `json:"customerId" description:"Customer id" validate:"required"`
	ProductId       string `json:"productId" description:"Product id" validate:"required"`
	PageNum         int    `json:"pageNum" description:"Page number"`
	PageRecordCount int    `json:"pageRecordCount" description:"Page record count"`
}

type IC000072Response struct {
	PageNum         int                `json:"pageNum" description:"Page number"`
	PageRecordCount int                `json:"pageRecordCount" description:"Page record count"`
	TotalCount      int                `json:"totalCount" description:"Total count"`
	Records         []LoanContractList `json:"records" description:"Records"`
}

type LoanContractList struct {
	LoanReceiptNum  string          `json:"loanReceiptNum" description:"Loan receipt number"`
	LoanStatus      string          `json:"loanStatus" description:"Loan status"`
	Currency        string          `json:"currency" description:"Currencuy"`
	LoanAmount      decimal.Decimal `json:"loanAmount" description:"Loan amount"`
	OpenAccountDate string          `json:"openAccountDate" description:"Open account date"`
	TotalPeriod     int             `json:"totalPeriod" description:"Total Period"`
	PaidPeriod      int             `json:"paidPeriod" description:"Paid Period"`
}

func (*IC000072Request) GetServiceKey() string {
	return "IC000072"
}
