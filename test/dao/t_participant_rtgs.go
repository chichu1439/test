package dao

type TPartRtgs struct {
	PartClearingCode     string `json:"part_clearing_code"`
	PartType             string `json:"part_type"`
	PartBic              string `json:"part_b_ic"`
	PartName             string `json:"part_name"`
	PartShortName        string `json:"part_short_name"`
	PartEnglishName      string `json:"part_english_name"`
	PartEnglishShortName string `json:"part_english_short_name"`
	PartNationality      string `json:"part_nationality"`
	PartRegion           string `json:"part_region"`
	PartContactNumber1   string `json:"part_contact_number_1"`
	ContactPerson1       string `json:"contact_person_1"`
	PartContactNumber2   string `json:"part_contact_number_2"`
	ContactPerson2       string `json:"contact_person_2"`
	PartContactAddress   string `json:"part_contact_address"`
	PartPostalCode       string `json:"part_postal_code"`
	PartEmail            string `json:"part_email"`
	PartCertDocType1     string `json:"part_cert_doc_type_1"`
	PartCertDocNumber1   string `json:"part_cert_doc_number_1"`
	PartCertDocType2     string `json:"part_cert_doc_type_2"`
	PartCertDocNumber2   string `json:"part_cert_doc_number_2"`
	PartIpAddress        string `json:"part_ip_address"`
	Status               string `json:"status"`
	CreateTime           string `json:"create_time"`
	UpdateTime           string `json:"update_time"`
}

func (t *TPartRtgs) TableName() string {
	return "t_part_rtgs"
}
