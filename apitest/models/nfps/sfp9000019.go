package models

//update Participant balance info
type FP900019I struct {
	PartClearingCode string   `json:"partClearingCode" validation:"required"`
	Currency         string   `json:"currency" validation:"required"`
	AlertBalance     *float64 `json:"alertBalance"`
	AlertMinuteTime  *int     `json:"alertMinuteTime"`
	Status           string   `json:"status"`
}
type FP900019O struct {
}

func (*FP900019I) GetServiceKey() string {
	return "FP900019"
}
