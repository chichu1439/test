package models

type FP000035I struct {
	FileName string `json:"fileName" validate:"required"`
}
type FP000035O struct {
	UrlStr string `json:"urlStr"`
}

func (*FP000035I) GetServiceKey() string {
	return "FP000035"
}
