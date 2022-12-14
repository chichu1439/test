package models

import (
	"github.com/shopspring/decimal"
)

type IC000018I struct {
	LoanAmount                 decimal.Decimal `json:"loanAmount" validate:"required" description:"Loan amount"`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             //贷款金额
	InterestCalculateStartDate string          `json:"interestCalculateStartDate" description:"Interest calculate start date" validate:"required"`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           //计算开始日期
	InterestCalculateEndDate   string          `json:"interestCalculateEndDate" description:"Interest calculate end date"`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   //计算结束日期
	LoanCycleCode              string          `json:"loanCycleCode" description:"Repayment cycle frequency:01-daily 03-fortnightly 04-monthly 06-quarterly 08-yearly"`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      //还款周期频率 01-daily 03-fortnightly 04-monthly 06-quarterly 08-yearly
	InterestRate               decimal.Decimal `json:"interestRate" validate:"required" description:"interest rate"`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         //利率
	RepayMethod                string          `json:"repayMethod" validate:"required" description:"Repay method:01-equal principal and interest 02-equal principal 03-due (one-time) repayment of principal and interest (interest shall be paid with principal) (one-time repayment of principal and interest upon maturity) 04-due repayment of principal and interest (installment payment of interest, due repayment of principal (interest first and then principal)) 05-equal principal and interest (the repayment amount of principal and interest in each period is equal, and the interest bearing principal of each period is the total principal of the loan)"` //01-等额本息 02-等额本金 03-到期（一次性）还本付息（息随本清)(到期一次性还本还息） 04-到期还本周期还息(分期付息到期还本（先息后本)) 05-等本等息（每期还本还息还款额都相等，每期计息的本金为贷款总本金）
	PeriodNum                  int             `json:"periodNum" description:"Number of periods"`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            //期数
	RepayDay                   string          `json:"repayDay" validate:"required" description:"Repayment date"`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            //还款日
	DaysOfYear                 int             `json:"daysOfYear" description:"Days of year" validate:"required"`
	DaysOfMonth                int             `json:"daysOfMonth" description:"Days of month"`
	NextRepayDate              string          `json:"nextRepayDate" description:"Next repay date"`
	BillDay                    string          `json:"billDay" description:"bill day"`
	IsSameFlag                 bool            `json:"isSameFlag"`
}

type IC000018O struct {
	RepayMethod                string            `json:"repayMethod" validate:"required" description:"Repay method:01-equal principal and interest 02-equal principal 03-due (one-time) repayment of principal and interest (interest shall be paid with principal) (one-time repayment of principal and interest upon maturity) 04-due repayment of principal and interest (installment payment of interest, due repayment of principal (interest first and then principal)) 05-equal principal and interest (the repayment amount of principal and interest in each period is equal, and the interest bearing principal of each period is the total principal of the loan)"` //01-等额本息 02-等额本金 03-到期（一次性）还本付息（息随本清)(到期一次性还本还息） 04-到期还本周期还息(分期付息到期还本（先息后本)) 05-等本等息（每期还本还息还款额都相等，每期计息的本金为贷款总本金）
	InterestCalculateStartDate string            `json:"interestCalculateStartDate" description:"Interest calculate start date"`
	InterestCalculateEndDate   string            `json:"interestCalculateEndDate" description:"Interest calculate end date"`
	PeriodNum                  int               `json:"periodNum" description:"Period number"`
	LoanAmount                 decimal.Decimal   `json:"loanAmount" description:"Loan amount"`
	PlanRepayTotalAmount       decimal.Decimal   `json:"planRepayTotalAmount" description:"Plan repay total amount"`
	PlanRepayTotalInterest     decimal.Decimal   `json:"planRepayTotalInterest" description:"Plan repay total interest"`
	InterestRate               decimal.Decimal   `json:"interestRate" description:"Interest rate"`
	FirstRepayDate             string            `json:"firstRepayDate"`
	FirstBillDate              string            `json:"firstBillDate"`
	Records                    []RepayPlanRecord `json:"records" description:"Records"`
}

type RepayPlanRecord struct {
	PeriodNum                  int             `json:"periodNum" description:"Period number"`
	InterestCalculateStartDate string          `json:"interestCalculateStartDate" description:"Interest calculate start date"`
	InterestCalculateEndDate   string          `json:"interestCalculateEndDate" description:"Interest calculate end date"`
	PlanRepayDate              string          `json:"planRepayDate" description:"Plan repay date"`
	PlanBillDate               string          `json:"planBillDate" description:"Plan bill date"`
	PlanRepayTotalAmount       decimal.Decimal `json:"planRepayTotalAmount" description:"Plan repay total amount"`
	PlanRepayPrinciple         decimal.Decimal `json:"planRepayPrinciple" description:"Plan repay principle"`
	PlanRepayInterest          decimal.Decimal `json:"planRepayInterest" description:"Plan repay interest"`
	MaintainPrinciple          decimal.Decimal `json:"maintainPrinciple" description:"Maintain principle"`
	DaysOfInterestCalculate    int             `json:"daysOfInterestCalculate" description:"Days of interest calculate"`
}

func (*IC000018I) GetServiceKey() string {
	return "IC000018"
}
