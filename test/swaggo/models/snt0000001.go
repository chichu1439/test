//Version: v1.0.0
package models

type NT000001I struct {
	SystemId        string `json:"systemId" validate:"required" description:"System Id(LN-loan;SV-saving)"`
	NotifyType      string `json:"notifyType" validate:"required" description:"Notify type(01-Loan Account Reminder;02-Loan Account Late Notification;03-Account Statement)"`
	ChannelType     string `json:"channelType" validate:"required" description:"Channel type(01-Email;02-SMS;03-Site;04-Mobile App;05-Call)"`
	TargetType      string `json:"targetType" validate:"required" description:"Target type(0-Customer;1-Internal)"`
	TemplateName    string `json:"templateName" description:"Template name" validate:"required"`
	TemplateTitle   string `json:"templateTitle" description:"Template title" validate:"required"`
	TemplateContent string `json:"templateContent" description:"Template content" validate:"required"`
}

type NT000001O struct {
	TemplateId int    `json:"templateId" description:"Template id"`
	RelateMcId string `json:"relateMcId" description:"Template id created by MC"`
}
