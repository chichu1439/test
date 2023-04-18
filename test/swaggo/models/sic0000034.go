package models

import "github.com/shopspring/decimal"

type IC000034I struct {
	GlobalSerialNum   string `json:"globalSerialNum" description:"Global serial number"`
	BusinessSerialNum string `json:"businessSerialNum" description:"Business serial number"`
	LoanReceiptNum    string `json:"loanReceiptNum" description:"Loan receipt number" validate:"required"`
	CustomerId        string `json:"customerId" description:"Customer id"`
}

type IC000034O struct {
	ApplyDate               string            `json:"applyDate" description:"Apply date"`
	GlobalSerialNum         string            `json:"globalSerialNum" description:"global serial num"`
	ApplySerialNum          string            `json:"applySerialNum" description:"Apply serial number"`
	LoanReceiptNum          string            `json:"loanReceiptNum" description:"Loan receipt number"`
	CustomerId              string            `json:"customerId" description:"Customer id"`
	LoanProductNum          string            `json:"loanProductNum" description:"Loan product number"`
	TransDate               string            `json:"transDate" description:"Trans date"`
	TransTime               string            `json:"transTime" description:"Trans time"`
	TransAmount             decimal.Decimal   `json:"transAmount" description:"Trans amount"`
	TransReason             string            `json:"transReason" description:"trans reason"`
	TransPrinciple          decimal.Decimal   `json:"transPrinciple" description:"Trans principle"`
	TransInterest           decimal.Decimal   `json:"transInterest" description:"Trans interest"`
	TransFee                decimal.Decimal   `json:"transFee" description:"Trans fee"`
	FeeTotal                decimal.Decimal   `json:"feeTotal" description:"Fee Total"`
	LoanChannel             string            `json:"loanChannel" description:"Loan channel"`
	BankName                string            `json:"bankName" description:"Bank name"`
	Bic                     string            `json:"bic" description:"Bic"`
	BankAccountName         string            `json:"bankAccountName" description:"Bank account name"`
	BankAccountNumber       string            `json:"bankAccountNumber" description:"Bank account number"`
	ChannelName             string            `json:"channelName" description:"Channel name"`
	OriginalTransDate       string            `json:"originalTransDate" description:"original trans date"`
	OriginalTransSerialNum  string            `json:"originalTransSerialNum" description:"original trans serial num"`
	RepayType               string            `json:"repayType"`
	FeePaidList             []FeePaiedInfo    `json:"feePaidList" description:"Fee paid list"`
	Maker                   string            `json:"maker" description:"Maker user id"`
	MakerComment            string            `json:"makerComment" description:"Maker comment"`
	Checker                 string            `json:"checker" description:"Checker user id"`
	CheckerComment          string            `json:"checkerComment" description:"Checker comment"`
	ApproveComment          string            `json:"approveComment" description:"Checker comment"`
	ApproveView             string            `json:"approveView" description:"Apply status:2-approved 3-rejected"`
	LoanStatus              string            `json:"loanStatus" description:"Loan status"`
	NormalInterest          decimal.Decimal   `json:"normalInterest"`
	PenaltyInterest         decimal.Decimal   `json:"penaltyInterest"`
	CompoundInterest        decimal.Decimal   `json:"compoundInterest"`
	PenaltyCompoundInterest decimal.Decimal   `json:"penaltyCompoundInterest"`
	LoanTransDetail         []LoanTransDetail `json:"loanTransDetail" description:"loan trans detail"`
	TransFlowList           []LoanTransFlow   `json:"transFlowList" description:" trans flow info"` //only include cover or waive trans flow
}

type LoanTransFlow struct {
	ApplyDate               string            `json:"applyDate" description:"Apply date"`
	GlobalSerialNum         string            `json:"globalSerialNum" description:"global serial num"`
	ApplySerialNum          string            `json:"applySerialNum" description:"Apply serial number"`
	LoanReceiptNum          string            `json:"loanReceiptNum" description:"Loan receipt number"`
	CustomerId              string            `json:"customerId" description:"Customer id"`
	LoanProductNum          string            `json:"loanProductNum" description:"Loan product number"`
	TransDate               string            `json:"transDate" description:"Trans date"`
	TransTime               string            `json:"transTime" description:"Trans time"`
	TransAmount             decimal.Decimal   `json:"transAmount" description:"Trans amount"`
	TransReason             string            `json:"transReason" description:"trans reason"`
	TransPrinciple          decimal.Decimal   `json:"transPrinciple" description:"Trans principle"`
	TransInterest           decimal.Decimal   `json:"transInterest" description:"Trans interest"`
	TransFee                decimal.Decimal   `json:"transFee" description:"Trans fee"`
	FeeTotal                decimal.Decimal   `json:"feeTotal" description:"Fee Total"`
	LoanChannel             string            `json:"loanChannel" description:"Loan channel"`
	BankName                string            `json:"bankName" description:"Bank name"`
	Bic                     string            `json:"bic" description:"Bic"`
	BankAccountName         string            `json:"bankAccountName" description:"Bank account name"`
	BankAccountNumber       string            `json:"bankAccountNumber" description:"Bank account number"`
	ChannelName             string            `json:"channelName" description:"Channel name"`
	OriginalTransDate       string            `json:"originalTransDate" description:"original trans date"`
	OriginalTransSerialNum  string            `json:"originalTransSerialNum" description:"original trans serial num"`
	RepayType               string            `json:"repayType"`
	FeePaidList             []FeePaiedInfo    `json:"feePaidList" description:"Fee paid list"`
	Maker                   string            `json:"maker" description:"Maker user id"`
	MakerComment            string            `json:"makerComment" description:"Maker comment"`
	Checker                 string            `json:"checker" description:"Checker user id"`
	CheckerComment          string            `json:"checkerComment" description:"Checker comment"`
	ApproveComment          string            `json:"approveComment" description:"Checker comment"`
	ApproveView             string            `json:"approveView" description:"Apply status:2-approved 3-rejected"`
	LoanStatus              string            `json:"loanStatus" description:"Loan status"`
	NormalInterest          decimal.Decimal   `json:"normalInterest"`
	PenaltyInterest         decimal.Decimal   `json:"penaltyInterest"`
	CompoundInterest        decimal.Decimal   `json:"compoundInterest"`
	PenaltyCompoundInterest decimal.Decimal   `json:"penaltyCompoundInterest"`
	LoanTransDetail         []LoanTransDetail `json:"loanTransDetail" description:"loan trans detail"`
}
type LoanTransDetail struct {
	AmountType  string          `json:"amountType"`
	TransAmount decimal.Decimal `json:"transAmount"`
}
type FeePaiedInfo struct {
	FeeType string          `json:"feeType"`
	FeePaid decimal.Decimal `json:"feePaid"`
}

func (*IC000034I) GetServiceKey() string {
	return "IC000034"
}
