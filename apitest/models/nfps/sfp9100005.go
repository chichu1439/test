// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP910005I struct {
	JournalId             int     `json:"journalId" validate:"required"`
	GlobalSerialNum       string  `json:"globalSerialNum" validate:"required"`
	MsgLevel              string  `json:"msgLevel"`
	OriginalMessage       []byte  `json:"originalMessage"`
	OriginalMsgDatetime   string  `json:"originalMsgDatetime"`
	OriginalMsgId         string  `json:"originalMsgId"`
	OriginalMsgType       string  `json:"originalMsgType"`
	ReturnMessage         []byte  `json:"returnMessage"`
	ReturnMsgDatetime     string  `json:"returnMsgDatetime"`
	ReturnMsgId           string  `json:"returnMsgId"`
	ReturnMsgType         string  `json:"returnMsgType"`
	TransStateCode        string  `json:"transStateCode"`
	TransStateDesc        string  `json:"transStateDesc"`
	TransId               string  `json:"transId"`
	ClrSysRef             string  `json:"clrSysRef"`
	EndToEndId            string  `json:"endToEndId"`
	AcceptDatetime        string  `json:"acceptDatetime"`
	CrMemberId            string  `json:"crMemberId"`
	DrMemberId            string  `json:"drMemberId"`
	InterbankStlmtAmt     float64 `json:"interbankStlmtAmt"`
	InterbankStlmtCcy     string  `json:"interbankStlmtCcy"`
	InterbankStlmtDate    string  `json:"interbankStlmtDate"`
	TransResultCode       string  `json:"transResultCode"`
	TransResultReason     string  `json:"transResultReason"`
	TransResultReasonInfo string  `json:"transResultReasonInfo"`
}

func (*FP910005I) GetServiceKey() string {
	return "FP910005"
}

type FP910005O struct {
	Status string `json:"status"`
}
