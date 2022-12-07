package models

type PF000004I struct {
	FeeId string `json:"feeId" validate:"required" description:"Fee Id"`
}

type PF000004O struct {
}

func (*PF000004I) GetServiceKey() string {
	return "PF000004"
}
