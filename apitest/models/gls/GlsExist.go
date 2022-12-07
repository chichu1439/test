package models

type GlsExistRequest struct {
	Dimension GlsDimension `json:"dimension" description:"Dimension"`
	Element   GlsData      `json:"element" description:"Element"`
}

type GlsExistResponse struct {
	ErrorCode uint   `json:"errorCode" description:"Error code"`
	ErrorMsg  string `json:"errorMsg" description:"Error msg"`
	Data      bool   `json:"data" description:"Data"`
}

func (*GlsExistRequest) GetServiceKey() string {
	return "GlsExist"
}
