//Version: v1.0.0
package models

type NT000012I struct {
	TargetType     string `json:"targetType" description:"Target type(0-Customer;1-Internal)"`
	ChannelType    string `json:"channelType" description:"Channel type(01-Email;02-SMS;03-Site;04-Mobile App;05-Call)"`
	TargetKeyType  string `json:"targetKeyType" description:"TargetKey type(0-Customer;1-Account;2-Certificate ID;3-Organization;4-Manager)"`
	TargetKeyValue string `json:"targetKeyValue" description:"Target key value" `
	ChannelValue   string `json:"channelValue" description:"Channel value" `
	SystemId       string `json:"systemId" description:"System Id(LN-Loan;SV-Saving)"`
	SourceType     string `json:"sourceType" description:"Source type" `
}

type NT000012O struct {
}
