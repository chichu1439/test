// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP913031Request struct {
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
	MessageContent        []byte `json:"messageContent"`
}

func (*FP913031Request) GetServiceKey() string {
	return "FP913031"
}

type FP913031Response struct {
	Status string `json:"status"`
}
