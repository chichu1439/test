package dao

type T_fee_code_info struct {
	FeeCode           string  `orm:"column(fee_code);pk" description:"费用代码"`
	FeeName           string  `orm:"column(fee_name);size(120);null" description:"费用说明"`
	Currency          string  `orm:"column(currency);size(3);null" description:"币种--枚举值CM0001（参考元数据）"`
	FeeApplyRule      string  `orm:"column(fee_apply_rule);size(1);null" description:"费用应用规则 C-计数器;A-金额"`
	FeeCountFlag      string  `orm:"column(fee_count_flag);size(1);null"`
	FeeOption         string  `orm:"column(fee_option);size(1);null" description:"费用选项 L –分量层级;A –总量层级;N-无层级"`
	TopFeeAmount      float64 `orm:"column(top_fee_amount);null;digits(18);decimals(2)" description:"最高费用金额"`
	ButtonFeeAmount   float64 `orm:"column(button_fee_amount);null;digits(18);decimals(2)" description:"最低费用金额"`
	FeeAmount         float64 `orm:"column(fee_amount);null;digits(18);decimals(2)" description:"费用金额"`
	FeeRate           float64 `orm:"column(fee_rate);null;digits(9);decimals(6)" description:"费率"`
	Level             int     `orm:"column(level);null" description:"层级数"`
	ReduceLvl         string  `orm:"column(reduce_lvl);size(1);null" description:"减免级别 0-最大值;1-最小值;2-汇总;"`
	CriticalPointFlag string  `orm:"column(critical_point_flag);size(1);null"`
	ReduceCode1       string  `orm:"column(reduce_code1);size(5);null" description:"减免代码1"`
	ReduceRelation1   string  `orm:"column(reduce_relation1);size(1);null" description:"减免关系1 O-或;A-与"`
	ReduceCode2       string  `orm:"column(reduce_code2);size(5);null" description:"减免代码2"`
	ReduceRelation2   string  `orm:"column(reduce_relation2);size(1);null" description:"减免关系2 O-或;A-与"`
	ReduceCode3       string  `orm:"column(reduce_code3);size(5);null" description:"减免代码3"`
	ReduceRelation3   string  `orm:"column(reduce_relation3);size(1);null" description:"减免关系3 O-或;A-与"`
	AddFeeFlag        string  `orm:"column(add_fee_flag);size(1);null"`
	AddFeeAmount      float64 `orm:"column(add_fee_amount);null;digits(18);decimals(2)" description:"附加手续费金额"`
	AddFeeRate        float64 `orm:"column(add_fee_rate);null;digits(9);decimals(6)" description:"附加手续费率"`
	Flag1             string  `orm:"column(flag1);size(2);null" description:"标志1"`
	Flag2             string  `orm:"column(flag2);size(2);null" description:"标志2"`
	FeeDealRuleOption string  `orm:"column(fee_deal_rule_option);size(1);null" description:"费用计算处理规则选项 C-调用计算程序计算;S-使用费用代码对应规则"`
	CallPrograme      string  `orm:"column(call_programe);size(10);null" description:"调用程序名"`
	Remark1           string  `orm:"column(remark1);size(10);null" description:"备用1"`
	Remark2           string  `orm:"column(remark2);size(10);null" description:"备用2"`
	Remark3           string  `orm:"column(remark3);size(10);null" description:"备用3"`
	LastMaintDate     string  `orm:"column(last_maint_date);type(date);null" description:"最后更新日期"`
	LastMaintTime     string  `orm:"column(last_maint_time);type(time);null" description:"最后更新时间"`
	LastMaintBrno     string  `orm:"column(last_maint_brno);size(20);null" description:"最后更新机构"`
	LastMaintTell     string  `orm:"column(last_maint_tell);size(20);null" description:"最后更新柜员"`
	TccState          int     `orm:"column(tcc_state);null" description:"tcc状态"`
}

func (t *T_fee_code_info) TableName() string {
	return "t_fee_code_info"
}
