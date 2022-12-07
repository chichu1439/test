package models

import "github.com/shopspring/decimal"

type IC000034I struct {
	BusinessSerialNum      string `json:"businessSerialNum" description:"Business serial number" validate:"required"`
	LoanReceiptNum         string `json:"loanReceiptNum" description:"Loan receipt number" validate:"required"`
	ApprovalStatus         string `json:"approvalStatus" description:"Approval status"`
	MakeLoanApplySerialNum string `json:"makeLoanApplySerialNum" description:"Make loan apply serial number"`
	CustomerId             string `json:"customerId" description:"Customer id"`
}

type IC000034O struct {
	ApplyDate               string          `json:"applyDate" description:"Apply date"`
	ApplySerialNum          string          `json:"applySerialNum" description:"Apply serial number"`
	LoanReceiptNum          string          `json:"loanReceiptNum" description:"Loan receipt number"`
	CustomerId              string          `json:"customerId" description:"Customer id"`
	LoanProductNum          string          `json:"loanProductNum" description:"Loan product number"`
	TransDate               string          `json:"transDate" description:"Trans date"`
	TransTime               string          `json:"transTime" description:"Trans time"`
	TransAmount             decimal.Decimal `json:"transAmount" description:"Trans amount"`
	TransPrinciple          decimal.Decimal `json:"transPrinciple" description:"Trans principle"`
	TransInterest           decimal.Decimal `json:"transInterest" description:"Trans interest"`
	LoanChannel             string          `json:"loanChannel" description:"Loan channel"`
	BankName                string          `json:"bankName" description:"Bank name"`
	Bic                     string          `json:"bic" description:"Bic"`
	BankAccountName         string          `json:"bankAccountName" description:"Bank account name"`
	BankAccountNumber       string          `json:"bankAccountNumber" description:"Bank account number"`
	ChannelName             string          `json:"channelName" description:"Channel name"`
	TransSerialNum          string          `json:"transSerialNum" description:"Trans serial number"`
	TransType               string          `json:"transType" description:"Trans type:03-RPYM 04-DISB 05-WTOF"`
	TransDescribe           string          `json:"transDescribe" description:"Trans describe:1.Repayment 2.Disbursement"`
	FrontEndFee             decimal.Decimal `json:"frontEndFee" description:"Front End Fee"`
	StampDutyFee            decimal.Decimal `json:"stampDutyFee" description:"Stamp Duty"`
	CollectionFee           decimal.Decimal `json:"collectionFee" description:"Collection Fee"`
	PrepaymentFee           decimal.Decimal `json:"prepaymentFee" description:"Prepayment Fee"`
	FeeTotal                decimal.Decimal `json:"feeTotal" description:"Fee Total"`
	Tax                     string          `json:"tax" description:"Tax"`
	MakerUserId             string          `json:"makerUserId" description:"Maker user id"`
	MakerComment            string          `json:"makerComment" description:"Maker comment"`
	CheckerUserId           string          `json:"checkerUserId" description:"Checker user id"`
	CheckerComment          string          `json:"checkerComment" description:"Checker comment"`
	ApplyStatus             string          `json:"applyStatus" description:"Apply status:2-approved 3-rejected"`
	LoanStatus              string          `json:"loanStatus" description:"Loan status"`
	NormalInterest          decimal.Decimal `json:"normalInterest"`
	PenaltyInterest         decimal.Decimal `json:"penaltyInterest"`
	CompoundInterest        decimal.Decimal `json:"compoundInterest"`
	PenaltyCompoundInterest decimal.Decimal `json:"penaltyCompoundInterest"`
	DisbursementAmount      decimal.Decimal `json:"disbursementAmount"`
}

func (*IC000034I) GetServiceKey() string {
	return "IC000034"
}
