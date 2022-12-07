//
//  Copyright 2020 Formssi Inc.
//
// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//
//  Create date  :  2020/7/27
//
//  Programmer   :  cenzhuocheng
//
//  Last Update  :  2020/7/27  [czc]
//
// ----------------------------------------------------------
//  Interface:
//
//  Function:
//
//  Struct:
//
// ==========================================================
package dao

import (
	"time"
)

type TFPSysChargeStd struct {
	ChargeStdId          string    `orm:"column(charge_std_id);pk;" json:"ChargeStdId,omitempty"`
	ServiceType          string    `orm:"column(service_type)" json:"ServiceType,omitempty"`
	ChargingTarget       string    `orm:"column(charging_target)" json:"ChargingTarget,omitempty"`
	ChargingMode         string    `orm:"column(charging_mode)" json:"ChargingMode,omitempty"`
	ChargingUnit         string    `orm:"column(charging_unit)" json:"ChargingUnit,omitempty"`
	ChargingCurrency     string    `orm:"column(charging_currency)" json:"ChargingCurrency,omitempty"`
	ChargeAmount         float64   `orm:"column(charge_amount)" json:"ChargeAmount,omitempty"`
	ChargeRate           float64   `orm:"column(charge_rate)" json:"ChargeRate,omitempty"`
	ChargeAmountMin      float64   `orm:"column(charge_amount_min)" json:"ChargeAmountMin,omitempty"`
	ChargeAmountMax      float64   `orm:"column(charge_amount_max)" json:"ChargeAmountMax,omitempty"`
	ChargeAdjustmentRate float64   `orm:"column(charge_adjustment_rate)" json:"ChargeAdjustmentRate,omitempty"`
	EffectiveDate        string    `orm:"column(effective_date);" json:"EffectiveDate,omitempty"`
	State                string    `orm:"column(state)" json:"State,omitempty"`
	LastMaintDate        time.Time `orm:"auto_now;type(date);column(last_maint_date);null;" json:"-"`
	LastMaintTime        time.Time `orm:"auto_now;type(time);column(last_maint_time);null;" json:"-"`
	LastMaintBrno        string    `orm:"column(last_maint_brno);null;" json:"LastMaintBrno,omitempty"`
	LastMaintTell        string    `orm:"column(last_maint_tell);null;" json:"LastMaintTell,omitempty"`
	TccState             int       `orm:"column(tcc_state);" json:"TccState"`
}

func (this *TFPSysChargeStd) TableName() string {
	return "t_fp_sys_charge_std"
} //end of TableName()
