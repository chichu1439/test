package models

type SCU0000011I struct {
}

type SCU0000011O struct {
	TotalCustomerNum          int `json:"totalCustomerNum" description:"Total customer number"`
	TotalCustomerCurrentMonth int `json:"totalCustomerCurrentMonth" description:"Total customer current month"`
	TotalCustomerBlack        int `json:"totalCustomerBlack" description:"Total customer black"`
}

func (*SCU0000011I) GetServiceKey() string {
	return "CU000011"
}
