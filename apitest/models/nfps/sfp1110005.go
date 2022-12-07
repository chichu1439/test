package models

type FP111005I struct {
	FileName   string `json:"fileName"`
	FileType   string `json:"fileType"`
	BankCode   string `json:"bankCode"`
	InFileName string `json:"inFileName"`
}

type FP111005O struct {
	Status string `json:"status"`
}

func (*FP111005I) GetServiceKey() string {
	return "FP111005"
}
