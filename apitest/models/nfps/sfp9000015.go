package models

//Settlement Account detail inquiry
type FP900015I struct {
	PartClearingCode string `json:"partClearingCode" validate:"required"`
	Currency         string `json:"currency"`
	PageNo           int    `json:"pageNo"`       //查询页数
	PageRecCount     int    `json:"pageRecCount"` //每页记录数
	Limit            int    `json:"limit"`        //优先使用偏移量查询
	Offset           int    `json:"offset"`
}
type FP900015O struct {
	TotalRecCount int             `json:"totalRecCount"`
	TotalPage     int             `json:"totalPage"`
	Records       []FP900015OItem `json:"records"`
}

type FP900015OItem struct {
	PartClearingCode string  `json:"partClearingCode"`
	PartName         string  `json:"partName"`
	PartType         string  `json:"partType"`
	Currency         string  `json:"currency"`
	Balance          float64 `json:"balance"`
	AlertBalance     float64 `json:"alertBalance"`
	AlertMinuteTime  int     `json:"alertMinuteTime"`
	HighWaterMark    float64 `json:"HighWaterMark"`
	LowWaterMark     float64 `json:"LowWaterMark"`
	Status           string  `json:"status"`
}

func (*FP900015I) GetServiceKey() string {
	return "FP900015"
}
