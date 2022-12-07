package models

//Inquire Participant relation list

type FP000007I struct {
	//DirectPartClearingCode   string `json:"directPartClearingCode"`
	//IndirectPartClearingCode string `json:"indirectPartClearingCode"`
	PartClearingCode   string `json:"partClearingCode" description:"Part clearing code"`
	Currency           string `json:"currency" description:"Currency"`
	EffectiveDateStart string `json:"effectiveDateStart" description:"Effective date start"`
	EffectiveDateEnd   string `json:"effectiveDateEnd" description:"Effective date end"`
	PageNo             int    `json:"pageNo" description:"Page no"`              //查询页数
	PageRecCount       int    `json:"pageRecCount" description:"Page rec count"` //每页记录数
}

type FP000007O struct {
	TotalRecCount int             `json:"totalRecCount" description:"Total rec count"`
	TotalPage     int             `json:"totalPage" description:"Total page"`
	Records       []FP000007OItem `json:"records" description:"Records"`
}

type FP000007OItem struct {
	DirectPartClearingCode   string `json:"directPartClearingCode" description:"Direct part clearing code"`
	PartTypeDp               string `json:"partTypeDp" description:"Part type dp"`
	DirectPartName           string `json:"directPartName" description:"Direct part name"`
	IndirectPartClearingCode string `json:"indirectPartClearingCode" description:"Indirect part clearing code"`
	PartTypeIp               string `json:"partTypeIp" description:"Part type ip"`
	IndirectPartName         string `json:"indirectPartName" description:"Indirect part name"`
	Currency                 string `json:"currency" description:"Currency"`
	EffectiveDate            string `json:"effectiveDate" description:"Effective date"`
	Remark                   string `json:"remark" description:"Remark"`
	Status                   string `json:"status" description:"Status"`
}

func (*FP000007I) GetServiceKey() string {
	return "FP000007"
}
