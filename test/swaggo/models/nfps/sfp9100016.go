// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP910016I struct {
	PartClearCode string `json:"partClearCode" validate:"required"`
	PartName      string `json:"partName" validate:"required"`
	MsgId         string `json:"msgId" validate:"required"`
	TransId       string `json:"transId" validate:"required"`
	ClrSysRef     string `json:"clrSysRef" validate:"required"`
}

type FP910016O struct {
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
	PayerName             string  `json:"payerName"`
	PayerAcctID           string  `json:"payerAcctId"`
	PayeePartID           string  `json:"payeePartId"`
	PayeePartName         string  `json:"payeePartName"`
	PayeeName             string  `json:"payeeName"`
	PayeeAcctID           string  `json:"payeeAcctId"`
	Currency              string  `json:"currency"`
	Amount                float64 `json:"amount"`
	CategoryPurpose       string  `json:"categoryPurpose"`
}

func (*FP910016I) GetServiceKey() string {
	return "FP910016"
}
