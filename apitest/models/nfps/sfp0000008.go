package models

//inquire Participant relation detail

type FP000008I struct {
	DirectPartClearingCode   string `json:"directPartClearingCode" description:"Direct part clearing code" validate:"required"`
	IndirectPartClearingCode string `json:"indirectPartClearingCode" description:"Indirect part clearing code" validate:"required"`
	Currency                 string `json:"currency" description:"Currency" validate:"required"`
}

type FP000008O struct {
	DirectPartClearingCode   string `json:"directPartClearingCode" description:"Direct part clearing code"`
	PartTypeDp               string `json:"partTypeDp" description:"Part type dp"`
	DirectPartName           string `json:"directPartName" description:"Direct part name"`
	IndirectPartClearingCode string `json:"indirectPartClearingCode" description:"Indirect part clearing code"`
	PartTypeIp               string `json:"partTypeIp" description:"Part type ip"`
	IndirectPartName         string `json:"indirectPartName" description:"Indirect part name"`
	Currency                 string `json:"currency" description:"Currency"`
	EffectiveDate            string `json:"effectiveDate" description:"Effective date"`
	Remark                   string `json:"remark" description:"Remark"`
	Status                   string `json:"status" description:"Status"`
}

func (*FP000008I) GetServiceKey() string {
	return "FP000008"
}
