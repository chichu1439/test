package models

type SPD0000014I struct {
	ProductId string `json:"productId" validate:"required" description:"Product Id"`
}

type SPD0000014O struct {
}

func (*SPD0000014I) GetServiceKey() string {
	return "PD000014"
}
