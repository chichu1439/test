//Version: v1
package models

type PD000114I struct {
	FeeType string `json:"feeType"`
}

type PD000114O struct {
	Records []PD000114OItem `json:"records"`
}

type PD000114OItem struct {
	FeeName        string `json:"feeName"`
	FeeType        string `json:"feeType"`
	AccountingCode string `json:"accountingCode"`
	IsAmortization string `json:"isAmortization"`
}
