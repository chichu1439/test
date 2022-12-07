package models

type FPX00001I struct {
	Flag          string `json:"flag"` //D-decrypt,E-encrypt
	SendingBankId string `json:"sendingBankId"`
	Data          string `json:"data"`
}

type FPX00001O struct {
	Key        string `json:"key"`
	SourceData string `json:"sourceData"`
	HexData    string `json:"hexData"`
}

func (*FPX00001I) GetServiceKey() string {
	return "FPX00001"
}
