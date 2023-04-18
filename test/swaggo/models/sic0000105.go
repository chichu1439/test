package models

import "github.com/shopspring/decimal"

type SIC0000105I struct {
	RestructuringApplySerialNum string `json:"restructuringApplySerialNum"`
	LoanReceiptNum              string `json:"loanReceiptNum"`
	CustomerId                  string `json:"customerId"`
	OprType                     string `json:"oprType" description:"1-restructuring;2-refinance"`
	PageNum                     int    `json:"pageNum" description:"Page number"`
	PageRecordCount             int    `json:"pageRecordCount" description:"Page record count"`
}

type SIC0000105O struct {
	PageNum         int               `json:"pageNum" description:"Page number"`
	PageRecordCount int               `json:"pageRecordCount" description:"Page record count"`
	TotalCount      int               `json:"totalCount" description:"Total count"`
	Records         []SIC0000105OList `json:"records"`
}

type SIC0000105OList struct {
	ApplyTime                   string             `json:"applyTime"`
	ApproveView                 string             `json:"approveView"`
	Maker                       string             `json:"maker"`
	MakerComment                string             `json:"makerComment"`
	ApplyDate                   string             `json:"applyDate"`
	Checker                     string             `json:"checker"`
	ApproveComment              string             `json:"approveComment"`
	ApproveDate                 string             `json:"approveDate"`
	OprType                     string             `json:"oprType" description:"1-restructuring;2-refinance"`
	CustomerId                  string             `json:"customerId"`
	CustomerName                string             `json:"customerName"`
	LoanReceiptNum              string             `json:"loanReceiptNum"`
	OriginalLoanReceiptNum      string             `json:"originalLoanReceiptNum"`
	ProductId                   string             `json:"productId"`
	ProductName                 string             `json:"productName"`
	PrincipalWrittenOff         decimal.Decimal    `json:"principalWrittenOff"`
	RestructuringApplySerialNum string             `json:"restructuringApplySerialNum"`
	RestructuringMethodList     string             `json:"restructuringMethodList"`
	RefinancingPeriod           int                `json:"refinancingPeriod"`
	RefinancingPrincipal        decimal.Decimal    `json:"refinancingPrincipal"`
	RestructuredLoanAmount      decimal.Decimal    `json:"restructuredLoanAmount"`
	RestructuredPrincipalAmount decimal.Decimal    `json:"restructuredPrincipalAmount"`
	NeedRepayTotalAmount        decimal.Decimal    `json:"needRepayTotalAmount"`
	ChangeContextList           []SIC0000105OList2 `json:"changeContextList"`
	TotalInterest               decimal.Decimal    `json:"totalInterest"`
	NormalInterest              decimal.Decimal    `json:"normalInterest"`
	PenaltyInterest             decimal.Decimal    `json:"penaltyInterest"`
	CompoundInterest            decimal.Decimal    `json:"compoundInterest"`
	PenaltyCompoundInterest     decimal.Decimal    `json:"penaltyCompoundInterest"`
	CollectionFee               decimal.Decimal    `json:"collectionFee"`
	TotalWaivedAmount           decimal.Decimal    `json:"totalWaivedAmount"`
	StampDutyFee                decimal.Decimal    `json:"stampDutyFee"`
	FrontEndFee                 decimal.Decimal    `json:"frontEndFee"`
	DisbursementAmount          decimal.Decimal    `json:"disbursementAmount"`
	IdType                      string             `json:"idType"`
	IdNum                       string             `json:"idNum"`
	IdNation                    string             `json:"idNation"`
	LoanPurpose                 string             `json:"loanPurpose"`
	NewPrincipalAmount          decimal.Decimal    `json:"newPrincipalAmount"`
	MaturityDate                string             `json:"maturityDate"`
	OutstandingPrincipal        decimal.Decimal    `json:"outstandingPrincipal"`
}
type SIC0000105OList2 struct {
	ChangeElement   string `json:"changeElement"`
	ElementOldValue string `json:"elementOldValue"`
	ElementNewValue string `json:"elementNewValue"`
}

func (*SIC0000105I) GetServiceKey() string {
	return "IC000105"
}
