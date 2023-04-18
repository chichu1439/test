package models

//Settlement Account detail inquiry

type FP000018I struct {
	PartClearingCode string `json:"partClearingCode" description:"Part clearing code" validate:"required"`
	Currency         string `json:"currency" description:"Currency"`
	PageNo           int    `json:"pageNo" description:"Page no"`              //查询页数
	PageRecCount     int    `json:"pageRecCount" description:"Page rec count"` //每页记录数
}

type FP000018O struct {
	TotalRecCount int             `json:"totalRecCount" description:"Total rec count"`
	TotalPage     int             `json:"totalPage" description:"Total page"`
	Records       []FP000018OItem `json:"records" description:"Records"`
}

type FP000018OItem struct {
	PartClearingCode string  `json:"partClearingCode" description:"Part clearing code"`
	PartName         string  `json:"partName" description:"Part name"`
	PartType         string  `json:"partType" description:"Part type"`
	Currency         string  `json:"currency" description:"Currency"`
	Balance          float64 `json:"balance" description:"Balance"`
	AlertBalance     float64 `json:"alertBalance" description:"Alert balance"`
	AlertMinuteTime  int     `json:"alertMinuteTime" description:"Alert minute time"`
	HighWaterMark    float64 `json:"highWaterMark" description:"high Water Mark"`
	LowWaterMark     float64 `json:"lowWaterMark" description:"low Wate rMark"`
	Status           string  `json:"status" description:"Status"`
}

func (*FP000018I) GetServiceKey() string {
	return "FP000018"
}
