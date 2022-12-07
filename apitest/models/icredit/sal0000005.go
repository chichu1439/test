package models

import "github.com/shopspring/decimal"

type AL000005I struct {
	AccountingAccountNum      string                       `json:"accountingAccountNum" description:"Accounting account number" validate:"required"`
	TLoanBalanceChangeDetail  *TLoanBalanceChangeDetailBo  `json:"tLoanBalanceChangeDetail" description:"T loan balance change detail"`
	TLoanBalanceChangeHistory *TLoanBalanceChangeHistoryBo `json:"tLoanBalanceChangeHistory" description:"T loan balance change history"`
}

type AL000005O struct {
}

type TLoanBalanceChangeDetailBo struct {
	TransDate                string          `json:"transDate" description:"Trans date"`                                 // ",
	BusinessSerialNum        string          `json:"businessSerialNum" description:"Business serial number"`             // ",
	AccountingAccountNum     string          `json:"accountingAccountNum" description:"Accounting account number"`       // ",
	Period                   string          `json:"period" description:"Period"`                                        // ",
	TransType                string          `json:"transType" description:"Trans type"`                                 // 03-RPYM 04-DISB 05-WTOF ",
	TransMethod              string          `json:"transMethod" description:"Trans method"`                             // O-联机;B-批量 ",
	ProductId                string          `json:"productId" description:"Product id"`                                 // ",
	TransCurrency            string          `json:"transCurrency" description:"Trans currency"`                         // CM0001 ",
	TransAmount              decimal.Decimal `json:"transAmount" description:"Trans amount"`                             // ",
	DuePayTotalAmount        decimal.Decimal `json:"duePayTotalAmount" description:"Due pay total amount"`               // ",
	DuePayPrincipal          decimal.Decimal `json:"duePayPrincipal" description:"Due pay principal"`                    // ",
	Balance                  decimal.Decimal `json:"balance" description:"Balance"`                                      // ",
	RepayType                string          `json:"repayType" description:"Repay type"`                                 // 01-提前部分还款 02-提前结清还款 03-逾期转正常还款 04-正常还款 还当期 05-月结还款 06-逾期追扣 07-减免还款 11-减值收回 12-部分减值收回 13-核销收回 ",
	PositiveReverseTransFlag string          `json:"positiveReverseTransFlag" description:"Positive reverse trans flag"` // 1-正交易 2-反交易 3-不入账交易 ",
	RectificationTransFlag   string          `json:"rectificationTransFlag" description:"Rectification trans flag"`      // N-正常交易 C-冲销交易 B-已冲正交易 ",
	OriginalTransDate        string          `json:"originalTransDate" description:"Original trans date"`                // ",
	OriginalTransSerialNum   string          `json:"originalTransSerialNum" description:"Original trans serial number"`  // ",
	FinalModifyDate          string          `json:"finalModifyDate" description:"Final modify date"`                    // ",
	FinalModifyTime          string          `json:"finalModifyTime" description:"Final modify time"`                    // ",
	FinalModifyOrgz          string          `json:"finalModifyOrgz" description:"Final modify organization"`                    // ",
	FinalModifyUser          string          `json:"finalModifyUser" description:"Final modify user"`                    // ",
	TccStatus                int             `json:"tccStatus" description:"Tcc status"`                                 // ",
}

type TLoanBalanceChangeHistoryBo struct {
	BalanceDate             string `json:"balanceDate" description:"Balance date"`                          // ",
	AccountingAccountDate   string `json:"accountingAccountDate" description:"Accounting account date"`     // ",
	Period                  string `json:"period" description:"Period"`                                     // ",
	ProductId               string `json:"productId" description:"Product id"`                              // ",
	AccountingAccountStatus string `json:"accountingAccountStatus" description:"Accounting account status:0:normal 1:overdue 2:impaired 3:write_off 4:closed"` // 0-正常;1-逾期;2-减值;3-核销;4-结清 ",
	Currency                string `json:"currency" description:"Currency"`                                 // CM0001 ",
	Balance                 string `json:"balance" description:"Balance"`                                   // ",
	OpenAccountOrgz         string `json:"openAccountOrgz" description:"Open account organization"`                 // ",
	FinalModifyDate         string `json:"finalModifyDate" description:"Final modify date"`                 // ",
	FinalModifyTime         string `json:"finalModifyTime" description:"Final modify time"`                 // ",
	FinalModifyOrgz         string `json:"finalModifyOrgz" description:"Final modify organization"`                 // ",
	FinalModifyUser         string `json:"finalModifyUser" description:"Final modify user"`                 // ",
	TccStatus               int    `json:"tccStatus" description:"Tcc status"`                              // 0-有效 1-无效 ",
}

func (*AL000005I) GetServiceKey() string {
	return "AL000005"
}
