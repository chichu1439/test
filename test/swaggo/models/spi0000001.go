//Version: v0.0.1
package models

type PI000001I struct {
	InterestId    string  `json:"interestId" validate:"required" description:"Interest Id"`
	SystemId      string  `json:"systemId" validate:"required" description:"System Id(LN-icredit;SV-isaving)"`
	ModifySeq     string  `json:"modifySeq" description:"Modify Sequnce"`
	InterestName  string  `json:"interestName" validate:"required" description:"Interest name"`
	Term          string  `json:"term" validate:"required" description:"Term"`
	Currency      string  `json:"currency" validate:"required" description:"Currency(THB)"`
	InterestRate  float64 `json:"interestRate" validate:"required" description:"Interest rate"`
	SourceType    string  `json:"sourceType" validate:"required" description:"Source type"`
	EffectiveDate string  `json:"effectiveDate" description:"Effective date"`
	ExpireDate    string  `json:"expireDate" description:"Expire date"`
}

type PI000001O struct {
	InterestId string `json:"interestId" description:"Interest Id"`
}

func (*PI000001I) GetServiceKey() string {
	return "PI000001"
}
