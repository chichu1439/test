//
//  Copyright 2020 Formssi Inc.
//
// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//
//  Create date  :  2020/8/7
//
//  Programmer   :  cenzhuocheng
//
//  Last Update  :  2020/8/7  [czc]
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

type TFPFieldParse struct {
	StdField    string `orm:"pk" json:"StdField,omitempty"`
	TableField  string `json:"TableField,omitempty"`
	DBTabelName string `orm:"column(db_table_name)" json:"DBTabelName,omitempty"`
	Database    string `orm:"column(database)" json:"DbName,omitempty"`
	TccState    int
}

func (TFPFieldParse) TableName() string {
	return "t_fp_field_parse"
}
