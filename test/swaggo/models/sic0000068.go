package models

import "github.com/shopspring/decimal"

type IC000068Request struct {
	ProductId       string          `json:"productId"`
	ContractAmount  decimal.Decimal `json:"contractAmount"`
	RefinanceAmount decimal.Decimal `json:"refinanceAmount"`
}

type IC000068Response struct {
	DisbursementAmount decimal.Decimal       `json:"disbursementAmount"`
	Records            []IC000068ResponseFee `json:"records"`
}
type IC000068ResponseFee struct {
	FeeType   string          `json:"feeType"`
	FeeAmount decimal.Decimal `json:"feeAmount"`
	OnTop     string          `json:"onTop"`
}
