package models

type PD000078I struct {
	ProductId  string `json:"productId" validate:"required" description:"Product Id"`
	NotifyType string `json:"notifyType" validate:"required" description:"Notify type(01-Loan Account Reminder;02-Loan Account Late Notification;03-Account Statement)"`
}

type PD000078O struct {
	ListStrategy []NotificationStrategy `json:"listStrategy" description:"List strategy"`
}

type NotificationStrategy struct {
	StrategyId int `json:"strategyId" description:"Strategy id"`
}

func (*PD000078I) GetServiceKey() string {
	return "PD000078"
}
