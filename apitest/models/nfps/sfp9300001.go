// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP930001I struct {
	TransId            string  `json:"transId" validate:"required"`
	ClrSysRef          string  `json:"clrSysRef" validate:"required"`
	NettingId          string  `json:"nettingId"`
	TransType          string  `json:"transType" validate:"required"`
	PaymentPurposeCode string  `json:"paymentPurposeCode"`
	TransDirection     string  `json:"transDirection" validate:"required"`
	TransPart          string  `json:"transPart" validate:"required"`
	TransAcctNo        string  `json:"transAcctNo" `
	CounterpartyPart   string  `json:"counterpartyPart" `
	CounterpartyAcctNo string  `json:"counterpartyAcctNo" `
	Currency           string  `json:"currency" validate:"required"`
	TransAmount        float64 `json:"transAmount" validate:"required"`
	AccountingType     string  `json:"accountingType" validate:"required"`
	AccountingStatus   string  `json:"accountingStatus" `
	StlmtDate          string  `json:"stlmtDate" validate:"required"`
	StlmtTime          string  `json:"stlmtTime" validate:"required"`
}

type FP930001O struct {
	Status string `json:"status" `
}

func (*FP930001I) GetServiceKey() string {
	return "FP930001"
}
