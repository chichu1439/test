package dao

type T_fee_plan struct {
	FeePlanNumber       string  `orm:"column(fee_plan_number);pk" description:"服务费计划号"`
	FeePlanName         string  `orm:"column(fee_plan_name);size(120)" description:"服务费计划说明"`
	Currency            string  `orm:"column(currency);size(3)" description:"币种--枚举值CM0001（参考元数据枚举值）"`
	TopFeeAmount        float64 `orm:"column(top_fee_amount);null;digits(18);decimals(2)" description:"最高费用金额"`
	ButtonFeeAmount     float64 `orm:"column(button_fee_amount);null;digits(18);decimals(2)" description:"最低费用金额"`
	FeeUserDefined      string  `orm:"column(fee_user_defined);size(1)" description:"费用自定义选项--S-系统定义;U-用户自定义;"`
	FeeAbstractCode     string  `orm:"column(fee_abstract_code);size(3)" description:"服务费摘要码--枚举值CM0011---TRF-转账 TUP-充值 WDR- 提现 OSP-商城支付 SCP-扫码支付 RDM-理财赎回 PCH-理财申购 LON-贷款借款 RPP-还款-本金 RPI-还款-利息 RPF-还款-手续费 CBT-跨行转账-手续费 CTC-跨行转账-加急手续费"`
	ReduceLvl           string  `orm:"column(reduce_lvl);size(1);null" description:"减免级别 0-最大值;1-最小值;2-汇总;"`
	ReduceCode1         string  `orm:"column(reduce_code1);size(5);null" description:"减免代码1"`
	ReduceRelation1     string  `orm:"column(reduce_relation1);size(1);null" description:"减免关系1 O-或;A-与"`
	ReduceCode2         string  `orm:"column(reduce_code2);size(5);null" description:"减免代码2"`
	ReduceRelation2     string  `orm:"column(reduce_relation2);size(1);null" description:"减免关系2 O-或;A-与"`
	ReduceCode3         string  `orm:"column(reduce_code3);size(5);null" description:"减免代码3"`
	PostOption          string  `orm:"column(post_option);size(1);null" description:"过账选项 0-单独过账;1-与其它费用一起过账"`
	ReduceOption        string  `orm:"column(reduce_option);size(2);null" description:"减免选项 0-不使用费用和计划减免代码;1-使用费用和计划减免代码;2-仅使用费用减免代码;3-仅使用计划减免代码"`
	CarryDownOption     string  `orm:"column(carry_down_option);size(1);null" description:"结转选项 1-收取服务费,其余减免;2-收取服务费,其余结转;3-服务费全部减免;4-服务费全部结转;5-收取服务费,金额不足则拒绝;"`
	BalanceOption       string  `orm:"column(balance_option);size(1);null" description:"余额选项 0-当前余额;1-不含透支可用余额;2-含透支可用余额;"`
	ChargeControlOption string  `orm:"column(charge_control_option);size(1);null" description:"扣收控制选项 0-遵循合约借贷控制;1-忽略合约借贷控制;"`
	OverdueFineCode     string  `orm:"column(overdue_fine_code);size(5);null" description:"滞纳金代码"`
	OverdueFineSubject  string  `orm:"column(overdue_fine_subject);size(8);null" description:"滞纳金核算科目号"`
	OverdueFineDetail   string  `orm:"column(overdue_fine_detail);size(4);null" description:"滞纳金核算细目"`
	AddFeeCalOption     string  `orm:"column(add_fee_cal_option);size(1);null"`
	AddFeeAmount        string  `orm:"column(add_fee_amount);size(18);null" description:"附加手续费金额"`
	AddFeeRate          string  `orm:"column(add_fee_rate);size(9);null" description:"附加手续费率"`
	LastMaintDate       string  `orm:"column(last_maint_date);type(date);null" description:"最后更新日期"`
	LastMaintTime       string  `orm:"column(last_maint_time);type(time);null" description:"最后更新时间"`
	LastMaintBrno       string  `orm:"column(last_maint_brno);size(20);null" description:"最后更新机构"`
	LastMaintTell       string  `orm:"column(last_maint_tell);size(20);null" description:"最后更新柜员"`
	TccState            int     `orm:"column(tcc_state);null" description:"tcc状态"`
}

func (t *T_fee_plan) TableName() string {
	return "t_fee_plan"
}
