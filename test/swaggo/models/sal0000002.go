package models

type AL000002I struct {
	AccountingAccountNum        string                          `json:"accountingAccountNum" description:"Accounting account number"`
	TLoanAcctgAcct              *TLoanAcctgAcctStr              `json:"tLoanAcctgAcct" description:"T loan acctg acct"`
	TAcctLoanAcctgPeriods       []TAcctLoanAcctgPeriodStr       `json:"tAcctLoanAcctgPeriods" description:"T acct loan acctg periods"`
	TLoanInterestAccrual        *TLoanInterestAccrualStr        `json:"tLoanInterestAccrual" description:"T loan interest accrual"`
	TLoanInterestAccrualPeriods []TLoanInterestAccrualPeriodStr `json:"tLoanInterestAccrualPeriods" description:"T loan interest accrual periods"`
}

type AL000002O struct {
}

type TLoanAcctgAcctStr struct {
	AccountingAccountNum    string `json:"accountingAccountNum" description:"Accounting account number"`                                                       // ",
	LoanReceiptNum          string `json:"loanReceiptNum" description:"Loan receipt number"`                                                                   // ",
	AccountingAccountStatus string `json:"accountingAccountStatus" description:"Accounting account status:0:normal 1:overdue 2:impaired 3:write_off 4:closed"` // 0-正常;1-逾期;2-减值;3-核销;4-结清 ",
	ProductId               string `json:"productId" description:"Product id"`                                                                                 // ",
	CustomerId              string `json:"customerId" description:"Customer id"`                                                                               // ",
	Currency                string `json:"currency" description:"Currency"`                                                                                    // CM0001 ",
	RepayType               string `json:"repayType" description:"Repay type"`                                                                                 // 01-等额本金 02-等额本息 03-到期（一次性）还本付息 04-到期还本周期还息 05-等本等息 06-随借随还 ",
	LoanAmount              string `json:"loanAmount" description:"Loan amount"`                                                                               // ",
	Balance                 string `json:"balance" description:"Balance"`                                                                                      // ",
	RepayCycle              string `json:"repayCycle" description:"Repay cycle:01- daily 02- weekly 04- Monthly 08- Yearly"`                                   // 01-天 02-周 03-半月/双周 04-月 05-双月 06-季 07-半年 08-年 ",
	RepayFrequency          string `json:"repayFrequency" description:"Repay frequency"`                                                                       // ",
	RepayMethod             string `json:"repayMethod" description:"Repay method"`                                                                             // 01-等额本金 02-等额本息 03-到期（一次性）还本付息 04-到期还本周期还息 05-等本等息 06-随借随还 ",
	NextRepaymentDate       string `json:"nextRepaymentDate" description:"Next repayment date"`                                                                // ",
	TotalPeriod             string `json:"totalPeriod" description:"Total period"`                                                                             // ",
	CurrentPeriod           string `json:"currentPeriod" description:"Current period"`                                                                         // ",
	DrawTotalAmount         string `json:"drawTotalAmount" description:"Draw total amount"`                                                                    // ",
	DrawBeginDate           string `json:"drawBeginDate" description:"Draw begin date"`                                                                        // ",
	DrawTotalPeriod         string `json:"drawTotalPeriod" description:"Draw total period"`                                                                    // ",
	MaturityDate            string `json:"maturityDate" description:"Maturity date"`                                                                           // ",
	OpenAccountDate         string `json:"openAccountDate" description:"Open account date"`                                                                    // ",
	OpenAccountOrgz         string `json:"openAccountOrgz" description:"Open account organization"`                                                            // ",
	GracePeriod             string `json:"gracePeriod" description:"Grace period"`                                                                             // ",
	AutoRepayFlag           string `json:"autoRepayFlag" description:"Auto repay flag"`                                                                        // 1-Auto repay 2-Manual repay ",
	RepayLocalBankFlag      string `json:"repayLocalBankFlag" description:"Repay local bank flag"`                                                             // I-行内 O-行外 ",
	RepayBankNum            string `json:"repayBankNum" description:"Repay bank number"`                                                                       // ",
	RepayBankName           string `json:"repayBankName" description:"Repay bank name"`                                                                        // ",
	RepayAccountNum         string `json:"repayAccountNum" description:"Repay account number"`                                                                 // ",
	RepayAccountName        string `json:"repayAccountName" description:"Repay account name"`                                                                  // ",
	BadLoanStatus           string `json:"badLoanStatus" description:"Bad loan status"`                                                                        // 0-正常;1-不良 ",
	CompoundInterestFlag    string `json:"compoundInterestFlag" description:"Compound interest flag"`                                                          // 0-No 1-Yes ",
	LastActiveDate          string `json:"lastActiveDate" description:"Last active date"`                                                                      // ",
}

