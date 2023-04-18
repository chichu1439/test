//Version: v1
package models

import "github.com/shopspring/decimal"

type PF000011Request struct {
	SystemId          string          `json:"systemId"`
	FeeType           string          `json:"feeType" description:"5-Transfer Fee"  validate:"required"`
	Currency          string          `json:"currency" description:"Currency(THB)"  validate:"required"`
	FeeCondition      decimal.Decimal `json:"feeCondition"`
	DiscountCondition decimal.Decimal `json:"discountCondition"`
	ProductId         string          `json:"productId"`
}

type PF000011Response struct {
	FeeId         string          `json:"feeId"`
	ProductId     string          `json:"productId"`
	Currency      string          `json:"currency" description:"Currency(THB)"`
	FeeType       string          `json:"feeType"`
	MinFeeAmount  decimal.Decimal `json:"minFeeAmount"`
	MaxFeeAmount  decimal.Decimal `json:"maxFeeAmount"`
	CalcValue     decimal.Decimal `json:"calcValue"`
	DiscountValue decimal.Decimal `json:"discountValue"`
}

func (*PF000011Request) GetServiceKey() string {
	return "PF000011"
}
