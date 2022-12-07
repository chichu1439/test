package models

type GlsDcnListRequest struct {
	Dimension GlsDimension `json:"dimension" description:"Dimension"`
	SuTypes   []SuTypes    `json:"suTypes" description:"Su types"`
}

type GlsDcnListResponse struct {
	ErrorCode uint                `json:"errorCode" description:"Error code"`
	ErrorMsg  string              `json:"errorMsg" description:"Error msg"`
	Data      map[string][]string `json:"data" description:"Data"`
}

type SuTypes struct {
	Topic  GlsTopic `json:"topic" description:"Topic"`
	SuType string   `json:"suType" description:"Su type"`
}

func (*GlsDcnListRequest) GetServiceKey() string {
	return "GlsDcnList"
}
