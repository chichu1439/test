package models

import "github.com/shopspring/decimal"

type PD000084I struct {
	SystemId string `json:"systemId" validate:"required" description:"system Id(LN-icredit;SV-isaving)"`
}

type PD000084O struct {
	ListProduct []ProductFeeInfo `json:"listProduct" description:"List product"`
}

type ProductFeeInfo struct {
	ProductId string           `json:"productId" description:"Product Id"`
	Currency  string           `json:"currency" description:"Currency(THB)"`
	ListFee   []FeeServiceInfo `json:"listFee" description:"List fee"`
}

type FeeServiceInfo struct {
	ServiceId   string        `json:"serviceId" description:"Service id"`
	FeatureId   string        `json:"featureId" description:"Feature id"`
	ListFeeType []FeeItemInfo `json:"listFeeType" description:"List fee type"`
}

type FeeItemInfo struct {
	FeeId         string          `json:"feeId" description:"Fee id"`
	MaxFee        decimal.Decimal `json:"maxFee" description:"Max fee"`
	MinFee        decimal.Decimal `json:"minFee" description:"Min fee"`
	FeeCalcMethod string          `json:"feeCalcMethod" description:"fee Calc Method:0-By Percent;1-Fixed Amount"`
	CalcValue     decimal.Decimal `json:"calcValue" description:"Calc value"`
	PayerType     string          `json:"payerType" description:"payer Type:01-Payee Participant;02-Payer Participant"`
	EffectiveDate string          `json:"effectiveDate" description:"Effective date"`
	ExpireDate    string          `json:"expireDate" description:"Expire date"`
}

func (*PD000084I) GetServiceKey() string {
	return "PD000084"
}
