package models

type SPD0000044I struct {
	ProductId string `json:"productId" validate:"required" description:"Product Id"`
}

type SPD0000044O struct {
	ListInterest []InterestLoanInfo02 `json:"listInterest" description:"List interest"`
}

type InterestLoanInfo02 struct {
	InterestType string               `json:"interestType" description:"interest Type:0-normal 1-Withdrawal in Advance"` //0-normal 1-Withdrawal in Advance
	ListStrategy []InterestStrategy02 `json:"listStrategy" description:"List strategy"`
}

func (*SPD0000044I) GetServiceKey() string {
	return "spd0000044"
}
