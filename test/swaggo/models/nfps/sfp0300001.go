package models

// 轧差金额，更新余额及清算流水状态

type FP030001I struct {
	TransPart      string  `json:"partClearingCode" description:"Trans part"`
	StlmtDate      string  `json:"stlmtDate" description:"Stlmt date"`
	Currency       string  `json:"currency" description:"Currency"`
	TransDirection string  `json:"debCrtFlag" description:"Trans direction"`
	TransAmount    float64 `json:"transAmount" description:"Trans amount"`
	AccountingType string  `json:"accountingType" description:"Accounting type"`
	StlmtTime      string  `json:"stlmtTime" description:"Stlmt time"`
	Id             int     `json:"id" description:"Id"` //last netting min id
}

type FP030001O struct {
	PartClearingCode string  `json:"partClearingCode" description:"Part clearing code"`
	Currency         string  `json:"currency" description:"Currency"`
	Balance          float64 `json:"balance" description:"Balance"`
	Status           string  `json:"status"`
}

func (*FP030001I) GetServiceKey() string {
	return "FP030001"
}
