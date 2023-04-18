package models

type FP900036I struct {
	Records []FP900036IItem `json:"records"`
}
type FP900036IItem struct {
	DateOfReport         string `json:"dateOfReport"`
	ReportName           string `json:"reportName"`
	ReportType           string `json:"reportType"`
	ReportClassification string `json:"reportClassification"`
	ReportGenerateTime   string `json:"reportGenerateTime"`
	ReportGenerateDate   string `json:"reportGenerateDate"`
	ReportFileType       string `json:"reportFileType"`
	ReportFileId         string `json:"reportFileId"`
	ReportFileKey        string `json:"reportFileKey"`
}
type FP900036O struct {
	Status string
}

func (*FP900036I) GetServiceKey() string {
	return "FP900036"
}
