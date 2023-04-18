package models

import "github.com/shopspring/decimal"

type AL000063I struct {
	Id        string `json:"id" description:"Id" validate:"required"`
	HashValue string `json:"hashValue" description:"Hash value" validate:"required"`
}

type AL000063O struct {
	PaymentResult      string          `json:"paymentResult" description:"Payment result:F-Failure S-Succeed U-unknow"`
	LinkStatus         string          `json:"linkStatus" description:"Link status:1-valid 2-invalid"`
	ErrorMsg           string          `json:"errorMsg" description:"Error message:If failure return the error message"`
	LoanReceiptNum     string          `json:"loanReceiptNum" description:"Loan receipt number"`
	OutStandingBalance decimal.Decimal `json:"outStandingBalance" description:"Out standing balance"`
	Principal          decimal.Decimal `json:"principal" description:"Principal"`
	IntAmount          decimal.Decimal `json:"intAmount" description:"Int amount"`
}

func (*AL000063I) GetServiceKey() string {
	return "AL000063"
}
