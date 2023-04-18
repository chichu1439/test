package models

type CM000045I struct {
	TaskId               string `json:"taskId" description:"Task id" validate:"required"`
	LoanReceiptNum       string `json:"loanReceiptNum" description:"Loan receipt number" validate:"required"`
	LoanPeriod           int    `json:"loanPeriod" description:"Loan period"`
	TaskStartDate        string `json:"taskStartDate" description:"Task start date" validate:"required"`
	CurrentDueDate       string `json:"currentDueDate" description:"Current due date" validate:"required"`
	GracePeriod          int    `json:"gracePeriod" description:"Grace period"`
	TaskUserId           string `json:"taskUserId" description:"Task user id"`
	TaskUserName         string `json:"taskUserName" description:"Task user name"`
	CollectionAmountType string `json:"collectionAmountType" description:"Collection amount type"`
	AddScenario          string `json:"addScenario" description:"Add scenario" validate:"required"`
}

type CM000045O struct {
}

func (*CM000045I) GetServiceKey() string {
	return "CM000045"
}
