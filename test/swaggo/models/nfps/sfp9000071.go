package models

type FP900071I struct {
	Id int64 `json:"id" validate:"required"`
}

type FP900071O struct {
	Status string `json:"status"`
}

func (*FP900071I) GetServiceKey() string {
	return "FP900071"
}
