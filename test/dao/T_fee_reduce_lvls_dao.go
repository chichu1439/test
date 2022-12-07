package dao

type T_fee_reduce_lvls struct {
	ReduceCode      string  `orm:"column(reduce_code);pk;size(30)" description:"减免代码"`
	ReduceDimension string  `orm:"column(reduce_dimension);size(3)" description:"减免维度--枚举值CM0042---000-当前余额 001-交易计数器 002-借记金额 003-借记次数 004-账户状态 005-客户类别 006-客户等级 007-日均余额 008-交易日期段 009-交易时间段"`
	LevelSequnce    int     `orm:"column(level_sequnce)" description:"层级数"`
	LevelCountFlag  string  `orm:"column(level_count_flag);size(1)" description:"减免计算选项 "O-固定金额;F-费率""`
	FeeRate         float64 `orm:"column(fee_rate);null;digits(9);decimals(0)" description:"减免费率"`
	FeeAmount       float64 `orm:"column(fee_amount);null;digits(18);decimals(2)" description:"费用金额"`
	LevelRatio      string  `orm:"column(level_ratio);size(19)" description:"层级系数 金额：0000000000000030000；表示300"`
	Remark          string  `orm:"column(remark);size(120);null" description:"说明"`
	LastMaintDate   string  `orm:"column(last_maint_date);type(date);null" description:"最后更新日期"`
	LastMaintTime   string  `orm:"column(last_maint_time);type(time);null" description:"最后更新时间"`
	LastMaintBrno   string  `orm:"column(last_maint_brno);size(20);null" description:"最后更新机构"`
	LastMaintTell   string  `orm:"column(last_maint_tell);size(20);null" description:"最后更新柜员"`
	TccState        int     `orm:"column(tcc_state);null" description:"tcc状态"`
}

func (t *T_fee_reduce_lvls) TableName() string {
	return "t_fee_reduce_lvls"
}
