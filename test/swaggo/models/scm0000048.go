package models

type CM000048I struct {
	TaskId         string `json:"taskId" description:"Task id" `
	LoanReceiptNum string `json:"loanReceiptNum" description:"Loan receipt number" `
	TaskStatus     string `json:"taskStatus" description:"Task status" `
	//RepaymentStatus string `json:"repaymentStatus" `
	PageNo       int `json:"pageNo" description:"Page no" validate:"required"`
	PageRecCount int `json:"pageRecCount" description:"Page rec count" validate:"required"`
}

type CM000048O struct {
	PageNo       int        `json:"pageNo" description:"Page no"`
	PageTotCount int        `json:"pageTotCount" description:"Page tot count"`
	ListTask     []ListTask `json:"listTask" description:"List task"`
}

type ListTask2 struct {
	TaskId               string `json:"taskId" description:"Task id"`
	LoanReceiptNum       string `json:"loanReceiptNum" description:"Loan receipt number"`
	LoanPeriod           int    `json:"loanPeriod" description:"Loan period" `
	TaskStatus           string `json:"taskStatus" description:"Task status"`
	TaskStartDate        string `json:"taskStartDate" description:"Task start date"`
	TaskEndDate          string `json:"taskEndDate" description:"Task end date"`
	AddScenario          string `json:"addScenario" description:"Add scenario"`
	CollectionAmountType string `json:"collectionAmountType" description:"Collection amount type"`
	TaskUserId           string `json:"taskUserId" description:"Task user id"`
	TaskUserName         string `json:"taskUserName" description:"Task user name"`
	CurrentDueDate       string `json:"currentDueDate" description:"Current due date"`
}

func (*CM000048I) GetServiceKey() string {
	return "CM000048"
}
