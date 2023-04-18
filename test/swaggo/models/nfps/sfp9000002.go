package models

//Edit Participant
type FP900002I struct {
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
	PartCertDocType1     string `json:"partCertDocType1"`
	PartCertDocNumber1   string `json:"partCertDocNumber1"`
	PartCertDocType2     string `json:"partCertDocType2"`
	PartCertDocNumber2   string `json:"partCertDocNumber2"`
	PartIpAddress        string `json:"partIpAddress"`
	AccountingType       string `json:"accountingType"`
	Status               string `json:"status"`
}
type FP900002O struct {
	PartClearingCode string `json:"partClearingCode"`
}

func (*FP900002I) GetServiceKey() string {
	return "FP900002"
}
