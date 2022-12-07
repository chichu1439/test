package models

//Settlement Account Opening
type FP900012I struct {
	PartClearingCode string             `json:"partClearingCode"`
	UpperPart        string             `json:"upperPart"`
	PartType         string             `json:"partType"`
	Records          []FP900012ICurItem `json:"records"`
}
type FP900012ICurItem struct {
	Currency        string  `json:"currency"`
	SettAcctNo      string  `json:"settAcctNo"`
	SettAcctBalance float64 `json:"settAcctBalance"` //deprecated
	AcctBalance     float64 `json:"acctBalance"`     //deprecated
	Status          string  `json:"status"`
}
type FP900012O struct {
}

func (*FP900012I) GetServiceKey() string {
	return "FP900012"
}
