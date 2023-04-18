package models

import "github.com/shopspring/decimal"

type PF000023I struct {
	StrDate           string                `json:"strDate" description:"Start date"`
	EndDate           string                `json:"endDate" description:"End date"`
	PlanningFeeIdList map[int]PF000023IItem `json:"planningFeeIdList" description:"Fee id list"`
	HisFeeIdList      map[int]PF000023IItem `json:"hisFeeIdList" description:"Fee id list"`
}
type PF000023IItem struct {
	FeeId     string `json:"feeId"`
	ModifySeq string `json:"modifySeq"`
}
type PF000023O struct {
	ListProdFeeItemHis      map[int]ProdFeeItemHis `json:"prodFeeItemHis" description:"List prod fee item his"`
	ListProdFeeItemPlanning map[int]ProdFeeItemHis `json:"prodFeeItemPlanning" description:"List planning prod fee item "`
}

type ProdFeeItemHis struct {
	SystemId        string            `json:"systemId"`
	FeeId           string            `json:"feeId"`
	Currency        string            `json:"currency"`
	FeeCalcMethod   string            `json:"feeCalcMethod"`
	MaxFeeAmount    decimal.Decimal   `json:"maxFeeAmount"`
	MinFeeAmount    decimal.Decimal   `json:"minFeeAmount"`
	EffectiveDate   string            `json:"effectiveDate"`
	ExpireDate      string            `json:"expireDate"`
	Status          string            `json:"status"`
	FinalModifyDate string            `json:"finalModifyDate"`
	FinalModifyOrgz string            `json:"finalModifyOrgz"`
	FinalModifyTime string            `json:"finalModifyTime"`
	FinalModifyUser string            `json:"finalModifyUser"`
	Uid             int               `json:"uid"`
	FeeRules        []ProdFeeRulesHis `json:"feeRules"`
}
type ProdFeeRulesHis2 struct {
	Uid             int             `json:"uid"`
	FeeId           string          `json:"feeId"`
	IntervalNo      int             `json:"intervalNo"`
	UpperInterval   decimal.Decimal `json:"upperInterval"`
	LowerInterval   decimal.Decimal `json:"lowerInterval"`
	UpperCondition  string          `json:"upperCondition"`
	LowerCondition  string          `json:"lowerCondition"`
	CalcValue       decimal.Decimal `json:"calcValue"`
	ThresholdFlag   string          `json:"thresholdFlag"`
	ThresholdType   string          `json:"thresholdType"`
	FinalModifyDate string          `json:"finalModifyDate"`
	FinalModifyOrgz string          `json:"finalModifyOrgz"`
	FinalModifyTime string          `json:"finalModifyTime"`
	FinalModifyUser string          `json:"finalModifyUser"`
}

func (*PF000023I) GetServiceKey() string {
	return "PF000023"
}
