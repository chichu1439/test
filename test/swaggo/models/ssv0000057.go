//Version: v1
package models

import "github.com/shopspring/decimal"

type SV000057Request struct {
	TransactionNumber     string          `json:"transactionNumber" validate:"required"`
	ReversalFeeFlag       string          `json:"reversalFeeFlag" validate:"required"`
	CheckType             string          `json:"checkType" validate:"required"`
	AgreementNumber       string          `json:"agreementNumber" validate:"required"`
	DebitSettleInterest   decimal.Decimal `json:"amountInterestDebit" validate:"required"`
	DebitAccrualInterest  decimal.Decimal `json:"accruedInterestDebit" validate:"required"`
	CreditSettleInterest  decimal.Decimal `json:"amountInterestCredit" validate:"required"`
	CreditAccrualInterest decimal.Decimal `json:"accruedInterestCredit" validate:"required"`
}

type SV000057Response struct {
	Remark string `json:"remark" description:"Remark"`
	Status string `json:"status"`
}

func (*SV000057Request) GetServiceKey() string {
	return "SV000057"
}
