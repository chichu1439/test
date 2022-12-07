package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type IC000061I struct {
	CustomerId           string `json:"customerId" validate:"max=20" description:"Customer id"`
	ProductId            string `json:"productId" validate:"max=10" description:"Product id"`
	CreditApplySerialNum string `json:"creditApplySerialNum"validate:"max=32""`
}

type IC000061O struct {
	CreditLineInf                     CustomerProductCreditLine          `json:"creditLineInf" description:"Credit line inf"`
	NormalLoanPricing                 *NormalLoanPricing                 `json:"normalLoanPricing" description:"Normal loan pricing"`
	PenaltyInterest                   *PenaltyInterest                   `json:"penaltyInterest" description:"Penalty interest"`
	CompoundInterest                  *CompoundInterest                  `json:"compoundInterest" description:"Compound interest"`
	CompoundInterestOfPenaltyInterest *CompoundInterestOfPenaltyInterest `json:"compoundInterestOfPenaltyInterest" description:"Compound interest of penalty interest"`
	LoanLinkAccount                   LinkAccount                        `json:"loanLinkAccount" description:"Loan link account"`
	Guarantor                         Pguarantor                         `json:"pguarantor" description:"Guarantor" `
}

type CustomerProductCreditLine struct {
	CreditApplySerialNum    string          `json:"creditApplySerialNum" description:"Credit apply serial number"`
	CustomerId              string          `json:"customerId" description:"Customer id"`
	ProductId               string          `json:"productId" description:"Product id"`
	Currency                string          `json:"currency" description:"Currency"`
	CustomerQuotaType       string          `json:"customerQuotaType" description:"Customer quota type"`
	PermanentQuotaTotal     decimal.Decimal `json:"permanentQuotaTotal" description:"Permanent quota total"`
	PermanentQuotaUsed      decimal.Decimal `json:"permanentQuotaUsed" description:"Permanent quota used"`
	PermanentQuotaFrozen    decimal.Decimal `json:"permanentQuotaFrozen" description:"Permanent quota frozen"`
	PermanentQuotaStatus    string          `json:"permanentQuotaStatus" description:"Permanent quota status"`
	TemporaryQuotaTotal     decimal.Decimal `json:"temporaryQuotaTotal" description:"Temporary quota total"`
	TemporaryQuotaUsed      decimal.Decimal `json:"temporaryQuotaUsed" description:"Temporary quota used"`
	TemporaryQuotaFrozen    decimal.Decimal `json:"temporaryQuotaFrozen" description:"Temporary quota frozen"`
	TemporaryQuotaStatus    string          `json:"temporaryQuotaStatus" description:"Temporary quota status"`
	CustomerOverlimitFlag   string          `json:"customerOverlimitFlag" description:"Customer overlimit flag"`
	CustomerOverlimitAmount decimal.Decimal `json:"customerOverlimitAmount" description:"Customer overlimit amount"`
	CustomerOverlimitMax    decimal.Decimal `json:"customerOverlimitMax" description:"Customer overlimit max"`
	SystemOverlimitAmount   decimal.Decimal `json:"systemOverlimitAmount" description:"System overlimit amount"`
	TemporaryQuotaLimit     decimal.Decimal `json:"temporaryQuotaLimit" description:"Temporary quota limit"`
	EffectiveDate           string          `json:"effectiveDate" description:"Effective date"`
	ExpireDate              string          `json:"expireDate" description:"Expire date"`
	UpperAdjustmentTimes    int             `json:"upperAdjustmentTimes" description:"Upper adjustment times"`
	LowerAdjustmentTimes    int             `json:"lowerAdjustmentTimes" description:"Lower adjustment times"`
	ExtensionTimes          int             `json:"extensionTimes" description:"Extension times"`
	Status                  string          `json:"status" description:"Status"`
	ApproveStatus           string          `json:"approveStatus" description:"Approve status"`
	CreateTime              time.Time       `json:"createTime" description:"Create time"`
	FinalModifyOrgz         string          `json:"finalModifyTime" description:"Final modify organization"`
	FinalModifyUser         string          `json:"finalModifyUser" description:"Final modify user"`
	ArrangementPurpose      string          `json:"arrangementPurpose" description:"Arrangement purpose"`
	AvailableCreditLine     decimal.Decimal `json:"availableCreditLine" description:"Available credit line"`
	RevolvingFlag           string          `json:"revolvingFlag" description:"Revolving flag"`
	CreditLineUsed          decimal.Decimal `json:"creditLineUsed" description:"CreditLineUsed"`
}

