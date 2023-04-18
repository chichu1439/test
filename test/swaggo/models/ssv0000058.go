//Version: v1
package models

import "github.com/shopspring/decimal"

type SV000058Request struct {
	GlobalBizSeq string `json:"globalBizSeq" description:"Global business number sequence" validate:"required"`
}

type SV000058Response struct {
	DebitSettleInterest   decimal.Decimal `json:"debitSettleInterest" `
	DebitAccrualInterest  decimal.Decimal `json:"debitAccrualInterest"`
	CreditSettleInterest  decimal.Decimal `json:"creditSettleInterest"`
	CreditAccrualInterest decimal.Decimal `json:"creditAccrualInterest"`
}

func (*SV000058Request) GetServiceKey() string {
	return "SV000058"
}
