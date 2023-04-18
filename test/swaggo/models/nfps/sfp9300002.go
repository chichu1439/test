// Package models will define request and response message struct
// Version: v0.0.1
package models

//清算流水状态更新
type FP930002I struct {
	TransPart      string  `json:"transPart" validate:"required"`
	Currency       string  `json:"currency" validate:"required"`
	DebCrtFlag     string  `json:"transDirection" validate:"required"`
	StlmtDate      string  `json:"stlmtDate" validate:"required"`
	StlmtTime      string  `json:"stlmtTime" validate:"required"`
	AccountingType string  `json:"accountingType" validate:"required"`
	TransAmount    float64 `json:"transAmount" validate:"required"`
	Id             int     `json:"id"` //last netting min id
	MaxId          int     `json:"maxId"`
}

type FP930002O struct {
	Status string `json:"status"`
}

func (*FP930002I) GetServiceKey() string {
	return "FP930002"
}
