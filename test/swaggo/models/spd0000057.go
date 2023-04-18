package models

type SPD0000057I struct {
	ProductId    string `json:"productId" description:"Product id" validate:"required"`
	InterestType string `json:"interestType" validate:"required" description:"interest Type:0-normal 1-Withdrawal in Advance"` //0-normal 1-Withdrawal in Advance
}

type SPD0000057O struct {
}

func (*SPD0000057I) GetServiceKey() string {
	return "PD000057"
}
