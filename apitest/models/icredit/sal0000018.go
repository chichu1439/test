package models

type AL000018I struct {
	RegId                string `json:"regId" description:"Reg id" validate:"required"`
	LoanReceiptNum       string `json:"loanReceiptNum" description:"Loan receipt number" validate:"required"`
	AccountingAccountNum string `json:"accountingAccountNum" description:"Accounting account number" validate:"required"`
	Description          string `json:"description" description:"Description"`
}

type AL000018O struct {
}

func (*AL000018I) GetServiceKey() string {
	return "AL000018"
}
