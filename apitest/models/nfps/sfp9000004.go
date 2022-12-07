package models

// inquire Participant detail
type FP900004I struct {
	PartClearingCode string    `json:"partClearingCode"`
	PartType         string    `json:"partType"`
	PartList         []PartInf `json:"partList"` // for multi query
}
type PartInf struct {
	PartClearingCode string `json:"partClearingCode"`
	PartType         string `json:"partType"`
}
type FP900004O struct {
	PartClearingCode     string          `json:"partClearingCode"`
	PartType             string          `json:"partType"`
	PartBic              string          `json:"partBic"`
	PartName             string          `json:"partName"`
	PartShortName        string          `json:"partShortName"`
	PartEnglishName      string          `json:"partEnglishName"`
	PartEnglishShortName string          `json:"partEnglishShortName"`
	PartNationality      string          `json:"partNationality"`
	PartRegion           string          `json:"partRegion"`
	PartContactNumber1   string          `json:"partContactNumber1"`
	ContactPerson1       string          `json:"contactPerson1"`
	PartContactNumber2   string          `json:"partContactNumber2"`
	ContactPerson2       string          `json:"contactPerson2"`
	PartContactAddress   string          `json:"partContactAddress"`
	PartPostalCode       string          `json:"partPostalCode"`
	PartEmail            string          `json:"partEmail"`
	PartCertDocType1     string          `json:"partCertDocType1"`
	PartCertDocNumber1   string          `json:"partCertDocNumber1"`
	PartCertDocType2     string          `json:"partCertDocType2"`
	PartCertDocNumber2   string          `json:"partCertDocNumber2"`
	PartIpAddress        string          `json:"partIpAddress"`
	AccountingType       string          `json:"accountingType"`
	AccountingBatchAmt   float64         `json:"accountingBatchAmt"`
	Status               string          `json:"status"`
	SignStatus           string          `json:"signStatus"`
	Records              []FP900004OList `json:"records"`
}
type FP900004OList struct {
	PartClearingCode     string  `json:"partClearingCode"`
	PartType             string  `json:"partType"`
	PartBic              string  `json:"partBic"`
	PartName             string  `json:"partName"`
	PartShortName        string  `json:"partShortName"`
	PartEnglishName      string  `json:"partEnglishName"`
	PartEnglishShortName string  `json:"partEnglishShortName"`
	PartNationality      string  `json:"partNationality"`
	PartRegion           string  `json:"partRegion"`
	PartContactNumber1   string  `json:"partContactNumber1"`
	ContactPerson1       string  `json:"contactPerson1"`
	PartContactNumber2   string  `json:"partContactNumber2"`
	ContactPerson2       string  `json:"contactPerson2"`
	PartContactAddress   string  `json:"partContactAddress"`
	PartPostalCode       string  `json:"partPostalCode"`
	PartEmail            string  `json:"partEmail"`
	PartCertDocType1     string  `json:"partCertDocType1"`
	PartCertDocNumber1   string  `json:"partCertDocNumber1"`
	PartCertDocType2     string  `json:"partCertDocType2"`
	PartCertDocNumber2   string  `json:"partCertDocNumber2"`
	PartIpAddress        string  `json:"partIpAddress"`
	AccountingType       string  `json:"accountingType"`
	AccountingBatchAmt   float64 `json:"accountingBatchAmt"`
	Status               string  `json:"status"`
	SignStatus           string  `json:"signStatus"`
}

func (*FP900004I) GetServiceKey() string {
	return "FP900004"
}