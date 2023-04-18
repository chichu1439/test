/**
 * @Author: fangwen
 * @Description:
 * @File: sfp1000007
 * @Version: 1.0.0
 * @Date: 2022/3/18 1:45 AM
 */

package models

import (
	"encoding/json"
)

type FP900076I struct {
	Filter                  map[string]interface{} `json:"filter"`
	Keyword                 map[string]interface{} `json:"keyword"`
	ParticipantClearingCode []string               `json:"participantClearingCode"`
	Sort                    string                 `json:"sort"`
	Order                   string                 `json:"order"`
	StartTime               int64                  `json:"startTime"`
	EndTime                 int64                  `json:"endTime"`
	PageSize                int                    `json:"pageSize" validate:"required"`
	Page                    int                    `json:"page" validate:"required"`
}

type FP900076O struct {
	Data  []AlertDetil `json:"data"`
	Total int64        `json:"total"`
}

func (*FP900076I) GetServiceKey() string {
	return "FP900076"
}

type QSpec struct {
	CiType     string                 `json:"ciType"`
	Fields     []string               `json:"fields,omitempty"`
	TarCiType  []string               `json:"tarCiType,omitempty"`
	Filter     map[string]interface{} `json:"filter,omitempty"`
	OrderBy    []string               `json:"orderBy,omitempty"`
	AdvFilters []AddFilters           `json:"advFilters"`
	//Org string `json:"org"`
	IsPaging    bool          `json:"isPaging,omitempty"`
	StartIndex  int64         `json:"startIndex"`
	PageSize    int64         `json:"pageSize"`
	EnvIDs      []interface{} `json:"envIDs,omitempty"`
	WsIDs       []interface{} `json:"wsIDs,omitempty"`
	GroutBy     []string      `json:"groupBy"`
	RefDataMode string        `json:"refDataMode"`
}

type AggregateFilter struct {
	TabName   string                 `json:"tabName"`
	Field     string                 `json:"field"`
	Filter    map[string]interface{} `json:"filter"`
	TimeSlice [2]int64               `json:"timeSlice"`
}

type AddFilters struct {
	Name     string      `json:"name"`
	Values   interface{} `json:"values"`
	Operator string      `json:"operator"`
}

type MuResponseSpec struct {
	ErrorCode string   `json:"errorCode"`
	ErrorMsg  string   `json:"errorMsg"`
	Response  PageData `json:"response"`
}
type PageData struct {
	Total int64           `json:"Total"`
	Data  json.RawMessage `json:"Data"`
}
