package models

type SCU0000010I struct {
	CustomerId     string `json:"customerId" description:"Customer id"`
	CustomerStatus string `json:"customerStatus" description:"customer Status:0:Normal 4:Pending 9:Cancel "`
}

type SCU0000010O struct {
}

func (*SCU0000010I) GetServiceKey() string {
	return "CU000010"
}
