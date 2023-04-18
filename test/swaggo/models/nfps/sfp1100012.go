// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP110012I struct {
	PartClearingCode string  `json:"partClearingCode"`
	FromMemberId     string  `json:"fromMemberId"`
	ToMemberId       string  `json:"toMemberId"`
	MsgId            string  `json:"msgId"`
	ClrSysRef        string  `json:"clrSysRef"`
	TransId          string  `json:"transId"`
	EvtParam         string  `json:"evtParam"`
	EvtTm            string  `json:"evtTm"`
	Currency         string  `json:"currency"`
	Balance          float64 `json:"balance"`
	ReSend           string  `json:"reSend"` //Y-重发
	MsgService       string  `json:"msgService"`
}

//BBLT Available balance below threshold
//Noti. Acc. Type	4!A	Indicating the type of account in payment notification It could be one of the following values:
// SACC – FPS settlement account
// CPAC – CP’s account with its SP
type FP110012O struct {
}

func (*FP110012I) GetServiceKey() string {
	return "FP110012"
}

type FP110012OutgressRequest struct {
	Body []byte `json:"body"`
}

type FP110012OutgressResponse struct {
	Body []byte `json:"body"`
}

func (*FP110012OutgressRequest) GetServiceKey() string {
	return "FP210012"
}
