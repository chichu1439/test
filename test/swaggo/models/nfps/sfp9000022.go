package models

type FP900022I struct {
	BatchSequence string `json:"batchSequence"`
	BatchDate     string `json:"batchDate"`
	BatchTime     string `json:"batchTime"`
	BatchType     string `json:"batchType"`
	ExecuteId     string `json:"executeId"`
	ProjectId     string `json:"projectId"`
	FlowId        string `json:"flowId"`
	JobId         string `json:"jobId"`
	ErrorRecord   string `json:"errorRecord"`
	ErrSource     string `json:"errSource"`
	ErrorCode     string `json:"errorCode"`
	ServiceCode   string `json:"serviceCode"`
}
type FP900022O struct {
	Status string `json:"status"`
}

func (*FP900022I) GetServiceKey() string {
	return "FP900022"
}