type TAcctLoanAcctgPeriodStr struct {
	AccountingAccountNum   string `json:"accountingAccountNum" description:"Accounting account number"`   // ",
	Period                 int    `json:"period" description:"Period"`                                    // ",
	CurrentAccountStatus   string `json:"currentAccountStatus" description:"Current account status"`      // 0-正常;1-逾期;2-减值;3-核销;4-结清 ",
	PeriodRepayStatus      string `json:"periodRepayStatus" description:"Period repay status"`            // O-Unpaid 1-Partial Paid 2-Total Paid 3-Waived ",
	PrincipalRepayStatus   string `json:"principalRepayStatus" description:"Principal repay status"`      // O-Unpaid 1-Partial Paid 2-Total Paid 3-Waived ",
	FeeRepayStatus         string `json:"feeRepayStatus" description:"Fee repay status"`                  // O-Unpaid 1-Partial Paid 2-Total Paid 3-Waived ",
	CurrentPeriodBeginDate string `json:"currentPeriodBeginDate" description:"Current period begin date"` // ",
	CurrentPeriodEndDate   string `json:"currentPeriodEndDate" description:"Current period end date"`     // ",
	CurrentRepayDate       string `json:"currentRepayDate" description:"Current repay date"`              // ",
	Currency               string `json:"currency" description:"Currency"`                                // CM0001 ",
	PlanRepayPrincipal     string `json:"planRepayPrincipal" description:"Plan repay principal"`          // ",
	PlanRepayInterest      string `json:"planRepayInterest" description:"Plan repay interest"`            // ",
	ActualRepayDate        string `json:"actualRepayDate" description:"Actual repay date"`                // ",
	ActualRepayPrincipal   string `json:"actualRepayPrincipal" description:"Actual repay principal"`      // ",
	ActualRepayInterest    string `json:"actualRepayInterest" description:"Actual repay interest"`        // ",
	InterestWaived         string `json:"interestWaived" description:"Interest waived"`                   // ",
	PrincipalWaived        string `json:"principalWaived" description:"Principal waived"`                 // ",
	PrincipalImpaired      string `json:"principalImpaired" description:"Principal impaired"`             // ",
	PrincipalWriteoff      string `json:"principalWriteoff" description:"Principal writeoff"`             // ",
	FeeAmount              string `json:"feeAmount" description:"Fee amount"`                             // ",
	FeeRepay               string `json:"feeRepay" description:"Fee repay"`                               // ",
	FeeWaived              string `json:"feeWaived" description:"Fee waived"`                             // ",
	OverdueStartDate       string `json:"overdueStartDate" description:"Overdue start date"`              // ",
	OverdueDays            string `json:"overdueDays" description:"Overdue days"`                         // ",
	CloseDate              string `json:"closeDate" description:"Close date"`                             // ",
	CurrentUnpaidPrincipal string `json:"currentUnpaidPrincipal" description:"Current unpaid principal"`  // ",
	MaintainPrinciple      string `json:"maintainPrinciple" description:"Maintain principle"`             // ",
	LastActiveDate         string `json:"lastActiveDate" description:"Last active date"`                  // ",
	InterestWriteoff       string `json:"interestWriteoff" description:"Interest writeoff"`               // ",
}

