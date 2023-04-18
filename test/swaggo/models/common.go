package models

type ListSrategy struct {
	StrategyId     int    `json:"strategyId" description:"Strategy id"`
	StrategyName   string `json:"strategyName" description:"Strategy name"`
	TargetType     string `json:"targetType" description:"Target type"`
	NotifyTimeType string `json:"notifyTimeType" description:"Notify time type"`
	TimeUnit       string `json:"timeUnit" description:"Time unit"`
	TimeValue      int    `json:"timeValue" description:"Time value"`
	IntervalFlag   string `json:"intervalFlag" description:"Interval flag"`
	IntervalUnit   string `json:"intervalUnit" description:"Interval unit"`
	IntervalValue  int    `json:"intervalValue" description:"Interval value"`
	NotifyTimes    int    `json:"notifyTimes" description:"Notify times"`
}

type ListChannelResponse struct {
	MessageId      string `json:"messageId" description:"Message id"`
	ChannelType    string `json:"channelType" description:"Channel type"`
	ChannelValue   string `json:"channelValue" description:"Channel value"`
	Order          int    `json:"order" description:"Order"`
	RetryTimes     int    `json:"retryTimes" description:"Retry times"`
	RetryDoneTimes int    `json:"retryDoneTimes" description:"Retry done times"`
	ErrorCode      string `json:"errorCode" description:"Error code"`
	ErrorMsg       string `json:"errorMsg" description:"Error msg"`
	Data           Data   `json:"data" description:"Data"`
}

type Data struct {
	DeliveryTime string `json:"deliveryTime" description:"Delivery time"`
	ResponseTime string `json:"responseTime" description:"Response time"`
	Data         bool   `json:"data" description:"Data"`
}

type ListTemplate struct {
	TemplateId    int    `json:"templateId" description:"Template id"`
	ChannelType   string `json:"channelType" description:"Channel type"`
	TargetType    string `json:"targetType" description:"Target type"`
	TemplateName  string `json:"templateName" description:"Template name"`
	TemplateTitle string `json:"templateTitle" description:"Template title"`
	RelateMcId    string `json:"relateMcId" description:"Relate mc id"`
}

type ListChannelInfo struct {
	ChannelType string `json:"channelType" description:"Channel type"`
	TemplateId  int    `json:"templateId" description:"Template id"`
	MustSend    string `json:"mustSend" description:"Must send"`
	Order       int    `json:"order" description:"Order"`
	RetryTimes  int    `json:"retryTimes" description:"Retry times"`
}

type ListChannelObj2 struct {
	ChannelType     string `json:"channelType" description:"Channel type"`
	TemplateId      int    `json:"templateId" description:"Template id"`
	RelateMcId      string `json:"relateMcId" description:"Relate mc id"`
	TemplateName    string `json:"templateName" description:"Template name"`
	TemplateContent string `json:"templateContent" description:"Template content"`
	MustSend        string `json:"mustSend" description:"Must send"`
	Order           int    `json:"order" description:"Order"`
	RetryTimes      int    `json:"retryTimes" description:"Retry times"`
}

type ListChannelTarget struct {
	MessageId      string `json:"messageId" description:"Message id"`
	ChannelType    string `json:"channelType" description:"Channel type"`
	ChannelValue   string `json:"channelValue" description:"Channel value"`
	Order          int    `json:"order" description:"Order"`
	RetryTimes     int    `json:"retryTimes" description:"Retry times"`
	TemplateId     int    `json:"templateId" description:"Template id"`
	RelateMcId     string `json:"relateMcId" description:"Relate mc id"`
	TemplateTitle  string `json:"templateTitle" description:"Template title"`
	TargetKeyType  string `json:"targetKeyType" description:"Target key type"`
	TargetKeyValue string `json:"targetKeyValue" description:"Target key value"`
}

type TReportRegisterBo struct {
	DateOfReport         string `json:"dateOfReport" description:"Date of report"`                // ",
	ReportName           string `json:"reportName" description:"Report name"`                     // ",
	SequenceNum          int    `json:"sequenceNum" description:"Sequence number"`                // ",
	ReportType           string `json:"reportType" description:"Report type"`                     // IL001-Regulatory IL002-National Credit Bureau IL003-Disbursement IL004-Repayment IL005-Other ",
	ReportClassification string `json:"reportClassification" description:"Report classification"` // ",
	ReportGenerateTime   string `json:"reportGenerateTime" description:"Report generate time"`    // ",
	ReportGenerateDate   string `json:"reportGenerateDate" description:"Report generate date"`    // ",
	ReportFileType       string `json:"reportFileType" description:"Report file type"`            // csv excel ",
	ReportFileId         string `json:"reportFileId" description:"Report file id"`                // ",
	ReportFileKey        string `json:"reportFileKey" description:"Report file key"`              // ",
}

type ListStrategyDetail struct {
	StrategyId        int                 `json:"strategyId" description:"Strategy id" validate:"required"`
	TargetType        string              `json:"targetType" description:"Target type" validate:"required"`
	TriggerType       string              `json:"triggerType" description:"Trigger type" validate:"required"`
	PeriodUnit        string              `json:"periodUnit" description:"Period unit" `
	PeriodValue       int                 `json:"periodValue" description:"Period value" `
	PeriodDay         int                 `json:"periodDay" description:"Period day" `
	NotifyTimeType    string              `json:"notifyTimeType" description:"Notify time type" validate:"required"`
	TimeUnit          string              `json:"timeUnit" description:"Time unit" `
	TimeValue         int                 `json:"timeValue" description:"Time value" `
	IntervalUnit      string              `json:"intervalUnit" description:"Interval unit" `
	IntervalValue     int                 `json:"intervalValue" description:"Interval value" `
	NotifyTimes       int                 `json:"notifyTimes" description:"Notify times" `
	ChannelSend       string              `json:"channelSend" description:"Channel send" validate:"required"`
	AtLeastCount      int                 `json:"atLeastCount" description:"At least count" `
	ListChannelTarget []ListChannelTarget `json:"listChannelTarget" description:"List channel target" validate:"required"`
}
