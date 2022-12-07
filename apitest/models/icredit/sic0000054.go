package models

type IC000054I struct {
	MakeLoanApplySerialNum string `json:"makeLoanApplySerialNum" description:"Make loan apply serial number" validate:"required"`
	LoanReceiptNum         string `json:"loanReceiptNum" description:"Loan receipt number" validate:"required"`
	CustomerId             string `json:"customerId" description:"Customer id" validate:"required"`
	ApproveStatus          string `json:"approveStatus" description:"Approve status" validate:"required"`
	ApproveComment         string `json:"approveComment" description:"Approve comment"`
}

type IC000054O struct {
	MakeLoanApplySerialNum string `json:"makeLoanApplySerialNum" description:"Make loan apply serial number"`
	LoanReceiptNum         string `json:"loanReceiptNum" description:"Loan receipt number"`
}

func (*IC000054I) GetServiceKey() string {
	return "IC000054"
}
