package dao

//
//
//  文件字段表
//

type TFPFileField struct {
	RecordsId   int    `orm:"column(record_id);pk;" json:"RecordsId,omitempty"`
	DBTableName string `orm:"column(table_name);" json:"DBTableName,omitempty"`
	TableField  string `orm:"column(table_field);" json:"TableField,omitempty"`
	FileField   string `orm:"column(file_field);" json:"FileField,omitempty"`
	FieldSeq    int    `orm:"column(field_seq);" json:"FieldSeq,omitempty"`
}

//
//  获取表名
//  Get table name.
//
func (this *TFPFileField) TableName() string {
	return "t_fp_file_field"
} //end of TableName()
