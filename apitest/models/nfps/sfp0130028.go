package models

//mocker parameter adjust
type FP013028I struct {
	MockConfig []MockConfig `json:"mockConfig"`
}

type FP013028O struct {
}

func (*FP013028I) GetServiceKey() string {
	return "FP013028"
}

type MockList struct {
	MockConfigs []MockConfig `json:"mockConfigs"`
}
type MockConfig struct {
	Path                    string `json:"path"`
	PartClearingCode        string `json:"partClearingCode"`
	ReasonCode              string `json:"reasonCode"`
	ReasonPrtry             string `json:"reasonPrtry"`
	Delay                   string `json:"delay"` //0:1000:90,1000:3000:5,3000:5000:5
	RequestIso20022MsgType  string `json:"requestIso20022MsgType"`
	ResponseIso20022MsgType string `json:"responseIso20022MsgType"`
}
