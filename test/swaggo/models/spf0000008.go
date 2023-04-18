package models

import "github.com/shopspring/decimal"

type PF000008I struct {
	SystemId             string            `json:"systemId" validate:"required" description:"System Id(LN-icredit;SV-isaving)"`
	StrategyId           string            `json:"strategyId" description:"Strategy id" validate:"required"`
	StrategyName         string            `json:"strategyName" description:"Strategy name"`
	Currency             string            `json:"currency" description:"Currency(THB)"`
	MaxDiscount          decimal.Decimal   `json:"maxDiscount" description:"Max discount"`
	MaxFee               decimal.Decimal   `json:"maxFee" description:"Max fee"`
	MinDiscount          decimal.Decimal   `json:"minDiscount" description:"Min discount"`
	MinFee               decimal.Decimal   `json:"minFee" description:"Min fee"`
	MultipleCalcType     string            `json:"multipleCalcType" description:"multiple Calc Type：0-any;1-all"`
	MultipleDiscountType string            `json:"multipleDiscountType" description:"multiple Discount Type：0-Min Discount;1-Max Discount;2-Sum"`
	MultipleMatchType    string            `json:"multipleMatchType" description:"multiple Match Type：0-Min Fee;1-Max Fee;2-Sum"`
	EffectiveDate        string            `json:"effectiveDate" description:"Effective date"`
	ExpireDate           string            `json:"expireDate" description:"Expire date"`
	ListParmMatch        []FeeParamMatch   `json:"listParmMatch" description:"List parm match"` //
	ListParmCalc         []FeeParmCalc     `json:"listParmCalc" description:"List parm calc"`
	ListParmDiscount     []FeeParmDiscount `json:"listParmDiscount" description:"List parm discount"`
}

type PF000008O struct {
}

func (*PF000008I) GetServiceKey() string {
	return "PF000008"
}
