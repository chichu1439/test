//Version: v1
package models

type PD000105Request struct {
	Behavior               string `json:"behavior" validate:"required" description:"behavior(01-Repayment;02-Close Repayment;03-Refinance Repayment)"`
	RepaymentHierarchyType string `json:"repaymentHierarchyType" validate:"required" description:"repayment hierarchy type" `
	RepaymentHierarchyName string `json:"repaymentHierarchyName" validate:"required" description:"repayment hierarchy name" `
	Description            string `json:"description" description:"description " `
	RepaymentOrder         string `json:"repaymentOrder" validate:"required" description:"repayment order ,format； repayment priority：repayment Hierarchy" `
}

type PD000105Response struct {
	Code                   string `json:"code"`
	RepaymentHierarchyName string `json:"repaymentHierarchyName" description:"repayment hierarchy name" `
}

func (*PD000105Request) GetServiceKey() string {
	return "PD000105"
}
