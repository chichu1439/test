//
//  Copyright 2020 Formssi Inc.
//
// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//
//  Create date  :  2020/7/9
//
//  Programmer   :  cenzhuocheng
//
//  Last Update  :  2020/7/9  [czc]
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

//
//  文件存储表
//  t_fp_file_save
//
type TFPFileSave struct {
	FileId         string    `orm:"column(file_id);pk;" json:"FileId,omitempty"`
	FileType       string    `orm:"column(file_type);" json:"FileType,omitempty"`
	Currency       string    `orm:"column(currency);" json:"Currency,omitempty"`
	FileName       string    `orm:"column(file_name);" json:"FileName,omitempty"`
	FilePath       string    `orm:"column(file_path);" json:"FilePath,omitempty"`
	FileState      string    `orm:"column(file_state);" json:"FileState,omitempty"`
	FileCreateDate string    `orm:"column(file_create_date);" json:"FileCreateDate,omitempty"`
	FileDosId      string    `orm:"column(file_dos_id);" json:"FileDosId,omitempty"`
	BatchId        string    `orm:"column(batch_id);" json:"BatchId,omitempty"`
	ParticipantId  string    `orm:"column(participant_id);" json:"ParticipantId,omitempty"`
	LastMaintDate  time.Time `orm:"auto_now;type(date);column(last_maint_date);null;" json:"-"`
	LastMaintTime  time.Time `orm:"auto_now;type(time);column(last_maint_time);null;" json:"-"`
	LastMaintBrno  string    `orm:"column(last_maint_brno);null;" json:"LastMaintBrno,omitempty"`
	LastMaintTell  string    `orm:"column(last_maint_tell);null;" json:"LastMaintTell,omitempty"`
	TccState       int       `orm:"column(tcc_state);null;" json:"TccState"`
}

//
//  获取表名
//  Get table name.
//
func (this *TFPFileSave) TableName() string {
	return "t_fp_file_save"
} //end of TableName()
