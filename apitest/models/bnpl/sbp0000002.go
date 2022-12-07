//Version: v1
package models

import "github.com/shopspring/decimal"

type BP000002Request struct {
	CustomerId    string          `json:"customerId" description:"Customer ID" validate:"required, max=20"`
	TotalAmount   decimal.Decimal `json:"totalAmount" description:"Total Amount" validate:"required"`
	PeriodNum     int             `json:"periodNum" description:"Period Number" validate:"required"`
	LoanProductId string          `json:"loanProductId" description:"Loan Product ID" validate:"required,max=10"`
}

type BP000002Response struct {
	RepaymentPlanList []RepaymentPlanInfo `json:"repaymentPlanList" description:"Repayment Plan List"`
}

func (*BP000002Request) GetServiceKey() string {
	return "BP000002"
}
