package models

type CM000037I struct {
	ReportType           string `json:"reportType" description:"ReportType"`                                       //报表类型
	ReportClassification string `json:"reportClassification" validate:"max=40" description:"ReportClassification"` //报表类别
	ReportName           string `json:"reportName" description:"ReportName"`                                       //报表名称
	DateFrom             string `json:"dateFrom" description:"DateFrom"`                                           //开始日期
	DateTo               string `json:"dateTo" description:"DateTo"`                                               //截止日期
	OrderByDate          string `json:"orderByDate" validate:"max=10" description:"OrderByDate"`                   //报表文件生成日期排序条件
	OrderByTime          string `json:"orderByTime" validate:"max=10" description:"OrderByTime"`                   //报表文件生成时间排序条件
	OrderByDateOfReport  string `json:"orderByDateOfReport" validate:"max=10" description:"OrderByDateOfReport"`   //报表数据产生时间排序条件
	OrderByFileType      string `json:"orderByFileType" validate:"max=10" description:"OrderByFileType"`           //报表文件格式排序条件
	PageNum              int    `json:"pageNum" description:"PageNum"`                                             //当前页
	PageSize             int    `json:"pageSize" description:"PageSize"`                                           //每页总数
}

type CM000037O struct {
	PageNum   int                `json:"pageNum" description:"PageNum" `     //当前页
	PageTotal int                `json:"pageTotal" description:"PageTotal" ` //总条数
	Records   []ReportRegisterBo `json:"Records" description:"Records" `
}

type ReportRegisterBo struct {
	ReportName         string `json:"reportName" description:"Report name"`                  //报表名称
	ReportGenerateDate string `json:"reportGenerateDate" description:"Report generate date"` //报表文件生成日期
	ReportGenerateTime string `json:"reportGenerateTime" description:"Report generate time"` //报表文件生成时间
	DateOfReport       string `json:"dateOfReport" description:"Date of report"`             //报表数据产生时间
	ReportFileType     string `json:"reportFileType" description:"Report file type"`         //报表文件格式
	ReportFileKey      string `json:"reportFileKey" description:"Report file key"`           //报表文件key
}

func (*CM000037I) GetServiceKey() string {
	return "CM000037"
}
