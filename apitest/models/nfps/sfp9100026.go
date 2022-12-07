package models

type FP910026I struct {
	ParticipantCode string `json:"participantCode"`
	GeneratedDate   string `json:"generatedDate"`
	ReportType      string `json:"reportType"`
	PageNo          int    `json:"pageNo"`
	PageSize        int    `json:"pageSize"`
}
type FP910026IItem struct {
	ReportType    string `json:"reportType"`
	ReportDate    string `json:"reportDate"`
	ReportName    string `json:"reportName"`
	GeneratedDate string `json:"generatedDate"`
}
type FP910026O struct {
	TotalCount int             `json:"totalCount"`
	Records    []FP910026IItem `json:"records"`
}

func (*FP910026I) GetServiceKey() string {
	return "FP910026"
}
