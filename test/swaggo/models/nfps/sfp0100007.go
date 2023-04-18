package models

// swagger:parameters ViewMessageListRequest

type FP010007I struct {
	PartClearingCode string `json:"partClearingCode" description:"Part clearing code" describe:"参与者编号" validate:"required,max=35"`
	MessageDate      string `json:"messageDate" description:"Message date" describe:"报文创建日期"`
	Currency         string `json:"currency" description:"Currency" describe:"货币"`
	ServiceType      string `json:"serviceType" description:"Service type" describe:"服务类型"`
	MessageType      string `json:"messageType" description:"Message type" describe:"报文类型"`
	SendReceived     string `json:"sendReceived" description:"Send received" describe:"报文方向"`
	MessageId        string `json:"messageId" description:"Message id" describe:"报文编号"`
	PageNo           int    `json:"pageNo" description:"Page no"`              //查询页数
	PageRecCount     int    `json:"pageRecCount" description:"Page rec count"` //每页记录数
}

// swagger:response ViewMessageListResponse

type FP010007O struct {
	TotalRecCount int             `json:"totalRecCount" description:"Total rec count"`
	TotalPage     int             `json:"totalPage" description:"Total page"`
	Records       []FP010007OItem `json:"records" description:"Records"`
}

// swagger:response FP010007OItem

type FP010007OItem struct {
	//RecordId        int `json:"recordId"`
	GlobalSerialNum string `json:"globalSerialNum" description:"Global serial number"`
	TransId         string `json:"transId" description:"Trans id"`
	MessageDate     string `json:"messageDate" description:"Message date"`
	Currency        string `json:"currency" description:"Currency"`
	ServiceType     string `json:"serviceType" description:"Service type"`
	MessageType     string `json:"messageType" description:"Message type"`
	SendReceived    string `json:"sendReceived" description:"Send received"`
	MessageId       string `json:"messageId" description:"Message id"`
	Message         string `json:"message" description:"Message"`
}

func (*FP010007I) GetServiceKey() string {
	return "FP010007"
}
