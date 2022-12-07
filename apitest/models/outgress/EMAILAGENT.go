package outgress

type EmailAgentI struct {
	MessageID string `json:"messageID" description:"Message i d" validate:"required"`
}

type EmailAgentO struct {
	ErrorCode string          `json:"errorCode" description:"Error code" validate:"required"`
	ErrorMsg  string          `json:"errorMsg" description:"Error msg" validate:"required"`
	Data      EmailAgentOData `json:"data" description:"Data" validate:"required"`
}

type EmailAgentOData struct {
	DeliveryTime string `json:"deliveryTime" description:"Delivery time" validate:"required"`
	ResponseTime string `json:"responseTime" description:"Response time" validate:"required"`
	Data         bool   `json:"data" description:"Data" validate:"required"`
}

func (i *EmailAgentI) GetServiceKey() string {
	return "EMAILAGENT"
}
