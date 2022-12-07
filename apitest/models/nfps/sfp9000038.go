package models

type FP900038I struct {
	ParticipantCode string `json:"participantCode"`
	GeneratedDate   string `json:"generatedDate"`
	ReportType      string `json:"reportType"`
	PageNo          int    `json:"pageNo"`
	PageSize        int    `json:"pageSize"`
}
type FP900038IItem struct {
	ReportType    string `json:"reportType"`
	ReportDate    string `json:"reportDate"`
	ReportName    string `json:"reportName"`
	GeneratedDate string `json:"generatedDate"`
}
type FP900038O struct {
	TotalCount int             `json:"totalCount"`
	Records    []FP900038IItem `json:"records"`
}

func (*FP900038I) GetServiceKey() string {
	return "FP900038"
}
