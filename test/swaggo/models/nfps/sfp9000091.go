package models

type FP900091I struct {
	FileStatus string `json:"fileStatus"`
	FileName   string `json:"fileName"`
}
type FP900091O struct {
	Status string `json:"status"`
}

func (*FP900091I) GetServiceKey() string {
	return "FP900091"
}
