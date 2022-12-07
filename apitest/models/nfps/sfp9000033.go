package models

type FP900033I struct {
	Head04      string `json:"head04"`
	Mti         string `json:"mti"`
	ProcessCode string `json:"processCode"`
}
type FP900033O struct {
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

func (*FP900033I) GetServiceKey() string {
	return "FP900033"
}
