package models

import "github.com/shopspring/decimal"

type SIC0000028I struct {
	ModifyType             string                  `json:"modifyType" validate:"required" description:"Modify type:2-Interest Rate;3-Period Distribution,4-Loan Extension,5- Bill rule"` //1-Loan 2-Transaction Account 3-Marital Detail 4-Contact 5-Guarantor 6-Fee 7-emergencyInfo
//	ApplyComment           string                  `json:"applyComment" description:"Apply comment"`
//	LoanCycle              string                  `json:"loanCycle" description:"Loan cycle:01 day 02 week 03 half month / double week 04 month 05 double month 06 quarter 07 half year 08 year"`
	MakeLoanApplySerialNum string                  `json:"makeLoanApplySerialNum" description:"Make loan apply serial number"`
	LoanReceiptNum         string                  `json:"loanReceiptNum" description:"Loan receipt number"`
	//LoanProductId          string                  `json:"loanProductId" description:"Loan product id"`
	CustomerId             string                  `json:"customerId" description:"Customer id"`
//	LoanAmount             decimal.Decimal         `json:"loanAmount" description:"Loan amount"`
//	LoanCurrency           string                  `json:"loanCurrency" description:"Loan currency"`
	PeriodNum              int                     `json:"periodNum" description:"Period number"`
	RepayMethod            string                  `json:"repayMethod" description:"Repay method"`
	LoanApplyStatus        string                  `json:"loanApplyStatus" description:"Loan apply status:01-loan application accepted 02-loan approval passed 03-loan approval rejected 04-contract signed successfully 05-Loan in progress 06-Loan succeeded 09-Loan failed 10-closed"`
//	CustomerName           string                  `json:"customerName" description:"Customer name"`
//	IdNation               string                  `json:"idNation" description:"Id nation"`
//	IdType                 string                  `json:"idType" description:"Id type"`
//	IdNum                  string                  `json:"idNum" description:"Id number"`
//	CustomerStatus         string                  `json:"customerStatus" description:"Customer status"`
//	CustomerGrade          string                  `json:"customerGrade" description:"Customer grade"`
//	Marriage               string                  `json:"marriage" description:"Marriage"`
//	ProductName            string                  `json:"productName" description:"Product name"`
//	LoanPurpose            string                  `json:"loanPurpose" description:"Loan purpose"`
//	AppDate                string                  `json:"appDate" description:"App date"`
//	DrawDate               string                  `json:"drawDate" description:"Draw date"`
	MatureDate             string                  `json:"matureDate" description:"Mature date"`
	RepayDay               string                  `json:"repayDay" description:"Repay day"`
	DisburseStatus         string                  `json:"disburseStatus" description:"Disburse status"`
	TransferChannel        string                  `json:"transferChannel" description:"Transfer channel"`
	RepayCycle             string                  `json:"repayCycle" description:"Repay cycle:01- daily 02- weekly 04- Monthly 08- Yearly"`
	AutoReplyFlag          string                  `json:"autoReplyFlag" description:"Auto reply flag:01-Auto repayment"`
	Maker                  string                  `json:"maker" description:"Maker"`
	MakerComment           string                  `json:"makerComment" description:"Maker comment"`
	MakeLoanOrgzNum        string                  `json:"makeLoanOrgzNum" description:"Make loan orgz number"`
	GracePeriod            int                     `json:"gracePeriod" description:"Grace period"`
	OverDueAmount          decimal.Decimal         `json:"overDueAmount" description:"Over due amount"`
	OutstandingBalance     decimal.Decimal         `json:"outstandingBalance" description:"Outstanding balance"`
	Principal              decimal.Decimal         `json:"principal" description:"Principal"`
	NextRepayDate          string                  `json:"nextRepayDate" description:"Next repay date"`
	RiskLevel              string                  `json:"riskLevel" description:"Risk level"`
	InstallAmount          decimal.Decimal         `json:"installAmount" description:"Install amount"`
	InstallPrincipal       decimal.Decimal         `json:"installPrincipal" description:"Install principal"`
	InstallInterest        decimal.Decimal         `json:"installInterest" description:"Install interest"`
	ListInterest           []MicroLoanInterest     `json:"listInterest" description:"List interest"`
	ListFee                []MicroLoanFee          `json:"listFee" description:"List fee"`
	ChangeContextList      []LoanChangeContextList `json:"changeContextList" description:"Change context list" validate:"required"`
	ChangeContext          ChangeContext           `json:"changeContext" description:"Change context"`
	BillRuleInfo           *BillRuleInfo           `json:"billRuleInfo" description:"Bill rule information"`
}

