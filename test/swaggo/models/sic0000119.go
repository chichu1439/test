//Version: v1
package models

import "github.com/shopspring/decimal"

type IC000119Request struct {
	CustomerId      string          `json:"customerId" validate:"required"`
	LoanReceiptNum  string          `json:"loanReceiptNum" validate:"required"`
	RefinancePeriod int             `json:"refinancePeriod" validate:"required"`
	RefinanceAmount decimal.Decimal `json:"refinanceAmount" validate:"required"`
	LoanPurpose     string          `json:"loanPurpose" validate:"required"`
	MatureDate      string          `json:"matureDate" validate:"required"`
	FrontEndFee     decimal.Decimal `json:"frontEndFee"`
	StampDutyFee    decimal.Decimal `json:"stampDutyFee"`
	MakerComment    string          `json:"makerComment" validate:"required"`
}

type IC000119Response struct {
	ApplySerialNum string `json:"applySerialNum"`
	LoanReceiptNum string `json:"loanReceiptNum"`
}

func (*IC000119Request) GetServiceKey() string {
	return "IC000119"
}
