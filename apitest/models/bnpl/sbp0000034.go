//Version: v1
package models

type BP000034Request struct {
	ProviderId string `json:"providerId" description:"Provider ID" validate:"required,max=20"`
}

type BP000034Response struct {
	ProviderId                string `json:"providerId" description:"Provider ID"`
	LoanProductId             string `json:"loanProductId" description:"Loan Product ID"`
	LoanPurpose               string `json:"loanPurpose" description:"Loan Purpose"`
	InstallmentList           string `json:"installmentList" description:"Installment List"`
	DefaultPeriod             int    `json:"defaultPeriod" description:"Default Period: InstallmentList split by “|”, piking up the first"`
	RepaymentFrequency        string `json:"repaymentFrequency" description:"Repayment Frequency"`
	RepaymentMethod           string `json:"repaymentMethod" description:"Repayment Method"`
	DefaultRepaymentFrequency string `json:"RepaymentFrequency split by “|“, piking up the first "`
	DefaultRepaymentMethod    string `json:"defaultRepaymentMethod" description:"RepaymentMethod split by “|“, piking up the first"`
}

func (*BP000034Request) GetServiceKey() string {
	return "BP000034"
}
