package models

//inquire RTGS Participant information

type FP000011I struct {
	PartClearingCode string `json:"partClearingCode" description:"Part clearing code" validate:"required"`
	PartType         string `json:"partType" description:"Part type"`
	Flag             string `json:"flag" description:"Flag"` //1-RTGS参与者信息，2-RTGS参与者清算账号信息
	Currency         string `json:"currency" description:"Currency"`
}

type FP000011O struct {
	FP000011OPartItem FP000011OPartItem  `json:"fp000011OPartItem" description:"F p000011 o part item"`
	Records           []FP000011OCurItem `json:"records" description:"Records"`
}

type FP000011OPartItem struct {
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
	PartCertDocType1     string `json:"partCertDocType1" description:"Part cert doc type1"`
	PartCertDocNumber1   string `json:"partCertDocNumber1" description:"Part cert doc number1"`
	PartCertDocType2     string `json:"partCertDocType2" description:"Part cert doc type2"`
	PartCertDocNumber2   string `json:"partCertDocNumber2" description:"Part cert doc number2"`
	PartIpAddress        string `json:"partIpAddress" description:"Part ip address"`
}

type FP000011OCurItem struct {
	PartCurrency    string  `json:"partCurrency" description:"Part currency"`
	PartSettleAcct  string  `json:"partSettleAcct" description:"Part settle acct"`
	SettAcctBalance float64 `json:"settAcctBalance" description:"Sett acct balance"`
}

func (*FP000011I) GetServiceKey() string {
	return "FP000011"
}
