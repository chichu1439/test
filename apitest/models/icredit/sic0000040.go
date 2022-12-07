//Version: v1.0.0
package models

import "github.com/shopspring/decimal"

type IC000040I struct {
	OrgNum string `json:"orgNum" description:"Org number" validate:"required"`
}

type IC000040O struct {
	OverdueAmt map[string]decimal.Decimal `json:"overdueAmt" description:"Overdue amt"`
}

func (i *IC000040I) GetServiceKey() string {
	return "IC000040"
}
