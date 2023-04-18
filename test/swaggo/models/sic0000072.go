package models

import "github.com/shopspring/decimal"

type IC000072Request struct {
	CustomerId      string `json:"customerId" description:"Customer id"`
	ProductId       string `json:"productId" description:"Product id"`
	LoanReceiptNum  string `json:"loanReceiptNum" description:"Loan Receipt Num"`
	LoanStatus      string `json:"loanStatus"`
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
	Balance         decimal.Decimal `json:"balance" description:"Balance"`
	InterestRate    decimal.Decimal `json:"interestRate" description:"Interest Rate"`
	ProductId       string          `json:"productId" description:"product Id"`
	CustomerId      string          `json:"customerId" description:"Customer id"`
	DrawdownDate    string          `json:"drawdownDate" description:"Drawdown Date"`
	MaturityDate    string          `json:"maturityDate" description:"Maturity Date"`
	RepayMethod     string          `json:"repayMethod" description:"Repay Method"`
	RepayCycle      string          `json:"repayCycle" description:"Repay Cycle"`
}

type LoanAccountList struct {
	LoanReceiptNum  string          `json:"loanReceiptNum" description:"Loan receipt number"`
	LoanStatus      string          `json:"loanStatus" description:"Loan status"`
	Currency        string          `json:"currency" description:"Currencuy"`
	LoanAmount      decimal.Decimal `json:"loanAmount" description:"Loan amount"`
	TotalPeriod     int             `json:"totalPeriod" description:"Total Period"`
	PaidPeriod      int             `json:"paidPeriod" description:"Paid Period"`
	OpenAccountDate string          `json:"openAccountDate" description:"Open account date"`
	Balance         decimal.Decimal `json:"balance" description:"Balance"`
	MaturityDate    string          `json:"maturityDate" description:"Maturity Date"`
	InterestRate    decimal.Decimal `json:"interestRate" description:"Interest Rate"`
	ProductId       string          `json:"productId" description:"product Id"`
	CustomerId      string          `json:"customerId" description:"Customer id"`
}

func (*IC000072Request) GetServiceKey() string {
	return "IC000072"
}
