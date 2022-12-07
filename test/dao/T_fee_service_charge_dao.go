package dao

type T_fee_service_charge struct {
	FeePlanNumber    string `orm:"column(fee_plan_number);pk;size(15)" description:"服务费计划号"`
	FeeSequenNumber  int    `orm:"column(fee_sequen_number)" description:"收费顺序号"`
	FeeCode          string `orm:"column(fee_code);size(30)" description:"费用代码"`
	FeeAbsctractCode string `orm:"column(fee_absctract_code);size(3)" description:"费用摘要码--枚举值CM0011---TRF-转账 TUP-充值 WDR- 提现 OSP-商城支付 SCP-扫码支付 RDM-理财赎回 PCH-理财申购 LON-贷款借款 RPP-还款-本金 RPI-还款-利息 RPF-还款-手续费 CBT-跨行转账-手续费 CTC-跨行转账-加急手续费"`
	PostOption       string `orm:"column(post_option);size(1)" description:"过账选项 0-单独过账;1-与其它费用一起过账"`
	CarryDownOption  string `orm:"column(carry_down_option);size(1)" description:"结转选项 1-收取服务费,其余减免;2-收取服务费,其余结转;3-服务费全部减免;4-服务费全部结转;5-收取服务费,金额不足则拒绝;"`
	FeeAmountType    string `orm:"column(fee_amount_type);size(3)" description:"费用金额类型 001-发生额;002-账户余额;003-日均余额"`
	SubjectNo        string `orm:"column(subject_no);size(8);null" description:"科目号"`
	SubjectDetail    string `orm:"column(subject_detail);size(4);null" description:"科目细目"`
	Flag1            string `orm:"column(flag1);size(1);null" description:"标志1"`
	Flag2            string `orm:"column(flag2);size(2);null" description:"标志2"`
	Back1            string `orm:"column(back1);size(10);null" description:"备用1"`
	Back2            string `orm:"column(back2);size(10);null" description:"备用2"`
	Back3            string `orm:"column(back3);size(10);null" description:"备用3"`
	LastMaintDate    string `orm:"column(last_maint_date);type(date);null" description:"最后更新日期"`
	LastMaintTime    string `orm:"column(last_maint_time);type(time);null" description:"最后更新时间"`
	LastMaintBrno    string `orm:"column(last_maint_brno);size(20);null" description:"最后更新机构"`
	LastMaintTell    string `orm:"column(last_maint_tell);size(20);null" description:"最后更新柜员"`
	TccState         int    `orm:"column(tcc_state);null" description:"tcc状态"`
}

func (t *T_fee_service_charge) TableName() string {
	return "t_fee_service_charge"
}
