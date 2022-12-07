package dao

type T_fee_reduce_info struct {
	ReduceCode        string  `orm:"column(reduce_code);pk" description:"减免代码"`
	ReduceName        string  `orm:"column(reduce_name);size(120);null" description:"减免说明"`
	ReduceCodeType    string  `orm:"column(reduce_code_type);size(1);null" description:"减免代码类型 S-单一减免代码;C-组合减免代码"`
	ReduceDimension   string  `orm:"column(reduce_dimension);size(3);null" description:"减免维度 000-当前余额; 001-交易计数器; 002-借记金额; 003-借记次数; 004-账户状态; 005-客户类别; 006-客户等级: 007-日均余额"`
	ReduceType        string  `orm:"column(reduce_type);size(1);null" description:"减免类型 "V-按值;M-多值;R-按范围;Y-分层;O-固定""`
	TotalLvls         int     `orm:"column(total_lvls);null" description:"总层数"`
	CriticalPointFlag string  `orm:"column(critical_point_flag);size(1);null" description:"临界点标识 0-包含;1-不包含； 2－期限包含金额不包含； 3－金额包含期限不包含"`
	ReduceMode        string  `orm:"column(reduce_mode);size(1);null" description:"减免方式 A-固定金额;F-百分比;"`
	ReduceAmount      float64 `orm:"column(reduce_amount);null;digits(18);decimals(2)" description:"减免金额"`
	ReducePercent     float64 `orm:"column(reduce_percent);null;digits(9);decimals(6)" description:"减免百分比"`
	Remark1           string  `orm:"column(remark1);size(10);null" description:"备用1"`
	Remark2           string  `orm:"column(remark2);size(10);null" description:"备用2"`
	Remark3           string  `orm:"column(remark3);size(10);null" description:"备用3"`
	LastMaintDate     string  `orm:"column(last_maint_date);type(date);null" description:"最后更新日期"`
	LastMaintTime     string  `orm:"column(last_maint_time);type(time);null" description:"最后更新时间"`
	LastMaintBrno     string  `orm:"column(last_maint_brno);size(20);null" description:"最后更新机构"`
	LastMaintTell     string  `orm:"column(last_maint_tell);size(20);null" description:"最后更新柜员"`
	TccState          int     `orm:"column(tcc_state);null" description:"tcc状态"`
}

func (t *T_fee_reduce_info) TableName() string {
	return "t_fee_reduce_info"
}
