package controller

// Controller example
type Controller struct {
}

// NewController example
func NewController() *Controller {
	return &Controller{}
}

type Response struct {
	Code     string      `json:"code" example:"0"`
	Message  string      `json:"message" example:""`
	Response interface{} `json:"response"`
}

type Error struct {
	ErrorCode  string `json:"errorCode" example:"160005"`
	ErrorMsg   string `json:"errorMsg" example:"Please retry later, request backend failed"`
	ErrorStack string `json:"errorStack" example:"<no value>"`
}
