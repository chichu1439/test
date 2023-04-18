package models

type SV000031I struct {
	AgreementId   string `json:"agreementId" description:"Agreement Id" validate:"required,max=32"`
	AgreementType string `json:"agreementType" description:"Agreement type" validate:"required"`
	PageNo        int    `json:"pageNo" description:"Page no" validate:"required,min=1"`
	PageRecCount  int    `json:"pageRecCount" description:"Page rec count" validate:"required,min=1"`
}

type SV000031OData struct {
	FreezeBusinessNumber string  `json:"freezeBusinessNumber" description:"Freeze business number" validate:"max=30"`
	SerialNumber         string  `json:"serialNumber" description:"Serial number"`
	FreezeType           string  `json:"freezeType" description:"Frozen state" validate:"max=1"`
	FreezeAmount         float64 `json:"freezeAmount" validate:"omitempty,gte=0.0" description:"Frozen amount" validate:"max=18"`
	FreezeOrganizeType   string  `json:"freezeOrganizeType" description:"Freeze Organize Type" validate:"max=1"`
	EffectiveDate        string  `json:"effectiveDate" description:"Effective date" validate:"max=10"`
	FreezeOweDate        string  `json:"freezeOweDate" description:"Freeze expiry date" validate:"max=10"`
}

type SV000031O struct {
	TotalRecord int             `json:"totalRecord" description:"Total record"`
	TotalPages  int             `json:"totalPages" description:"Total pages"`
	Records     []SV000031OData `json:"records" description:"Records"`
}

func (*SV000031I) GetServiceKey() string {
	return "SV000031"
}
