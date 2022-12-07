package models

type GlsCreateGlsDataArgs struct {
	Option    int            `json:"option" description:"Option" validate:"omitempty"`
	Dimension *DataDimension `json:"dimension" description:"Dimension" validate:"required"`
	Elements  []GlsData      `json:"elements" description:"Elements" validate:"required"`
}

type DataDimension struct {
	GlsDimension
	Topic  GlsTopic `json:"topic" description:"Topic" validate:"required"`
	SuType string   `json:"suType" description:"Su type" validate:"required"`
}

type GlsDimension struct {
	Tenant      string `json:"tenant" description:"Tenant" validate:"omitempty"`
	Workspace   string `json:"workspace" description:"Workspace" validate:"omitempty"`
	Environment string `json:"environment" description:"Environment" validate:"omitempty"`
}

type GlsTopic struct {
	TopicType string `json:"topicType" description:"Topic type" validate:"omitempty"`
	TopicId   string `json:"topicId" description:"Topic id" validate:"omitempty"`
}

type GlsData struct {
	GlsType  string `json:"glsType" description:"Gls type" validate:"required"`
	GlsClass string `json:"glsClass" description:"Gls class" validate:"omitempty"`
	DataId   string `json:"dataId" description:"Data id" validate:"required"`
}

type GlsCreateResponse struct {
	ErrorCode string        `json:"errorCode" description:"Error code"`
	ErrorMsg  string        `json:"errorMsg" description:"Error msg"`
	Response  *GlsPrimarySu `json:"response" description:"Response"`
}

//// Create Response

type GlsPrimarySu struct {
	SuType string `json:"suType" description:"Su type" validate:"required"`
	SuId   string `json:"suId" description:"Su id" validate:"required"`
}

func (*GlsCreateGlsDataArgs) GetServiceKey() string {
	return "GlsCreate"
}
