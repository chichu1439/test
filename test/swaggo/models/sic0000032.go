package models

type IC000032I struct {
	LoanReceiptNum            string `json:"loanReceiptNum" description:"Loan receipt number"`
	CustomerId                string `json:"customerId" description:"Customer id"`
	CustomerName              string `json:"customerName" description:"Customer name"`
	ChangeApplyStartDate      string `json:"changeApplyStartDate" description:"Change apply start date"`
	ChangeApplyEndDate        string `json:"changeApplyEndDate" description:"Change apply end date"`
	LoanElementClassification string `json:"loanElementClassification" description:"Loan element classification"`
	PageNum                   int    `json:"pageNum" description:"Page number"`
	PageRecordCount           int    `json:"pageRecordCount" description:"Page record count"`
}

type IC000032O struct {
	PageNum         int                 `json:"pageNum" description:"Page number"`
	PageRecordCount int                 `json:"pageRecordCount" description:"Page record count"`
	TotalCount      int                 `json:"totalCount" description:"Total count"`
	Records         []LoanChangeHistory `json:"records" description:"Records"`
}

type LoanChangeHistory struct {
	ChangeApplyDate           string `json:"changeApplyDate" description:"Change apply date"`
	LoanReceiptNum            string `json:"loanReceiptNum" description:"Loan receipt number"`
	LoanChangeApplySerialNum  string `json:"loanChangeApplySerialNum" description:"Loan change apply serial number"`
	ChangeApplyTime           string `json:"changeApplyTime" description:"Change apply time"`
	ElementNewValue           string `json:"elementNewValue" description:"Element new value"`
	ElementOldValue           string `json:"elementOldValue" description:"Element old value"`
	LoanElement               string `json:"loanElement" description:"Loan element"`
	LoanElementClassification string `json:"loanElementClassification" description:"Loan element classification"`
	ApplyUser                 string `json:"applyUser" description:"Apply user"`
	ApproveUser               string `json:"approveUser" description:"Approve user"`
}

func (*IC000032I) GetServiceKey() string {
	return "IC000032"
}
