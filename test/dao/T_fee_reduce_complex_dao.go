package dao

type T_fee_reduce_complex struct {
	CombReduceCode  string `orm:"column(comb_reduce_code);pk" description:"组合减免代码"`
	CombReduceName  string `orm:"column(comb_reduce_name);size(120);null" description:"减免说明"`
	ReduceCode1     string `orm:"column(reduce_code1);size(5);null" description:"减免代码1"`
	ReduceRelation1 string `orm:"column(reduce_relation1);size(1);null" description:"减免关系1"`
	ReduceCode2     string `orm:"column(reduce_code2);size(5);null" description:"减免代码2 O-或;A-与"`
	ReduceRelation2 int    `orm:"column(reduce_relation2);null" description:"减免关系2"`
	ReduceCode3     string `orm:"column(reduce_code3);size(5);null" description:"减免代码3 O-或;A-与"`
	Flag1           string `orm:"column(flag1);size(2);null" description:"标志1"`
	Flag2           string `orm:"column(flag2);size(2);null" description:"标志2"`
	Remark1         string `orm:"column(remark1);size(10);null" description:"备用1"`
	Remark2         string `orm:"column(remark2);size(10);null" description:"备用2"`
	Remark3         string `orm:"column(remark3);size(10);null" description:"备用3"`
	LastMaintDate   string `orm:"column(last_maint_date);type(date);null" description:"最后更新日期"`
	LastMaintTime   string `orm:"column(last_maint_time);type(time);null" description:"最后更新时间"`
	LastMaintBrno   string `orm:"column(last_maint_brno);size(20);null" description:"最后更新机构"`
	LastMaintTell   string `orm:"column(last_maint_tell);size(20);null" description:"最后更新柜员"`
	TccState        int    `orm:"column(tcc_state);null" description:"tcc状态"`
}

func (t *T_fee_reduce_complex) TableName() string {
	return "t_fee_reduce_complex"
}
