package models

import "github.com/shopspring/decimal"

type PF000006I struct {
	SystemId     string   `json:"systemId" validate:"required" description:"System Id(LN-icredit;SV-isaving)"`
	StrategyName string   `json:"strategyName" description:"Strategy name"`
	Currency     string   `json:"currency" description:"Currency(THB)"`
	ListStrategy []string `json:"listStrategy" description:"List strategy" `
}

type PF000006O struct {
	ListStrategy []FeeStrategy `json:"listStrategy" description:"List strategy"`
}

type FeeStrategy struct {
	SystemId             string            `json:"systemId" validate:"required" description:"System Id(LN-icredit;SV-isaving)"`
	StrategyId           string            `json:"strategyId" description:"Strategy id" validate:"required"`
	StrategyName         string            `json:"strategyName" description:"Strategy name"`
	Currency             string            `json:"currency" description:"Currency"`
	MaxDiscount          decimal.Decimal   `json:"maxDiscount" description:"Max discount"`
	MaxFee               decimal.Decimal   `json:"maxFee" description:"Max fee"`
	MinDiscount          decimal.Decimal   `json:"minDiscount" description:"Min discount"`
	MinFee               decimal.Decimal   `json:"minFee" description:"Min fee"`
	MultipleCalcType     string            `json:"multipleCalcType" description:"multiple Calc Type：0-any;1-all"`
	MultipleDiscountType string            `json:"multipleDiscountType" description:"multiple Discount Type：0-Min Discount;1-Max Discount;2-Sum"`
	MultipleMatchType    string            `json:"multipleMatchType" description:"multiple Match Type：0-Min Fee;1-Max Fee;2-Sum"`
	EffectiveDate        string            `json:"effectiveDate" description:"Effective date"`
	ExpireDate           string            `json:"expireDate" description:"Expire date"`
	Status               string            `json:"status" description:"0-Normal"`
	ListFeeParmMatch     []FeeParamMatch   `json:"listFeeParmMatch" description:"List fee parm match"`
	ListFeeParmCalc      []FeeParmCalc     `json:"listFeeParmCalc" description:"List fee parm calc"`
	ListFeeParmDiscount  []FeeParmDiscount `json:"listFeeParmDiscount" description:"List fee parm discount"`
}

func (*PF000006I) GetServiceKey() string {
	return "PF000006"
}
