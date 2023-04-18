package models

//Participant balance adjust（redis）

type FP000033I struct {
	Records []FP000033IAdjItem `json:"records" description:"Records"`
}

type FP000033IAdjItem struct {
	PartClearingCode string  `json:"partClearingCode" description:"Part clearing code"`
	DebCrtFlag       string  `json:"debCrtFlag" description:"Deb crt flag"`
	TransAmount      float64 `json:"transAmount" description:"Trans amount"`
	Currency         string  `json:"currency" description:"Currency"`
}

type FP000033O struct {
	Records []FP000033OBalItem `json:"records" description:"Records"`
}

type FP000033OBalItem struct {
	PartClearingCode string  `json:"partClearingCode" description:"Part clearing code"`
	Currency         string  `json:"currency" description:"Currency"`
	Balance          float64 `json:"balance" description:"Balance"`
}

func (*FP000033I) GetServiceKey() string {
	return "FP000033"
}
