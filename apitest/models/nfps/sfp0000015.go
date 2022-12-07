package models

//Settlement Account inquiry

type FP000015I struct {
	PartClearingCode string `json:"partClearingCode" description:"Part clearing code"`
	PartType         string `json:"partType" description:"Part type"`
	Currency         string `json:"currency" description:"Currency"`
	PageNo           int    `json:"pageNo" description:"Page no"`              //查询页数
	PageRecCount     int    `json:"pageRecCount" description:"Page rec count"` //每页记录数
}

type FP000015O struct {
	TotalRecCount int             `json:"totalRecCount" description:"Total rec count"`
	TotalPage     int             `json:"totalPage" description:"Total page"`
	Records       []FP000015OItem `json:"records" description:"Records"`
}

type FP000015OItem struct {
	PartClearingCode string  `json:"partClearingCode" description:"Part clearing code"`
	Name             string  `json:"name" description:"Name"`
	Currency         string  `json:"currency" description:"Currency"`
	SettAcctNo       string  `json:"settAcctNo" description:"Sett acct no"`
	AcctBalance      float64 `json:"acctBalance" description:"Acct balance"`        //账户余额
	DpBalance        float64 `json:"dpBalance" description:"Dp balance"`            //直参余额/可用余额
	IpTotalBalance   float64 `json:"ipTotalBalance" description:"Ip total balance"` //间参合计余额
	AlertBalance     float64 `json:"alertBalance" description:"Alert balance"`
	AlertMinuteTime  int     `json:"alertMinuteTime" description:"Alert minute time"`
	HighWaterMark    float64 `json:"highWaterMark" description:"high Water Mark"`
	LowWaterMark     float64 `json:"lowWaterMark" description:"low Wate rMark"`
	Status           string  `json:"status" description:"Status"`
}

func (*FP000015I) GetServiceKey() string {
	return "FP000015"
}
