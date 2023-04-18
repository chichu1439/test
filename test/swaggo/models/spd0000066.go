package models

import "github.com/shopspring/decimal"

type SPD0000066I struct {
	ProductId string `json:"productId" description:"Product id" validate:"required"`
	TopicId   string `json:"topicId" description:"Topic id" validate:"required"`
}

type SPD0000066O struct {
	ServiceId         string          `json:"serviceId" description:"Service id"`
	DiscountMethod    string          `json:"discountMethod" description:"Discount method:0-By Each Fee 1-By Total Fee"`                             //0-By Each Fee 1-By Total Fee
	MultiDiscountType string          `json:"multiDiscountType" description:"Multi discount type:0-Min Discount 1-Max Discount 2-Sum"`               //0-Min Discount 1-Max Discount 2-Sum ,When Discount Method=0
	DiscountCalcType  string          `json:"discountCalcType" description:"Discount calc type:0-by percent 1-fixed amount ,When Discount Method=1"` //0-by percent 1-fixed amount ,When Discount Method=1
	MaxFeeAmount      decimal.Decimal `json:"maxFeeAmount" description:"Max fee amount"`
	MinFeeAmount      decimal.Decimal `json:"minFeeAmount" description:"Min fee amount"`
	MaxDiscountAmount decimal.Decimal `json:"maxDiscountAmount" description:"Max discount amount"`
	MinDiscountAmount decimal.Decimal `json:"minDiscountAmount" description:"Min discount amount"`
	ListFeeId         []ListFeeId     `json:"listFeeId" description:"List fee id"`
}

type ListFeeId struct {
	FeeId string `json:"feeId" description:"Fee id"`
}

func (*SPD0000066I) GetServiceKey() string {
	return "PD000066"
}
