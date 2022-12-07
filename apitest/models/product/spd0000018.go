package models

type SPD0000018I struct {
	ProductId string              `json:"productId" validate:"required" description:"Product Id"`
	ListFee   []FeeSavingScenario `json:"listFee" description:"List fee" validate:"required,dive"`
}

type SPD0000018O struct {
}

func (*SPD0000018I) GetServiceKey() string {
	return "spd0000018"
}
