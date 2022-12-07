package dao

type T_fee_evaluate_entry struct {
	RecordId        int    `orm:"column(record_id);pk" description:"记录ID"`
	FeeEvaluateType string `orm:"column(fee_evaluate_type);size(6)" description:"费用试算类型--枚举值CM0041---000001-存款 000002-取款 000003-转账 000004-汇款 000005-票据出售 000006-票据挂失 000007-存款证明打印"`
	ChannelCode     string `orm:"column(channel_code);size(4)" description:"渠道号--枚举值CM0003---TRM	柜面 IBS	网银 BPS	大额支付系统 HPS	小额支付系统 PGS	省金服 BCH	批次 EXG	POS收单系统交换模块 MID-中间业务平台 TEL-电话银行 MBW-手机银行 SDS-后台流程再造 CMS-对公信贷管理 PMS-个贷管理 EBS-单证系统 CDT-信用卡 SSR-社保系统 MCC-现金管理 ATM-ATM POS-POS YLT-银联 EFS-金融服务平台二代 QSS-汇兑及清算系统 CBM-商票系统 EBM-电子商业汇票系统 TPS-总对总国库系统 CRM-贵宾系统 IQM-智能排队机（叫号机） CIS-后台账户管理 BCS-后台银承 HDX-电子回单箱 BMS-电票系统 MBL-短信平台"`
	FuncCode        string `orm:"column(func_code);size(5)" description:"功能代码"`
	//FinancialFlag        string    `orm:"column(financial_flag);size(1)" description:"金融非金融标志 F-金融 N-非金融"`
	//CashFlag             string    `orm:"column(cash_flag);size(1);null" description:"现金非现金标志 C-现金 T-非现金"`
	//CurrencyExchangeFlag string    `orm:"column(currency_exchange_flag);size(1);null" description:"币种钞汇标识 1-同币种同钞汇; 2-同币种钞折汇; 3-结汇; 4-售汇; 5-兑换""`
	//DebitDiffPalceFlag   string    `orm:"column(debit_diff_palce_flag);size(1);null" description:"借记同城异地标识 1-同网点; 2-同城同分行; 3-同城跨分行; 4-同城跨行; 5-异地同分行; 6-异地跨分行; 7-异地跨行""`
	//CredtDiffPalceFlag   string    `orm:"column(credt_diff_palce_flag);size(1);null" description:"贷记同城异地标识 1-同网点; 2-同城同分行; 3-同城跨分行; 4-同城跨行; 5-异地同分行; 6-异地跨分行; 7-异地跨行"`
	ProdNo        string `orm:"column(prod_no);size(8);null" description:"产品编号"`
	Currency      string `orm:"column(currency);size(3);null" description:"币种--枚举值CM0001（参考元数据枚举值）"`
	MediumType    string `orm:"column(medium_type);size(4);null" description:"费用试算凭证类型--枚举值MD0001（参考元数据枚举值）"`
	MediumTypeDet string `orm:"column(medium_type_det);size(3);null" description:"费用试算凭证类型细分"`
	BranchNo      string `orm:"column(branch_no);size(20);null" description:"分行号"`
	UrgentFlag    string `orm:"column(urgent_flag);size(1);null" description:"加急标志 Y-加急;N-非加急"`
	CustDiff      string `orm:"column(cust_diff);size(5);null" description:"客户区分--枚举值CU0019（参考元数据枚举值）"`
	AttrDetail    string `orm:"column(attr_detail);size(1);null" description:"属性细分--枚举值CU0006---1——个人 2——单位 3——同业 4——内部"`
	Remark1       string `orm:"column(remark1);size(10);null" description:"自定义1"`
	Remark2       string `orm:"column(remark2);size(10);null" description:"自定义2"`
	Remark3       string `orm:"column(remark3);size(10);null" description:"自定义3"`
	EfftDate      string `orm:"column(efft_date);type(date);null" description:"生效日期"`
	ExpDate       string `orm:"column(exp_date);type(date);null" description:"到期日期"`
	EfftFlag      string `orm:"column(efft_flag);size(1);null" description:"生效标识  Y-是 N-否"`
	Remark        string `orm:"column(remark);size(200);null" description:"功能服务费说明"`
	FeePlanNumber string `orm:"column(fee_plan_number);size(15);null" description:"服务费计划号"`
	LastMaintDate string `orm:"column(last_maint_date);type(date);null" description:"最后修改日期"`
	LastMaintTime string `orm:"column(last_maint_time);type(time);null" description:"最后修改时间"`
	LastMaintBrno string `orm:"column(last_maint_brno);size(20);null" description:"最后修改行所"`
	LastMaintTell string `orm:"column(last_maint_tell);size(20);null" description:"最后修改柜员"`
	TccState      int    `orm:"column(tcc_state);null"`
}

func (t *T_fee_evaluate_entry) TableName() string {
	return "t_fee_evaluate_entry"
}
