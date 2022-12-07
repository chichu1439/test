package models

type SMP9CU0004I struct {
	CustomerNo string `json:"customerNo"`
	MobileNo   string `json:"mobileNo"`
}

type SMP9CU0004O struct {
	CustomerNo   string `json:"customerNo"`
	MobileNo     string `json:"mobileNo"`
	CitizenId    string `json:"citizenId"`
	Email        string `json:"email"`
	Language     string `json:"language"`
	SessionToken string `json:"sessionToken"`
	LoginMethod  int    `json:"loginMethod"`
}

func (i *SMP9CU0004I) GetServiceKey() string {
	return "MP9CU004"
}
