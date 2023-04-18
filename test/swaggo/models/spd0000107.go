//Version: v1
package models

type PD000107Request struct {
	Code                   string `json:"code" validate:"required" description:"code"`
	RepaymentHierarchyType string `json:"repaymentHierarchyType" description:"repayment hierarchy type" `
	RepaymentHierarchyName string `json:"repaymentHierarchyName" description:"repayment hierarchy name" `
	RepaymentOrder         string `json:"repaymentOrder" description:"repayment order ,format； repayment priority：repayment Hierarchy" `
	Description            string `json:"description" description:"description " `
}

type PD000107Response struct {
	Code string `json:"code" description:"code"`
}

func (*PD000107Request) GetServiceKey() string {
	return "PD000107"
}
