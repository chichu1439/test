package models

type SV000032I struct {
	AgreementId   string `json:"agreementId" description:"Agreement Id" validate:"required,max=30"`
	AgreementType string `json:"agreementType" description:"Agreement type" validate:"required,max=5"`
	DeductionMode string `json:"deductionMode" description:"Deduction Mode" validate:"max=1"`
	PageNo        int    `json:"pageNo" description:"Page no"`
	PageRecCount  int    `json:"pageRecCount" description:"Page rec count"`
}

type SV000032OData struct {
	DeductionBusinessNumber string  `validate:"max=30" json:"deductionBusinessNumber" description:"Deduction business number" validate:"required,max=30"`
	DeductionDate           string  `json:"deductionDate" description:"Deduction date" validate:"required,max=10"`
	DeductionAmount         float64 `json:"deductionAmount" description:"Deduction amount" validate:"required,max=18"`
	DeductionDocnumber      string  `json:"deductionDocnumber" description:"Deduction docnumber"`
}

type SV000032O struct {
	TotalRecord int             `json:"totalRecord" description:"Total record"`
	TotalPages  int             `json:"totalPages" description:"Total pages"`
	Records     []SV000032OData `json:"records" description:"Records"`
}

func (*SV000032I) GetServiceKey() string {
	return "SV000032"
}
