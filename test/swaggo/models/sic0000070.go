package models

import "github.com/shopspring/decimal"

type IC000070Request struct {
	LoanReceiptNum  string          `json:"loanReceiptNum" description:"Loan Receipt Number" validate:"required"`
	TransAmount     decimal.Decimal `json:"transAmount" description:"Transaction amount" validate:"required"`
	LoanChannel     string          `json:"loanChannel" description:"Loan channel"  validate:"required"`
	MakerUserId     string          `json:"makerUserId" description:"Maker User Id"`
	MakerComment    string          `json:"makerComment" description:"Maker Comment"`
	RepaymentStatus string          `json:"repaymentStatus" description:"Repayment Status 01-total 02-Partial 03-by bill 04-down payment 05- overdue payment"`
	TransSerialNum  string          `json:"transSerialNum" description:"Transaction Serial Number"`
	BankNum         string          `json:"bankNum"`
	BankName        string          `json:"bankName"`
	AccountNum      string          `json:"accountNum"`
	AccountName     string          `json:"accountName"`
}

type IC000070Response struct {
}

func (*IC000070Request) GetServiceKey() string {
	return "IC000070"
}
