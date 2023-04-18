package models

type FP9000044I struct {
	MsgId   string `json:"msgId" validate:"required"`
	TransId string `json:"transId" validate:"required"`
}

type LogSearchReq struct {
	// Type      string   `json:"type"` //日志类型 audit|platform|app|passOss|redisAudit|redisQuery|dbSlow|sys|haProxy
	StartTime int64    `json:"startTime"`
	EndTime   int64    `json:"endTime"`
	Fields    []string `json:"fields,omitempty"`
	//in: body
	Filter  map[string]string `json:"filter"`
	Filters [][]string        `json:"filters"`
	Sort    string            `json:"sort,omitempty"`
	Asc     bool              `json:"asc"`
	Page    int               `json:"page,omitempty"`
	Size    int               `json:"size,omitempty"`
}

type Pagination struct {
	Page  int         `json:"page"`
	Size  int         `json:"size"`
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}

func (*FP9000044I) GetServiceKey() string {
	return "FP900044"
}
