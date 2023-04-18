package models

type CM000050I struct {
	TaskId     string `json:"taskId" description:"Task id" validate:"required"`
	TaskStatus string `json:"taskStatus" description:"Task status" `
	//TaskUserId   string `json:"taskUserId" `
	//TaskUserName string `json:"taskUserName" `
	FinalModifyOrgz string `json:"finalModifyOrgz" description:"Final modify organization" `
	FinalModifyUser string `json:"finalModifyUser" description:"Final modify user"`
}

type CM000050O struct {
}

func (*CM000050I) GetServiceKey() string {
	return "CM000050"
}
