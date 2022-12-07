package models

type SMP9CU0003I struct {
	MobileNo     string `json:"mobileNo" validate:"required"`
	CustomerNo   string `json:"customerNo" validate:"required"`
	SessionToken string `json:"sessionToken"`
	CitizenId    string `json:"citizenId"`
	DeviceUuid   string `json:"deviceUuid"`
	//TitleNameTh           string        `json:"titleNameTh"`
	//TitleNameEn           string        `json:"titleNameEn"`
	//NameEn                string        `json:"nameEn" validate:"required"`
	//LastnameEn            string        `json:"lastnameEn" validate:"required"`
	//Email                 string        `json:"email" validate:"required"`
	//Gender                string        `json:"gender"`
	//Occupation            string        `json:"occupation" validate:"required"`
	//OccupationOther       string        `json:"occupationOther"`
	//IncomeRange           string        `json:"incomeRange"`
	//SourceOfIncome        string        `json:"sourceOfIncome"`
	//SourceOfIncomeCountry string        `json:"sourceOfIncomeCountry"`
	//AddressType           []AddressType `json:"addressType"`
}

type SMP9CU0003O struct {
	CustomerNo string `json:"customerNo"`
}

func (i *SMP9CU0003I) GetServiceKey() string {
	return "MP9CU003"
}
