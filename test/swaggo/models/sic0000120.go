//Version: v1
package models

type IC000120Request struct {
	CustomerId     string `json:"customerId" validate:"required"`
	LoanReceiptNum string `json:"loanReceiptNum" validate:"required"`
	ApplySerialNum string `json:"applySerialNum" validate:"required" `
	ApproveView    string `json:"approveView" description:"approve View(2-agree;3-refuse)" validate:"required"`
	ApproveComment string `json:"approveComment" description:"Approve comment" validate:"required"`
	SourceChannel  string `json:"sourceChannel"`
}

type IC000120Response struct {
	ApplySerialNum    string `json:"applySerialNum" validate:"required"`
	LoanReceiptNum    string `json:"loanReceiptNum"`
	NewLoanReceiptNum string `json:"newLoanReceiptNum"`
}

func (*IC000120Request) GetServiceKey() string {
	return "IC000120"
}
