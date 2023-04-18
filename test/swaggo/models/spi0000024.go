//Version: v0.0.1
package models

type PI000024I struct {
	InterestName string `json:"interestName"`
	SystemId     string `json:"systemId"`
	Term         string `json:"term"`
	Currency     string `json:"currency"`
	InterestId   string `json:"interestId"`
	StartExpDate string `json:"startExpDate"`
	EndExpDate   string `json:"endExpDate"`
}

type PI000024O struct {
	ListRate []BaseInterestRate `json:"listRate"`
}

func (*PI000024I) GetServiceKey() string {
	return "PI000024"
}
