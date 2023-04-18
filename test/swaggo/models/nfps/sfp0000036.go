package models

type FP000036I struct {
	Records []FP000036IAdjItem `json:"records"`
}
type FP000036O struct {
	AdjustBalResult []FP000036OTemp `json:"adjustBalResult"`
}
type FP000036IAdjItem struct {
	PartClearingCode string  `json:"partClearingCode"`
	TransDate        string  `json:"transDate"`
	Currency         string  `json:"currency"`
	DebCrtFlag       string  `json:"debCrtFlag"`
	TransAmount      float64 `json:"transAmount"`
}
type FP000036OTemp struct {
	PartClearingCode string  `json:"partClearingCode"`
	Currency         string  `json:"currency"`
	DebCrtFlag       string  `json:"debCrtFlag"`
	TransAmount      float64 `json:"transAmount"`
	Balance          float64 `json:"balance"`
}

func (*FP000036I) GetServiceKey() string {
	return "FP000036"
}

type FP000036Outgress struct {
	Body []byte `json:"body"`
}

func (*FP000036Outgress) GetServiceKey() string {
	return "FP200036"
}

type FP000036OutgressResponse struct {
	Body []byte `json:"body"`
}
