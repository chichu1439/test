package models

import (
	"fmt"
)

type CommonResponse struct {
	ReturnCode    string      `json:"returnCode" description:"Return code"`
	ReturnMessage string      `json:"returnMessage" description:"Return message"`
	Data          interface{} `json:"data"`
}

const (
	successCode = "0"
	successMsg  = "success"
)

type Float64 float64

// MarshalJSON to keep leading zero of the number
func (f Float64) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%f", f)), nil
}
