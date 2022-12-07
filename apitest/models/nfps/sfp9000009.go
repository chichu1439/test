package models

//Inquire Participant Agency Relationship list
type FP900009I struct {
	//DirectPartName           string `json:"directPartName"`
	//IndirectPartName         string `json:"indirectPartName"`
	DirectPartClearingCode   string `json:"directPartClearingCode"`
	IndirectPartClearingCode string `json:"indirectPartClearingCode"`
	Currency                 string `json:"currency"`
	EffectiveDateStart       string `json:"effectiveDateStart"`
	EffectiveDateEnd         string `json:"effectiveDateEnd"`
	PageNo                   int    `json:"pageNo"`       //查询页数
	PageRecCount             int    `json:"pageRecCount"` //每页记录数
}
type FP900009O struct {
	TotalRecCount int             `json:"totalRecCount"`
	TotalPage     int             `json:"totalPage"`
	Records       []FP900009OItem `json:"records"`
}
type FP900009OItem struct {
	DirectPartClearingCode   string `json:"directPartClearingCode"`
	DirectPartName           string `json:"directPartName"`
	IndirectPartClearingCode string `json:"indirectPartClearingCode"`
	IndirectPartName         string `json:"indirectPartName"`
	Currency                 string `json:"currency"`
	EffectiveDate            string `json:"effectiveDate"`
	Remark                   string `json:"remark"`
	Status                   string `json:"status"`
}

func (*FP900009I) GetServiceKey() string {
	return "FP900009"
}
