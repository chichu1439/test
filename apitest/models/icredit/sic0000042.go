//Version: v1.0.0
package models

type IC000042I struct {
	LoanReceiptNum    string  `json:"loanReceiptNum" description:"Loan receipt number" validate:"required"`
	TaskId            string  `json:"taskId" description:"Task id" validate:"required"`
	ContactType       string  `json:"contactType" validate:"required" description:"Contact type:01- Home phone 02- Office phone 04- EMAIL 09- Mobile 99- other"`
	ContactMethod     string  `json:"contactMethod" description:"Contact method" validate:"required"`
	ContactNumber     string  `json:"contactNumber" description:"Contact number" validate:"required"`
	ContactPersonType string  `json:"contactPersonType" description:"Contact person type" validate:"required"`
	ContactPersonName string  `json:"contactPersonName" description:"Contact person name" validate:"required"`
	ContactStatus     string  `json:"contactStatus" description:"Contact status" validate:"required"`
	PromisePayDate    string  `json:"promisePayDate" description:"Promise pay date" `
	PromisePayAmount  float64 `json:"promisePayAmount" description:"Promise pay amount" `
	PromiseInfo       string  `json:"promiseInfo" description:"Promise information" `
	FamilyReason      string  `json:"familyReason" description:"Family reason" `
	BusinessReason    string  `json:"businessReason" description:"Business reason" `
	JobReason         string  `json:"jobReason" description:"Job reason" `
	Remark            string  `json:"remark" description:"Remark" `
	CurrentAddress    string  `json:"currentAddress" description:"Current address" `
	// TransactionRecord TransactionRecord `json:"transactionRecord" description:"Transaction record" `
}

type IC000042O struct {
	TaskId string `json:"taskId" description:"Task id"`
}

func (*IC000042I) GetServiceKey() string {
	return "IC000042"
}
