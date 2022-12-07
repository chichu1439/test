package models

type FP900042I struct {
	FeeId           string `json:"feeId"`
	ModifySeq       string `json:"modifySeq"`
	FeeIdNew        string `json:"feeIdNew"`
	ModifySeqNew    string `json:"modifySeqNew"`
	FinalModifyOrgz string `json:"finalModifyOrgz"`
	FinalModifyUser string `json:"finalModifyUser"`
	DeleteFlag      bool   `json:"deleteFlag"`
	ExpireDate      string `json:"expireDate"`
}
type FP900042O struct {
	Status string `json:"status"`
}

func (*FP900042I) GetServiceKey() string {
	return "FP900042"
}
