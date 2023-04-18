package models

// swagger:parameters inquireParticipantListRequest
// in: body

type FP000001I struct {
	PartClearingCode string `json:"partClearingCode" description:"Part clearing code"`
	PartType         string `json:"partType" description:"Part type"`
	PartName         string `json:"partName" description:"Part name"`
	PageNo           int    `json:"pageNo" description:"Page no"`              //查询页数
	PageRecCount     int    `json:"pageRecCount" description:"Page rec count"` //每页记录数
}

// swagger:response inquireParticipantListResponse

type FP000001O struct {
	TotalRecCount int             `json:"totalRecCount" description:"Total rec count"`
	TotalPage     int             `json:"totalPage" description:"Total page"`
	Records       []FP000001OItem `json:"records" description:"Records"`
}

// swagger:model FP000001OItem

type FP000001OItem struct {
	PartClearingCode     string `json:"partClearingCode" description:"Part clearing code"`
	PartType             string `json:"partType" description:"Part type"`
	PartBic              string `json:"partBic" description:"Part bic"`
	PartName             string `json:"partName" description:"Part name"`
	PartShortName        string `json:"partShortName" description:"Part short name"`
	PartEnglishName      string `json:"partEnglishName" description:"Part english name"`
	PartEnglishShortName string `json:"partEnglishShortName" description:"Part english short name"`
	Status               string `json:"status" description:"Status"` //P-待生效，N-正常，S-停用
	//PartNationality      string `json:"partNationality"`
	//PartRegion           string `json:"partRegion"`
	//PartContactNumber1   string `json:"partContactNumber1"`
	//ContactPerson1       string `json:"contactPerson1"`
	//PartContactNumber2   string `json:"partContactNumber2"`
	//ContactPerson2       string `json:"contactPerson2"`
	//PartContactAddress   string `json:"partContactAddress"`
	//PartPostalCode       string `json:"partPostalCode"`
	//PartEmail            string `json:"partEmail"`
}

func (*FP000001I) GetServiceKey() string {
	return "FP000001"
}
