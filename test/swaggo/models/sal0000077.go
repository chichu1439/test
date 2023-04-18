//Version: v1
package models

import "github.com/shopspring/decimal"

type AL000077Request struct {
	LoanReceiptNumList []string `json:"loanReceiptNumList"`
}
type AL000077Response struct {
	AccountInfoList []AccountInfo1 `json:"accountInfoList"`
}

type AccountInfo1 struct {
	AccountingAccountNum string          `json:"accountingAccountNum" description:"Accounting account number"` // ",
	LoanReceiptNum       string          `json:"loanReceiptNum" description:"Loan receipt number"`             // ",
	Balance              decimal.Decimal `json:"balance" description:"Balance"`                                // ",
	RepayCycle           string          `json:"repayCycle" description:"Repay cycle"`                         // 01-天 02-周 03-半月/双周 04-月 05-双月 06-季 07-半年 08-年 ",
	RepayFrequency       int             `json:"repayFrequency" description:"Repay frequency"`                 // ",
	RepayMethod          string          `json:"repayMethod" description:"Repay method"`                       // 01-等额本金 02-等额本息 03-到期（一次性）还本付息 04-到期还本周期还息 05-等本等息 06-随借随还 ",
	RepayDay             string          `json:"repayDay" description:"Repay day"`                             // ",
	InterestRate         decimal.Decimal `json:"interestRate"`
}

func (*AL000077Request) GetServiceKey() string {
	return "AL000077"
}
