package models

//Inquire Participant list
type FP900021I struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}
type FP900021O struct {
	TotalRecCount int             `json:"totalRecCount"`
	Records       []FP900021OItem `json:"records"`
}
type FP900021OItem struct {
	PartClearingCode string `json:"partClearingCode"`
	PartName         string `json:"partName"`
	PartType         string `json:"partType"`
	Status           string `json:"status"`
	Currency         string `json:"currency"`
	PartIpAddress    string `json:"partIpAddress"`
}

func (*FP900021I) GetServiceKey() string {
	return "FP900021"
}