type NormalLoanPricing struct {
	InterestRate         decimal.Decimal `json:"interestRate" description:"Interest rate"`
	FloatDirection       string          `json:"floatDirection" description:"Float direction"`
	FloatType            string          `json:"floatType" description:"Float type"`
	FloatValue           decimal.Decimal `json:"floatValue" description:"Float value"`
	ExecutedInterestRate decimal.Decimal `json:"executedInterestRate" description:"Executed interest rate"`
}

type PenaltyInterest struct {
	InterestRate         decimal.Decimal `json:"interestRate" description:"Interest rate"`
	FloatDirection       string          `json:"floatDirection" description:"Float direction"`
	FloatType            string          `json:"floatType" description:"Float type"`
	FloatValue           decimal.Decimal `json:"floatValue" description:"Float value"`
	ExecutedInterestRate decimal.Decimal `json:"executedInterestRate" description:"Executed interest rate"`
}

type CompoundInterest struct {
	InterestRate         decimal.Decimal `json:"interestRate" description:"Interest rate"`
	FloatDirection       string          `json:"floatDirection" description:"Float direction"`
	FloatType            string          `json:"floatType" description:"Float type"`
	FloatValue           decimal.Decimal `json:"floatValue" description:"Float value"`
	ExecutedInterestRate decimal.Decimal `json:"executedInterestRate" description:"Executed interest rate"`
}

type CompoundInterestOfPenaltyInterest struct {
	InterestRate         decimal.Decimal `json:"interestRate" description:"Interest rate"`
	FloatDirection       string          `json:"floatDirection" description:"Float direction"`
	FloatType            string          `json:"floatType" description:"Float type"`
	FloatValue           decimal.Decimal `json:"floatValue" description:"Float value"`
	ExecutedInterestRate decimal.Decimal `json:"executedInterestRate" description:"Executed interest rate"`
}

type LinkAccount struct {
	LinkAccountUseType    string `json:"linkAccountUseType" description:"Link account use type"`
	LocalBankFlag         string `json:"localBankFlag" description:"Local bank flag"`
	BankNum               string `json:"bankNum" description:"Bank number"`
	BankName              string `json:"bankName" description:"Bank name"`
	AccountNum            string `json:"accountNum" description:"Account number"`
	AccountName           string `json:"accountName" description:"Account name"`
	ExpireMonth           string `json:"expireMonth" description:"Expire month"`
	ExpireYear            string `json:"expireYear" description:"Expire year"`
	SecurityCode          string `json:"securityCode" description:"Security code"`
	CreditCardNumber      string `json:"creditCardNumber" description:"Credit card number"`
	CardholderName        string `json:"cardholderName" description:"Cardholder name"`
	LinkAccountUsePurpose string `json:"linkAccountUsePurpose" description:"Link account use purpose"`
}

type Pguarantor struct {
	GuarantorName           string `json:"guarantorName" description:"Guarantor name"`
	GuarantorRelation       string `json:"guarantorRelation" description:"Guarantor relation"`
	GuarantorIdCountry      string `json:"guarantorIdCountry" description:"Guarantor id country"`
	GuarantorIdType         string `json:"guarantorIdType" description:"Guarantor id type"`
	GuarantorIdNum          string `json:"guarantorIdNum" description:"Guarantor id number"`
	GuarantorTelephoneNum   string `json:"guarantorTelephoneNum" description:"Guarantor telephone number"`
	GuarantorAddressCountry string `json:"guarantorAddressCountry" description:"Guarantor address country"`
	GuarantorAddressType    string `json:"guarantorAddressType" description:"Guarantor address type"`
	GuarantorAddress1       string `json:"guarantorAddress1" description:"Guarantor address1"`
	GuarantorAddress2       string `json:"guarantorAddress2" description:"Guarantor address2"`
	GuarantorAddress3       string `json:"guarantorAddress3" description:"Guarantor address3"`
	GuarantorPostCode       string `json:"guarantorPostCode" description:"Guarantor post code"`
}

func (*IC000061I) GetServiceKey() string {
	return "IC000061"
}
