package models

import "github.com/shopspring/decimal"

type IC000001I struct {
	LoanReceiptNum            string          `json:"loanReceiptNum" description:"Loan receipt number" validate:"required"`
	TransAmount               decimal.Decimal `json:"transAmount" description:"Trans amount" validate:"required"`
	LoanStatus                string          `json:"loanStatus" validate:"required" description:"Loan Status"`
	LoanChannel               string          `json:"loanChannel" description:"Loan channel"`
	BankName                  string          `json:"bankName" description:"Bank name"`
	CounterpartyAccountName   string          `json:"counterpartyAccountName" description:"Counterparty account name"`
	CounterpartyAccountNum    string          `json:"counterpartyAccountNum" description:"Counterparty account number"`
	CounterpartyBankName      string          `json:"counterpartyBankName" description:"Counterparty bank name"`
	CounterpartyBankNum       string          `json:"counterpartyBankNum" description:"Counterparty bank number"`
	CounterpartyLocalBankFlag string          `json:"counterpartyLocalBankFlag" description:"Counterparty local bank flag"`
	CounterpartyTransDate     string          `json:"counterpartyTransDate" description:"Counterparty trans date"`
	ChannelName               string          `json:"channelName" description:"Channel name"`
	TransSerialNum            string          `json:"transSerialNum" description:"Trans serial number"`
	TransType                 string          `json:"transType" description:"Trans type:03-RPYM 04-DISB 05-WTOF"`
	TransDescribe             string          `json:"transDescribe" description:"Trans describe:1.Repayment 2.Disbursement"`
	MakerUserId               string          `json:"makerUserId" description:"Maker user id"`
	MakerComment              string          `json:"makerComment" description:"Maker comment"`
	ApplyStatus               string          `json:"applyStatus" description:"Apply status"`
}

type FileListBo struct {
	FileId   string `json:"fileId" description:"File id"`
	FileName string `json:"fileName" description:"File name"`
}

type IC000001O struct {
	Status         string `json:"status" description:"Status"`
	ApplySerialNum string `json:"applySerialNum" description:"Apply serial number"`
}

func (*IC000001I) GetServiceKey() string {
	return "IC000001"
}
