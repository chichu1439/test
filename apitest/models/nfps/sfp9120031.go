// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP912031I struct {
	MsgId                 string `json:"msgId" validate:"required"`
	StatusCode            string `json:"statusCode"`
	TransDate             string `json:"transDate"`
	TransStatus           string `json:"transStatus"`
	TransRejectCode       string `json:"transRejectCode"`
	TransRejectReason     string `json:"transRejectReason"`
	TransRejectReasonInfo string `json:"transRejectReasonInfo"`
	SettleDate            string `json:"settleDate"`
	SettleTime            string `json:"settleTime"`
	SettleStatus          string `json:"settleStatus"`
}

func (*FP912031I) GetServiceKey() string {
	return "FP912031"
}

type FP912031O struct {
	Status string `json:"status"`
}
