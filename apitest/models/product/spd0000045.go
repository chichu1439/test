package models

type SPD0000045I struct {
	ProductId    string               `json:"productId" validate:"required" description:"Product Id"`
	ListInterest []InterestLoanInfo02 `json:"listInterest" description:"List interest"`
}

type SPD0000045O struct {
}

func (*SPD0000045I) GetServiceKey() string {
	return "spd0000045"
}
