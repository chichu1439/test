package models

import "github.com/shopspring/decimal"

type PD000080I struct {
	StrDate         string `json:"strDate" description:"Start date"`
	EndDate         string `json:"endDate" description:"End date"`
	ProductId       string `json:"productId" description:"Product Id"`
	PageNum         int    `json:"pageNum" description:"Page number"`
	PageRecordCount int    `json:"pageRecordCount" description:"Page record count"`
}

type PD000080O struct {
	HistoryRecords  HistoryRecordsTemp  `json:"historyRecords"`
	PlanningRecords HistoryRecordsTemp2 `json:"planningRecords"`
}
type HistoryRecordsTemp struct {
	PageNum         int             `json:"pageNum" description:"Page number"`
	PageRecordCount int             `json:"pageRecordCount" description:"Page record count"`
	TotalCount      int             `json:"totalCount" description:"Total count"`
	Records         []ProdFeeItem02 `json:"records" description:"Records"`
}
type HistoryRecordsTemp2 struct {
	TotalCount int             `json:"totalCount" description:"Total count"`
	Records    []ProdFeeItem02 `json:"records" description:"Records"`
}
type ProdFeeItem02 struct {
	FeeType       string            `json:"feeType"`
	CalcBasisType string            `json:"calcBasisType"`
	PayerType     string            `json:"payerType"`
	RemarkInfo    string            `json:"remarkInfo"`
	ChangeType    string            `json:"changeType"`
	MaxDiscount   decimal.Decimal   `json:"maxDiscount"`
	MinDiscount   decimal.Decimal   `json:"minDiscount"`
	Status        string            `json:"status"`
	FeeId         string            `json:"feeId"`
	EffectiveDate string            `json:"effectiveDate"`
	ExpireDate    string            `json:"expireDate"`
	SystemId      string            `json:"systemId"`
	Currency      string            `json:"currency"`
	MaxFeeAmount  decimal.Decimal   `json:"maxFeeAmount"`
	MinFeeAmount  decimal.Decimal   `json:"minFeeAmount"`
	FeeCalcMethod string            `json:"feeCalcMethod"`
	FeeRules      []ProdFeeRulesHis `json:"feeRules"`
	Uid           int               `json:"uid"`
	RecordTime    string            `json:"recordTime"`
	RecordDate    string            `json:"recordDate"`
}
type ProdFeeRulesHis struct {
	Uid            int             `json:"uid"`
	IntervalNo     int             `json:"intervalNo"`
	UpperInterval  decimal.Decimal `json:"upperInterval"`
	LowerInterval  decimal.Decimal `json:"lowerInterval"`
	UpperCondition string          `json:"upperCondition"`
	LowerCondition string          `json:"lowerCondition"`
	CalcValue      decimal.Decimal `json:"calcValue"`
	ThresholdFlag  string          `json:"thresholdFlag"`
	ThresholdType  string          `json:"thresholdType"`
}

func (*PD000080I) GetServiceKey() string {
	return "PD000080"
}
