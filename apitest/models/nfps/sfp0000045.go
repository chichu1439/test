package models

import (
	"github.com/go-playground/validator/v10"
	"time"
)



type FP000045I struct {
	Filter map[string]interface{} `json:"filter"`
	Keyword map[string]interface{} `json:"keyword"`
	ParticipantClearingCode string `json:"participantClearingCode"`
	StartTime int64 `json:"startTime"`
	EndTime int64 `json:"endTime"`
	Sort     string `json:"sort"`
	Order    string `json:"order"`
	PageSize int `json:"pageSize" validate:"required"`
	Page int `json:"pageNum" validate:"required"`
}

type FP000045O struct {
	Data []AlertDetil `json:"data"`
	Total int64 `json:"total"`
}

type AlertDetil struct {
	Id int64 `json:"id"`
	AlertType string `json:"alertType"`
	ParticipantClearingCode string `json:"participantClearingCode"`
	AlertEntity string `json:"alertEntity"`
	AlertDetails string `json:"alertDetails"`
	Level string `json:"alertLevel"`
	TimeStart time.Time `json:"createTime"`
	AlertName string `json:"alertName"`
	Label string `json:"label"`
}

func (*FP000045I) GetServiceKey() string {
	return "FP000045"
}


func (d *FP000045I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}