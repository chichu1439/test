package models

type SCU0000006I struct {
	CustomerId string `json:"customerId" description:"Customer id" validate:"required"`
}

type SCU0000006O struct {
	CustomerStatus string `json:"customerStatus" description:"Customer Status:0:Normal 4:Pending 9:Cancel "`
	IdExpireDate   string `json:"idExpireDate" description:"Id expire date"`
}

func (*SCU0000006I) GetServiceKey() string {
	return "CU000006"
}
