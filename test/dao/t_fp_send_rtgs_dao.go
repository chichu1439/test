//
//  Copyright 2020 Formssi Inc.
//
// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//
//  Create date  :  2020/8/20
//
//  Programmer   :  cenzhuocheng
//
//  Last Update  :  2020/8/20  [czc]
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

type TFPSendRTGS struct {
	RecordId int    `orm:"pk" `
	RtgsData string `json:"RtgsData,omitempty"`
	TccState int
}

func (TFPSendRTGS) TableName() string {
	return "t_fp_send_rtgs"
}
