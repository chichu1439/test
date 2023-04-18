package models

type SPD0000056I struct {
	ProductId string `json:"productId" description:"Product id" validate:"required"`
	FeeType   string `json:"feeType" description:"Fee type" validate:"required"`
}

type SPD0000056O struct {
}

func (*SPD0000056I) GetServiceKey() string {
	return "PD000056"
}
