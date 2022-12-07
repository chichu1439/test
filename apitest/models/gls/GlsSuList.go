package models

type GlsSuListRequest struct {
	Dimension GlsDimension `json:"dimension" description:"Dimension"`
	SuTypes   []SuTypes    `json:"suTypes" description:"Su types"`
}

type GlsSuListResponse struct {
	Dimension GlsDimension `json:"dimension" description:"Dimension"`
	SuTypes   []SuTypeInfo `json:"suTypes" description:"Su types"`
}

type SuTypeInfo struct {
	Topic  GlsTopic `json:"topic" description:"Topic"`
	SuType string   `json:"suType" description:"Su type"`
	SuList []string `json:"suList" description:"Su list"`
}

func (*GlsSuListRequest) GetServiceKey() string {
	return "GlsSuList"
}
