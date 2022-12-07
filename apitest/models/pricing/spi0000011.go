//Version: v0.0.1
package models

type PI000011I struct {
	InterestName  string  `json:"interestName" validate:"required" description:"Interest name"`
	Term          string  `json:"term" validate:"required" description:"Term"`
	Currency      string  `json:"currency" validate:"required" description:"Currency(THB)"`
	InterestRate  float64 `json:"interestRate" validate:"required" description:"Interest rate"`
	SourceType    string  `json:"sourceType" validate:"required" description:"Source type"`
	EffectiveDate string  `json:"effectiveDate" description:"Effective date"`
	ExpireDate    string  `json:"expireDate" description:"Expire date"`
}

type PI000011O struct {
	InterestId string `json:"interestId" description:"Interest Id"`
}

func (*PI000011I) GetServiceKey() string {
	return "PI000011"
}
