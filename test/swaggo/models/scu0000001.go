package models

type SCU0000001I struct {
	IdNum    string `json:"idNum" description:"Id number" validate:"required"`
	IdType   string `json:"idType" validate:"required" description:"id Type 101-Citizen ID 104-Passport"`
	IdNation string `json:"idNation" description:"Id nation"`
}

type SCU0000001O struct {
	NewFlag              int    `json:"newFlag" description:"New flag"`
	CustomerId           string `json:"customerId" description:"Customer id"`
	CustomerName         string `json:"customerName" description:"Customer Name"`
	IdNum                string `json:"idNum" description:"Id number"`
	IdType               string `json:"idType" description:"id Type 101-Citizen ID 104-Passport"`
	IdNation             string `json:"idNation" description:"Id nation"`
	CustomerStatus       string `json:"customerStatus" description:"Customer Status"`
	MutipleCustomerFlage string `json:"mutipleCustomerFlage" description:"Mutiple Customer Flage(Y-yes;N-no)"`
}

func (*SCU0000001I) GetServiceKey() string {
	return "CU000001"
}
