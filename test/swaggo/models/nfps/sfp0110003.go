package models

import "github.com/go-playground/validator/v10"

type SF011003I struct {
	FileName string `json:"fileName" validate:"required"`
	FileType string `json:"fileType" validate:"required"`
}

type SF011003O struct {
	Records []FP110030ResponseRecord `json:"records"`
}

type FP110030ResponseRecord struct {
	MsgId         string  `json:"msgId"`
	TransId       string  `json:"transId"`
	ClrSysRef     string  `json:"clrSysRef"`
	TransDate     string  `json:"transDate"`
	TransType     string  `json:"transType"`
	TransStatus   string  `json:"transStatus"`
	Currency      string  `json:"currency"`
	Amount        float64 `json:"amount"`
	InstructionId string  `json:"instructionId"`
	EndToEndId    string  `json:"endToEndId"`
	MsgDirection  string  `json:"msgDirection"`
}

func (d *SF011003I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}
