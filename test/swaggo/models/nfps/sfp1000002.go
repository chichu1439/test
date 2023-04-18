package models

type FP100002I struct {
	PartClearingCode string `json:"partClearingCode"`
	PartType         string `json:"partType"`
	Currency         string `json:"currency"`
}
type FP100002O struct {
	PartClearingCode string          `json:"partClearingCode"`
	PartType         string          `json:"partType"`
	PartBic          string          `json:"partBic"`
	Records          []FP100002OItem `json:"records"`
}
type FP100002OItem struct {
	PartCurrency    string  `json:"partCurrency"`
	PartSettleAcct  string  `json:"partSettleAcct"`
	SettAcctBalance float64 `json:"settAcctBalance"`
}

func (*FP100002I) GetServiceKey() string {
	return "FP100002"
}
