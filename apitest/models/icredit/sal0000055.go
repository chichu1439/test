package models

import "github.com/shopspring/decimal"

type AL000055I struct {
	AccountingAccountNum string `json:"accountingAccountNum" description:"Accounting account number" validate:"required"`
}

type AL000055O struct {
	WriteoffAmount decimal.Decimal `json:"writeoffAmount" description:"Writeoff amount"`
	WriteoffType   string          `json:"writeoffType" description:"Writeoff type:0-partial 1-total"`
	Status         string          `json:"status" description:"Status:01-to be approve 02-approved 03-rejected 04-cancelled"`
}

func (*AL000055I) GetServiceKey() string {
	return "AL000055"
}
