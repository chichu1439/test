package models

//Settlement Account inquiry
type FP900084I struct {
	PartClearingCode string `json:"partClearingCode"`
	PartType         string `json:"partType"`
	Currency         string `json:"currency"`
	PageNo           int    `json:"pageNo"`       //查询页数
	PageRecCount     int    `json:"pageRecCount"` //每页记录数
}
type FP900084O struct {
	TotalRecCount int             `json:"totalRecCount"`
	TotalPage     int             `json:"totalPage"`
	FP900084OItem []FP900084OItem `json:"fp900013OItem"`
}
type FP900084OItem struct {
	PartClearingCode string  `json:"partClearingCode"`
	Name             string  `json:"name"`
	PartType         string  `json:"partType"`
	Currency         string  `json:"currency"`
	SettAcctNo       string  `json:"settAcctNo"`
	SettAcctBalance  float64 `json:"settAcctBalance"` //Deprecated
	AcctBalance      float64 `json:"acctBalance"`     //账户可用余额
	IpTotalBalance   float64 `json:"ipTotalBalance"`  //间参合计余额
	AlertBalance     float64 `json:"alertBalance"`
	AlertMinuteTime  int     `json:"alertMinuteTime"`
	HightWaterMark   float64 `json:"HightWaterMark"`
	LowWaterMark     float64 `json:"LowWaterMark"`
	Status           string  `json:"status"`
}

func (*FP900084I) GetServiceKey() string {
	return "FP900084"
}
