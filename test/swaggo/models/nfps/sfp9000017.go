package models

// Participant balance adjust(db)
type FP900017I struct {
	PartClearingCode string             `json:"partClearingCode"` //Deprecated
	TransDate        string             `json:"transDate"`        //Deprecated
	Currency         string             `json:"currency"`         //Deprecated
	DebCrtFlag       string             `json:"debCrtFlag"`       //Deprecated
	TransAmount      float64            `json:"transAmount"`      //Deprecated
	AccountingType   string             `json:"accountingType"`   //Deprecated 1-online 2-batch,default online
	Balance          float64            `json:"balance"`          //Deprecated
	Records          []FP900017IAdjItem `json:"records"`
}
type FP900017IAdjItem struct {
	AdjustBalance
}
type AdjustBalance struct {
	PartClearingCode string  `json:"partClearingCode" validation:"required"`
	TransDate        string  `json:"transDate" validation:"required"`
	Currency         string  `json:"currency" validation:"required"`
	DebCrtFlag       string  `json:"debCrtFlag" validation:"required"` //D-调减，C-调增
	TransAmount      float64 `json:"transAmount" validation:"required"`
	AccountingType   string  `json:"accountingType"` // 1-online 2-batch,default online
	Balance          float64 `json:"balance"`        //Deprecated
	TraceId          string  `json:"traceId"`        //Deprecated,no need give value
}
type FP900017O struct {
	AdjustBalResult []AdjustBalResult `json:"adjustBalResult"`
}
type AdjustBalResult struct {
	PartClearingCode string  `json:"partClearingCode"`
	Currency         string  `json:"currency"`
	DebCrtFlag       string  `json:"debCrtFlag"`
	TransAmount      float64 `json:"transAmount"`
	Balance          float64 `json:"balance"`
	AlertBalance     float64 `json:"alertBalance"`
	AlertHighWater   float64 `json:"alertHighWater"`
	AlertLowWater    float64 `json:"alertLowWater"`
	AlertMinuteTime  int     `json:"alertMinuteTime"`
}

func (*FP900017I) GetServiceKey() string {
	return "FP900017"
}
