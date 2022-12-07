package models

type FP910008I struct {
	TransId    string `json:"transId"`
	MsgId      string `json:"msgId"`
	PartCode   string `json:"partCode"`
	MsgService string `json:"msgService"`
}
type FP910008O struct {
	TotalCount  int             `json:"totalCount"`
	JournalList []FP910008OItem `json:"journalList"`
}
type FP910008OItem struct {
	GlobalSerialNum     string  `json:"globalSerialNum"`
	TransId             string  `json:"transId"`
	MsgId               string  `json:"msgId"`
	ClrSysRef           string  `json:"clrSysRef"`
	EndToEndId          string  `json:"endToEndId"`
	MsgType             string  `json:"msgType"`
	MsgDefId            string  `json:"msgDefId"`
	MsgLevel            int64   `json:"msgLevel"`
	MsgDirection        string  `json:"msgDirection"`
	Message             []byte  `json:"message"`
	MsgDate             string  `json:"msgDate"`
	MsgDatetime         string  `json:"msgDatetime"`
	MsgService          string  `json:"msgService"`
	FromMemberId        string  `json:"fromMemberId"`
	ToMemberId          string  `json:"toMemberId"`
	AcceptanceDatetime  string  `json:"acceptanceDatetime"`
	InterbankStlmtAmt   float64 `json:"interbankStlmtAmt"`
	InterbankStlmtCcy   string  `json:"interbankStlmtCcy"`
	InterbankStlmtDate  string  `json:"interbankStlmtDate"`
	ResultStatus        string  `json:"resultStatus"`
	ResultReason        string  `json:"resultReason"`
	ResultDetailsReason string  `json:"resultDetailsReason"`
	CreateTime          string  `json:"createTime"`
}

func (*FP910008I) GetServiceKey() string {
	return "FP910008"
}
