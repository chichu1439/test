package models

import "github.com/shopspring/decimal"

type IC000060I struct {
	IdType               string `json:"idType" description:"Id type"`
	IdNumber             string `json:"idNumber" description:"Id number"`
	CustomerId           string `json:"customerId" description:"Customer Id"`
	Currency             string `json:"currency" description:"Currency"`
	CreditApplySerialNum string `json:"creditLineReferenceNumber" description:"Credit apply serial number"`
	CreditLineStatus     string `json:"creditLineStatus" description:"Credit Line Status(0-inactive;1-active;2-fozen;3-expired)"`
}

type IC000060O struct {
	CustomerProductQuotaList []CustomerProductQuotaList `json:"customerProductQuotaList" description:"Customer product quota list" `
}

type CustomerProductQuotaList struct {
	CreditApplySerialNum string          `json:"creditLineReferenceNumber" description:"Credit apply serial number"`
	Cif                  string          `json:"cif" description:"Cif"`
	CustomerName         string          `json:"customerName" description:"Customer name"`
	IdType               string          `json:"idType" description:"Id type"`
	IdNumber             string          `json:"idNumber" description:"Id number"`
	IdNation             string          `json:"idNation" description:"Id nation"`
	ProductID            string          `json:"productID" description:"Product id"`
	ProductName          string          `json:"productName" description:"Product name"`
	Currency             string          `json:"currency" description:"Currency"`
	CreditLine           decimal.Decimal `json:"creditLine" description:"Credit line"`
	CreditLineStatus     string          `json:"creditLineStatus" description:"Credit line status"`
	CustomerStatus       string          `json:"customerStatus" description:"Customer status"`
}

func (*IC000060I) GetServiceKey() string {
	return "IC000060"
}
