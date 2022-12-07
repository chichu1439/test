package models

type SPD0000055I struct {
	ProductId    string `json:"productId" description:"Product id" validate:"required"`
	InterestType string `json:"interestType" description:"Interest type" validate:"required"`
}

type SPD0000055O struct {
}

func (*SPD0000055I) GetServiceKey() string {
	return "spd0000055"
}
