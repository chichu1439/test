//Version: v1.0.0
package models

type NT000013I struct {
	SystemId         string      `json:"systemId" validate:"required" description:"LN-Loan;SV-Saving"`
	NotifyType       string      `json:"notifyType" validate:"required" description:"01-Loan Account Reminder;02-Loan Account Late Notification;03-Account Statement"`
	ProductId        string      `json:"productId" description:"Product id" validate:"required"`
	CustomerId       string      `json:"customerId" description:"Customer id" validate:"required"`
	AccountId        string      `json:"accountId" description:"Account id" `
	ManagerUser      string      `json:"managerUser" description:"Manager user" `
	ChannelType      string      `json:"channelType" description:"01-Email;02-SMS;03-Site;04-Mobile App;05-Call"`
	ChannelValue     string      `json:"channelValue" description:"Channel value" `
	Body             interface{} `json:"body" validate:"required" description:"key-value pair for template content, such as:{“customerName”:”Josh”,“transAmt“:1113}"`
	AttachFileIDList []string    `json:"attachFileIDList" description:"file url list, such as:[“file1.doc“,”file2.doc”]""`
	BizDate          string      `json:"bizDate" validate:"required" description:"The business date (yyyy-mm-dd)"`
	BizSeqNo         string      `json:"bizSeqNo" validate:"required" description:"The business flow"`
	SourceChannel    string      `json:"sourceChannel" description:"Source channel" `
}

type NT000013O struct {
	Status string `json:"status" description:"Status" `
	Data   bool   `json:"data" description:"Data" `
}

func (i *NT000013I) GetServiceKey() string {
	return "NT000013"
}
