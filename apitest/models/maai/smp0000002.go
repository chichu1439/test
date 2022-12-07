package models

type SMP0000002I struct {
	Version string `json:"version" description:"version" validate:"required,max=10"`
}
type SMP0000002O struct {
	ForceUpdate bool `json:"forceUpdate" description:"force update flag"`
}

func (i SMP0000002I) GetServiceKey() string {
	return "MP000002"
}
