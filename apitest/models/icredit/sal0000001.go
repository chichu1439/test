package models

import "github.com/shopspring/decimal"

type AL000001I struct {
	AccountingAccountNum string `json:"accountingAccountNum" description:"Accounting account number"`
	TransDate            string `json:"transDate" description:"Trans date"`
	Period               int    `json:"period" description:"Period"`
}

type AL000001O struct {
	TLoanInterestTypes          []TLoanInterestType          `json:"tLoanInterestTypes" description:"Loan interest types"`
	TLoanAcctgAcct              *TLoanAcctgAcct              `json:"tLoanAcctgAcct" description:"Loan accounting account"`
	TLoanInterestAccrual        *TLoanInterestAccrual        `json:"tLoanInterestAccrual" description:"Loan interest accrual"`
	TAcctLoanAcctgPeriods       []TAcctLoanAcctgPeriod       `json:"tAcctLoanAcctgPeriods" description:"Account loan accounting periods"`
	TLoanInterestAccrualPeriods []TLoanInterestAccrualPeriod `json:"tLoanInterestAccrualPeriods" description:"Loan interest accrual periods"`
	BillRuleInfo                *BillRuleInfo                `json:"billRuleInfo"`
	TLoanFeePeriod              []TLoanFeePeriod             `json:"TLoanFeePeriod"`
}

type TLoanAcctgAcct struct {
	AccountingAccountNum    string          `json:"accountingAccountNum" description:"Accounting account number"`                                                       // ",
	LoanReceiptNum          string          `json:"loanReceiptNum" description:"Loan receipt number"`                                                                   // ",
	AccountingAccountStatus string          `json:"accountingAccountStatus" description:"Accounting account status:0:normal 1:overdue 2:impaired 3:write_off 4:closed"` // 0-正常;1-逾期;2-减值;3-核销;4-结清 ",
	ProductId               string          `json:"productId" description:"Product id"`                                                                                 // ",
	CustomerId              string          `json:"customerId" description:"Customer id"`                                                                               // ",
	Currency                string          `json:"currency" description:"Currency"`                                                                                    // CM0001 ",
	RepayType               string          `json:"repayType" description:"Repay type"`                                                                                 // 01-等额本金 02-等额本息 03-到期（一次性）还本付息 04-到期还本周期还息 05-等本等息 06-随借随还 ",
	LoanAmount              decimal.Decimal `json:"loanAmount" description:"Loan amount"`                                                                               // ",
	Balance                 decimal.Decimal `json:"balance" description:"Balance"`                                                                                      // ",
	RepayCycle              string          `json:"repayCycle" description:"Repay cycle"`                                                                               // 01-天 02-周 03-半月/双周 04-月 05-双月 06-季 07-半年 08-年 ",
	RepayFrequency          int             `json:"repayFrequency" description:"Repay frequency"`                                                                       // ",
	RepayMethod             string          `json:"repayMethod" description:"Repay method"`                                                                             // 01-等额本金 02-等额本息 03-到期（一次性）还本付息 04-到期还本周期还息 05-等本等息 06-随借随还 ",
	NextRepaymentDate       string          `json:"nextRepaymentDate" description:"Next repayment date"`                                                                // ",
	TotalPeriod             int             `json:"totalPeriod" description:"Total period"`                                                                             // ",
	RepayDay                string          `json:"repayDay" description:"Repay day"`                                                                                   // ",
	CurrentPeriod           int             `json:"currentPeriod" description:"Current period"`                                                                         // ",
	DrawTotalAmount         decimal.Decimal `json:"drawTotalAmount" description:"Draw total amount"`                                                                    // ",
	DrawBeginDate           string          `json:"drawBeginDate" description:"Draw begin date"`                                                                        // ",
	DrawTotalPeriod         int             `json:"drawTotalPeriod" description:"Draw total period"`                                                                    // ",
	MaturityDate            string          `json:"maturityDate" description:"Maturity date"`                                                                           // ",
	OpenAccountDate         string          `json:"openAccountDate" description:"Open account date"`                                                                    // ",
	OpenAccountOrgz         string          `json:"openAccountOrgz" description:"Open account organization"`                                                            // ",
	GracePeriod             int             `json:"gracePeriod" description:"Grace period"`                                                                             // ",
	AutoRepayFlag           string          `json:"autoRepayFlag" description:"Auto repay flag:1-Auto repay 2-Manual repay"`                                            // 1-Auto repay 2-Manual repay ",
	RepayLocalBankFlag      string          `json:"repayLocalBankFlag" description:"Repay local bank flag"`                                                             // I-行内 O-行外 ",
	RepayBankNum            string          `json:"repayBankNum" description:"Repay bank number"`                                                                       // ",
	RepayBankName           string          `json:"repayBankName" description:"Repay bank name"`                                                                        // ",
	RepayAccountNum         string          `json:"repayAccountNum" description:"Repay account number"`                                                                 // ",
	RepayAccountName        string          `json:"repayAccountName" description:"Repay account name"`                                                                  // ",
	BadLoanStatus           string          `json:"badLoanStatus" description:"Bad loan status"`                                                                        // 0-正常;1-不良 ",
	CompoundInterestFlag    string          `json:"compoundInterestFlag" description:"Compound interest flag:0-No 1-Yes"`                                               // 0-No 1-Yes ",
	LastActiveDate          string          `json:"lastActiveDate" description:"Last active date"`                                                                      // ",
	//FinalModifyDate         string          `json:"finalModifyDate"`         //    ",
	//FinalModifyTime         string          `json:"finalModifyTime"`         //    ",
	//FinalModifyOrgz         string          `json:"finalModifyOrgz"`         //    ",
	//FinalModifyUser         string          `json:"finalModifyUser"`         //    ",
	//TccStatus               int             `json:"tccStatus"`               // 0-有效 1-无效   ",
}

