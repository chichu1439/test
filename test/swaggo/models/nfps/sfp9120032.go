// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP912032I struct {
	InstructionId         string `json:"instructionId"`
	EndToEndId            string `json:"endToNndId"`
	StatusCode            string `json:"statusCode"`
	TransDate             string `json:"transDate"`
	TransStatus           string `json:"transStatus"`
	TransRejectCode       string `json:"transRejectCode"`
	TransRejectReason     string `json:"transRejectReason"`
	TransRejectReasonInfo string `json:"transRejectReasonInfo"`
	SettleDate            string `json:"settleDate"`
	SettleTime            string `json:"settleTime"`
	SettleStatus          string `json:"settleStatus"`
	Behavior              string `json:"behavior"`
}

func (*FP912032I) GetServiceKey() string {
	return "FP912032"
}

type FP912032O struct {
	Status string `json:"status"`
}
