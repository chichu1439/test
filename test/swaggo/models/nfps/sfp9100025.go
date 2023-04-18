package models

type FP910025I struct {
	DateTimeStart string `json:"dateTimeStart"`
	DateTimeEnd   string `json:"dateTimeEnd"`
}
type FP910025O struct {
	TotalCount int            `json:"totalCount"`
	RecordsOut []FP910025Item `json:"recordsOut"`
	RecordsIn  []FP910025Item `json:"recordsIn"`
}
type FP910025Item struct {
	ParticipantCode string `json:"participantCode"`
	Count           int    `json:"count"`
	Currency        string `json:"currency"`
	TransType       string `json:"transType"`
}

func (*FP910025I) GetServiceKey() string {
	return "FP910025"
}