type TAcctLoanAcctgPeriod struct {
	AccountingAccountNum   string          `json:"accountingAccountNum" description:"Accounting account number"`                                            // ",
	Period                 int             `json:"period" description:"Period"`                                                                             // ",
	CurrentAccountStatus   string          `json:"currentAccountStatus" description:"Current account status"`                                               // 0-正常;1-逾期;2-减值;3-核销;4-结清 ",
	PeriodRepayStatus      string          `json:"periodRepayStatus" description:"Period repay status:O-Unpaid 1-Partial Paid 2-Total Paid 3-Waived"`       // O-Unpaid 1-Partial Paid 2-Total Paid 3-Waived ",
	PrincipalRepayStatus   string          `json:"principalRepayStatus" description:"Principal repay status:O-Unpaid 1-Partial Paid 2-Total Paid 3-Waived"` // O-Unpaid 1-Partial Paid 2-Total Paid 3-Waived ",
	FeeRepayStatus         string          `json:"feeRepayStatus" description:"Fee repay status:O-Unpaid 1-Partial Paid 2-Total Paid 3-Waived"`             // O-Unpaid 1-Partial Paid 2-Total Paid 3-Waived ",
	CurrentPeriodBeginDate string          `json:"currentPeriodBeginDate" description:"Current period begin date"`                                          // ",
	CurrentPeriodEndDate   string          `json:"currentPeriodEndDate" description:"Current period end date"`                                              // ",
	CurrentRepayDate       string          `json:"currentRepayDate" description:"Current repay date"`                                                       // ",
	Currency               string          `json:"currency" description:"Currency"`                                                                         // CM0001 ",
	PlanRepayPrincipal     decimal.Decimal `json:"planRepayPrincipal" description:"Plan repay principal"`                                                   // ",
	PlanRepayInterest      decimal.Decimal `json:"planRepayInterest" description:"Plan repay interest"`                                                     // ",
	ActualRepayDate        string          `json:"actualRepayDate" description:"Actual repay date"`                                                         // ",
	ActualRepayPrincipal   decimal.Decimal `json:"actualRepayPrincipal" description:"Actual repay principal"`                                               // ",
	ActualRepayInterest    decimal.Decimal `json:"actualRepayInterest" description:"Actual repay interest"`                                                 // ",
	InterestWaived         decimal.Decimal `json:"interestWaived" description:"Interest waived"`                                                            // ",
	PrincipalWaived        decimal.Decimal `json:"principalWaived" description:"Principal waived"`                                                          // ",
	PrincipalImpaired      decimal.Decimal `json:"principalImpaired" description:"Principal impaired"`                                                      // ",
	PrincipalWriteoff      decimal.Decimal `json:"principalWriteoff" description:"Principal writeoff"`                                                      // ",
	FeeAmount              decimal.Decimal `json:"feeAmount" description:"Fee amount"`                                                                      // ",
	FeeRepay               decimal.Decimal `json:"feeRepay" description:"Fee repay"`                                                                        // ",
	FeeWaived              decimal.Decimal `json:"feeWaived" description:"Fee waived"`                                                                      // ",
	OverdueStartDate       string          `json:"overdueStartDate" description:"Overdue start date"`                                                       // ",
	OverdueDays            int             `json:"overdueDays" description:"Overdue days"`                                                                  // ",
	CloseDate              string          `json:"closeDate" description:"Close date"`                                                                      // ",
	CurrentUnpaidPrincipal decimal.Decimal `json:"currentUnpaidPrincipal" description:"Current unpaid principal"`                                           // ",
	MaintainPrinciple      decimal.Decimal `json:"maintainPrinciple" description:"Maintain principle"`                                                      // ",
	LastActiveDate         string          `json:"lastActiveDate" description:"Last active date"`                                                           // ",
	InterestWriteoff       decimal.Decimal `json:"interestWriteoff" description:"Interest writeoff"`                                                        // ",
	//FinalModifyDate         string          `json:"finalModifyDate"`         //    ",
	//FinalModifyTime         string          `json:"finalModifyTime"`         //    ",
	//FinalModifyOrgz         string          `json:"finalModifyOrgz"`         //    ",
	//FinalModifyUser         string          `json:"finalModifyUser"`         //    ",
	//TccStatus               int             `json:"tccStatus"`               // 0-有效 1-无效   ",
}

