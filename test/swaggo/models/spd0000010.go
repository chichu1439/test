package models

type SPD0000010I struct {
	ProductId string `json:"productId" description:"Product id" validate:"required"`
}

type SPD0000010O struct {
}

func (*SPD0000010I) GetServiceKey() string {
	return "PD000010"
}
