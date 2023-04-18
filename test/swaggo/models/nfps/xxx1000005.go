package models

type XX100005I struct {
	Flag          string `json:"flag"` //D-decrypt,E-encrypt
	SendingBankId string `json:"sendingBankId"`
	Data          string `json:"data"`
}

type XX100005O struct {
	Key        string `json:"key"`
	SourceData string `json:"sourceData"`
	HexData    string `json:"hexData"`
}

func (*XX100005I) GetServiceKey() string {
	return "XX100005"
}
