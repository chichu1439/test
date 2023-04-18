// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP110005I struct {
	TransId            string  `json:"transId" validate:"required"`
	ClrSysRef          string  `json:"clrSysRef" validate:"required"`
	TransType          string  `json:"transType" validate:"required"`
	PaymentPurposeCode string  `json:"paymentPurposeCode"`
	DrMemberId         string  `json:"drMemberId" validate:"required"`
	DrAcctingType      string  `json:"drAcctingType"`
	DrBatchAcctingAmt  float64 `json:"drBatchAcctingAmt"`
	CrMemberId         string  `json:"crMemberId" validate:"required"`
	CrAcctingType      string  `json:"crAcctingType"`
	CrBatchAcctingAmt  float64 `json:"crBatchAcctingAmt"`
	StlmtCcy           string  `json:"stlmtCcy" validate:"required"`
	StlmtAmt           float64 `json:"stlmtAmt" validate:"required"`
	ReversalFlag       string  `json:"reversalFlag" description:"N-normal,R-reversal"`
	// Deprecated
	BusinessType string `json:"businessType"`
}

type FP110005O struct {
	StlmtResult string `json:"stlmtResult"`
	StlmtDate   string `json:"stlmtDate"`
	StlmtTime   string `json:"stlmtTime"`
}

func (*FP110005I) GetServiceKey() string {
	return "FP110005"
}
