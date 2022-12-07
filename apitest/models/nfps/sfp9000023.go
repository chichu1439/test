package models

type FP900023I struct {
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
type FP900023O struct {
	Status string `json:"status"`
	Id     int    `json:"id"`
}

func (*FP900023I) GetServiceKey() string {
	return "FP900023"
}
