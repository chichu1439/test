package models

import "github.com/shopspring/decimal"

type IC000086I struct {
	BusinessSerialNum string `json:"businessSerialNum" description:"Business serial number" validate:"required"`
	LoanReceiptNum    string `json:"loanReceiptNum" description:"Loan receipt number" validate:"required"`
	CustomerId        string `json:"customerId" description:"Customer id"`
}

type IC000086O struct {
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
	RepaymentType           string          `json:"repaymentType"`
	FeeTotal                decimal.Decimal `json:"feeTotal" description:"Fee Total"`
	Maker                   string          `json:"maker" description:"Maker user id"`
	MakerComment            string          `json:"makerComment" description:"Maker comment"`
	Checker                 string          `json:"checker" description:"Checker user id"`
	CheckerComment          string          `json:"checkerComment" description:"Checker comment"`
	ApproveComment          string          `json:"approveComment" description:"Checker comment"`
	ApproveView             string          `json:"approveView" description:"Apply status:2-approved 3-rejected"`
	RepaymentStatus         string          `json:"repaymentStatus" description:"repayment Status"`
	NormalInterest          decimal.Decimal `json:"normalInterest"`
	PenaltyInterest         decimal.Decimal `json:"penaltyInterest"`
	CompoundInterest        decimal.Decimal `json:"compoundInterest"`
	PenaltyCompoundInterest decimal.Decimal `json:"penaltyCompoundInterest"`
	DisbursementAmount      decimal.Decimal `json:"disbursementAmount"`
	DiscountPercentage      decimal.Decimal `json:"discountPercentage" description:"Discount percentage"`
	TotalOutstanding        decimal.Decimal `json:"totalOutstanding"`
	ReasonForCancelling     string          `json:"reasonForCancelling"`
	ListFee                 []ApplyFeeInfo  `json:"listFee"`
}
type ApplyFeeInfo struct {
	FeeAmount decimal.Decimal `json:"feeAmount"`
	FeeType   string          `json:"feeType"`
	ApplyType string          `json:"applyType"`
}

func (*IC000086I) GetServiceKey() string {
	return "IC000086"
}