type TLoanInterestAccrual struct {
	AccountingAccountNum  string          `json:"accountingAccountNum" description:"Accounting account number"` // ",
	Currency              string          `json:"currency" description:"Currency"`                              // ",
	InterestAccountStatus string          `json:"interestAccountStatus" description:"Interest account status"`  // 0-正常;1-逾期;2-减值;3-核销;4-结清 ",
	InterestBeginDate     string          `json:"interestBeginDate" description:"Interest begin date"`          // ",
	InterestEndDate       string          `json:"interestEndDate" description:"Interest end date"`              // ",
	LastAccrualDate       string          `json:"lastAccrualDate" description:"Last accrual date"`              // ",
	NextAccrualDate       string          `json:"nextAccrualDate" description:"Next accrual date"`              // ",
	LastSettlementDate    string          `json:"lastSettlementDate" description:"Last settlement date"`        // ",
	NextSettlementDate    string          `json:"nextSettlementDate" description:"Next settlement date"`        // ",
	InterestRepaymentDate string          `json:"interestRepaymentDate" description:"Interest repayment date"`  // ",
	InterestPaid          decimal.Decimal `json:"interestPaid" description:"Interest paid"`                     // ",
	InterestWaived        decimal.Decimal `json:"interestWaived" description:"Interest waived"`                 // ",
	InterestImpaired      decimal.Decimal `json:"interestImpaired" description:"Interest impaired"`             // ",
	InterestWriteoff      decimal.Decimal `json:"interestWriteoff" description:"Interest writeoff"`             // ",
	InterestAfterWriteoff decimal.Decimal `json:"interestAfterWriteoff" description:"Interest after writeoff"`  // ",
	InterestAccrual       decimal.Decimal `json:"interestAccrual" description:"Interest accrual"`               // ",
	InterestSettlement    decimal.Decimal `json:"interestSettlement" description:"Interest settlement"`         // ",
	LastCalcDate          string          `json:"lastCalcDate" description:"Last calc date"`
	InterestCalc          decimal.Decimal `json:"interestCalc" description:"Interest calc"`
	//FinalModifyDate       string          `json:"finalModifyDate"`       //    ",
	//FinalModifyTime       string          `json:"finalModifyTime"`       //    ",
	//FinalModifyOrgz       string          `json:"finalModifyOrgz"`       //    ",
	//FinalModifyUser       string          `json:"finalModifyUser"`       //    ",
	//TccStatus             int             `json:"tccStatus"`             // 0-有效 1-无效   ",
}

