package models

type AL000006I struct {
	AccountingAccountNum string `json:"accountingAccountNum" description:"Accounting account number" validate:"required"`
	CustomerId           string `json:"customerId" description:"Customer id"`
	NextRepaymentDate    string `json:"nextRepaymentDate" description:"Next repayment date"`
	TotalPeriod          int    `json:"totalPeriod" description:"Total period"`
	MaturityDate         string `json:"maturityDate" description:"Maturity date"`
	GracePeriod          *int   `json:"gracePeriod" description:"Grace period"`
	AutoRepayFlag        string `json:"autoRepayFlag" description:"Auto repay flag:1-Auto repay 2-Manual repay"` //1-Auto repay 2-Manual repay
	RepayLocalBankFlag   string `json:"repayLocalBankFlag" description:"Repay local bank flag"`
	RepayBankNum         string `json:"repayBankNum" description:"Repay bank number"`
	RepayBankName        string `json:"repayBankName" description:"Repay bank name"`
	RepayAccountNum      string `json:"repayAccountNum" description:"Repay account number"`
	RepayAccountName     string `json:"repayAccountName" description:"Repay account name"`
	FinalModifyOrgz      string `json:"finalModifyOrgz" description:"Final modify organization"`
	FinalModifyUser      string `json:"finalModifyUser" description:"Final modify user"`
}

type AL000006O struct {
}

func (*AL000006I) GetServiceKey() string {
	return "AL000006"
}
