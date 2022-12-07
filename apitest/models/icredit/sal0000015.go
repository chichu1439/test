package models

import "github.com/shopspring/decimal"

type AL000015I struct {
	RegId                string `json:"regId" description:"Reg id" validate:"required"`
	LoanReceiptNum       string `json:"loanReceiptNum" description:"Loan receipt number" validate:"required"`
	AccountingAccountNum string `json:"accountingAccountNum" description:"Accounting account number" validate:"required"`
}

type AL000015O struct {
	Balance         decimal.Decimal      `json:"balance" description:"Balance"`                                                                                                                        //
	DrawTotalAmount decimal.Decimal      `json:"drawTotalAmount" description:"Draw total amount"`                                                                                                      //
	DrawBeginDate   string               `json:"drawBeginDate" description:"Draw begin date"`                                                                                                          //
	DrawTotalPeriod int                  `json:"drawTotalPeriod" description:"Draw total period"`                                                                                                      //
	MaturityDate    string               `json:"maturityDate" description:"Maturity date"`                                                                                                             //
	WriteoffType    string               `json:"writeoffType" description:"0-partial 1-total"`                                                                                                         // 0-partial 1-total
	WriteoffAmount  decimal.Decimal      `json:"writeoffAmount" description:"Writeoff amount"`                                                                                                         //
	ReasonCode      string               `json:"reasonCode" description:"Reason code:01-not repaid after repeated collections 02-borrower is bankrupt 03-borrower is missing 04-borrower is dead 06-other reason"` // 01-not repaid after repeated collections 02-borrower is bankrupt 03-borrower is missing 04-borrower is dead 06-other reason
	ProofFile       string               `json:"proofFile" description:"file url list[file1.jpg,file2.jpg…]"`                                                                                          // file url list[file1.jpg,file2.jpg…]
	Status          string               `json:"status" description:"Status:01-to be approve 02-approved 03-rejected 04-cancelled"`                                                                           // 01-to be approve 02-approved 03-rejected 04-cancelled
	BizDate         string               `json:"bizDate" description:"Business date"`                                                                                                                       //
	BizSeqNo        string               `json:"bizSeqNo" description:"Business serial number"`                                                                                                                    //
	LastActiveDate  string               `json:"lastActiveDate" description:"Last active date"`                                                                                                        //
	ListRecord      []TLoanWriteoffRcdBo `json:"listRecord" description:"List record"`
}

type TLoanWriteoffRcdBo struct {
	Description     string `json:"description" description:"Description"`
	FlowType        string `json:"flowType" description:"Flow type"`
	LastActiveDate  string `json:"lastActiveDate" description:"Last active date"`
	FinalModifyDate string `json:"finalModifyDate" description:"Final modify date"`
	FinalModifyTime string `json:"finalModifyTime" description:"Final modify time"`
	FinalModifyOrgz string `json:"finalModifyOrgz" description:"Final modify organization"`
	FinalModifyUser string `json:"finalModifyUser" description:"Final modify user"`
}

func (*AL000015I) GetServiceKey() string {
	return "AL000015"
}
