package models

import "github.com/shopspring/decimal"

type PF000005I struct {
	SystemId             string            `json:"systemId" validate:"required" description:"System Id(LN-icredit;SV-isaving)"`
	StrategyName         string            `json:"strategyName" validate:"required" description:"Strategy name"`
	Currency             string            `json:"currency" validate:"required" description:"Currency(THB)"`
	MaxDiscount          decimal.Decimal   `json:"maxDiscount" validate:"required" description:"Max discount"`
	MaxFee               decimal.Decimal   `json:"maxFee" validate:"required" description:"Max fee"`
	MinDiscount          decimal.Decimal   `json:"minDiscount" validate:"required" description:"Min discount"`
	MinFee               decimal.Decimal   `json:"minFee" validate:"required" description:"Min fee"`
	MultipleCalcType     string            `json:"multipleCalcType" validate:"required" description:"Multiple Calculate Type(0-Min Discount;1-Max Discount;2-Sum)"`
	MultipleDiscountType string            `json:"multipleDiscountType" validate:"required" description:"Min fee(0-Min Fee;1-Max Fee;2-Sum)"`
	MultipleMatchType    string            `json:"multipleMatchType" validate:"required" description:"multiple Match Type(0-any;1-all)"`
	EffectiveDate        string            `json:"effectiveDate" description:"Effective date"`
	ExpireDate           string            `json:"expireDate" description:"Expire date"`
	ListParmMatch        []FeeParamMatch   `json:"listParmMatch" description:"List parm match"` //
	ListParmCalc         []FeeParmCalc     `json:"listParmCalc" description:"List parm calc"`
	ListParmDiscount     []FeeParmDiscount `json:"listParmDiscount" description:"List parm discount"`
}

type FeeParamMatch struct {
	ParmCode           string           `json:"parmCode" description:"Parm code"`
	ParmValue          string           `json:"parmValue" description:"Parm value"`
	ParmMatchType      string           `json:"parmMatchType" description:"parm Match Type：0-Greater than 1-Greater than or equal to 2-Less than 3-Less than or equal to 4-equal 5-not equal 6-range"` //0-Greater than 1-Greater than or equal to 2-Less than 3-Less than or equal to 4-equal 5-not equal 6-range
	RangeCalculateType string           `json:"rangeCalculateType" description:"range Calculate Type：(0-exclude end;1-exclude start;2-include both)"`
	ListRange          []PriceParmRange `json:"listRange" description:"List range"`
}

type FeeParmCalc struct {
	CalcType           string           `json:"calcType" description:"calc Type：0-By Percent;1-Fixed Amount"`
	CalcValue          decimal.Decimal  `json:"calcValue" description:"Calc value"`
	GroupMatchType     string           `json:"groupMatchType" description:"group Match Type：0-any;1-all"`
	ParmCode           string           `json:"parmCode" description:"Parm code"`
	ParmMatchType      string           `json:"parmMatchType" description:"parm Match Type：(0-Greater than;1-Greater than or equal to;2-Less than;3-Less than or equal to;4-equal;5-not equal;6-range)"`
	ParmType           string           `json:"parmType" description:"parm Type：0-single 1-group" ` //0-single 1-group
	ParmValue          string           `json:"parmValue" description:"Parm value"`
	RangeCalculateType string           `json:"rangeCalculateType" description:"range Calculate Type：0-exclude end;1-exclude start;2-include both"`
	ListRange          []PriceParmRange `json:"listRange" description:"List range"`
	ListGroup          []PriceParmGroup `json:"listGroup" description:"List group"`
}

type FeeParmDiscount struct {
	DiscountType       string           `json:"discountType" description:"discount Type：0-By Percent;1-Fixed Amount"`
	DiscountValue      decimal.Decimal  `json:"discountValue" description:"Discount value" `
	GroupMatchType     string           `json:"groupMatchType" description:"group Match Type：0-Greater than;1-Greater than or equal to;2-Less than;3-Less than or equal to;4-equal;5-not equal;6-range"`
	ParmCode           string           `json:"parmCode" description:"Parm code"`
	ParmMatchType      string           `json:"parmMatchType" description:"parm Match Type：0-any;1-all"`
	ParmType           string           `json:"parmType" description:"parm Type：0-single;1-group"`
	ParmValue          string           `json:"parmValue" description:"Parm value"`
	RangeCalculateType string           `json:"rangeCalculateType" description:"range Calculate Type：0-exclude end;1-exclude start;2-include both"`
	ListRange          []PriceParmRange `json:"listRange" description:"List range"`
	ListGroup          []PriceParmGroup `json:"listGroup" description:"List group"`
}

type PriceParmRange struct {
	ParmCode   string `json:"parmCode" description:"Parm code"`
	RangeStart string `json:"rangeStart" description:"Range start"`
	RangeEnd   string `json:"rangeEnd" description:"Range end"`
	Result     string `json:"result" description:"Result"`
	ResultType string `json:"resultType" description:"Result type"`
}

type PriceParmGroup struct {
	GroupName          string           `json:"groupName" description:"Group name"`
	ParmCode           string           `json:"parmCode" description:"Parm code"`
	ParmValue          string           `json:"parmValue" description:"Parm value"`
	ParmMatchType      string           `json:"parmMatchType" description:"parm Match Type：0-Greater than;1-Greater than or equal to;2-Less than;3-Less than or equal to;4-equal;5-not equal;6-range"`
	RangeCalculateType string           `json:"rangeCalculateType" description:"range Calculate Type：0-exclude end;1-exclude start;2-include both"`
	ListRange          []PriceParmRange `json:"listRange" description:"List range"`
}

type PF000005O struct {
	StrategyId string `json:"strategyId" description:"Strategy id"`
}

func (*PF000005I) GetServiceKey() string {
	return "PF000005"
}
