package models

//Create Part
type FP900102I struct {
	BankIdArray []string `json:"bankIdArray"`
	AllBank     bool     `json:"allBank"`
	BeginTime   string   `json:"beginTime"`
	EndTime     string   `json:"endTime"`
}

type SampleData struct {
	MatchCount   int64  `json:"matchCount"`
	UnMatchCount int64  `json:"unMatchCount"`
	Tag          string `json:"tag"`
	DateTime     string `json:"dateTime"`
}

type FP900102O struct {
	SampleDataArray []SampleData `json:"sampleDataArray"`
}

func (*FP900102I) GetServiceKey() string {
	return "FP900102"
}
