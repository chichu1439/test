package models

import "github.com/shopspring/decimal"

type PF000010I struct {
	ProductId   string          `json:"productId" description:"Product id" validate:"required"`
	TopicId     string          `json:"topicId" description:"Topic id" validate:"required"`
	CalcAmount  decimal.Decimal `json:"calcAmount" description:"Calc amount" validate:"required"`
	ListBizParm []ListBizParam  `json:"listBizParm" description:"List biz parm" validate:"required,dive"`
}

type ListBizParam struct {
	ParmCode  string `json:"parmCode" description:"Parm code" validate:"required"`
	ParmValue string `json:"parmValue" description:"Parm value" validate:"required"`
}

type PF000010O struct {
	ServiceId         string          `json:"serviceId" description:"Service id"`
	FeeAmount         decimal.Decimal `json:"feeAmount" description:"Fee amount"`           //sum of listFee.feeAmount
	DiscountAmount    decimal.Decimal `json:"discountAmount" description:"Discount amount"` //sum of listFee.discountAmount
	FeeAmountActual   decimal.Decimal `json:"feeAmountActual" description:"Fee amount actual"`
	MaxFeeAmount      decimal.Decimal `json:"maxFeeAmount" description:"Max fee amount"`
	MinFeeAmount      decimal.Decimal `json:"minFeeAmount" description:"Min fee amount"`
	MaxDiscountAmount decimal.Decimal `json:"maxDiscountAmount" description:"Max discount amount"`
	MinDiscountAmount decimal.Decimal `json:"minDiscountAmount" description:"Min discount amount"`
	ListFee           []ListFee       `json:"listFee" description:"List fee"`
}

type ListFee struct {
	FeeId              string          `json:"feeId" description:"Fee id"`
	FeeName            string          `json:"feeName" description:"Fee name"`
	StrategyId         string          `json:"strategyId" description:"Strategy id"`
	FeeAmount          decimal.Decimal `json:"feeAmount" description:"Fee amount"`              //fee calculation result
	DiscountAmount     decimal.Decimal `json:"discountAmount" description:"Discount amount"`    //discount calculation result
	FeeAmountActual    decimal.Decimal `json:"feeAmountActual" description:"Fee amount actual"` //actual fee amount need to pay: fee amount-discount amount.
	IsAmortization     string          `json:"isAmortization" description:"Is amortization"`
	AmortizationPeriod string          `json:"amortizationPeriod" description:"Amortization period"`
	AmortizationTimes  *int            `json:"amortizationTimes" description:"Amortization times"`
	ChargeType         string          `json:"chargeType" description:"Charge type"`
	ChargePeriod       string          `json:"chargePeriod" description:"Charge period"`
	ChargePeriodValue  *int            `json:"chargePeriodValue" description:"Charge period value"`
	ChargePeriodDay    *int            `json:"chargePeriodDay" description:"Charge period day"`
	MaxFeeAmount       decimal.Decimal `json:"maxFeeAmount" description:"Max fee amount"`
	MinFeeAmount       decimal.Decimal `json:"minFeeAmount" description:"Min fee amount"`
	AccountingCode     string          `json:"accountingCode" description:"Accounting code"`
}

func (*PF000010I) GetServiceKey() string {
	return "PF000010"
}
