package models

//Create Part
type FP900001I struct {
	PartClearingCode     string `json:"partClearingCode" validate:"required"`
	PartType             string `json:"partType" validate:"required"`
	PartBic              string `json:"partBic" validate:"required"`
	PartName             string `json:"partName" validate:"required"`
	PartShortName        string `json:"partShortName" validate:"required"`
	PartEnglishName      string `json:"partEnglishName" validate:"required"`
	PartEnglishShortName string `json:"partEnglishShortName" validate:"required"`
	PartNationality      string `json:"partNationality" validate:"required"`
	PartRegion           string `json:"partRegion" validate:"required"`
	PartContactNumber1   string `json:"partContactNumber1" validate:"required"`
	ContactPerson1       string `json:"contactPerson1" validate:"required"`
	PartContactNumber2   string `json:"partContactNumber2"`
	ContactPerson2       string `json:"contactPerson2"`
	PartContactAddress   string `json:"partContactAddress" validate:"required"`
	PartPostalCode       string `json:"partPostalCode" validate:"required"`
	PartEmail            string `json:"partEmail" validate:"required"`
	PartCertDocType1     string `json:"partCertDocType1" validate:"required"`
	PartCertDocNumber1   string `json:"partCertDocNumber1" validate:"required"`
	PartCertDocType2     string `json:"partCertDocType2"`
	PartCertDocNumber2   string `json:"partCertDocNumber2"`
	PartIpAddress        string `json:"partIpAddress" validate:"required"`
	Status               string `json:"status"`
}

type FP900001O struct {
	PartClearingCode string `json:"partClearingCode"`
}

func (*FP900001I) GetServiceKey() string {
	return "FP900001"
}
