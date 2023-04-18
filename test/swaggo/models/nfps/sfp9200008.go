package models

type DownRequstParam struct {
	FileName         string `json:"fileName"`
	PartClearingCode string `json:"partClearingCode"`
}

type ResetSvcData struct {
	Data string `json:"data"`
}

type ResetSvcConversionData struct {
	Data []byte `json:"data"`
}
