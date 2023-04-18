// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP910022I struct {
	GlobalSerialNum     string  `json:"globalSerialNum" ` // Deprecated
	TransId             string  `json:"transId" `         // Deprecated
	MsgId               string  `json:"msgId" validate:"required"`
	ClrSysRef           string  `json:"clrSysRef" `           // Deprecated
	EndToEndId          string  `json:"endToEndId" `          // Deprecated
	MsgType             string  `json:"msgType" `             // Deprecated
	MsgDefId            string  `json:"msgDefId" `            // Deprecated
	MsgLevel            int64   `json:"msgLevel" `            // Deprecated
	MsgDirection        string  `json:"msgDirection" `        // Deprecated
	Message             []byte  `json:"message" `             // Deprecated
	MsgDate             string  `json:"msgDate" `             // Deprecated
	MsgDatetime         string  `json:"msgDatetime" `         // Deprecated
	MsgService          string  `json:"msgService" `          // Deprecated
	FromMemberId        string  `json:"fromMemberId" `        // Deprecated
	ToMemberId          string  `json:"toMemberId" `          // Deprecated
	AcceptanceDatetime  string  `json:"acceptanceDatetime" `  // Deprecated
	InterbankStlmtAmt   float64 `json:"interbankStlmtAmt" `   // Deprecated
	InterbankStlmtCcy   string  `json:"interbankStlmtCcy" `   // Deprecated
	InterbankStlmtDate  string  `json:"interbankStlmtDate" `  // Deprecated
	ResultStatus        string  `json:"resultStatus" `        // Deprecated
	ResultReason        string  `json:"resultReason" `        // Deprecated
	ResultDetailsReason string  `json:"resultDetailsReason" ` // Deprecated
}

type FP910022O struct {
	Status string `json:"status"`
}

func (*FP910022I) GetServiceKey() string {
	return "FP910022"
}
