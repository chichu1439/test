package models

type FP1110001I struct {
	FileName string `json:"fileName"`
	FileType string `json:"fileType"`
	BankCode string `json:"bankCode"`
}

type FP111000O struct {
}

func (*FP1110001I) GetServiceKey() string {
	return "FP111001"
}
