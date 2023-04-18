package models

type FP111004I struct {
	FileName string `json:"fileName"`
	BankCode string `json:"bankCode"`
	FileType string `json:"fileType"`
}

type OutDoneGateway struct {
	FileName string `json:"fileName"`
	//BankCode string `json:"bankCode"`
	//FileType string `json:"fileType"`
}

type InDoneGateway struct {
	Body []byte `json:"body"`
}

type InDoneGatewayJson struct {
	Status string `json:"status"`
}

type FP111004O struct {
	Status string `json:"status"`
}

func (*FP111004I) GetServiceKey() string {
	return "FP111004"
}

func (*OutDoneGateway) GetServiceKey() string {
	return "FP210030"
}
