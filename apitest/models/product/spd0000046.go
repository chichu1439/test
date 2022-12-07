package models

import "github.com/shopspring/decimal"

type SPD0000046I struct {
	ProductId string `json:"productId" description:"Product id" validate:"required"`
}

type SPD0000046O struct {
	ListFee []FeeLoanInfo02 `json:"listFee" description:"List fee"`
}

type FeeLoanInfo02 struct {
	FeeType               string               `json:"feeType" description:"fee Type:0-account opening"` //0-account opening
	MaxFee                decimal.Decimal      `json:"maxFee" description:"Max fee"`
	MinFee                decimal.Decimal      `json:"minFee" description:"Min fee"`
	DiscountMethod        string               `json:"discountMethod" description:"discount Method:0-By Each Fee 1-By Total Fee"`               //0-By Each Fee 1-By Total Fee
	DiscountGroupType     string               `json:"discountGroupType" description:"discount Group Type:0-Min Discount 1-Max Discount 2-Sum"` //0-Min Discount 1-Max Discount 2-Sum
	DiscountCalculateType string               `json:"discountCalculateType" description:"discount Calculate Type:0-rate 1-fixed amount"`       //0-rate 1-fixed amount
	MaxDiscount           decimal.Decimal      `json:"maxDiscount" description:"Max discount"`
	MinDiscount           decimal.Decimal      `json:"minDiscount" description:"Min discount"`
	Status                string               `json:"status" description:"status:0-Normal"`
	ListStrategy          []StrategyLoanInfo02 `json:"listStrategy" description:"List strategy"`
}

type StrategyLoanInfo02 struct {
	FeeStrategy string `json:"feeStrategy" description:"Fee strategy"`
	Status      string `json:"status" description:"0-Normal"`
}

func (*SPD0000046I) GetServiceKey() string {
	return "spd0000046"
}
