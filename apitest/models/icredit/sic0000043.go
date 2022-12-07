//Version: v1.0.0
package models

type IC000043I struct {
	LoanReceiptNum string `json:"loanReceiptNum" description:"Loan receipt number" validate:"required"`
	TaskId         string `json:"taskId" description:"Task id" validate:"required"`
	ContactType    string `json:"contactType" description:"Contact type:01-Home phone 02-Office phone 04-EMAIL 09-Mobile 99-other"`
	ContactMethod  string `json:"contactMethod" description:"Contact method" `
	ContactStatus  string `json:"contactStatus" description:"Contact status" `
	PageNo         int    `json:"pageNo" description:"Page no" validate:"required"`
	PageRecCount   int    `json:"pageRecCount" description:"Page rec count" validate:"required"`
}

type IC000043O struct {
	PageNo       int       `json:"pageNo" description:"Page no"`
	PageTotCount int       `json:"pageTotCount" description:"Page tot count"`
	ListContact  []Contact `json:"listContact" description:"List contact"`
}

type Contact struct {
	Uid              int     `json:"uid" description:"Uid" `
	ContactIndex     int     `json:"contactIndex" description:"Contact index" `
	ContactType      string  `json:"contactType" description:"Contact type" `
	ContactMethod    string  `json:"contactMethod" description:"Contact method" `
	ContactNumber    string  `json:"contactNumber" description:"Contact number" `
	ContactStatus    string  `json:"contactStatus" description:"Contact status" `
	ContactTime      string  `json:"contactTime" description:"Contact time" `
	PromisePayDate   string  `json:"promisePayDate" description:"Promise pay date" `
	PromisePayAmount float64 `json:"promisePayAmount" description:"Promise pay amount" `
}

func (*IC000043I) GetServiceKey() string {
	return "IC000043"
}
