package models

//Participant balance debit/credit

type FP000019I struct {
	PartClearingCode string  `json:"partClearingCode" description:"Part clearing code"`
	PartType         string  `json:"partType" description:"Part type"`
	DebCrtFlag       string  `json:"debCrtFlag" description:"Deb crt flag"`
	TransAmount      float64 `json:"transAmount" description:"Trans amount"`
	Currency         string  `json:"currency" description:"Currency"`
}

type FP000019O struct {
	Balance float64 `json:"balance" description:"Balance"`
}

func (*FP000019I) GetServiceKey() string {
	return "FP000019"
}
