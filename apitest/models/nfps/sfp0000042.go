package models

type FP000042I struct {
	FeeId      string `json:"feeId"`
	ModifySeq  string `json:"modifySeq"`
	DeleteFlag bool   `json:"deleteFlag"`
	ExpireDate string `json:"expireDate"`
}
type FP000042O struct {
	Status string `json:"status"`
}

func (*FP000042I) GetServiceKey() string {
	return "FP000042"
}
