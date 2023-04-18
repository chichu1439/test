package models

//inquire other Participant detail

type FP000010I struct {
	PartClearingCode string `json:"partClearingCode" description:"Part clearing code" validate:"required"`
	PartType         string `json:"partType" description:"Part type"`
}

type FP000010O struct {
	PartClearingCode     string `json:"partClearingCode" description:"Part clearing code"`
	PartType             string `json:"partType" description:"Part type"`
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
	Status               string `json:"status" description:"Status"` //P-待生效，N-正常，S-停用
}

func (*FP000010I) GetServiceKey() string {
	return "FP000010"
}
