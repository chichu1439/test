//Version: v1.0.0
package models

type IC000039I struct {
	LoanReceiptNum string `json:"loanReceiptNum" description:"Loan receipt number" validate:"required"`
	TaskId         string `json:"taskId" description:"Task id" validate:"required"`
	TaskStatus     string `json:"taskStatus" description:"Task status" `
	//TaskUserId   string `json:"taskUserId" `
	//TaskUserName string `json:"taskUserName" `
}

type IC000039O struct {
}

func (*IC000039I) GetServiceKey() string {
	return "IC000039"
}
