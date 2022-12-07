package models

//Inquire Participant list
type FP900003I struct {
	PartClearingCode string        `json:"partClearingCode"`
	PartType         string        `json:"partType"`
	PartBic          string        `json:"partBic"`
	PartName         string        `json:"partName"`
	PartShortName    string        `json:"partShortName"`
	PartCertDoc      []PartCertDoc `json:"partCertDoc"`
	PageNo           int           `json:"pageNo"`       //查询页数
	PageRecCount     int           `json:"pageRecCount"` //每页记录数
}
type PartCertDoc struct {
	PartCertDocType   string `json:"partCertDocType"`
	PartCertDocNumber string `json:"partCertDocNumber"`
}
type FP900003O struct {
	TotalRecCount int             `json:"totalRecCount"`
	TotalPage     int             `json:"totalPage"`
	Records       []FP900003OItem `json:"records"`
}
type FP900003OItem struct {
	PartClearingCode     string `json:"partClearingCode"`
	PartType             string `json:"partType"`
	PartBic              string `json:"partBic"`
	PartName             string `json:"partName"`
	PartShortName        string `json:"partShortName"`
	PartEnglishName      string `json:"partEnglishName"`
	PartEnglishShortName string `json:"partEnglishShortName"`
	Status               string `json:"status"`
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

func (*FP900003I) GetServiceKey() string {
	return "FP900003"
}
