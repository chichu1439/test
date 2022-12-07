package models

type ResponseTemplate struct {
	ErrorCode string `json:"errorCode"`
	ErrorMsg string `json:"errorMsg"`
	Response interface{} `json:"response"`
}