package dao

import (
	"time"
)

type T_loan_contract_suppt_info_form struct {
	LoanDubilNo                      string    `orm:"column(loan_dubil_no);pk" description:"贷款借据号"`
	CustNo                           string    `orm:"column(cust_no);size(14);null" description:"客户编号"`
	LoanTotlPridnum                  int       `orm:"column(loan_totl_pridnum);null" description:"贷款总期数"`
	CurrExecPridnum                  int       `orm:"column(curr_exec_pridnum);null" description:"当前执行期数"`
	LoanRemainPridnum                int       `orm:"column(loan_remain_pridnum);null" description:"贷款剩余期数"`
	PrinDueBgnDt                     time.Time `orm:"column(prin_due_bgn_dt);type(date);null" description:"本金拖欠开始日期"`
	IntrDueBgnDt                     time.Time `orm:"column(intr_due_bgn_dt);type(date);null" description:"利息拖欠开始日期"`
	AccmDuePridnum                   int       `orm:"column(accm_due_pridnum);null" description:"累计拖欠期数"`
	AccmWrtoffAmt                    float64   `orm:"column(accm_wrtoff_amt);null;digits(18);decimals(2)" description:"累计核销金额"`
	LsttrmCompEtimLnRiskClsfDt       string    `orm:"column(lsttrm_comp_etim_ln_risk_clsf_dt);type(date);null" description:"上期机评贷款风险分类日期"`
	LsttrmCompEtimLnRiskClsfCd       string    `orm:"column(lsttrm_comp_etim_ln_risk_clsf_cd);size(2);null" description:"上期机评贷款风险分类代码  1,正常2,关注3,次级4,可疑5,损失11,正常一级12,正常二级13,正常三级14,正常四级15,正常五级21,关注一级22,关注二级23,关注三级31,次级一级32,次级二级41,可疑级51,损失级记录机器评测的贷款风险分类结果,主要有正常,关注,次级,可疑,损失等,"`
	CurtprdCompEtimLnRiskClsfDt      string    `orm:"column(curtprd_comp_etim_ln_risk_clsf_dt);type(date);null" description:"本期机评贷款风险分类日期"`
	CurtprdCompEtimLnRiskClsfCd      string    `orm:"column(curtprd_comp_etim_ln_risk_clsf_cd);size(2);null" description:"本期机评贷款风险分类代码  1,正常2,关注3,次级4,可疑5,损失11,正常一级12,正常二级13,正常三级14,正常四级15,正常五级21,关注一级22,关注二级23,关注三级31,次级一级32,次级二级41,可疑级51,损失级"`
	LsttrmArtgclIdtfyLoanRiskClsfDt  string    `orm:"column(lsttrm_artgcl_idtfy_loan_risk_clsf_dt);type(date);null" description:"上期人工认定贷款风险分类日期"`
	LsttrmArtgclIdtfyLoanRiskClsfCd  string    `orm:"column(lsttrm_artgcl_idtfy_loan_risk_clsf_cd);size(2);null" description:"上期人工认定贷款风险分类代码  1,正常2,关注3,次级4,可疑5,损失11,正常一级12,正常二级13,正常三级14,正常四级15,正常五级21,关注一级22,关注二级23,关注三级31,次级一级32,次级二级41,可疑级51,损失级"`
	CurtprdArtgclIdtfyLoanRiskClsfDt string    `orm:"column(curtprd_artgcl_idtfy_loan_risk_clsf_dt);type(date);null" description:"本期人工认定贷款风险分类日期"`
	CurtprdArtgclIdtfyLoanRiskClsfCd string    `orm:"column(curtprd_artgcl_idtfy_loan_risk_clsf_cd);size(2);null" description:"本期人工认定贷款风险分类代码  1,正常2,关注3,次级4,可疑5,损失11,正常一级12,正常二级13,正常三级14,正常四级15,正常五级21,关注一级22,关注二级23,关注三级31,次级一级32,次级二级41,可疑级51,损失级"`
	FinlModfyDt                      string    `orm:"column(finl_modfy_dt);type(date);null" description:"最后修改日期"`
	FinlModfyTm                      string    `orm:"column(finl_modfy_tm);type(time);null" description:"最后修改时间"`
	FinlModfyOrgNo                   string    `orm:"column(finl_modfy_org_no);size(4);null" description:"最后修改机构号"`
	FinlModfyTelrNo                  string    `orm:"column(finl_modfy_telr_no);size(6);null" description:"最后修改柜员号"`
	TccState                         int       `orm:"column(tcc_state);null"`
}

func (t *T_loan_contract_suppt_info_form) TableName() string {
	return "t_loan_contract_suppt_info_form"
}
