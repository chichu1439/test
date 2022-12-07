// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP110011I struct {
	PartClearingCode string `json:"partClearingCode" validate:"required"`
	GlobalSerialNum  string `json:"globalSerialNum" validate:"required"`
	TransId          string `json:"transId" validate:"required"`
	MessageId        string `json:"messageId" validate:"required"`
}

type FP110011O struct {
}

func (*FP110011I) GetServiceKey() string {
	return "FP110011"
}

type FP110011OutgressRequest struct {
	Body []byte `json:"body"`
}

type FP110011OutgressResponse struct {
	Body []byte `json:"body"`
}

func (*FP110011OutgressRequest) GetServiceKey() string {
	return "FP210011"
}
