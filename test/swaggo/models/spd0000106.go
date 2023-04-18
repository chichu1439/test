//Version: v1
package models

type PD000106Request struct {
	Code                   string `json:"code" description:"code"`
	Behavior               string `json:"behavior" description:"behavior(01-Repayment;02-Close Repayment;03-Refinance Repayment)"`
	RepaymentHierarchyType string `json:"repaymentHierarchyType" description:"repayment hierarchy type" `
	RepaymentHierarchyName string `json:"repaymentHierarchyName" description:"repayment hierarchy name" `
}

type PD000106Response struct {
	ListRepaymentHierarchy []RepaymentHierarchy
}
type RepaymentHierarchy struct {
	Code                   string `json:"code"`
	Behavior               string `json:"behavior"`
	RepaymentHierarchyType string `json:"repaymentHierarchyType"`
	RepaymentHierarchyName string `json:"repaymentHierarchyName" description:"repayment hierarchy name" `
	RepaymentOrder         string `json:"repaymentOrder"`
	Description            string `json:"description"`
	CreateDate             string `json:"createDate" description:"create date"`
	UpdateDate             string `json:"updateDate" description:"update date"`
}

func (*PD000106Request) GetServiceKey() string {
	return "PD000106"
}
