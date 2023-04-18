package models

import "github.com/shopspring/decimal"

type PD000084I struct {
}

type PD000084O struct {
	ListProduct []ProductFeeInfo `json:"listProduct" description:"List product"`
}

type ProductFeeInfo struct {
	ProductId string             `json:"productId"`
	ListFee   []PD000084OFeeItem `json:"listFee"`
}
type PD000084OFeeItem struct {
	FeeId         string                  `json:"feeId"`
	FeeType       string                  `json:"feeType"`
	CalcBasisType string                  `json:"calcBasisType"`
	PayerType     string                  `json:"payerType"`
	MaxDiscount   decimal.Decimal         `json:"maxDiscount"`
	MinDiscount   decimal.Decimal         `json:"minDiscount"`
	EffectiveDate string                  `json:"effectiveDate"`
	ExpireDate    string                  `json:"expireDate"`
	Status        string                  `json:"statusstatus"`
	RemarkInfo    string                  `json:"remarkInfo"`
	ChangeType    string                  `json:"changeType"`
	SystemId      string                  `json:"systemId"`
	Currency      string                  `json:"currency"`
	MaxFeeAmount  decimal.Decimal         `json:"maxFeeAmount"`
	MinFeeAmount  decimal.Decimal         `json:"minFeeAmount"`
	FeeCalcMethod string                  `json:"feeCalcMethod"`
	FeeRules      []PD000084OFeeRulesItem `json:"feeRules"`
}
type PD000084OFeeRulesItem struct {
	IntervalNo     int             `json:"intervalNo"`
	ThresholdFlag  string          `json:"thresholdFlag"`
	ThresholdType  string          `json:"thresholdType"`
	UpperCondition string          `json:"upperCondition"`
	LowerCondition string          `json:"lowerCondition"`
	UpperInterval  decimal.Decimal `json:"upperInterval"`
	LowerInterval  decimal.Decimal `json:"lowerInterval"`
	CalcValue      decimal.Decimal `json:"calcValue"`
}

func (*PD000084I) GetServiceKey() string {
	return "PD000084"
}
