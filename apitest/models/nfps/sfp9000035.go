package models

type FP900035I struct {
	AccountingType string `json:"accountingType"` //1-online,2-batch
}
type FP900035O struct {
}

func (*FP900035I) GetServiceKey() string {
	return "FP900035"
}
