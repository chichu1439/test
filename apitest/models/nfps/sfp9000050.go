//Version: v1
package models

type FP900050I struct {
	PageNum         int `json:"pageNum"`
	PageRecordCount int `json:"pageRecordCount"`
}

type FP900050O struct {
	PageNum         int            `json:"pageNum"`
	PageRecordCount int            `json:"pageRecordCount"`
	TotalCount      int            `json:"totalCount"`
	ScenarioList    []ScenarioInfo `json:"scenarioList"`
}
type ScenarioInfo struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

func (*FP900050I) GetServiceKey() string {
	return "FP900050"
}
