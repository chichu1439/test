// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP910002I struct {
	PartClearCode string `json:"partClearCode" validate:"required"`
	PartName      string `json:"partName" validate:"required"`
	TransDate     string `json:"transDate" validate:"required"`
	TransType     string `json:"transType" validate:""`
	TransId       string `json:"transId" `
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

type FP910002O struct {
	TotalRecord int                      `json:"totalRecord"`
	TotalPage   int                      `json:"totalPage"`
	Records     []FP910002ResponseRecord `json:"records"`
}

type FP910002ResponseRecord struct {
	InstructionId string  `json:"instructionId"`
	EndToEndId    string  `json:"endToEndId"`
	MsgId         string  `json:"msgId"`
	TransId       string  `json:"transId"`
	ClrSysRef     string  `json:"clrSysRef"`
	TransDate     string  `json:"transDate"`
	TransType     string  `json:"transType"`
	TransStatus   string  `json:"transStatus"`
	Currency      string  `json:"currency"`
	Amount        float64 `json:"amount"`
	MsgDirection  string  `json:"msgDirection"`
}

func (*FP910002I) GetServiceKey() string {
	return "FP910002"
}
