//Version: v1
package models

type PD000123I struct {
	ProductId       string `json:"productId"`
	PageNum         int    `json:"pageNum"`
	PageRecordCount int    `json:"pageRecordCount"`
}

type PD000123O struct {
	HistoryRecords  HistoryRecords  `json:"historyRecords"`
	PlanningRecords PlanningRecords `json:"planningRecords"`
}

type Records struct {
	Uid            int              `json:"uid"`
	ProductId      string           `json:"productId"`
	NotifyType     string           `json:"notifyType"`
	StrategyId     int              `json:"strategyId"`
	EffectiveDate  string           `json:"effectiveDate"`
	ExpireDate     string           `json:"expireDate"`
	Status         string           `json:"status"`
	TargetType     string           `json:"targetType"`
	StrategyName   string           `json:"strategyName"`
	TriggerType    string           `json:"triggerType"`
	PeriodUnit     string           `json:"periodUnit"`
	PeriodValue    int              `json:"periodValue"`
	PeriodDay      int              `json:"periodDay"`
	NotifyTimeType string           `json:"notifyTimeType"`
	TimeUnit       string           `json:"timeUnit"`
	TimeValue      int              `json:"timeValue"`
	IntervalUnit   string           `json:"intervalUnit"`
	IntervalValue  int              `json:"intervalValue"`
	NotifyTimes    int              `json:"notifyTimes"`
	ChannelSelect  string           `json:"channelSelect"`
	ChannelSend    string           `json:"channelSend"`
	AtLeastCount   int              `json:"atLeastCount"`
	RecordTime     string           `json:"recordTime"`
	RecordDate     string           `json:"recordDate"`
	ListChannelObj []ListChannelObj `json:"listChannelObj" description:"List channel obj"`
}

type HistoryRecords struct {
	PageNum         int       `json:"pageNum"`
	PageRecordCount int       `json:"pageRecordCount"`
	TotalCount      int       `json:"totalCount"`
	Records         []Records `json:"records"`
}

type PlanningRecords struct {
	TotalCount int       `json:"totalCount"`
	Records    []Records `json:"records"`
}

func (*PD000123I) GetServiceKey() string {
	return "PD000123"
}
