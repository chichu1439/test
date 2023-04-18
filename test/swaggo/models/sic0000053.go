package models

type IC000053I struct {
	ReportType           string `json:"reportType" validate:"required,max=5" description:"ReportType"`             //报表类型
	ReportClassification string `json:"reportClassification" validate:"max=40" description:"ReportClassification"` //报表类别
	ReportName           string `json:"reportName" validate:"required,max=200" description:"ReportName"`           //报表名称
	DateFrom             string `json:"dateFrom" description:"DateFrom"`                                           //开始日期
	DateTo               string `json:"dateTo" description:"DateTo"`                                               //截止日期
	OrderByDate          string `json:"orderByDate" validate:"max=10" description:"OrderByDate"`                   //报表文件生成日期排序条件
	OrderByTime          string `json:"orderByTime" validate:"max=10" description:"OrderByTime"`                   //报表文件生成时间排序条件
	OrderByDateOfReport  string `json:"orderByDateOfReport" validate:"max=10" description:"OrderByDateOfReport"`   //报表数据产生时间排序条件
	OrderByFileType      string `json:"orderByFileType" validate:"max=10" description:"OrderByFileType"`           //报表文件格式排序条件
	PageNum              int    `json:"pageNum" description:"PageNum"`                                             //当前页
	PageRecordCount      int    `json:"pageRecordCount" description:"PageRecordCount"`
}

type IC000053O struct {
	PageNum   int                `json:"pageNum" description:"PageNum" `     //当前页
	PageTotal int                `json:"pageTotal" description:"PageTotal" ` //总条数
	Records   []ReportRegisterBo `json:"records" description:"Records" `
}

func (*IC000053I) GetServiceKey() string {
	return "IC000053"
}
