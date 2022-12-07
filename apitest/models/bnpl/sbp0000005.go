//Version: v1
package models

import "github.com/shopspring/decimal"

type BP000005Request struct {
	LoanReceiptNum string `json:"loanReceiptNum" description:"Loan receipt number" validate:"required"`
}

type BP000005Response struct {
	LoanInfo          LoanInfo        `json:"loanInfo" description:"Loan info"`
	RepaymentPlanList []RepaymentPlan `json:"repaymentPlanList" description:"Repayment plan list"`
}

type LoanInfo struct {
	LoanReceiptNum                                   string          `json:"loanReceiptNum" description:"Loan Receipt Number"`
	AccountingAccountNum                             string          `json:"accountingAccountNum" description:"Accounting Account Number"`
	ReceiptStatus                                    string          `json:"receiptStatus" description:"Receipt Status"`
	CustomerId                                       string          `json:"customerId" description:"Customer Id"`
	ProductId                                        string          `json:"productId" description:"Product Id"`
	PartialDisbursement                              string          `json:"partialDisbursement" description:"Partial Disbursement"`
	MakeLoanOrgzNum                                  string          `json:"makeLoanOrgzNum" description:"Make Loan Orgz Num"`
	ChargeCompoundFlag                               string          `json:"chargeCompoundFlag" description:"Charge Compound Flag"`
	LoanMethod                                       string          `json:"loanMethod" description:"Loan Method"`
	RepayMethod                                      string          `json:"repayMethod" description:"Repay Method"`
	LoanCycle                                        string          `json:"loanCycle" description:"Loan Cycle"`
	LoanPeriod                                       int             `json:"loanPeriod" description:"Loan Period"`
	PermitAdvanceRepayTime                           int             `json:"permitAdvanceRepayTime" description:"Permit Advance Repay Time"`
	AdvanceRepayLimitStartDate                       string          `json:"advanceRepayLimitStartDate" description:"Advance Repay Limit Start Date"`
	AdvanceRepayLimitDay                             int             `json:"advanceRepayLimitDay" description:"Advance Repay Limit Day"`
	LimitTermInsidePermitPartRepayFlag               string          `json:"limitTermInsidePermitPartRepayFlag" description:"Limit Term Inside Permit Part Repay Flag"`
	LimitTermInsidePermitPaymentOffFlag              string          `json:"limitTermInsidePermitPaymentOffFlag" description:"Limit Term Inside Permit Payment Off Flag"`
	OpenAccountDate                                  string          `json:"openAccountDate" description:"Open Account Date"`
	FirstDisburseDate                                string          `json:"firstDisburseDate" description:"First Disburse Date"`
	StartDate                                        string          `json:"startDate" description:"Start Date"`
	Currency                                         string          `json:"currency" description:"Currency"`
	LoanAmount                                       decimal.Decimal `json:"loanAmount" description:"LoanAmount"`
	AdvanceRepayTimes                                int             `json:"advanceRepayTimes" description:"Advance Repay Times"`
	RepayCycle                                       string          `json:"repayCycle" description:"Repay Cycle"`
	GraceTermInterestFlag                            string          `json:"graceTermInterestFlag" description:"Grace Term Interest Flag"`
	MaxGraceTermDay                                  int             `json:"maxGraceTermDay" description:"Max Grace Term Day"`
	GraceTotalDay                                    int             `json:"graceTotalDay" description:"Grace Total Day"`
	CurrentPeriod                                    int             `json:"currentPeriod" description:"Current Period"`
	ExtendFlag                                       string          `json:"extendFlag" description:"Extend Flag"`
	ExtendTime                                       int             `json:"extendTime" description:"Extend Time"`
	MatureDate                                       string          `json:"matureDate" description:"Mature Date"`
	ClearDate                                        string          `json:"clearDate" description:"Clear Date"`
	LoanTotalPeriod                                  int             `json:"loanTotalPeriod" description:"Loan Total Period"`
	CurrentExecutePeriod                             int             `json:"currentExecutePeriod" description:"Current Execute Period"`
	LoanMaintainPeriod                               int             `json:"loanMaintainPeriod" description:"Loan Maintain Period"`
	PrincipleOweStartDate                            string          `json:"principleOweStartDate" description:"Principle Owe Start Date"`
	InterestOverdueStartDate                         string          `json:"interestOverdueStartDate" description:"Interest Overdue Start Date"`
	AccumulateOverduePeriod                          int             `json:"accumulateOverduePeriod" description:"Accumulate Overdue Period"`
	AccumulateWriteOffAmount                         decimal.Decimal `json:"accumulateWriteOffAmount" description:"Accumulate Write Off Amount"`
	LastTermEstimateLoanRiskClassifiedDate           string          `json:"lastTermEstimateLoanRiskClassifiedDate" description:"Last Term Estimate Loan Risk Classified Date"`
	LastTermEstimateLoanRiskClassifiedCode           string          `json:"lastTermEstimateLoanRiskClassifiedCode" description:"Last Term Estimate Loan Risk Classified Code"`
	CurrentEstimateLoanRiskClassifiedDate            string          `json:"currentEstimateLoanRiskClassifiedDate" description:"Current Estimate Loan Risk Classified Date"`
	CurrentEstimateLoanRiskClassifiedCode            string          `json:"currentEstimateLoanRiskClassifiedCode" description:"Current Estimate Loan Risk Classified Code"`
	LastTermArtificialIdentifyLoanRiskClassifiedDate string          `json:"lastTermArtificialIdentifyLoanRiskClassifiedDate" description:"Last Term Artificial Identify Loan Risk Classified Date"`
	LastTermArtificialIdentifyLoanRiskClassifiedCode string          `json:"lastTermArtificialIdentifyLoanRiskClassifiedCode" description:"Last Term Artificial Identify LoanRisk Classified Code"`
	CurrentArtificialIdentifyLoanRiskClassifiedDate  string          `json:"currentArtificialIdentifyLoanRiskClassifiedDate" description:"Current Artificial Identify Loan Risk Classified Date"`
	CurrentArtificialIdentfyLoanRiskClassifiedCode   string          `json:"currentArtificialIdentfyLoanRiskClassifiedCode" description:"Current Artificial Identfy Loan Risk Classified sCode"`
	TotalPeriod                                      int             `json:"totalPeriod" description:"Total Period"`
	PaidPeriod                                       int             `json:"paidPeriod" description:"Paid Period"`
}

