//Version: 1
package models

type SAC0000005I struct {
	FreezeBusinessNumber string  `json:"freezeBusinessNumber" description:"Freeze business number" validate:"required"`
	Account              string  `json:"account" description:"Account" validate:"required"`
	Amount               float64 `json:"amount" description:"Amount" validate:"required"`
	OverAmtFreezeFlag    string  `json:"overAmtFreezeFlag" description:"Over amt freeze flag"` //超额冻结标识，Y-允许超额冻结，N-不允许超额冻结
	TransDate            string  `json:"transDate" description:"Trans date" validate:"required"`
}

type SAC0000005O struct {
	AmountFreezing  float64 `json:"amountFreezing" description:"Amount freezing"`
	AmountCurrent   float64 `json:"amountCurrent" description:"Amount current"`
	AmountLast      float64 `json:"amountLast" description:"Amount last"`
	AmountAvailable float64 `json:"amountAvailable" description:"Amount available"`
	AccountStatus   string  `json:"accountStatus" description:"Account status"`
}

func (*SAC0000005I) GetServiceKey() string {
	return "sac0000005"
}
