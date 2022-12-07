package models

type SMP9CM0003I struct {
	Version string `json:"version" validate:"required"`
}
type SMP9CM0003O struct {
	ForceUpdate  bool   `json:"forceUpdate"`
	DownloadLink string `json:"downloadLink"`
}

func (i SMP9CM0003I) GetServiceKey() string {
	return "MP9CM003"
}
