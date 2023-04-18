package models

type FP900074I struct {
	Id int64 `json:"Id" validate:"required"`
}

type FP900074O struct {
	Status string `json:"status"`
}

func (*FP900074I) GetServiceKey() string {
	return "FP900074"
}
