package models

// parameters ViewMessageListRequest
type FP910014I struct {
	PartClearingCode string `json:"partClearingCode"`
	MessageDate      string `json:"messageDate"`
	Currency         string `json:"currency"`
	ServiceType      string `json:"serviceType"`
	MessageType      string `json:"messageType"`
	SendReceived     string `json:"sendReceived"`
	MessageId        string `json:"messageId"`
	PageNo           int    `json:"pageNo"`       //查询页数
	PageRecCount     int    `json:"pageRecCount"` //每页记录数
}

// response ViewMessageListResponse
type FP910014O struct {
	TotalRecCount int             `json:"totalRecCount"`
	TotalPage     int             `json:"totalPage"`
	Records       []FP910014OItem `json:"records"`
}

// swagger:response ViewMessageListResponse
type FP910014OItem struct {
	RecordId     int    `json:"recordId"`
	MessageDate  string `json:"messageDate"`
	Currency     string `json:"currency"`
	ServiceType  string `json:"serviceType"`
	MessageType  string `json:"messageType"`
	MessageId    string `json:"messageId"`
	SendReceived string `json:"sendReceived"`
	Message      string `json:"message"`
}

func (*FP910014I) GetServiceKey() string {
	return "FP910014"
}
