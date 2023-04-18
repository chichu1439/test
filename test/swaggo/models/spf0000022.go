package models

type PF000022I struct {
	FeeTypeId string `json:"feeTypeId"`
}

type PF000022O struct {
}

func (*PF000022I) GetServiceKey() string {
	return "PF000022"
}
