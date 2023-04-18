package models

type SIC0000029I struct {
	MakeLoanApplySerialNum string `json:"makeLoanApplySerialNum" description:"Make loan apply serial number" validate:"required"`
	LoanReceiptNum         string `json:"loanReceiptNum" description:"Loan receipt number" validate:"required"`
	CustomerId             string `json:"customerId" description:"Customer id" validate:"required"`
	ApproveStatus          string `json:"approveStatus" validate:"required" description:"Approval status:1-to be approved 2-approved 3-rejected"`
	ApproveComment         string `json:"approveComment" description:"Approve comment"`
}

type SIC0000029O struct {
	MakeLoanApplySerialNum string `json:"makeLoanApplySerialNum" description:"Make loan apply serial number"`
	LoanReceiptNum         string `json:"loanReceiptNum" description:"Loan receipt number"`
}

func (*SIC0000029I) GetServiceKey() string {
	return "IC000029"
}
