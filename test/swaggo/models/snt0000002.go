//Version: v1.0.0
package models

type NT000002I struct {
	SystemId    string `json:"systemId" validate:"required" description:"System Id(LN-loan;SV-saving)"`
	NotifyType  string `json:"notifyType" validate:"required" description:"Notify type(01-Loan Account Reminder;02-Loan Account Late Notification;03-Account Statement)"`
	ChannelType string `json:"channelType" description:"Channel type(01-Email;02-SMS;03-Site;04-Mobile App;05-Call)"`
}

type NT000002O struct {
	ListTemplate []ListTemplate `json:"listTemplate" description:"List template"`
}
