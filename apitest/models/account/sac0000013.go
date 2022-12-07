//Version: v0.0.1
package models

type AC000013I struct {
	AcctList []HistoryBalanceQuery `json:"acctList" description:"Acct list" `
}

type AC000013O struct {
	ResultList []ResultHistoryBalance `json:"resultList" description:"Result list"`
}

type ResultHistoryBalance struct {
	AcctNo      string  `json:"acctNo" description:"Acct no" `
	BalanceDate string  `json:"balanceDate" description:"Balance date" `
	Balance     float64 `json:"balance" description:"Balance" `
	Currency    string  `json:"currency" description:"Currency" `
}

type HistoryBalanceQuery struct {
	AcctNo      string `json:"acctNo" description:"Acct no" validate:"required"`
	BalanceDate string `json:"balanceDate" description:"Balance date" validate:"required"`
}

func (*AC000013I) GetServiceKey() string {
	return "sac0000013"
}
