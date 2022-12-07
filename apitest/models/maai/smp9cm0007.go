package models

// swagger:parameters SMP9CM0007I
type SMP9CM0007I struct {
	System       string `json:"system" validate:"max=45"`
	ApiName      string `json:"apiName" validate:"max=45"`
	ResponseCode string `json:"responseCode" validate:"max=45"`
	Language     string `json:"language"`
	ErrorCode    string `json:"errorCode" validate:"max=45"`
	ErrorMsg     string `json:"errorMsg" validate:"max=45"`
	ErrorMsgTh   string `json:"errorMsgTh" validate:"max=255"`
}

// swagger:parameters SMP9CM0007O
type SMP9CM0007O struct {
	ErrorMsgs []ErrorMsg `json:"errorMsgs"`
}

func (*SMP9CM0007I) GetServiceKey() string {
	return "MP9CM007"
}

type ErrorMsg struct {
	Id           string `json:"id" validate:"required,max=45"`
	System       string `json:"system" validate:"max=45"`
	ApiName      string `json:"apiName" validate:"max=45"`
	Language     string `json:"language"`
	ErrorCode    string `json:"errorCode" validate:"max=45"`
	ErrorMsg     string `json:"errorMsg" validate:"max=45"`
	ErrorMsgTh   string `json:"errorMsgTh" validate:"max=255"`
	ResponseCode string `json:"responseCode"`
}
