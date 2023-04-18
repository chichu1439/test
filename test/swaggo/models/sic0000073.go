package models

import "github.com/shopspring/decimal"

type IC000073Request struct {
	LoanReceiptNum string `json:"loanReceiptNum" description:"Loan receipt number" validate:"required"`
	CustomerId     string `json:"customerId" description:"Customer ID"`
}

type IC000073Response struct {
	ContractDetail    *ContractDetail     `json:"contractDetail" description:"Contract detail"`
	CustomerInfo      *CustomerInfo       `json:"customerInfo" description:"Customer info"`
	ProductInfo       *ProductInfo        `json:"productInfo" description:"Product info"`
	AccountInfo       *AccountInfo        `json:"accountInfo" description:"Account info"`
	ListInterest      []MicroLoanInterest `json:"listInterest" description:"List interest"`
	RepaymentPlanList []RepaymentPlan     `json:"repaymentPlanList" description:"Repayment plan"`
}

type ContractDetail struct {
	LoanReceiptNum                                   string          `json:"loanReceiptNum" description:"Loan receipt number"`
	AccountingAccountNum                             string          `json:"accountingAccountNum" description:"Accounting account number"`
	ReceiptStatus                                    string          `json:"receiptStatus" description:"Receipt status"`
	CustomerId                                       string          `json:"customerId" description:"Customer ID"`
	ProductId                                        string          `json:"productId" description:"Product ID"`
	PartialDisbursement                              string          `json:"partialDisbursement" description:"Partial disbursement"`
	MakeLoanOrgzNum                                  string          `json:"makeLoanOrgzNum" description:"Make loan organization number"`
	ChargeCompoundFlag                               string          `json:"chargeCompoundFlag" description:"Charge compound flag"`
	LoanMethod                                       string          `json:"loanMethod" description:"Loan method"`
	LoanCycle                                        string          `json:"loanCycle" description:"Loan cycle"`
	LoanPeriod                                       int             `json:"loanPeriod" description:"Loan period"`
	PermitAdvanceRepayTime                           int             `json:"permitAdvanceRepayTime" description:"Permit advance repay time"`
	AdvanceRepayLimitStartDate                       string          `json:"advanceRepayLimitStartDate" description:"Advance repay limit start date"`
	AdvanceRepayLimitDay                             int             `json:"advanceRepayLimitDay" description:"Advance Repay limit day"`
	LimitTermInsidePermitPartRepayFlag               string          `json:"limitTermInsidePermitPartRepayFlag" description:"Limit term inside permit part repay flag"`
	LimitTermInsidePermitPaymentOffFlag              string          `json:"limitTermInsidePermitPaymentOffFlag" description:"Limit term inside permit payment off flag"`
	OpenAccountDate                                  string          `json:"openAccountDate" description:"Open account date"`
	FirstDisburseDate                                string          `json:"firstDisburseDate" description:"First disburse date"`
	StartDate                                        string          `json:"startDate" description:"Start date"`
	Currency                                         string          `json:"currency" description:"Currency"`
	LoanAmount                                       decimal.Decimal `json:"loanAmount" description:"Loan amount"`
	AdvanceRepayTimes                                int             `json:"advanceRepayTimes" description:"Advance repay times"`
	RepayCycle                                       string          `json:"repayCycle" description:"Repay cycle"`
	GraceTermInterestFlag                            string          `json:"graceTermInterestFlag" description:"Grace term interest flag"`
	MaxGraceTermDay                                  int             `json:"maxGraceTermDay" description:"Max grace term day"`
	GraceTotalDay                                    int             `json:"graceTotalDay" description:"Grace total day"`
	CurrentPeriod                                    int             `json:"currentPeriod" description:"Current period"`
	ExtendFlag                                       string          `json:"extendFlag" description:"Extend flag"`
	ExtendTime                                       int             `json:"extendTime" description:"Extend time"`
	MatureDate                                       string          `json:"matureDate" description:"Mature date"`
	ClearDate                                        string          `json:"clearDate" description:"Clear date"`
	LoanTotalPeriod                                  int             `json:"loanTotalPeriod" description:"Loan total period"`
	CurrentExecutePeriod                             int             `json:"currentExecutePeriod" description:"Current execute period"`
	LoanMaintainPeriod                               int             `json:"loanMaintainPeriod" description:"Loan maintain period"`
	PrincipleOweStartDate                            string          `json:"principleOweStartDate" description:"Principle owe start date"`
	InterestOverdueStartDate                         string          `json:"interestOverdueStartDate" description:"Interest overdue start date"`
	AccumulateOverduePeriod                          int             `json:"accumulateOverduePeriod" description:"Accumulate overdue period"`
	AccumulateWriteOffAmount                         decimal.Decimal `json:"accumulateWriteOffAmount" description:"Accumulate write off amount"`
	LastTermEstimateLoanRiskClassifiedDate           string          `json:"lastTermEstimateLoanRiskClassifiedDate" description:"Last term estimate loan risk classified date"`
	LastTermEstimateLoanRiskClassifiedCode           string          `json:"lastTermEstimateLoanRiskClassifiedCode" description:"Last term estimate loan risk classified code"`
	CurrentEstimateLoanRiskClassifiedDate            string          `json:"currentEstimateLoanRiskClassifiedDate" description:"Current estimate loan risk classified date"`
	CurrentEstimateLoanRiskClassifiedCode            string          `json:"currentEstimateLoanRiskClassifiedCode" description:"Current estimate loan risk classified code"`
	LastTermArtificialIdentifyLoanRiskClassifiedDate string          `json:"lastTermArtificialIdentifyLoanRiskClassifiedDate" description:"Last term artificial identify loan risk classified date"`
	LastTermArtificialIdentifyLoanRiskClassifiedCode string          `json:"lastTermArtificialIdentifyLoanRiskClassifiedCode" description:"Last term artificial identify loan risk classified code"`
	CurrentArtificialIdentifyLoanRiskClassifiedDate  string          `json:"currentArtificialIdentifyLoanRiskClassifiedDate" description:"Current artificial identify loan risk classified date"`
	CurrentArtificialIdentifyLoanRiskClassifiedCode  string          `json:"currentArtificialIdentifyLoanRiskClassifiedCode" description:"Current artificial identify loan risk classified code"`
	LoanPurpose                                      string          `json:"loanPurpose"`
	HasRestructuring                                 string          `json:"hasRestructuring"`
	LastRestructuringDate                            string          `json:"lastRestructuringDate"`
	RepayMethod                                      string          `json:"repayMethod"`
}

