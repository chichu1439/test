//Version: v1.0.0
package models

type NT000006I struct {
	StrategyName    string            `json:"strategyName" description:"Strategy name" validate:"required"`
	SystemId        string            `json:"systemId" validate:"required" description:"LN-loan;SV-saving"`
	NotifyType      string            `json:"notifyType" validate:"required" description:"Notify type(01-Loan Account Reminder;02-Loan Account Late Notification;03-Account Statement)"`
	TargetType      string            `json:"targetType" validate:"required" description:"Target type(0-Customer;1-Internal)"`
	TriggerType     string            `json:"triggerType" validate:"required" description:"Trigger type(0-Transaction;1-Contract Status;2-Natural Period)"`
	PeriodUnit      string            `json:"periodUnit" description:"Period unit(0-Day;1-Month;2-Quarter;3-Year)"`
	PeriodValue     int               `json:"periodValue" description:"Period value" `
	PeriodDay       int               `json:"periodDay" description:"Period day" `
	NotifyTimeType  string            `json:"notifyTimeType" validate:"required" description:"Notify time type(0-Instant;1-Before;2-After)"`
	TimeUnit        string            `json:"timeUnit" description:"Time unit" `
	TimeValue       int               `json:"timeValue" description:"Time value" `
	IntervalUnit    string            `json:"intervalUnit" description:"Interval unit" `
	IntervalValue   int               `json:"intervalValue" description:"Interval value(0-Day;1-Hour;2-Minitue;3-Second)"`
	NotifyTimes     int               `json:"notifyTimes" description:"Notify times" `
	ChannelSelect   string            `json:"channelSelect" validate:"required" description:"Channel select(0-Customize;1-Fixed)"`
	ChannelSend     string            `json:"channelSend" validate:"required" description:"Channel send(0-Must;1-At Least)"`
	AtLeastCount    int               `json:"atLeastCount" description:"At least count" `
	ListChannelInfo []ListChannelInfo `json:"listChannelInfo" description:"List channel information" validate:"required"`
}

type NT000006O struct {
	StrategyId int64 `json:"strategyId" description:"Strategy id"`
}
