package models

import (
	"github.com/shopspring/decimal"
)

type IC000062I struct {
	CreditLineInf   CustomerProductCreditInf `json:"creditLineInf" description:"Credit line inf"`
	LoanLinkAccount LinkAccountInf           `json:"loanLinkAccount" description:"Loan link account"`
	Guarantor       PguarantorInf            `json:"pguarantor" description:"Guarantor"`
}

type Interest struct {
	InterestType       string          `json:"interestType" description:"Interest type:O-Normal Loan 1-Overdue Loan"`
	InterestId         string          `json:"interestId" description:"Interest id"`
	InterestStrategyId string          `json:"interestStrategyId" description:"Interest strategy id"`
	BaseInterestRate   decimal.Decimal `json:"baseInterestRate" description:"Base Interest Rate"`
	FloatValue         decimal.Decimal `json:"floatValue" description:"Float value"`
	InterestRate       decimal.Decimal `json:"interestRate" description:"Interest rate"`
	FloatDirection     string          `json:"floatDirection" description:"Float direction"`
	FloatMethod        string          `json:"floatMethod" description:"Float method"`
}

type IC000062O struct {
	CreditApplySerialNum string `json:"creditApplySerialNum" description:"Credit apply serial number"`
}

type CustomerProductCreditInf struct {
	CustomerId                 string          `json:"customerId" description:"Customer id"`
	ProductId                  string          `json:"productId" description:"Product id"`
	RevolvingCreditLine        decimal.Decimal `json:"revolvingCreditLine" description:"Revolving credit line"`
	ProductCurrency            string          `json:"productCurrency" description:"Product currency"`
	CustomerCreditLineCurrency string          `json:"customerCreditLineCurrency" description:"Customer credit line currency"`
	CreditLine                 decimal.Decimal `json:"creditLine" description:"Credit line"`
	EffectiveDate              string          `json:"effectiveDate" description:"Effective date"`
	ExpiredDate                string          `json:"expiredDate" description:"Expired date"`
	ArrangementPurpose         string          `json:"arrangementPurpose" description:"Arrangement purpose"`
	TransactionUser            string          `json:"transactionUser" description:"Transaction user"`
	TransactionBran            string          `json:"transactionBran" description:"Transaction bran"`
	QuotaReleaseFlag           string          `json:"quotaReleaseFlag"`
}

type NormalLoanPricingInf struct {
	InterestRate         decimal.Decimal `json:"interestRate" description:"Interest rate"`
	FloatDirection       string          `json:"floatDirection" description:"Float direction"`
	FloatType            string          `json:"floatType" description:"Float type"`
	FloatValue           decimal.Decimal `json:"floatValue" description:"Float value"`
	ExecutedInterestRate decimal.Decimal `json:"executedInterestRate" description:"Executed interest rate"`
}

type LinkAccountInf struct {
	LocalBankFlag    string `json:"localBankFlag" description:"Local bank flag"`
	BankNum          string `json:"bankNum" description:"Bank number"`
	BankName         string `json:"bankName" description:"Bank name"`
	AccountNum       string `json:"accountNum" description:"Account number"`
	AccountName      string `json:"accountName" description:"Account name"`
	ExpireMonth      string `json:"expireMonth" description:"Expire month"`
	ExpireYear       string `json:"expireYear" description:"Expire year"`
	SecurityCode     string `json:"securityCode" description:"Security code"`
	CreditCardNumber string `json:"creditCardNumber" description:"Credit card number"`
	CardholderName   string `json:"cardholderName" description:"Cardholder name"`
}

type PguarantorInf struct {
	GuaranteeMethod         string `json:"guaranteeMethod" description:"Guarantee method"`
	GuarantorName           string `json:"guarantorName" description:"Guarantor name"`
	GuarantorRelation       string `json:"guarantorRelation"`
	GuarantorIdCountry      string `json:"guarantorIdCountry" description:"Guarantor id country"`
	GuarantorIdType         string `json:"guarantorIdType" description:"Guarantor id type"`
	GuarantorIdNum          string `json:"guarantorIdNum" description:"Guarantor id number"`
	GuarantorTelephoneNum   string `json:"guarantorTelephoneNum" description:"Guarantor telephone number"`
	GuarantorAddressCountry string `json:"guarantorAddressCountry"`
	GuarantorAddressType    string `json:"guarantorAddressType"`
	GuarantorAddress1       string `json:"guarantorAddress1"`
	GuarantorAddress2       string `json:"guarantorAddress2"`
	GuarantorAddress3       string `json:"guarantorAddress3"`
	GuarantorPostCode       string `json:"guarantorPostCode"`
}

func (*IC000062I) GetServiceKey() string {
	return "IC000062"
}