type RepaymentPlan struct {
	AccountingAccountNum   string          `json:"accountingAccountNum" description:"Accounting account number"`
	Period                 int             `json:"period" description:"Period"`
	CurrentAccountStatus   string          `json:"currentAccountStatus" description:"Current account status"`
	PeriodRepayStatus      string          `json:"periodRepayStatus" description:"Period repay status"`
	PrincipalRepayStatus   string          `json:"principalRepayStatus" description:"Principal repay status"`
	FeeRepayStatus         string          `json:"feeRepayStatus" description:"Fee repay status"`
	CurrentPeriodBeginDate string          `json:"currentPeriodBeginDate" description:"Current period begin date"`
	CurrentPeriodEndDate   string          `json:"currentPeriodEndDate" description:"Current period end date"`
	CurrentRepayDate       string          `json:"currentRepayDate" description:"Current repay date"`
	Currency               string          `json:"currency" description:"Currency"`
	PlanRepayPrincipal     decimal.Decimal `json:"planRepayPrincipal" description:"Plan repay principal"`
	PlanRepayInterest      decimal.Decimal `json:"planRepayInterest" description:"Plan repay interest"`
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
	MaintainPrinciple      decimal.Decimal `json:"maintainPrinciple" description:"Maintain principle"`
	LastActiveDate         string          `json:"lastActiveDate" description:"Last active date"`
	InterestWriteoff       decimal.Decimal `json:"interestWriteoff" description:"Interest writeoff"`
}

func (*BP000005Request) GetServiceKey() string {
	return "BP000005"
}
