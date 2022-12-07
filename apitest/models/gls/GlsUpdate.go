package models

type GlsUpdatePrimaryParamter struct {
	Primary   GlsData            `json:"primary" description:"Primary" validate:"required"`
	Dimension DataDimension      `json:"dimension" description:"Dimension" validate:"required"`
	Elements  []GlsUpdateElement `json:"elements" description:"Elements" validate:"required"`
}

type GlsUpdateElement struct {
	GlsData
	SourceClass string `json:"sourceClass" description:"Source class" validate:"omitempty"`
	SourceId    string `json:"sourceId" description:"Source id" validate:"omitempty"`
	State       int    `json:"state" description:"State" validate:"omitempty"`
}

// Update & Remove Response

type GlsResponse struct {
	Result bool `json:"result" description:"Result" validate:"required"`
}

func (*GlsUpdatePrimaryParamter) GetServiceKey() string {
	return "GlsUpdate"
}
