package models

//Participant balance debit/credit

type FP000020I struct {
	PartClearingCode string   `json:"partClearingCode" description:"Part clearing code"`
	AlertBalance     *float64 `json:"alertBalance" description:"Alert balance"`
	AlertMinuteTime  *int     `json:"alertMinuteTime" description:"Alert minute time"`
	Currency         string   `json:"currency" description:"Currency"`
}

type FP000020O struct {
	AlertBalance float64 `json:"alertBalance" description:"Alert balance"`
}

func (*FP000020I) GetServiceKey() string {
	return "FP000020"
}
