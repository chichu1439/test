package models

type FP100005I struct {
	SendingBankId string `json:"sendingBankId"`
	Data          string `json:"data"`
}

type FP100005O struct {
	SourceData string `json:"sourceData"`
}

func (*FP100005I) GetServiceKey() string {
	return "FP100005"
}
