package models

type SPD0000058I struct {
	ProductId string `json:"productId" description:"Product id" validate:"required"`
	FeeType   string `json:"feeType" validate:"required" description:"fee Type:01-Default Application Fees;02-Default Processing/Origination Fees;03-Default Overdue Fees;04-Default Anticipated Repayment Fees;05-Default Extension Fees;06-Default Repayment Plan Adjustment Fees;07-Default Tax Fees"`
}

type SPD0000058O struct {
}

func (*SPD0000058I) GetServiceKey() string {
	return "spd0000058"
}
