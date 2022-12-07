package dao

//func init() {
//	orm.RegisterModel(new(TParamManagementDetail))
//}

type TParamManagementDetail struct {
	ParamName     string `orm:"column(param_name);size(50);pk"`
	Ccy           string `orm:"column(ccy);size(3);null"`
	CurrParam     string `orm:"column(curr_param);size(35);null"`
	NewParam      string `orm:"column(new_param);size(35);null"`
	OrgEffectDate string `orm:"column(org_effect_date);null"`
	NewEffectDate string `orm:"column(new_effect_date);null"`
	LastMaintDate string `orm:"column(last_maint_date);type(date);null"`
	LastMaintTime string `orm:"column(last_maint_time);type(time);null"`
	LastMaintBrno string `orm:"column(last_maint_brno);null"`
	LastMaintTell string `orm:"column(last_maint_tell);null"`
	TccState      int    `orm:"column(tcc_state);null" description:"tcc状态 0-Normal，1-Insert"`
}

func (t *TParamManagementDetail) TableName() string {
	return "t_param_management_detail"
}
