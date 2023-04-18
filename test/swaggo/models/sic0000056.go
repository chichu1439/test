package models

import "github.com/shopspring/decimal"

type IC000056I struct {
	LoanReceiptNum string           `validate:"required,max=32" json:"loanReceiptNum" description:"Loan receipt number"`
	CalcObj        decimal.Decimal  `validate:"max=(18,2)" json:"calcObj" description:"Calc obj"`
	ListFeeType    []LoanFeeInfoReq `json:"loanFeeInfo" description:"List fee type"`
}

type IC000056O struct {
	ListFeeType []LoanFeeInfoRsp `json:"listFeeType" description:"List fee type"`
}

type LoanFeeInfoReq struct {
	FeeType string `validate:"required,max=2" json:"feeType" description:"Fee type"`
}

type LoanFeeInfoRsp struct {
	FeeAmount     decimal.Decimal `validate:"max=(18,2)" json:"feeAmount" description:"Fee amount"`
	FeeType       string          `validate:"max=2" json:"feeType" description:"Fee type"`
	MaxFee        decimal.Decimal `validate:"max=(10,2)" json:"maxFee" description:"Max fee"`
	MinFee        decimal.Decimal `validate:"max=(10,2)" json:"minFee" description:"Min fee"`
	FeeCalcMethod string          `validate:"max=1" json:"feeCalcMethod" description:"Fee calc method"`
	CalcValue     decimal.Decimal `validate:"max=(10,2)" json:"calcValue" description:"Calc value"`
	PayerType     string          ` validate:"max=1" json:"payerType" description:"Payer type"`
}

func (*IC000056I) GetServiceKey() string {
	return "IC000056"
}
