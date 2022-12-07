package dao

type TpOptionInfo struct {
	OptnId         string `orm:"column(optn_id);pk"`
	OptnNm         string `orm:"column(optn_nm);null"`
	AvalFlg        string `orm:"column(aval_flg);null"`
	PareClsId      string `orm:"column(pare_cls_id);null"`
	CrtTelrNo      string `orm:"column(crt_telr_no);null"`
	CrtDt          string `orm:"column(crt_dt);null"`
	CrtTm          string `orm:"column(crt_tm);null"`
	FinlModyDt     string `orm:"column(finl_modfy_dt)"`
	FinlModyTm     string `orm:"column(finl_modfy_tm)"`
	FinlModyOrgNo  string `orm:"column(finl_modfy_org_no)"`
	FinlModyTelrNo string `orm:"column(finl_modfy_telr_no)"`
	TccState       string `orm:"column(tcc_state)"`
	OptionClass    string `orm:"column(option_class)"`
}

func (t *TpOptionInfo) TableName() string {
	return "tp_option_info"
}
