//Version: v0.0.1
package models

type PI000002I struct {
	InterestName string `json:"interestName" description:"Uid"`
	SystemId     string `json:"systemId" description:"System Id(LN-icredit;SV-isaving)"`
	InterestId   string `json:"interestId" description:"Interest Id"`
	Term         string `json:"term" description:"Term"`
	Currency     string `json:"currency" description:"Currency(THB)"`
}

type PI000002O struct {
	ListRate []BaseInterestRate `json:"listRate" description:"List rate"`
}

type BaseInterestRate struct {
	Uid           int64   `json:"uid" description:"Uid"`
	InterestId    string  `json:"interestId" description:"Interest Id"`
	SystemId      string  `json:"systemId" description:"System Id(LN-icredit;SV-isaving)"`
	InterestName  string  `json:"interestName" description:"Interest Name"`
	Term          string  `json:"term" description:"Term"`
	Currency      string  `json:"currency" description:"Currency(THB)"`
	InterestRate  float64 `json:"interestRate" description:"Interest rate"`
	SourceType    string  `json:"sourceType" description:"Source type"`
	EffectiveDate string  `json:"effectiveDate" description:"Effective date"`
	ExpireDate    string  `json:"expireDate" description:"Expire date"`
}

func (*PI000002I) GetServiceKey() string {
	return "PI000002"
}
