package models

//Participant balance adjust

type FP000016I struct {
	DirectPartClearingCode string                `json:"directPartClearingCode" description:"Direct part clearing code" validate:"required"`
	Currency               string                `json:"currency" description:"Currency" validate:"required"`
	Records                []FP000016IAdjustItem `json:"records" description:"Records"`
}

type FP000016IAdjustItem struct {
	PartClearingCode string  `json:"partClearingCode" description:"Part clearing code" validate:"required"`
	Currency         string  `json:"currency" description:"Currency" validate:"required"`
	DebCrtFlag       string  `json:"debCrtFlag" description:"Deb crt flag" validate:"required"` //D-调减,C-调增
	TransAmount      float64 `json:"transAmount" description:"Trans amount" validate:"required"`
}

type FP000016O struct {
	Records []FP000016OResItem `json:"records" description:"Records"`
}

type FP000016OResItem struct {
	PartClearingCode string  `json:"partClearingCode" description:"Part clearing code"`
	Currency         string  `json:"currency" description:"Currency"`
	Balance          float64 `json:"balance" description:"Balance"`
}

func (*FP000016I) GetServiceKey() string {
	return "FP000016"
}