type TLoanInterestAccrualStr struct {
	AccountingAccountNum   string `json:"accountingAccountNum" description:"Accounting account number"` // ",
	Currency               string `json:"currency" description:"Currency"`                              // ",
	InterestAccountStatus  string `json:"interestAccountStatus" description:"Interest account status"`  // 0-正常;1-逾期;2-减值;3-核销;4-结清 ",
	InterestBeginDate      string `json:"interestBeginDate" description:"Interest begin date"`          // ",
	InterestEndDate        string `json:"interestEndDate" description:"Interest end date"`              // ",
	LastAccrualDate        string `json:"lastAccrualDate" description:"Last accrual date"`              // ",
	LastCalcDate           string `json:"lastCalcDate" description:"Last calc date"`
	NextAccrualDate        string `json:"nextAccrualDate" description:"Next accrual date"`             // ",
	LastSettlementDate     string `json:"lastSettlementDate" description:"Last settlement date"`       // ",
	NextSettlementDate     string `json:"nextSettlementDate" description:"Next settlement date"`       // ",
	InterestRepaymentDate  string `json:"interestRepaymentDate" description:"Interest repayment date"` // ",
	InterestPaid           string `json:"interestPaid" description:"Interest paid"`                    // ",
	InterestWaived         string `json:"interestWaived" description:"Interest waived"`                // ",
	InterestImpaired       string `json:"interestImpaired" description:"Interest impaired"`            // ",
	InterestWriteoff       string `json:"interestWriteoff" description:"Interest writeoff"`            // ",
	InterestAfterWriteoff  string `json:"interestAfterWriteoff" description:"Interest after writeoff"` // ",
	InterestAccrual        string `json:"interestAccrual" description:"Interest accrual"`              // ",
	InterestCalc           string `json:"interestCalc" description:"Interest calc"`
	InterestSettlement     string `json:"interestSettlement" description:"Interest settlement"` // ",
	PenaltyInterestAccrual string `json:"penaltyInterestAccrual"`
	PenaltyInterestCalc    string `json:"penaltyInterestCalc"`
	PenaltyInterestPaid    string `json:"penaltyInterestPaid"`
}

type TLoanInterestAccrualPeriodStr struct {
	AccountingAccountNum  string `json:"accountingAccountNum" description:"Accounting account number"` // ",
	Period                int    `json:"period" description:"Period"`                                  // ",
	InterestType          string `json:"interestType" description:"Interest type"`                     // O-Normal Loan 1-Overdue Loan ",
	CompoundFlag          string `json:"compoundFlag" description:"Compound flag"`                     // 0-No 1-Yes ",
	Currency              string `json:"currency" description:"Currency"`                              // ",
	InterestAccountStatus string `json:"interestAccountStatus" description:"Interest account status"`  // 0-正常;1-逾期;2-减值;3-核销;4-结清 ",
	InterestRepayStatus   string `json:"interestRepayStatus" description:"Interest repay status"`      // O-Unpaid 1-Partial Paid 2-Total Paid 3-Waived ",
	InterestBeginDate     string `json:"interestBeginDate" description:"Interest begin date"`          // ",
	InterestEndDate       string `json:"interestEndDate" description:"Interest end date"`              // ",
	LastAccrualDate       string `json:"lastAccrualDate" description:"Last accrual date"`              // ",
	LastCalcDate          string `json:"lastCalcDate" description:"Last calc date"`
	NextAccrualDate       string `json:"nextAccrualDate" description:"Next accrual date"`             // ",
	LastSettlementDate    string `json:"lastSettlementDate" description:"Last settlement date"`       // ",
	NextSettlementDate    string `json:"nextSettlementDate" description:"Next settlement date"`       // ",
	InterestRepaymentDate string `json:"interestRepaymentDate" description:"Interest repayment date"` // ",
	InterestPaid          string `json:"interestPaid" description:"Interest paid"`                    // ",
	InterestWaived        string `json:"interestWaived" description:"Interest waived"`                // ",
	InterestImpaired      string `json:"interestImpaired" description:"Interest impaired"`            // ",
	InterestWriteoff      string `json:"interestWriteoff" description:"Interest writeoff"`            // ",
	InterestAfterWriteoff string `json:"interestAfterWriteoff" description:"Interest after writeoff"` // ",
	InterestAccrual       string `json:"interestAccrual" description:"Interest accrual"`              // ",
	InterestCalc          string `json:"interestCalc" description:"Interest calc"`
	InterestOverdue       string `json:"interestOverdue" description:"Interest overdue"`
	InterestSettlement    string `json:"interestSettlement" description:"Interest settlement"` // ",
}

func (*AL000002I) GetServiceKey() string {
	return "AL000002"
}
