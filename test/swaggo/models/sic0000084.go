//Version: v1
package models

import (
	"github.com/shopspring/decimal"
)

type IC000084Request struct {
	CustomerId                  string          `json:"customerId"`
	LoanReceiptNum              string          `json:"loanReceiptNum" validate:"required"`
	RestructuringMethodList     string          `json:"restructuringMethodList" validate:"required" description:"1-Partial Write-off;2-Refinance;3-Adjust Interest Rate;4-Period Distribution;5-Maturity Date Extension;6-Bill Setting"`
	PrincipalWrittenOff         decimal.Decimal `json:"principalWrittenOff"`
	RefinancingPrincipal        decimal.Decimal `json:"refinancingPrincipal"`
	RestructuredLoanAmount      decimal.Decimal `json:"restructuredLoanAmount"`
	RestructuredPrincipalAmount decimal.Decimal `json:"restructuredPrincipalAmount"`
	NeedRepayTotalAmount        decimal.Decimal `json:"needRepayTotalAmount" validate:"required"`
	SourceChannle               string          `json:"sourceChannle"`
	MakerComment                string          `json:"makerComment"`
	FrontEndFee                 decimal.Decimal `json:"frontEndFee"`
	StampDutyFee                decimal.Decimal `json:"stampDutyFee"`
	ChangeContextList           []ChangeContext `json:"changeContextList"`
}
type ChangeContext struct {
	ChangeElement   string `json:"changeElement"`
	ElementOldValue string `json:"elementOldValue"`
	ElementNewValue string `json:"elementNewValue"`
}
type IC000084Response struct {
	RestructuringApplySerialNum string `json:"restructuringApplySerialNum"`
	CustomerId                  string `json:"customerId"`
	LoanReceiptNum              string `json:"loanReceiptNum"`
}
