//Version: v1
package models

import "github.com/shopspring/decimal"

type AL000069Request struct {
	Records map[string][]string `json:"records"`
}

type AL000069Response struct {
	RecordList []AL000069ResponseItem `json:"recordList"`
}
type AL000069ResponseItem struct {
	ProductId          string          `json:"productId"`
	RiskClassifiedCode string          `json:"riskClassifiedCode"`
	LoanTotalPeriod    decimal.Decimal `json:"loanTotalPeriod"`
}

func (*AL000069Request) GetServiceKey() string {
	return "AL000069"
}
