package models

type SPD0000016I struct {
	ProductId    string         `json:"productId" validate:"required" description:"Product Id"`
	ListInterest []InterestInfo `json:"listInterest" description:"List interest" validate:"required,dive"`
}

type SPD0000016O struct {
}

func (*SPD0000016I) GetServiceKey() string {
	return "spd0000016"
}
