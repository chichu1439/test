//Version: v1
package models

type IC000082Request struct {
	CustomerId                  string `json:"customerId"`
	RestructuringApplySerialNum string `json:"restructuringApplySerialNum"`
	LoanReceiptNum              string `json:"loanReceiptNum" description:"Loan receipt number"`
	ApplyType                   string `json:"applyType" description:"1 - loan application;2 - modification application; 3 - repayment application; 4 - secondary loan application; 5- restrcturing loan;"`
	ApplyMode                   string `json:"applyMode" description:"If there is no any pending application(t_loan_apply_flow) within today, Go to check product status and customer status and valid period."`
}

type IC000082Response struct {
	CustomerValid string      `json:"customerValid"`
	ProductValid  string      `json:"productValid"`
	IsPending     string      `json:"isPending"`
	ApplyList     []ApplyInfo `json:"applyList"`
}
type ApplyInfo struct {
	ApplyDate              string `json:"applyDate"`
	BizSeq                 string `json:"bizSeq"`
	ApplySerialNum         string `json:"applySerialNum"`
	MakeLoanApplySerialNum string `json:"makeLoanApplySerialNum"`
	LoanReceiptNum         string `json:"loanReceiptNum"`
	CreateTime             string `json:"createTime"`
	OperateUserId          string `json:"operateUserId"`
	OperateOrgzNum         string `json:"operateOrgzNum"`
	OperateComment         string `json:"operateComment"`
	OperateStatus          string `json:"operateStatus"`
	ApplyType              string `json:"applyType"`
	SourceChannel          string `json:"sourceChannel"`
}

func (*IC000082Request) GetServiceKey() string {
	return "IC000082"
}
