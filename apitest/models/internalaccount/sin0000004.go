//Version: v1.0.0.0
package models

type IN000004I struct {
	SerialNumber            string // 内部户顺序号
	BranchesNumber          string // 机构代号
	Currency                string // 币别
	InternalAccount         string // 内部户账号
	AccountPayableNumber    string // 挂账编号
	AccountPayableAge       int    // 挂帐账龄
	ClearingStatus          string // 挂销账状态1-未销账(未销账与部分销账统一为未销账) 2-全部销账3-已冲正 默认为1-未销账
	AccountPayableBeginDate string // 挂账起始日期
	AccountPayableEndDate   string // 挂账结束日期
	MaturityDateFrom        string // 挂账起始到期日
	MaturityDateTo          string // 挂账结束到期日
	PageNo                  int    // 查询页数
	PageRecCount            int    // 每页记录数
}

type IN000004O struct {
	PageNo       int             `json:"pageNo" description:"Page no"`
	PageRecCount int             `json:"pageRecCount" description:"Page rec count"`
	TotalRecord  int             `json:"totalRecord" description:"Total record"`
	TotalPages   int             `json:"totalPages" description:"Total pages"`
	Item         []IN000004OItem `json:"item" description:"Item"`
}

type IN000004OItem struct {
	AccountPayableNumber    string  `json:"accountPayableNumber" description:"Account payable number"`        //挂账编号
	InternalAccount         string  `json:"internalAccount" description:"Internal account"`                   //内部户账号
	SerialNumber            string  `json:"serialNumber" description:"Serial number"`                         //内部户顺序号
	TypeApAr                string  `json:"typeApAr" description:"Type ap ar"`                                //挂销账种类
	MaturityDate            string  `json:"maturityDate" description:"Maturity date"`                         //挂账到期日
	AccountingBranches      string  `json:"accountingBranches" description:"Accounting branches"`             //核算机构
	RegisterSerialNumberMax int     `json:"registerSerialNumberMax" description:"Register serial number max"` //已挂帐最大序号
	CleanSerialNumberMax    int     `json:"cleanSerialNumberMax" description:"Clean serial number max"`       //已销账最大序号
	TranAmount              float64 `json:"tranAmount" description:"Tran amount"`                             //挂帐总金额
	AccountPayableBalance   float64 `json:"accountPayableBalance" description:"Account payable balance"`      //挂帐余额
	RegDateFirstTime        string  `json:"regDateFirstTime" description:"Reg date first time"`               //首次挂帐日期
	RegDateLastTime         string  `json:"regDateLastTime" description:"Reg date last time"`                 //末次挂帐日期
	LastTransactionDate     string  `json:"lastTransactionDate" description:"Last transaction date"`          //最后交易日期
	OnCancelFlag            string  `json:"onCancelFlag" description:"On cancel flag"`                        //挂帐记录状态
}

func (*IN000004I) GetServiceKey() string {
	return "IN000004"
}