type CustomerInfo struct {
	CustomerName   string `json:"customerName" description:"Customer name"`
	IdNation       string `json:"idNation" description:"ID nation"`
	IdType         string `json:"idType" description:"ID type"`
	IdNum          string `json:"idNum" description:"ID number"`
	IdExpireDate   string `json:"idExpireDate" description:"ID expire date"`
	CustomerStatus string `json:"customerStatus" description:"Customer status"`
	Marriage       string `json:"marriage" description:"Marriage"`
	CustomerGrade  string `json:"customerGrade" description:"Customer grade"`
}

type AccountInfo2 struct {
	AccountingAccountNum    string          `json:"accountingAccountNum" description:"Accounting account number"`
	LoanReceiptNum          string          `json:"loanReceiptNum" description:"Loan receipt number"`
	AccountingAccountStatus string          `json:"accountingAccountStatus" description:"Accounting account status"`
	ProductId               string          `json:"productId" description:"Product ID"`
	CustomerId              string          `json:"customerId" description:"Customer ID"`
	Currency                string          `json:"currency" description:"Currency"`
	LoanAmount              decimal.Decimal `json:"loanAmount" description:"Loan amount"`
	Balance                 decimal.Decimal `json:"balance" description:"Balance"`
	RepayCycle              string          `json:"repayCycle" description:"Repay cycle"`
	RepayFrequency          int             `json:"repayFrequency" description:"Repay frequency"`
	RepayMethod             string          `json:"repayMethod" description:"Repay method"`
	NextRepaymentDate       string          `json:"nextRepaymentDate" description:"Next repayment date"`
	TotalPeriod             int             `json:"totalPeriod" description:"Total period"`
	RepayDay                string          `json:"repayDay" description:"Repay day"`
	CurrentPeriod           int             `json:"currentPeriod" description:"Current period"`
	DrawTotalAmount         decimal.Decimal `json:"drawTotalAmount" description:"Draw total amount"`
	DrawBeginDate           string          `json:"drawBeginDate" description:"Draw begin date"`
	DrawTotalPeriod         int             `json:"drawTotalPeriod" description:"Draw total period"`
	MaturityDate            string          `json:"maturityDate" description:"Maturity date"`
	OpenAccountDate         string          `json:"openAccountDate" description:"Open account date"`
	OpenAccountOrgz         string          `json:"openAccountOrgz" description:"Open account organization"`
	GracePeriod             int             `json:"gracePeriod" description:"Grace period"`
}

