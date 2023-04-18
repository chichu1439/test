//Version: v1
package models

import (
	"github.com/shopspring/decimal"
)

type AL000073Request struct {
	WriteOffRefinancingFlag     string              `json:"writeOffRefinancingFlag" description:"0-Write off 1- refinancing"`
	ChangeAmount                decimal.Decimal     `json:"changeAmount" description:"refinancing or wirte off amount"`
	ProductId                   string              `json:"productId"`
	RestructuringApplySerialNum string              `json:"restructuringApplySerialNum" validate:"required"`
	TransDate                   string              `json:"transDate" validate:"required"`
	AccountingAccountNum        string              `json:"accountingAccountNum" validate:"required"`
	LoanReceiptNum              string              `json:"loanReceiptNum" validate:"required"`
	Currency                    string              `json:"currency" validate:"required"`
	LoanAmount                  decimal.Decimal     `json:"loanAmount" description:"Loan amount" validate:"required"`
	MatureDate                  string              `json:"matureDate" description:"Mature date" validate:"required"`
	LoanPeriod                  int                 `json:"loanPeriod" description:"Loan period"`
	RepaymentMethod             string              `json:"repaymentMethod" description:"Repayment method:01-Fixed Installment;02-Fixed Principal" validate:"required"`
	RepaymentCycle              string              `json:"repaymentCycle" description:"Repayment cycle"`
	RepayDay                    string              `json:"repayDay" description:"Repay day"`
	InterestRate                decimal.Decimal     `json:"interestRate" description:"Interest rate"`
	FloatValue                  decimal.Decimal     `json:"floatValue" description:"float Value"`
	GracePeriod                 int                 `json:"gracePeriod"`
	YearDays                    string              `json:"yearDays" validate:"required" description:"0-360 1-365 2-Actual days"`
	BillRule                    BillRuleInfo        `json:"billRuleInfo"`
	ListFee                     []FeeTransDetail    `json:"listFee"`
	WaiveAmountList             []WaiveAmountDetail `json:"waiveAmountList"`
	RefinacePrincipal           decimal.Decimal     `json:"refinacePrincipal"`
	GlobalSerialNum             string              `json:"globalSerialNum" description:"global serial num"`
	FinalModifyOrgz             string              `json:"finalModifyOrgz"`
	FinalModifyUser             string              `json:"finalModifyUser"`
}
type WaiveAmountDetail struct {
	SeqNum      int             `json:"seqNum"`
	AmountType  string          `json:"amountType"`
	TransAmount decimal.Decimal `json:"transAmount"`
	WaiveAmount decimal.Decimal `json:"waiveAmount"`
}
type AL000073Response struct {
	TotalPeriod int `json:"totalPeriod"`
}

func (*AL000073Request) GetServiceKey() string {
	return "AL000073"
}
