package models

type CM000023I struct {
	ListReportRegister []ReportRegister `json:"listReportRegister" description:"List report register"`
	DateOfReport       string           `json:"dateOfReport" description:"Date of report"`
	ReportGenerateDate string           `json:"reportGenerateDate" description:"Report generate date"`
	ReportGenerateTime string           `json:"reportGenerateTime" description:"Report generate time"`
	FinalModifyOrgzNum string           `json:"finalModifyOrgzNum" description:"Final modify orgz number"`
	FinalModifyUserNum string           `json:"finalModifyUserNum" description:"Final modify user number"`
}

type ReportRegister struct {
	ReportClassification string `json:"reportClassification" description:"Report classification"`
	ReportFileId         string `json:"reportFileId" description:"Report file id"`
	ReportFileKey        string `json:"reportFileKey" description:"Report file key"`
	ReportFileType       string `json:"reportFileType" description:"Report file type"`
	ReportName           string `json:"reportName" description:"Report name"`
	ReportType           string `json:"reportType" description:"Report type"`
}

type CM000023O struct {
}

func (*CM000023I) GetServiceKey() string {
	return "CM000023"
}
