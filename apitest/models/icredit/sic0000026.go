package models

type IC000026I struct {
	LoanReceiptNum string `json:"loanReceiptNum" description:"Loan receipt number" validate:"required"`
	ApplySerialNum string `json:"applySerialNum" description:"Apply serial number" validate:"required"`
	CustomerId     string `json:"customerId" description:"Customer id" validate:"required"`
	ApplyStatus    string `json:"applyStatus" validate:"required,lte=3,gt=0" description:"Apply status:2-approved 3-rejected"` //2-审批通过 3-审批拒绝
	CheckerComment string `json:"checkerComment" description:"Checker comment"`
}

type IC000026O struct {
	Status string `json:"status" description:"Status"`
}

func (*IC000026I) GetServiceKey() string {
	return "IC000026"
}
