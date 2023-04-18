package models

type SPD0000015I struct {
	ProductId string `json:"productId" description:"Product Id"`
}

type SPD0000015O struct {
	ListInterest []ProductInterestSaving02 `json:"listInterest" description:"List interest"`
}

type ProductInterestSaving02 struct {
	InterestType string               `json:"interestType" description:"interest Type:0-normal loan;1-overdue loan"`
	ListStrategy []InterestStrategy02 `json:"listStrategy" description:"List strategy"`
}

type InterestStrategy02 struct {
	InterestStrategy string `json:"interestStrategy" description:"Interest strategy"`
	Status           string `json:"status" description:"status:0-normal"`
}

func (*SPD0000015I) GetServiceKey() string {
	return "PD000015"
}
