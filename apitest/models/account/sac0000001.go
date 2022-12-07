//Version: v0.0.1
package models

type SAC0000001I struct {
	CustomerId         string  `json:"customerId" description:"Customer id" validate:"required"`
	Account            string  `json:"account" description:"Account" validate:"required"`
	Amount             float64 `json:"amount" description:"Amount"`
	AllowWithdrawTimes int     `json:"allowWithdrawTimes" description:"Allow withdraw times" validate:"required"`
	Currency           string  `json:"currency" description:"Currency" validate:"required"`
	ObjectPublic       string  `json:"objectPublic" description:"Object public"`
	BranchNum          string  `json:"branchNum" description:"Branch number" validate:"required"`
	TellerId           string  `json:"tellerId" description:"Teller id" validate:"required"`
}

type SAC0000001O struct {
	DcnId        string `json:"dcnId" description:"Dcn id"`
	AccountingId string `json:"accountingId" description:"Accounting id"`
}

func (*SAC0000001I) GetServiceKey() string {
	return "sac0000001"
}
