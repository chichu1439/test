package models

type FP912007I struct {
	TransId      string `json:"transId"`
	MsgId        string `json:"msgId"`
	PartCode     string `json:"partCode"`
	PageNo       int    `json:"pageNo" `
	PageRecCount int    `json:"pageRecCount" `
}
type FP912007O struct {
	TotalRecCount int             `json:"totalRecCount"`
	TotalPage     int             `json:"totalPage"`
	Records       []FP912007OItem `json:"records"`
}
type FP912007OItem struct {
	Id             int64   `json:"id"`
	MsgId          string  `json:"msgId"`
	MsgType        string  `json:"msgType"`
	TransId        string  `json:"transId"`
	InstructionId  string  `json:"instructionId"`
	EndToEndId     string  `json:"endToEndId"`
	TransNo        string  `json:"transNo"`
	CreateDatetime string  `json:"createDatetime"`
	TransCurrency  string  `json:"transCurrency"`
	TransAmt       float64 `json:"transAmt"`
	TransType      string  `json:"transType"`
	CreateUnixTime int64   `json:"createUnixTime"`
	RecordType     string  `json:"recordType"`
	CreateTime     string  `json:"createTime"`
	UpdateTime     string  `json:"updateTime"`
}

func (*FP912007I) GetServiceKey() string {
	return "FP912007"
}
