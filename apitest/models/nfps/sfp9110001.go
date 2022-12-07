package models

type FP911001I struct {
	FileTrans     []FileTrans `json:"file_trans"`
	PartClearCode string
}

type FP911001O struct {
	TotalRecord int                      `json:"totalRecord"`
	TotalPage   int                      `json:"totalPage"`
	Records     []FP110030ResponseRecord `json:"records"`
}

func (*FP911001I) GetServiceKey() string {
	return "FP911001"
}
