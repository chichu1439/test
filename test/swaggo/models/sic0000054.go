package models

type IC000054I struct {
	EditLoanApplySerialNum string `json:"editLoanApplySerialNum" description:"Edit loan apply serial number" validate:"required"`
	LoanReceiptNum         string `json:"loanReceiptNum" description:"Loan receipt number" validate:"required"`
	CustomerId             string `json:"customerId" description:"Customer id" validate:"required"`
	ApproveView            string `json:"approveView" description:"approve View" validate:"required"`
	ApproveComment         string `json:"approveComment" description:"Approve comment"`
}

type IC000054O struct {
	LoanReceiptNum string `json:"loanReceiptNum" description:"Loan receipt number"`
}

func (*IC000054I) GetServiceKey() string {
	return "IC000054"
}
