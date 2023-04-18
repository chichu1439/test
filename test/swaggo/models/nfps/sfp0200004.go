package models

type FP020004I struct {
	AccountKeyNum    string `json:"accountKeyNum"`
	PartClearingCode string `json:"partClearingCode"`
	PageNo           int    `json:"pageNo"`
	PageSize         int    `json:"pageSize"`
}
type FP020004O struct {
	TotalCount     int              `json:"totalCount"`
	AddressingList []AddressingInfo `json:"addressingList"`
}

func (*FP020004I) GetServiceKey() string {
	return "FP020004"
}
