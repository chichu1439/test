package models

type GlsRemoveRequest struct {
	GlsType  string `json:"glsType" description:"Gls type" validate:"required"`
	GlsClass string `json:"glsClass" description:"Gls class" validate:"omitempty"`
	DataId   string `json:"dataId" description:"Data id" validate:"required"`
}

type GlsRemoveResponse struct {
	ErrorCode string      `json:"errorCode" description:"Error code"`
	ErrorMsg  string      `json:"errorMsg" description:"Error msg"`
	Data      GlsResponse `json:"data" description:"Data"`
}

func (*GlsRemoveRequest) GetServiceKey() string {
	return "GlsRemove"
}
