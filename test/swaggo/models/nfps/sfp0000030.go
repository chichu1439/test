package models

type FP000030I struct {
	ReportType              string `json:"reportType"`
	ParticipantClearingCode string `json:"participantClearingCode"`
	ParticipantName         string `json:"participantName"`
	GeneratedDate           string `json:"generatedDate" validate:"required"`
	PageNo                  int    `json:"pageNo"`
	PageSize                int    `json:"pageSize"`
}
type FP000030O struct {
	TotalCount int             `json:"totalCount"`
	Records    []FP000030OItem `json:"records"`
}
type FP000030OItem struct {
	ReportType    string `json:"reportType"`
	ReportDate    string `json:"reportDate"`
	ReportName    string `json:"reportName"`
	GeneratedDate string `json:"generatedDate"`
}

func (*FP000030I) GetServiceKey() string {
	return "FP000030"
}
