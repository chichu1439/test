//Version: v1.0.0.0
package models

type IN000006I struct {
	InternalAccount      string  `json:"internalAccount" description:"Internal account" validate:"required"` //核算机构（X位）+币别（X位）+ 科目号（X位） + 序号（X位）。其中核算机构、币种、科目号以乐高公共的为准。顺序号设置为4位专用类型的内部户账户，在内部户维护时写入内部户编号。专用类型的内部户设定不自动开户
	AccountingItemNumber string  `json:"accountingItemNumber" description:"Accounting item number"`
	SerialNumber         string  `json:"serialNumber" description:"Serial number"`
	TranDate             string  `json:"tranDate" description:"Tran date"`
	FlowNo               string  `json:"flowNo" description:"Flow no"`
	DebCrtFlag           string  `json:"debCrtFlag" description:"Deb crt flag"`
	TranAmount           float64 `json:"tranAmount" description:"Tran amount"`
	Currency             string  `json:"currency" description:"Currency"`
	OnCancelMark         string  `json:"onCancelMark" description:"On cancel mark"`
	OnAccountNumber      string  `json:"onAccountNumber" description:"On account number"`
	ReverseCnlFlag       string  `json:"reverseCnlFlag" description:"Reverse cnl flag"`         //"N"-noraml txn,"R"-reverse txn
	BranchesNumber       string  `json:"branchesNumber" description:"Branches number"`          //机构
	TellerNumber         string  `json:"tellerNumber" description:"Teller number"`              //柜员
	OrgTranDate          string  `json:"orgTranDate" description:"Org tran date"`               // for reverse txn
	OrgGlobalBizSeqNo    string  `json:"orgGlobalBizSeqNo" description:"Org global biz seq no"` // for reverse txn
	SyncFlag             string  `json:"syncFlag" description:"Sync flag"`                      // "S"-synchronous，"A"-asynchronous
}

type IN000006O struct {
	TxnDate         string `json:"txnDate" description:"Txn date"`
	FlowNo          string `json:"flowNo" description:"Flow no"`
	SuccessFlag     string `json:"successFlag" description:"Success flag"`
	ErrCode         string `json:"errCode" description:"Err code"`
	OnAccountNumber string `json:"onAccountNumber" description:"On account number"`
}

func (*IN000006I) GetServiceKey() string {
	return "IN000006"
}
