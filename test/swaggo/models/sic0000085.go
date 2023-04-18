//Version: v1
package models

type IC000085Request struct {
	CustomerId                  string `json:"customerId" validate:"required"`
	LoanReceiptNum              string `json:"loanReceiptNum" validate:"required"`
	RestructuringApplySerialNum string `json:"restructuringApplySerialNum" validate:"required"`
	ApproveView                 string `json:"approveView" description:"approve View(2-agree;3-refuse)" validate:"required"`
	ApproveComment              string `json:"approveComment" description:"Approve comment" validate:"required"`
	SourceChannel               string `json:"sourceChannel"`
}

type IC000085Response struct {
	RestructuringApplySerialNum string `json:"restructuringApplySerialNum"`
	LoanReceiptNum              string `json:"loanReceiptNum"`
}
