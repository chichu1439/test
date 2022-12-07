//Version: v1.0.0.0
package models

type IN000001I struct {
	AccountingItemNumber string `json:"accountingItemNumber" description:"Accounting item number"` //科目号
	SerialNumber         string `json:"serialNumber" description:"Serial number"`                  //内部户顺序号
	AccountName          string `json:"accountName" description:"Account name"`                    //账户名称
	PageNo               string `json:"pageNo" description:"Page no"`                              //查询页数
	PageRecCount         string `json:"pageRecCount" description:"Page rec count"`                 //每页记录数
}

type IN000001O struct {
	TotalRecCount int             `json:"totalRecCount" description:"Total rec count"`
	TotalPage     int             `json:"totalPage" description:"Total page"`
	Item          []IN000001OItem `json:"item" description:"Item"`
}

type IN000001OItem struct {
	AccountingItemNumber string `json:"accountingItemNumber" description:"Accounting item number"` //科目号
	SerialNumber         string `json:"serialNumber" description:"Serial number"`                  //内部户顺序号
	AccountName          string `json:"accountName" description:"Account name"`                    //账户名称
	Status               string `json:"status" description:"Status"`                               //记录状态P-待生效，N-正常，S-停用
}

func (*IN000001I) GetServiceKey() string {
	return "IN000001"
}
