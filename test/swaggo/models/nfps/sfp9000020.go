package models

//
type FP900020I struct {
	KeyTraceID string
	Records    []FP900020IItem `json:"records"`
}
type FP900020IItem struct {
	Sequence    int     `json:"sequence"`
	PartId      string  `json:"partId"`
	Currency    string  `json:"currency"`
	TransId     string  `json:"transId"`
	SettAcctNo  string  `json:"settAcctNo"`
	TransDate   string  `json:"transDate"`
	TransTime   string  `json:"transTime"`
	CrDrSign    string  `json:"crDrSign"`
	TransType   string  `json:"transType"`
	TransAmt    float64 `json:"transAmt"`
	AcctBalance float64 `json:"acctBalance"`
	FrMmbId     string  `json:"frMmbId"`
	ToMmbId     string  `json:"toMmbId"`
	TccState    int     `json:"tccState"`
}
type FP900020O struct {
}

func (*FP900020I) GetServiceKey() string {
	return "FP900020"
}
