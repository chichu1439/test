// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP910024I struct {
	MemberId     string `json:"memberId" validate:"required"`
	MsgService   string `json:"msgService" `
	MsgType      string `json:"msgType" `
	MsgDirection string `json:"msgDirection" `
	MsgId        string `json:"msgId" `
	TransId      string `json:"transId" `
	MsgDate      string `json:"msgDate"`
	Currency     string `json:"currency"`
	PageNo       int    `json:"pageNo" `
	PageRecCount int    `json:"pageRecCount" `
}

type FP910024O struct {
	TotalRecord int64                    `json:"totalRecord"`
	TotalPage   int64                    `json:"totalPage"`
	Records     []FP910024ResponseRecord `json:"records"`
}

type FP910024ResponseRecord struct {
	GlobalSerialNum     string  `json:"globalSerialNum"`
	TransId             string  `json:"transId"`
	MsgId               string  `json:"msgId"`
	ClrSysRef           string  `json:"clrSysRef"`
	MsgService          string  `json:"msgService"`
	MsgType             string  `json:"msgType"`
	MsgDirection        string  `json:"msgDirection"`
	MsgDate             string  `json:"msgDate"`
	AcceptanceDatetime  string  `json:"acceptanceDatetime"`
	InterbankStlmtAmt   float64 `json:"interbankStlmtAmt"`
	InterbankStlmtCcy   string  `json:"interbankStlmtCcy"`
	InterbankStlmtDate  string  `json:"interbankStlmtDate"`
	ResultStatus        string  `json:"resultStatus"`
	ResultReason        string  `json:"resultReason"`
	ResultDetailsReason string  `json:"resultDetailsReason"`
}

func (*FP910024I) GetServiceKey() string {
	return "FP910024"
}