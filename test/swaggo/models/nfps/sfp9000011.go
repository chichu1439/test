package models

//inquire RTGS Participant info
type FP900011I struct {
	PartClearingCode string `json:"partClearingCode"`
	PartType         string `json:"partType"`
	Flag             string `json:"flag"` //1-RTGS参与者信息，2-RTGS参与者清算账号信息
	Currency         string `json:"currency"`
}
type FP900011O struct {
	PartClearingCode  string             `json:"partClearingCode"`
	PartType          string             `json:"partType"`
	FP000011OPartItem FP000011OPartItem  `json:"fp000011OPartItem"`
	Records           []FP900011OCurItem `json:"records"`
}
type FP900011OCurItem struct {
	PartCurrency    string  `json:"partCurrency"`
	PartSettleAcct  string  `json:"partSettleAcct"`
	SettAcctBalance float64 `json:"settAcctBalance"`
}
type FP000011OPartItem struct {
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
}

func (*FP900011I) GetServiceKey() string {
	return "FP900011"
}
