package models

import "github.com/shopspring/decimal"

type IC000069Request struct {
	CustomerId   string          `json:"customerId" description:"Customer id" validate:"required"`
	ProductId    string          `json:"productId" description:"Product id" validate:"required"`
	LoanAmount   decimal.Decimal `json:"loanAmount" description:"Loan Amount" validate:"required"`
	PeriodNum    int             `json:"periodNum" description:"Period Num" validate:"required"`
	RepayCycle   string          `json:"repayCycle" description:"Repay Cycle" validate:"required"`
	RepayMethod  string          `json:"repayMethod" description:"Repay Method" validate:"required"`
	GracePeriod  int             `json:"gracePeriod" description:"Grace Period"`
	LoanPurpose  string          `json:"loanPurpose" description:"Loan Purpose"`
	AppDate      string          `json:"appDate" description:"App Date"`
	Maker        string          `json:"maker" description:"Maker"`
	MakerComment string          `json:"maker Comment" description:"Maker Comment"`
}
type IC000069Response struct {
	ContractDetail    *ContractDetail `json:"contractDetail" description:"Contract Detail"`
	CustomerInfo      *CustomerInfo   `json:"customerInfo" description:"Customer Info"`
	ProductInfo       *ProductInfo    `json:"productInfo " description:"Product Info "`
	AccountInfo       *AccountInfo    `json:"accountInfo" description:"Account Info"`
	RepaymentPlanList []RepaymentPlan `json:"repaymentPlanList" description:"Repayment Plan List"`
}
type SIC8800035O struct {
	LoanReceiptNum                                   string          `json:"loanReceiptNum"`
	AccountingAccountNum                             string          `json:"accountingAccountNum"`
	ReceiptStatus                                    string          `json:"receiptStatus"`
	CustomerId                                       string          `json:"customerId"`
	ProductNum                                       string          `json:"productNum"`
	PartialDisbursement                              string          `json:"partialDisbursement"`
	MakeLoanOrgzNum                                  string          `json:"makeLoanOrgzNum"`
	ChargeCompoundFlag                               string          `json:"chargeCompoundFlag"`
	LoanMethod                                       string          `json:"loanMethod"`
	RepayMethod                                      string          `json:"repayMethod"`
	LoanCycle                                        string          `json:"loanCycle"`
	LoanPeriod                                       int             `json:"loanPeriod"`
	PermitAdvanceRepayTime                           int             `json:"permitAdvanceRepayTime"`
	AdvanceRepayLimitStartDate                       string          `json:"advanceRepayLimitStartDate"`
	AdvanceRepayLimitDay                             int             `json:"advanceRepayLimitDay"`
	LimitTermInsidePermitPartRepayFlag               string          `json:"limitTermInsidePermitPartRepayFlag"`
	LimitTermInsidePermitPaymentOffFlag              string          `json:"limitTermInsidePermitPaymentOffFlag"`
	OpenAccountDate                                  string          `json:"openAccountDate"`
	FirstDisburseDate                                string          `json:"firstDisburseDate"`
	StartDate                                        string          `json:"startDate"`
	Currency                                         string          `json:"currency"`
	LoanAmount                                       decimal.Decimal `json:"loanAmount"`
	AdvanceRepayTimes                                int             `json:"advanceRepayTimes"`
	RepayCycle                                       string          `json:"repayCycle"`
	GraceTermInterestFlag                            string          `json:"graceTermInterestFlag"`
	MaxGraceTermDay                                  int             `json:"maxGraceTermDay"`
	GraceTotalDay                                    int             `json:"graceTotalDay"`
	CurrentPeriod                                    int             `json:"currentPeriod"`
	ExtendFlag                                       string          `json:"extendFlag"`
	ExtendTime                                       int             `json:"extendTime"`
	MatureDate                                       string          `json:"matureDate"`
	ClearDate                                        string          `json:"clearDate"`
	LoanTotalPeriod                                  int             `json:"loanTotalPeriod"`
	CurrentExecutePeriod                             int             `json:"currentExecutePeriod"`
	LoanMaintainPeriod                               int             `json:"loanMaintainPeriod"`
	PrincipleOweStartDate                            string          `json:"principleOweStartDate"`
	InterestOverdueStartDate                         string          `json:"interestOverdueStartDate"`
	AccumulateOverduePeriod                          int             `json:"accumulateOverduePeriod"`
	AccumulateWriteOffAmount                         decimal.Decimal `json:"accumulateWriteOffAmount"`
	LastTermEstimateLoanRiskClassifiedDate           string          `json:"lastTermEstimateLoanRiskClassifiedDate"`
	LastTermEstimateLoanRiskClassifiedCode           string          `json:"lastTermEstimateLoanRiskClassifiedCode"`
	CurrentEstimateLoanRiskClassifiedDate            string          `json:"currentEstimateLoanRiskClassifiedDate"`
	CurrentEstimateLoanRiskClassifiedCode            string          `json:"currentEstimateLoanRiskClassifiedCode"`
	LastTermArtificialIdentifyLoanRiskClassifiedDate string          `json:"lastTermArtificialIdentifyLoanRiskClassifiedDate"`
	LastTermArtificialIdentifyLoanRiskClassifiedCode string          `json:"lastTermArtificialIdentifyLoanRiskClassifiedCode"`
	CurrentArtificialIdentifyLoanRiskClassifiedDate  string          `json:"currentArtificialIdentifyLoanRiskClassifiedDate"`
	CurrentArtificialIdentifyLoanRiskClassifiedCode  string          `json:"currentArtificialIdentifyLoanRiskClassifiedCode"`
}

func (*IC000069Request) GetServiceKey() string {
	return "IC000069"
}
