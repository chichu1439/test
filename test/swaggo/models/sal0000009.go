package models

import (
	"github.com/shopspring/decimal"
)

type AL000009I struct {
	ListProductId    []string `json:"listProductId" description:"List product id"`
	BatchDate        string   `json:"batchDate" description:"Batch date"`
	FirstDateOfMonth string   `json:"lastDateOfMonth" description:"First date of month"`
}

type AL000009O struct {
	CalculateResult      map[string]AccountProductCalcResult      `json:"calculateResult" description:"Calculate result"`
	CalculateMonthResult map[string]AccountProductCalcMonthResult `json:"calculateMonthResult" description:"Calculate month result"`
}

type AccountProductCalcResult struct {
	SumUnpaidBalance         decimal.Decimal `json:"sumUnpaidBalance" description:"Sum unpaid balance"`
	SumDrawTotalAmount       decimal.Decimal `json:"sumDrawTotalAmount" description:"Sum draw total amount"`
	SumAccrualBalance        decimal.Decimal `json:"sumAccrualBalance" description:"The total amount of interest accrued on the day"`
	SumInterestOrig          decimal.Decimal `json:"sumInterestOrig" description:"Sum interest orig"`
	SumInterest              decimal.Decimal `json:"sumInterest" description:"Sum interest"`
	SumPenalty               decimal.Decimal `json:"sumPenalty" description:"Sum penalty"`
	SumCompound              decimal.Decimal `json:"sumCompound" description:"Sum compound"`
	SumCompoundAndSumPenalty decimal.Decimal `json:"sumCompoundAndSumPenalty" description:"Sum compound and sum penalty"`
	SumFee                   decimal.Decimal `json:"sumFee" description:"Sum fee"`
}

type AccountProductCalcMonthResult struct {
	SumMonthUnpaidBalance         decimal.Decimal `json:"sumMonthUnpaidBalance" description:"Sum month unpaid balance"`
	SumMonthDrawTotalAmount       decimal.Decimal `json:"sumMonthDrawTotalAmount" description:"Sum month draw total amount"`
	SumMonthInterestOrig          decimal.Decimal `json:"sumMonthInterestOrig" description:"Sum month interest orig"`
	SumMonthInterest              decimal.Decimal `json:"sumMonthInterest" description:"Sum month interest"`
	SumMonthPenalty               decimal.Decimal `json:"sumMonthPenalty" description:"Sum month penalty"`
	SumMonthCompound              decimal.Decimal `json:"sumMonthCompound" description:"Sum month compound"`
	SumMonthCompoundAndSumPenalty decimal.Decimal `json:"sumMonthCompoundAndSumPenalty" description:"Sum month compound and sum penalty"`
	SumMonthFee                   decimal.Decimal `json:"sumMonthFee" description:"Sum month fee"`
}

func (*AL000009I) GetServiceKey() string {
	return "AL000009"
}
