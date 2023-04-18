package models

type RequestSfp0200009 struct {
	Size             int    `json:"size"`
	Page             int    `json:"page"`
	PartClearingCode string `json:"partClearingCode"`
	FileName         string `json:"fileName"`
}

type ResponseSfp0200009 struct {
	SumTotal int64      `json:"sumTotal"`
	Data     []RowsData `json:"data"`
}

type RowsData struct {
	CreatedDate int64  `json:"createdDate"`
	FileName    string `json:"fileName"`
}
