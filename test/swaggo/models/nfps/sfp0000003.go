package models

// swagger:parameters editParticipantDetailRequest

type FP000003I struct {
	PartClearingCode     string `json:"partClearingCode" description:"Part clearing code" validate:"required"`
	PartType             string `json:"partType" description:"Part type"`
	PartName             string `json:"partName" description:"Part name"`
	PartShortName        string `json:"partShortName" description:"Part short name"`
	PartEnglishName      string `json:"partEnglishName" description:"Part english name"`
	PartEnglishShortName string `json:"partEnglishShortName" description:"Part english short name"`
	PartNationality      string `json:"partNationality" description:"Part nationality"`
	PartRegion           string `json:"partRegion" description:"Part region"`
	PartContactNumber1   string `json:"partContactNumber1" description:"Part contact number1"`
	ContactPerson1       string `json:"contactPerson1" description:"Contact person1"`
	PartContactNumber2   string `json:"partContactNumber2" description:"Part contact number2"`
	ContactPerson2       string `json:"contactPerson2" description:"Contact person2"`
	PartContactAddress   string `json:"partContactAddress" description:"Part contact address"`
	PartPostalCode       string `json:"partPostalCode" description:"Part postal code"`
	PartEmail            string `json:"partEmail" description:"Part email"`
	Status               string `json:"status" description:"Status"` //P-待生效，N-正常，S-停用
	//PartCertDocType1     string `json:"partCertDocType1"`
	//PartCertDocNumber1   string `json:"partCertDocNumber1"`
	//PartCertDocType2     string `json:"partCertDocType2"`
	//PartCertDocNumber2   string `json:"partCertDocNumber2"`
	//PartIpAddress        string `json:"partIpAddress"`
}

// swagger:response editParticipantDetailResponse

type FP000003O struct {
}

func (*FP000003I) GetServiceKey() string {
	return "FP000003"
}
