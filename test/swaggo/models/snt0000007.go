//Version: v1.0.0
package models

type NT000007I struct {
	SystemId   string `json:"systemId" validate:"required" description:"LN-loan;SV-saving"`
	NotifyType string `json:"notifyType" description:"Notify type(01-Loan Account Reminder;02-Loan Account Late Notification;03-Account Statement)"`
	TargetType string `json:"targetType" description:"Target type(0-Customer;1-Internal)"`
}

type NT000007O struct {
	// ListSrategy []ListSrategy `json:"listSrategy" description:"List srategy"`
}
