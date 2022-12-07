package models

import (
	"github.com/shopspring/decimal"
	"sort"
)

type AL000014I struct {
	RegId           string   `json:"regId" description:"Reg id"`
	WriteOffStatus  []string `json:"writeOffStatus" description:"Write off status:01-to be approve 02-approved 03-rejected 04-cancelled"`
	StrDate         string   `json:"strDate" description:"Str date"`
	EndDate         string   `json:"endDate" description:"End date"`
	PageNum         int      `json:"pageNum" description:"Page number"`
	PageRecordCount int      `json:"pageRecordCount" description:"Page record count"`
}

type AL000014O struct {
	PageNum         int                       `json:"pageNum" description:"Page number"`
	PageRecordCount int                       `json:"pageRecordCount" description:"Page record count"`
	TotalCount      int                       `json:"totalCount" description:"Total count"`
	Records         []TLoanWriteoffRegisterBo `json:"records" description:"Records"`
}

type TLoanWriteoffRegisterBo struct {
	RegId                string          `json:"regId" description:"Reg id"`                                                                                                                                       //
	CustomerId           string          `json:"customerId" description:"Customer id"`                                                                                                                             //
	LoanReceiptNum       string          `json:"loanReceiptNum" description:"Loan receipt number"`                                                                                                                 //
	AccountingAccountNum string          `json:"accountingAccountNum" description:"Accounting account number"`                                                                                                     //
	Balance              decimal.Decimal `json:"balance" description:"Balance"`                                                                                                                                    //
	DrawTotalAmount      decimal.Decimal `json:"drawTotalAmount" description:"Draw total amount"`                                                                                                                  //
	DrawBeginDate        string          `json:"drawBeginDate" description:"Draw begin date"`                                                                                                                      //
	DrawTotalPeriod      int             `json:"drawTotalPeriod" description:"Draw total period"`                                                                                                                  //
	MaturityDate         string          `json:"maturityDate" description:"Maturity date"`                                                                                                                         //
	WriteoffType         string          `json:"writeoffType" description:"Writeoff type:0-partial 1-total"`                                                                                                       // 0-partial 1-total
	WriteoffAmount       decimal.Decimal `json:"writeoffAmount" description:"Writeoff amount"`                                                                                                                     //
	ReasonCode           string          `json:"reasonCode" description:"Reason code:01-not repaid after repeated collections 02-borrower is bankrupt 03-borrower is missing 04-borrower is dead 06-other reason"` // 01-not repaid after repeated collections 02-borrower is bankrupt 03-borrower is missing 04-borrower is dead 06-other reason
	ProofFile            string          `json:"proofFile" description:"file url list[file1.jpg,file2.jpg…]"`                                                                                                      // file url list[file1.jpg,file2.jpg…]
	Status               string          `json:"status" description:"Status:01-to be approve 02-approved 03-rejected 04-cancelled"`                                                                                // 01-to be approve 02-approved 03-rejected 04-cancelled
	BizDate              string          `json:"bizDate" description:"Business date"`                                                                                                                                   //
	BizSeqNo             string          `json:"bizSeqNo" description:"Business serial number"`                                                                                                                                //
	LastActiveDate       string          `json:"lastActiveDate" description:"Last active date"`                                                                                                                    //
}

//排序  start

type ListSortAl14 struct {
	List []TLoanWriteoffRegisterBo `json:"list"`
}

func (m *ListSortAl14) Len() int {
	return len(m.List)
}

func (m *ListSortAl14) Less(i, j int) bool {
	b := m.List[i].RegId > m.List[j].RegId
	return b
}

func (m *ListSortAl14) Swap(i, j int) {
	m.List[i], m.List[j] = m.List[j], m.List[i]
}

func (*AL000014I) Sort(maps []TLoanWriteoffRegisterBo) []TLoanWriteoffRegisterBo {
	listSort := ListSortAl14{}
	listSort.List = maps
	sort.Sort(&listSort)
	return listSort.List
}

func (*AL000014I) GetServiceKey() string {
	return "AL000014"
}
