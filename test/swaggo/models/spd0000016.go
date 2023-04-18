package models

type SPD0000016I struct {
	ProductId    string         `json:"productId" validate:"required" description:"Product Id"`
	ListInterest []InterestInfo `json:"listInterest" description:"List interest" validate:"required,dive"`
}
type InterestInfo struct {
	InterestType string             `json:"interestType" description:"Interest type" validate:"required"` //0-normal 1-Withdrawal in Advance
	ListStrategy []InterestStrategy `json:"listStrategy" description:"List strategy" validate:"dive"`
}

type InterestStrategy struct {
	InterestStrategy string `json:"interestStrategy" description:"Interest strategy" validate:"required"`
}
type SPD0000016O struct {
}

func (*SPD0000016I) GetServiceKey() string {
	return "PD000016"
}
