package models

import "github.com/go-playground/validator/v10"

type NewFiles struct {
	DateTime  string   `json:"dateTime"`
	FileNames []string `json:"fileNames"`
}

type Results struct {
	Status bool `json:"status"`
}

type FP000044I struct {
	ParticipantClearingCode string `json:"participantClearingCode"`
	Details                 bool   `json:"details"`
	Page                    int    `json:"pageNum"`
	PageSize                int    `json:"pageSize" validate:"required"`
}

type FP000044O struct {
	Data []ParticipantList `json:"data"`
	Rows int64             `json:"rows"`
}

type ParticipantList struct {
	ParticipantClearingCode string `json:"participantClearingCode" xorm:"'participant_clearing_code'"`
	ContactName             string `json:"contactName" xorm:"'contact_name'"`
	ContactMethod           string `json:"contactMethod" xorm:"'contact_method'"`
	ContactInformation      string `json:"contactInformation" xorm:"'contact_information'"`
	Priority                string `json:"priority" xorm:"'priority'"`
	Necessary               string `json:"necessary" xorm:"'necessary'"`
	Remark                  string `json:"remark" xorm:"'remark'"`
	Id                      int64  `json:"id" xorm:"'id'"`
	Label                   string `json:"label"`
	Count                   int    `json:"count"`
}

func (*FP000044I) GetServiceKey() string {
	return "FP000044"
}

func (d *FP000044I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}
