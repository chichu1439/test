package dao

type T_loan_loan_cont struct {
	LoanDubilNo                    string  `orm:"column(loan_dubil_no);size(32);pk" description:"贷款借据号  智能贷款中贷款账号和借据号相同?27位贷款账号生成规则:合同号+序号(3位)+尾号(2位)编号规则:尾号两位使用6?8(66?68?86?88,4个数字随机拼接到序号最后)序号:合同项下借据的序号递增举例:合同号0401132019101400000168的第一笔借据,随机获取到88为尾号?生成贷款账号:040113201910140000016800188"`
	AcctiAcctNo                    string  `orm:"column(accti_acct_no);size(32);null" description:"核算账号  16位核算账号生成规则:机构码4位+币种代码3位+系统顺序号或流水号9位编号规则:举例:机构:0401,币种代码:156或CNY生成贷款核算账号:0401156010000006"`
	ContrtStusCd                   string  `orm:"column(contrt_stus_cd);size(2);null" description:"合约状态代码  01-正常02-逾期03-减值04-部分核销05-核销06-结清07-减免结清"`
	CustNo                         string  `orm:"column(cust_no);size(14);null" description:"客户编号  02-逾期"`
	LoanProdtNo                    string  `orm:"column(loan_prodt_no);size(20);null" description:"贷款产品编号  8位产品编号生成规则:产品设计阶段'1位'+产品分类1段'1位'+产品分类2段'1位'+序号'5位'产品设计阶段:B基础产品,S可售产品,产品分类'1段':0金融,1衍生,产品分类'1段'+,产品分类'2段',,,,,,,金融类产品:01:存款,02:贷款,03:,理财,,,,,,,,,,,,,,,,,,,,,,,,,,,,04:基金,05:保险,,,,,,,,衍生:11:积分,12:红包,,13:,卡券,14:商品举例:产品设计阶段:可售产品+产品分类1段:金融,产品分类:'金融'贷款,序号:00001生成产品编号:S0200001"`
	LoanProdtVersNo                string  `orm:"column(loan_prodt_vers_no);size(5);null" description:"贷款产品版本编号  5位产品版本号生成规则:产品编号独立,修改产品参数产品编号不变,提升版本号,整形数,顺序加,转为前面填充0的5位字符串存储,举例:原贷款产品O0200001版本号00010最高授信额度50,000.00,修改为最高授信额度100,000.00,版本号变更为00011,"`
	MakelnOrgNo                    string  `orm:"column(makeln_org_no);size(4);null" description:"放款机构号"`
	IndvCrtfTypCd                  string  `orm:"column(indv_crtf_typ_cd);size(4);null" description:"个人证件类型代码  参考标准代码:CD0015"`
	IndvCrtfNo                     string  `orm:"column(indv_crtf_no);size(20);null" description:"个人证件号码"`
	RevneCmpdCd                    string  `orm:"column(revne_cmpd_cd);size(1);null" description:"计收复利代码     0-不收复利,,1-收取复利,2-由全局参数控制"`
	MakelnManrCd                   string  `orm:"column(makeln_manr_cd);size(1);null" description:"放款方式代码  0-自主支付,,1-受托支付,2-混合支付"`
	RelaAcctDtrmnManrCd            string  `orm:"column(rela_acct_dtrmn_manr_cd);size(2);null" description:"关联账户确定方式代码  01-客户层,02-产品层,03-合同层,04-借据层"`
	RepayDayDtrmnManrCd            string  `orm:"column(repay_day_dtrmn_manr_cd);size(2);null" description:"还款日确定方式代码       01-客户层,02-产品层,03-合同层,04-借据层"`
	LoanDeadl                      int     `orm:"column(loan_deadl);null" description:"贷款周期"`
	LoanDeadlCycCd                 string  `orm:"column(loan_deadl_cyc_cd);null" description:"贷款周期代码"`
	PmitAdvRepayTms                int     `orm:"column(pmit_adv_repay_tms);null" description:"允许提前还款次数"`
	AdvRepayLimitBgnDt             string  `orm:"column(adv_repay_limit_bgn_dt);type(date);null" description:"提前还款限制开始日期"`
	AdvRepayLimitDays              int     `orm:"column(adv_repay_limit_days);null" description:"提前还款限制天数"`
	LimitTrmInsidPmitPartlRepayFlg string  `orm:"column(limit_trm_insid_pmit_partl_repay_flg);size(1);null" description:"限制期内允许部分还款标志  0否1是"`
	LimitTrmInsidPmitPayOffFlg     string  `orm:"column(limit_trm_insid_pmit_pay_off_flg);size(1);null" description:"限制期内允许结清标志  0否1是"`
	OpenAcctDt                     string  `orm:"column(open_acct_dt);type(date);null" description:"开户日期  记录客户首次办理登记的具体日期,"`
	FsttmForsprtDt                 string  `orm:"column(fsttm_forsprt_dt);type(date);null" description:"首次支用日期"`
	BegintDt                       string  `orm:"column(begint_dt);type(date);null" description:"起息日期  是否应在计息中心记录"`
	OrgnlMatrDt                    string  `orm:"column(orgnl_matr_dt);type(date);null" description:"原始到期日期"`
	CurCd                          string  `orm:"column(cur_cd);size(3);null" description:"币种代码  参考标准代码:CD0040"`
	LoanAmt                        float64 `orm:"column(loan_amt);null;digits(18);decimals(2)" description:"贷款金额  如果多次贷放,或者信用卡模式,此金额如何处理?"`
	EmbFlg                         string  `orm:"column(emb_flg);size(1);null" description:"挪用标志  0否1是"`
	DecdEmbDt                      string  `orm:"column(decd_emb_dt);type(date);null" description:"判定挪用日期"`
	MansbjTypCd                    string  `orm:"column(mansbj_typ_cd);size(3);null" description:"主体类型代码  01-农户02-非农个人03-农村个体工商户90-农户小额91-农村个体户小额92-农村小微企业小额93-非农小微企业小额94-非农个体工商户小额"`
	CtrtNo                         string  `orm:"column(ctrt_no);size(32);null" description:"合同号  22位合同号生成规则:机构号4位+业务属性2位+年月日8位+顺序号6位+尾号2位编号规则:1,尾号两位使用6,8'66,68,86,88,4个数字随机拼接到序号最后'2,业务属性:,02-微贷,,,13-个贷,14-小企业,01-公司,08-承兑,02-普贷举例:机构:0401,,业务属性:个人贷款,日期:20191014生成合同号:0401132019101400000168"`
	LoanGuarManrCd                 string  `orm:"column(loan_guar_manr_cd);size(2);null" description:"贷款担保方式代码  参考标准代码:CD0090"`
	OthConsmTypCd                  string  `orm:"column(oth_consm_typ_cd);size(3);null" description:"其他消费类型代码      1-个人日常消费,2-装修,3-旅游,4-教育,5-医疗,6-其他消费,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,"`
	RepayManrCd                    string  `orm:"column(repay_manr_cd);size(3);null" description:"还款方式代码      参考标准代码:CD0062标准基础上新增:51-按月付息按季还本52-按月付息按半年还本53-按月付息按年还本"`
	IntStlDayDtrmnManrCd           string  `orm:"column(int_stl_day_dtrmn_manr_cd);size(3);null" description:"结息日确定方式代码    1-统一定日,2-按户定日,3-按指定日期4-按产品,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,"`
	AdvRepayTms                    int     `orm:"column(adv_repay_tms);null" description:"提前还款次数"`
	TranNormlLoanDt                string  `orm:"column(tran_norml_loan_dt);type(date);null" description:"转正常贷款日期"`
	BldInstltRepayFlg              string  `orm:"column(bld_instlt_repay_flg);size(1);null" description:"建立分期还款标志  描述是否建立分期还款,0否1是,"`
	IntStlPrtyTypCd                string  `orm:"column(int_stl_prty_typ_cd);size(1);null" description:"结息优先类型代码  1-以计划还息金额为结息金额,2-以计提金额为结息金额"`
	LoanIntrtAdjCycCd              string  `orm:"column(loan_intrt_adj_cyc_cd);size(3);null" description:"贷款利率调整周期代码"`
	LoanIntrtAdjCycQty             int     `orm:"column(loan_intrt_adj_cyc_qty);null" description:"贷款利率调整周期数量"`
	StpDeductFlg                   string  `orm:"column(stp_deduct_flg);size(1);null" description:"停止批扣标志  描述银行与客户达成减免协议后,停止通过批量扣款方式进行追扣,由客户经理通过其他交易手工发起扣款,0否1是,"`
	StpDeductRsnTypCd              string  `orm:"column(stp_deduct_rsn_typ_cd);size(1);null" description:"停止批扣原因类型代码  描述停止批扣的原因类型信息,1-建立减免计划,2-确认无法收回"`
	StpDeductCnfrmDt               string  `orm:"column(stp_deduct_cnfrm_dt);type(date);null" description:"停止批扣确认日期"`
	LoanIntrtAdjManrCd             string  `orm:"column(loan_intrt_adj_manr_cd);size(3);null" description:"贷款利率调整方式代码  1-对日,2-周期开始日,3-固定"`
	ExpdayPayOffManrCd             string  `orm:"column(expday_pay_off_manr_cd);size(2);null" description:"到期日结清方式代码  0-系统自动1-手工"`
	GraceTrmIntacrFlg              string  `orm:"column(grace_trm_intacr_flg);size(1);null" description:"宽限期计息标志  0否1是"`
	EvrpridMaxGraceTrmDays         int     `orm:"column(evrprid_max_grace_trm_days);null" description:"每期最大宽限期天数"`
	ContrtPrdGraceTrmTotlDays      int     `orm:"column(contrt_prd_grace_trm_totl_days);null" description:"合约期间宽限期总天数"`
	AdvRepayColtfeFlg              string  `orm:"column(adv_repay_coltfe_flg);size(1);null" description:"提前还款收费标志  0否1是"`
	ColtfeManrCd                   string  `orm:"column(coltfe_manr_cd);size(1);null" description:"收费方式代码  1:按笔2:按金额比例,"`
	SglColtfeAmt                   float64 `orm:"column(sgl_coltfe_amt);null;digits(18);decimals(2)" description:"单笔收费金额"`
	ColtfeAmtCrdnlnbrCd            string  `orm:"column(coltfe_amt_crdnlnbr_cd);size(4);null" description:"收费金额基数代码  1:还款金额2:贷款余额3:借款金额"`
	ColtfeRatio                    float64 `orm:"column(coltfe_ratio);null;digits(9);decimals(6)" description:"收费比例"`
	AcrdgRatioColtfeCeilAmt        float64 `orm:"column(acrdg_ratio_coltfe_ceil_amt);null;digits(18);decimals(2)" description:"按比例收费上限金额"`
	AcrdgRatioColtfeFloorAmt       float64 `orm:"column(acrdg_ratio_coltfe_floor_amt);null;digits(18);decimals(2)" description:"按比例收费下限金额"`
	CurrExecTmprd                  int     `orm:"column(curr_exec_tmprd);null" description:"当前执行期次"`
	RepayPlanAdjFlg                string  `orm:"column(repay_plan_adj_flg);size(1);null" description:"还款计划调整标志  0否1是"`
	RestFlg                        string  `orm:"column(rest_flg);size(1);null" description:"停息标志  0否1是"`
	RestDt                         string  `orm:"column(rest_dt);type(date);null" description:"停息日期"`
	TranDvalDt                     string  `orm:"column(tran_dval_dt);type(date);null" description:"转减值日期"`
	TranDvalFlg                    string  `orm:"column(tran_dval_flg);size(1);null" description:"转减值标志  0否1是"`
	ExtsnTms                       int     `orm:"column(extsn_tms);null" description:"展期次数  展期的总次数,"`
	WaitExtsnFlg                   string  `orm:"column(wait_extsn_flg);size(1);null" description:"待展期标志  0-否1-是"`
	CurrLoanRiskClsfCd             string  `orm:"column(curr_loan_risk_clsf_cd);size(2);null" description:"当前贷款风险分类代码  1,正常2,关注3,次级4,可疑5,损失11,正常一级12,正常二级13,正常三级14,正常四级15,正常五级21,关注一级22,关注二级23,关注三级31,次级一级32,次级二级41,可疑级51,损失级"`
	MatrDt                         string  `orm:"column(matr_dt);type(date);null" description:"到期日期  记录到期结清的日期,"`
	PayOffDt                       string  `orm:"column(pay_off_dt);type(date);null" description:"结清日期"`
	FinlModfyDt                    string  `orm:"column(finl_modfy_dt);type(date);null" description:"最后修改日期  N"`
	FinlModfyTm                    string  `orm:"column(finl_modfy_tm);type(time);null" description:"最后修改时间  N"`
	FinlModfyOrgNo                 string  `orm:"column(finl_modfy_org_no);size(4);null" description:"最后修改机构号  N"`
	FinlModfyTelrNo                string  `orm:"column(finl_modfy_telr_no);size(6);null" description:"最后修改柜员号  N"`
	TccState                       int8    `orm:"column(tcc_state)" description:"TCC状态值[0:正常 1:新增 2:删除]"`
}

func (t *T_loan_loan_cont) TableName() string {
	return "t_loan_loan_cont"
}
