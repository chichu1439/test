// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP912002I struct {
	PartClearCode string `json:"partClearCode" validate:"required"`
	PartName      string `json:"partName" validate:"required"`
	TransDate     string `json:"transDate" validate:"required"`
	TransType     string `json:"transType" validate:""`
	TransStatus   string `json:"transStatus" `
	TypeId        int    `json:"typeId"`
	Currency      string `json:"currency" `
	RefNum        string `json:"refNum" `
	MsgId         string `json:"msgId"`
	InstructionId string `json:"instructionId"`
	EndToEndId    string `json:"endToEndId"`
	PageNo        int    `json:"pageNo" `
	PageRecCount  int    `json:"pageRecCount" `
}

type FP912002O struct {
	TotalRecord int                      `json:"totalRecord"`
	TotalPage   int                      `json:"totalPage"`
	Records     []FP912002ResponseRecord `json:"records"`
}

type FP912002ResponseRecord struct {
	MsgId         string  `json:"msgId"`
	InstructionId string  `json:"instructionId"`
	EndToEndId    string  `json:"endToEndId"`
	TransId       string  `json:"transId"`
	ClrSysRef     string  `json:"clrSysRef"`
	TransDate     string  `json:"transDate"`
	TransType     string  `json:"transType"`
	TransStatus   string  `json:"transStatus"`
	Currency      string  `json:"currency"`
	Amount        float64 `json:"amount"`
	MsgDirection  string  `json:"msgDirection"`
}

func (*FP912002I) GetServiceKey() string {
	return "FP912002"
}
