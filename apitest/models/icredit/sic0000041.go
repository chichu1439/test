//Version: v1.0.0
package models

type IC000041I struct {
	OrgNum string `json:"orgNum" description:"Org number" validate:"required"`
}

type IC000041O struct {
	CollecTaskTotal float64      `json:"collecTaskTotal" description:"Collec task total"`
	ListStatus      []ListStatus `json:"listStatus" description:"List status"`
}

type ListStatus struct {
	RepaymentStatus string  `json:"repaymentStatus" description:"Repayment status"`
	TaskStatus      string  `json:"taskStatus" description:"Task status" `
	Percent         float64 `json:"percent" description:"Percent" `
	TotalCount      int     `json:"totalCount" description:"Total count" `
}

func (*IC000041I) GetServiceKey() string {
	return "IC000041"
}
