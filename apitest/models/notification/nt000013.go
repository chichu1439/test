package models

type NT000013I struct {
	SystemId         string      `json:"systemId" description:"System id" validate:"required"`
	NotifyType       string      `json:"notifyType" description:"Notify type" validate:"required"`
	ProductId        string      `json:"productId" description:"Product id" validate:"required"`
	CustomerId       string      `json:"customerId" description:"Customer id" validate:"required"`
	AccountId        string      `json:"accountId" description:"Account id" `
	ManagerUser      string      `json:"managerUser" description:"Manager user" `
	ChannelType      string      `json:"channelType" description:"Channel type" `
	ChannelValue     string      `json:"channelValue" description:"Channel value" `
	Body             interface{} `json:"body" validate:"required"`
	AttachFileIDList []string    `json:"attachFileIDList" `
	BizDate          string      `json:"bizDate" validate:"required"`
	BizSeqNo         string      `json:"bizSeqNo" validate:"required"`
	SourceChannel    string      `json:"sourceChannel" `
}

type NT000013O struct {
}

func (i *NT000013I) GetServiceKey() string {
	return "NT000013"
}
