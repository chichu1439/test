//Version: 1
package models

type SAC0000004I struct {
	AccountArray []SAC0000004IAccountInfo `json:"accountArray" description:"Account array"`
}

type SAC0000004IAccountInfo struct {
	CustomerId    string  `json:"customerId" description:"Customer id" validate:"required"`
	ContractId    string  `json:"contractId" description:"Contract id" validate:"required"`
	Account       string  `json:"account" description:"Account" validate:"required"`
	Amount        float64 `json:"amount" description:"Amount" validate:"required"`
	TransDate     string  `json:"transDate" description:"Trans date" validate:"required"`
	TransSequence string  `json:"transSequence" description:"Trans sequence" validate:"required"`
	FlowSequence  int     `json:"flowSequence" description:"Flow sequence" validate:"required"`
	DeductionFlag string  `json:"deductionFlag" description:"Deduction flag"`
}

type SAC0000004O struct {
	ResultArray []SAC0000004OResultInfo `json:"resultArray" description:"Result array"`
}

type SAC0000004OResultInfo struct {
	AmountFreezing  float64 `json:"amountFreezing" description:"Amount freezing"`
	AmountCurrent   float64 `json:"amountCurrent" description:"Amount current"`
	AmountLast      float64 `json:"amountLast" description:"Amount last"`
	AmountAvailable float64 `json:"amountAvailable" description:"Amount available"`
	AccountStatus   string  `json:"accountStatus" description:"Account status"`
}

func (*SAC0000004I) GetServiceKey() string {
	return "sac0000004"
}
