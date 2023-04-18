package models

import "github.com/shopspring/decimal"

type PF000003I struct {
	FeeId         string               `json:"feeId"`
	SystemId      string               `json:"systemId"`
	Currency      string               `json:"currency"`
	MaxFeeAmount  decimal.Decimal      `json:"maxFeeAmount"`
	MinFeeAmount  decimal.Decimal      `json:"minFeeAmount"`
	FeeCalcMethod string               `json:"feeCalcMethod"`
	Remark        string               `json:"remark"`
	FeeRules      []PF000003IRulesItem `json:"feeRules"`
	FeeName       string               `json:"feeName"`
	FeeTypeId     string               `json:"feeTypeId"`
	CalcBasisType string               `json:"calcBasisType"`
	PayerType     string               `json:"payerType"`
}

type PF000003IRulesItem struct {
	IntervalNo     int             `json:"intervalNo"`
	ThresholdFlag  string          `json:"thresholdFlag"`
	ThresholdType  string          `json:"thresholdType"`
	UpperCondition string          `json:"upperCondition"`
	LowerCondition string          `json:"lowerCondition"`
	UpperInterval  decimal.Decimal `json:"upperInterval"`
	LowerInterval  decimal.Decimal `json:"lowerInterval"`
	CalcValue      decimal.Decimal `json:"calcValue"`
}
type PF000003O struct {
	FeeId string `json:"feeId"`
}

func (*PF000003I) GetServiceKey() string {
	return "PF000003"
}