type TLoanInterestAccrualPeriod struct {
	AccountingAccountNum  string          `json:"accountingAccountNum" description:"Accounting account number"`                                          // ",
	Period                int             `json:"period" description:"Period"`                                                                           // ",
	InterestType          string          `json:"interestType" description:"Interest type:O-Normal Loan 1-Overdue Loan"`                                 // O-Normal Loan 1-Overdue Loan ",
	CompoundFlag          string          `json:"compoundFlag" description:"Compound flag:0-No 1-Yes"`                                                   // 0-No 1-Yes ",
	Currency              string          `json:"currency" description:"Currency"`                                                                       // ",
	InterestAccountStatus string          `json:"interestAccountStatus" description:"Interest account status"`                                           // 0-正常;1-逾期;2-减值;3-核销;4-结清 ",
	InterestRepayStatus   string          `json:"interestRepayStatus" description:"Interest repay status:O-Unpaid 1-Partial Paid 2-Total Paid 3-Waived"` // O-Unpaid 1-Partial Paid 2-Total Paid 3-Waived ",
	InterestBeginDate     string          `json:"interestBeginDate" description:"Interest begin date"`                                                   // ",
	InterestEndDate       string          `json:"interestEndDate" description:"Interest end date"`                                                       // ",
	LastAccrualDate       string          `json:"lastAccrualDate" description:"Last accrual date"`                                                       // ",
	NextAccrualDate       string          `json:"nextAccrualDate" description:"Next accrual date"`                                                       // ",
	LastSettlementDate    string          `json:"lastSettlementDate" description:"Last settlement date"`                                                 // ",
	NextSettlementDate    string          `json:"nextSettlementDate" description:"Next settlement date"`                                                 // ",
	InterestRepaymentDate string          `json:"interestRepaymentDate" description:"Interest repayment date"`                                           // ",
	InterestPaid          decimal.Decimal `json:"interestPaid" description:"Interest paid"`                                                              // ",
	InterestWaived        decimal.Decimal `json:"interestWaived" description:"Interest waived"`                                                          // ",
	InterestImpaired      decimal.Decimal `json:"interestImpaired" description:"Interest impaired"`                                                      // ",
	InterestWriteoff      decimal.Decimal `json:"interestWriteoff" description:"Interest writeoff"`                                                      // ",
	InterestAfterWriteoff decimal.Decimal `json:"interestAfterWriteoff" description:"Interest after writeoff"`                                           // ",
	InterestAccrual       decimal.Decimal `json:"interestAccrual" description:"Interest accrual"`                                                        // ",
	InterestSettlement    decimal.Decimal `json:"interestSettlement" description:"Interest settlement"`                                                  // ",
	LastCalcDate          string          `json:"lastCalcDate" description:"Last calc date"`
	InterestCalc          decimal.Decimal `json:"interestCalc" description:"Interest calc"`
	//FinalModifyDate       string          `json:"finalModifyDate"`       //    ",
	//FinalModifyTime       string          `json:"finalModifyTime"`       //    ",
	//FinalModifyOrgz       string          `json:"finalModifyOrgz"`       //    ",
	//FinalModifyUser       string          `json:"finalModifyUser"`       //    ",
	//TccStatus             int             `json:"tccStatus"`             // 0-有效 1-无效   ",
}

