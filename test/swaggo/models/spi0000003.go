//Version: v0.0.1
package models

type PI000003I struct {
	Uid           int64   `json:"uid" validate:"required" description:"Uid"`
	ModifySeq     string  `json:"modifySeq" description:"Modify sequnce"`
	InterestId    string  `json:"interestId" description:"Interest Id"`
	InterestName  string  `json:"interestName" description:"Interest Name"`
	Term          string  `json:"term" description:"Term"`
	Currency      string  `json:"currency" description:"Currency"`
	InterestRate  float64 `json:"interestRate" description:"Interest rate"`
	SourceType    string  `json:"sourceType" description:"source type"`
	EffectiveDate string  `json:"effectiveDate" description:"Effective date"`
	ExpireDate    string  `json:"expireDate" description:"Expire date"`
}

type PI000003O struct {
}

func (*PI000003I) GetServiceKey() string {
	return "PI000003"
}
