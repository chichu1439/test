package models

type FP000034I struct {
	AccountingType string `json:"accountingType" description:"Accounting type"` //1-online,2-batch
}

type FP000034O struct {
}

func (*FP000034I) GetServiceKey() string {
	return "FP000034"
}
