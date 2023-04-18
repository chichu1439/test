package models

import "github.com/shopspring/decimal"

type AL000019I struct {
	SuId                    string   `json:"suId" description:"Su id"`
	AccountingAccountNums   []string `json:"accountingAccountNums" description:"Accounting account nums"`
	AccountingAccountStatus []string `json:"accountingAccountStatus" description:"Accounting account status:0:normal 1:overdue 2:impaired 3:write_off 4:closed"`
}

type AL000019O struct {
	Records []TLoanAcctgAcctBo `json:"records" description:"Records"`
}

type TLoanAcctgAcctBo struct {
	AccountingAccountNum    string          `json:"accountingAccountNum" description:"Accounting account number"`                                                       //
	LoanReceiptNum          string          `json:"loanReceiptNum" description:"Loan receipt number"`                                                                   //
	AccountingAccountStatus string          `json:"accountingAccountStatus" description:"Accounting account status:0:normal 1:overdue 2:impaired 3:write_off 4:closed"` // 0-正常;1-逾期;2-减值;3-核销;4-结清
	ProductId               string          `json:"productId" description:"Product id"`                                                                                 //
	CustomerId              string          `json:"customerId" description:"Customer id"`                                                                               //
	Currency                string          `json:"currency" description:"Currency"`                                                                                    // CM0001
	RepayType               string          `json:"repayType" description:"Repay type"`                                                                                 // 01-等额本金 02-等额本息 03-到期（一次性）还本付息 04-到期还本周期还息 05-等本等息 06-随借随还
	LoanAmount              decimal.Decimal `json:"loanAmount" description:"Loan amount"`                                                                               //
	Balance                 decimal.Decimal `json:"balance" description:"Balance"`                                                                                      //
	RepayCycle              string          `json:"repayCycle" description:"Repay cycle:01- daily 02- weekly 04- Monthly 08- Yearly"`                                   // 01-天 02-周 03-半月/双周 04-月 05-双月 06-季 07-半年 08-年
	RepayFrequency          int             `json:"repayFrequency" description:"Repay frequency"`                                                                       //
	NextRepaymentDate       string          `json:"nextRepaymentDate" description:"Next repayment date"`                                                                //
	TotalPeriod             string          `json:"totalPeriod" description:"Total period"`                                                                             //
	CurrentPeriod           string          `json:"currentPeriod" description:"Current period"`                                                                         //
	DrawTotalAmount         decimal.Decimal `json:"drawTotalAmount" description:"Draw total amount"`                                                                    //
	DrawBeginDate           string          `json:"drawBeginDate" description:"Draw begin date"`                                                                        //
	DrawTotalPeriod         int             `json:"drawTotalPeriod" description:"Draw total period"`                                                                    //
	MaturityDate            string          `json:"maturityDate" description:"Maturity date"`                                                                           //
	OpenAccountDate         string          `json:"openAccountDate" description:"Open account date"`                                                                    //
	OpenAccountOrgz         string          `json:"openAccountOrgz" description:"Open account organization"`                                                            //
	GracePeriod             int             `json:"gracePeriod" description:"Grace period"`                                                                             //
	AutoRepayFlag           string          `json:"autoRepayFlag" description:"Auto repay flag:1-Auto repay 2-Manual repay"`                                            // 1-Auto repay 2-Manual repay
	RepayLocalBankFlag      string          `json:"repayLocalBankFlag" description:"Repay local bank flag"`                                                             // I-行内 O-行外
	RepayBankNum            string          `json:"repayBankNum" description:"Repay bank number"`                                                                       //
	RepayBankName           string          `json:"repayBankName" description:"Repay bank name"`                                                                        //
	RepayAccountNum         string          `json:"repayAccountNum" description:"Repay account number"`                                                                 //
	RepayAccountName        string          `json:"repayAccountName" description:"Repay account name"`                                                                  //
	BadLoanStatus           string          `json:"badLoanStatus" description:"Bad loan status:0-normal;1-poor"`                                                        // 0-正常;1-不良
	LastActiveDate          string          `json:"lastActiveDate" description:"Last active date"`                                                                      //
	PrincipalWriteoff       decimal.Decimal `json:"principalWriteoff" description:"Principal writeoff"`
}

func (*AL000019I) GetServiceKey() string {
	return "AL000019"
}
