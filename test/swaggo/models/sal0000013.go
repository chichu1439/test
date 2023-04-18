package models

type AL000013I struct {
	RegId                string `json:"regId" description:"Reg id" validate:"required"`
	LoanReceiptNum       string `json:"loanReceiptNum" description:"Loan receipt number" validate:"required"`
	AccountingAccountNum string `json:"accountingAccountNum" description:"Accounting account number" validate:"required"`
	ApproveView          string `json:"approveView" description:"Approve type:0-approved 1-rejected"`
	ApproveComment       string `json:"approveComment" description:"approve comment"`
}

type AL000013O struct {
	RegId string `json:"regId" description:"Reg id"`
}

func (*AL000013I) GetServiceKey() string {
	return "AL000013"
}
