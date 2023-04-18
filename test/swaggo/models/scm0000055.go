//Version: v1
package models

type CM000055Request struct {
	DisburseDate string `json:"disburseDate" description:"Disburse Date" validate:"required"`
}

type CM000055Response struct {
	FirstRepaymentDate string `json:"firstRepaymentDate"  description:"First Repayment Date"`
	RepaymentDay       string `json:"repaymentDay" description:"Repayment Day"`
}

func (*CM000055Request) GetServiceKey() string {
	return "CM000055"
}
