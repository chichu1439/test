//Version: v1
package models

import "github.com/shopspring/decimal"

type BP000007Request struct {
	LoanReceiptNum string `json:"loanReceiptNum" description:"Loan receipt number" validate:"required"`
	LoanCustomerId string `json:"loanCustomerId" description:"Loan Customer Id" validate:"required"`
	RepaymentType  string `json:"repaymentType" description:"Repayment Type" validate:"required"`
}

type BP000007Response struct {
	NeedRepayTotalAmount               decimal.Decimal `json:"needRepayTotalAmount" description:"Need Repay Total Amount"`
	NeedRepayBalance                   decimal.Decimal `json:"needRepayBalance" description:"Need Repay Balance"`
	NeedRepayTotalPrinciple            decimal.Decimal `json:"needRepayTotalPrinciple" description:"Need Repay Total Principle"`
	NeedRepayImpairmentPrinciple       decimal.Decimal `json:"needRepayImpairmentPrinciple" description:"Need Repay Impairment Principle"`
	NeedRepayOverduePrinciple          decimal.Decimal `json:"needRepayOverduePrinciple" description:"Need Repay Overdue Principle"`
	NeedRepayTotalInterest             decimal.Decimal `json:"needRepayTotalInterest" description:"Need Repay Total Interest"`
	NeedRepayNormalInterest            decimal.Decimal `json:"needRepayNormalInterest" description:"Need Repay Normal Interest"`
	NeedRepayNormalInterestTotal       decimal.Decimal `json:"needRepayNormalInterestTotal" description:"Need Repay Normal Interest Total"`
	NeedRepayTotalCompoundInterest     decimal.Decimal `json:"needRepayTotalCompoundInterest" description:"Need Repay Total Compound Interest"`
	NeedRepayNormalCompoundInterest    decimal.Decimal `json:"needRepayNormalCompoundInterest" description:"Need Repay Normal Compound Interest"`
	NeedRepayOverDueInterest           decimal.Decimal `json:"needRepayOverDueInterest" description:"Need Repay Over Due Interest"`
	NeedRepayOverDueCompoundInterest   decimal.Decimal `json:"needRepayOverDueCompoundInterest" description:"Need Repay Over Due Compound Interest"`
	NeedRepayOverDueNormalInterest     decimal.Decimal `json:"needRepayOverDueNormalInterest" description:"Need Repay Over Due Normal Interest"`
	NeedRepayOverDueRemainPlanInterest decimal.Decimal `json:"needRepayOverDueRemainPlanInterest" description:"Need Repay Over Due Remain Plan Interest"`
	NeedRepayCollectionFee             decimal.Decimal `json:"needRepayCollectionFee" description:"Need Repay Collection Fee"`
	NeedRepayTotalFee                  decimal.Decimal `json:"needRepayTotalFee" description:"Need Repay Total Fee"`
}

func (*BP000007Request) GetServiceKey() string {
	return "BP000007"
}
