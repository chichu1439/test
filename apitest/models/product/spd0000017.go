package models

import "github.com/shopspring/decimal"

type SPD0000017I struct {
	ProductId string `json:"productId" description:"Product Id"`
}

type SPD0000017O struct {
	ListFee []FeeSavingScenario `json:"listFee" description:"List fee"`
}

type FeeSavingScenario struct {
	FeeType               string          `json:"feeType" description:"fee Type:0-account opening"`
	MaxFee                decimal.Decimal `json:"maxFee" description:"Max fee"`
	MinFee                decimal.Decimal `json:"minFee" description:"Min fee"`
	DiscountMethod        string          `json:"discountMethod" description:"discount Method:0-By Each Fee;1-By Total Fee"`
	DiscountGroupType     string          `json:"discountGroupType" description:"discount Group Type:0-Min Discount;1-Max Discount;2-Sum"`
	DiscountCalculateType string          `json:"discountCalculateType" description:"discount Calculate Type:0-By Percent;1-Fixed Amount"`
	MaxDiscount           decimal.Decimal `json:"maxDiscount" description:"Max discount"`
	MinDiscount           decimal.Decimal `json:"minDiscount" description:"Min discount"`
	ListStrategy          []FeeSaving     `json:"listStrategy" description:"List strategy"`
}

type FeeSaving struct {
	StrategyId string `json:"strategyId" description:"Strategy id"`
}

func (*SPD0000017I) GetServiceKey() string {
	return "spd0000017"
}
