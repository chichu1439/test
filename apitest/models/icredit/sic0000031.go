package models

type IC000031I struct {
	CustomerId                    string `json:"customerId" description:"Customer id"`     // 客户号
	IdCountry                     string `json:"idCountry" description:"Id country"`       // 发证国家
	IdType                        string `json:"idType" description:"Id type"`             // 证件类型
	IdNum                         string `json:"IdNum" description:"Id number"`            // 证件号码
	CustomerName                  string `json:"customerName" description:"Customer name"` //客户名称
	CustomerElementClassification string `json:"customerElementClassification" description:"Customer element classification"`
	CustomerElement               string `json:"customerElement" description:"Customer element"`
	DateStart                     string `json:"dateStart" description:"Date start"`
	DateEnd                       string `json:"dateEnd" description:"Date end"`
	PageNum                       int    `json:"pageNum" description:"Page number"`
	PageRecordCount               int    `json:"pageRecordCount" description:"Page record count"`
}

type IC000031O struct {
	PageNum         int                   `json:"pageNum" description:"Page number"`
	PageRecordCount int                   `json:"pageRecordCount" description:"Page record count"`
	TotalCount      int                   `json:"totalCount" description:"Total count"`
	Records         []CustomerHistoryList `json:"records" description:"Records"`
}

type CustomerHistoryList struct {
	ChangeSerialNum string `json:"changeSerialNum" description:"Change serial number"` // 申请流水号
	ChangeDate      string `json:"changeDate" description:"Change date"`               // 申请日期
	ChangeSequnce   int    `json:"changeSequnce" description:"Change sequnce"`         // 变更序号
	ChangeTime      string `json:"changeTime" description:"Change time"`               // 变更时间
	CustomerId      string `json:"customerId" description:"Customer id"`               // 客户号
	CustomerName    string `json:"customerName" description:"Customer name"`           //客户名称
	CustomerElement string `json:"customerElement" description:"Customer element"`     // 客户要素
	ElementOldValue string `json:"elementOldValue" description:"Element old value"`    // 变更前值
	ElementNewValue string `json:"elementNewValue" description:"Element new value"`    // 变更后值
	ChangeUser      string `json:"changeUser" description:"Change user"`               // 变更用户
	FinalModifyDate string `json:"finalModifyDate" description:"Final modify date"`    // 最后修改日期
	FinalModifyTime string `json:"finalModifyTime" description:"Final modify time"`    // 最后修改时间
	FinalModifyOrgz string `json:"finalModifyOrgz" description:"Final modify organization"`    // 最后修改机构号
	FinalModifyUser string `json:"finalModifyUser" description:"Final modify user"`    // 最后修改用户号
}

func (*IC000031I) GetServiceKey() string {
	return "IC000031"
}
