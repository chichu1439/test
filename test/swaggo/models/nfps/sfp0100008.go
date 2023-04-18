package models

// swagger:parameters viewMessageDetailRequest

type FP010008I struct {
	PartClearingCode string `json:"partClearingCode" description:"Part clearing code" validate:"required"`
	GlobalSerialNum  string `json:"globalSerialNum" description:"Global serial number" validate:"required"`
	TransId          string `json:"transId" description:"Trans id" validate:"required"`
	MessageId        string `json:"messageId" description:"Message id" validate:"required"`
}

// swagger:response viewMessageDetailResponse

type FP010008O struct {
	Message string `json:"message" description:"Message"`
}

func (*FP010008I) GetServiceKey() string {
	return "FP010008"
}
