package models


type FP020007I struct {
	FileName string `json:"fileName"`
	PartClearingCode string `json:"partClearingCode"`
	Body string `json:"body"`
}

func (*FP020007I) GetServiceKey() string {
	return "FP020007"
}

// todo 暂时用一下
type ResponseStatus struct {
	Status string
}

