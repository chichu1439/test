package models

type FP111006I struct {
	ResponsePacs002 []*FPSB2BPacs002
	FileName        string `json:"fileName"`
	FileType        string `json:"fileType"`
	BankCode        string `json:"bankCode"`
	InFileName      string `json:"inFileName"`
}

type FP111006O struct {
	Status string `json:"status"`
}

func (*FP111006I) GetServiceKey() string {
	return "FP111006"
}
