//Version: v1.0.0
package models

type IC000038I struct {
	LoanReceiptNum string `json:"loanReceiptNum" description:"Loan receipt number" validate:"required"`
	TaskId         string `json:"taskId" description:"Task id" validate:"required"`
}

type IC000038O struct {
	LoanReceiptNum string `json:"loanReceiptNum" description:"Loan receipt number"`
	LoanPeriod     int    `json:"loanPeriod" description:"Loan period"`
	TaskStatus     string `json:"taskStatus" description:"Task status"`
	GracePeriod    int    `json:"gracePeriod" description:"Grace period"`
	TaskStartDate  string `json:"taskStartDate" description:"Task start date"`
	TaskEndDate    string `json:"taskEndDate" description:"Task end date"`
	ContactTimes   int    `json:"contactTimes" description:"Contact times"`
	ResponseTimes  int    `json:"responseTimes" description:"Response times"`
	TaskUserId     string `json:"taskUserId" description:"Task user id"`
	TaskUserName   string `json:"taskUserName" description:"Task user name"`
}

func (i *IC000038I) GetServiceKey() string {
	return "IC000038"
}
