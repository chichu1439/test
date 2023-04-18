package models

type PF000027I struct {
	FeeTypeId       string `json:"feeTypeId"`
	FeeTypeName     string `json:"feeTypeName"`
	PageNum         int    `json:"pageNum"`
	PageRecordCount int    `json:"pageRecordCount"`
}

type PF000027O struct {
	PageNum         int                `json:"pageNum"`
	PageRecordCount int                `json:"pageRecordCount"`
	TotalCount      int                `json:"totalCount"`
	Records         []PF000027ORecords `json:"records"`
}
type PF000027ORecords struct {
	FeeTypeId       string `json:"feeTypeId"`
	FeeTypeName     string `json:"feeTypeName"`
	AccountingCode  string `json:"accountingCode"`
	IsAmortization  string `json:"isAmortization"`
	Description     string `json:"description"`
	CreateDate      string `json:"createDate"`
	FinalModifyOrgz string `json:"finalModifyOrgz"`
	FinalModifyUser string `json:"finalModifyUser"`
	FinalModifyDate string `json:"finalModifyDate"`
	FinalModifyTime string `json:"finalModifyTime"`
}

func (*PF000027I) GetServiceKey() string {
	return "PF000027"
}
