package models

//Create Participant

type FP000006I struct {
	PartClearingCode     string `json:"partClearingCode" description:"Part clearing code" validate:"required"`
	PartType             string `json:"partType" description:"Part type" validate:"required"`
	PartBic              string `json:"partBic" description:"Part bic"`
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
	PartCertDocType1     string `json:"partCertDocType1" description:"Part cert doc type1" validate:"required,max=50"`
	PartCertDocNumber1   string `json:"partCertDocNumber1" description:"Part cert doc number1" validate:"required,max=120"`
	PartCertDocType2     string `json:"partCertDocType2" description:"Part cert doc type2" validate:"required,max=50"`
	PartCertDocNumber2   string `json:"partCertDocNumber2" description:"Part cert doc number2" validate:"required,max=120"`
	PartIpAddress        string `json:"partIpAddress" description:"Part ip address"`
}

type FP000006O struct {
	PartClearingCode string `json:"partClearingCode" description:"Part clearing code"`
}

func (*FP000006I) GetServiceKey() string {
	return "FP000006"
}
