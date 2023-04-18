// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP912016I struct {
	PartClearCode string `json:"partClearCode" validate:"required"`
	PartName      string `json:"partName" validate:"required"`
	MsgId         string `json:"msgId"`
	TransId       string `json:"transId"`
	ClrSysRef     string `json:"clrSysRef"`
}

type FP912016O struct {
	ClrSysRef             string  `json:"clrSysRef"`
	ClrSysId              string  `json:"clrSysId"`
	TransId               string  `json:"transId"`
	MsgId                 string  `json:"msgId"`
	InstructionId         string  `json:"instructionId"`
	EndToEndId            string  `json:"endToEndId"`
	TransNo               string  `json:"transNo"`
	TransType             string  `json:"transType"`
	TransStatus           string  `json:"transStatus"`
	TransRejectReason     string  `json:"transRejectReason"`
	TransRejectReasonInfo string  `json:"transRejectReasonInfo"`
	TransDatetime         string  `json:"transDatetime"`
	TransDirection        string  `json:"transDirection"`
	PayerPartID           string  `json:"payerPartId"`
	PayerPartName         string  `json:"payerPartName"`
	PayeePartID           string  `json:"payeePartId"`
	PayeePartName         string  `json:"payeePartName"`
	Currency              string  `json:"currency"`
	Amount                float64 `json:"amount"`
	CategoryPurpose       string  `json:"categoryPurpose"`
}

func (*FP912016I) GetServiceKey() string {
	return "FP912016"
}
