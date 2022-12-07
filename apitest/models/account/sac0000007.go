package models

type AC000007I struct {
	Account       string `json:"account" description:"Account" validate:"required"`
	AccountStatus string `json:"accountStatus" description:"Account status" validate:"oneof=9"`
}

type AC000007O struct {
	Status string `json:"status" description:"Status"`
}

func (*AC000007I) GetServiceKey() string {
	return "sac0000007"
}
