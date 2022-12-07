//Version: v1
package models

import "github.com/shopspring/decimal"

type BP000031Request struct {
	OrderId    string `json:"orderId" description:"Order ID" validate:"required,max=20"`
	MerchantId string `json:"merchantId" description:"Merchant ID" validate:"required,max=20"`
}

type BP000031Response struct {
	OrderId          string          `json:"orderId" description:"Order ID"`
	MerchantId       string          `json:"merchantId" description:"Merchant ID"`
	MerchantName     string          `json:"merchantName" description:"Merchant Name"`
	TotalOrderAmount decimal.Decimal `json:"totalOrderAmount" description:"Total Order Amount"`
}

func (*BP000031Request) GetServiceKey() string {
	return "BP000031"
}
