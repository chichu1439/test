package models

// parameters ViewMessageDetailRequest
type FP910015I struct {
	//RecordId     int    `json:"recordId"`
	GlobalSerialNum string `json:"globalSerialNum"`
	TransId         string `json:"transId"`
	MessageId       string `json:"messageId"`
}

// response ViewMessageDetailResponse
type FP910015O struct {
	MessageDate  string `json:"messageDate"`
	Currency     string `json:"currency"`
	ServiceType  string `json:"serviceType"`
	MessageType  string `json:"messageType"`
	SendReceived string `json:"sendReceived"`
	MessageId    string `json:"messageId"`
	Message      string `json:"message"`
}

func (*FP910015I) GetServiceKey() string {
	return "FP910015"
}
