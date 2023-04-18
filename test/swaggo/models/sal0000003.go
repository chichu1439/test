package models

import "github.com/shopspring/decimal"

type AL000003I struct {
	LoanReceiptNum     string          `json:"loanReceiptNum" description:"Loan receipt number"`
	CustomerId         string          `json:"customerId" description:"Customer id"`
	TransAmount        decimal.Decimal `json:"transAmount" description:"Trans amount"`
	Behavior           string          `json:"behavior" description:"01- Repayment,02-Close Repayment"`
	RepaymentStatus    string          `json:"repaymentStatus" description:"01-Total Repayment,02-Partial Repayment,03-Bill Repayment,04-Due Amount Repayment"`
	DiscountPercentage decimal.Decimal `json:"discountPercentage"`
}

type AL000003O struct {
	NeedRepayTotalAmount               decimal.Decimal `json:"needRepayTotalAmount" description:"Need repay total amount"`
	NeedRepayBalance                   decimal.Decimal `json:"needRepayBalance" description:"Need repay balance"`
	NeedRepayTotalPrinciple            decimal.Decimal `json:"needRepayTotalPrinciple" description:"Need repay total principle"`
	NeedRepayImpairmentPrinciple       decimal.Decimal `json:"needRepayImpairmentPrinciple" description:"Need repay impairment principle"`
	NeedRepayOverduePrinciple          decimal.Decimal `json:"needRepayOverduePrinciple" description:"Need repay overdue principle"`
	NeedRepayTotalInterest             decimal.Decimal `json:"needRepayTotalInterest" description:"Need repay total interest"`
	NeedRepayNormalInterest            decimal.Decimal `json:"needRepayNormalInterest" description:"Need repay normal interest"`
	NeedRepayNormalInterestTotal       decimal.Decimal `json:"needRepayNormalInterestTotal" description:"Need repay normal interest total"`
	NeedRepayTotalCompoundInterest     decimal.Decimal `json:"needRepayTotalCompoundInterest" description:"Need repay total compound interest"`
	NeedRepayNormalCompoundInterest    decimal.Decimal `json:"needRepayNormalCompoundInterest" description:"Need repay normal compound interest"`
	NeedRepayOverDueInterest           decimal.Decimal `json:"needRepayOverDueInterest" description:"Need repay over due interest"`
	NeedRepayOverDueCompoundInterest   decimal.Decimal `json:"needRepayOverDueCompoundInterest" description:"Need repay over due compound interest"`
	NeedRepayOverDueNormalInterest     decimal.Decimal `json:"needRepayOverDueNormalInterest" description:"Need repay over due normal interest"`
	NeedRepayOverDueRemainPlanInterest decimal.Decimal `json:"needRepayOverDueRemainPlanInterest" description:"Need repay over due remain plan interest"`
	NeedRepayCollectionFee             decimal.Decimal `json:"needRepayCollectionFee" description:"need repay colleciont fee"`
	NeedRepayTotalFee                  decimal.Decimal `json:"needRepayTotalFee"`
	NeedRepayPPI                       decimal.Decimal `json:"needRepayPPI"`
	NeedRepayDueTotalAmount            decimal.Decimal `json:"needRepayDueTotalAmount"`
	DueInterestAndPrinciple            decimal.Decimal `json:"dueInterestAndPrinciple"`
	PreSettleRepaymentFee              decimal.Decimal `json:"preSettleRepaymentFee"`
	PreSettleAmount                    decimal.Decimal `json:"preSettleAmount"`
	PreSettleTotalAmount               decimal.Decimal `json:"preSettleTotalAmount"`
	BillDetail                         BillDetail      `json:"billDetail"`
	NeedRepayFeeTypeList               []FeeTypeInfo   `json:"needRepayFeeTypeList"`
	CurrentPeriod                      int             `json:"currentPeriod"`
	InstallmentAmount                  decimal.Decimal `json:"installmentAmount"` //current period,planRepayPrinciple +planRepayInterest
	DueDateFlag                        string          `json:"dueDateFlag"`       //1- Before Due Date 2- On due date 3- After due date

	RemainTotalAmount              decimal.Decimal `json:"remainTotalAmount" description:"Remain total amount"`
	Balance                        decimal.Decimal `json:"balance" description:"Balance"`
	RepayAmount                    decimal.Decimal `json:"repayAmount" description:"Repay amount"`
	OutstandingBalance             decimal.Decimal `json:"outstandingBalance" description:"Outstanding balance"`
	OutstandingPPI                 decimal.Decimal `json:"outstandingPPI"`
	RepayTotalPrinciple            decimal.Decimal `json:"repayTotalPrinciple" description:"Repay total principle"`
	RepayImpairmentPrinciple       decimal.Decimal `json:"repayImpairmentPrinciple" description:"Repay impairment principle"`
	RepayOverduePrinciple          decimal.Decimal `json:"repayOverduePrinciple" description:"Repay overdue principle"`
	RepayTotalInterest             decimal.Decimal `json:"repayTotalInterest" description:"Repay total interest"`
	RepayNormalInterestTotal       decimal.Decimal `json:"repayNormalInterestTotal" description:"Repay normal interest total"`
	RepayNormalInterest            decimal.Decimal `json:"repayNormalInterest" description:"Repay normal interest"`
	RepayTotalCompoundInterest     decimal.Decimal `json:"repayTotalCompoundInterest" description:"Repay total compound interest"`
	RepayNormalCompoundInterest    decimal.Decimal `json:"repayNormalCompoundInterest" description:"Repay normal compound interest"`
	RepayOverDueInterest           decimal.Decimal `json:"repayOverDueInterest" description:"Repay over due interest"`
	RepayOverDueCompoundInterest   decimal.Decimal `json:"repayOverDueCompoundInterest" description:"Repay over due compound interest"`
	RepayOverDueNormalInterest     decimal.Decimal `json:"repayOverDueNormalInterest" description:"Repay over due normal interest"`
	RepayOverDueRemainPlanInterest decimal.Decimal `json:"repayOverDueRemainPlanInterest" description:"Repay over due remain plan interest"`
	RepayCollectionFee             decimal.Decimal `json:"repayCollectionFee" description:"Repay collection fee"`
	RepayPrepaymentFee             decimal.Decimal `json:"repayPrepaymentFee" description:"Repay Prepayment fee"`
	RepayTotalFee                  decimal.Decimal `json:"repayTotalFee"`
	BillRepayDetail                BillRepayDetail `json:"billRepayDetail"`
	RepayPPIFee                    decimal.Decimal `json:"repayPPIFee"`
	DiscountInfo                   DiscountObj     `json:"discountInfo"`
	RepayFeeTypeList               []FeeTypeInfo   `json:"repayFeeTypeList"`
}

