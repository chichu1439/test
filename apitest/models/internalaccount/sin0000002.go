//Version: v1.0.0.0
package models

type IN000002I struct {
	AccountingItemNumber string  `json:"accountingItemNumber" description:"Accounting item number" validate:"required"`
	SerialNumber         string  `json:"serialNumber" description:"Serial number" validate:"required"`
	BranchesNumber       string  `json:"branchesNumber" description:"Branches number" validate:"required"`
	Currency             string  `json:"currency" description:"Currency" validate:"required"`
	AccountName          string  `json:"accountName" description:"Account name"`
	DeadlineApAr         int     `json:"deadlineApAr" description:"Deadline ap ar"`
	TermUnit             string  `json:"termUnit" description:"Term unit"`
	DebitLimit           float64 `json:"debitLimit" description:"Debit limit"`
	CreditLimit          float64 `json:"creditLimit" description:"Credit limit"`
	ExpiryDate           string  `json:"expiryDate" description:"Expiry date"`
}

type IN000002O struct {
	InternalAccount string `json:"internalAccount" description:"Internal account"`
}

func (*IN000002I) GetServiceKey() string {
	return "IN000002"
}
