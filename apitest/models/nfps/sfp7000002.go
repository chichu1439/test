// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP700002I struct {
}

type FP700002O struct {
	IndirectPartClearingCode string `json:"indirectPartClearingCode"`
	Currency                 string `json:"currency"`
	DirectPartClearingCode   string `json:"directPartClearingCode"`
}
