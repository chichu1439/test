//Version: v1.0.0
package models

type NT000009I struct {
	StrategyName    string            `json:"strategyName" validate:"required" description:"Strategy Name"`
	StrategyId      int               `json:"strategyId" validate:"required" description:"Strategy Id"`
	NotifyType      string            `json:"notifyType" validate:"required" description:"01-Loan Account Reminder;02-Loan Account Late Notification;03-Account Statement"`
	TargetType      string            `json:"targetType" description:"Target type(0-Customer;1-Internal)"`
	TriggerType     string            `json:"triggerType" description:"Trigger type(0-Transaction;1-Contract Status;2-Natural Period)"`
	PeriodUnit      string            `json:"periodUnit" description:"Period unit(0-Day;1-Month;2-Quarter;3-Year)"`
	PeriodValue     int               `json:"periodValue" description:"Period value" `
	PeriodDay       int               `json:"periodDay" description:"Period day" `
	NotifyTimeType  string            `json:"notifyTimeType" description:"Notify time type(0-Instant;1-Before;2-After)"`
	TimeUnit        string            `json:"timeUnit" description:"Time unit" `
	TimeValue       int               `json:"timeValue" description:"Time value" `
	IntervalUnit    string            `json:"intervalUnit" description:"Interval unit(0-Day;1-Hour;2-Minitue;3-Second)"`
	IntervalValue   int               `json:"intervalValue" description:"Interval value" `
	NotifyTimes     int               `json:"notifyTimes" description:"Notify times" `
	ChannelSelect   string            `json:"channelSelect" description:"Channel select(0-Customize;1-Fixed)"`
	ChannelSend     string            `json:"channelSend" description:"Channel send(0-Must;1-At Least)"`
	AtLeastCount    int               `json:"atLeastCount" description:"At least count" `
	ListChannelInfo []ListChannelInfo `json:"listChannelInfo" description:"List channel information" validate:"required"`
}

type NT000009O struct {
}

func (i *NT000009I) GetServiceKey() string {
	return "NT000009"
}
