//Version: v0.0.1
package models

import "github.com/shopspring/decimal"

type AL000066I struct {
	LoanReceiptNum  string `json:"loanReceiptNum" description:"Loan receipt number" validate:"required"`
	CustomerId      string `json:"customerId" description:"Customer id" validate:"required"`
	TransDate       string `json:"transDate" description:"Trans date" validate:"required"`
	Period          int    `json:"period" description:"Period" validate:"required"`
	RecalculateType string `json:"recalculateType" description:"Recalculate type"`
}

type AL000066O struct {
	NeedRepayTotalAmount             decimal.Decimal `json:"needRepayTotalAmount" description:"Need repay total amount"`
	NeedRepayBalance                 decimal.Decimal `json:"needRepayBalance" description:"Need repay balance"`
	NeedRepayTotalPrinciple          decimal.Decimal `json:"needRepayTotalPrinciple" description:"Need repay total principle"`
	NeedRepayImpairmentPrinciple     decimal.Decimal `json:"needRepayImpairmentPrinciple" description:"Need repay impairment principle"`
	NeedRepayOverduePrinciple        decimal.Decimal `json:"needRepayOverduePrinciple" description:"Need repay overdue principle"`
	NeedRepayTotalInterest           decimal.Decimal `json:"needRepayTotalInterest" description:"Need repay total interest"`
	NeedRepayNormalInterest          decimal.Decimal `json:"needRepayNormalInterest" description:"Need repay normal interest"`
	NeedRepayTotalCompoundInterest   decimal.Decimal `json:"needRepayTotalCompoundInterest" description:"Need repay total compound interest"`
	NeedRepayNormalCompoundInterest  decimal.Decimal `json:"needRepayNormalCompoundInterest" description:"Need repay normal compound interest"`
	NeedRepayOverDueInterest         decimal.Decimal `json:"needRepayOverDueInterest" description:"Need repay over due interest"`
	NeedRepayOverDueCompoundInterest decimal.Decimal `json:"needRepayOverDueCompoundInterest" description:"Need repay over due compound interest"`
	Balance                          decimal.Decimal `json:"balance" description:"Balance"`
	RepayAmount                      decimal.Decimal `json:"repayAmount" description:"Repay amount"`
	OutstandingBalance               decimal.Decimal `json:"outstandingBalance" description:"Outstanding balance"`
	RepayTotalPrinciple              decimal.Decimal `json:"repayPrinciple" description:"Repay total principle"`
	RepayImpairmentPrinciple         decimal.Decimal `json:"repayImpairmentPrinciple" description:"Repay impairment principle"`
	RepayOverduePrinciple            decimal.Decimal `json:"repayOverduePrinciple" description:"Repay overdue principle"`
	RepayTotalInterest               decimal.Decimal `json:"repayTotalInterest" description:"Repay total interest"`
	RepayNormalInterest              decimal.Decimal `json:"repayInterest" description:"Repay normal interest"`
	RepayTotalCompoundInterest       decimal.Decimal `json:"repayTotalCompoundInterest" description:"Repay total compound interest"`
	RepayNormalCompoundInterest      decimal.Decimal `json:"repayNormalCompoundInterest" description:"Repay normal compound interest"`
	RepayOverDueInterest             decimal.Decimal `json:"repayOverDueInterest" description:"Repay over due interest"`
	RepayOverDueCompoundInterest     decimal.Decimal `json:"repayOverDueCompoundInterest" description:"Repay over due compound interest"`
	NeedRepayOverDueNormalInterest   decimal.Decimal `json:"needRepayOverDueNormalInterest"`
	EarlyRepaymentFlag               bool            `json:"earlyRepaymentFlag" description:"Early repayment flag"`
}

func (*AL000066I) GetServiceKey() string {
	return "AL000066"
}
