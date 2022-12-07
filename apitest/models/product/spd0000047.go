package models

type SPD0000047I struct {
	ProductId string        `json:"productId" description:"Product Id"`
	ListFee   []FeeLoanInfo `json:"listFee" description:"List fee"`
}

type SPD0000047O struct {
}

func (*SPD0000047I) GetServiceKey() string {
	return "spd0000047"
}
