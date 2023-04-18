package models

import "github.com/shopspring/decimal"

type SIC0000023I struct {
	MakeLoanApplySerialNum string `json:"makeLoanApplySerialNum" description:"Make loan apply serial number"`
	LoanReceiptNum         string `json:"loanReceiptNum" description:"Loan receipt number"`
	CustomerId             string `json:"customerId" description:"Customer id" validate:"required"`
}

type SIC0000023O struct {
	// customer
	CustomerId   string `json:"customerId" description:"Customer id"`
	CustomerName string `json:"customerName" description:"Customer name"`
	// CustomerStatus string `json:"customerStatus" description:"Customer status"`
	IdNation string `json:"idNation" description:"Id nation"`
	IdType   string `json:"idType" description:"Id type"`
	IdNum    string `json:"idNum" description:"Id number"`
	// product
	ProductId   string `json:"productId"`
	ProductName string `json:"productName" description:"Product name"`
	Currency    string `json:"currency" description:"currency of the loan"`
	ProductType string `json:"productType"`
	// transaction account
	BankName          string `json:"bankName"`
	Bic               string `json:"bic"`
	BankAccountName   string `json:"bankAccountName"`
	BankAccountNumber string `json:"bankAccountNumber"`
	CreditCardNumber  string `json:"creditCardNumber"`
	CardHoldName      string `json:"cardHoldName"`
	ExpireMonth       string `json:"expireMonth"`
	ExpireYear        string `json:"expireYear"`
	SecurityCode      string `json:"securityCode"`
	// customer-guarantee
	GuaranteeMethod string                  `json:"guaranteeMethod"`
	GuaranteeInfo   SIC0000023GuaranteeInfo `json:"guaranteeInfo"`
	// loan Amount
	PrincipalAmount decimal.Decimal `json:"principalAmount"`
	ContractAmount  decimal.Decimal `json:"contractAmount" description:"Loan amount"`
	// StampDuty          decimal.Decimal `json:"stampDuty"`
	// FrontEndFee        decimal.Decimal `json:"frontEndFee"`
	DisbursementAmount decimal.Decimal `json:"disbursementAmount"`
	// PPI                decimal.Decimal
	// loan account
	LoanPurpose            string `json:"loanPurpose" description:"Loan purpose"`
	LoanReceiptNum         string `json:"loanReceiptNum" description:"Loan receipt number"`
	MakeLoanApplySerialNum string `json:"makeLoanApplySerialNum" description:"Make loan apply serial number"`
	// repayment
	RepayMethod           string          `json:"repayMethod" description:"Repay method"`
	BillDay               string          `json:"billDay" description:"Bill day"`
	GracePeriod           int             `json:"gracePeriod" description:"Grace period"`
	SameDayTreatment      string          `json:"sameDayTreatment" description:"Bill day and repayment day in the same day(Y-same date;N-repaymen date in the next month of bill date)"`
	BillCycle             string          `json:"billCycle" description:"Bill cycle(04-by month)"`
	BillRepaymentType     string          `json:"billRepaymentType" description:"Bill repayment type(0- Percentage of Principal；1- Free Minimum Amount；2- Interest Minimum Amount；3- Remaining Outstanding Amount；)"`
	PercentageOfPrincipal decimal.Decimal `json:"percentageOfPrincipal" description:"Percentage of principal"`
	FixMinimumAmount      decimal.Decimal `json:"fixMinimumAmount" description:"Fix minimum amount"`
	RepayDay              string          `json:"repayDay" description:"repay date"`
	// date
	PeriodNum  int    `json:"periodNum" description:"Period number"`
	AppDate    string `json:"appDate" description:"App date"`
	DrawDate   string `json:"drawDate" description:"Draw date"`
	MatureDate string `json:"matureDate" description:"Mature date"`
	// interest pricing
	ListInterest []MicroLoanInterest `json:"listInterest" description:"List interest"`
	// maker and maker common
	Maker             string    `json:"maker" description:"Maker"`
	MakerComment      string    `json:"makerComment" description:"Maker comment"`
	Checker           string    `json:"checker"`
	ApproveComment    string    `json:"approveComment"`
	ApproveView       string    `json:"approveView"`
	RepayCycle        string    `json:"repayCycle" description:"Repay cycle"`
	LoanApplyStatus   string    `json:"loanApplyStatus" description:"Loan apply status"`
	RestructuringFlag string    `json:"restructuringFlag"`
	LoanCurrency      string    `json:"loanCurrency"`
	ListFee           []FeeInfo `json:"listFee"`
	GraceType         string    `json:"graceType"`
}

type SIC0000023GuaranteeInfo struct {
	Name        string `json:"name"`
	IdType      string `json:"idType"`
	IdNation    string `json:"idNation"`
	IdNum       string `json:"idNum"`
	PhoneNumber string `json:"phoneNumber"`
}

func (*SIC0000023I) GetServiceKey() string {
	return "IC000023"
}
