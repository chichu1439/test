//Version: v1
package models

import "github.com/shopspring/decimal"

type AL000072Request struct {
	LoanReceiptNum []string `json:"loanReceiptNum" validate:"required"`
}

type AL000072Response struct {
	SumAccountAmountList []SumAccountAmount `json:"sumAccountAmountList"`
}

type SumAccountAmount struct {
	AccountingAccountNum                 string          `json:"accountingAccountNum"`
	LoanReceiptNum                       string          `json:"loanReceiptNum"`
	InterestRate                         decimal.Decimal `json:"interestRate"`
	TotalPrinciple                       decimal.Decimal `json:"totalPrinciple"`
	TotalInterest                        decimal.Decimal `json:"totalInterest"`
	SumNormalInterest                    decimal.Decimal `json:"sumNormalInterest"`
	SumPenaltyInterest                   decimal.Decimal `json:"sumPenaltyInterest"`
	SumCompoundInterest                  decimal.Decimal `json:"sumCompoundInterest"`
	SumCompoundInterestOfPenaltyInterest decimal.Decimal `json:"sumCompoundInterestOfPenaltyInterest"`
}

func (*AL000072Request) GetServiceKey() string {
	return "AL000072"
}
