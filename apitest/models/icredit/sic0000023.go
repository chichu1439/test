package models

import "github.com/shopspring/decimal"

type SIC0000023I struct {
	MakeLoanApplySerialNum string `json:"makeLoanApplySerialNum" description:"Make loan apply serial number"`
	LoanReceiptNum         string `json:"loanReceiptNum" description:"Loan receipt number"`
	CustomerId             string `json:"customerId" description:"Customer id" validate:"required"`
}

type SIC0000023O struct {
	ModifyType              string              `json:"modifyType" description:"Modify type"`
	LoanCycle               string              `json:"loanCycle" description:"Loan cycle"`
	MakeLoanApplySerialNum  string              `json:"makeLoanApplySerialNum" description:"Make loan apply serial number"`
	LoanReceiptNum          string              `json:"loanReceiptNum" description:"Loan receipt number"`
	LoanProductId           string              `json:"loanProductId" description:"Loan product id"`
	CustomerId              string              `json:"customerId" description:"Customer id"`
	LoanAmount              decimal.Decimal     `json:"loanAmount" description:"Loan amount"`
	Currency                string              `json:"currency" description:"currency of the loan"`
	PeriodNum               int                 `json:"periodNum" description:"Period number"`
	RepayMethod             string              `json:"repayMethod" description:"Repay method"`
	LoanApplyStatus         string              `json:"loanApplyStatus" description:"Loan apply status"`
	AccountingAccountStatus string              `json:"accountingAccountStatus" description:"Accounting account status:0:normal 1:overdue 2:impaired 3:write_off 4:closed"` //0:normal 1:overdue 2:impaired 3:write_off 4:closed
	CustomerName            string              `json:"customerName" description:"Customer name"`
	IdNation                string              `json:"idNation" description:"Id nation"`
	IdType                  string              `json:"idType" description:"Id type"`
	IdNum                   string              `json:"idNum" description:"Id number"`
	IdExpireDate            string              `json:"idExpireDate" description:"Id expire date"`
	CustomerStatus          string              `json:"customerStatus" description:"Customer status"`
	Marriage                string              `json:"marriage" description:"Marriage"`
	CustomerGrade           string              `json:"customerGrade" description:"Customer grade"`
	ProductName             string              `json:"productName" description:"Product name"`
	LoanPurpose             string              `json:"loanPurpose" description:"Loan purpose"`
	AppDate                 string              `json:"appDate" description:"App date"`
	DrawDate                string              `json:"drawDate" description:"Draw date"`
	MatureDate              string              `json:"matureDate" description:"Mature date"`
	RepayDay                string              `json:"repayDay" description:"Repay day"`
	DisburseStatus          string              `json:"disburseStatus" description:"Disburse status"`
	TransferChannel         string              `json:"transferChannel" description:"Transfer channel"`
	RepayCycle              string              `json:"repayCycle" description:"Repay cycle"`
	AutoReplyFlag           string              `json:"autoReplyFlag" description:"Auto reply flag"`
	Maker                   string              `json:"maker" description:"Maker"`
	MakerComment            string              `json:"makerComment" description:"Maker comment"`
	MakeLoanOrgzNum         string              `json:"makeLoanOrgzNum" description:"Make loan orgz number"`
	ApproveInfo             ApproveInfo02       `json:"approveInfo" description:"Approve information"`
	GracePeriod             int                 `json:"gracePeriod" description:"Grace period"`
	OverDueAmount           decimal.Decimal     `json:"overDueAmount" description:"Over due amount"`
	OutstandingBalance      decimal.Decimal     `json:"outstandingBalance" description:"Outstanding balance"`
	Principal               decimal.Decimal     `json:"principal" description:"Principal"`
	FixedInterestRate       decimal.Decimal     `json:"fixedInterestRate" description:"Fixed interest rate"`
	NextRepayDate           string              `json:"nextRepayDate" description:"Next repay date"`
	RiskLevel               string              `json:"riskLevel" description:"Risk level"`
	InstallAmount           decimal.Decimal     `json:"installAmount" description:"Install amount"`
	InstallPrincipal        decimal.Decimal     `json:"installPrincipal" description:"Install principal"`
	InstallInterest         decimal.Decimal     `json:"installInterest" description:"Install interest"`
	ListInterest            []MicroLoanInterest `json:"listInterest" description:"List interest"`
	ListFee                 []MicroLoanFee      `json:"listFee" description:"List fee"`
	BillRuleInof            *BillRuleInfo       `json:"billRuleInfo" description:"Bill rule information"`
	DisburseAmount          decimal.Decimal     `json:"disburseAmount"`
}

type ApproveInfo02 struct {
	MakeLoanUser        string                  `json:"makeLoanUser" description:"Make loan user"`
	MakeLoanComment     string                  `json:"makeLoanComment" description:"Make loan comment"`
	CheckerUser         string                  `json:"checkerUser" description:"Checker user"`
	CheckerComment      string                  `json:"checkerComment" description:"Checker comment"`
	CheckerReviewStatus string                  `json:"checkerReviewStatus" description:"Checker review status"`
	ChangeContextList   []LoanChangeContextList `json:"changeContextList" description:"Change context list"`
	ChangeContext       *ChangeContext          `json:"changeContext" description:"Change context"`
	ModifyType          string                  `json:"modifyType" description:"Modify type"`
	ApplyDate           string                  `json:"applyDate" description:"Apply date"`
}

func (*SIC0000023I) GetServiceKey() string {
	return "sic0000023"
}
