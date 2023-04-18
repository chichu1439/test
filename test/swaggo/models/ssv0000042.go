//Version: v0.0.1
package models

import "github.com/shopspring/decimal"

type SV000042I struct {
	AgreementId     string `json:"agreementId" description:"Contract number" validate:"required,max=50"` //合约号
	AgreementType   string `json:"agreementType" description:"Contract type" validate:"required,max=5"`  //合约类型
	EndSettleDate   string `json:"endSettleDate"`
	ProcessType     string `json:"processType"` //"1"-试算 "2"-结息
	StartSettleDate string `json:"startSettleDate"`
	Currency        string `json:"currency"`
	AccountNum      string `json:"accountNum"`
}

type SV000042O struct {
	SettleInterest       decimal.Decimal        `json:"settleInterest"`
	InterestCalcDateRang []InterestCalcDateRang `json:"interestCalcDateRang"`
}

//type AccountList struct {
//	AccountNumber string
//	AccountName   string
//	Currency      string
//	AccountStatus string
//}
type Interest2 struct {
	InterestCalcDateRang []InterestCalcDateRang `json:"interestCalcDateRang"`
}
type InterestCalcDateRang struct {
	StartDate      string          `json:"startDate"`
	EndDate        string          `json:"endDate"`
	CalcIntBal     decimal.Decimal `json:"calcIntBal"`
	CalcIntDays    int64           `json:"calcIntDays"`
	InterestRate   decimal.Decimal `json:"interestRate"`
	InterestAmount decimal.Decimal `json:"interestAmount"`
}

func (*SV000042I) GetServiceKey() string {
	return "SV000042"
}
