//
//  Copyright 2020 Formssi Inc.
//
// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//
//  Create date  :  2020/7/2
//
//  Programmer   :  cenzhuocheng
//
//  Last Update  :  2020/7/6  [czc]
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

//
//  报文登记表
//  t_fp_message_register
//
type TFpMessageRegister struct {
	MsgId             string `orm:"column(msg_id);size(320);pk;" json:"MsgId,omitempty"`
	MsgPaticipate     string `orm:"column(msg_paticipate);size(320);null;" json:"MsgPaticipate,omitempty"`
	MsgType           string `orm:"column(msg_type);size(320);null;" json:"MsgType,omitempty"`
	ServiceType       string `orm:"column(service_type);size(320);null;" json:"ServiceType,omitempty"`
	MsgDirection      string `orm:"column(msg_direction);size(320);null;" json:"MsgDirection,omitempty"`
	MsgCreateDate     string `orm:"column(msg_create_date);size(320);null;" json:"MsgCreateDate,omitempty"`
	RedeliverCount    string `orm:"column(redeliver_count);size(320);null;" json:"RedeliverCount,omitempty"`
	MsgData           string `orm:"column(msg_data);size(320);null;" json:"MsgData,omitempty"`
	MsgState          string `orm:"column(msg_state);size(320);null;" json:"MsgState,omitempty"`
	MsgResponseState  string `orm:"column(msg_response_state);size(320);null;" json:"MsgResponseState,omitempty"`
	MsgResponseReason string `orm:"column(msg_response_reason);size(320);null;" json:"MsgResponseReason,omitempty"`
	FrMmbId           string `orm:"column(fr_mmb_id);size(320);null;" json:"FrMmbId,omitempty"`
	ToMmbId           string `orm:"column(to_mmb_id);size(320);null;" json:"ToMmbId,omitempty"`
	Currency          string `orm:"column(currency);size(320);null;" json:"Currency,omitempty"`
	TranNo            string `orm:"column(tran_no);" json:"TranNo,omitempty"`
	ClrSysRef         string `orm:"column(clr_sys_ref);size(35);null;" json:"ClrSysRef,omitempty"`
	ChannelCode       string `orm:"column(channel_code);size(35);null;" json:"ChannelCode,omitempty"`
	LastMaintDate     string `orm:"column(last_maint_date);size(320);null;" json:"last_maint_date,omitempty"`
	LastMaintTime     string `orm:"column(last_maint_time);size(320);null;" json:"LastMaintTime,omitempty"`
	LastMaintBrno     string `orm:"column(last_maint_brno);size(320);null;" json:"LastMaintBrno,omitempty"`
	LastMaintTell     string `orm:"column(last_maint_tell);size(320);null;" json:"LastMaintTell,omitempty"`
	TccState          int    `orm:"column(tcc_state);size(320);null;" json:"TccState"`
}

//
//  获取表名
//  Get table name.
//
func (t *TFpMessageRegister) TableName() string {
	return "t_fp_message_register"
}
