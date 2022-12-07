//Version: v1.0.0
package models

type IC000036I struct {
	LoanReceiptNum string `json:"loanReceiptNum" description:"Loan receipt number" validate:"required"`
	//LoanPeriod           Int    `json:"loanPeriod" validate:"required"`   `json:"int"`
	TaskStartDate        string `json:"taskStartDate" description:"Task start date" validate:"required"`
	TaskUserId           string `json:"taskUserId" description:"Task user id"`
	TaskUserName         string `json:"taskUserName" description:"Task user name"`
	CollectionAmountType string `json:"collectionAmountType" description:"Collection amount type"`
}

type IC000036O struct {
	TaskId string `json:"taskId" description:"Task id"`
}
