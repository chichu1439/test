package models

type PI000011I struct {
	InterestName string `json:"interestName" description:"Uid"`
	SystemId     string `json:"systemId" description:"System Id(LN-icredit;SV-isaving)"`
	InterestId   string `json:"interestId" description:"Interest Id"`
	Term         string `json:"term" description:"Term"`
	Currency     string `json:"currency" description:"Currency(THB)"`
}

type PI000011O struct {
	ListRate []BaseInterestRate `json:"listRate" description:"List rate"`
}

func (*PI000011I) GetServiceKey() string {
	return "PI000011"
}
