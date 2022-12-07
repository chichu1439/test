// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP910012I struct {
	TransId      string  `json:"transId" `
	DrMemberId   string  `json:"drMemberId" `
	CrMemberId   string  `json:"crMemberId" `
	StlmtCcy     string  `json:"stlmtCcy" `
	StlmtAmt     float64 `json:"stlmtAmt" `
	BusinessType string  `json:"businessType" `
}

type FP910012O struct {
	StlmtResult string `json:"stlmtResult"`
	StlmtDate   string `json:"stlmtDate"`
	StlmtTime   string `json:"stlmtTime"`
}

func (*FP910012I) GetServiceKey() string {
	return "FP910012"
}
