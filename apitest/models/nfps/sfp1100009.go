package models

type FP110009I struct {
	TransId   string `json:"transId" validate:"required"`
	MsgId     string `json:"msgId" validate:"required"`
	ClrSysRef string `json:"clrSysRef" validate:"required"`
	MmbId     string `json:"mmbId" validate:"required"`
}

type FP110009O struct {
	EndToEndId        string  `json:"endToEndId"`
	TransDatetime     string  `json:"transDatetime"`
	InterbankStlmtAmt float64 `json:"interbankStlmtAmt"`
	InterbankStlmtCcy string  `json:"interbankStlmtCcy"`
	TransResultCode   string  `json:"transResultCode"`
	TransResultReason string  `json:"transResultReason"`
}

func (*FP110009I) GetServiceKey() string {
	return "FP110009"
}
