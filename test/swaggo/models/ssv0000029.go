package models

type SV000029I struct {
	MediaType    string `json:"mediaType" validate:"max=4" description:"Medium type"`
	MediaNumber  string `json:"mediaNumber" validate:"max=40" description:"Medium number"`
	Currency     string `json:"currency"`
	CashtranFlag string `json:"cashTransFlag"`
}

type SV000029O struct {
	AccountList []AccountInfo `json:"accountList"`
}
type AccountInfo struct {
	AgreementId       string `json:"agreementId"`
	AgreementType     string `json:"agreementType"`
	AccountName       string `json:"accountName"`
	Currency          string `json:"currency" description:"Currency"`                           //币种
	CashtranFlag      string `json:"cashTransFlag" description:"Flag of cash and remittance"`   //钞汇标识
	FreezeType        string `json:"freezeType" description:"Frozen state"`                     //冻结状态
	AgreementStatus   string `json:"agreementStatus" description:"Conttract state"`             //合约状态
	AccountingAccount string `json:"accountingAccount" description:"Accounting account number"` //核算账号
	WithdrawMethod    string `json:"withdrawMethod" description:"Withdrawal method"`            //支取方式
	DepcreFlag        string `json:"depcreFlag" description:"Debit credit control mark"`        //借贷记控制标记
	CustomerId        string `json:"customerId"`
}

func (*SV000029I) GetServiceKey() string {
	return "SV000029"
}
