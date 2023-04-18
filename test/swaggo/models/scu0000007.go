package models

type SCU0000007I struct {
	CustomerId   string `json:"customerId" description:"Customer id"`
	ContractType string `json:"contractType" description:"contract Type:01- Home phone 02- Office phone 04- EMAIL 09- Mobile 99- other"`
}

type SCU0000007O struct {
	ContractCount int `json:"contractCount" description:"Contract count"`
}

func (*SCU0000007I) GetServiceKey() string {
	return "CU000007"
}
