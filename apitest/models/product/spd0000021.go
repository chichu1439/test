package models

type SPD0000021I struct {
	ProductId string `json:"productId" description:"Product Id"`
}

type SPD0000021O struct {
}

func (*SPD0000021I) GetServiceKey() string {
	return "spd0000021"
}
