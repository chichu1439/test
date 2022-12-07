package models

type SIC0000007I struct {
	MakeLoanApplySerialNum string `json:"makeLoanApplySerialNum" validate:"required" description:"Make Loan Apply SerialNum"`
	CustomerId             string `json:"customerId" description:"Customer id" validate:"required"`
	ApproveView            string `json:"approveView" validate:"required" description:"Approve View"`
	ApproveComment         string `json:"approveComment" description:"Approve Comment" `
}

type SIC0000007O struct {
	MakeLoanApplySerialNum string `json:"makeLoanApplySerialNum" description:"Make Loan Apply SerialNum"`
	LoanReceiptNum         string `json:"loanReceiptNum" description:"Loan Receipt number"`
}

func (*SIC0000007I) GetServiceKey() string {
	return "sic0000007"
}
