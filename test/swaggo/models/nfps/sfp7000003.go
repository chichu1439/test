// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP700003I struct {
}

type FP700003O struct {
	PartClearingCode string
	Currency         string
	BeginningBalance float64
	EndingBalance    float64
	SettAcctNo       string
}
