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
	PageNum         int             `json:"pageNum" description:"Page number"`
	PageRecordCount int             `json:"pageRecordCount" description:"Page record count"`
	TotalCount      int             `json:"totalCount" description:"Total count"`
	Records         []ProdFeeItem02 `json:"records" description:"Records"`
}

type ProdFeeService02 struct {
	Uid         int             `json:"uid" description:"Uid"`
	ServiceId   string          `json:"serviceId" description:"Service id"`
	FeatureId   string          `json:"featureId" description:"Feature id"`
	ListFeeType []ProdFeeItem02 `json:"listFeeType" description:"List fee type"`
}

type ProdFeeItem02 struct {
	Uid             int             `json:"uid" description:"Uid"`
	ServiceId       string          `json:"serviceId" description:"Service id"`
	FeatureId       string          `json:"featureId" description:"Feature id"`
	FeeId           string          `json:"feeId" description:"Fee id"`
	FeeName         string          `json:"feeName" description:"Fee name"`
	MaxFee          decimal.Decimal `json:"maxFee" description:"Max fee"`
	MinFee          decimal.Decimal `json:"minFee" description:"Min fee"`
	EffectiveDate   string          `json:"effectiveDate" description:"Effective date"`
	ExpireDate      string          `json:"expireDate" description:"Expire date"`
	Status          string          `json:"status" description:"status:0-inactive;1-active;2-expired;3-deleted;4-updated"`
	ModifySeq       string          `json:"modifySeq" description:"Modify seq"`
	FinalModifyDate string          `json:"finalModifyDate" description:"Final modify date"`
	FinalModifyTime string          `json:"finalModifyTime" description:"Final modify time"`
	FinalModifyOrgz string          `json:"finalModifyOrgz" description:"Final modify organization"`
	FinalModifyUser string          `json:"finalModifyUser" description:"Final modify user"`
	FeeCalcMethod   string          `json:"feeCalcMethod" description:"feeCalcMethod:0-By Percent;1-Fixed Amount""`
	CalcValue       decimal.Decimal `json:"calcValue" description:"Calc value"`
	PayerType       string          `json:"payerType" description:"payerType:01-Payee Participant;02-Payer Participant"`
}

func (*PD000080I) GetServiceKey() string {
	return "PD000080"
}
