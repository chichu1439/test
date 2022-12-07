package models

import "github.com/shopspring/decimal"

type AL000058I struct {
	TransactionDate       string                 `validate:"required,max=8"  json:"transactionDate" description:"Transaction date"`
	LoanReceiptNum        string                 `validate:"required,max=32"  json:"loanReceiptNum" description:"Loan receipt number"`
	LocalBankFlag         string                 `json:"localBankFlag"`
	AccountingaAccountNum string                 `validate:"required,max=30"  json:"accountingaAccountNum" description:"Accountinga account number"`
	Period                int                    `validate:"required,max=6"  json:"period" description:"Period"`
	Currency              string                 `validate:"required,max=3"   json:"currency" description:"Currency"`
	CollecMethod          string                 `validate:"required,max=1"  json:"collecMethod" description:"Collec method:0-Directly Collect,1-Add to Bill"` // 0-Directly Collect,1-Add to Bill
	CollectAccount        CollectAccount         `json:"collectAccount" description:"Collect account"`
	ListFeeType           []LoanFeeCollectionReq `json:"listFeeType" description:"List fee type"`
}

type AL000058O struct {
}

type LoanFeeCollectionReq struct {
	FeeId               string          `validate:"max=10" json:"feeId" description:"Fee id"`
	FeeAmount           decimal.Decimal `json:"feeAmount" description:"Fee amount"`
	FeeStrategyId       string          `validate:"max=10" json:"feeStrategyId" description:"Fee strategy id"`
	FeeType             string          `validate:"max=2" json:"feeType" description:"Fee type"`
	BeFeeDeductionFlag  string          `validate:"required,max=1" json:"beFeeDeductionFlag" description:"Be fee deduction flag"` //
	FeeAbstractCode     string          `validate:"required,max=3" json:"feeAbstractCode" description:"Fee abstract code"`
	FeeAmortizationFlag string          `validate:"required,max=1" json:"feeAmortizationFlag" description:"Fee amortization flag"`
	AmortizePeriodType  string          `validate:"required,max=2" json:"amortizePeriodType" description:"Amortize period type"`
	AmortizationTimes   int             `validate:"max=6" json:"amortizationTimes" description:"Amortization times"`
}

type CollectAccount struct {
	LocalBankFlag string `validate:"max=1" json:"localBankFlag" description:"Local bank flag:I-Internal Account,O-External Account"` //I-Internal Account,O-External Account
	BankNum       string `validate:"max=30" json:"bankNum" description:"Bank number"`
	BankName      string `validate:"max=60" json:"bankName" description:"Bank name"`
	AccountNum    string `validate:"max=40" json:"accountNum" description:"Account number"`
	AccountName   string `validate:"max=60" json:"accountName" description:"Account name"`
	AccountType   string `validate:"max=6" json:"accountType" description:"Account type"`
}

func (*AL000058I) GetServiceKey() string {
	return "AL000058"
}
