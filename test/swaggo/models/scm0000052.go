//Version: v1.0.0
package models

type CM000052I struct {
	OrgNum string `json:"orgNum" description:"Org number" validate:"required"`
}

type CM000052O struct {
	CollecTaskTotal float64       `json:"collecTaskTotal" description:"Collec task total"`
	ListStatus      []ListStatus2 `json:"listStatus" description:"List status"`
}

type ListStatus2 struct {
	RepaymentStatus string  `json:"repaymentStatus" description:"Repayment status"`
	TaskStatus      string  `json:"taskStatus" description:"Task status" `
	Percent         float64 `json:"percent" description:"Percent" `
	TotalCount      int     `json:"totalCount" description:"Total count" `
}

func (*CM000052I) GetServiceKey() string {
	return "CM000052"
}
