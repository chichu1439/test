package models

import "github.com/shopspring/decimal"

type IC000015I struct {
	ProductId string `json:"productId" description:"Product id"`
}

type IC000015O struct {
	Records []RiskClassification `json:"records" description:"Records"`
}

type RiskClassification struct {
	ProductId        string          `json:"productId" description:"Product id"`
	RiskClsfCdPerMon string          `json:"riskClsfCdPerMon" description:"Risk level:1-Normal 2-Attention 3-Secondary 4-Suspicious 5-Loss"` //1-Normal 2-Attention 3-Secondary 4-Suspicious 5-Loss
	TotalMon         decimal.Decimal `json:"totalMon" description:"Total month"`
	CurrtMonth       string          `json:"currtMonth" description:"Currt month"` //2021-05
}

func (*IC000015I) GetServiceKey() string {
	return "IC000015"
}
