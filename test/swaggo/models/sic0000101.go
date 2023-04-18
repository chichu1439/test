package models

import "github.com/shopspring/decimal"

type SIC0000101I struct {
	LoanReceiptNum string `json:"loanReceiptNum" validate:"required"`
	CustomerId     string `json:"customerId" validate:"required"`
}

type SIC0000101O struct {
	Currency              string                    `json:"currency"`
	LoanReceiptNum        string                    `json:"loanReceiptNum"`
	LoanStatus            string                    `json:"loanStatus"`
	LoanPurpose           string                    `json:"loanPurpose"`
	ProductId             string                    `json:"productId"`
	RepayMethod           string                    `json:"repayMethod"`
	FixMinimumAmount      decimal.Decimal           `json:"fixMinimumAmount"`
	PercentageOfPrincipal decimal.Decimal           `json:"percentageOfPrincipal"`
	RepayCycle            string                    `json:"repayCycle"`
	BillDay               string                    `json:"billDay"`
	RepayDay              string                    `json:"repayDay"`
	SameDayTreatment      string                    `json:"sameDayTreatment"`
	GraceType             string                    `json:"graceType"`
	GracePeriod           int                       `json:"gracePeriod"`
	DrawDate              string                    `json:"drawDate"`
	MatureDate            string                    `json:"matureDate"`
	ListInterest          []SIC0000101OInterestList `json:"listInterest"`
	LoanAmount            decimal.Decimal           `json:"loanAmount"`
	PrincipalAmount       decimal.Decimal           `json:"principalAmount"`
	ContractAmount        decimal.Decimal           `json:"contractAmount"`
	CreditLineStatus      string                    `json:"creditLineStatus"`
	AvailableCreditLine   decimal.Decimal           `json:"availableCreditLine"`
	BankName              string                    `json:"bankName"`
	Bic                   string                    `json:"bic"`
	BankAccountNumber     string                    `json:"bankAccountNumber"`
	BankAccountName       string                    `json:"bankAccountName"`
	CreditCardNumber      string                    `json:"creditCardNumber"`
	CardHoldName          string                    `json:"cardHoldName"`
	ExpireMonth           string                    `json:"expireMonth"`
	ExpireYear            string                    `json:"expireYear"`
	SecurityCode          string                    `json:"securityCode"`
	ContractList          []SIC0000101OContractList `json:"contractList"`
	GuaranteeInfo         SIC0000023GuaranteeInfo   `json:"guaranteeInfo"`
	CustomerId            string                    `json:"customerId"`
	PeriodNum             int                       `json:"periodNum"`
	CurrentPeriod         int                       `json:"currentPeriod"`
	BillCycle             string                    `json:"billCycle"`
	BillRepaymentType     string                    `json:"billRepaymentType"`
	RestructuringFlag     string                    `json:"restructuringFlag"`
	LastRestructuringDate string                    `json:"lastRestructuringDate"`
	Maker                 string                    `json:"maker"`
	NextRepayDate         string                    `json:"nextRepayDate"`
	ListFee               []FeeInfo                 `json:"listFee"`
}
type SIC0000101OInterestList struct {
	InterestType     string          `json:"interestType"`
	CustomerGrade    string          `json:"customerGrade"`
	InterestRate     decimal.Decimal `json:"interestRate"`
	FloatDirection   string          `json:"floatDirection"`
	FloatMethod      string          `json:"floatMethod"`
	FloatValue       decimal.Decimal `json:"floatValue"`
	BaseInterestRate decimal.Decimal `json:"baseInterestRate"`
	YearDays         string          `json:"yearDays"`
}
type SIC0000101OContractList struct {
	GenerateDate string `json:"generateDate"`
	Version      string `json:"version"`
	FileKey      string `json:"fileKey"`
}

func (*SIC0000101I) GetServiceKey() string {
	return "IC000101"
}
