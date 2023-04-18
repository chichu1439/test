//Version: v1.0.0
package models

type CM000051I struct {
	OrgNum string `json:"orgNum" description:"Org number" validate:"required"`
}

type CM000051O struct {
	OverdueAmt float64 `json:"overdueAmt" description:"Overdue amt"`
}

func (i *CM000051I) GetServiceKey() string {
	return "CM000051"
}
