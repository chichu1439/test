package models

//Participant balance adjust
type FP900014I struct {
	PartClearingCode string             `json:"partClearingCode"` //Deprecated
	PartType         string             `json:"partType"`         //Deprecated
	DebCrtFlag       string             `json:"debCrtFlag"`       //Deprecated
	TransAmount      float64            `json:"transAmount"`      //Deprecated
	Currency         string             `json:"currency"`         //Deprecated
	AccountingType   string             `json:"accountingType"`   //Deprecated 1-online 2-batch,default online
	Records          []FP900014IAdjItem `json:"records"`
	//
	//OutPartClearingCode string
	//OutPartType         string
	//OutUpperPart        string
	//InPartClearingCode  string
	//InPartType          string
	//InUpperPart         string
}
type FP900014IAdjItem struct {
	PartClearingCode string  `json:"partClearingCode"`
	DebCrtFlag       string  `json:"debCrtFlag"`
	TransAmount      float64 `json:"transAmount"`
	Currency         string  `json:"currency"`
	AccountingType   string  `json:"accountingType"` // 1-online 2-batch,default online
}
type FP900014O struct {
	Balance float64            `json:"balance"` //Deprecated
	Records []FP900014OBalItem `json:"records"`
}
type FP900014OBalItem struct {
	PartClearingCode string  `json:"partClearingCode"`
	Currency         string  `json:"currency"`
	Balance          float64 `json:"balance"`
}

func (*FP900014I) GetServiceKey() string {
	return "FP900014"
}
