package models

type FP900025I struct {
	Id            int    `json:"id"`
	BatchSequence string `json:"batchSequence"`
	BatchInitDate string `json:"batchInitDate"`
	BatchInitTime string `json:"batchInitTime"`
	BatchType     string `json:"batchType"`
	ExecuteId     string `json:"executeId"`
	ProjectId     string `json:"projectId"`
	FlowId        string `json:"flowId"`
	JobId         string `json:"jobId"`
	SysSource     string `json:"sysSource"`
	ServiceCode   string `json:"serviceCode"`
	BatchEndDate  string `json:"batchEndDate"`
	BatchEndTime  string `json:"batchEndTime"`
}
type FP900025O struct {
	Status string `json:"status"`
}

func (*FP900025I) GetServiceKey() string {
	return "FP900025"
}
