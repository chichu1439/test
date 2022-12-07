//Version: v1.0.0.0
package models

type IN000012I struct {
	AccountPayableNumber string `json:"accountPayableNumber" description:"Account payable number"`
	CheckType            string `json:"checkType" description:"Check type"` //1-AR,2-AP 1-挂帐，2-销账
	PageNo               int    `json:"pageNo" description:"Page no"`
	PageRecCount         int    `json:"pageRecCount" description:"Page rec count"`
}

type IN000012O struct {
	PageNo       int               `json:"pageNo" description:"Page no"`
	PageRecCount int               `json:"pageRecCount" description:"Page rec count"`
	TotalRecord  int               `json:"totalRecord" description:"Total record"`
	TotalPages   int               `json:"totalPages" description:"Total pages"`
	Record       []IN000012ORecord `json:"record" description:"Record"`
}

type IN000012ORecord struct {
	AccountPayableNumber    string  `json:"accountPayableNumber" description:"Account payable number"`       //挂账编号
	TransactionDate         string  `json:"transactionDate" description:"Transaction date"`                  //挂销帐日期
	ApArSerialNumber        int     `json:"ApArSerialNumber" description:"Ap ar serial number"`              //挂销帐序号
	TypeApAr                string  `json:"typeApAr" description:"Type ap ar"`                               //挂销账种类
	TransactionBranches     string  `json:"transactionBranches" description:"Transaction branches"`          //挂销帐机构号
	InternalAccount         string  `json:"internalAccount" description:"Internal account"`                  //内部户账号
	SerialNumber            string  `json:"serialNumber" description:"Serial number"`                        //内部户顺序号
	Currency                string  `json:"currency" description:"Currency"`                                 //币别
	DebCrtFlag              string  `json:"debCrtFlag" description:"Deb crt flag"`                           //借贷标志
	TranAmount              float64 `json:"tranAmount" description:"Tran amount"`                            //挂销帐金额
	Status                  string  `json:"status" description:"Status"`                                     //状态
	CounterpartyAccount     string  `json:"counterpartyAccount" description:"Counterparty account"`          //对方账号
	CounterpartyAccountName string  `json:"counterpartyAccountName" description:"Counterparty account name"` //对手方账户名称
	CounterpartyBranches    string  `json:"counterpartyBranches" description:"Counterparty branches"`        //对方机构
	CustomerId              string  `json:"customerId" description:"Customer id"`                            //客户号
	TellerNumber            string  `json:"tellerNumber" description:"Teller number"`                        //交易柜员
	ChannelTransaction      string  `json:"channelTransaction" description:"Channel transaction"`            //发起渠道
	SrcBizseqNo             string  `json:"srcBizseqNo" description:"Src bizseq no"`                         //业务流水号
	CounterpartyName        string  `json:"counterpartyName" description:"Counterparty name"`                //付款人名称
	AuditTeller             string  `json:"auditTeller" description:"Audit teller"`                          //授权柜员
}

func (*IN000012I) GetServiceKey() string {
	return "IN000012"
}
