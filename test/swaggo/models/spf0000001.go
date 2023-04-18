package models

import "github.com/shopspring/decimal"

type PF000001I struct {
	SystemId      string          `json:"systemId"`
	Currency      string          `json:"currency"`
	MaxFeeAmount  decimal.Decimal `json:"maxFeeAmount"`
	MinFeeAmount  decimal.Decimal `json:"minFeeAmount"`
	FeeCalcMethod string          `json:"feeCalcMethod"`
	Remark        string          `json:"remark"`
	FeeTypeId     string          `json:"feeTypeId"`
	FeeName       string          `json:"feeName"`
	CalcBasisType string          `json:"calcBasisType"`
	PayerType     string          `json:"payerType"`
	FeeRules      []FeeRulesItem  `json:"feeRules"`
}

type FeeRulesItem struct {
	ThresholdFlag  string          `json:"thresholdFlag"`
	ThresholdType  string          `json:"thresholdType"`
	UpperCondition string          `json:"upperCondition"`
	LowerCondition string          `json:"lowerCondition"`
	UpperInterval  decimal.Decimal `json:"upperInterval"`
	LowerInterval  decimal.Decimal `json:"lowerInterval"`
	CalcValue      decimal.Decimal `json:"calcValue"`
}
type PF000001O struct {
	FeeId   string `json:"feeId"`
	FeeName string `json:"feeName"`
}

func (*PF000001I) GetServiceKey() string {
	return "PF000001"
}
