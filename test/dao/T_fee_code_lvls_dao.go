package dao

type T_fee_code_lvls struct {
	FeeCode        string  `orm:"column(fee_code);pk;size(30)" description:"费用代码"`
	LevelSequnce   int     `orm:"column(level_sequnce)" description:"层级数"`
	LevelCountFlag string  `orm:"column(level_count_flag);size(1)"`
	FeeRate        float64 `orm:"column(fee_rate);null;digits(9);decimals(6)" description:"费率"`
	FeeAmount      float64 `orm:"column(fee_amount);null;digits(18);decimals(2)" description:"费用金额"`
	LevelRatio     int     `orm:"column(level_ratio);size(19)" description:"层级系数 金额：0000000000000030000；表示300"`
	Remark         string  `orm:"column(remark);size(120);null" description:"说明"`
	LastMaintDate  string  `orm:"column(last_maint_date);type(date);null" description:"最后更新日期"`
	LastMaintTime  string  `orm:"column(last_maint_time);type(time);null" description:"最后更新时间"`
	LastMaintBrno  string  `orm:"column(last_maint_brno);size(20);null" description:"最后更新机构"`
	LastMaintTell  string  `orm:"column(last_maint_tell);size(20);null" description:"最后更新柜员"`
	TccState       int     `orm:"column(tcc_state);null" description:"tcc状态"`
}

func (t *T_fee_code_lvls) TableName() string {
	return "t_fee_code_lvls"
}
