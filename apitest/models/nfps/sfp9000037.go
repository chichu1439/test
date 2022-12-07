package models

type FP900037I struct {
	PartId    string `json:"partId"`
	BatchDate string `json:"batchDate"`
	Currency  string `json:"currency"`
}
type FP900037O struct {
	Record []FP900037OItem
}
type FP900037OItem struct {
	PartClearingCode string  `json:"partClearingCode"`
	SettAcctNo       string  `json:"settAcctNo"`
	Currency         string  `json:"currency"`
	BeginningBalance float64 `json:"beginningBalance"`
	EndingBalance    float64 `json:"endingBalance"`
}

func (*FP900037I) GetServiceKey() string {
	return "FP900037"
}
