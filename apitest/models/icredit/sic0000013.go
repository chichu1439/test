package models

import "github.com/shopspring/decimal"

type IC000013I struct {
	ProductId string `json:"productId" description:"Product id"`
	Currency  string `json:"currency" description:"Currency"`
}

type IC000013O struct {
	Records []BalanceStatisticsBean `json:"records" description:"Records"`
}

type BalanceStatisticsBean struct {
	BalancePerMon     decimal.Decimal `json:"balancePerMon" description:"Loan balance at the end of the month"`     //当月月底贷款余额
	NLoanAmountPerMon decimal.Decimal `json:"nLoanAmountPerMon" description:"New loan amount of the current month"` //当月新增贷款金额
	//SeqGrowthPerMon   int             `json:"seqGrowthPerMon"`   //当月新增贷款账户数量
	CurrentMonth string `json:"currentMonth" description:"Month of the month (2021-05)"` //当月月份(2021-05)
	Currency     string `json:"currency" description:"Currency"`
}

func (*IC000013I) GetServiceKey() string {
	return "IC000013"
}
