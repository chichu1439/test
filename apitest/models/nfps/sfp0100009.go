package models

// swagger:parameters redeliverMessageRequest

type FP010009I struct {
	PartClearingCode string `json:"partClearingCode" description:"Part clearing code" validate:"required"`
	GlobalSerialNum  string `json:"globalSerialNum" description:"Global serial number" validate:"required"`
	TransId          string `json:"transId" description:"Trans id" validate:"required"`
	MessageId        string `json:"messageId" description:"Message id" validate:"required"`
}

// swagger:response redeliverMessageResponse

type FP010009O struct {
}

func (*FP010009I) GetServiceKey() string {
	return "FP010009"
}
