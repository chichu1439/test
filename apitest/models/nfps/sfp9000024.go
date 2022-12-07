package models

type FP900024I struct {
	Head04          string `json:"head04"`
	Mti             string `json:"mti"`
	ProcessCode     string `json:"processCode"`
	ChgDirection    string `json:"chgDirection"`
	MessageType     string `json:"messageType"`
	ServiceCode     string `json:"serviceCode"`
	Description     string `json:"description"`
	ServiceTopic    string `json:"serviceTopic"`
	MessageTypeResp string `json:"messageTypeResp"`
}
type FP900024O struct {
	Status string `json:"status"`
}

func (*FP900024I) GetServiceKey() string {
	return "FP900024"
}