type TLoanInterestType struct {
	LoanReceiptNum       string          `json:"loanReceiptNum" description:"Loan receipt number"`                                                                                       // ",
	AccountingAccountNum string          `json:"accountingAccountNum" description:"Accounting account number"`                                                                           // ",
	InterestType         string          `json:"interestType" description:"Interest type:0-normal loan 1-overdue loan"`                                                                  // 0-normal loan 1-overdue loan ",
	InterestId           string          `json:"interestId" description:"Interest id"`                                                                                                   // ",
	StrategyId           string          `json:"strategyId" description:"Strategy id"`                                                                                                   // ",
	BaseInterestRate     decimal.Decimal `json:"baseInterestRate" description:"Base interest rate"`                                                                                      // ",
	FloatDirection       string          `json:"floatDirection" description:"Float direction:0-None 1-Up 2-Down"`                                                                        // 0-None 1-Up 2-Down ",
	FloatMethod          string          `json:"floatMethod" description:"Float method:0-By Percent 1-BP Value"`                                                                         // 0-By Percent 1-BP Value ",
	FloatValue           decimal.Decimal `json:"floatValue" description:"Float value"`                                                                                                   // ",
	InterestRate         decimal.Decimal `json:"interestRate" description:"Interest rate"`                                                                                               // ",
	CalculateType        string          `json:"calculateType" description:"Calculate type:0-Actual days 1-Year to month to day 2-Year to month"`                                        // 0-Actual days 1-Year to month to day 2-Year to month ",
	RateExecType         string          `json:"rateExecType" description:"Rate exec type:0-By Value Date 1-By Calc Date"`                                                               // 0-By Value Date 1-By Calc Date ",
	FloatExecType        string          `json:"floatExecType" description:"Float exec type:0-By Value Date 1-By Calc Date 2-Fixed float"`                                               // 0-By Value Date 1-By Calc Date 2-Fixed float ",
	MonthDays            string          `json:"monthDays" description:"Month days:0-30 1-Actual days"`                                                                                  // 0-30 1-Actual days ",
	YearDays             string          `json:"yearDays" description:"Year days:0-360 1-365 2-Actual days"`                                                                             // 0-360 1-365 2-Actual days ",
	DecimalFlag          string          `json:"decimalFlag" description:"Decimal flag:0-no 1-yes"`                                                                                      // 0-no 1-yes ",
	AccrualPeriodType    string          `json:"accrualPeriodType" description:"Accrual period type:01-day 02-week 03-month 04-odd month 05-even month 06-quarter 07-half year 08-year"` // 01-day 02-week 03-month 04-odd month 05-even month 06-quarter 07-half year 08-year ",
	AccrualPeriodValue   int             `json:"accrualPeriodValue" description:"Accrual period value"`                                                                                  // ",
	AccrualPeriodDay     int             `json:"accrualPeriodDay" description:"Accrual period day"`                                                                                      // ",
	SettlePeriodType     string          `json:"settlePeriodType" description:"Settle period type:01-day 02-week 03-month 04-odd month 05-even month 06-quarter 07-half year 08-year"`   // 01-day 02-week 03-month 04-odd month 05-even month 06-quarter 07-half year 08-year ",
	SettlePeriodValue    int             `json:"settlePeriodValue" description:"Settle period value"`                                                                                    // ",
	SettlePeriodDay      int             `json:"settlePeriodDay" description:"Settle period day"`                                                                                        // ",
	Status               string          `json:"status" description:"Status"`                                                                                                            // ",
	TaskOrgzNum          string          `json:"taskOrgzNum" description:"Task orgz number"`                                                                                             // ",
	TaskUserId           string          `json:"taskUserId" description:"Task user id"`                                                                                                  // ",
	TaskUserName         string          `json:"taskUserName" description:"Task user name"`                                                                                              // ",
	CreateTime           string          `json:"createTime" description:"Create time"`                                                                                                   // ",
	//FinalModifyDate      string          `json:"finalModifyDate"`      //    ",
	//FinalModifyTime      string          `json:"finalModifyTime"`      //    ",
	//FinalModifyOrgz      string          `json:"finalModifyOrgz"`      //    ",
	//FinalModifyUser      string          `json:"finalModifyUser"`      //    ",
	//TccStatus            int             `json:"tccStatus"`            //    ",
}

type TLoanFeePeriod struct {
	AccountingAccountNum string          `json:"accountingAccountNum" description:"Accounting account number"` // ",
	Period               int             `json:"period"`
	FeeType              string          `json:"feeType"`
	FeeAccountStatus     string          `json:"feeAccountStatus"`
	FeeAmount            decimal.Decimal `json:"feeAmount"`
	FeeWaived            decimal.Decimal `json:"feeWaived"`
	FeePaid              decimal.Decimal `json:"FeePaid"`
}

func (*AL000001I) GetServiceKey() string {
	return "AL000001"
}
