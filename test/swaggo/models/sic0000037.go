//Version: v1.0.0
package models

import (
	//"git.sirius.io/banking/common/constant"
	"github.com/shopspring/decimal"
	//"sort"
	//"time"
)

type IC000037I struct {
	TaskId         string `json:"taskId" description:"Task id"`
	LoanReceiptNum string `json:"loanReceiptNum" description:"Loan receipt number"`
	TaskStatus     string `json:"taskStatus" description:"Task status"`
	//RepaymentStatus string `json:"repaymentStatus"`
	PageNo       int `json:"pageNo" description:"Page no" validate:"required,min=1"`
	PageRecCount int `json:"pageRecCount" description:"Page rec count" validate:"required,min=1"`
}

type IC000037O struct {
	PageNo       int        `json:"pageNo" description:"Page no"`
	PageTotCount int        `json:"pageTotCount" description:"Page tot count"`
	ListTask     []ListTask `json:"listTask" description:"List task"`
}

type ListTask struct {
	TaskId         string `json:"taskId" description:"Task id"`
	LoanReceiptNum string `json:"loanReceiptNum" description:"Loan receipt number"`
	TaskStatus     string `json:"taskStatus" description:"Task status"`
	//RepaymentStatus string          `json:"repaymentStatus"`
	CurrentDueDate string `json:"currentDueDate" description:"Current due date"`
	//PromisePayDate  string          `json:"promisePayDate"`
	//LoanAmount      decimal.Decimal `json:"loanAmount"`
	//ExecuteInterest decimal.Decimal `json:"executeInterest"`
	TaskCreatedDate      string          `json:"taskCreatedDate" description:"Task created date"`
	TaskEndDate          string          `json:"taskEndDate" description:"Task end date"`
	Currency             string          `json:"currency" description:"Currency"`
	PrincipalRepayable   decimal.Decimal `json:"principalRepayable" description:"Principal repayable"`
	InterestRepayable    decimal.Decimal `json:"interestRepayable" description:"Interest repayable"`
	FeeRepayable         decimal.Decimal `json:"feeRepayable" description:"Fee repayable"`
	TotalRepayableAmount decimal.Decimal `json:"totalRepayableAmount" description:"Total repayable amount"`
	WriteOffFlag         string          `json:"writeOffFlag" description:"Write off flag"`
	CustomerId           string          `json:"customerId" description:"Customer id"`
	TaskUserId           string          `json:"taskUserId" description:"Task user id"`
	TaskUserName         string          `json:"taskUserName" description:"Task user name"`
	ProductId            string          `json:"productId" description:"Product id"`
	AccountingAccountNum string          `json:"accountingAccountNum" description:"Accounting account number"`
}
