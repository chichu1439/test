//Version: v1
package models

import "github.com/shopspring/decimal"

type PF000012Request struct {
	FeeType           string          `json:"feeType"  validate:"required"`
	Currency          string          `json:"currency" description:"Currency(THB)"  validate:"required"`
	CalcAmount        decimal.Decimal `json:"calcAmount"  validate:"required"`
	FeeCondition      decimal.Decimal `json:"feeCondition"`
	DiscountCondition decimal.Decimal `json:"discountCondition"`
	RemainDiscount    decimal.Decimal `json:"remainDiscount"`
	ProductId         string          `json:"productId"`
	PaymentInAccount  string          `json:"paymentInAccount" `                       //转出内部户账号
	CltInAccount      string          `json:"cltInAccount" `                           //转入内部户账号
	SourceFunds       string          `json:"sourceFunds" description:"Source funds" ` //资金来源 1-客户账户；2-内部账户；3-合约
	GoingFunds        string          `json:"goingFunds" description:"Going funds" `   //资金去向 1-客户账户；2-内部账户；3-合约
}

type PF000012Response struct {
	FeeAmount           decimal.Decimal `json:"feeAmount"`
	FeeReceivableAmount decimal.Decimal `json:"feeReceivableAmount"`
	FeeDiscountAmount   decimal.Decimal `json:"feeDiscountAmount"`
	RemainDiscount      decimal.Decimal `json:"remainDiscount"`
	FeeDiscountFlag     int             `json:"feeDiscountFlag" description:"1:yes 0:no"`
}

func (*PF000012Request) GetServiceKey() string {
	return "PF000012"
}
