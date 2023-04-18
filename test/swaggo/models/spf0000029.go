package models

type PF000029I struct {
}

type PF000029O struct {
	FeeTypeIdList map[string]string `json:"feeTypeIdList"`
}

func (*PF000029I) GetServiceKey() string {
	return "PF000029"
}
