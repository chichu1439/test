package models

import "github.com/shopspring/decimal"

type IC000068Request struct {
	CustomerId           string          `json:"customerId" validate:"required"`
	ProductId            string          `json:"productId" validate:"required"`
	CreditApplySerialNum string          `json:"creditApplySerialNum" validate:"required"`
	LoanAmount           decimal.Decimal `json:"loanAmount"`
}

type IC000068Response struct {
	Records []IC000068ResponseFee `json:"records"`
}
type IC000068ResponseFee struct {
	FeeType   string          `json:"feeType"`
	FeeAmount decimal.Decimal `json:"feeAmount"`
}
