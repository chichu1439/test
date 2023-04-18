package models

// swagger:parameters inquireParticipantListRequest
// in: body

type FP000102I struct {
	BankIdArray []string `json:"bankIdArray"`
	AllBank     bool     `json:"allBank"`
	BeginTime   string   `json:"beginTime"`
	EndTime     string   `json:"endTime"`
}

// swagger:response inquireParticipantListResponse

type FP000102O struct {
	SampleDataArray []SampleData `json:"sampleDataArray"`
}

//type SampleData struct {
//	MatchCount   int64  `json:"matchCount"`
//	UnMatchCount int64  `json:"unMatchCount"`
//	Tag          string `json:"tag"`
//	DateTime     string `json:"dateTime"`
//}

func (*FP000102I) GetServiceKey() string {
	return "FP000102"
}
