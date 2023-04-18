package models

import "github.com/shopspring/decimal"

type PD000082I struct {
	StrDate         string `json:"strDate" description:"Start date"`
	EndDate         string `json:"endDate" description:"End date"`
	ProductId       string `json:"productId" description:"Product Id"`
	PageNum         int    `json:"pageNum" description:"Page number"`
	PageRecordCount int    `json:"pageRecordCount" description:"Page record count"`
}

type PD000082O struct {
	PageNum         int                   `json:"pageNum" description:"Page number"`
	PageRecordCount int                   `json:"pageRecordCount" description:"Page record count"`
	TotalCount      int                   `json:"totalCount" description:"Total count"`
	Records         []ProdLimitControlHis `json:"records" description:"Records"`
}

type ProdLimitControlHis struct {
	Uid               int             `json:"uid" description:"Uid"`
	ProductId         string          `json:"productId" description:"Product id"`
	ServiceId         string          `json:"serviceId" description:"Service id"`
	FeatureId         string          `json:"featureId" description:"Feature id"`
	LimitTarget       string          `json:"limitTarget" description:"limit Target:01-Balance;02-Trans Times;03-Trans Amount;04-Free Fee Times;05-Free Fee Amount;06-Free Interest Times;07-Free Interest Amount;08-Free Tax Times;09-Free Tax Amount;10-Grace Period"`
	LimitMethod       string          `json:"limitMethod" description:"limit Method:0-Direct Limit"`
	ControlMethod     string          `json:"controlMethod" description:"Control method"`
	UpperLimitControl string          `json:"upperLimitControl" description:"upper Limit Control:0-Reject;1-Prompt;2-Record Only"`
	UpperLimitValue   decimal.Decimal `json:"upperLimitValue" description:"Upper limit value"`
	LowerLimitControl string          `json:"lowerLimitControl" description:"lower Limit Control:0-Reject;1-Prompt;2-Record Only"`
	LowerLimitValue   decimal.Decimal `json:"lowerLimitValue" description:"Lower limit value"`
	LimitStrategyId   string          `json:"limitStrategyId" description:"Limit strategy id"`
	Status            string          `json:"status" description:"status:0-inactive;1-active;2-expired;3-deleted;4-updated"`
	ModifySeq         string          `json:"modifySeq" description:"Modify seq"`
	EffectiveDate     string          `json:"effectiveDate" description:"Effective date"`
	ExpireDate        string          `json:"expireDate" description:"Expire date"`
	LastModifyDate    string          `json:"lastModifyDate" description:"Last modify date"`
	LastModifyTime    string          `json:"lastModifyTime" description:"Last modify time"`
	LastModifyOrg     string          `json:"lastModifyOrg" description:"Last modify org"`
	LastModifyUser    string          `json:"lastModifyUser" description:"Last modify user"`
}

func (*PD000082I) GetServiceKey() string {
	return "PD000082"
}