type LoanChangeContextList struct {
	ChangeItem                string `json:"changeItem" description:"Change item"`
	LoanElementClassification string `json:"loanElementClassification" description:"Loan element classification"`
	BeforeValue               string `json:"beforeValue" description:"Before value"`
	AfterValue                string `json:"afterValue" description:"After value"`
}

type ChangeContext struct {
	CustomerGrade          string              `json:"customerGrade" description:"Customer grade"`
	ApplyComment           string              `json:"applyComment" description:"Apply comment"`
	LoanCycle              string              `json:"loanCycle" description:"Loan cycle"`
	MakeLoanApplySerialNum string              `json:"makeLoanApplySerialNum" description:"Make loan apply serial number"`
	LoanReceiptNum         string              `json:"loanReceiptNum" description:"Loan receipt number"`
	LoanProductId          string              `json:"loanProductId" description:"Loan product id"`
	CustomerId             string              `json:"customerId" description:"Customer id"`
	LoanAmount             decimal.Decimal     `json:"loanAmount" description:"Loan amount"`
	LoanCurrency           string              `json:"loanCurrency" description:"Loan currency"`
	PeriodNum              int                 `json:"periodNum" description:"Period number"`
	RepayMethod            string              `json:"repayMethod" description:"Repay method"`
	LoanApplyStatus        string              `json:"loanApplyStatus" description:"Loan apply status"`
	CustomerName           string              `json:"customerName" description:"Customer name"`
	IdNation               string              `json:"idNation" description:"Id nation"`
	IdType                 string              `json:"idType" description:"Id type"`
	IdNum                  string              `json:"idNum" description:"Id number"`
	CustomerStatus         string              `json:"customerStatus" description:"Customer status"`
	Marriage               string              `json:"marriage" description:"Marriage"`
	ProductName            string              `json:"productName" description:"Product name"`
	LoanPurpose            string              `json:"loanPurpose" description:"Loan purpose"`
	AppDate                string              `json:"appDate" description:"App date"`
	DrawDate               string              `json:"drawDate" description:"Draw date"`
	MatureDate             string              `json:"matureDate" description:"Mature date"`
	RepayDay               string              `json:"repayDay" description:"Repay day"`
	DisburseStatus         string              `json:"disburseStatus" description:"Disburse status"`
	TransferChannel        string              `json:"transferChannel" description:"Transfer channel"`
	RepayCycle             string              `json:"repayCycle" description:"Repay cycle"`
	AutoReplyFlag          string              `json:"autoReplyFlag" description:"Auto reply flag"`
	Maker                  string              `json:"maker" description:"Maker"`
	MakerComment           string              `json:"makerComment" description:"Maker comment"`
	MakeLoanOrgzNum        string              `json:"makeLoanOrgzNum" description:"Make loan orgz number"`
	GracePeriod            int                 `json:"gracePeriod" description:"Grace period"`
	OverDueAmount          decimal.Decimal     `json:"overDueAmount" description:"Over due amount"`
	OutstandingBalance     decimal.Decimal     `json:"outstandingBalance" description:"Outstanding balance"`
	Principal              decimal.Decimal     `json:"principal" description:"Principal"`
	FixedInterestRate      decimal.Decimal     `json:"fixedInterestRate" description:"Fixed interest rate"`
	NextRepayDate          string              `json:"nextRepayDate" description:"Next repay date"`
	RiskLevel              string              `json:"riskLevel" description:"Risk level"`
	InstallAmount          decimal.Decimal     `json:"installAmount" description:"Install amount"`
	InstallPrincipal       decimal.Decimal     `json:"installPrincipal" description:"Install principal"`
	InstallInterest        decimal.Decimal     `json:"installInterest" description:"Install interest"`
	ListInterest           []MicroLoanInterest `json:"listInterest" description:"List interest"`
	ListFee                []MicroLoanFee      `json:"listFee" description:"List fee"`
	BillRuleInfo           *BillRuleInfo       `json:"billRuleInfo" description:" Bill rule information"`
}

type SIC0000028O struct {
	MakeLoanApplySerialNum string `json:"makeLoanApplySerialNum" description:"Make loan apply serial number"`
	LoanReceiptNum         string `json:"loanReceiptNum" description:"Loan receipt number"`
}

func (*SIC0000028I) GetServiceKey() string {
	return "sic0000028"
}
