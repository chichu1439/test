package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type PF000002I struct {
	ListFeeId       []string `json:"listFeeId"`
	FeeTypeName     string   `json:"feeTypeName"`
	FeeTypeId       string   `json:"feeTypeId"`
	FeeName         string   `json:"feeName"`
	PageNum         int      `json:"pageNum"`
	PageRecordCount int      `json:"pageRecordCount"`
}

type PF000002O struct {
	PageNum         int                `json:"pageNum" description:"Page number"`
	PageRecordCount int                `json:"pageRecordCount" description:"Page record count"`
	TotalCount      int                `json:"totalCount" description:"Total count"`
	Records         []PriceFeeItemList `json:"records" description:"Records"`
}

type PriceFeeItemList struct {
	FeeId           string                  `json:"feeId"`
	SystemId        string                  `json:"systemId"`
	Currency        string                  `json:"currency"`
	MaxFeeAmount    decimal.Decimal         `json:"maxFeeAmount"`
	MinFeeAmount    decimal.Decimal         `json:"minFeeAmount"`
	FeeCalcMethod   string                  `json:"feeCalcMethod"`
	Remark          string                  `json:"remark"`
	FeeTypeId       string                  `json:"feeTypeId"`
	FeeTypeName     string                  `json:"feeTypeName"`
	AccountingCode  string                  `json:"accountingCode"`
	IsAmortization  string                  `json:"isAmortization"`
	FeeName         string                  `json:"feeName"`
	CalcBasisType   string                  `json:"calcBasisType"`
	PayerType       string                  `json:"payerType"`
	FinalModifyDate time.Time               `json:"finalModifyDate"`
	FinalModifyOrgz string                  `json:"finalModifyOrgz"`
	FinalModifyTime time.Time               `json:"finalModifyTime"`
	FinalModifyUser string                  `json:"finalModifyUser"`
	FeeRules        []PF000002OFeeRulesItem `json:"feeRules"`
}
type PF000002OFeeRulesItem struct {
	IntervalNo     int             `json:"intervalNo"`
	ThresholdFlag  string          `json:"thresholdFlag"`
	ThresholdType  string          `json:"thresholdType"`
	UpperCondition string          `json:"upperCondition"`
	LowerCondition string          `json:"lowerCondition"`
	UpperInterval  decimal.Decimal `json:"upperInterval"`
	LowerInterval  decimal.Decimal `json:"lowerInterval"`
	CalcValue      decimal.Decimal `json:"calcValue"`
}

func (*PF000002I) GetServiceKey() string {
	return "PF000002"
}
