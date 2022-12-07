package models

type IC000035I struct {
	CustomerId     string `json:"customerId" description:"Customer id"`             // 客户号
	LoanReceiptNum string `json:"loanReceiptNum" description:"Loan receipt number"` //借据号
	ApplySerialNum string `json:"ApplySerialNum" description:"Apply serial number"` //申请流水号
}

type IC000035O struct {
	Status string `json:"status" description:"Status"` //成功失败状态
}

func (*IC000035O) GetServiceKey() string {
	return "IC000035"
}
