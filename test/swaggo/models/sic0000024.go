package models

import (
	"github.com/shopspring/decimal"
	"sort"
)

type IC000024I struct {
	LoanReceiptNum    string `json:"loanReceiptNum" description:"Loan receipt number" validate:"required"`
	CustomerId        string `json:"customerId" description:"Customer id"`
	BusinessSerialNum string `json:"businessSerialNum" description:"Business serial number"`
	GlobalSerialNum   string `json:"globalSerialNum" description:"Global serial number"`
	TransDate         string `json:"transDate" description:"Trans date"`
	PageNum           int    `json:"pageNum" description:"Page number"`
	PageSize          int    `json:"pageSize" description:"Page size" validate:"required,gt=0"`
}

type IC000024O struct {
	PageNum         int               `json:"pageNum" description:"Page number"`
	PageRecordCount int               `json:"pageRecordCount" description:"Page record count"`
	TotalCount      int               `json:"totalCount" description:"Total count"`
	LoanReceiptNum  string            `json:"loanReceiptNum" description:"Loan receipt number"`
	CustomerId      string            `json:"customerId" description:"Customer id"`
	Records         []IC000024ORecord `json:"records" description:"Records"`
}

type IC000024ORecord struct {
	TransDate                 string          `json:"transDate" description:"Trans date"`
	TransTime                 string          `json:"transTime" description:"Trans time"`
	BusinessSerialNum         string          `json:"businessSerialNum" description:"Business serial number"`
	GlobalSerialNum           string          `json:"globalSerialNum" description:"Global serial number"`
	TransType                 string          `json:"transType" description:"Trans type:03-RPYM 04-DISB 05-WTOF"` //"03-RPYM 04-DISB 05-WTOF"
	RepayType                 string          `json:"repayType"`
	TransCurrency             string          `json:"transCurrency" description:"Trans currency"`
	TransAmount               decimal.Decimal `json:"transAmount" description:"Trans amount"`
	ActualPayPrincipal        decimal.Decimal `json:"actualPayPrincipal" description:"Actual pay principal"`
	ActualPayInterest         decimal.Decimal `json:"actualPayInterest" description:"Actual pay interest"`
	ActualPayFee              decimal.Decimal `json:"actualPayFee" description:"Actual pay fee"`
	Balance                   decimal.Decimal `json:"balance" description:"Balance"`
	CounterpartyLocalBankFlag string          `json:"counterpartyLocalBankFlag" description:"Counterparty local bank flag"`
	CounterpartyBankNum       string          `json:"counterpartyBankNum" description:"Counterparty bank number"`
	CounterpartyBankName      string          `json:"counterpartyBankName" description:"Counterparty bank name"`
	CounterpartyAccountNum    string          `json:"counterpartyAccountNum" description:"Counterparty account number"`
	CounterpartyAccountName   string          `json:"counterpartyAccountName" description:"Counterparty account name"`
	CounterpartyTransDate     string          `json:"counterpartyTransDate" description:"Counterparty trans date"`
	CounterpartySerialNum     string          `json:"counterpartySerialNum" description:"Counterparty serial number"`
	ApplyStatus               string          `json:"applyStatus" description:"Apply status"`
	MakerUserId               string          `json:"makerUserId"`
	MakerComment              string          `json:"makerComment"`
	CheckerUserId             string          `json:"checkerUserId"`
	CheckerComment            string          `json:"checkerComment"`
	FeeTotal                  decimal.Decimal `json:"feeTotal"`
	CancelFlag                bool            `json:"cancelFlag"`
	OrigBusinessSerialNum     string          `json:"origBusinessSerialNum"`
	OrigRepayType             string          `json:"origRepayType"`
}

//排序  start

type ListSort struct {
	List []IC000024ORecord `json:"list"`
}

func (m *ListSort) Len() int {
	return len(m.List)
}

func (m *ListSort) Less(i, j int) bool {
	b := m.List[i].TransDate > m.List[j].TransDate
	if m.List[i].TransDate == m.List[j].TransDate {
		b = m.List[i].TransTime > m.List[j].TransTime
	}
	return b
}

func (m *ListSort) Swap(i, j int) {
	m.List[i], m.List[j] = m.List[j], m.List[i]
}

func (*IC000024I) Sort(maps []IC000024ORecord) []IC000024ORecord {
	listSort := ListSort{}
	listSort.List = maps
	sort.Sort(&listSort)
	return listSort.List
}

func (*IC000024I) GetServiceKey() string {
	return "IC000024"
}
