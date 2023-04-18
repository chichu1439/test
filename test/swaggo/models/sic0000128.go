//Version: v1
package models

type IC000128Request struct {
	LoanReceiptNum            string `json:"loanReceiptNum" description:"Loan receipt number" validate:"required"`
	ApplySerialNum            string `json:"applySerialNum" description:"Apply serial number" validate:"required"`
	OriginalBusinessSerialNum string `json:"originalBusinessSerialNum" description:"original business serial num" `
	OriginalGlobalSerialNum   string `json:"originalGlobalSerialNum"description:"original global serial number" validate:"required"`
	CustomerId                string `json:"customerId" description:"Customer Id"`
	ApproveComment            string `json:"approveComment" description:"Checker comment"`
	ApproveView               string `json:"approveView" description:"Apply status" validate:"required"`
}

type IC000128Response struct {
	LoanReceiptNum         string `json:"loanReceiptNum" description:"Loan receipt number" `
	OriginalLoanReceiptNum string `json:"originalLoanReceiptNum" description:"Loan receipt number" `
	Status                 string `json:"status"`
}

func (*IC000128Request) GetServiceKey() string {
	return "IC000128"
}
