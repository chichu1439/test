package models

type FP900008I struct {
	PartClearingCode string `json:"partClearingCode"`
	PartType         string `json:"partType"`
	PartName         string `json:"partName"`
	PageNo           string `json:"pageNo"`       //查询页数
	PageRecCount     string `json:"pageRecCount"` //每页记录数
}
type FP900008O struct {
	TotalRecCount int             `json:"totalRecCount"`
	TotalPage     int             `json:"totalPage"`
	Records       []FP900008OItem `json:"records"`
}
type FP900008OItem struct {
	PartClearingCode     string `json:"partClearingCode"`
	PartType             string `json:"partType"`
	PartBic              string `json:"partBic"`
	PartName             string `json:"partName"`
	PartShortName        string `json:"partShortName"`
	PartEnglishName      string `json:"partEnglishName"`
	PartEnglishShortName string `json:"partEnglishShortName"`
	PartNationality      string `json:"partNationality"`
	PartRegion           string `json:"partRegion"`
	PartContactNumber1   string `json:"partContactNumber1"`
	ContactPerson1       string `json:"contactPerson1"`
	PartContactNumber2   string `json:"partContactNumber2"`
	ContactPerson2       string `json:"contactPerson2"`
	PartContactAddress   string `json:"partContactAddress"`
	PartPostalCode       string `json:"partPostalCode"`
	PartEmail            string `json:"partEmail"`
	Status               string `json:"status"`
}

func (*FP900008I) GetServiceKey() string {
	return "FP900008"
}
