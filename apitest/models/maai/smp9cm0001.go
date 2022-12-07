package models

type SMP9CM0001I struct {
	DocGroup      []string `json:"docGroup"`
	LatestVersion int      `json:"latestVersion"`
}
type SMP9CM0001O struct {
	Documents []SMP9CM0001OList `json:"documents"`
}
type SMP9CM0001OList struct {
	DocGroup         string `json:"docGroup"`
	DocId            string `json:"docId"`
	Remark           string `json:"remark"`
	DocVersion       int    `json:"docVersion"`
	DocContentTextId string `json:"docContentTextId"`
}

func (i SMP9CM0001I) GetServiceKey() string {
	return "MP9CM001"
}
