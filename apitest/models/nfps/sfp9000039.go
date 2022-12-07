package models

type FP900039I struct {
	ProductId string `json:"productId"`
	Currency  string `json:"currency"`
	PageNo    int    `json:"pageNo"`
	PageSize  int    `json:"pageSize"`
}
type FP900039O struct {
	TotalCount int             `json:"totalCount"`
	Records    []FP000039IItem `json:"records"`
}

func (*FP900039I) GetServiceKey() string {
	return "FP900039"
}