type BillDetail struct {
	NeedRepayBillAmountTotal         decimal.Decimal `json:"needRepayBillAmountTotal"`
	NeedRepayTotalPrinciple          decimal.Decimal `json:"needRepayTotalPrinciple"`
	NeedRepayTotalInterest           decimal.Decimal `json:"needRepayTotalInterest"`
	NeedRepayDueInterest             decimal.Decimal `json:"needRepayDueInterest"`
	NeedRepayDuePrinciple            decimal.Decimal `json:"needRepayDuePrinciple"`
	NeedRepayCollectionFee           decimal.Decimal `json:"needRepayCollectionFee"`
	NeedRepayTotalFee                decimal.Decimal `json:"needRepayTotalFee"`
	NeedRepayOverduePrincipal        decimal.Decimal `json:"needRepayOverduePrincipal"`
	NeedRepayOverdueInterest         decimal.Decimal `json:"needRepayOverdueInterest"`
	NeedRepayPenaltyInterest         decimal.Decimal `json:"needRepayPenaltyInterest"`
	NeedRepayNormalCompoundInterest  decimal.Decimal `json:"needRepayNormalCompoundInterest"`
	NeedRepayPenaltyCompoundInterest decimal.Decimal `json:"needRepayPenaltyCompoundInterest"`
}
type BillRepayDetail struct {
	RepayBillTotalAmount         decimal.Decimal `json:"repayBillTotalAmount"`
	RepayTotalPrinciple          decimal.Decimal `json:"repayTotalPrinciple"`
	RepayTotalInterest           decimal.Decimal `json:"repayTotalInterest"`
	RepayDueInterest             decimal.Decimal `json:"repayDueInterest"`
	RepayDuePrinciple            decimal.Decimal `json:"repayDuePrinciple"`
	RepayOverduePrincipal        decimal.Decimal `json:"repayOverduePrincipal"`
	RepayOverdueInterest         decimal.Decimal `json:"repayOverdueInterest"`
	RepayCollectionFee           decimal.Decimal `json:"repayCollectionFee"`
	RepayTotalFee                decimal.Decimal `json:"repayTotalFee"`
	RepayPenaltyInterest         decimal.Decimal `json:"repayPenaltyInterest"`
	RepayNormalCompoundInterest  decimal.Decimal `json:"repayNormalCompoundInterest"`
	RepayPenaltyCompoundInterest decimal.Decimal `json:"repayPenaltyCompoundInterest"`
}
type FeeTypeInfo struct {
	FeeType   string          `json:"feeType"`
	FeeAmount decimal.Decimal `json:"feeAmount"`
}

type DiscountObj struct {
	NeedRepayTotalAmountAfterDiscount    decimal.Decimal `json:"needRepayTotalAmountAfterDiscount"`
	DiscountInterestTotal                decimal.Decimal `json:"discountInterestToal"`
	DiscountNormalInterestTotal          decimal.Decimal `json:"discountNormalInterestTotal"`
	DiscountPenaltyInterest              decimal.Decimal `json:"discountPenaltyInterest"`
	DiscountCompoundInterest             decimal.Decimal `json:"discountCompoundInterest"`
	DiscountCompoundPenaltyInterest      decimal.Decimal `json:"discountCompoundPenaltyInterest"`
	NormalInterestTotalAfterDiscount     decimal.Decimal `json:"normalInterestTotalAfterDiscount"`
	PenaltyInterestAfterDiscount         decimal.Decimal `json:"penaltyInterestAfterDiscount"`
	CompoundInterestAfterDiscount        decimal.Decimal `json:"compoundInterestAfterDiscount"`
	CompoundPenaltyInterestAfterDiscount decimal.Decimal `json:"compoundPenaltyInterestAfterDiscount"`
	TotalInterestAfterDiscount           decimal.Decimal `json:"totalInterestAfterDiscount"`
}

func (*AL000003I) GetServiceKey() string {
	return "AL000003"
}
