package models

//Participant balance adjust
type FP910029I struct {
	PartClearingCode string             `json:"partClearingCode"`
	PartType         string             `json:"partType"`
	DebCrtFlag       string             `json:"debCrtFlag"`
	TransAmount      float64            `json:"transAmount"`
	Currency         string             `json:"currency"`
	AccountingType   string             `json:"accountingType"`
	Records          []FP910029IAdjItem `json:"records"`
}
type FP910029IAdjItem struct {
	PartClearingCode string  `json:"partClearingCode"`
	DebCrtFlag       string  `json:"debCrtFlag"`
	TransAmount      float64 `json:"transAmount"`
	Currency         string  `json:"currency"`
	AccountingType   string  `json:"accountingType"` // 1-online 2-batch,default online
}
type FP910029O struct {
	Balance float64            `json:"balance"` //Deprecated
	Records []FP910029OBalItem `json:"records"`
}
type FP910029OBalItem struct {
	PartClearingCode string  `json:"partClearingCode"`
	Currency         string  `json:"currency"`
	Balance          float64 `json:"balance"`
}

func (*FP910029I) GetServiceKey() string {
	return "FP910029"
}
