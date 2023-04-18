//Version: v1
package models

import "github.com/shopspring/decimal"

type AL000075Request struct {
	ApplySerialNum string `json:"applySerialNum"`
}
type AL000075Response struct {
	ListInterest          []AL000075InterestList    `json:"listInterest"`
	ContractList          []SIC0000101OContractList `json:"contractList"`
	RepaymentPlan         []AL000075RepayPlan       `json:"repaymentPlan"`
	LoanRefNum            string                    `json:"loanRefNum"`
	LoanStatus            string                    `json:"loanStatus"`
	ArrangementPurpose    string                    `json:"arrangementPurpose"`
	ProductId             string                    `json:"productId"`
	RepayMethod           string                    `json:"repayMethod"`
	FreeMinimumAmount     decimal.Decimal           `json:"freeMinimumAmount"`
	PercentageOfPrincipal decimal.Decimal           `json:"percentageOfPrincipal"`
	RepayFrequency        int                       `json:"repayFrequency"`
	BillingDate           string                    `json:"billingDate"`
	RepayDate             string                    `json:"repayDate"`
	SameMonth             string                    `json:"sameMonth"`
	GracePeriod           int                       `json:"gracePeriod"`
	AppDate               string                    `json:"appDate"`
	DrawDate              string                    `json:"drawDate"`
	MatuirtyDate          string                    `json:"matuirtyDate"`
	LoanAmount            decimal.Decimal           `json:"loanAmount"`
	PrincipalAmount       decimal.Decimal           `json:"principalAmount"`
	CreditLine            decimal.Decimal           `json:"creditLine"`
	CreditLineStatus      string                    `json:"creditLineStatus"`
	AvailableCreditLine   decimal.Decimal           `json:"availableCreditLine"`
	Cif                   string                    `json:"cif"`
	IdType                string                    `json:"idType"`
	CerificateId          string                    `json:"cerificateId"`
	CustomerName          string                    `json:"customerName"`
	CustomerStatus        string                    `json:"customerStatus"`
	Nationality           string                    `json:"nationality"`
	BankName              string                    `json:"bankName"`
	Bic                   string                    `json:"bic"`
	BankAccNum            string                    `json:"bankAccNum"`
	BankAccName           string                    `json:"bankAccName"`
	CreditCardNum         string                    `json:"creditCardNum"`
	CardName              string                    `json:"cardName"`
	ExpiryMonth           string                    `json:"expiryMonth"`
	ExpiryYear            string                    `json:"expiryYear"`
	SecurityCode          string                    `json:"securityCode"`
	CreditName            string                    `json:"creditName"`
	CreditIdType          string                    `json:"creditIdType"`
	CreditCertificateId   string                    `json:"creditCertificateId"`
	CreditNationality     string                    `json:"creditNationality"`
	CreditPhoneNumber     string                    `json:"creditPhoneNumber"`
}
type AL000075RepayPlan struct {
	RepayPeriod          int             `json:"repayPeriod"`
	InstallmentAmount    decimal.Decimal `json:"installmentAmount"`
	PrincipalPayment     decimal.Decimal `json:"principalPayment"`
	InterestPayment      decimal.Decimal `json:"interestPayment"`
	RepaymentDate        string          `json:"repaymentDate"`
	OutstandingPrincipal decimal.Decimal `json:"outstandingPrincipal"`
}
type AL000075InterestList struct {
	InterestType         string
	CustomerGrade        string
	InterestRate         decimal.Decimal
	FloatDirection       string
	FloatType            string
	FloatValue           decimal.Decimal
	ExecutedInterestRate decimal.Decimal
}

func (*AL000075Request) GetServiceKey() string {
	return "AL000075"
}
