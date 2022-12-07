//
//  Copyright 2020 Formssi Inc.
//
// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//
//  Create date  :  2020/8/6
//
//  Programmer   :  cenzhuocheng
//
//  Last Update  :  2020/8/6  [czc]
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

type TFPTableName struct {
	TransactionType string `orm:"column(transaction_type);pk;" json:"TransactionType,omitempty" pk:"1"`
	DBName          string `orm:"column(db_name);" json:"DBName,omitempty"`
	DataTableName   string `orm:"column(data_table_name);" json:"DataTableName,omitempty"`
	Purpose         string `orm:"column(purpose);" json:"Purpose,omitempty" pk:"2"`
	TccState        int
}

func (t *TFPTableName) TableName() string {
	return "t_fp_tablename"
}
