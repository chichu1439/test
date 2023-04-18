package models

import (
	"encoding/json"
)

type FP100014I struct {
	Keyword     map[string]interface{} `json:"keyword"`
	WorkSpaceId string                 `json:"workSpaceId"`
	SceneName   string                 `json:"sceneName"`
	Sort        string                 `json:"sort"`
	Order       string                 `json:"order"`
	PageSize    int64                  `json:"pageSize"`
	Page        int64                  `json:"pageNum"`
}

type FP100014O struct {
	Data  []ApiDetail `json:"data"`
	Total int64       `json:"total"`
}

func (*FP100014I) GetServiceKey() string {
	return "FP100014"
}

type ApiDetail struct {
	Id                      int64  `json:"id"`
	ApiId                   string `json:"apiID"`
	Name                    string `json:"name"`
	Category                string `json:"category"`
	Version                 string `json:"version"`
	Type                    string `json:"type"`
	Protocol                string `json:"protocol"`
	Format                  string `json:"format"`
	Endpoint                string `json:"endpoint"`
	Path                    string `json:"path"`
	ExternalPath            string `json:"externalPath"`
	Method                  string `json:"method"`
	TimeoutS                string `json:"timeoutS"`
	OnlineTime              string `json:"onlineTime"`
	MaintainTime            string `json:"maintainTime"`
	StatusTime              string `json:"statusTime"`
	PathRegExpressionEnable bool   `json:"pathRegExpressionEnable"`
	EnableAuthentication    bool   `json:"enableAuthentication"`
	EnableAccessControl     bool   `json:"enableAccessControl"`
	EnableTrafficControl    bool   `json:"enableTrafficControl"`
	Description             string `json:"description"`
	Priority                string `json:"priority"`
	EventId                 string `json:"eventId"`
	EventPkId               int    `json:"eventPkId"`
	Status                  string `json:"status"`
	ScopeId                 string `json:"scopeId"`
	CreatorId               string `json:"creatorId"`
	CreateTime              string `json:"createTime"`
	UpdatorId               string `json:"updatorId"`
	UpdateTime              string `json:"updateTime"`
}

type QuerySpec struct {
	CiType     string      `json:"ciType"`
	OrderBy    []string    `json:"orderBy,omitempty"`
	AdvFilters []AddFilter `json:"advFilters"`
	IsPaging   bool        `json:"isPaging,omitempty"`
	StartIndex int64       `json:"startIndex"`
	PageSize   int64       `json:"pageSize"`
}

type AddFilter struct {
	Name     string      `json:"name"`
	Values   interface{} `json:"values"`
	Operator string      `json:"operator"`
}

type MuResponse struct {
	ErrorCode string   `json:"errorCode"`
	ErrorMsg  string   `json:"errorMsg"`
	Response  PageData `json:"response"`
}
type PageDatas struct {
	Total int64           `json:"Total"`
	Data  json.RawMessage `json:"Data"`
}
