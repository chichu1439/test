//Version: v1.0.0.0
package models

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type SPR0000010_CANCELI struct {
	ScenarioId            string                 `json:"scenarioId" validate:"required,max=10" description:"Scene ID"`                                                                                                                                                //场景ID sting Y 10位
	ScenName              string                 `json:"scenName" validate:"required,max=50" description:"Scene name"`                                                                                                                                                // 场景名称 sting Y 50位
	DiscountStg           string                 `json:"discountStg" validate:"required,oneof=0 1" description:"Preferential strategy(0-item by item fee discount 1-total fee discount)"`                                                                         //优惠策略 string Y "0-逐项费用优惠 1-合计费用优惠"
	DiscountSumType       string                 `json:"discountSumType" validate:"required,oneof=0 1 2" description:"Preferential combination method(0-minimum discount 1-maximum discount 2-discount superposition * valid when the discount strategy is 1)"` // 优惠组合方式 string Y "0-最低优惠 1-最高优惠 2-优惠叠加 *优惠策略为1时有效"
	DiscountCalculateType string                 `json:"discountCalculateType" validate:"required,oneof=0 1" description:"Discount calculation method(0-specify the discount proportion 1-specify the discount amount * valid when the discount policy is 1)"`    // 优惠计算方式 float Y "0-指定优惠比例 1-指定优惠金额 *优惠策略为1时有效"
	DiscountMax           float64                `json:"discountMax" validate:"required" description:"Highest discount"`                                                                                                                                              // 最高优惠 10, 2 Y
	DiscountMin           float64                `json:"discountMin" validate:"required" description:"Lowest discount"`                                                                                                                                               // 最低优惠 10, 2 Y
	EffectDate            string                 `json:"effectDate" validate:"required" description:"effective date"`                                                                                                                                                 //生效日期 date Y
	AStrategy             []SPR0000010IAStrategy `json:"aStrategy" validate:"required" description:"Expense strategy array"`                                                                                                                                          // 费用策略数组 Array Y
}

type SPR0000010IAStrategy struct {
	StrategyId string `json:"strategyId" validate:"required" description:"PI000006OStrategy ID"` //策略ID string
}

type SPR0000010_CANCELO struct {
	ScenarioId string `json:"scenarioId" description:"Scene ID"` //场景ID sting Y 10位
}

func (*SPR0000010_CANCELI) GetServiceKey() string {
	return "spr0000010"
}

// customFunc 自定义字段级别校验方法
func (*SPR0000010_CANCELI) CustomFunc(fl validator.FieldLevel) bool {
	date, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}
	if date.Before(time.Now()) {
		return false
	}
	return true
}
