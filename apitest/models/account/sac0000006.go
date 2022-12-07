package models

type AC000006I struct {
	FreezeBusinessNumber string  `json:"freezeBusinessNumber" description:"Freeze business number" validate:"required"`
	Account              string  `json:"account" description:"Account" validate:"required"`
	Amount               float64 `json:"amount" description:"Amount" validate:"required"`
	TransDate            string  `json:"transDate" description:"Trans date" validate:"required"`
}

type AC000006O struct {
	AmountFreezing  float64 `json:"amountFreezing" description:"Amount freezing"`
	AmountCurrent   float64 `json:"amountCurrent" description:"Amount current"`
	AmountLast      float64 `json:"amountLast" description:"Amount last"`
	AmountAvailable float64 `json:"amountAvailable" description:"Amount available"`
	AccountStatus   string  `json:"accountStatus" description:"Account status"`
}

func (*AC000006I) GetServiceKey() string {
	return "sac0000006"
}
