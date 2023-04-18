//Version: v1
package models

type IC000115Request struct {
	ApplySerialNum string `json:"applySerialNum" description:"apply serial num" validate:"required"`
	LoanReceiptNum string `json:"loanReceiptNum" description:"loan receipt num"  validate:"required"`
	ApproveComment string `json:"approveComment" description:"Checker comment"`
	ApproveView    string `json:"approveView" description:"Apply status" validate:"required"`
}

type IC000115Response struct {
	LoanReceiptNum string `json:"loanReceiptNum" description:"loan receipt num"`
}

func (*IC000115Request) GetServiceKey() string {
	return "IC000115"
}
