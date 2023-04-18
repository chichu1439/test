//Version: v1
package models

type IC000125Request struct {
	CustomerId     string `json:"customerId" validate:"required"`
	LoanReceiptNum string `json:"loanReceiptNum" validate:"required"`
	ApplySerialNum string `json:"applySerialNum" validate:"required" `
	ApproveView    string `json:"approveView" description:"approve View(2-agree;3-refuse)" validate:"required"`
	ApproveComment string `json:"approveComment" description:"Approve comment" validate:"required"`
}

type IC000125Response struct {
	ApplySerialNum string `json:"applySerialNum" validate:"required"`
	LoanReceiptNum string `json:"loanReceiptNum"`
}

func (*IC000125Request) GetServiceKey() string {
	return "IC000125"
}
