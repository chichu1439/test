package models

// Participant balance adjust(db)

type FP000032I struct {
	Records []FP000032IAdjItem `json:"records" description:"Records"`
}

type FP000032IAdjItem struct {
	PartClearingCode string  `json:"partClearingCode" description:"Part clearing code"`
	TransDate        string  `json:"transDate" description:"Trans date"`
	Currency         string  `json:"currency" description:"Currency"`
	DebCrtFlag       string  `json:"debCrtFlag" description:"Deb crt flag"`
	TransAmount      float64 `json:"transAmount" description:"Trans amount"`
	TraceId          string  `json:"traceId" description:"Trace id"`
}

type FP000032O struct {
	AdjustBalResult []AdjustBalResult `json:"adjustBalResult" description:"Adjust bal result"`
}

type AdjustBalResult2 struct {
	PartClearingCode string  `json:"partClearingCode" description:"Part clearing code"`
	Currency         string  `json:"currency" description:"Currency"`
	DebCrtFlag       string  `json:"debCrtFlag" description:"Deb crt flag"`
	TransAmount      float64 `json:"transAmount" description:"Trans amount"`
	Balance          float64 `json:"balance" description:"Balance"`
}

func (*FP000032I) GetServiceKey() string {
	return "FP000032"
}
