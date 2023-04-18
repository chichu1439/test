package models

type FP900041I struct {
	ProductId string `json:"productId"`
	Currency  string `json:"currency"`
	ServiceId string `json:"serviceId"`
	FeatureId string `json:"featureId"`
	Status    string `json:"status"`
}
type FP900041O struct {
	TotalCount int             `json:"totalCount"`
	Records    []FP000041IItem `json:"records"`
}

func (*FP900041I) GetServiceKey() string {
	return "FP900041"
}
