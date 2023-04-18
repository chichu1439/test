// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP110008I struct {
	TransId            string  `json:"transId" validate:"required"`
	ClrSysRef          string  `json:"clrSysRef" validate:"required"`
	TransType          string  `json:"transType" validate:"required"`
	PaymentPurposeCode string  `json:"paymentPurposeCode"`
	DrMemberId         string  `json:"drMemberId" validate:"required"`
	CrMemberId         string  `json:"crMemberId" validate:"required"`
	StlmtCcy           string  `json:"stlmtCcy" validate:"required"`
	StlmtAmt           float64 `json:"stlmtAmt" validate:"required"`
	// Deprecated
	BusinessType string `json:"businessType"`
}

type FP110008O struct {
	StlmtResult string `json:"stlmtResult"`
	StlmtDate   string `json:"stlmtDate"`
	StlmtTime   string `json:"stlmtTime"`
}

func (*FP110008I) GetServiceKey() string {
	return "FP110008"
}
