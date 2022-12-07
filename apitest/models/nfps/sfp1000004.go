package models

type FP100004Outgress struct {
	Body []byte `json:"body"`
}

func (*FP100004Outgress) GetServiceKey() string {
	return "FP200004"
}

type FP100004OutgressResponse struct {
	Body []byte `json:"body"`
}
