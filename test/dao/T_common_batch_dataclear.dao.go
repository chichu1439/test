package dao

type T_common_batch_dataclear struct {
	TableName      string `orm:"column(table_name);size(50);pk" description:"数据表名"`
	HisTableName   string `orm:"column(his_table_name);size(50);null" description:"历史表名"`
	KeepCycs       string `orm:"column(keep_cycs);size(2);null" description:"保留周期 01-按天 02-按周 03-按半月/双周 04-按月 05-按双月 06-按季 07-按半年 08-按年"`
	KeepFreq       int    `orm:"column(keep_freq);size(5);null" description:"保留周期频率"`
	FieldName      string `orm:"column(field_name);size(50);null" description:"日期字段名(清除条件)"`
	OtherCondition string `orm:"column(other_condition);size(256);null" description:"其他条件(sql语句)"`
}
