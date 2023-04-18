// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP910023I struct {
	GlobalSerialNum     string  `json:"globalSerialNum" validate:"required"`
	TransId             string  `json:"transId" validate:"required"`
	MsgId               string  `json:"msgId" `
	ClrSysRef           string  `json:"clrSysRef" `
	EndToEndId          string  `json:"endToEndId" `
	MsgType             string  `json:"msgType" `
	MsgDefId            string  `json:"msgDefId" `
	MsgLevel            int64   `json:"msgLevel" `
	MsgDirection        string  `json:"msgDirection" `
	Message             []byte  `json:"message" `
	MsgDate             string  `json:"msgDate" `
	MsgDatetime         string  `json:"msgDatetime" `
	MsgService          string  `json:"msgService" `
	FromMemberId        string  `json:"fromMemberId" `
	ToMemberId          string  `json:"toMemberId" `
	AcceptanceDatetime  string  `json:"acceptanceDatetime" `
	InterbankStlmtAmt   float64 `json:"interbankStlmtAmt" `
	InterbankStlmtCcy   string  `json:"interbankStlmtCcy" `
	InterbankStlmtDate  string  `json:"interbankStlmtDate" `
	ResultStatus        string  `json:"resultStatus" `
	ResultReason        string  `json:"resultReason" `
	ResultDetailsReason string  `json:"resultDetailsReason" `
}

type FP910023O struct {
	Status string `json:"status"`
}

func (*FP910023I) GetServiceKey() string {
	return "FP910023"
}
