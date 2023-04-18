//Version: v1.0.0.0
package models

import "github.com/shopspring/decimal"

type SV000034I struct {
	AgreementId          string `json:"agreementId" validate:"required"`
	AgreementType        string `json:"agreementType" validate:"required"`
	FreezeBusinessNumber string `json:"freezeBusinessNumber" validate:"required"`
	PageNo               int    `json:"pageNo"`
	PageRecCount         int    `json:"pageRecCount"`
}

type SV000034O struct {
	TotalRecord int              `json:"totalRecord"`
	TotalPage   int              `json:"totalPage"`
	Records     []FrozenRelation `json:"records"`
}

type FrozenRelation struct {
	BusinessNumber string          `json:"businessNumber"`
	DocNumber      string          `json:"docNumber"`
	EffectiveDate  string          `json:"effectiveDate"`
	ControlType    string          `json:"controlType"`
	OrganizeType   string          `json:"organizeType"`
	Amount         decimal.Decimal `json:"amount"`
}

func (*SV000034I) GetServiceKey() string {
	return "SV000034"
}
