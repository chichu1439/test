package models

type FP930003I struct {
	ParticipantCode string   `json:"participantCode"`
	GeneratedDate   string   `json:"generatedDate"`
	ReportType      []string `json:"reportType"`
	ParticipantName string   `json:"participantName"`
	PageNo          int      `json:"pageNo"`
	PageSize        int      `json:"pageSize"`
}
type FP930003IItem struct {
	ReportType    string `json:"reportType"`
	ReportDate    string `json:"reportDate"`
	ReportName    string `json:"reportName"`
	GeneratedDate string `json:"generatedDate"`
}
type FP930003O struct {
	TotalCount int             `json:"totalCount"`
	Records    []FP930003IItem `json:"records"`
}

func (*FP930003I) GetServiceKey() string {
	return "FP930003"
}