type RepaymentPlan struct {
	AccountingAccountNum   string          `json:"accountingAccountNum" description:"Accounting account number"`
	Period                 int             `json:"period" description:"Period Number"`
	CurrentAccountStatus   string          `json:"currentAccountStatus" description:"Current account status"`
	PeriodRepayStatus      string          `json:"periodRepayStatus" description:"Period repay status"`
	PrincipalRepayStatus   string          `json:"principalRepayStatus" description:"Principal repay status"`
	FeeRepayStatus         string          `json:"feeRepayStatus" description:"Fee repay status"`
	CurrentPeriodBeginDate string          `json:"currentPeriodBeginDate" description:"Current period begin date"`
	CurrentPeriodEndDate   string          `json:"currentPeriodEndDate" description:"Current period end date"`
	CurrentRepayDate       string          `json:"currentRepayDate" description:"Current repay date"`
	Currency               string          `json:"currency" description:"Currency"`
	Installment            decimal.Decimal `json:"installment" description:"Plan Repayment Principal"`
	PlanRepayPrincipal     decimal.Decimal `json:"planRepayPrincipal" description:"Plan Repayment Principal"`
	PlanRepayInterest      decimal.Decimal `json:"planRepayInterest" description:"Plan Repayment Interest"`
	ActualRepayDate        string          `json:"actualRepayDate" description:"Actual repay date"`
	ActualRepayPrincipal   decimal.Decimal `json:"actualRepayPrincipal" description:"Actual repay principal"`
	ActualRepayInterest    decimal.Decimal `json:"actualRepayInterest" description:"Actual repay interest"`
	InterestWaived         decimal.Decimal `json:"interestWaived" description:"Interest waived"`
	PrincipalWaived        decimal.Decimal `json:"principalWaived" description:"Principal waived"`
	PrincipalImpaired      decimal.Decimal `json:"principalImpaired" description:"Principal impaired"`
	PrincipalWriteoff      decimal.Decimal `json:"principalWriteoff" description:"Principal writeoff"`
	FeeAmount              decimal.Decimal `json:"feeAmount" description:"Fee amount"`
	FeeRepay               decimal.Decimal `json:"feeRepay" description:"Fee repay"`
	FeeWaived              decimal.Decimal `json:"feeWaived" description:"Fee waived"`
	OverdueStartDate       string          `json:"overdueStartDate" description:"Overdue start date"`
	OverdueDays            int             `json:"overdueDays" description:"Overdue days"`
	CloseDate              string          `json:"closeDate" description:"Close date"`
	CurrentUnpaidPrincipal decimal.Decimal `json:"currentUnpaidPrincipal" description:"Current unpaid principal"`
	MaintainPrinciple      decimal.Decimal `json:"maintainPrinciple" description:"Maintain Principle"`
	LastActiveDate         string          `json:"lastActiveDate" description:"Last active date"`
	InterestWriteoff       decimal.Decimal `json:"interestWriteoff" description:"Interest writeoff"`
}

type ProductInfo2 struct {
	ProductId     string `json:"productId" description:"Product ID"`
	ProductType   string `json:"productType" description:"Product type"`
	CustomerType  string `json:"customerType" description:"Customer type"`
	ProductName   string `json:"productName" description:"Product name"`
	Currency      string `json:"currency" description:"Currency"`
	Version       string `json:"version" description:"Version"`
	Remark        string `json:"remark" description:"Remark"`
	EffectiveDate string `json:"effectiveDate" description:"Effective date"`
	ExpireDate    string `json:"expireDate" description:"Expire date"`
}

func (*IC000073Request) GetServiceKey() string {
	return "IC000073"
}
