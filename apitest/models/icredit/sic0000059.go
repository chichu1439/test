package models

import (
	"github.com/shopspring/decimal"
)

type IC000059I struct {
	CustomerId    string `json:"customerId" description:"Customer id" validate:"required"`
	RevolvingFlag string `json:"revolvingFlag" description:"Revolving flag" validate:"required"`
	ProductId     string `json:"productId" description:"Product id" validate:"required"`
}

type IC000059O struct {
	CreditLine           decimal.Decimal   `json:"creditLine" description:"Credit line"`
	CreditStatus         string            `json:"creditStatus" description:"Credit status"`
	CreditLineAvailable  decimal.Decimal   `json:"creditLineAvailable" description:"Credit line available"`
	CreditApplySerialNum string            `json:"creditApplySerialNum" description:"Credit apply serial number"`
	CustomerQuotaInfo    CustomerQuotaInfo `json:"customerQuotaInfo" description:"Customer quota information" `
	ProductQuotaInfo     ProductQuotaInfo  `json:"productQuotaInfo" description:"Product quota information" `
	LoanLinkAccount      []LoanLinkAccount `json:"loanLinkAccount" description:"Loan link account" `
	Guarantor            []Guarantor       `json:"guarantor" description:"Guarantor" `
}

type CustomerQuotaInfo struct {
	PermanentQuotaTotal     decimal.Decimal `json:"permanentQuotaTotal" description:"Permanent quota total"`
	PermanentQuotaUsed      decimal.Decimal `json:"permanentQuotaUsed" description:"Permanent quota used"`
	PermanentQuotaFrozen    decimal.Decimal `json:"permanentQuotaFrozen" description:"Permanent quota frozen"`
	TemporaryQuotaTotal     decimal.Decimal `json:"temporaryQuotaTotal" description:"Temporary quota total"`
	TemporaryQuotaUsed      decimal.Decimal `json:"temporaryQuotaUsed" description:"Temporary quota used"`
	TemporaryQuotaFrozen    decimal.Decimal `json:"temporaryQuotaFrozen" description:"Temporary quota frozen"`
	CustomerOverlimitFlag   string          `json:"customerOverlimitFlag" description:"Customer overlimit flag"`
	CustomerOverlimitAmount decimal.Decimal `json:"customerOverlimitAmount" description:"Customer overlimit amount"`
	SystemOverlimitAmount   decimal.Decimal `json:"systemOverlimitAmount" description:"System overlimit amount"`
	TemporaryQuotaLimit     decimal.Decimal `json:"temporaryQuotaLimit" description:"Temporary quota limit"`
	ExpireDate              string          `json:"expireDate" description:"Expire date"`
	Status                  string          `json:"status" description:"Status"`
}

type ProductQuotaInfo struct {
	PermanentQuotaTotal     decimal.Decimal `json:"permanentQuotaTotal" description:"Permanent quota total"`
	PermanentQuotaUsed      decimal.Decimal `json:"permanentQuotaUsed" description:"Permanent quota used"`
	PermanentQuotaFrozen    decimal.Decimal `json:"permanentQuotaFrozen" description:"Permanent quota frozen"`
	TemporaryQuotaTotal     decimal.Decimal `json:"temporaryQuotaTotal" description:"Temporary quota total"`
	TemporaryQuotaUsed      decimal.Decimal `json:"temporaryQuotaUsed" description:"Temporary quota used"`
	TemporaryQuotaFrozen    decimal.Decimal `json:"temporaryQuotaFrozen" description:"Temporary quota frozen"`
	CustomerOverlimitFlag   string          `json:"customerOverlimitFlag" description:"Customer overlimit flag"`
	CustomerOverlimitAmount decimal.Decimal `json:"customerOverlimitAmount" description:"Customer overlimit amount"`
	SystemOverlimitAmount   decimal.Decimal `json:"systemOverlimitAmount" description:"System overlimit amount"`
	TemporaryQuotaLimit     decimal.Decimal `json:"temporaryQuotaLimit" description:"Temporary quota limit"`
	ExpireDate              string          `json:"expireDate" description:"Expire date"`
	Status                  string          `json:"status" description:"Status"`
}

type LoanLinkAccount struct {
	LinkAccountUseType string `json:"linkAccountUseType" description:"Link account use type"`
	LocalBankFlag      string `json:"localBankFlag" description:"Local bank flag"`
	BankNum            string `json:"bankNum" description:"Bank number"`
	BankName           string `json:"bankName" description:"Bank name"`
	AccountNum         string `json:"accountNum" description:"Account number"`
	AccountName        string `json:"accountName" description:"Account name"`
	ExpireMonth        string `json:"expireMonth" description:"Expire month"`
	ExpireYear         string `json:"expireYear" description:"Expire year"`
	SecurityCode       string `json:"securityCode" description:"Security code"`
}

type Guarantor struct {
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

func (*IC000059I) GetServiceKey() string {
	return "IC000059"
}
